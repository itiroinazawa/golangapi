# Go API using Clean Architecture

Golang API created using Clean Architecture to describe a simple functionality of a user payment transaction

### Prerequisites

Docker 
Docker Compose
Go

## Getting Started

To run the application run the following command

```
docker-compose up --build
```

This will create an docker image of the application, a MySQL instance to run the application and create the necessary tables

## Executing the API

After the application is up and running, you can use it using postman in the following endpoints:

(POST)  http://localhost:8000/accounts
(GET)   http://localhost:8000/accounts/{id}
(POST)  http://localhost:8000/transactions
(GET)   http://localhost:8000/accounts/{id}/transactions

## Running the tests

To run the tests, execute the following command on your terminal:

```
go test ./... 
```

This will run all existing tests recursively from the folder and all subfolders
Only one test file were create for example: ./app/interfaces/api_test.go