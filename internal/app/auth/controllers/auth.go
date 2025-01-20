package controllers

import "github.com/gin-gonic/gin"

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

}

func (userController *UserController) Login(context *gin.Context) {

}

func (userController *UserController) Refresh(context *gin.Context) {

}
