package interfaces

import (
	"fmt"

	"github.com/itiroinazawa/golangapi/app/model"
)

type DbAccountRepo DbRepo

func NewDbAccountRepo(dbHandlers map[string]DbHandler) *DbAccountRepo {
	dbAccountRepo := new(DbAccountRepo)
	dbAccountRepo.dbHandlers = dbHandlers
	dbAccountRepo.dbHandler = dbHandlers["DbAccountRepo"]
	return dbAccountRepo
}

func (repo *DbAccountRepo) GetAccount(id int) model.Account {
	command := fmt.Sprintf("SELECT id, document_number FROM Account where id = '%d'", id)
	row := repo.dbHandler.Query(command)

	var acc model.Account

	for row.Next() {
		row.Scan(&acc.ID, &acc.DocumentNumber)
	}

	return acc
}

func (repo *DbAccountRepo) CreateAccount(documentNumber string, availableCreditLimit float32) error {

	command := fmt.Sprintf("INSERT INTO Account (document_number, available_credit_limit) VALUES ('%v', '%f')", documentNumber, availableCreditLimit)
	repo.dbHandler.Execute(command)
	return nil
}

func (repo *DbAccountRepo) UpdateAvailableLimit(accountID int, availableCreditLimit float32) error {

	command := fmt.Sprintf("UPDATE Account SET available_credit_limit='%f' WHERE id = '%v", availableCreditLimit, accountID)
	repo.dbHandler.Execute(command)
	return nil
}
