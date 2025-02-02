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

func NewUserModule(db adapterTypes.Storage, wr types.WalletRepository, tr types.TransactionRepository) *UserModule {
	ur := db.GetUserRepository()
	us := service.NewUserService(ur, wr, tr)
	uc := controller.NewUserController(us)

	return &UserModule{
		Repository: ur,
		Service:    us,
		Controller: uc,
	}
}
