package middleware

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"school-platform/services/api-gateway/internal/response"
)

func TestRequireAuthRejectsMissingToken(t *testing.T) {
	validator, _ := newTestJWTValidator(t)
	called := false
	handler := RequestContext(RequireAuth(validator)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusNoContent)
	})))

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/protected", nil))

	requireErrorEnvelope(t, recorder, http.StatusUnauthorized, response.CodeUnauthorized)
	if called {
		t.Fatal("protected handler should not be called")
	}
}

func TestRequireAuthRejectsInvalidToken(t *testing.T) {
	validator, _ := newTestJWTValidator(t)
	handler := RequestContext(RequireAuth(validator)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})))

	request := httptest.NewRequest(http.MethodGet, "/protected", nil)
	request.Header.Set("Authorization", "Bearer invalid-token")
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)

	requireErrorEnvelope(t, recorder, http.StatusUnauthorized, response.CodeUnauthorized)
}

func TestRequireAuthStoresActorContext(t *testing.T) {
	validator, privateKey := newTestJWTValidator(t)
	accessToken := signTestAccessToken(t, privateKey, AccessTokenClaims{
		FoundationID: "foundation-1",
		SchoolID:     "school-1",
		Roles:        []string{"guru"},
		Permissions:  []string{"academic.grade.manage"},
		Scope: map[string]any{
			"class_ids": []string{"class-1"},
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "user-1",
			Issuer:    "school-platform-identity",
			Audience:  jwt.ClaimStrings{"school-platform"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})

	var actorContext ActorContext
	var accessTokenFromContext string
	handler := RequestContext(RequireAuth(validator)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ok bool
		actorContext, ok = ActorContextFromContext(r.Context())
		if !ok {
			t.Fatal("expected actor context in request context")
		}
		accessTokenFromContext, ok = AccessTokenFromContext(r.Context())
		if !ok {
			t.Fatal("expected raw access token in request context")
		}
		w.WriteHeader(http.StatusNoContent)
	})))

	request := httptest.NewRequest(http.MethodGet, "/protected", nil)
	request.Header.Set("Authorization", "Bearer "+accessToken)
	request.Header.Set(HeaderRequestID, "request-test")
	request.Header.Set(HeaderCorrelationID, "correlation-test")
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, recorder.Code)
	}
	if actorContext.UserID != "user-1" || actorContext.FoundationID != "foundation-1" || actorContext.SchoolID != "school-1" {
		t.Fatalf("unexpected actor context: %#v", actorContext)
	}
	if actorContext.RequestID != "request-test" || actorContext.CorrelationID != "correlation-test" {
		t.Fatalf("expected request metadata in actor context, got %#v", actorContext)
	}
	if accessTokenFromContext != accessToken {
		t.Fatalf("expected access token from context")
	}
}

func TestAccessLogDoesNotLogAuthorizationToken(t *testing.T) {
	validator, privateKey := newTestJWTValidator(t)
	accessToken := signTestAccessToken(t, privateKey, AccessTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "user-1",
			Issuer:    "school-platform-identity",
			Audience:  jwt.ClaimStrings{"school-platform"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})

	var logs bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&logs, nil))
	handler := RequestContext(AccessLog(logger)(RequireAuth(validator)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))))

	request := httptest.NewRequest(http.MethodGet, "/protected", nil)
	request.Header.Set("Authorization", "Bearer "+accessToken)
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, recorder.Code)
	}
	if bytes.Contains(logs.Bytes(), []byte(accessToken)) {
		t.Fatal("authorization token leaked into logs")
	}
}

func newTestJWTValidator(t *testing.T) (*JWTValidator, ed25519.PrivateKey) {
	t.Helper()
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("generate Ed25519 key: %v", err)
	}
	validator, err := NewJWTValidator(publicKey, "school-platform-identity", "school-platform")
	if err != nil {
		t.Fatalf("create JWT validator: %v", err)
	}
	return validator, privateKey
}

func signTestAccessToken(t *testing.T, privateKey ed25519.PrivateKey, claims AccessTokenClaims) string {
	t.Helper()
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims).SignedString(privateKey)
	if err != nil {
		t.Fatalf("sign test access token: %v", err)
	}
	return tokenString
}

func requireErrorEnvelope(t *testing.T, recorder *httptest.ResponseRecorder, expectedStatus int, expectedCode string) {
	t.Helper()
	if recorder.Code != expectedStatus {
		t.Fatalf("expected status %d, got %d", expectedStatus, recorder.Code)
	}
	var envelope response.Envelope
	if err := json.Unmarshal(recorder.Body.Bytes(), &envelope); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if envelope.Error == nil || envelope.Error.Code != expectedCode {
		t.Fatalf("unexpected error envelope: %#v", envelope)
	}
}
