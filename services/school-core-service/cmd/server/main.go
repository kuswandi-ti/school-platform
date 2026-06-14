package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"school-platform/services/school-core-service/internal/app"
	"school-platform/services/school-core-service/internal/config"
	"school-platform/services/school-core-service/internal/logger"
	"syscall"
)

func main() {
	c, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	l := logger.New(c)
	if err = app.New(c, l).Run(ctx); err != nil {
		l.Error("school core service stopped with error", "error", err)
		os.Exit(1)
	}
}
