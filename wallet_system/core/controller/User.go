package controller

import (
	"errors"
	"go-wallet-system/pkg/utils"
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
	if vali := utils.Validate(name, `^[a-zA-Z0-9_]{3,20}$`); !vali {
		return errors.New("invalid name format")
	}

	return uc.userService.UserRegistration(name)
}

func (uc UserController) UserInfo(name string) error {
	// user input validation
	if vali := utils.Validate(name, `^[a-zA-Z0-9_]{3,20}$`); !vali {
		return errors.New("invalid name format")
	}

	return uc.userService.UserInfo(name)
}
