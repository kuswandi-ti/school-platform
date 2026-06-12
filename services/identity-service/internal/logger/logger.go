package logger

import (
	"log/slog"
	"os"
)

func New(serviceName, environment string) *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil)).With(
		slog.String("service", serviceName),
		slog.String("environment", environment),
	)
}
