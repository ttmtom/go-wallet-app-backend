package controller

import (
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
		return share.InvalidNameInputError
	}
	floatAmount := share.AmountValidationAndConversation(amount)
	if floatAmount == nil {
		return share.InvalidAmountInputError
	}

	return wc.walletService.Deposit(username, *floatAmount)
}

func (wc WalletController) Withdraw(username string, amount string) error {
	if vali := share.UsernameValidation(username); !vali {
		return share.InvalidNameInputError
	}
	floatAmount := share.AmountValidationAndConversation(amount)
	if floatAmount == nil {
		return share.InvalidAmountInputError
	}

	return wc.walletService.Withdraw(username, *floatAmount)
}

func (wc WalletController) Transfer(fromUsername string, toUsername string, amount string) error {
	if vali := share.UsernameValidation(fromUsername); !vali {
		return share.InvalidNameInputError
	}
	if vali := share.UsernameValidation(toUsername); !vali {
		return share.InvalidNameInputError
	}
	floatAmount := share.AmountValidationAndConversation(amount)
	if floatAmount == nil {
		return share.InvalidAmountInputError
	}

	return wc.walletService.Transfer(fromUsername, toUsername, *floatAmount)
}
