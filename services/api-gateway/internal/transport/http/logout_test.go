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

func TestLogoutHandlerSuccessUsesStandardResponse(t *testing.T) {
	client := &logoutIdentityClientStub{response: &identityv1.LogoutResponse{}}
	recorder := performLogout(t, client, "access-token", `{"refresh_token":"refresh-token"}`)
	requireStatus(t, recorder, http.StatusOK)

	var envelope response.Envelope
	if err := json.Unmarshal(recorder.Body.Bytes(), &envelope); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if envelope.Error != nil || envelope.Meta != nil || envelope.Data == nil {
		t.Fatalf("unexpected standard envelope: %#v", envelope)
	}
	if client.request.GetAccessToken() != "access-token" || client.request.GetRefreshToken() != "refresh-token" {
		t.Fatalf("unexpected logout request: %#v", client.request)
	}
	if client.request.GetCorrelationId() != "correlation-test" {
		t.Fatalf("expected correlation id propagation, got %q", client.request.GetCorrelationId())
	}
}

func TestLogoutHandlerRequiresBearerToken(t *testing.T) {
	recorder := performLogout(t, &logoutIdentityClientStub{}, "", `{"refresh_token":"refresh-token"}`)
	requireError(t, recorder, http.StatusUnauthorized, response.CodeUnauthorized)
}

func TestLogoutHandlerRejectsInvalidSession(t *testing.T) {
	recorder := performLogout(t, &logoutIdentityClientStub{
		err: status.Error(codes.Unauthenticated, "invalid session"),
	}, "access-token", `{"refresh_token":"refresh-token"}`)
	requireError(t, recorder, http.StatusUnauthorized, response.CodeUnauthorized)
}

type logoutIdentityClientStub struct {
	request  *identityv1.LogoutRequest
	response *identityv1.LogoutResponse
	err      error
}

func (s *logoutIdentityClientStub) Login(context.Context, *identityv1.LoginRequest, ...grpc.CallOption) (*identityv1.LoginResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *logoutIdentityClientStub) Refresh(context.Context, *identityv1.RefreshRequest, ...grpc.CallOption) (*identityv1.RefreshResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *logoutIdentityClientStub) Logout(_ context.Context, request *identityv1.LogoutRequest, _ ...grpc.CallOption) (*identityv1.LogoutResponse, error) {
	s.request = request
	return s.response, s.err
}

func performLogout(t *testing.T, client *logoutIdentityClientStub, accessToken, body string) *httptest.ResponseRecorder {
	t.Helper()
	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/logout", bytes.NewBufferString(body))
	request.Header.Set("Content-Type", "application/json")
	if accessToken != "" {
		request.Header.Set("Authorization", "Bearer "+accessToken)
	}
	request.Header.Set(middleware.HeaderRequestID, "request-test")
	request.Header.Set(middleware.HeaderCorrelationID, "correlation-test")
	recorder := httptest.NewRecorder()
	middleware.RequestContext(LogoutHandler(client)).ServeHTTP(recorder, request)
	return recorder
}
