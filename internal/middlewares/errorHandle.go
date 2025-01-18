package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-base/internal/response"
	"net/http"
)

func ErrorHandle() gin.HandlerFunc {
	return func(context *gin.Context) {
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
	}
}
