package service

import (
	"errors"
	"go-wallet-system/test/mock"
	"go-wallet-system/wallet_system/core/model"
	"go-wallet-system/wallet_system/core/service"
	"go-wallet-system/wallet_system/core/share"
	"go-wallet-system/wallet_system/core/types"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
)

type userRegistrationInput struct {
	username string
}

type userRegistrationExpectedOutput struct {
	err error
}

func TestUserService_UserRegistration(t *testing.T) {
	username := "testingId"

	testCases := []struct {
		desc  string
		mocks func(
			userRepo *mock.MockUserRepository,
			walletRepo *mock.MockWalletRepository,
			transactionRepo *mock.MockTransactionRepository,
		)
		input    userRegistrationInput
		expected userRegistrationExpectedOutput
	}{
		{
			desc: "Success",
			mocks: func(
				userRepo *mock.MockUserRepository,
				walletRepo *mock.MockWalletRepository,
				transactionRepo *mock.MockTransactionRepository,
			) {
				userRepo.EXPECT().Create(gomock.Eq(model.NewUser(username))).Times(1).Return(nil)
				userRepo.EXPECT().FindByID(gomock.Eq(username)).Times(1).Return(nil)
				walletRepo.EXPECT().Create(gomock.Eq(model.NewWallet(username, 0))).Times(1).Return(nil)
			},
			input: userRegistrationInput{
				username: username,
			},
			expected: userRegistrationExpectedOutput{
				err: nil,
			},
		},
		{
			desc: "UserExisted",
			mocks: func(
				userRepo *mock.MockUserRepository,
				walletRepo *mock.MockWalletRepository,
				transactionRepo *mock.MockTransactionRepository,
			) {
				userRepo.EXPECT().FindByID(gomock.Eq(username)).Times(1).Return(model.NewUser(username))
			},
			input: userRegistrationInput{
				username: username,
			},
			expected: userRegistrationExpectedOutput{
				err: share.UserExistsError,
			},
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userRepo := mock.NewMockUserRepository(ctrl)
			walletRepo := mock.NewMockWalletRepository(ctrl)
			transactionRepo := mock.NewMockTransactionRepository(ctrl)

			tc.mocks(userRepo, walletRepo, transactionRepo)

			userService := service.NewUserService(userRepo, walletRepo, transactionRepo)

			err := userService.UserRegistration(username)

			if tc.expected.err != nil {
				if !errors.Is(err, tc.expected.err) {
					t.Errorf("[case: %s] expected to get error %v; got %v", tc.desc, tc.expected.err, err)
				}
				return
			}
		})
	}
}

type userInfoInput struct {
	username string
}

type userInfoExpectedOutput struct {
	userInfo *types.UserInfoResponse
	err      error
}

