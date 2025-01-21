package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-base/internal/app/auth/repositories"
	"go-base/internal/app/auth/services"
	"go-base/internal/app/auth/validators"
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
// @Param        req  body      models.RegisterRequest true "Register Request"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /auth/register [post]
func (userController *UserController) Register(context *gin.Context) {
	var requestBody validators.RegisterRequest
	_ = context.ShouldBindBodyWith(&requestBody, binding.JSON)

	fmt.Println(requestBody)

	isExistEmail := userController.UserService.CheckExistEmail(requestBody.Email)
	fmt.Println(isExistEmail)
}

// Login godoc
// @Summary      Login
// @Description  login a user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        req  body      models.LoginRequest true "Login Request"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /auth/login [post]
func (userController *UserController) Login(context *gin.Context) {
}

// Refresh godoc
// @Summary      Refresh
// @Description  refreshes a user token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        req  body      models.RefreshRequest true "Refresh Request"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /auth/refresh [post]
func (userController *UserController) Refresh(context *gin.Context) {

}
