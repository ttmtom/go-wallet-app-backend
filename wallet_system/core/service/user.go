package service

import (
	"go-wallet-system/wallet_system/core/model"
	"go-wallet-system/wallet_system/core/share"
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
		return share.UserExistsError
	}

	newUser := model.NewUser(name)
	err := us.userRepository.Create(newUser)
	if err != nil {
		return err
	}

	newWallet := model.NewWallet(name, 0)
	err = us.walletRepository.Create(newWallet)
	if err != nil {
		return err
	}

	return nil
}

func (us UserService) UserInfo(name string) (*types.UserInfoResponse, error) {
	user := us.userRepository.FindByID(name)
	if user == nil {
		return nil, share.UserNotFoundError
	}
	wallet := us.walletRepository.FindById(name)
	histories := us.transactionRepository.GetAllByUserID(name)
	userInfo := &types.UserInfoResponse{
		Wallet:               wallet,
		TransactionHistories: histories,
	}
	return userInfo, nil
}
