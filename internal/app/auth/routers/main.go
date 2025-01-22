package routers

import (
	"github.com/gin-gonic/gin"
	"go-base/internal/app/auth/controllers"
	"go-base/internal/app/auth/request"
)

func LoadAuthModuleRouter(r *gin.Engine) *gin.RouterGroup {
	userController := controllers.NewUserController()

	groupAuth := r.Group("/auth")
	{
		groupAuth.POST("/login", request.LoginValidator(), userController.Login)
		groupAuth.POST("/register", request.RegisterValidator(), userController.Register)
		groupAuth.POST("/refresh", request.RefreshValidator(), userController.Refresh)
	}
	return groupAuth
}
