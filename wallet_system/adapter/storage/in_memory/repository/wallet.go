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

func (wr WalletRepository) Create(wallet *model.Wallet) error {
	if _, exists := wr.store[wallet.Username]; exists {
		return errors.New("wallet already exists")
	}
	wr.store[wallet.Username] = wallet
	return nil
}

func (wr WalletRepository) FindById(id string) *model.Wallet {
	if wallet, exists := wr.store[id]; exists {
		return wallet
	}
	return nil
}

func (wr WalletRepository) Update(wallet *model.Wallet) error {
	if _, exists := wr.store[wallet.Username]; !exists {
		return errors.New("wallet not found")
	}

	wr.store[wallet.Username] = wallet

	return nil
}
