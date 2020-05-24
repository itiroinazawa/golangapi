package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/itiroinazawa/golangapi/app/infrastructure"
	"github.com/itiroinazawa/golangapi/app/interfaces"
	"github.com/itiroinazawa/golangapi/app/usecases"
)

func main() {

	router := mux.NewRouter()

	dbHandler := infrastructure.NewMySQlHandler()

	handlers := make(map[string]interfaces.DbHandler)
	handlers["DbAccountRepo"] = dbHandler
	handlers["DbOperationTypeRepo"] = dbHandler
	handlers["DbTransactionRepo"] = dbHandler

	interactor := new(usecases.Interactor)
	interactor.AccountRepository = interfaces.NewDbAccountRepo(handlers)
	interactor.TransactionRepository = interfaces.NewDbTransactionRepo(handlers)
	//interactor.OperationTypeRepository = interfaces.NewDbOperationTypeRepo(handlers)

	apiHandler := interfaces.ApiHandler{}
	apiHandler.Interactor = interactor

	router.HandleFunc("/accounts", apiHandler.CreateAccount).Methods("POST")
	router.HandleFunc("/accounts/{accountId}", apiHandler.GetAccount)
	router.HandleFunc("/accounts/{accountId}/transactions", apiHandler.GetTransactionsByAccountID)
	router.HandleFunc("/transactions", apiHandler.CreateTransaction).Methods("POST")

	fmt.Println("Starting api-server - port: 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
