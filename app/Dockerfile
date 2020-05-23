FROM golang:latest

RUN mkdir /app 
ADD . /app/
WORKDIR /app 

ENV GOPATH=/app

RUN apt-get update
RUN apt-get install vim -y
RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/gorilla/mux"
RUN go get "github.com/itiroinazawa/golangapi"

RUN go build main.go

ENTRYPOINT ["./main"]

EXPOSE 8000