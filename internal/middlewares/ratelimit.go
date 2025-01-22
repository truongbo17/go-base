package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/redis"
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

func RateLimit() gin.HandlerFunc {
	rate := limiter.Rate{
		Period: 1 * time.Hour,
		Limit:  2,
	}

	client := redis2.ClientRedis

	store, err := redis.NewStoreWithOptions(client, limiter.StoreOptions{
		Prefix: "your_own_prefix",
	})
	if err != nil {
		panic(err)
	}

	options := []mgin.Option{
		mgin.WithKeyGetter(keyGetterIP),
		mgin.WithLimitReachedHandler(limitReachedHandler),
	}

	middleware := mgin.NewMiddleware(limiter.New(store, rate), options...)

	return middleware
}
