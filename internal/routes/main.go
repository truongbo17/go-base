package routes

import (
	"github.com/gin-gonic/gin"
	"go-base/internal/app/auth/routers"
	"go-base/internal/middlewares"
)

var Router *gin.Engine

func Init() {
	Router = gin.Default()

	Router.Use(middlewares.RequestID())
	Router.Use(middlewares.RequestLogger())
	Router.Use(middlewares.ErrorHandle())
	Router.Use(middlewares.Cors())

	LoadPublicRouter(Router)

	routers.LoadAuthModuleRouter(Router)

	LoadSwaggerRouter(Router)
}
