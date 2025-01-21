package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-base/internal/infra/logger"
	"go-base/internal/response"
	"go-base/internal/utils"
	"net/http"
	"runtime/debug"
)

type RequestLogStack struct {
	RequestID string `json:"request_id"`
	Stack     string `json:"stack"`
	Err       any    `json:"err"`
}

func ErrorHandle() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestId := context.GetString("x-request-id")
		logApp := logger.LogrusLogger

		defer func() {
			if rec := recover(); rec != nil {
				logApp.WithFields(logrus.Fields{
					"stack":     "\n" + utils.FormatStackTrace(debug.Stack()),
					"requestId": requestId,
				}).Errorln(rec)

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

		status := context.Writer.Status()
		if status >= 400 {
			context.JSON(status, response.BaseResponse{
				Status:     false,
				StatusCode: status,
				RequestId:  requestId,
				Data:       nil,
				Message:    http.StatusText(status),
				Error:      context.Errors.String(),
			})
			return
		}

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
