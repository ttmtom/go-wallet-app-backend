package repository

import (
	"go-wallet-system/wallet_system/core/model"
	coreTypes "go-wallet-system/wallet_system/core/types"
)

type TransactionRepository struct {
}

func NewTransactionRepository() coreTypes.TransactionRepository {
	return &TransactionRepository{}
}

func (t TransactionRepository) CreateTransaction(transaction *model.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (t TransactionRepository) GetTransactionsByUserID(userID string) ([]*model.Transaction, error) {
	//TODO implement me
	panic("implement me")
}
