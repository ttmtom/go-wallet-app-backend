package main

import (
	"fmt"
	"go-wallet-system/pkg/logger"
	"go-wallet-system/wallet_system"
	"go-wallet-system/wallet_system/adapter/storage/in_memory"
	"os"
)

func main() {
	logger.Init()
	db := in_memory.New()
	ws := wallet_system.New(db)

	quit := make(chan int)

	go func() {
		run(ws, quit)
	}()

	<-quit
	fmt.Println("Good bye")
}

func run(ws *wallet_system.WalletSystem, quit chan int) {
	fmt.Println("-----------------------------")
	fmt.Println("Go wallet backend system testing")
	fmt.Println("-----------------------------")
	fmt.Println("command option")
	fmt.Println("1. register: user Register and init wallet")
	fmt.Println("2. deposit: Deposit money")
	fmt.Println("3. withdraw: Withdraw money")
	fmt.Println("4. transaction: Withdraw money")
	fmt.Println("5. info: Check balance")
	fmt.Println("6. exit: end the system")
	fmt.Println("-----------------------------")

	for {
		var cmd string
		fmt.Println("Enter command")
		getUserInput(&cmd)
		fmt.Println("Command:", cmd)
		switch cmd {
		case "register":
			registerHandler(ws)
		case "deposit":
			// Implement deposit logic here
		case "withdraw":
			// Implement withdrawal logic here
		case "transaction":
			// Implement transaction logic here
		case "info":
			userInfoHandler(ws)
		case "exit":
			quit <- 1
			return
		default:
			fmt.Println("Invalid command. Please try again.")
		}
		fmt.Println("-----------------------------")
	}
}

func getUserInput(p *string) {
	_, err := fmt.Scan(p)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
}

func registerHandler(ws *wallet_system.WalletSystem) {
	fmt.Println("Enter username")
	var username string
	getUserInput(&username)

	err := ws.User.UserRegister(username)
	if err != nil {
		fmt.Println("error on user register: ", err.Error())
	} else {
		fmt.Println("user register done")
	}
}

func userInfoHandler(ws *wallet_system.WalletSystem) {
	fmt.Println("Enter username")
	var username string
	getUserInput(&username)
	err := ws.User.UserInfo(username)
	if err != nil {
		fmt.Println("error on get user info: ", err.Error())
	} else {
		fmt.Println("user info end")
	}
}
