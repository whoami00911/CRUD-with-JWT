package service

import (
	"webPractice1/internal/domain"
	"webPractice1/internal/repository"
	"webPractice1/pkg/hasher"
	"webPractice1/pkg/logger"
)

type Autherization interface {
	CreateUser(user domain.User) int
	//GenHashPass(password string) string
	GenToken(user, password string) (string, error)
}

type CRUDList interface {
	AddEntity(ar domain.AssetData)
	DeleteAllEntitiesDB()
	DeleteEntityDB(ip string)
	GetEntity(ip string) *domain.AssetData
	GetEntities() []domain.AssetData
	UpdateEntity(ar domain.AssetData)
}

type Service struct {
	Autherization
	CRUDList
}

func NewService(repos *repository.Repository, hash *hasher.Hash, logger *logger.Logger) *Service {
	return &Service{
		Autherization: NewAuthService(repos.Authorization, hash, logger),
		CRUDList:      NewServiceCRUD(repos.CRUDList),
	}
}
