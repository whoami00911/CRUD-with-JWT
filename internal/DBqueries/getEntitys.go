package DBqueries

import (
	"database/sql"
	"fmt"
	"webPractice1/internal/netHttp"
	"webPractice1/pkg/errorPrinter"
)

func GetEntity(db *sql.DB, ip string) *netHttp.AssetData {
	asset := netHttp.NewAsset()
	asset.Mu.Lock()
	defer asset.Mu.Unlock()
	tx, err := db.Begin()
	if err != nil {
		errorPrinter.PrintCallerFunctionName(err)
	}
	query, err := tx.Query(fmt.Sprintf(`SELECT * FROM "AbuseEntity" WHERE "ipAddress"='%s'`, ip))
	if err != nil {
		errorPrinter.PrintCallerFunctionName(err)
	}
	defer query.Close()
	if err != nil {
		errorPrinter.PrintCallerFunctionName(err)
	}
	defer func() {
		if err != nil {
			tx.Rollback() // Rollback only if there was an error
			errorPrinter.PrintCallerFunctionName(err)
		} else {
			tx.Commit() // Commit if no errors occurred
		}
	}()
	for query.Next() {
		if err := query.Scan(&asset.Asset.IPAddress, &asset.Asset.IsPublic, &asset.Asset.IPVersion, &asset.Asset.IsWhitelisted,
			&asset.Asset.AbuseConfidenceScore, &asset.Asset.CountryCode, &asset.Asset.CountryName, &asset.Asset.UsageType,
			&asset.Asset.ISP, &asset.Asset.IsTor, &asset.Asset.IsDb); err != nil {
			errorPrinter.PrintCallerFunctionName(err)
		}
	}
	return &asset.Asset
}
func GetEntitys(db *sql.DB) []netHttp.AssetData {
	var assets []netHttp.AssetData
	tx, err := db.Begin()
	if err != nil {
		errorPrinter.PrintCallerFunctionName(err)
	}
	query, err := tx.Query(`SELECT * FROM "AbuseEntity"`)
	if err != nil {
		errorPrinter.PrintCallerFunctionName(err)
	}
	defer query.Close()
	if err != nil {
		errorPrinter.PrintCallerFunctionName(err)
	}
	defer func() {
		if err != nil {
			tx.Rollback() // Rollback only if there was an error
			errorPrinter.PrintCallerFunctionName(err)
		} else {
			tx.Commit() // Commit if no errors occurred
		}
	}()
	for query.Next() {
		var asset netHttp.AssetData
		if err := query.Scan(&asset.IPAddress, &asset.IsPublic, &asset.IPVersion, &asset.IsWhitelisted,
			&asset.AbuseConfidenceScore, &asset.CountryCode, &asset.CountryName, &asset.UsageType,
			&asset.ISP, &asset.IsTor, &asset.IsDb); err != nil {
			errorPrinter.PrintCallerFunctionName(err)
		}
		assets = append(assets, asset)
	}
	return assets
}
