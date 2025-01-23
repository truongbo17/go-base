package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-base/config"
	"go-base/internal/app/auth/model"
	"go-base/internal/app/auth/repositories"
	"go-base/internal/app/auth/services"
	"net/http"
	"strings"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader(config.HeaderAuth)
		if token == "" {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		token = strings.Replace(token, config.TokenType, "", 1)
		token = strings.TrimSpace(token)

		tokenRepository := repositories.NewTokenRepository()
		authService := services.NewAuthService(tokenRepository)
		tokenModel, err := authService.VerifyToken(token, model.TokenTypeAccess)
		if err != nil {
			_ = context.Error(err)
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		context.Set("userId", tokenModel.User)

		context.Next()
	}
}
