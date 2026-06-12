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
	HTTPPort        string
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
		HTTPPort:        stringFromEnv("HTTP_PORT", "8081"),
		LogLevel:        strings.ToLower(stringFromEnv("LOG_LEVEL", "info")),
		ShutdownTimeout: time.Duration(shutdownSeconds) * time.Second,
	}
	cfg.HTTPAddr = httpAddrFromEnv(cfg.HTTPPort)

	if cfg.ServiceName == "" {
		return Config{}, fmt.Errorf("SERVICE_NAME is required")
	}
	if cfg.HTTPPort == "" {
		return Config{}, fmt.Errorf("HTTP_PORT is required")
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

func httpAddrFromEnv(port string) string {
	addr := strings.TrimSpace(os.Getenv("HTTP_ADDR"))
	if addr != "" {
		return addr
	}
	return ":" + port
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
