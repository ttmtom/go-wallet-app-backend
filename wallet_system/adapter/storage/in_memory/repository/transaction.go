package repository

import (
	"go-wallet-system/wallet_system/core/model"
	coreTypes "go-wallet-system/wallet_system/core/types"
)

type TransactionRepository struct {
	store map[string]*model.Transaction
}

func NewTransactionRepository() coreTypes.TransactionRepository {
	store := make(map[string]*model.Transaction)
	return &TransactionRepository{store: store}
}

func (t TransactionRepository) CreateTransaction(transaction *model.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (t TransactionRepository) GetTransactionsByUserID(userID string) ([]*model.Transaction, error) {
	//TODO implement me
	panic("implement me")
}
