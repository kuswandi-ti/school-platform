package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct{ ServiceName, AppEnv, GRPCAddr, DatabaseURL string }

func Load() (Config, error) {
	c := Config{ServiceName: get("SERVICE_NAME", "school-core-service"), AppEnv: get("APP_ENV", "local"), GRPCAddr: get("GRPC_ADDR", ":9102"), DatabaseURL: get("DATABASE_URL", "")}
	if c.DatabaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is required")
	}
	return c, nil
}
func get(k, d string) string {
	v := strings.TrimSpace(os.Getenv(k))
	if v == "" {
		return d
	}
	return v
}
