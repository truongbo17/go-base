package server

import (
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go-base/config"
	"go-base/internal/infra/logger"
	"go-base/routers"
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
	routers.Init()
	r := routers.Router

	logger.Init()
	log := logger.LogrusLogger

	config.Init()
	EnvConfig := config.EnvConfig

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
