package asynq

import (
	"fmt"
	asynqPackage "github.com/hibiken/asynq"
	"go-base/config"
	"go-base/internal/infra/logger"
)

func InitClient() {
	EnvConfig := config.EnvConfig
	configRedis := EnvConfig.CacheConfig
	logApp := logger.LogrusLogger

	rdb := asynqPackage.NewClient(asynqPackage.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%s", configRedis.RedisHost, configRedis.RedisPort),
		Username: configRedis.RedisUsername,
		Password: configRedis.RedisPassword,
	})
	defer func(rdb *asynqPackage.Client) {
		err := rdb.Close()
		if err != nil {
			panic(err)
		}
	}(rdb)
	logApp.Infoln("Success init asynq queue.")
}
