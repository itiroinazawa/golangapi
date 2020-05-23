package model

// Account - User information
type Account struct {
	ID             int    `db:"id" json:"id"`
	DocumentNumber string `db:"document_number" json:"document_number"`
}

type AccountRepository interface {
	GetAccount(id int) Account
	CreateAccount(documentNumber string) error
}
