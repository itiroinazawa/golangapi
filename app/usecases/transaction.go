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

	acc := interactor.AccountRepository.GetAccount(accountID)

	if ot.IsNegativeOperator {
		if !isAccountLimitAvailable(acc.AvailableCreditLimit, amount) {
			return fmt.Errorf("Limite não disponivel")
		}
	}

	err := interactor.TransactionRepository.CreateTransaction(accountID, operationTypeID, parsedAmount)

	if err == nil {

		newLimit := acc.AvailableCreditLimit - amount
		interactor.AccountRepository.UpdateAvailableLimit(accountID, newLimit)
	}

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

func isAccountLimitAvailable(availableCreditLimit float32, amount float32) bool {

	if availableCreditLimit >= amount {
		return true
	}

	return false
}

func parseAmountByOperationType(ot model.OperationType, amount float32) float32 {

	if ot.ID == 0 {
		fmt.Println("Not valid operation type")
	} else if ot.IsNegativeOperator {
		amount = amount * float32(-1)
	}

	return amount
}
