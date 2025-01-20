package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-base/internal/infra/cache"
	"time"
)

type UserController struct{}

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
	err := cache.Cache.Set("key1", "value11111", 10*time.Second)
	if err != nil {
		fmt.Println("Set error:", err)
	}

	value, err := cache.Cache.Get("key112")
	if err != nil {
		fmt.Println("Get error:", err)
	} else {
		fmt.Println("Get value:", value)
	}

	context.String(200, "OK")
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
