package usecases

import "fmt"

// Account - User information
type Account struct {
	ID             int    `json:"id"`
	DocumentNumber string `json:"document_number"`
}

func (interactor *Interactor) CreateAccount(documentNumber string) error {

	err := interactor.AccountRepository.CreateAccount(documentNumber)
	return err
}

func (interactor *Interactor) GetAccount(id int) (Account, error) {
	var acc Account

	mAcc := interactor.AccountRepository.GetAccount(id)
	acc = Account{ID: mAcc.ID, DocumentNumber: mAcc.DocumentNumber}

	if acc.ID == 0 {
		err := fmt.Errorf("Account not found")
		return acc, err
	}

	return acc, nil
}
