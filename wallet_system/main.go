package wallet_system

import (
	adapter "go-wallet-system/wallet_system/adapter/storage/types"
	"go-wallet-system/wallet_system/core/types"
)

type WalletSystem struct {
	User        types.UserController
	Wallet      types.WalletController
	Transaction types.TransactionController
}

func New(db adapter.Storage) *WalletSystem {

	return &WalletSystem{}
}
