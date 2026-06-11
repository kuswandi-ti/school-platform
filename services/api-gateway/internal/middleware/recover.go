package middleware

import (
	"log/slog"
	"net/http"

	"school-platform/services/api-gateway/internal/response"
)

func Recover(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if recovered := recover(); recovered != nil {
					logger.Error(
						"panic recovered",
						slog.Any("panic", recovered),
						slog.String("request_id", RequestIDFromContext(r.Context())),
						slog.String("correlation_id", CorrelationIDFromContext(r.Context())),
					)
					response.Error(w, http.StatusInternalServerError, response.CodeInternalError, "Internal server error.", nil)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
