package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-base/internal/app/auth/model"
	"go-base/internal/app/auth/repositories"
	"go-base/internal/app/auth/services"
	"net/http"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenRepository := repositories.NewTokenRepository()
		authService := services.NewAuthService(tokenRepository)

		token := context.GetHeader("Bearer-Token")
		tokenModel, err := authService.VerifyToken(token, model.TokenTypeAccess)
		if err != nil {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		context.Set("userId", tokenModel.User)

		context.Next()
	}
}
