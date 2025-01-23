package schedule

import (
	redislock "github.com/go-co-op/gocron-redis-lock/v2"
	"github.com/go-co-op/gocron/v2"
	"go-base/config"
	"go-base/internal/infra/logger"
	"go-base/internal/infra/redis"
	"time"
)

func Init() {
	client := redis.ClientRedis

	locker, err := redislock.NewRedisLocker(client, redislock.WithTries(config.DefaultScheduleLockRedisRetry))
	if err != nil {
		panic(err)
	}

	s, err := gocron.NewScheduler(gocron.WithDistributedLocker(locker))
	if err != nil {
		panic(err)
	}
	logApp := logger.LogrusLogger

	// Register job for schedule/cron at here.
	_, err = s.NewJob(
		gocron.DailyJob(1, gocron.NewAtTimes(gocron.NewAtTime(12, 0, 0))),
		gocron.NewTask(func(time string) {
			logApp.Infoln("Success" + time)
		}, time.Now().Format(time.DateTime)))
	if err != nil {
		logApp.Errorln(err)
	}

	s.Start()

	logApp.Infoln("Success init schedule/cron")
}
