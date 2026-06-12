package httptransport

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	identityv1 "school-platform/packages/proto/gen/go/identity/v1"
	"school-platform/services/api-gateway/internal/middleware"
	"school-platform/services/api-gateway/internal/response"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestLogoutHandlerSuccessUsesStandardResponse(t *testing.T) {
	client := &logoutIdentityClientStub{response: &identityv1.LogoutResponse{}}
	recorder := performLogout(t, client, `{"refresh_token":"refresh-token"}`)
	requireStatus(t, recorder, http.StatusOK)

	var envelope response.Envelope
	if err := json.Unmarshal(recorder.Body.Bytes(), &envelope); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if envelope.Error != nil || envelope.Meta != nil || envelope.Data == nil {
		t.Fatalf("unexpected standard envelope: %#v", envelope)
	}
	if client.request.GetAccessToken() == "" || client.request.GetRefreshToken() != "refresh-token" {
		t.Fatalf("unexpected logout request: %#v", client.request)
	}
	if client.request.GetCorrelationId() != "correlation-test" {
		t.Fatalf("expected correlation id propagation, got %q", client.request.GetCorrelationId())
	}
}

func TestLogoutHandlerRequiresBearerToken(t *testing.T) {
	publicKey, _, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("generate Ed25519 key: %v", err)
	}
	validator, err := middleware.NewJWTValidator(publicKey, "school-platform-identity", "school-platform")
	if err != nil {
		t.Fatalf("create JWT validator: %v", err)
	}
	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/logout", bytes.NewBufferString(`{"refresh_token":"refresh-token"}`))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set(middleware.HeaderRequestID, "request-test")
	request.Header.Set(middleware.HeaderCorrelationID, "correlation-test")
	recorder := httptest.NewRecorder()
	handler := middleware.RequestContext(middleware.RequireAuth(validator)(LogoutHandler(&logoutIdentityClientStub{})))
	handler.ServeHTTP(recorder, request)
	requireError(t, recorder, http.StatusUnauthorized, response.CodeUnauthorized)
}

func TestLogoutHandlerRejectsInvalidSession(t *testing.T) {
	recorder := performLogout(t, &logoutIdentityClientStub{
		err: status.Error(codes.Unauthenticated, "invalid session"),
	}, `{"refresh_token":"refresh-token"}`)
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

func performLogout(t *testing.T, client *logoutIdentityClientStub, body string) *httptest.ResponseRecorder {
	t.Helper()
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("generate Ed25519 key: %v", err)
	}
	validator, err := middleware.NewJWTValidator(publicKey, "school-platform-identity", "school-platform")
	if err != nil {
		t.Fatalf("create JWT validator: %v", err)
	}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodEdDSA, middleware.AccessTokenClaims{
		FoundationID: "foundation-1",
		SchoolID:     "school-1",
		Roles:        []string{"guru"},
		Permissions:  []string{"academic.grade.manage"},
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "user-1",
			Issuer:    "school-platform-identity",
			Audience:  jwt.ClaimStrings{"school-platform"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}).SignedString(privateKey)
	if err != nil {
		t.Fatalf("sign access token: %v", err)
	}

	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/logout", bytes.NewBufferString(body))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+accessToken)
	request.Header.Set(middleware.HeaderRequestID, "request-test")
	request.Header.Set(middleware.HeaderCorrelationID, "correlation-test")
	recorder := httptest.NewRecorder()
	handler := middleware.RequestContext(middleware.RequireAuth(validator)(LogoutHandler(client)))
	handler.ServeHTTP(recorder, request)
	return recorder
}
