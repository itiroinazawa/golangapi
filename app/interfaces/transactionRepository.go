package interfaces

import (
	"fmt"
	"time"

	"github.com/itiroinazawa/golangapi/model"
)

type DbTransactionRepo DbRepo

func NewDbTransactionRepo(dbHandlers map[string]DbHandler) *DbTransactionRepo {
	dbTransactionRepo := new(DbTransactionRepo)
	dbTransactionRepo.dbHandlers = dbHandlers
	dbTransactionRepo.dbHandler = dbHandlers["DbTransactionRepo"]
	return dbTransactionRepo
}

func (repo *DbTransactionRepo) GetTransactionsByAccountID(accountID int) []model.Transaction {
	command := fmt.Sprintf("SELECT id, account_id, operation_type_id, amount, event_date FROM Transaction where account_id = '%d'", accountID)
	row := repo.dbHandler.Query(command)

	var tr model.Transaction
	var transactions []model.Transaction

	for row.Next() {
		row.Scan(&tr.ID, &tr.AccountID, &tr.OperationTypeID, &tr.Amount, &tr.EventDate)
		transactions = append(transactions, tr)
	}

	return transactions
}

func (repo *DbTransactionRepo) CreateTransaction(accountID int, operationTypeID int, amount float32) error {
	command := fmt.Sprintf("INSERT INTO Transaction (account_id, operation_type_id, amount, event_date) VALUES ('%d','%d','%f','%v')", accountID, operationTypeID, amount, time.Now())
	repo.dbHandler.Execute(command)
	return nil
}
