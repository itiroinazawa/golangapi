package model

import "time"

// Transaction payment
type Transaction struct {
	ID              int       `db:"id"`
	AccountID       int       `db:"account_id" json:"account_id"`
	OperationTypeID int       `db:"operation_type_id" json:"operation_type_id"`
	Amount          float32   `db:"amount" json:"amount"`
	EventDate       time.Time `db:"event_date" json:"event_date"`
}

type TransactionRepository interface {
	CreateTransaction(accountID int, operationTypeID int, amount float32) error
	GetTransactionsByAccountID(accountID int) []Transaction
}
