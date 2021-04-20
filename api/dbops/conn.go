package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:xiaxichenspassword19950203.@tcp(localhost:3306)/videos_server?utf8mb4")
	if err != nil {
		panic(err)
	}
}