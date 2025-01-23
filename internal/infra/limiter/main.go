package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"github.com/ulule/limiter/v3/drivers/store/redis"
	"go-base/config"
	"go-base/internal/infra/logger"
	redis2 "go-base/internal/infra/redis"
	"net/http"
	"strconv"
)

func limitReachedHandler(context *gin.Context) {
	context.AbortWithStatus(http.StatusTooManyRequests)
}

var store limiter.Store

func keyGetterUserID(c *gin.Context) string {
	return strconv.FormatInt(c.MustGet("userId").(int64), 16)
}

func keyGetterIP(context *gin.Context) string {
	return context.ClientIP()
}

func Limit(rate limiter.Rate) gin.HandlerFunc {
	options := []mgin.Option{
		mgin.WithKeyGetter(keyGetterIP),
		mgin.WithLimitReachedHandler(limitReachedHandler),
	}

	return mgin.NewMiddleware(limiter.New(store, rate), options...)
}

func InitLimiterStore(storeCache string) {
	options := limiter.StoreOptions{
		Prefix: config.CacheKeyRateLimit,
	}
	if storeCache == config.CacheStoreRedis {
		client := redis2.ClientRedis
		storeRedis, err := redis.NewStoreWithOptions(client, options)
		if err != nil {
			panic(err)
		}

		store = storeRedis
	} else {
		store = memory.NewStoreWithOptions(options)
	}

	logApp := logger.LogrusLogger
	logApp.Infoln("Success init limiter middleware with store " + storeCache)
}
