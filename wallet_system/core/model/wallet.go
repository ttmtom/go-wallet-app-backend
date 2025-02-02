package model

/**
* Wallet represents a user's wallet with a username and balance.
* user wallet would auto created when user register and the initial balance is set to zero.
 */

type Wallet struct {
	Username string
	Balance  float64
}

func NewWallet(username string, balance float64) *Wallet {
	return &Wallet{
		Username: username,
		Balance:  balance,
	}
}
