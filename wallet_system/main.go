package wallet_system

import (
	adapter "go-wallet-system/wallet_system/adapter/interfaces"
	"go-wallet-system/wallet_system/core/interfaces"
)

type WalletSystem struct {
	User        interfaces.UserController
	Wallet      interfaces.WalletController
	Transaction interfaces.TransactionController
}

func New(db adapter.Storage) *WalletSystem {

	return &WalletSystem{}
}
