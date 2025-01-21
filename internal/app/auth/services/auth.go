package services

import (
	"go-base/internal/app/auth/repositories"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

// CheckExistEmail @email string
// @description Check exist email
// @return bool
func (userService *UserService) CheckExistEmail(email string) bool {
	user, err := userService.UserRepository.FindByEmail(email)
	if err != nil {
		panic("Error get email.")
	}

	return user != nil
}
