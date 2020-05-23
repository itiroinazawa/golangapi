package infrastructure

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/itiroinazawa/golangapi/app/interfaces"
)

type MySQlHandler struct {
	Conn *sql.DB
}

func (handler *MySQlHandler) Execute(statement string) {
	handler.Conn.Exec(statement)
}

func (handler *MySQlHandler) Query(statement string) interfaces.Row {
	//fmt.Println(statement)
	rows, err := handler.Conn.Query(statement)
	if err != nil {
		fmt.Println(err)
		return new(SqliteRow)
	}
	row := new(SqliteRow)
	row.Rows = rows
	return row
}

type SqliteRow struct {
	Rows *sql.Rows
}

func (r SqliteRow) Scan(dest ...interface{}) {
	r.Rows.Scan(dest...)
}

func (r SqliteRow) Next() bool {
	return r.Rows.Next()
}

func NewMySQlHandler() *MySQlHandler {
	conn, _ := sql.Open("mysql", "test:test@tcp(db:3306)/pismo_db?parseTime=true")
	mySQLHandler := new(MySQlHandler)
	mySQLHandler.Conn = conn
	return mySQLHandler
}
