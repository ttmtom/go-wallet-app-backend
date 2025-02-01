package types

type WalletRepository interface {
	GetBalance(userId string) (float64, error)
	SendTransaction(from, to string, amount float64) error
	Deposit(userId string, amount float64) error
	Withdraw(userId string, amount float64) error
}

type WalletService struct {
}

type WalletController struct {
}
