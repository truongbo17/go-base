package routes

import (
	"github.com/gin-gonic/gin"
	"go-base/internal/app/auth/routers"
	"go-base/internal/middlewares"
	"go-base/internal/response"
	"net/http"
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

func handleRouterStatusError() {
	Router.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, response.BaseResponse{
			Status:     false,
			StatusCode: http.StatusNotFound,
			RequestId:  context.GetString("x-request-id"),
			Data:       nil,
			Message:    "404 Not Found",
			Error:      nil,
		})
	})

	Router.NoMethod(func(context *gin.Context) {
		context.JSON(http.StatusMethodNotAllowed, response.BaseResponse{
			Status:     false,
			StatusCode: http.StatusMethodNotAllowed,
			RequestId:  context.GetString("x-request-id"),
			Data:       nil,
			Message:    "405 Method Not Allowed",
			Error:      nil,
		})
	})
}
