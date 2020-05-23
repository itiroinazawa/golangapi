package interfaces

import (
	"fmt"

	"github.com/itiroinazawa/golangapi/model"
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

func (repo *DbAccountRepo) CreateAccount(documentNumber string) error {

	command := fmt.Sprintf("INSERT INTO Account (document_number) VALUES ('%v')", documentNumber)
	repo.dbHandler.Execute(command)
	return nil
}
