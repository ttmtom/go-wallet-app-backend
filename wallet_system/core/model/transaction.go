package model

/** Transaction struct,
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
*   tx := Transaction{
*       ID:        1,
*       From:      "123456789",
*       To:        "987654321",
*       Amount:    100.0,
*       Timestamp: time.Now().Format("2006-01-02T15:04:05Z"),
 */

type Transaction struct {
	ID        int
	From      *string
	To        *string
	Amount    float64
	Timestamp string
}
