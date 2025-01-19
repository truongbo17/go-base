package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go-base/config"
	_ "go-base/docs"
	"go-base/internal/infra/logger"
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

const (
	// DebugMode app env debug stg.
	DebugMode = "debug"
	// ReleaseMode app env debug production.
	ReleaseMode = "release"
)

const (
	debugCode = iota
	releaseCode
)

var (
	ginMode = debugCode
)

func start() {
	config.Init()
	EnvConfig := config.EnvConfig

	appEnv := EnvConfig.AppConfig.Env

	switch appEnv {
	case DebugMode:
		ginMode = debugCode
	case ReleaseMode:
		ginMode = releaseCode
	default:
		panic("APP_ENV unknown: " + appEnv + " (available mode: debug release)")
	}

	if appEnv == ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	routes.Init()
	r := routes.Router

	logger.Init()
	log := logger.LogrusLogger

	server := &http.Server{
		Addr:    ":" + EnvConfig.AppConfig.Port,
		Handler: r,
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
