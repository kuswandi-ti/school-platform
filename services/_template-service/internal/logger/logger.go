package logger

import (
	"log/slog"
	"os"

	"school-platform/services/_template-service/internal/config"
)

func New(cfg config.Config) *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: parseLevel(cfg.LogLevel),
	})

	return slog.New(handler).With(
		slog.String("service", cfg.ServiceName),
		slog.String("environment", cfg.AppEnv),
	)
}

func parseLevel(value string) slog.Level {
	switch value {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
