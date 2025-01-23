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
		Limit:  100,
	}

	return limiter2.Limit(rate)
}

func RateLoginPublic() gin.HandlerFunc {
	rate := limiter.Rate{
		Period: 1 * time.Hour,
		Limit:  5,
	}

	return limiter2.Limit(rate)
}

func RateRegisterPublic() gin.HandlerFunc {
	rate := limiter.Rate{
		Period: 1 * time.Hour * 24 * 7,
		Limit:  311,
	}

	return limiter2.Limit(rate)
}

func RateRefreshPublic() gin.HandlerFunc {
	rate := limiter.Rate{
		Period: 1 * time.Hour * 24 * 7,
		Limit:  1,
	}

	return limiter2.Limit(rate)
}
