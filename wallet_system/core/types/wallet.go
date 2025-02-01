package types

import "go-wallet-system/wallet_system/core/model"

type WalletRepository interface {
	Create(wallet *model.Wallet) error
	FindById(id string) *model.Wallet
}

type WalletService interface {
	GetBalance(userId string) (float64, error)
	SendTransaction(from, to string, amount float64) error
	Deposit(userId string, amount float64) error
	Withdraw(userId string, amount float64) error
}

type WalletController interface {
}
