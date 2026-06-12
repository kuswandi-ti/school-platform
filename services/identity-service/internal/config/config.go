package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	ServiceName       string
	AppEnv            string
	GRPCAddr          string
	DatabaseURL       string
	JWTPrivateKeyPath string
	JWTIssuer         string
	JWTAudience       string
	AccessTokenTTL    time.Duration
	RefreshTokenTTL   time.Duration
}

func Load() (Config, error) {
	accessMinutes, err := positiveInt("ACCESS_TOKEN_TTL_MINUTES", 15)
	if err != nil {
		return Config{}, err
	}
	refreshDays, err := positiveInt("REFRESH_TOKEN_TTL_DAYS", 30)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{
		ServiceName:       env("SERVICE_NAME", "identity-service"),
		AppEnv:            env("APP_ENV", "local"),
		GRPCAddr:          env("GRPC_ADDR", ":9101"),
		DatabaseURL:       env("DATABASE_URL", ""),
		JWTPrivateKeyPath: env("JWT_PRIVATE_KEY_PATH", ""),
		JWTIssuer:         env("JWT_ISSUER", "school-platform-identity"),
		JWTAudience:       env("JWT_AUDIENCE", "school-platform"),
		AccessTokenTTL:    time.Duration(accessMinutes) * time.Minute,
		RefreshTokenTTL:   time.Duration(refreshDays) * 24 * time.Hour,
	}
	if cfg.DatabaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is required")
	}
	if cfg.JWTPrivateKeyPath == "" {
		return Config{}, fmt.Errorf("JWT_PRIVATE_KEY_PATH is required")
	}
	return cfg, nil
}

func env(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}

func positiveInt(key string, fallback int) (int, error) {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback, nil
	}
	parsed, err := strconv.Atoi(value)
	if err != nil || parsed <= 0 {
		return 0, fmt.Errorf("%s must be a positive integer", key)
	}
	return parsed, nil
}
