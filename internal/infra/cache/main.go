package cache

import (
	"go-base/config"
	"go-base/internal/infra/logger"
	"go-base/internal/infra/redis"
	"time"
)

type ICache interface {
	Set(key string, value interface{}, ttl time.Duration) error
	Get(key string) (interface{}, error)
	Delete(key string) error
}

var Cache ICache

func InitCache(storeCache string) {
	if storeCache == config.CacheStoreRedis {
		Cache = NewRedisCache(redis.ClientRedis)
	} else {
		Cache = NewLocalCache(5*time.Minute, 10*time.Minute)
	}

	logApp := logger.LogrusLogger
	logApp.Infoln("Success init cache with store " + storeCache)
}
