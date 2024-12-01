package DBqueries

import (
	"database/sql"
	"fmt"
	"webPractice1/pkg/errorPrinter"
)

func DeleteAllEntitiesDB(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		errorPrinter.PrintCallerFunctionName(err)
		return
	}
	_, err = tx.Exec(`DELETE FROM "AbuseEntity"`)
	if err != nil {
		errorPrinter.PrintCallerFunctionName(err)
		return
	}
	if err != nil {
		errorPrinter.PrintCallerFunctionName(err)
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback() // Rollback only if there was an error
			errorPrinter.PrintCallerFunctionName(err)
			return
		} else {
			tx.Commit() // Commit if no errors occurred
		}
	}()
}

func DeleteEntityDB(db *sql.DB, ip string) {
	tx, err := db.Begin()
	if err != nil {
		errorPrinter.PrintCallerFunctionName(err)
		return
	}
	_, err = tx.Exec(fmt.Sprintf(`DELETE FROM "AbuseEntity" WHERE "ipAddress" = '%s'`, ip))
	if err != nil {
		errorPrinter.PrintCallerFunctionName(err)
		return
	}
	if err != nil {
	}
	defer func() {
		if err != nil {
			tx.Rollback() // Rollback only if there was an error
			errorPrinter.PrintCallerFunctionName(err)
			return
		} else {
			tx.Commit() // Commit if no errors occurred
		}
	}()
}
