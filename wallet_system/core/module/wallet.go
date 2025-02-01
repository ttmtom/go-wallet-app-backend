package module

import (
	adapterTypes "go-wallet-system/wallet_system/adapter/storage/types"
	"go-wallet-system/wallet_system/core/controller"
	"go-wallet-system/wallet_system/core/service"
	"go-wallet-system/wallet_system/core/types"
)

type WalletModule struct {
	Repository types.WalletRepository
	Service    types.WalletService
	Controller types.WalletController
}

func NewWalletModule(db adapterTypes.Storage) *WalletModule {
	wr := db.GetWalletRepository()
	ws := service.NewWalletService()
	wc := controller.NewWallController()

	return &WalletModule{Repository: wr, Service: ws, Controller: wc}
}
