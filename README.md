# wallet system backend

A simple centralized wallet app backend system.

## Requirements 
### Functional
#### User
1. User can create an account
2. User can view their account info (wallet balance, transaction history)

#### Wallet
1. User can deposit money into their wallet
2. User can withdraw money from their wallet
3. User can transfer money to other user's wallet

### Non-functional 
1. Data Consistency: 
   - Ensure that all transaction operations (deposits, withdrawals, transfers) are atomic.
   - Be careful on the floating point numbers when dealing with money.
2. Maintainability & Extensibility:
   - The codebase should be well-documented, modular, and easy to understand. 

## Design
The system start on `wallet_system/main.go` as the entry point.
The main components of the system are `User`, `Wallet`, `Transaction`.
Each component has its own service, controller, and repository.

### Language
Golang are chosen as the programming language for this coding test due to the following reasons:
1. Concurrency
2. Performance
3. Strong Typing

### Storage
For this coding test, in-memory storage would used to store user data and transactions.
However, an interface `Storage` were provided that can be easily replaced with a more robust database system by implementing the `Storage` interface.


### Architecture
The architecture of the system is based on the MVC (Model-View-Controller) pattern.
- **Model**: Represents the data structure of the application, such as `User`, `Wallet`, `Transaction`.
- **Service**: Contains business logic for manipulating the model. 
- **Repository**: Handles data persistence and retrieval from storage.

Client should interact with the controller, which will then call the service to perform business logic.
```go
// `wallet_system/main.go`

   type WalletSystem struct {
       User   types.UserController
       Wallet types.WalletController
   }
```
User controller provided interface for user operations, such as creating a new user or get user information
```go
// `wallet_system/core/types/user.go`
   type UserController interface {
       UserRegister(name string) error
       GetUserInfo(name string) (info *UserInfoResponse, err error)
   }
```
Wallet controller provided interface for wallet operations, deposit, withdraw and transfer
```go
// `wallet_system/core/types/wallet.go`
   type WalletController interface {
       Deposit(username string, amount string) error
       Withdraw(username string, amount string) error
       Transfer(fromUsername string, toUsername string, amount string) error
   }
```

The repository is responsible for interacting with the storage layer to persist data.

## Testing
Unit tests are written for service and controller.
```shell
  go test ./test/...
```

testing with cli

```shell
  go run cmd/test.go
```

example:
```
> go run cmd/test.go                                                                                                                                                            [2:22:55] 
-----------------------------
Go wallet backend system testing
-----------------------------
command option
1. register: user register and wallet init
2. deposit: Deposit money
3. withdraw: Withdraw money
4. transfer: Transfer money to user
5. info: Check balance
6. exit: Exit the system
-----------------------------
Enter command
> register
Command: register
Enter username
> aaa
user register done
-----------------------------
Enter command
> register
Command: register
Enter username
> bbb
user register done
-----------------------------
Enter command
> deposit
Command: deposit
Enter username
> aaa
Enter amount to deposit
> 123.123
deposit done
-----------------------------
Enter command
> transfer
Command: transfer
Enter from username
> aaa
Enter to username
> bbb
Enter amount to transfer
> 100
transaction done
-----------------------------
Enter command
> info
Command: info
Enter username
> aaa
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
Username:  aaa
Wallet Balance: 23.123000
Transaction history
Type: Deposit,  Amount: 123.123000, Timestamp: 2025-02-03 02:30:14 +0800 HKT
Type: Transfer, From: aaa, To: bbb, Amount: 100.000000, Timestamp: 2025-02-03 02:30:27 +0800 HKT
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
user info end
-----------------------------
Enter command
> info
Command: info
Enter username
> bbb
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
Username:  bbb
Wallet Balance: 100.000000
Transaction history
Type: Transfer, From: aaa, To: bbb, Amount: 100.000000, Timestamp: 2025-02-03 02:30:27 +0800 HKT
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
user info end
-----------------------------
```

## Review and Improve

TODO list
- [ ] Implement a mq system to handle concurrent transactions efficiently, on failure, transactions should be retried. 
- [ ] Implement a user authentication system to ensure that only authorized users can access their wallets. 
- [ ] Implement a transaction history feature to keep track of all transactions made by users, (includes the transaction status, timestamp, etc.). 
- [ ] improve the database object design, e.g. update the primary to hash id, the relation between tables, etc. 

## Work plan
Total time spent on this task is 3 days (Approximately 15 hours in total)

2-3 hours for understanding the requirements and designing the system architecture.

4-5 hours for coding and implementing the system.

4-5 hours for testing and debugging the system.

2 hour for documentation and final review.
