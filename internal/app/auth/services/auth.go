package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"go-base/config"
	"go-base/internal/app/auth/model"
	"go-base/internal/app/auth/repositories"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
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

func (authService AuthService) GenerateAccessTokens(user *model.User) (*model.Token, *model.Token, error) {
	configAuth := config.EnvConfig.AuthConfig

	accessExpiresAt := time.Now().Add(time.Duration(configAuth.JWTAccessExpirationMinutes) * time.Minute)
	refreshExpiresAt := time.Now().Add(time.Duration(configAuth.JWTRefreshExpirationDays) * time.Minute)

	accessToken, err := authService.createToken(user, model.TokenTypeAccess, accessExpiresAt)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := authService.createToken(user, model.TokenTypeRefresh, refreshExpiresAt)
	if err != nil {
		return nil, nil, err
	}

	return accessToken, refreshToken, nil
}

func (authService AuthService) createToken(user *model.User, tokenType string, expiresAt time.Time) (*model.Token, error) {
	claims := &model.UserClaims{
		Email: user.Email,
		Type:  tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			Subject:   strconv.Itoa(int(user.ID)),
		},
	}

	configAuth := config.EnvConfig.AuthConfig

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(configAuth.JWTSecretKey))
	if err != nil {
		return nil, errors.New("cannot create access token")
	}

	tokenModel := &model.Token{
		Token:     tokenString,
		Type:      &tokenType,
		ExpiresAt: expiresAt,
	}
	err = authService.TokenRepository.Create(tokenModel)
	if err != nil {
		return nil, errors.New("cannot save access token to db")
	}

	return tokenModel, nil
}

func (authService AuthService) VerifyToken(accessToken string) {
	return
}
