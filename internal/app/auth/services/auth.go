package services

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"go-base/config"
	"go-base/internal/app/auth/model"
	"go-base/internal/app/auth/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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
	refreshExpiresAt := time.Now().Add(time.Duration(configAuth.JWTRefreshExpirationDays) * time.Hour * 24)

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
			Subject:   fmt.Sprintf("%v", user.ID),
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
		User:      user.ID,
		ExpiresAt: expiresAt,
	}
	err = authService.TokenRepository.Create(tokenModel)
	if err != nil {
		return nil, errors.New("cannot save access token to db")
	}

	return tokenModel, nil
}

func (authService AuthService) VerifyToken(token string, tokenType string) (*model.Token, error) {
	claims := &model.UserClaims{}
	configAuth := config.EnvConfig.AuthConfig

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(configAuth.JWTSecretKey), nil
	})
	if err != nil || claims.Type != tokenType {
		return nil, errors.New("not valid token parse")
	}

	if time.Now().Sub(claims.ExpiresAt.Time) > 10*time.Second {
		return nil, errors.New("token is expired")
	}

	userId, _ := primitive.ObjectIDFromHex(claims.Subject)
	condition := map[string]interface{}{
		"user": userId,
	}
	tokenModel, err := authService.TokenRepository.FindOneByCondition(condition)
	if err != nil || tokenModel == nil {
		return nil, errors.New("not valid token")
	}

	return tokenModel, nil
}

func (authService AuthService) DeleteTokenByUser(user uint) error {
	condition := map[string]interface{}{
		"user": user,
	}
	return authService.TokenRepository.DeleteByCondition(condition)
}

func (authService AuthService) RevokeTokenByUser(user uint) {
	condition := map[string]interface{}{
		"user": user,
	}
	update := map[string]interface{}{
		"expires_at": time.Now().Add(-3 * time.Minute),
	}
	err := authService.TokenRepository.UpdateByCondition(condition, update)
	if err != nil {
		panic(err)
	}
}
