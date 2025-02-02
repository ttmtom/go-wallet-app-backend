package service

import (
	"errors"
	"github.com/shopspring/decimal"
	"go-wallet-system/test/mock"
	"go-wallet-system/wallet_system/core/model"
	"go-wallet-system/wallet_system/core/service"
	"go-wallet-system/wallet_system/core/share"
	"go.uber.org/mock/gomock"
	"testing"
)

type walletDepositInput struct {
	username string
	amount   float64
}

type walletDepositOutput struct {
	err error
}

func TestWalletService_Deposit(t *testing.T) {
	username := "testingId"

	walletBalance := decimal.NewFromFloat(123.123)
	depositAmount := decimal.NewFromFloat(10.0)
	newAmount := walletBalance.Add(depositAmount)

	mockWallet := model.NewWallet(username, walletBalance.InexactFloat64())
	mockUpdatedWallet := model.NewWallet(username, newAmount.InexactFloat64())
	mockTransaction := model.NewTransaction("deposit", &username, nil, &username, depositAmount.InexactFloat64())

	testCases := []struct {
		desc  string
		mocks func(
			walletRepo *mock.MockWalletRepository,
			transactionRepo *mock.MockTransactionRepository,
		)
		input    walletDepositInput
		expected walletDepositOutput
	}{
		{
			desc: "Success",
			mocks: func(
				walletRepo *mock.MockWalletRepository,
				transactionRepo *mock.MockTransactionRepository,
			) {
				// search for the existing wallet
				walletRepo.EXPECT().FindById(gomock.Eq(username)).Times(1).Return(mockWallet)
				// update the wallet balance and save it
				walletRepo.EXPECT().Update(gomock.Eq(mockUpdatedWallet)).Times(1).Return(nil)
				// insert the transaction histories
				transactionRepo.EXPECT().Insert(gomock.Eq(mockTransaction)).Times(1).Return(nil)
			},
			input: walletDepositInput{
				username: username,
			},
			expected: walletDepositOutput{
				err: nil,
			},
		},
		{
			desc: "WalletNotFoundError",
			mocks: func(
				walletRepo *mock.MockWalletRepository,
				transactionRepo *mock.MockTransactionRepository,
			) {
				// search for the existing wallet and the wallet not found
				walletRepo.EXPECT().FindById(gomock.Eq(username)).Times(1).Return(nil)
			},
			input: walletDepositInput{
				username: username,
			},
			expected: walletDepositOutput{
				err: share.WalletNotFoundError,
			},
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			walletRepo := mock.NewMockWalletRepository(ctrl)
			transactionRepo := mock.NewMockTransactionRepository(ctrl)

			tc.mocks(walletRepo, transactionRepo)

			walletService := service.NewWalletService(walletRepo, transactionRepo)

			err := walletService.Deposit(username, depositAmount.InexactFloat64())

			if err != nil && tc.expected.err == nil ||
				err == nil && tc.expected.err != nil ||
				tc.expected.err != nil && !errors.Is(err, tc.expected.err) {
				t.Errorf("[case: %s] expected to get error %v; got %v", tc.desc, tc.expected.err, err)
			}
		})
	}
}

type walletWithdrawInput struct {
	username string
	amount   float64
}

type walletWithdrawOutput struct {
	err error
}

