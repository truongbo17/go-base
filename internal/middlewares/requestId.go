package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-base/config"
)

func RequestID() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := uuid.New().String()

		context.Set(config.HeaderRequestID, id)
		context.Header(config.HeaderRequestID, id)

		context.Next()
	}
}
