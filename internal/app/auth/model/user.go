package model

import (
	models "go-base/internal/infra/model"
)

type User struct {
	models.BasicModel

	Name     string `gorm:"size:255;not null"`
	Email    string `gorm:"size:100;unique;not null"`
	Username string `gorm:"unique;not null"`
	Password string `json:"password"`
}
