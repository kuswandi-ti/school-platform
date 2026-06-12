package config

import "testing"

func TestLoadUsesHTTPPortWhenHTTPAddrIsUnset(t *testing.T) {
	t.Setenv("HTTP_ADDR", "")
	t.Setenv("HTTP_PORT", "18080")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("expected config to load, got error: %v", err)
	}

	if cfg.HTTPAddr != ":18080" {
		t.Fatalf("expected HTTP address :18080, got %q", cfg.HTTPAddr)
	}
}

func TestLoadPrefersHTTPAddrOverHTTPPort(t *testing.T) {
	t.Setenv("HTTP_ADDR", "127.0.0.1:19090")
	t.Setenv("HTTP_PORT", "18080")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("expected config to load, got error: %v", err)
	}

	if cfg.HTTPAddr != "127.0.0.1:19090" {
		t.Fatalf("expected HTTP address 127.0.0.1:19090, got %q", cfg.HTTPAddr)
	}
}
