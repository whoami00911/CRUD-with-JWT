package DBqueries

import (
	"database/sql"
	"webPractice1/internal/netHttp"
	"webPractice1/pkg/errorPrinter"
)

func AddEntity(db *sql.DB, ar netHttp.AssetData) {
	tx, err := db.Begin()
	if err != nil {
		errorPrinter.PrintCallerFunctionName(err)
		return
	}
	ar.IsDb = true
	_, err = tx.Exec(
		`INSERT INTO "AbuseEntity" ("ipAddress", "isPublic", "iPVersion", "isWhitelisted", "abuseConfidenceScore", "countryCode", "countryName", "usageType", "isFromDB", "isTor", "isp")
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, ar.IPAddress, ar.IsPublic, ar.IPVersion, ar.IsWhitelisted, ar.AbuseConfidenceScore,
		ar.CountryCode, ar.CountryName, ar.UsageType, ar.IsDb, ar.IsTor, ar.ISP)
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
