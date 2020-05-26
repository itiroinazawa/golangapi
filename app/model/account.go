package model

// Account - User information
type Account struct {
	ID                   int     `db:"id" json:"id"`
	DocumentNumber       string  `db:"document_number" json:"document_number"`
	AvailableCreditLimit float32 `db:"available_credit_limit" json:"available_credit_limit"`
}

type AccountRepository interface {
	GetAccount(id int) Account
	CreateAccount(documentNumber string, availableCreditLimit float32) error
	UpdateAvailableLimit(accountID int, availableCreditLimit float32) error
}
