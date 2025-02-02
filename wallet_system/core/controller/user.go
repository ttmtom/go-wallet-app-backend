package controller

import (
	"errors"
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
		return errors.New("invalid name format")
	}

	return uc.userService.UserRegistration(name)
}

func (uc UserController) GetUserInfo(name string) (*types.UserInfo, error) {
	// user input validation
	if vali := share.UsernameValidation(name); !vali {
		return nil, errors.New("invalid name format")
	}

	return uc.userService.UserInfo(name)
}
