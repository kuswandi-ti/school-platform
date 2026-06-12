package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	ServiceName        string
	AppEnv             string
	HTTPAddr           string
	LogLevel           string
	CORSAllowedOrigins []string
	JWTPublicKeyPath   string
	JWTIssuer          string
	JWTAudience        string
	ShutdownTimeout    time.Duration
	GRPCTargets        GRPCTargets
}

type GRPCTargets struct {
	Identity      string
	SchoolCore    string
	Admission     string
	Academic      string
	Finance       string
	Communication string
	Reporting     string
}

func Load() (Config, error) {
	shutdownSeconds, err := intFromEnv("SHUTDOWN_TIMEOUT_SECONDS", 10)
	if err != nil {
		return Config{}, err
	}
	httpAddr := httpAddrFromEnv()

	cfg := Config{
		ServiceName:        stringFromEnv("SERVICE_NAME", "api-gateway"),
		AppEnv:             stringFromEnv("APP_ENV", "local"),
		HTTPAddr:           httpAddr,
		LogLevel:           strings.ToLower(stringFromEnv("LOG_LEVEL", "info")),
		CORSAllowedOrigins: listFromEnv("CORS_ALLOWED_ORIGINS", []string{"http://localhost:3000"}),
		JWTPublicKeyPath:   stringFromEnv("JWT_PUBLIC_KEY_PATH", "./secrets/jwt/public.pem"),
		JWTIssuer:          stringFromEnv("JWT_ISSUER", "school-platform-identity"),
		JWTAudience:        stringFromEnv("JWT_AUDIENCE", "school-platform"),
		ShutdownTimeout:    time.Duration(shutdownSeconds) * time.Second,
		GRPCTargets: GRPCTargets{
			Identity:      stringFromEnv("IDENTITY_GRPC_ADDR", "localhost:9101"),
			SchoolCore:    stringFromEnv("SCHOOL_CORE_GRPC_ADDR", "localhost:9102"),
			Admission:     stringFromEnv("ADMISSION_GRPC_ADDR", "localhost:9103"),
			Academic:      stringFromEnv("ACADEMIC_GRPC_ADDR", "localhost:9104"),
			Finance:       stringFromEnv("FINANCE_GRPC_ADDR", "localhost:9105"),
			Communication: stringFromEnv("COMMUNICATION_GRPC_ADDR", "localhost:9106"),
			Reporting:     stringFromEnv("REPORTING_GRPC_ADDR", "localhost:9107"),
		},
	}

	if cfg.ServiceName == "" {
		return Config{}, fmt.Errorf("SERVICE_NAME is required")
	}
	if cfg.HTTPAddr == "" {
		return Config{}, fmt.Errorf("HTTP_ADDR is required")
	}

	return cfg, nil
}

func httpAddrFromEnv() string {
	if value := strings.TrimSpace(os.Getenv("HTTP_ADDR")); value != "" {
		return value
	}

	port := stringFromEnv("HTTP_PORT", "8080")
	return ":" + strings.TrimPrefix(port, ":")
}

func stringFromEnv(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}

func listFromEnv(key string, fallback []string) []string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}

	parts := strings.Split(value, ",")
	items := make([]string, 0, len(parts))
	for _, part := range parts {
		item := strings.TrimSpace(part)
		if item != "" {
			items = append(items, item)
		}
	}
	if len(items) == 0 {
		return fallback
	}

	return items
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
