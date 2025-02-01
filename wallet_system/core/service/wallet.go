package service

import "go-wallet-system/wallet_system/core/types"

type WalletService struct {
}

func NewWalletService() types.WalletService {
	return &WalletService{}
}

func (w WalletService) GetBalance(userId string) (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (w WalletService) SendTransaction(from, to string, amount float64) error {
	//TODO implement me
	panic("implement me")
}

func (w WalletService) Deposit(userId string, amount float64) error {
	//TODO implement me
	panic("implement me")
}

func (w WalletService) Withdraw(userId string, amount float64) error {
	//TODO implement me
	panic("implement me")
}
