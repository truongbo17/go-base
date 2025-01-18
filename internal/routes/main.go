package routes

import (
	"github.com/gin-gonic/gin"
	"go-base/internal/middlewares"
)

var Router *gin.Engine

func Init() {
	Router = gin.Default()

	Router.Use(middlewares.RequestID())
	Router.Use(middlewares.ErrorHandle())
	Router.Use(middlewares.Cors())
	Router.Use(middlewares.RequestLogger())

	LoadPublicRouter(Router)
}
