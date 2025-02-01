package repository

import (
	"errors"
	"go-wallet-system/wallet_system/core/model"
	coreTypes "go-wallet-system/wallet_system/core/types"
)

type UserRepository struct {
	store map[string]*model.User
}

func NewUserRepository() coreTypes.UserRepository {
	store := make(map[string]*model.User)

	return &UserRepository{store: store}
}

func (u UserRepository) FindUserByID(id string) (*model.User, error) {
	if user, exists := u.store[id]; exists {
		return user, nil
	}
	return nil, errors.New("user not found")
}

func (u UserRepository) CreateUser(user *model.User) error {
	if _, exists := u.store[user.ID]; exists {
		return errors.New("user already exists")
	}
	u.store[user.ID] = user
	return nil
}
