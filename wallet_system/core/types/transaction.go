package types

import "go-wallet-system/wallet_system/core/model"

type TransactionRepository interface {
	Insert(transaction *model.Transaction) error
	GetAllByUserID(userID string) []*model.Transaction
}
