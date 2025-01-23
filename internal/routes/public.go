package routes

import (
	"github.com/gin-gonic/gin"
	"go-base/config"
	"go-base/internal/middlewares"
)

// LoadPublicRouter sets up public routes.
//
// @Summary Ping endpoint
// @Description Responds with "pong" and the request ID.
// @Tags Public APi
// @Accept json
// @Produce plain
// @Success 200 {string} string "pong: <x-request-id>"
// @Router /ping [get]
func LoadPublicRouter(r *gin.Engine) *gin.RouterGroup {
	public := r.Group("/")
	public.Use(middlewares.RateLimitPublic())
	{
		public.GET("ping", func(context *gin.Context) {
			context.String(200, "pong: "+context.GetString(config.HeaderRequestID))
		})
	}
	return public
}
