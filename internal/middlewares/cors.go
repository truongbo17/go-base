package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-base/config"
	"strings"
)

func Cors() gin.HandlerFunc {
	EnvConfig := config.EnvConfig
	patternAllowOrigin := EnvConfig.CorsConfig.AllowOrigin
	allowOrigin := []string{"*"}

	if patternAllowOrigin != "" {
		allowOrigin = strings.Split(EnvConfig.CorsConfig.AllowOrigin, ",")
		if allowOrigin == nil || len(allowOrigin) == 0 {
			allowOrigin = []string{"*"}
		}
	}

	return cors.New(cors.Config{
		AllowOrigins:     allowOrigin,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})
}
