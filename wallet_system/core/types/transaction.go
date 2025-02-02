package types

import "go-wallet-system/wallet_system/core/model"

//go:generate mockgen -source=transaction.go -destination=../../../test/mock/transaction.go -package=mock

type TransactionRepository interface {
	Insert(transaction *model.Transaction) error
	GetAllByUserID(userID string) []*model.Transaction
}
