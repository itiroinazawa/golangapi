package usecases

import (
	"time"
)

// Transaction payment
type Transaction struct {
	ID              int       `json:"id"`
	AccountID       int       `json:"account_id"`
	OperationTypeID int       `json:"operation_type_id"`
	Amount          float32   `json:"amount"`
	EventDate       time.Time `json:"event_date"`
}

func (interactor *Interactor) CreateTransaction(accountID int, operationTypeID int, amount float32) error {
	err := interactor.TransactionRepository.CreateTransaction(accountID, operationTypeID, amount)
	return err
}

func (interactor *Interactor) GetTransactionsByAccountID(accountID int) ([]Transaction, error) {

	mTransactions := interactor.TransactionRepository.GetTransactionsByAccountID(accountID)

	transactions := make([]Transaction, 0)

	for _, mtr := range mTransactions {
		tr := Transaction{ID: mtr.ID, AccountID: mtr.AccountID, OperationTypeID: mtr.OperationTypeID, Amount: mtr.Amount, EventDate: mtr.EventDate}
		transactions = append(transactions, tr)

	}

	return transactions, nil
}
