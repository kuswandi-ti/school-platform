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

func TestRefreshHandlerSuccessUsesStandardResponse(t *testing.T) {
	client := &refreshIdentityClientStub{response: &identityv1.RefreshResponse{
		AccessToken: "new-access-token", RefreshToken: "new-refresh-token",
		TokenType: "Bearer", ExpiresIn: 900,
	}}
	recorder := performRefresh(t, client, `{"refresh_token":"old-refresh-token"}`)
	requireStatus(t, recorder, http.StatusOK)

	var envelope response.Envelope
	if err := json.Unmarshal(recorder.Body.Bytes(), &envelope); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if envelope.Error != nil || envelope.Meta != nil || envelope.Data == nil {
		t.Fatalf("unexpected standard envelope: %#v", envelope)
	}
	if client.request.GetCorrelationId() != "correlation-test" {
		t.Fatalf("expected correlation id propagation, got %q", client.request.GetCorrelationId())
	}
}

func TestRefreshHandlerRejectsReusedExpiredOrRevokedToken(t *testing.T) {
	recorder := performRefresh(t, &refreshIdentityClientStub{
		err: status.Error(codes.Unauthenticated, "invalid refresh token"),
	}, `{"refresh_token":"invalid-refresh-token"}`)
	requireError(t, recorder, http.StatusUnauthorized, response.CodeUnauthorized)
}

func TestRefreshHandlerValidationError(t *testing.T) {
	recorder := performRefresh(t, &refreshIdentityClientStub{}, `{"refresh_token":""}`)
	requireError(t, recorder, http.StatusBadRequest, response.CodeValidationError)
}

type refreshIdentityClientStub struct {
	request  *identityv1.RefreshRequest
	response *identityv1.RefreshResponse
	err      error
}

func (s *refreshIdentityClientStub) Login(_ context.Context, _ *identityv1.LoginRequest, _ ...grpc.CallOption) (*identityv1.LoginResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *refreshIdentityClientStub) Refresh(_ context.Context, request *identityv1.RefreshRequest, _ ...grpc.CallOption) (*identityv1.RefreshResponse, error) {
	s.request = request
	return s.response, s.err
}

func (s *refreshIdentityClientStub) Logout(_ context.Context, _ *identityv1.LogoutRequest, _ ...grpc.CallOption) (*identityv1.LogoutResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func performRefresh(t *testing.T, client *refreshIdentityClientStub, body string) *httptest.ResponseRecorder {
	t.Helper()
	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/refresh", bytes.NewBufferString(body))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set(middleware.HeaderRequestID, "request-test")
	request.Header.Set(middleware.HeaderCorrelationID, "correlation-test")
	recorder := httptest.NewRecorder()
	middleware.RequestContext(RefreshHandler(client)).ServeHTTP(recorder, request)
	return recorder
}
