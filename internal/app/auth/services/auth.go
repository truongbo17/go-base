package services

import (
	"go-base/internal/app/auth/repositories"
)

type UserService struct{}

//var userRepository = repositories.NewUserRepository()

// CheckExistEmail @email string
// @description Check exist email
// @return bool
func (userService *UserService) CheckExistEmail(email string) bool {
	var userRepository = repositories.NewUserRepository()
	user, err := userRepository.FindByEmail(email)
	if err != nil {
		panic("Error get email.")
	}

	return user != nil
}
