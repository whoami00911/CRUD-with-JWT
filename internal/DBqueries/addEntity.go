package DBqueries

import (
	"database/sql"
	"fmt"
	netHttp "webPractice1/internal/transport"
	"webPractice1/pkg/logger"
)

func AddEntity(db *sql.DB, ar netHttp.AssetData) {
	logger := logger.GetLogger()
	tx, err := db.Begin()
	if err != nil {
		logger.Error(fmt.Sprintf("Truntransaction not started: %s", err))
		return
	}
	ar.IsDb = true
	_, err = tx.Exec(
		`INSERT INTO "AbuseEntity" ("ipAddress", "isPublic", "ipVersion", "isWhitelisted", "abuseConfidenceScore", "countryCode", "countryName", "usageType", "isFromDB", "isTor", "isp")
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, ar.IPAddress, ar.IsPublic, ar.IPVersion, ar.IsWhitelisted, ar.AbuseConfidenceScore,
		ar.CountryCode, ar.CountryName, ar.UsageType, ar.IsDb, ar.IsTor, ar.ISP)
	if err != nil {
		logger.Error(fmt.Sprintf("INSERT ERROR: %s", err))
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
