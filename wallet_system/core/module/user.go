package module

import (
	adapterTypes "go-wallet-system/wallet_system/adapter/storage/types"
	"go-wallet-system/wallet_system/core/controller"
	"go-wallet-system/wallet_system/core/service"
	"go-wallet-system/wallet_system/core/types"
)

type UserModule struct {
	Repository types.UserRepository
	Service    types.UserService
	Controller types.UserController
}

func NewUserModule(db adapterTypes.Storage, wm types.WalletRepository) *UserModule {
	ur := db.GetUserRepository()
	us := service.NewUserService(ur, wm)
	uc := controller.NewUserController(us)

	return &UserModule{
		Repository: ur,
		Service:    us,
		Controller: uc,
	}
}
