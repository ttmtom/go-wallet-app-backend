package contoller

import (
	"errors"
	"go-wallet-system/test/mock"
	"go-wallet-system/wallet_system/core/controller"
	"go-wallet-system/wallet_system/core/share"
	"go.uber.org/mock/gomock"
	"testing"
)

type walletDepositInput struct {
	username string
	amount   string
}

type walletDepositOutput struct {
	err error
}

func TestWalletController_WalletDeposit(t *testing.T) {
	username := "testingId"
	usernameTooLongCase := "abcdefghijklmnopqrstuvwxyz"
	usernameTooShortCase := "aa"
	usernameInvalidCase := "testingId@"

	amount := "10.0"
	floatAmount := 10.0
	amountLessThanZero := "-10.0"
	amountMoreThanSixDigi := "1.1234567"
	amountNotNumeric := "abc"
	amountNotNumeric2 := "1.123.123"

	testCases := []struct {
		desc  string
		mocks func(
			walletService *mock.MockWalletService,
		)
		input    walletDepositInput
		expected walletDepositOutput
	}{
		{
			desc: "Success",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should call the user registration method with the correct username and return nil error.
				walletService.EXPECT().Deposit(gomock.Eq(username), floatAmount).Times(1).Return(nil)
			},
			input: walletDepositInput{
				username: username,
				amount:   amount,
			},
			expected: walletDepositOutput{
				err: nil,
			},
		},
		{
			desc: "InvalidNameInputError_TooLong",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletDepositInput{
				username: usernameTooLongCase,
				amount:   amount,
			},
			expected: walletDepositOutput{
				err: share.InvalidNameInputError,
			},
		},
		{
			desc: "InvalidNameInputError_TooShort",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletDepositInput{
				username: usernameTooShortCase,
				amount:   amount,
			},
			expected: walletDepositOutput{
				err: share.InvalidNameInputError,
			},
		},
		{
			desc: "InvalidNameInputError_Invalid",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletDepositInput{
				username: usernameInvalidCase,
				amount:   amount,
			},
			expected: walletDepositOutput{
				err: share.InvalidNameInputError,
			},
		},
		{
			desc: "InvalidAmountInputError_LessThanZero",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletDepositInput{
				username: username,
				amount:   amountLessThanZero,
			},
			expected: walletDepositOutput{
				err: share.InvalidAmountInputError,
			},
		},
		{
			desc: "InvalidAmountInputError_ThanSixDigi",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletDepositInput{
				username: username,
				amount:   amountMoreThanSixDigi,
			},
			expected: walletDepositOutput{
				err: share.InvalidAmountInputError,
			},
		},
		{
			desc: "InvalidAmountInputError_NotNumeric",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletDepositInput{
				username: username,
				amount:   amountNotNumeric,
			},
			expected: walletDepositOutput{
				err: share.InvalidAmountInputError,
			},
		},
		{
			desc: "InvalidAmountInputError_amountNotNumeric2",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletDepositInput{
				username: username,
				amount:   amountNotNumeric2,
			},
			expected: walletDepositOutput{
				err: share.InvalidAmountInputError,
			},
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			walletService := mock.NewMockWalletService(ctrl)

			tc.mocks(walletService)

			userController := controller.NewWalletController(walletService)

			err := userController.Deposit(tc.input.username, tc.input.amount)

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
	amount   string
}

type walletWithdrawOutput struct {
	err error
}

func TestWalletController_WalletWithdraw(t *testing.T) {
	username := "testingId"
	usernameTooLongCase := "abcdefghijklmnopqrstuvwxyz"
	usernameTooShortCase := "aa"
	usernameInvalidCase := "testingId@"

	amount := "10.0"
	floatAmount := 10.0
	amountLessThanZero := "-10.0"
	amountMoreThanSixDigi := "1.1234567"
	amountNotNumeric := "abc"
	amountNotNumeric2 := "1.123.123"

	testCases := []struct {
		desc  string
		mocks func(
			walletService *mock.MockWalletService,
		)
		input    walletWithdrawInput
		expected walletWithdrawOutput
	}{
		{
			desc: "Success",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should call the user registration method with the correct username and return nil error.
				walletService.EXPECT().Withdraw(gomock.Eq(username), floatAmount).Times(1).Return(nil)
			},
			input: walletWithdrawInput{
				username: username,
				amount:   amount,
			},
			expected: walletWithdrawOutput{
				err: nil,
			},
		},
		{
			desc: "InvalidNameInputError_TooLong",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletWithdrawInput{
				username: usernameTooLongCase,
				amount:   amount,
			},
			expected: walletWithdrawOutput{
				err: share.InvalidNameInputError,
			},
		},
		{
			desc: "InvalidNameInputError_TooShort",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletWithdrawInput{
				username: usernameTooShortCase,
				amount:   amount,
			},
			expected: walletWithdrawOutput{
				err: share.InvalidNameInputError,
			},
		},
		{
			desc: "InvalidNameInputError_Invalid",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletWithdrawInput{
				username: usernameInvalidCase,
				amount:   amount,
			},
			expected: walletWithdrawOutput{
				err: share.InvalidNameInputError,
			},
		},
		{
			desc: "InvalidAmountInputError_LessThanZero",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletWithdrawInput{
				username: username,
				amount:   amountLessThanZero,
			},
			expected: walletWithdrawOutput{
				err: share.InvalidAmountInputError,
			},
		},
		{
			desc: "InvalidAmountInputError_ThanSixDigi",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletWithdrawInput{
				username: username,
				amount:   amountMoreThanSixDigi,
			},
			expected: walletWithdrawOutput{
				err: share.InvalidAmountInputError,
			},
		},
		{
			desc: "InvalidAmountInputError_NotNumeric",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletWithdrawInput{
				username: username,
				amount:   amountNotNumeric,
			},
			expected: walletWithdrawOutput{
				err: share.InvalidAmountInputError,
			},
		},
		{
			desc: "InvalidAmountInputError_amountNotNumeric2",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletWithdrawInput{
				username: username,
				amount:   amountNotNumeric2,
			},
			expected: walletWithdrawOutput{
				err: share.InvalidAmountInputError,
			},
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			walletService := mock.NewMockWalletService(ctrl)

			tc.mocks(walletService)

			userController := controller.NewWalletController(walletService)

			err := userController.Withdraw(tc.input.username, tc.input.amount)

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
	amount       string
}

type walletTransferOutput struct {
	err error
}

func TestWalletController_WalletTransfer(t *testing.T) {
	fromUsername := "fromName"
	fromUsernameTooLongCase := "abcdefghijklmnopqrstuvwxyz"
	fromUsernameTooShortCase := "aa"
	fromUsernameInvalidCase := "testingId@"

	toUsername := "toName"
	toUsernameTooLongCase := "abcdefghijklmnopqrstuvwxyz"
	toUsernameTooShortCase := "aa"
	toUsernameInvalidCase := "testingId@"

	amount := "10.0"
	floatAmount := 10.0
	amountLessThanZero := "-10.0"
	amountMoreThanSixDigi := "1.1234567"
	amountNotNumeric := "abc"
	amountNotNumeric2 := "1.123.123"

	testCases := []struct {
		desc  string
		mocks func(
			walletService *mock.MockWalletService,
		)
		input    walletTransferInput
		expected walletTransferOutput
	}{
		{
			desc: "Success",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should call the user registration method with the correct fromUsername and return nil error.
				walletService.EXPECT().Transfer(gomock.Eq(fromUsername), gomock.Eq(toUsername), gomock.Eq(floatAmount)).Times(1).Return(nil)
			},
			input: walletTransferInput{
				fromUsername: fromUsername,
				toUsername:   toUsername,
				amount:       amount,
			},
			expected: walletTransferOutput{
				err: nil,
			},
		},
		{
			desc: "InvalidNameInputError_SameName",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletTransferInput{
				fromUsername: fromUsername,
				toUsername:   fromUsername,
				amount:       amount,
			},
			expected: walletTransferOutput{
				err: share.InvalidNameInputError,
			},
		},
		{
			desc: "InvalidNameInputError_FromNameTooLong",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletTransferInput{
				fromUsername: fromUsernameTooLongCase,
				toUsername:   toUsername,
				amount:       amount,
			},
			expected: walletTransferOutput{
				err: share.InvalidNameInputError,
			},
		},
		{
			desc: "InvalidNameInputError_FromNameTooShort",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletTransferInput{
				fromUsername: fromUsernameTooShortCase,
				toUsername:   toUsername,
				amount:       amount,
			},
			expected: walletTransferOutput{
				err: share.InvalidNameInputError,
			},
		},
		{
			desc: "InvalidNameInputError_FromNameInvalid",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletTransferInput{
				fromUsername: fromUsernameInvalidCase,
				toUsername:   toUsername,
				amount:       amount,
			},
			expected: walletTransferOutput{
				err: share.InvalidNameInputError,
			},
		},

		{
			desc: "InvalidNameInputError_ToNameTooLong",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletTransferInput{
				fromUsername: fromUsername,
				toUsername:   toUsernameTooLongCase,
				amount:       amount,
			},
			expected: walletTransferOutput{
				err: share.InvalidNameInputError,
			},
		},
		{
			desc: "InvalidNameInputError_ToNameTooShort",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletTransferInput{
				fromUsername: fromUsername,
				toUsername:   toUsernameTooShortCase,
				amount:       amount,
			},
			expected: walletTransferOutput{
				err: share.InvalidNameInputError,
			},
		},
		{
			desc: "InvalidNameInputError_ToNameInvalid",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletTransferInput{
				fromUsername: fromUsername,
				toUsername:   toUsernameInvalidCase,
				amount:       amount,
			},
			expected: walletTransferOutput{
				err: share.InvalidNameInputError,
			},
		},
		{
			desc: "InvalidAmountInputError_LessThanZero",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletTransferInput{
				fromUsername: fromUsername,
				toUsername:   toUsername,
				amount:       amountLessThanZero,
			},
			expected: walletTransferOutput{
				err: share.InvalidAmountInputError,
			},
		},
		{
			desc: "InvalidAmountInputError_ThanSixDigi",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletTransferInput{
				fromUsername: fromUsername,
				toUsername:   toUsername,
				amount:       amountMoreThanSixDigi,
			},
			expected: walletTransferOutput{
				err: share.InvalidAmountInputError,
			},
		},
		{
			desc: "InvalidAmountInputError_NotNumeric",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletTransferInput{
				fromUsername: fromUsername,
				toUsername:   toUsername,
				amount:       amountNotNumeric,
			},
			expected: walletTransferOutput{
				err: share.InvalidAmountInputError,
			},
		},
		{
			desc: "InvalidAmountInputError_amountNotNumeric2",
			mocks: func(
				walletService *mock.MockWalletService,
			) {
				// it should not call the user registration method
			},
			input: walletTransferInput{
				fromUsername: fromUsername,
				toUsername:   toUsername,
				amount:       amountNotNumeric2,
			},
			expected: walletTransferOutput{
				err: share.InvalidAmountInputError,
			},
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			walletService := mock.NewMockWalletService(ctrl)

			tc.mocks(walletService)

			userController := controller.NewWalletController(walletService)

			err := userController.Transfer(tc.input.fromUsername, tc.input.toUsername, tc.input.amount)

			if err != nil && tc.expected.err == nil ||
				err == nil && tc.expected.err != nil ||
				tc.expected.err != nil && !errors.Is(err, tc.expected.err) {
				t.Errorf("[case: %s] expected to get error %v; got %v", tc.desc, tc.expected.err, err)
			}
		})
	}
}
