package routers

import "github.com/gin-gonic/gin"

func LoadAuthModuleRouter(r *gin.Engine) *gin.RouterGroup {
	groupAuth := r.Group("/auth")
	{
		groupAuth.POST("/login", func(context *gin.Context) {
			context.String(200, "login")
		})
		groupAuth.POST("/register", func(context *gin.Context) {
			context.String(200, "login")
		})
	}
	return groupAuth
}
