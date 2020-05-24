package usecases

import (
	"fmt"
	"time"

	"github.com/itiroinazawa/golangapi/app/model"
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
	ot := interactor.OperationTypeRepository.GetOperationTypeById(operationTypeID)
	parsedAmount := parseAmountByOperationType(ot, amount)
	fmt.Println("Amount2:", amount)
	err := interactor.TransactionRepository.CreateTransaction(accountID, operationTypeID, parsedAmount)
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

func parseAmountByOperationType(ot model.OperationType, amount float32) float32 {

	if ot.ID == 0 {
		fmt.Println("Not valid operation type")
	} else if ot.IsNegativeOperator {
		amount = amount * float32(-1)
	}

	fmt.Println("Amount1:", amount)
	return amount

}
