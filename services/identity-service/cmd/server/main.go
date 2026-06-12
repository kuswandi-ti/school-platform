package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"school-platform/services/identity-service/internal/app"
	"school-platform/services/identity-service/internal/config"
	"school-platform/services/identity-service/internal/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", slog.String("error", err.Error()))
		os.Exit(1)
	}
	log := logger.New(cfg.ServiceName, cfg.AppEnv)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := app.New(cfg, log).Run(ctx); err != nil {
		log.Error("identity service stopped with error", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
