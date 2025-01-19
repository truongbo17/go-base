package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-base/internal/infra/logger"
	"time"
)

type RequestLog struct {
	RequestID string  `json:"request_id"`
	ClientIP  string  `json:"client_ip"`
	UserAgent string  `json:"user_agent"`
	Method    string  `json:"method"`
	Path      string  `json:"path"`
	Latency   float64 `json:"latency"`
	Timestamp string  `json:"string"`
}

func RequestLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		timeNow := time.Now()
		requestId := context.GetString("x-request-id")
		clientIp := context.ClientIP()
		userAgent := context.Request.UserAgent()
		method := context.Request.Method
		path := context.Request.URL.Path

		context.Next()

		latency := time.Since(timeNow).Seconds()

		logData := RequestLog{
			RequestID: requestId,
			ClientIP:  clientIp,
			UserAgent: userAgent,
			Method:    method,
			Path:      path,
			Latency:   latency,
			Timestamp: timeNow.Format(time.DateTime),
		}

		log := logger.LogrusLogger
		log.Infof("RequestLog: %+v", logData)
	}
}
