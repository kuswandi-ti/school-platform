package httptransport

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	identityv1 "school-platform/packages/proto/gen/go/identity/v1"
	"school-platform/services/api-gateway/internal/middleware"
	"school-platform/services/api-gateway/internal/response"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestLoginHandlerSuccessUsesStandardResponse(t *testing.T) {
	client := &identityClientStub{response: &identityv1.LoginResponse{
		UserId: "user-id", DisplayName: "Test User", AccessToken: "access-token",
		RefreshToken: "refresh-token", TokenType: "Bearer", ExpiresIn: 900,
	}}
	recorder := performLogin(t, client, `{"email":"user@example.com","password":"valid password"}`)
	requireStatus(t, recorder, http.StatusOK)

	var envelope response.Envelope
	if err := json.Unmarshal(recorder.Body.Bytes(), &envelope); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if envelope.Error != nil || envelope.Meta != nil || envelope.Data == nil {
		t.Fatalf("unexpected standard envelope: %#v", envelope)
	}
	if bytes.Contains(recorder.Body.Bytes(), []byte("password_hash")) {
		t.Fatal("response must not expose password hash")
	}
	if client.request.GetCorrelationId() != "correlation-test" {
		t.Fatalf("expected correlation id propagation, got %q", client.request.GetCorrelationId())
	}
}

func TestLoginHandlerWrongPassword(t *testing.T) {
	recorder := performLogin(t, &identityClientStub{err: status.Error(codes.Unauthenticated, "invalid credentials")}, `{"email":"user@example.com","password":"wrong"}`)
	requireError(t, recorder, http.StatusUnauthorized, response.CodeUnauthorized)
}

func TestLoginHandlerInactiveUser(t *testing.T) {
	recorder := performLogin(t, &identityClientStub{err: status.Error(codes.PermissionDenied, "inactive")}, `{"email":"user@example.com","password":"valid password"}`)
	requireError(t, recorder, http.StatusForbidden, response.CodeForbidden)
}

func TestLoginHandlerValidationError(t *testing.T) {
	recorder := performLogin(t, &identityClientStub{}, `{"email":"not-an-email","password":""}`)
	requireError(t, recorder, http.StatusBadRequest, response.CodeValidationError)
}

type identityClientStub struct {
	request  *identityv1.LoginRequest
	response *identityv1.LoginResponse
	err      error
}

func (s *identityClientStub) Login(_ context.Context, request *identityv1.LoginRequest, _ ...grpc.CallOption) (*identityv1.LoginResponse, error) {
	s.request = request
	return s.response, s.err
}

func (s *identityClientStub) Refresh(_ context.Context, _ *identityv1.RefreshRequest, _ ...grpc.CallOption) (*identityv1.RefreshResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func performLogin(t *testing.T, client *identityClientStub, body string) *httptest.ResponseRecorder {
	t.Helper()
	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBufferString(body))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set(middleware.HeaderRequestID, "request-test")
	request.Header.Set(middleware.HeaderCorrelationID, "correlation-test")
	recorder := httptest.NewRecorder()
	middleware.RequestContext(LoginHandler(client)).ServeHTTP(recorder, request)
	return recorder
}

func requireStatus(t *testing.T, recorder *httptest.ResponseRecorder, expected int) {
	t.Helper()
	if recorder.Code != expected {
		t.Fatalf("expected status %d, got %d: %s", expected, recorder.Code, recorder.Body.String())
	}
}

func requireError(t *testing.T, recorder *httptest.ResponseRecorder, expectedStatus int, expectedCode string) {
	t.Helper()
	requireStatus(t, recorder, expectedStatus)
	var envelope response.Envelope
	if err := json.Unmarshal(recorder.Body.Bytes(), &envelope); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if envelope.Data != nil || envelope.Meta != nil || envelope.Error == nil || envelope.Error.Code != expectedCode {
		t.Fatalf("unexpected error envelope: %#v", envelope)
	}
}
