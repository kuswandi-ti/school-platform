package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestContextUsesIncomingHeaders(t *testing.T) {
	var gotRequestID string
	var gotCorrelationID string

	handler := RequestContext(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotRequestID = RequestIDFromContext(r.Context())
		gotCorrelationID = CorrelationIDFromContext(r.Context())
		w.WriteHeader(http.StatusNoContent)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(HeaderRequestID, "req-123")
	req.Header.Set(HeaderCorrelationID, "corr-123")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if gotRequestID != "req-123" {
		t.Fatalf("expected request ID req-123, got %q", gotRequestID)
	}
	if gotCorrelationID != "corr-123" {
		t.Fatalf("expected correlation ID corr-123, got %q", gotCorrelationID)
	}
	if rec.Header().Get(HeaderRequestID) != "req-123" {
		t.Fatalf("expected response request ID header")
	}
	if rec.Header().Get(HeaderCorrelationID) != "corr-123" {
		t.Fatalf("expected response correlation ID header")
	}
}

func TestRequestContextDefaultsCorrelationIDToRequestID(t *testing.T) {
	var gotRequestID string
	var gotCorrelationID string

	handler := RequestContext(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotRequestID = RequestIDFromContext(r.Context())
		gotCorrelationID = CorrelationIDFromContext(r.Context())
		w.WriteHeader(http.StatusNoContent)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if gotRequestID == "" {
		t.Fatal("expected generated request ID")
	}
	if gotCorrelationID == "" {
		t.Fatal("expected generated correlation ID")
	}
	if gotRequestID != gotCorrelationID {
		t.Fatalf("expected correlation ID to default to request ID, got request=%q correlation=%q", gotRequestID, gotCorrelationID)
	}
}
