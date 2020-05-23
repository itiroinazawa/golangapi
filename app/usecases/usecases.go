package usecases

import (
	"github.com/itiroinazawa/golangapi/model"
)

type Interactor struct {
	AccountRepository       model.AccountRepository
	TransactionRepository   model.TransactionRepository
	OperationTypeRepository model.OperationTypeRepository
}
