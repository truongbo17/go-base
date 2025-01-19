package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := uuid.New().String()

		context.Set("x-request-id", id)
		context.Header("x-request-id", id)

		context.Next()
	}
}
