package in_memory

import "go-wallet-system/wallet_system/adapter/interfaces"

type InMemoryStore struct {
}

func New() interfaces.Storage {
	return &InMemoryStore{}
}
