package types

import (
	coreTypes "go-wallet-system/wallet_system/core/types"
)

type Storage interface {
	GetUserRepository() coreTypes.UserRepository
	GetWalletRepository() coreTypes.WalletRepository
	GetTransactionRepository() coreTypes.TransactionRepository
}
