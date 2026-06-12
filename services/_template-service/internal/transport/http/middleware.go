package httptransport

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log/slog"
	"net/http"
	"time"
)

const (
	headerRequestID     = "X-Request-ID"
	headerCorrelationID = "X-Correlation-ID"
)

type contextKey string

const (
	requestIDContextKey     contextKey = "request_id"
	correlationIDContextKey contextKey = "correlation_id"
)

func RequestContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get(headerRequestID)
		if requestID == "" {
			requestID = newID()
		}

		correlationID := r.Header.Get(headerCorrelationID)
		if correlationID == "" {
			correlationID = requestID
		}

		w.Header().Set(headerRequestID, requestID)
		w.Header().Set(headerCorrelationID, correlationID)

		ctx := context.WithValue(r.Context(), requestIDContextKey, requestID)
		ctx = context.WithValue(ctx, correlationIDContextKey, correlationID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AccessLogMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
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

func RequestIDFromContext(ctx context.Context) string {
	value, _ := ctx.Value(requestIDContextKey).(string)
	return value
}

func CorrelationIDFromContext(ctx context.Context) string {
	value, _ := ctx.Value(correlationIDContextKey).(string)
	return value
}

func newID() string {
	var bytes [16]byte
	if _, err := rand.Read(bytes[:]); err != nil {
		return time.Now().UTC().Format("20060102150405.000000000")
	}
	return hex.EncodeToString(bytes[:])
}

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (r *statusRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}
