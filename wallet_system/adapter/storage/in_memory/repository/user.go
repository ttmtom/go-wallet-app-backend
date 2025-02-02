package repository

import (
	"go-wallet-system/wallet_system/core/model"
	"go-wallet-system/wallet_system/core/share"
	coreTypes "go-wallet-system/wallet_system/core/types"
)

type UserRepository struct {
	store map[string]*model.User
}

func NewUserRepository() coreTypes.UserRepository {
	store := make(map[string]*model.User)

	return &UserRepository{store: store}
}

func (ur UserRepository) FindByID(id string) *model.User {
	if user, exists := ur.store[id]; exists {
		return user
	}
	return nil
}

func (ur UserRepository) Create(user *model.User) error {
	if _, exists := ur.store[user.Name]; exists {
		return share.UserExistsError
	}
	ur.store[user.Name] = user
	return nil
}
