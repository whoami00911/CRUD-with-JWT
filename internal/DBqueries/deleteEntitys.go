package DBqueries

import (
	"database/sql"
	"fmt"
	"webPractice1/pkg/logger"
)

func DeleteAllEntitiesDB(db *sql.DB) {
	logger := logger.GetLogger()
	tx, err := db.Begin()
	if err != nil {
		logger.Error(fmt.Sprintf("Truntransaction not started: %s", err))
		return
	}
	_, err = tx.Exec(`DELETE FROM "AbuseEntity"`)
	if err != nil {
		logger.Error(fmt.Sprintf("DELETE IN DB ERROR: %s", err))
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback() // Rollback only if there was an error
			logger.Error(fmt.Sprintf("Something wrong with truntransaction: %s", err))
			return
		} else {
			tx.Commit() // Commit if no errors occurred
		}
	}()
}

func DeleteEntityDB(db *sql.DB, ip string) {
	logger := logger.GetLogger()
	tx, err := db.Begin()
	if err != nil {
		logger.Error(fmt.Sprintf("Truntransaction not started: %s", err))
		return
	}
	_, err = tx.Exec(fmt.Sprintf(`DELETE FROM "AbuseEntity" WHERE "ipAddress" = '%s'`, ip))
	if err != nil {
		logger.Error(fmt.Sprintf("DELETE IN DB ERROR: %s", err))
		return
	}
	if err != nil {
	}
	defer func() {
		if err != nil {
			tx.Rollback() // Rollback only if there was an error
			logger.Error(fmt.Sprintf("Something wrong with truntransaction: %s", err))
			return
		} else {
			tx.Commit() // Commit if no errors occurred
		}
	}()
}
