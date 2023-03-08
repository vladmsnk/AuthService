package app

import (
	"auth/vladmsnk/config"
	v1 "auth/vladmsnk/internal/controller/v1"
	"auth/vladmsnk/internal/repository"
	"auth/vladmsnk/internal/usecase"
	"auth/vladmsnk/pkg/logger"
	"auth/vladmsnk/pkg/postgres"
	httpserver "auth/vladmsnk/pkg/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg *config.Config) {
	lg := logger.New(cfg.Log.Level)

	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		lg.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	authUseCase := usecase.NewAuthUseCase(repository.New(pg), cfg.Auth)

	handler := gin.New()
	v1.NewRouter(handler, lg, authUseCase)
	server := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port),
		httpserver.ReadTimeout(time.Duration(cfg.HTTP.ReadTimeout)),
		httpserver.WriteTimeout(time.Duration(cfg.HTTP.WriteTimeout)),
		httpserver.ShutdownTimeout(time.Duration(cfg.HTTP.ShutdownTimeout)))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		lg.Info("app - Run - signal: " + s.String())
	case err = <-server.Notify():
		lg.Error(fmt.Errorf("app - Run - server.Notify: %w", err))
	}

	err = server.Shutdown()
	if err != nil {
		lg.Error(fmt.Errorf("app - Run - server.Shutdown: %w", err))
	}
}
