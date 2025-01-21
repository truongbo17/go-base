package repositories

import (
	"go-base/internal/app/auth/model"
	"go-base/internal/infra/database"
	"go-base/internal/infra/repository"
)

var db = database.DB

type UserRepository struct {
	*repository.BaseRepository[model.User]
}

func NewUserRepository() *UserRepository {
	db := database.DB
	return &UserRepository{
		BaseRepository: repository.NewBaseRepository[model.User](db),
	}
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
