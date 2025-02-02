package types

import "go-wallet-system/wallet_system/core/model"

type UserRepository interface {
	FindByID(id string) *model.User
	Create(user *model.User) error
}

type UserService interface {
	UserRegistration(name string) error
	UserInfo(name string) error
}

type UserController interface {
	UserRegister(name string) error
	GetUserInfo(name string) error
}