func TestWalletService_Withdrawal(t *testing.T) {
	username := "testingId"

	walletBalance := decimal.NewFromFloat(123.123)
	withdrawAmount := decimal.NewFromFloat(10.0)
	newAmount := walletBalance.Sub(withdrawAmount)

	mockWallet := model.NewWallet(username, walletBalance.InexactFloat64())
	mockUpdatedWallet := model.NewWallet(username, newAmount.InexactFloat64())
	mockTransaction := model.NewTransaction("withdraw", &username, &username, nil, withdrawAmount.InexactFloat64())

	mockInsufficientWallet := model.NewWallet(username, 0)

	testCases := []struct {
		desc  string
		mocks func(
			walletRepo *mock.MockWalletRepository,
			transactionRepo *mock.MockTransactionRepository,
		)
		input    walletWithdrawInput
		expected walletWithdrawOutput
	}{
		{
			desc: "Success",
			mocks: func(
				walletRepo *mock.MockWalletRepository,
				transactionRepo *mock.MockTransactionRepository,
			) {
				// search for the existing wallet
				walletRepo.EXPECT().FindById(gomock.Eq(username)).Times(1).Return(mockWallet)
				// update the wallet balance and save it
				walletRepo.EXPECT().Update(gomock.Eq(mockUpdatedWallet)).Times(1).Return(nil)
				// insert the transaction histories
				transactionRepo.EXPECT().Insert(gomock.Eq(mockTransaction)).Times(1).Return(nil)
			},
			input: walletWithdrawInput{
				username: username,
			},
			expected: walletWithdrawOutput{
				err: nil,
			},
		},
		{
			desc: "WalletNotFoundError",
			mocks: func(
				walletRepo *mock.MockWalletRepository,
				transactionRepo *mock.MockTransactionRepository,
			) {
				// search for the existing wallet and the wallet not found
				walletRepo.EXPECT().FindById(gomock.Eq(username)).Times(1).Return(nil)
			},
			input: walletWithdrawInput{
				username: username,
			},
			expected: walletWithdrawOutput{
				err: share.WalletNotFoundError,
			},
		},
		{
			desc: "InsufficientBalance",
			mocks: func(
				walletRepo *mock.MockWalletRepository,
				transactionRepo *mock.MockTransactionRepository,
			) {
				// search for the existing wallet and return an insufficient balance wallet
				walletRepo.EXPECT().FindById(gomock.Eq(username)).Times(1).Return(mockInsufficientWallet)
			},
			input: walletWithdrawInput{
				username: username,
			},
			expected: walletWithdrawOutput{
				err: share.InsufficientBalanceError,
			},
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			walletRepo := mock.NewMockWalletRepository(ctrl)
			transactionRepo := mock.NewMockTransactionRepository(ctrl)

			tc.mocks(walletRepo, transactionRepo)

			walletService := service.NewWalletService(walletRepo, transactionRepo)

			err := walletService.Withdraw(username, withdrawAmount.InexactFloat64())

			if err != nil && tc.expected.err == nil ||
				err == nil && tc.expected.err != nil ||
				tc.expected.err != nil && !errors.Is(err, tc.expected.err) {
				t.Errorf("[case: %s] expected to get error %v; got %v", tc.desc, tc.expected.err, err)
			}
		})
	}
}

type walletTransferInput struct {
	fromUsername string
	toUsername   string
	amount       float64
}

type walletTransferOutput struct {
	err error
}

