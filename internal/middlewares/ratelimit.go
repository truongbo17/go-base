package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	limiter2 "go-base/internal/infra/limiter"
	"time"
)

func RateGlobalLimit() gin.HandlerFunc {
	rate := limiter.Rate{
		Period: 1 * time.Hour,
		Limit:  1000,
	}

	return limiter2.Limit(rate)
}

func RateLimitPublic() gin.HandlerFunc {
	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  1000,
	}

	return limiter2.Limit(rate)
}
