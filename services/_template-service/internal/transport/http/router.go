package httptransport

import (
	"log/slog"
	"net/http"
)

func NewRouter(logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", HealthHandler())
	mux.HandleFunc("GET /readyz", ReadinessHandler())

	return RequestContextMiddleware(AccessLogMiddleware(logger, mux))
}
