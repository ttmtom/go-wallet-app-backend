package repository

import (
	"errors"
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

func (t TransactionRepository) Insert(transaction *model.Transaction) error {
	if _, exists := t.store[transaction.ID]; exists {
		return errors.New("unexpected error")
	}

	t.store[transaction.ID] = transaction
	return nil
}

func (t TransactionRepository) GetAllByUserID(userID string) []*model.Transaction {
	var transactions []*model.Transaction

	for _, transaction := range t.store {
		if (transaction.From != nil && *transaction.From == userID) || (transaction.To != nil && *transaction.To == userID) {
			transactions = append(transactions, transaction)
		}
	}

	return transactions
}
