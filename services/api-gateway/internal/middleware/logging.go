package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func AccessLog(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startedAt := time.Now()
			recorder := &statusRecorder{
				ResponseWriter: w,
				statusCode:     http.StatusOK,
			}

			next.ServeHTTP(recorder, r)

			logger.Info(
				"http request completed",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Int("status", recorder.statusCode),
				slog.Int64("duration_ms", time.Since(startedAt).Milliseconds()),
				slog.String("request_id", RequestIDFromContext(r.Context())),
				slog.String("correlation_id", CorrelationIDFromContext(r.Context())),
			)
		})
	}
}

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (r *statusRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}
