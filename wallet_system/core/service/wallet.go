package service

import (
	"github.com/shopspring/decimal"
	"go-wallet-system/wallet_system/core/model"
	"go-wallet-system/wallet_system/core/share"
	"go-wallet-system/wallet_system/core/types"
)

type WalletService struct {
	walletRepository      types.WalletRepository
	transactionRepository types.TransactionRepository
}

func NewWalletService(wr types.WalletRepository, tr types.TransactionRepository) types.WalletService {
	return &WalletService{walletRepository: wr, transactionRepository: tr}
}

func (w WalletService) Deposit(userId string, amount float64) error {
	wallet := w.walletRepository.FindById(userId)
	if wallet == nil {
		return share.WalletNotFoundError
	}

	insertAmount := decimal.NewFromFloat(amount)

	newAmount := insertAmount.Add(decimal.NewFromFloat(wallet.Balance))

	wallet.Balance = newAmount.InexactFloat64()

	err := w.walletRepository.Update(wallet)
	if err != nil {
		return err
	}

	transaction := model.NewTransaction("deposit", &userId, nil, &userId, insertAmount.InexactFloat64())

	err = w.transactionRepository.Insert(transaction)
	if err != nil {
		return err
	}

	return nil
}

func (w WalletService) Withdraw(userId string, amount float64) error {
	wallet := w.walletRepository.FindById(userId)
	if wallet == nil {
		return share.WalletNotFoundError
	}

	insertAmount := decimal.NewFromFloat(amount)
	walletAmount := decimal.NewFromFloat(wallet.Balance)

	if walletAmount.LessThan(insertAmount) {
		return share.InsufficientBalanceError
	}

	newAmount := walletAmount.Sub(insertAmount)

	wallet.Balance = newAmount.InexactFloat64()

	err := w.walletRepository.Update(wallet)
	if err != nil {
		return err
	}

	transaction := model.NewTransaction("withdraw", &userId, &userId, nil, insertAmount.InexactFloat64())

	err = w.transactionRepository.Insert(transaction)
	if err != nil {
		return err
	}

	return nil
}

func (w WalletService) Transfer(from, to string, amount float64) error {
	fromWallet := w.walletRepository.FindById(from)
	toWallet := w.walletRepository.FindById(to)
	if fromWallet == nil || toWallet == nil {
		return share.WalletNotFoundError
	}

	actionAmount := decimal.NewFromFloat(amount)
	fromWalletAmount := decimal.NewFromFloat(fromWallet.Balance)
	toWalletAmount := decimal.NewFromFloat(toWallet.Balance)

	if fromWalletAmount.LessThan(actionAmount) {
		return share.InsufficientBalanceError
	}

	newFromAmount := fromWalletAmount.Sub(actionAmount)
	newToAmount := toWalletAmount.Add(actionAmount)

	fromWallet.Balance = newFromAmount.InexactFloat64()
	toWallet.Balance = newToAmount.InexactFloat64()

	err := w.walletRepository.Update(fromWallet)
	if err != nil {
		return err
	}

	err = w.walletRepository.Update(toWallet)
	if err != nil {
		return err
	}

	transaction := model.NewTransaction("transfer", &from, &from, &to, actionAmount.InexactFloat64())
	err = w.transactionRepository.Insert(transaction)
	if err != nil {
		return err
	}
	return nil
}
