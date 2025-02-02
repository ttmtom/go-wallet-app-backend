package service

import (
	"errors"
	"fmt"
	"go-wallet-system/wallet_system/core/model"
	"go-wallet-system/wallet_system/core/types"
)

type UserService struct {
	userRepository        types.UserRepository
	walletRepository      types.WalletRepository
	transactionRepository types.TransactionRepository
}

func NewUserService(ur types.UserRepository, wr types.WalletRepository, tr types.TransactionRepository) types.UserService {
	return &UserService{userRepository: ur, walletRepository: wr, transactionRepository: tr}
}

func (us UserService) UserRegistration(name string) error {
	user := us.userRepository.FindByID(name)
	if user != nil {
		return errors.New("user already exists")
	}
	newUser := &model.User{Name: name}
	err := us.userRepository.Create(newUser)
	if err != nil {
		return err
	}

	newWallet := &model.Wallet{Username: name, Balance: 0}
	err = us.walletRepository.Create(newWallet)
	if err != nil {
		return err
	}
	return nil
}

func (us UserService) UserInfo(name string) error {
	user := us.userRepository.FindByID(name)
	if user == nil {
		return errors.New("user not found")
	}
	wallet := us.walletRepository.FindById(name)
	histories := us.transactionRepository.GetAllByUserID(name)
	fmt.Println("`````````````````````````````")
	fmt.Println("Username: ", user.Name)
	fmt.Println("Wallet Balance: ", wallet.Balance)
	fmt.Println("Transaction history")
	for _, v := range histories {
		fmt.Println(v)
	}
	fmt.Println("`````````````````````````````")
	return nil
}
