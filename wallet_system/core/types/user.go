package types

import "go-wallet-system/wallet_system/core/model"

//go:generate mockgen -source=user.go -destination=../../../test/mock/user.go -package=mock

type UserRepository interface {
	FindByID(id string) *model.User
	Create(user *model.User) error
}

type UserService interface {
	UserRegistration(name string) error
	UserInfo(name string) (userInfo *UserInfoResponse, err error)
}

type UserInfoResponse struct {
	Wallet               *model.Wallet
	TransactionHistories []*model.Transaction
}

type UserController interface {
	UserRegister(name string) error
	GetUserInfo(name string) (info *UserInfoResponse, err error)
}