func TestUserService_UserInfo(t *testing.T) {
	username := "testingId"

	mockUser := model.NewUser(username)
	mockWallet := model.NewWallet(username, 1)
	mockHistories := []*model.Transaction{
		model.NewTransaction("deposit", &username, nil, &username, 1),
	}

	mockUserCase2 := model.NewUser(username)
	mockWalletCase2 := model.NewWallet(username, 0)
	var mockHistoriesCase2 []*model.Transaction

	testCases := []struct {
		desc  string
		mocks func(
			userRepo *mock.MockUserRepository,
			walletRepo *mock.MockWalletRepository,
			transactionRepo *mock.MockTransactionRepository,
		)
		input    userInfoInput
		expected userInfoExpectedOutput
	}{
		{
			desc: "Success_UserInfo_case_1",
			mocks: func(
				userRepo *mock.MockUserRepository,
				walletRepo *mock.MockWalletRepository,
				transactionRepo *mock.MockTransactionRepository,
			) {
				userRepo.EXPECT().FindByID(gomock.Eq(username)).Times(1).Return(mockUser)
				walletRepo.EXPECT().FindById(gomock.Eq(username)).Times(1).Return(mockWallet)
				transactionRepo.EXPECT().GetAllByUserID(username).Times(1).Return(mockHistories)
			},
			input: userInfoInput{
				username: username,
			},
			expected: userInfoExpectedOutput{
				userInfo: &types.UserInfoResponse{
					Wallet:               mockWallet,
					TransactionHistories: mockHistories,
				},
				err: nil,
			},
		},
		{
			desc: "Success_UserInfo_case_2",
			mocks: func(
				userRepo *mock.MockUserRepository,
				walletRepo *mock.MockWalletRepository,
				transactionRepo *mock.MockTransactionRepository,
			) {
				userRepo.EXPECT().FindByID(gomock.Eq(username)).Times(1).Return(mockUserCase2)
				walletRepo.EXPECT().FindById(gomock.Eq(username)).Times(1).Return(mockWalletCase2)
				transactionRepo.EXPECT().GetAllByUserID(username).Times(1).Return(mockHistoriesCase2)
			},
			input: userInfoInput{
				username: username,
			},
			expected: userInfoExpectedOutput{
				userInfo: &types.UserInfoResponse{
					Wallet:               mockWalletCase2,
					TransactionHistories: mockHistoriesCase2,
				},
				err: nil,
			},
		},
		{
			desc: "UserNotFoundError",
			mocks: func(
				userRepo *mock.MockUserRepository,
				walletRepo *mock.MockWalletRepository,
				transactionRepo *mock.MockTransactionRepository,
			) {
				userRepo.EXPECT().FindByID(gomock.Eq(username)).Times(1).Return(nil)
			},
			input: userInfoInput{
				username: username,
			},
			expected: userInfoExpectedOutput{
				userInfo: nil,
				err:      share.UserNotFoundError,
			},
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userRepo := mock.NewMockUserRepository(ctrl)
			walletRepo := mock.NewMockWalletRepository(ctrl)
			transactionRepo := mock.NewMockTransactionRepository(ctrl)

			tc.mocks(userRepo, walletRepo, transactionRepo)

			userService := service.NewUserService(userRepo, walletRepo, transactionRepo)

			userInfo, err := userService.UserInfo(username)

			if tc.expected.err != nil {
				if !errors.Is(err, tc.expected.err) {
					t.Errorf("[case: %s] expected to get error %v; got %v", tc.desc, tc.expected.err, err)
				}
				return
			}
			if userInfo == nil && tc.expected.userInfo != nil {
				t.Errorf("[case: %s] expected to get user info %+v; got nil", tc.desc, *tc.expected.userInfo)
			} else if userInfo != nil && tc.expected.userInfo == nil {
				t.Errorf("[case: %s] expected to get nil user info; got %+v", tc.desc, userInfo)
			} else if userInfo != nil && tc.expected.userInfo != nil {
				if userInfo.Wallet.Balance != tc.expected.userInfo.Wallet.Balance ||
					userInfo.Wallet.Username != tc.expected.userInfo.Wallet.Username {
					t.Errorf("[case: %s] expected user wallet to be %+v; got %+v", tc.desc, tc.expected.userInfo.Wallet, userInfo.Wallet)
				}
				if len(userInfo.TransactionHistories) != len(tc.expected.userInfo.TransactionHistories) {
					t.Errorf("[case: %s] expected user wallet to be %+v; got %+v", tc.desc, tc.expected.userInfo.Wallet, userInfo.Wallet)
				} else {
					for i := range userInfo.TransactionHistories {
						if !reflect.DeepEqual(userInfo.TransactionHistories[i], tc.expected.userInfo.TransactionHistories[i]) {
							t.Errorf("[case: %s] expected user transaction history to be %+v; got %+v", tc.desc, tc.expected.userInfo.TransactionHistories[i], userInfo.TransactionHistories[i])
						}
					}
				}
			}
		})
	}
}
