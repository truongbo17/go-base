package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	cacheManager "go-base/internal/infra/cache"
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
	cache := cacheManager.Cache
	err := cache.Set("test123132", "ok", 20)
	value, err1 := cache.Get("test123132", "12")
	fmt.Println(err, err1, value)
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
