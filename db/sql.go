package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // postgres golang driver
)

var PgCon *sql.DB

//CreatePgConnection does connects to Postgres
func CreatePgConnection() *sql.DB {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	PgCon = db

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to PostGres .....")
	// return the connection
	return db
}

//LogAndQuery does print query and executes
func LogAndQuery(db *sql.DB, query string, args ...interface{}) *sql.Rows {
	fmt.Println(query)
	res, err := db.Query(query, args...)
	if err != nil {
		panic(err)
	}
	return res
}

