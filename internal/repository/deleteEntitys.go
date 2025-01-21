package repository

import (
	"fmt"
)

func (c *CRUD) DeleteAllEntitiesDB() {
	tx, err := c.db.Begin()
	if err != nil {
		c.logger.Error(fmt.Sprintf("Truntransaction not started: %s", err))
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			c.logger.Error(fmt.Sprintf("Something wrong with truntransaction: %s", err))
			return
		} else {
			tx.Commit()
		}
	}()
	_, err = tx.Exec(`DELETE FROM "AbuseEntity"`)
	if err != nil {
		c.logger.Error(fmt.Sprintf("DELETE IN DB ERROR: %s", err))
		return
	}
}

func (c *CRUD) DeleteEntityDB(ip string) {
	tx, err := c.db.Begin()
	if err != nil {
		c.logger.Error(fmt.Sprintf("Truntransaction not started: %s", err))
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			c.logger.Error(fmt.Sprintf("Something wrong with truntransaction: %s", err))
			return
		} else {
			tx.Commit()
		}
	}()
	_, err = tx.Exec(fmt.Sprintf(`DELETE FROM "AbuseEntity" WHERE "ipAddress" = '%s'`, ip))
	if err != nil {
		c.logger.Error(fmt.Sprintf("DELETE IN DB ERROR: %s", err))
		return
	}
	if err != nil {
		return
	}
}
