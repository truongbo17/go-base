package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
)

func ErrorHandle() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)

				debug.PrintStack()

				context.JSON(http.StatusInternalServerError, gin.H{
					"code":    http.StatusInternalServerError,
					"message": "Error",
				})

				context.Abort()
			}
		}()

		context.Next()
	}
}
