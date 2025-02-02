package model

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
