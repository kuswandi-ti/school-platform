package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	ServiceName     string
	AppEnv          string
	HTTPAddr        string
	LogLevel        string
	ShutdownTimeout time.Duration
}

func Load() (Config, error) {
	shutdownSeconds, err := intFromEnv("SHUTDOWN_TIMEOUT_SECONDS", 10)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{
		ServiceName:     stringFromEnv("SERVICE_NAME", "_template-service"),
		AppEnv:          stringFromEnv("APP_ENV", "local"),
		HTTPAddr:        stringFromEnv("HTTP_ADDR", ":8081"),
		LogLevel:        strings.ToLower(stringFromEnv("LOG_LEVEL", "info")),
		ShutdownTimeout: time.Duration(shutdownSeconds) * time.Second,
	}

	if cfg.ServiceName == "" {
		return Config{}, fmt.Errorf("SERVICE_NAME is required")
	}
	if cfg.HTTPAddr == "" {
		return Config{}, fmt.Errorf("HTTP_ADDR is required")
	}

	return cfg, nil
}

func stringFromEnv(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}

func intFromEnv(key string, fallback int) (int, error) {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback, nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be an integer: %w", key, err)
	}
	if parsed <= 0 {
		return 0, fmt.Errorf("%s must be greater than zero", key)
	}

	return parsed, nil
}
