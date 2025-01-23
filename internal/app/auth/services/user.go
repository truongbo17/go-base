package services

import (
	"errors"
	"fmt"
	"go-base/internal/app/auth/model"
	"go-base/internal/app/auth/repositories"
	"go-base/internal/app/auth/request"
	"go-base/internal/infra/cache"
	"gorm.io/gorm"
	"strconv"
	"time"
)

const CacheKeyCheckExistEmail string = "CACHE_KEY_CHECK_EXIST_EMAIL"

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
	isExist, err := cache.Cache.Get(CacheKeyCheckExistEmail)
	if err != nil {
		panic(err)
	}
	if isExist == nil {
		user, err := userService.UserRepository.FindByEmail(email)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = cache.Cache.Set(CacheKeyCheckExistEmail, false, time.Hour*24*365)
			if err != nil {
				panic(err)
			}

			return false
		}
		if err != nil {
			panic(err)
		}

		err = cache.Cache.Set(CacheKeyCheckExistEmail, user != nil, time.Hour*24*365)
		if err != nil {
			panic(err)
		}

		return user != nil
	}

	isExistCheck, _ := strconv.ParseBool(fmt.Sprintf("%t", isExist))

	return isExistCheck
}

func (userService *UserService) CreateUser(userRequest request.RegisterRequest) *model.User {
	newUser := &model.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: &userRequest.Password,
	}
	err := userService.UserRepository.Create(newUser)
	if err != nil {
		panic(err)
	} else {
		return newUser
	}
}

func (userService *UserService) GetUserByEmail(email string) (*model.User, error) {
	user, err := userService.UserRepository.FindByEmail(email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return user, err
}

func (userService *UserService) GetUserById(id uint) (*model.User, error) {
	user := &model.User{}
	err := userService.UserRepository.FindByID(id, user)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return user, err
}
