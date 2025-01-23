package routes

import (
	"github.com/gin-gonic/gin"
	"go-base/internal/infra/swagger"
)

func LoadSwaggerRouter(r *gin.Engine) *gin.RouterGroup {
	swaggerGroup := r.Group("/docs")
	{
		swaggerGroup.GET("/swagger/*any", swagger.CustomSwaggerWrapHandler())
	}
	return swaggerGroup
}
