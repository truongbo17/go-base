package models

import (
	"time"
)

type BasicWithDeleteModel struct {
	BasicIDModel
	BasicDateModel
	BasicCreateDateModel
	BasicUpdateDateModel
	BasicDeleteDateModel
}

type BasicModel struct {
	BasicIDModel
	BasicDateModel
	BasicCreateDateModel
	BasicUpdateDateModel
}

type BasicIDModel struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

type BasicDateModel struct {
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type BasicCreateDateModel struct {
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type BasicUpdateDateModel struct {
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type BasicDeleteDateModel struct {
	DeletedAt time.Time `json:"deleted_at"`
}
