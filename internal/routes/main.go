package routes

import (
	"github.com/gin-gonic/gin"
	"go-base/config"
	"go-base/internal/app/auth/routers"
	"go-base/internal/middlewares"
)

var Router *gin.Engine

func Init(appEnv string) {
	Router = gin.Default()

	Router.Use(middlewares.RequestID())
	Router.Use(middlewares.RequestLogger())
	Router.Use(middlewares.ErrorHandle())
	Router.Use(middlewares.Cors())

	LoadPublicRouter(Router)

	routers.LoadAuthModuleRouter(Router)

	if appEnv == config.DebugMode {
		LoadSwaggerRouter(Router)
	}
}
