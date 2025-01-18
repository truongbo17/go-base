package routes

import (
	"github.com/gin-gonic/gin"
	"go-base/internal/middlewares"
)

var Router *gin.Engine

func Init() {
	Router = gin.Default()

	Router.Use(middlewares.ErrorHandle())
	Router.Use(middlewares.Cors())

	LoadPublicRouter(Router)
}
