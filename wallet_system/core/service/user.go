package service

import (
	"errors"
	"fmt"
	"go-wallet-system/wallet_system/core/model"
	"go-wallet-system/wallet_system/core/types"
)

type UserService struct {
	userRepository   types.UserRepository
	walletRepository types.WalletRepository
}

func NewUserService(ur types.UserRepository, wr types.WalletRepository) types.UserService {
	return &UserService{userRepository: ur, walletRepository: wr}
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
	fmt.Println("`````````````````````````````")
	fmt.Println("Username: ", user.Name)
	fmt.Println("Wallet Balance: ", wallet.Balance)
	fmt.Println("`````````````````````````````")
	return nil
}
