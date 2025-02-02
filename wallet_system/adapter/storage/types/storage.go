package types

import (
	coreTypes "go-wallet-system/wallet_system/core/types"
)

/**
* Storage Interface defines the methods for accessing different repositories.
* This allows for a separation of concerns and makes it easier to test and maintain the code.
* Also it provide a way to switch between different storage implementations without changing the rest of the application.
 */

type Storage interface {
	GetUserRepository() coreTypes.UserRepository
	GetWalletRepository() coreTypes.WalletRepository
	GetTransactionRepository() coreTypes.TransactionRepository
}
