package repositories

import (
	"go-base/internal/app/auth/model"
	"go-base/internal/infra/database"
	"go-base/internal/infra/repository"
)

type TokenRepository struct {
	*repository.BaseRepository[model.Token]
}

func NewTokenRepository() *TokenRepository {
	db := database.DB
	return &TokenRepository{
		BaseRepository: repository.NewBaseRepository[model.Token](db),
	}
}
