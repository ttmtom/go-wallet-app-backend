package model

import (
	"fmt"
	"github.com/shopspring/decimal"
	"time"
)

/** Transfer struct,
* represents a financial transaction.
* Fields:
* - ID: unique identifier for the transaction.
* - From: pointer to the sender's account number (optional). If From is nil, the transaction is considered a deposit.
* - To: pointer to the receiver's account number (optional). If To is nil, the transaction is considered a withdrawal.
* - Amount: the amount of money transferred.
* - Timestamp: the time when the transaction occurred.
*
* This struct is used to store and manipulate transaction data in the application.
* It can be extended with additional fields or methods as needed for specific use cases.
*
* Example usage:
*   tx := Transfer{
*       ID:        1,
*       From:      "123456789",
*       To:        "987654321",
*       Amount:    100.0,
*       Timestamp: time.Now(),
 */

type Transaction struct {
	ID        string
	From      *string
	To        *string
	Amount    float64
	Timestamp int
}

func NewTransaction(actionType string, userId, from, to *string, amount float64) *Transaction {
	now := time.Now()

	return &Transaction{
		ID:        fmt.Sprintf("%s-%s-%d", actionType, userId, now.Unix()),
		From:      from,
		To:        to,
		Amount:    amount,
		Timestamp: int(now.Unix()),
	}
}

func (t *Transaction) String() string {
	actionType := "Transfer"
	if t.From == nil {
		actionType = "Deposit"
		return fmt.Sprintf("Type: %s,  Amount: %f, Timestamp: %s", actionType, decimal.NewFromFloat(t.Amount).InexactFloat64(), time.Unix(int64(t.Timestamp), 0))
	} else if t.To == nil {
		actionType = "Withdrawal"
		return fmt.Sprintf("Type: %s,  Amount: %f, Timestamp: %s", actionType, decimal.NewFromFloat(t.Amount).InexactFloat64(), time.Unix(int64(t.Timestamp), 0))
	}

	return fmt.Sprintf("Type: %s, From: %v, To: %v, Amount: %f, Timestamp: %s", actionType, *t.From, *t.To, decimal.NewFromFloat(t.Amount).InexactFloat64(), time.Unix(int64(t.Timestamp), 0))
}
