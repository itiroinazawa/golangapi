package model

// OperationType for Transaction
type OperationType struct {
	ID                 int    `db:"id" json:"id"`
	Description        string `db:"description" json:"description"`
	IsNegativeOperator bool   `db:"isNegativeOperator" json:"isNegativeOperator"`
}

type OperationTypeRepository interface {
	GetOperationTypeById(id int) error
}
