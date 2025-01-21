package repository

import (
	"database/sql"
	"webPractice1/pkg/logger"
)

type CRUD struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewCrudDbInicialize(db *sql.DB, log *logger.Logger) *CRUD {
	return &CRUD{
		db:     db,
		logger: log,
	}

}
