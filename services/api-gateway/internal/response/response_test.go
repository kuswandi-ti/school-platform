package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSONWritesStandardEnvelope(t *testing.T) {
	rec := httptest.NewRecorder()

	JSON(rec, http.StatusOK, map[string]string{"message": "pong"}, nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var body Envelope
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if body.Data == nil {
		t.Fatal("expected data to be set")
	}
	if body.Meta != nil {
		t.Fatalf("expected meta to be nil, got %#v", body.Meta)
	}
	if body.Error != nil {
		t.Fatalf("expected error to be nil, got %#v", body.Error)
	}
}

func TestErrorWritesStandardEnvelope(t *testing.T) {
	rec := httptest.NewRecorder()

	Error(rec, http.StatusNotFound, CodeNotFound, "Resource not found.", nil)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}

	var body Envelope
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if body.Data != nil {
		t.Fatalf("expected data to be nil, got %#v", body.Data)
	}
	if body.Error == nil {
		t.Fatal("expected error to be set")
	}
	if body.Error.Code != CodeNotFound {
		t.Fatalf("expected error code %q, got %q", CodeNotFound, body.Error.Code)
	}
}
