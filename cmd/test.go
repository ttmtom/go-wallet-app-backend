package main

import (
	"fmt"
	"github.com/shopspring/decimal"

	"go-wallet-system/pkg/utils"
	"go-wallet-system/wallet_system"
	"go-wallet-system/wallet_system/adapter/storage/in_memory"
	"go-wallet-system/wallet_system/adapter/storage/types"
)

func main() {
	db := in_memory.New()

	quit := make(chan int)

	go run(db, quit)

	<-quit
	fmt.Println("~ Good bye ~")
}

func run(db types.Storage, quit chan int) {
	ws := wallet_system.New(db)

	fmt.Println("-----------------------------")
	fmt.Println("Go wallet backend system testing")
	fmt.Println("-----------------------------")
	fmt.Println("command option")
	fmt.Println("1. register: user register and wallet init")
	fmt.Println("2. deposit: Deposit money")
	fmt.Println("3. withdraw: Withdraw money")
	fmt.Println("4. transfer: Transfer money to user")
	fmt.Println("5. info: Check balance")
	fmt.Println("6. exit: Exit the system")
	fmt.Println("-----------------------------")

	for {
		var cmd string
		fmt.Println("Enter command")
		utils.GetUserInput(&cmd)
		fmt.Println("Command:", cmd)
		switch cmd {
		case "register":
			registerHandler(ws)
		case "deposit":
			depositHandler(ws)
		case "withdraw":
			withdrawHandler(ws)
		case "transfer":
			transferHandler(ws)
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

func registerHandler(ws *wallet_system.WalletSystem) {
	fmt.Println("Enter username")
	var username string
	utils.GetUserInput(&username)

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
	utils.GetUserInput(&username)

	userInfo, err := ws.User.GetUserInfo(username)

	if err != nil {
		fmt.Println("error on get user info: ", err.Error())
	} else {

		fmt.Println("`````````````````````````````")
		fmt.Println("Username: ", userInfo.Wallet.Username)
		fmt.Println(fmt.Sprintf("Wallet Balance: %f", decimal.NewFromFloat(userInfo.Wallet.Balance).InexactFloat64()))
		fmt.Println("Transaction history")
		for _, v := range userInfo.TransactionHistories {
			fmt.Println(v)
		}
		fmt.Println("`````````````````````````````")
		fmt.Println("user info end")
	}
}

func depositHandler(ws *wallet_system.WalletSystem) {
	fmt.Println("Enter username")
	var username string
	utils.GetUserInput(&username)
	fmt.Println("Enter amount to deposit")
	var amount string
	utils.GetUserInput(&amount)

	err := ws.Wallet.Deposit(username, amount)

	if err != nil {
		fmt.Println("error on deposit: ", err.Error())
	} else {
		fmt.Println("deposit done")
	}
}

func withdrawHandler(ws *wallet_system.WalletSystem) {
	fmt.Println("Enter username")
	var username string
	utils.GetUserInput(&username)
	fmt.Println("Enter amount to withdraw")
	var amount string
	utils.GetUserInput(&amount)

	err := ws.Wallet.Withdraw(username, amount)

	if err != nil {
		fmt.Println("error on withdraw: ", err.Error())
	} else {
		fmt.Println("withdrawal done")
	}
}

func transferHandler(ws *wallet_system.WalletSystem) {
	fmt.Println("Enter from username")
	var from string
	utils.GetUserInput(&from)
	fmt.Println("Enter to username")
	var to string
	utils.GetUserInput(&to)
	fmt.Println("Enter amount to transfer")
	var amount string
	utils.GetUserInput(&amount)

	err := ws.Wallet.Transfer(from, to, amount)

	if err != nil {
		fmt.Println("error on transfer: ", err.Error())
	} else {
		fmt.Println("transaction done")
	}
}
