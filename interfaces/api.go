package interfaces

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/itiroinazawa/golangapi/usecases"
)

type Interactor interface {
	CreateAccount(documentNumber string) error
	CreateTransaction(accountID int, operationTypeID int, amount float32) error
	GetAccount(id int) (usecases.Account, error)
	GetTransactionsByAccountID(accountID int) ([]usecases.Transaction, error)
}

type ApiHandler struct {
	Interactor Interactor
}

func (handler ApiHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	accountID, _ := strconv.Atoi(params["accountId"])
	acc, err := handler.Interactor.GetAccount(accountID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(acc)
}

func (handler ApiHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {

	var acc usecases.Account

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &acc)

	err := handler.Interactor.CreateAccount(acc.DocumentNumber)

	if err != nil {
		fmt.Println("Error in CreateAccount: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (handler ApiHandler) GetTransactionsByAccountID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	accountID, _ := strconv.Atoi(params["accountId"])
	tr, err := handler.Interactor.GetTransactionsByAccountID(accountID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(tr)
}

func (handler ApiHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {

	var tr usecases.Transaction

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &tr)

	err := handler.Interactor.CreateTransaction(tr.AccountID, tr.OperationTypeID, tr.Amount)

	if err != nil {
		fmt.Println("Error in CreateTransaction: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
