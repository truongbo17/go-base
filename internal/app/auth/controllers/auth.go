package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/copier"
	"go-base/internal/app/auth/model"
	"go-base/internal/app/auth/repositories"
	"go-base/internal/app/auth/request"
	responseAuth "go-base/internal/app/auth/response"
	"go-base/internal/app/auth/services"
	"go-base/internal/response"
	"golang.org/x/crypto/bcrypt"
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

// Me godoc
// @Summary      Me
// @Description  Get me
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.UserInfo
// @Failure      400  {object}  response.BaseResponse
// @Router       /auth/me [get]
// @Security	 Authorization
func (userController *UserController) Me(context *gin.Context) {

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
		Message: "Register user successfully.",
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
	var requestBody request.LoginRequest
	_ = context.ShouldBindBodyWith(&requestBody, binding.JSON)

	user, err := userController.UserService.GetUserByEmail(requestBody.Email)
	if err != nil {
		panic(err)
	}

	res := &response.BaseResponse{
		Status:     false,
		StatusCode: 1001,
		RequestId:  context.GetString("x-request-id"),
		Data:       nil,
		Message:    "Email or password is wrong",
		Error:      nil,
	}
	if user == nil {
		context.JSON(http.StatusOK, res)
		return
	}

	password := *user.Password
	if user.Password == nil {
		password = ""
	}
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(requestBody.Password))
	if err != nil {
		res.StatusCode = 1002
		context.JSON(http.StatusOK, res)
		return
	}

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
		Data: responseAuth.UserLoginResponse{
			Token: responseAuth.Token{
				AccessToken:  accessToken.Token,
				RefreshToken: refreshToken.Token,
			},
			User: userInfo,
		},
		Message: "Login user successfully.",
		Error:   nil,
	})
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
	var requestBody request.RefreshRequest
	_ = context.ShouldBindBodyWith(&requestBody, binding.JSON)

	tokenModel, err := userController.AuthService.VerifyToken(requestBody.Token, model.TokenTypeRefresh)
	res := &response.BaseResponse{
		Status:     false,
		StatusCode: 1003,
		RequestId:  context.GetString("x-request-id"),
		Data:       nil,
		Message:    "",
		Error:      nil,
	}
	if err != nil {
		res.Message = err.Error()
		context.JSON(http.StatusOK, res)
		return
	}

	userController.AuthService.RevokeTokenByUser(tokenModel.User)

	user, err := userController.UserService.GetUserById(tokenModel.User)
	if err != nil {
		res.StatusCode = 1004
		res.Message = err.Error()
		context.JSON(http.StatusOK, res)
		return
	}
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
		Data: responseAuth.Token{
			AccessToken:  accessToken.Token,
			RefreshToken: refreshToken.Token,
		},
		Message: "Refresh token successfully.",
		Error:   nil,
	})
}
