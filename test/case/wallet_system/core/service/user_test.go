package service

import (
	"errors"
	"go-wallet-system/test/mock"
	"go-wallet-system/wallet_system/core/model"
	"go-wallet-system/wallet_system/core/service"
	"go.uber.org/mock/gomock"
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
			transaction *mock.MockTransactionRepository,
		)
		input    userRegistrationInput
		expected userRegistrationExpectedOutput
	}{
		{
			desc: "Success_UserRegistration",
			mocks: func(
				userRepo *mock.MockUserRepository,
				walletRepo *mock.MockWalletRepository,
				transaction *mock.MockTransactionRepository,
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

			if errors.Is(err, tc.expected.err) {
				return
			} else {
				t.Errorf("[case: %s] expected to get error %v; got %v", tc.desc, tc.expected.err, err)
			}
		})
	}
}
