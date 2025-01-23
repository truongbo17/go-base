package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go-base/config"
	_ "go-base/docs"
	"go-base/internal/infra/asynq"
	"go-base/internal/infra/cache"
	"go-base/internal/infra/database"
	"go-base/internal/infra/limiter"
	"go-base/internal/infra/logger"
	"go-base/internal/infra/redis"
	"go-base/internal/infra/schedule"
	"go-base/internal/routes"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	StartServerCmd = &cobra.Command{
		Use:   "server",
		Short: `Start the server`,
		Run: func(cmd *cobra.Command, args []string) {
			start()
		},
	}
)

func start() {
	config.Init()
	EnvConfig := config.EnvConfig

	appEnv := EnvConfig.AppConfig.Env

	if appEnv == config.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	logger.Init()
	log := logger.LogrusLogger

	log.Infoln("Success init config")

	redis.ConnectRedis()

	storeCache := EnvConfig.CacheConfig.CacheStore
	cache.InitCache(storeCache)

	limiter.InitLimiterStore(storeCache)

	database.ConnectDatabase(&EnvConfig.DatabaseConnection)

	if storeCache == config.CacheStoreRedis {
		schedule.Init()

		asynq.InitClient()
		asynq.InitServer()
	}

	routes.Init(appEnv)
	r := routes.Router

	server := &http.Server{
		Addr:         ":" + EnvConfig.AppConfig.Port,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Second * 30,
		Handler:      r,
	}

	log.Printf("Server is now listening at port: %s. Good luck!", EnvConfig.AppConfig.Port)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server listen error: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	i := <-quit
	log.Println("Server receive a signal: ", i.String())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %s", err)
	}
	log.Println("Server exiting.")
}
