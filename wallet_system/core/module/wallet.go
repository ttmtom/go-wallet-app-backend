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

func NewWalletModule(db adapterTypes.Storage, tr types.TransactionRepository) *WalletModule {
	wr := db.GetWalletRepository()
	ws := service.NewWalletService(wr, tr)
	wc := controller.NewWalletController(ws)

	return &WalletModule{Repository: wr, Service: ws, Controller: wc}
}
