package module

import (
	adapterTypes "go-wallet-system/wallet_system/adapter/storage/types"
	"go-wallet-system/wallet_system/core/types"
)

type TransactionModule struct {
	Repository types.TransactionRepository
	Service    types.TransactionService
}

func NewTransactionModule(db adapterTypes.Storage) *TransactionModule {
	tr := db.GetTransactionRepository()

	return &TransactionModule{Repository: tr}
}
