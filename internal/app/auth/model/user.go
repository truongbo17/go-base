package model

import (
	models "go-base/internal/infra/model"
)

type User struct {
	models.BasicModel

	Name     string  `json:"name" gorm:"size:255;not null"`
	Email    string  `json:"email" gorm:"size:100;unique;not null"`
	Username string  `json:"username" gorm:"unique;not null"`
	Password *string `json:"password" gorm:"null"`
}

func (User) TableName() string {
	return "users"
}
