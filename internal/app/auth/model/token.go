package model

import (
	models "go-base/internal/infra/model"
	"time"
)

type Token struct {
	models.BasicModel

	Token     string    `json:"token" gorm:"size:255;not null"`
	Type      *string   `json:"type" gorm:"size:20;null"`
	ExpiresAt time.Time `json:"expires_at" gorm:"autoCreateTime"`
}

func (Token) TableName() string {
	return "tokens"
}
