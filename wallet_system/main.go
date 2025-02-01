package wallet_system

import (
	adapterTypes "go-wallet-system/wallet_system/adapter/storage/types"
	"go-wallet-system/wallet_system/core/module"
	"go-wallet-system/wallet_system/core/types"
)

type WalletSystem struct {
	User   types.UserController
	Wallet types.WalletController
}

func New(db adapterTypes.Storage) *WalletSystem {
	//tm := module.NewTransactionModule(db)
	wm := module.NewWalletModule(db)
	um := module.NewUserModule(db, wm.Repository)

	return &WalletSystem{
		User:   um.Controller,
		Wallet: wm.Controller,
	}
}
