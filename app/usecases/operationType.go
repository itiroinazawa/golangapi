package usecases

import "fmt"

// OperationType for Transaction
type OperationType struct {
	ID                 int    `json:"id"`
	Description        string `json:"description"`
	IsNegativeOperator bool   `json:"isNegativeOperator"`
}

func (interactor *Interactor) GetOperationTypeByID(operationTypeID int) (OperationType, error) {

	mOt := interactor.OperationTypeRepository.GetOperationTypeById(operationTypeID)
	ot := OperationType{ID: mOt.ID, Description: mOt.Description, IsNegativeOperator: mOt.IsNegativeOperator}

	if ot.ID == 0 {
		err := fmt.Errorf("Operation Type not found")
		return ot, err
	}

	return ot, nil
}
