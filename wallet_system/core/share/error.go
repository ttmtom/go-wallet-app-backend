package share

import "errors"

var UserExistsError = errors.New("user already exists")

var UserNotFoundError = errors.New("user not found")

var InvalidNameInputError = errors.New("invalid name format")

var InvalidAmountInputError = errors.New("invalid amount format")

var WalletNotFoundError = errors.New("wallet not found")

var WalletAlreadyExistsError = errors.New("wallet already exists")

var InsufficientBalanceError = errors.New("insufficient amount")

var UnexpectedError = errors.New("unexpected error")
