package model

import (
	models "go-base/internal/infra/model"
	"time"
)

const (
	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"
)

type Token struct {
	models.BasicIDModel

	Token     string    `json:"token" gorm:"size:255;not null"`
	Type      *string   `json:"type" gorm:"size:20;null"`
	User      uint      `json:"user" gorm:"not null"`
	ExpiresAt time.Time `json:"expires_at" gorm:"autoCreateTime"`
}

func (Token) TableName() string {
	return "tokens"
}
