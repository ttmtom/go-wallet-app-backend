package interfaces

import "go-wallet-system/wallet_system/core/interfaces"

type Storage interface {
	interfaces.UserRepository
	interfaces.WalletRepository
	interfaces.TransactionRepository
}
