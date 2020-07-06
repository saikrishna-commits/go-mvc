package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//SqlCon does
var SqlCon *sql.DB

//CreateSqlConnection does
func CreateSqlConnection() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/atom")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	SqlCon = db
}
