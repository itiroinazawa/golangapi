package usecases

import (
	"github.com/itiroinazawa/golangapi/app/model"
)

type Interactor struct {
	AccountRepository       model.AccountRepository
	TransactionRepository   model.TransactionRepository
	OperationTypeRepository model.OperationTypeRepository
}
