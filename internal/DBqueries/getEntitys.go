package DBqueries

import (
	"database/sql"
	"fmt"
	netHttp "webPractice1/internal/transport"
	"webPractice1/pkg/logger"
)

func GetEntity(db *sql.DB, ip string) *netHttp.AssetData {
	logger := logger.GetLogger()
	asset := netHttp.NewAsset()
	asset.Mu.Lock()
	defer asset.Mu.Unlock()
	tx, err := db.Begin()
	if err != nil {
		logger.Error(fmt.Sprintf("Truntransaction not started: %s", err))
	}
	query, err := tx.Query(fmt.Sprintf(`SELECT * FROM "AbuseEntity" WHERE "ipAddress"='%s'`, ip))
	if err != nil {
		logger.Error(fmt.Sprintf("INSERT ERROR: %s", err))
	}
	defer query.Close()
	defer func() {
		if err != nil {
			tx.Rollback()
			logger.Error(fmt.Sprintf("Something wrong with truntransaction: %s", err))
		} else {
			tx.Commit()
		}
	}()
	for query.Next() {
		if err := query.Scan(&asset.Asset.IPAddress, &asset.Asset.IsPublic, &asset.Asset.IPVersion, &asset.Asset.IsWhitelisted,
			&asset.Asset.AbuseConfidenceScore, &asset.Asset.CountryCode, &asset.Asset.CountryName, &asset.Asset.UsageType,
			&asset.Asset.ISP, &asset.Asset.IsTor, &asset.Asset.IsDb); err != nil {
			logger.Error(fmt.Sprintf("Scan method error: %s", err))
		}
	}
	return &asset.Asset
}
func GetEntitys(db *sql.DB) []netHttp.AssetData {
	logger := logger.GetLogger()
	var assets []netHttp.AssetData
	tx, err := db.Begin()
	if err != nil {
		logger.Error(fmt.Sprintf("Truntransaction not started: %s", err))
	}
	query, err := tx.Query(`SELECT * FROM "AbuseEntity"`)
	if err != nil {
		logger.Error(fmt.Sprintf("INSERT ERROR: %s", err))
	}
	defer query.Close()
	defer func() {
		if err != nil {
			tx.Rollback()
			logger.Error(fmt.Sprintf("Something wrong with truntransaction: %s", err))
		} else {
			tx.Commit()
		}
	}()
	for query.Next() {
		var asset netHttp.AssetData
		if err := query.Scan(&asset.IPAddress, &asset.IsPublic, &asset.IPVersion, &asset.IsWhitelisted,
			&asset.AbuseConfidenceScore, &asset.CountryCode, &asset.CountryName, &asset.UsageType,
			&asset.ISP, &asset.IsTor, &asset.IsDb); err != nil {
			logger.Error(fmt.Sprintf("Scan method error: %s", err))
		}
		assets = append(assets, asset)
	}
	return assets
}
