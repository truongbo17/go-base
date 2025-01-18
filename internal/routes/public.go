package routes

import (
	"github.com/gin-gonic/gin"
)

func LoadPublicRouter(r *gin.Engine) *gin.RouterGroup {
	public := r.Group("/")
	{
		public.GET("ping", func(context *gin.Context) {
			context.String(200, "pong: "+context.GetString("x-request-id"))
		})
	}
	return public
}