func TestWalletService_Transfer(t *testing.T) {
	transferAmount := decimal.NewFromFloat(23.123)

	fromUsername := "fromUsername"
	fromWallBalance := decimal.NewFromFloat(123.123)

	fromWallet := model.NewWallet(fromUsername, fromWallBalance.InexactFloat64())

	toUsername := "toUsername"
	toWallBalance := decimal.NewFromFloat(0.99)

	toWallet := model.NewWallet(toUsername, toWallBalance.InexactFloat64())

	fromUpdatedWallet := model.NewWallet(fromUsername, fromWallBalance.Sub(transferAmount).InexactFloat64())
	toUpdatedWallet := model.NewWallet(toUsername, toWallBalance.Add(transferAmount).InexactFloat64())

	successTransaction := model.NewTransaction("transfer", &fromUsername, &fromUsername, &toUsername, transferAmount.InexactFloat64())

	mockInsufficientWallet := model.NewWallet(fromUsername, 0)

	testCases := []struct {
		desc  string
		mocks func(
			walletRepo *mock.MockWalletRepository,
			transactionRepo *mock.MockTransactionRepository,
		)
		input    walletTransferInput
		expected walletTransferOutput
	}{
		{
			desc: "Success",
			mocks: func(
				walletRepo *mock.MockWalletRepository,
				transactionRepo *mock.MockTransactionRepository,
			) {
				// search for the existing wallet
				walletRepo.EXPECT().FindById(gomock.Eq(fromUsername)).Times(1).Return(fromWallet)
				walletRepo.EXPECT().FindById(gomock.Eq(toUsername)).Times(1).Return(toWallet)

				// update the wallet balance and save it
				walletRepo.EXPECT().Update(gomock.Eq(fromUpdatedWallet)).Times(1).Return(nil)
				walletRepo.EXPECT().Update(gomock.Eq(toUpdatedWallet)).Times(1).Return(nil)

				// insert the transaction histories
				transactionRepo.EXPECT().Insert(gomock.Eq(successTransaction)).Times(1).Return(nil)
			},
			input: walletTransferInput{
				fromUsername: fromUsername,
				toUsername:   toUsername,
				amount:       transferAmount.InexactFloat64(),
			},
			expected: walletTransferOutput{
				err: nil,
			},
		},
		{
			desc: "WalletNotFoundError_FromWalletNotfound",
			mocks: func(
				walletRepo *mock.MockWalletRepository,
				transactionRepo *mock.MockTransactionRepository,
			) {
				// search for the existing wallet and the wallet not found
				walletRepo.EXPECT().FindById(gomock.Eq(fromUsername)).Times(1).Return(nil)
				walletRepo.EXPECT().FindById(gomock.Eq(toUsername)).Times(1).Return(toWallet)

			},
			input: walletTransferInput{
				fromUsername: fromUsername,
				toUsername:   toUsername,
				amount:       transferAmount.InexactFloat64(),
			},
			expected: walletTransferOutput{
				err: share.WalletNotFoundError,
			},
		},
		{
			desc: "WalletNotFoundError_ToWalletNotfound",
			mocks: func(
				walletRepo *mock.MockWalletRepository,
				transactionRepo *mock.MockTransactionRepository,
			) {
				// search for the existing wallet and the wallet not found
				walletRepo.EXPECT().FindById(gomock.Eq(fromUsername)).Times(1).Return(fromWallet)
				walletRepo.EXPECT().FindById(gomock.Eq(toUsername)).Times(1).Return(nil)

			},
			input: walletTransferInput{
				fromUsername: fromUsername,
				toUsername:   toUsername,
				amount:       transferAmount.InexactFloat64(),
			},
			expected: walletTransferOutput{
				err: share.WalletNotFoundError,
			},
		},
		{
			desc: "InsufficientBalance",
			mocks: func(
				walletRepo *mock.MockWalletRepository,
				transactionRepo *mock.MockTransactionRepository,
			) {
				// search for the existing wallet and the wallet not found
				walletRepo.EXPECT().FindById(gomock.Eq(fromUsername)).Times(1).Return(mockInsufficientWallet)
				walletRepo.EXPECT().FindById(gomock.Eq(toUsername)).Times(1).Return(toWallet)

			},
			input: walletTransferInput{
				fromUsername: fromUsername,
				toUsername:   toUsername,
				amount:       transferAmount.InexactFloat64(),
			},
			expected: walletTransferOutput{
				err: share.InsufficientBalanceError,
			},
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			walletRepo := mock.NewMockWalletRepository(ctrl)
			transactionRepo := mock.NewMockTransactionRepository(ctrl)

			tc.mocks(walletRepo, transactionRepo)

			walletService := service.NewWalletService(walletRepo, transactionRepo)

			err := walletService.Transfer(fromUsername, toUsername, transferAmount.InexactFloat64())

			if err != nil && tc.expected.err == nil ||
				err == nil && tc.expected.err != nil ||
				tc.expected.err != nil && !errors.Is(err, tc.expected.err) {
				t.Errorf("[case: %s] expected to get error %v; got %v", tc.desc, tc.expected.err, err)
			}
		})
	}
}
