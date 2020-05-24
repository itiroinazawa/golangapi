package interfaces

import (
	"fmt"

	"github.com/itiroinazawa/golangapi/app/model"
)

type DbOperationTypeRepo DbRepo

func NewDbOperationTypeRepo(dbHandlers map[string]DbHandler) *DbOperationTypeRepo {
	dbOperationTypeRepo := new(DbOperationTypeRepo)
	dbOperationTypeRepo.dbHandlers = dbHandlers
	dbOperationTypeRepo.dbHandler = dbHandlers["DbOperationTypeRepo"]
	return dbOperationTypeRepo
}

func (repo *DbOperationTypeRepo) GetOperationTypeById(operationTypeID int) model.OperationType {

	command := fmt.Sprintf("SELECT id, description, isNegativeOperator FROM OperationType where id = '%d'", operationTypeID)
	row := repo.dbHandler.Query(command)

	var ot model.OperationType

	for row.Next() {
		row.Scan(&ot.ID, &ot.Description, &ot.IsNegativeOperator)
	}

	return ot
}
