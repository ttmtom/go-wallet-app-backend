package controller

import (
	"errors"
	"go-wallet-system/wallet_system/core/share"
	"go-wallet-system/wallet_system/core/types"
)

type WalletController struct {
	walletService types.WalletService
}

func NewWalletController(ws types.WalletService) types.WalletController {
	return &WalletController{walletService: ws}
}

func (wc WalletController) Deposit(username string, amount string) error {
	if vali := share.UsernameValidation(username); !vali {
		return errors.New("invalid name format")
	}
	floatAmount := share.AmountValidationAndConversation(amount)
	if floatAmount == nil {
		return errors.New("invalid amount format")
	}

	return wc.walletService.Deposit(username, *floatAmount)
}

func (wc WalletController) Withdraw(username string, amount string) error {
	if vali := share.UsernameValidation(username); !vali {
		return errors.New("invalid name format")
	}
	floatAmount := share.AmountValidationAndConversation(amount)
	if floatAmount == nil {
		return errors.New("invalid amount format")
	}

	return wc.walletService.Withdraw(username, *floatAmount)
}

func (wc WalletController) Transfer(fromUsername string, toUsername string, amount string) error {
	if vali := share.UsernameValidation(fromUsername); !vali {
		return errors.New("invalid name format")
	}
	if vali := share.UsernameValidation(toUsername); !vali {
		return errors.New("invalid name format")
	}
	floatAmount := share.AmountValidationAndConversation(amount)
	if floatAmount == nil {
		return errors.New("invalid amount format")
	}

	return wc.walletService.Transfer(fromUsername, toUsername, *floatAmount)
}
