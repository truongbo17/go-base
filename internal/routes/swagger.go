package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func LoadSwaggerRouter(r *gin.Engine) *gin.RouterGroup {
	swagger := r.Group("/docs")
	{
		swagger.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return swagger
}
