package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/redis"
	"go-base/config"
	redis2 "go-base/internal/infra/redis"
	"net/http"
	"strconv"
	"time"
)

func limitReachedHandler(context *gin.Context) {
	context.AbortWithStatus(http.StatusTooManyRequests)
}

func keyGetterUserID(c *gin.Context) string {
	return strconv.FormatInt(c.MustGet("userId").(int64), 16)
}

func keyGetterIP(context *gin.Context) string {
	return context.ClientIP()
}

func limit(rate limiter.Rate) gin.HandlerFunc {
	client := redis2.ClientRedis
	store, err := redis.NewStoreWithOptions(client, limiter.StoreOptions{
		Prefix: config.CacheKeyRateLimit,
	})
	if err != nil {
		panic(err)
	}
	options := []mgin.Option{
		mgin.WithKeyGetter(keyGetterIP),
		mgin.WithLimitReachedHandler(limitReachedHandler),
	}

	return mgin.NewMiddleware(limiter.New(store, rate), options...)
}

func RateGlobalLimit() gin.HandlerFunc {
	rate := limiter.Rate{
		Period: 1 * time.Hour,
		Limit:  1000,
	}

	return limit(rate)
}

func RateLimitPublic() gin.HandlerFunc {
	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  1000,
	}

	return limit(rate)
}
