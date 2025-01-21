package service

import (
	"webPractice1/internal/domain"
	"webPractice1/internal/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (as *AuthService) CreateUser(user domain.User) (int, error) {
	return as.repo.CreateUser(user)
}
