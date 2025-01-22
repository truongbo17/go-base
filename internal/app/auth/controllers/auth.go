package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/copier"
	"go-base/internal/app/auth/repositories"
	"go-base/internal/app/auth/request"
	responseAuth "go-base/internal/app/auth/response"
	"go-base/internal/app/auth/services"
	"go-base/internal/response"
	"net/http"
)

type UserController struct {
	UserService *services.UserService
	AuthService *services.AuthService
}

func NewUserController() *UserController {
	userRepository := repositories.NewUserRepository()
	tokenRepository := repositories.NewTokenRepository()
	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService(tokenRepository)
	return &UserController{
		UserService: userService,
		AuthService: authService,
	}
}

// Register godoc
// @Summary      Register
// @Description  registers a user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        req  body      request.RegisterRequest true "Register Request"
// @Success      200  {object}  response.UserRegisterResponse
// @Failure      400  {object}  response.BaseResponse
// @Router       /auth/register [post]
func (userController *UserController) Register(context *gin.Context) {
	var requestBody request.RegisterRequest
	_ = context.ShouldBindBodyWith(&requestBody, binding.JSON)

	isExistEmail := userController.UserService.CheckExistEmail(requestBody.Email)
	if isExistEmail {
		context.JSON(http.StatusOK, response.BaseResponse{
			Status:     false,
			StatusCode: 1000,
			RequestId:  context.GetString("x-request-id"),
			Data:       nil,
			Message:    "Email already exist",
			Error:      nil,
		})
		return
	}

	decryptPassword, err := userController.AuthService.GeneratePassword(requestBody.Password)
	if err != nil {
		panic(err)
	}
	requestBody.Password = string(decryptPassword)

	user := userController.UserService.CreateUser(requestBody)

	accessToken, refreshToken, err := userController.AuthService.GenerateAccessTokens(user)
	if err != nil {
		panic(err)
	}

	userInfo := responseAuth.UserInfo{}
	_ = copier.Copy(&userInfo, &user)

	context.JSON(http.StatusOK, response.BaseResponse{
		Status:     true,
		StatusCode: http.StatusOK,
		RequestId:  context.GetString("x-request-id"),
		Data: responseAuth.UserRegisterResponse{
			Token: responseAuth.Token{
				AccessToken:  accessToken.Token,
				RefreshToken: refreshToken.Token,
			},
			User: userInfo,
		},
		Message: "Success",
		Error:   nil,
	})
}

// Login godoc
// @Summary      Login
// @Description  login a user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        req  body      request.LoginRequest true "Login Request"
// @Success      200  {object}  response.BaseResponse
// @Failure      400  {object}  response.BaseResponse
// @Router       /auth/login [post]
func (userController *UserController) Login(context *gin.Context) {
}

// Refresh godoc
// @Summary      Refresh
// @Description  refreshes a user token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        req  body      request.RefreshRequest true "Refresh Request"
// @Success      200  {object}  response.BaseResponse
// @Failure      400  {object}  response.BaseResponse
// @Router       /auth/refresh [post]
func (userController *UserController) Refresh(context *gin.Context) {

}
