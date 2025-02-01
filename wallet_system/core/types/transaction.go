package types

import "go-wallet-system/wallet_system/core/model"

type TransactionRepository interface {
	CreateTransaction(transaction *model.Transaction) error
	GetTransactionsByUserID(userID string) ([]*model.Transaction, error)
}

type TransactionService interface {
}

type TransactionController interface {
}
