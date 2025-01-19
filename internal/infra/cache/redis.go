package cache

import (
	"github.com/chenyahui/gin-cache/persist"
	"go-base/internal/infra/redis"
)

func InitCacheRedis() *persist.RedisStore {
	return persist.NewRedisStore(redis.ClientRedis)
}
