package routers

import (
	"github.com/gin-gonic/gin"
	"go-base/internal/app/auth/controllers"
	"go-base/internal/app/auth/validators"
)

func LoadAuthModuleRouter(r *gin.Engine) *gin.RouterGroup {
	var userController = new(controllers.UserController)

	groupAuth := r.Group("/auth")
	{
		groupAuth.POST("/login", validators.LoginValidator(), userController.Login)
		groupAuth.POST("/register", validators.RegisterValidator(), userController.Register)
		groupAuth.POST("/refresh", validators.RefreshValidator(), userController.Refresh)
	}
	return groupAuth
}
