package repository

import (
	"errors"
	"go-wallet-system/wallet_system/core/model"
	coreTypes "go-wallet-system/wallet_system/core/types"
)

type WalletRepository struct {
	store map[string]*model.Wallet
}

func NewWalletRepository() coreTypes.WalletRepository {
	store := make(map[string]*model.Wallet)

	return &WalletRepository{store: store}
}

func (w WalletRepository) GetBalance(userId string) (float64, error) {
	if wallet, exists := w.store[userId]; exists {
		return wallet.Balance, nil
	}
	return 0, errors.New("user not found")
}

func (w WalletRepository) SendTransaction(from, to string, amount float64) error {
	//TODO implement me
	panic("implement me")
}

func (w WalletRepository) Deposit(to string, amount float64) error {
	//TODO implement me
	panic("implement me")
}

func (w WalletRepository) Withdraw(to string, amount float64) error {
	//TODO implement me
	panic("implement me")
}
