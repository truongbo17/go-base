package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-base/config"
	"go-base/internal/infra/logger"
)

var ClientRedis *redis.Client

func ConnectRedis() *redis.Client {
	logApp := logger.LogrusLogger
	EnvConfig := config.EnvConfig
	configRedis := EnvConfig.CacheConfig

	fmt.Println(1332313123, configRedis)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", configRedis.RedisHost, configRedis.RedisPort),
		Username: configRedis.RedisUsername,
		Password: configRedis.RedisPassword,
	})

	ClientRedis = redisClient

	logApp.Infoln("Success connect to Redis.")

	return redisClient
}
