package main

import (
	"log/slog"
	"os"

	"school-platform/services/api-gateway/internal/app"
	"school-platform/services/api-gateway/internal/config"
	"school-platform/services/api-gateway/internal/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", slog.String("error", err.Error()))
		os.Exit(1)
	}

	log := logger.New(cfg)
	application := app.New(cfg, log)

	if err := application.Run(); err != nil {
		log.Error("api gateway stopped with error", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
