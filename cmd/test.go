package main

import (
	"go-wallet-system/pkg/logger"
	"go-wallet-system/wallet_system"
	"go-wallet-system/wallet_system/adapter/storage/in_memory"
)

func main() {
	logger.Init()
	db := in_memory.New()
	_ := wallet_system.New(db)

	// TODO add user input and command line arguments to interact with the system
}
