package controller

import (
	"go-wallet-system/wallet_system/core/share"
	"go-wallet-system/wallet_system/core/types"
)

type UserController struct {
	userService types.UserService
}

func NewUserController(userService types.UserService) types.UserController {
	return &UserController{userService: userService}
}

func (uc UserController) UserRegister(name string) error {
	// user input validation
	if vali := share.UsernameValidation(name); !vali {
		return share.InvalidNameInputError
	}

	return uc.userService.UserRegistration(name)
}

func (uc UserController) GetUserInfo(name string) (*types.UserInfoResponse, error) {
	// user input validation
	if vali := share.UsernameValidation(name); !vali {
		return nil, share.InvalidNameInputError
	}

	return uc.userService.UserInfo(name)
}
