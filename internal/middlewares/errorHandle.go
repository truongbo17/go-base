package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-base/internal/response"
	"log"
	"net/http"
)

func ErrorHandle() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Printf("Total Errors -> %d", len(context.Errors))

		requestId := context.GetString("x-request-id")

		defer func() {
			if rec := recover(); rec != nil {
				context.JSON(http.StatusInternalServerError, response.BaseResponse{
					Status:     false,
					StatusCode: http.StatusInternalServerError,
					RequestId:  requestId,
					Data:       nil,
					Message:    "Internal Server Error",
					Error:      rec,
				})
			}
		}()
		context.Next()

		if len(context.Errors) > 0 {
			context.JSON(http.StatusBadRequest, response.BaseResponse{
				Status:     false,
				StatusCode: http.StatusBadRequest,
				RequestId:  requestId,
				Data:       nil,
				Message:    "Bad Request",
				Error:      context.Errors.String(),
			})
		}

		status := context.Writer.Status()
		if status >= 400 {
			context.JSON(status, response.BaseResponse{
				Status:     false,
				StatusCode: http.StatusNotFound,
				RequestId:  requestId,
				Data:       nil,
				Message:    http.StatusText(status),
				Error:      nil,
			})
		}
	}
}
