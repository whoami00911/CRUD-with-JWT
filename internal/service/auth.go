package service

import (
	"errors"
	"fmt"
	"time"
	"webPractice1/internal/domain"
	"webPractice1/internal/repository"
	"webPractice1/pkg/hasher"
	"webPractice1/pkg/logger"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

var ErrUserNotFound = errors.New("user not found")

type AuthService struct {
	repo   repository.Authorization
	hash   *hasher.Hash
	logger *logger.Logger
}

type jwtToken struct {
	jwt.StandardClaims
	UserId int
}

func NewAuthService(repo repository.Authorization, hash *hasher.Hash, logger *logger.Logger) *AuthService {
	return &AuthService{
		repo:   repo,
		hash:   hash,
		logger: logger,
	}
}

func (as *AuthService) CreateUser(user domain.User) int {
	user.Password = as.hash.GenHashPass(user.Password)
	return as.repo.CreateUser(user)
}

func (as *AuthService) GenToken(user, password string) (string, error) {
	//хешировать здесь
	password = as.hash.GenHashPass(password)
	id := as.repo.GetUser(user, password)
	if id == 0 {
		return "", ErrUserNotFound
	}
	ttl, err := time.ParseDuration(viper.GetString("token.token_ttl"))
	if err != nil {
		as.logger.Error(fmt.Sprintf("error time parse: %s", err))
		return "", err
	}
	//создать токен здесь
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtToken{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: id,
	})
	tokenString, err := token.SignedString([]byte(viper.GetString("token.token_key")))
	if err != nil {
		as.logger.Error(fmt.Sprintf("failed to get token signed string: %s", err))
	}
	return tokenString, nil
}

func (as *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &jwtToken{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(viper.GetString("token.token_key")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*jwtToken)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
