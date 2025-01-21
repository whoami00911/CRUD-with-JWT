package repository

import "database/sql"

type AuthDatabase struct {
	db *sql.DB
}

func NewAuthUserDbInicialize(db *sql.DB) *AuthDatabase {
	return &AuthDatabase{db: db}
}
