package logger

import (
	"log/slog"
	"os"
	"school-platform/services/school-core-service/internal/config"
)

func New(c config.Config) *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})).With("service", c.ServiceName, "environment", c.AppEnv)
}
