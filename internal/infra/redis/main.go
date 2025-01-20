package redis

import (
	"context"
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

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", configRedis.RedisHost, configRedis.RedisPort),
		Username: configRedis.RedisUsername,
		Password: configRedis.RedisPassword,
	})

	ClientRedis = redisClient

	checkRedisConnection(redisClient)

	logApp.Infoln("Success connect to Redis.")

	return redisClient
}

func checkRedisConnection(client *redis.Client) {
	err := client.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
}
