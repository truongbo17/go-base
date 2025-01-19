package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-base/config"
	"go-base/internal/infra/logger"
)

var RedisClient *redis.Client

func ConnectRedis() *redis.Client {
	logApp := logger.LogrusLogger
	EnvConfig := config.EnvConfig
	configRedis := EnvConfig.CacheConfig

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", configRedis.RedisHost, configRedis.RedisPort),
		Username: configRedis.RedisUsername,
		Password: configRedis.RedisPassword,
	})

	RedisClient = redisClient

	logApp.Infoln("Success connect to Redis.")

	return redisClient
}
