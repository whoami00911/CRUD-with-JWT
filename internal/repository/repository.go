package repository

import (
	"database/sql"
	"webPractice1/internal/domain"
	"webPractice1/pkg/logger"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
}

type CRUDList interface {
	AddEntity(ar domain.AssetData)
	DeleteAllEntitiesDB()
	DeleteEntityDB(ip string)
	GetEntity(ip string) *domain.AssetData
	GetEntities() []domain.AssetData
	UpdateEntity(ar domain.AssetData)
}

type Repository struct {
	Authorization
	CRUDList
}

func NewRepository(db *sql.DB, log *logger.Logger) *Repository {
	return &Repository{
		Authorization: NewAuthUserDbInicialize(db),
		CRUDList:      NewCrudDbInicialize(db, log),
	}
}
