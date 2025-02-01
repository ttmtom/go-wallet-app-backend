package types

import "go-wallet-system/wallet_system/core/model"

type UserRepository interface {
	FindUserByID(id string) (*model.User, error)
	CreateUser(user *model.User) error
}

type UserService interface {
}

type UserController interface {
}
