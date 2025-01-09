package DBqueries

import (
	"database/sql"
	"fmt"
	netHttp "webPractice1/internal/transport"
	"webPractice1/pkg/logger"
)

func UpdateEntity(db *sql.DB, ar netHttp.AssetData) {
	logger := logger.GetLogger()
	tx, err := db.Begin()
	if err != nil {
		logger.Error(fmt.Sprintf("Truntransaction not started: %s", err))
		return
	}
	ar.IsDb = true
	_, err = tx.Exec(
		`UPDATE "AbuseEntity" SET "isPublic" = $1, "ipVersion" = $2, "isWhitelisted" = $3, "abuseConfidenceScore" = $4, "countryCode" = $5, "countryName" = $6, "usageType" = $7, "isFromDB" = $8, "isTor" = $9, "isp" = $10 WHERE "ipAddress" = $11`, ar.IsPublic, ar.IPVersion, ar.IsWhitelisted, ar.AbuseConfidenceScore,
		ar.CountryCode, ar.CountryName, ar.UsageType, ar.IsDb, ar.IsTor, ar.ISP, ar.IPAddress)
	if err != nil {
		logger.Error(fmt.Sprintf("UPDATE ERROR: %s", err))
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
