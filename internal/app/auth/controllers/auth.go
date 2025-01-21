package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-base/internal/app/auth/repositories"
	"go-base/internal/app/auth/services"
	"go-base/internal/app/auth/validators"
	"go-base/internal/response"
	"net/http"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController() *UserController {
	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	return &UserController{
		UserService: userService,
	}
}

// Register godoc
// @Summary      Register
// @Description  registers a user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        req  body      validators.RegisterRequest true "Register Request"
// @Success      200  {object}  response.BaseResponse
// @Failure      400  {object}  response.BaseResponse
// @Router       /auth/register [post]
func (userController *UserController) Register(context *gin.Context) {
	var requestBody validators.RegisterRequest
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
}

// Login godoc
// @Summary      Login
// @Description  login a user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        req  body      validators.LoginRequest true "Login Request"
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
// @Param        req  body      validators.RefreshRequest true "Refresh Request"
// @Success      200  {object}  response.BaseResponse
// @Failure      400  {object}  response.BaseResponse
// @Router       /auth/refresh [post]
func (userController *UserController) Refresh(context *gin.Context) {

}
