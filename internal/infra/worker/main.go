package worker

import (
	"fmt"
	"github.com/hibiken/asynq"
	"go-base/config"
	"go-base/internal/app/auth/jobs"
	"go-base/internal/infra/logger"
	"log"
)

var ClientWorker *asynq.Client

func InitClient() {
	EnvConfig := config.EnvConfig
	configRedis := EnvConfig.CacheConfig
	logApp := logger.LogrusLogger

	rdb := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%s", configRedis.RedisHost, configRedis.RedisPort),
		Username: configRedis.RedisUsername,
		Password: configRedis.RedisPassword,
	})
	defer func(rdb *asynq.Client) {
		err := rdb.Close()
		if err != nil {
			panic(err)
		}
	}(rdb)

	ClientWorker = rdb

	logApp.Infoln("Success init client asynq queue.")
}

func InitServer() {
	EnvConfig := config.EnvConfig
	configRedis := EnvConfig.CacheConfig
	logApp := logger.LogrusLogger

	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     fmt.Sprintf("%s:%s", configRedis.RedisHost, configRedis.RedisPort),
			Username: configRedis.RedisUsername,
			Password: configRedis.RedisPassword,
		},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	mux := asynq.NewServeMux()

	mux.HandleFunc(jobs.TypeEmailRegister, jobs.HandleSendMailRegister)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
	logApp.Infoln("Success init server asynq queue.")
}
