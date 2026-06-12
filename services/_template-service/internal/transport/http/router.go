package httptransport

import (
	"log/slog"
	"net/http"
)

func NewRouter(logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", HealthHandler())
	mux.HandleFunc("GET /readyz", ReadinessHandler())
	mux.HandleFunc("GET /metrics", MetricsHandler())

	return RequestContextMiddleware(AccessLogMiddleware(logger, mux))
}
