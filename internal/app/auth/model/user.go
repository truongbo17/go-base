package model

import (
	"github.com/golang-jwt/jwt/v4"
	models "go-base/internal/infra/model"
)

type User struct {
	models.BasicModel

	Name     string  `json:"name" gorm:"size:255;not null"`
	Email    string  `json:"email" gorm:"size:100;unique;not null"`
	Password *string `json:"password" gorm:"null"`
}

func (User) TableName() string {
	return "users"
}

type UserClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
	Type  string `json:"type"`
}
