package contoller

import (
	"errors"
	"go-wallet-system/test/mock"
	"go-wallet-system/wallet_system/core/controller"
	"go-wallet-system/wallet_system/core/model"
	"go-wallet-system/wallet_system/core/share"
	"go-wallet-system/wallet_system/core/types"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
)

type userRegisterInput struct {
	username string
}

type userRegisterOutput struct {
	err error
}

func TestUserController_UserRegister(t *testing.T) {
	username := "testingId"
	usernameTooLongCase := "abcdefghijklmnopqrstuvwxyz"
	usernameTooShortCase := "aa"
	usernameInvalidCase := "testingId@"

	testCases := []struct {
		desc  string
		mocks func(
			userService *mock.MockUserService,
		)
		input    userRegisterInput
		expected userRegisterOutput
	}{
		{
			desc: "Success",
			mocks: func(
				userService *mock.MockUserService,
			) {
				// it should call the service with the correct username and return nil error.
				userService.EXPECT().UserRegistration(gomock.Eq(username)).Times(1).Return(nil)
			},
			input: userRegisterInput{
				username: username,
			},
			expected: userRegisterOutput{
				err: nil,
			},
		},
		{
			desc: "InvalidNameInputError_TooLong",
			mocks: func(
				userService *mock.MockUserService,
			) {
				// it should not call the service
			},
			input: userRegisterInput{
				username: usernameTooLongCase,
			},
			expected: userRegisterOutput{
				err: share.InvalidNameInputError,
			},
		},
		{
			desc: "InvalidNameInputError_TooShort",
			mocks: func(
				userService *mock.MockUserService,
			) {
				// it should not call the service
			},
			input: userRegisterInput{
				username: usernameTooShortCase,
			},
			expected: userRegisterOutput{
				err: share.InvalidNameInputError,
			},
		},
		{
			desc: "InvalidNameInputError_Invalid",
			mocks: func(
				userService *mock.MockUserService,
			) {
				// it should not call the service
			},
			input: userRegisterInput{
				username: usernameInvalidCase,
			},
			expected: userRegisterOutput{
				err: share.InvalidNameInputError,
			},
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userService := mock.NewMockUserService(ctrl)

			tc.mocks(userService)

			userController := controller.NewUserController(userService)

			err := userController.UserRegister(tc.input.username)

			if err != nil && tc.expected.err == nil ||
				err == nil && tc.expected.err != nil ||
				tc.expected.err != nil && !errors.Is(err, tc.expected.err) {
				t.Errorf("[case: %s] expected to get error %v; got %v", tc.desc, tc.expected.err, err)
			}
		})
	}
}

type getUserInfoInput struct {
	username string
}

type getUserInfoOutput struct {
	userInfo *types.UserInfoResponse
	err      error
}

func TestUserController_GetUserInfo(t *testing.T) {
	username := "testingId"
	usernameTooLongCase := "abcdefghijklmnopqrstuvwxyz"
	usernameTooShortCase := "aa"
	usernameInvalidCase := "testingId@"

	mockUserInfo := &types.UserInfoResponse{
		Wallet:               model.NewWallet(username, 0),
		TransactionHistories: []*model.Transaction{},
	}

	testCases := []struct {
		desc  string
		mocks func(
			userService *mock.MockUserService,
		)
		input    getUserInfoInput
		expected getUserInfoOutput
	}{
		{
			desc: "Success",
			mocks: func(
				userService *mock.MockUserService,
			) {
				// it should call the user info method with the correct username and return mock user info.
				userService.EXPECT().UserInfo(gomock.Eq(username)).Times(1).Return(mockUserInfo, nil)
			},
			input: getUserInfoInput{
				username: username,
			},
			expected: getUserInfoOutput{
				userInfo: mockUserInfo,
				err:      nil,
			},
		},
		{
			desc: "InvalidNameInputError_TooLong",
			mocks: func(
				userService *mock.MockUserService,
			) {
				// it should not call the service
			},
			input: getUserInfoInput{
				username: usernameTooLongCase,
			},
			expected: getUserInfoOutput{
				userInfo: nil,
				err:      share.InvalidNameInputError,
			},
		},
		{
			desc: "InvalidNameInputError_TooShort",
			mocks: func(
				userService *mock.MockUserService,
			) {
				// it should not call the service
			},
			input: getUserInfoInput{
				username: usernameTooShortCase,
			},
			expected: getUserInfoOutput{
				userInfo: nil,
				err:      share.InvalidNameInputError,
			},
		},
		{
			desc: "InvalidNameInputError_Invalid",
			mocks: func(
				userService *mock.MockUserService,
			) {
				// it should not call the service
			},
			input: getUserInfoInput{
				username: usernameInvalidCase,
			},
			expected: getUserInfoOutput{
				userInfo: nil,
				err:      share.InvalidNameInputError,
			},
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userService := mock.NewMockUserService(ctrl)

			tc.mocks(userService)

			userController := controller.NewUserController(userService)

			userInfo, err := userController.GetUserInfo(tc.input.username)

			if err != nil && tc.expected.err == nil ||
				err == nil && tc.expected.err != nil ||
				tc.expected.err != nil && !errors.Is(err, tc.expected.err) {
				t.Errorf("[case: %s] expected to get error %v; got %v", tc.desc, tc.expected.err, err)
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
