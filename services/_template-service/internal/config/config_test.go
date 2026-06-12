package config

import (
	"testing"
	"time"
)

func TestLoad_UsesHTTPPort(t *testing.T) {
	t.Setenv("SERVICE_NAME", "sample-service")
	t.Setenv("APP_ENV", "test")
	t.Setenv("HTTP_PORT", "8181")
	t.Setenv("HTTP_ADDR", "")
	t.Setenv("LOG_LEVEL", "debug")
	t.Setenv("SHUTDOWN_TIMEOUT_SECONDS", "7")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("expected config to load, got error: %v", err)
	}

	if cfg.ServiceName != "sample-service" {
		t.Fatalf("expected service name sample-service, got %q", cfg.ServiceName)
	}
	if cfg.AppEnv != "test" {
		t.Fatalf("expected app env test, got %q", cfg.AppEnv)
	}
	if cfg.HTTPPort != "8181" {
		t.Fatalf("expected HTTP port 8181, got %q", cfg.HTTPPort)
	}
	if cfg.HTTPAddr != ":8181" {
		t.Fatalf("expected HTTP addr :8181, got %q", cfg.HTTPAddr)
	}
	if cfg.LogLevel != "debug" {
		t.Fatalf("expected log level debug, got %q", cfg.LogLevel)
	}
	if cfg.ShutdownTimeout != 7*time.Second {
		t.Fatalf("expected shutdown timeout 7s, got %s", cfg.ShutdownTimeout)
	}
}

func TestLoad_AllowsHTTPAddrOverride(t *testing.T) {
	t.Setenv("HTTP_PORT", "8181")
	t.Setenv("HTTP_ADDR", "127.0.0.1:9090")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("expected config to load, got error: %v", err)
	}

	if cfg.HTTPAddr != "127.0.0.1:9090" {
		t.Fatalf("expected HTTP addr override, got %q", cfg.HTTPAddr)
	}
}

func TestLoad_RejectsInvalidShutdownTimeout(t *testing.T) {
	t.Setenv("SHUTDOWN_TIMEOUT_SECONDS", "invalid")

	if _, err := Load(); err == nil {
		t.Fatal("expected invalid shutdown timeout to fail")
	}
}
