package services

import (
	"errors"
	"go-base/internal/app/auth/model"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
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
