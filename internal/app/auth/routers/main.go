package routers

import (
	"github.com/gin-gonic/gin"
	"go-base/internal/app/auth/controllers"
	"go-base/internal/app/auth/request"
	"go-base/internal/middlewares"
)

func LoadAuthModuleRouter(r *gin.Engine) *gin.RouterGroup {
	userController := controllers.NewUserController()

	groupAuth := r.Group("/auth")
	{
		groupAuth.POST(
			"/login",
			request.LoginValidator(),
			middlewares.RateLoginPublic(),
			userController.Login,
		)
		groupAuth.POST(
			"/login/google",
			request.LoginValidator(),
			middlewares.RateLoginPublic(),
			userController.LoginGoogle,
		)
		groupAuth.POST(
			"/register",
			request.RegisterValidator(),
			middlewares.RateRegisterPublic(),
			userController.Register,
		)
		groupAuth.POST(
			"/refresh",
			request.RefreshValidator(),
			middlewares.RateRefreshPublic(),
			userController.Refresh,
		)
		groupAuth.GET(
			"/me",
			middlewares.JWTMiddleware(),
			userController.Me,
		)
	}
	return groupAuth
}
