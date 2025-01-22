package middlewares

import (
	"github.com/gin-gonic/gin"
	libredis "github.com/redis/go-redis/v9"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/redis"
	"log"
	"time"
)

func RateLimit() gin.HandlerFunc {
	rate := limiter.Rate{
		Period: 1 * time.Hour,
		Limit:  1000,
	}

	option, err := libredis.ParseURL("redis://localhost:6379/0")
	if err != nil {
		log.Fatal(err)
	}
	client := libredis.NewClient(option)

	store, err := redis.NewStoreWithOptions(client, limiter.StoreOptions{
		Prefix: "your_own_prefix",
	})
	if err != nil {
		panic(err)
	}

	middleware := mgin.NewMiddleware(limiter.New(store, rate))
	return middleware
}
