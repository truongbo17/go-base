package cache

import (
	"github.com/chenyahui/gin-cache/persist"
	"github.com/go-redis/redis/v8"
)

func InitCacheRedis() *persist.RedisStore {
	return persist.NewRedisStore(redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	}))
}
