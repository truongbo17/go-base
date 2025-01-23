package asynq

import (
	"fmt"
	asynqPackage "github.com/hibiken/asynq"
	"go-base/config"
	"go-base/internal/infra/logger"
	"log"
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
	logApp.Infoln("Success init client asynq queue.")
}

func InitServer() {
	EnvConfig := config.EnvConfig
	configRedis := EnvConfig.CacheConfig
	logApp := logger.LogrusLogger

	srv := asynqPackage.NewServer(
		asynqPackage.RedisClientOpt{
			Addr:     fmt.Sprintf("%s:%s", configRedis.RedisHost, configRedis.RedisPort),
			Username: configRedis.RedisUsername,
			Password: configRedis.RedisPassword,
		},
		asynqPackage.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)
	mux := asynqPackage.NewServeMux()
	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
	logApp.Infoln("Success init server asynq queue.")
}
