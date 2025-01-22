package services

import (
	"errors"
	"go-base/internal/app/auth/model"
	"go-base/internal/app/auth/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	TokenRepository *repositories.TokenRepository
}

func NewAuthService(tokenRepository *repositories.TokenRepository) *AuthService {
	return &AuthService{
		TokenRepository: tokenRepository,
	}
}

func (authService AuthService) GeneratePassword(plainPassword string) ([]byte, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("cannot generate hashed password")
	}

	return password, nil
}

func (authService AuthService) GenerateAccessTokens(user *model.User) (string, string, error) {
	return "", "", nil
}
