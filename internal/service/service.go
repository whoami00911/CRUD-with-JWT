package service

import (
	"webPractice1/internal/domain"
	"webPractice1/internal/repository"
)

type Autherization interface {
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

type Service struct {
	Autherization
	CRUDList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autherization: NewAuthService(repos.Authorization),
		CRUDList:      NewServiceCRUD(repos.CRUDList),
	}
}
