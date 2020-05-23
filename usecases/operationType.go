package usecases

// OperationType for Transaction
type OperationType struct {
	ID                 int    `json:"id"`
	Description        string `json:"description"`
	IsNegativeOperator bool   `json:"isNegativeOperator"`
}
