package in_memory

import (
	"go-wallet-system/wallet_system/adapter/storage/in_memory/repository"
	"go-wallet-system/wallet_system/adapter/storage/types"
	coreTypes "go-wallet-system/wallet_system/core/types"
)

type InMemoryStore struct {
	User        coreTypes.UserRepository
	Wallet      coreTypes.WalletRepository
	Transaction coreTypes.TransactionRepository
}

func New() types.Storage {
	userRepo := repository.NewUserRepository()
	walletRepo := repository.NewWalletRepository()
	transactionRepo := repository.NewTransactionRepository()

	return &InMemoryStore{
		User:        userRepo,
		Wallet:      walletRepo,
		Transaction: transactionRepo,
	}
}

func (s *InMemoryStore) GetUserRepository() coreTypes.UserRepository {
	return s.User
}

func (s *InMemoryStore) GetWalletRepository() coreTypes.WalletRepository {
	return s.Wallet
}

func (s *InMemoryStore) GetTransactionRepository() coreTypes.TransactionRepository {
	return s.Transaction
}
