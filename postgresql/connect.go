package connect

import (
	"database/sql"
	"fmt"
	"webPractice1/cmd/errorPrinter"

	_ "github.com/lib/pq"
)

func PostgresqlConnect() *sql.DB {
	const (
		host     = "127.0.0.1"
		port     = 5433
		user     = "postgres"
		password = "12345"
		dbname   = "webPractice1"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		errorPrinter.PrintCallerFunctionName(err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		errorPrinter.PrintCallerFunctionName(err)
	}

	fmt.Println("Successfully connected!")
	return db
}
