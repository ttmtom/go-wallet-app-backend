package types

import "go-wallet-system/wallet_system/core/model"

type WalletRepository interface {
	Create(wallet *model.Wallet) error
	FindById(id string) *model.Wallet
	Update(wallet *model.Wallet) error
}

type WalletService interface {
	Deposit(userId string, amount float64) error
	Withdraw(userId string, amount float64) error
	Transfer(from, to string, amount float64) error
}

type WalletController interface {
	Deposit(username string, amount string) error
	Withdraw(username string, amount string) error
	Transfer(fromUsername string, toUsername string, amount string) error
}
