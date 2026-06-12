package grpctransport

import (
	"context"
	"errors"
	"strings"

	identityv1 "school-platform/packages/proto/gen/go/identity/v1"
	"school-platform/services/identity-service/internal/usecase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IdentityServer struct {
	identityv1.UnimplementedIdentityServiceServer
	login   *usecase.Login
	refresh *usecase.Refresh
}

func NewIdentityServer(login *usecase.Login, refresh *usecase.Refresh) *IdentityServer {
	return &IdentityServer{login: login, refresh: refresh}
}

func (s *IdentityServer) Refresh(ctx context.Context, request *identityv1.RefreshRequest) (*identityv1.RefreshResponse, error) {
	if request.GetRefreshToken() == "" {
		return nil, status.Error(codes.InvalidArgument, "refresh token is required")
	}
	output, err := s.refresh.Execute(ctx, usecase.RefreshInput{
		RefreshToken: request.GetRefreshToken(),
		IPAddress:    optionalString(request.GetIpAddress()),
		UserAgent:    optionalString(request.GetUserAgent()),
	})
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrInvalidRefreshToken),
			errors.Is(err, usecase.ErrRefreshTokenReused),
			errors.Is(err, usecase.ErrRefreshTokenExpired):
			return nil, status.Error(codes.Unauthenticated, "invalid refresh token")
		case errors.Is(err, usecase.ErrUserInactive):
			return nil, status.Error(codes.PermissionDenied, "user is not active")
		default:
			return nil, status.Error(codes.Internal, "refresh failed")
		}
	}
	return &identityv1.RefreshResponse{
		AccessToken:  output.AccessToken,
		RefreshToken: output.RefreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    output.ExpiresIn,
	}, nil
}

func (s *IdentityServer) Login(ctx context.Context, request *identityv1.LoginRequest) (*identityv1.LoginResponse, error) {
	if strings.TrimSpace(request.GetEmail()) == "" || request.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "email and password are required")
	}

	output, err := s.login.Execute(ctx, usecase.LoginInput{
		Email:     request.GetEmail(),
		Password:  request.GetPassword(),
		IPAddress: optionalString(request.GetIpAddress()),
		UserAgent: optionalString(request.GetUserAgent()),
	})
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrInvalidCredentials):
			return nil, status.Error(codes.Unauthenticated, "invalid credentials")
		case errors.Is(err, usecase.ErrUserInactive):
			return nil, status.Error(codes.PermissionDenied, "user is not active")
		default:
			return nil, status.Error(codes.Internal, "login failed")
		}
	}

	return &identityv1.LoginResponse{
		UserId:       output.UserID.String(),
		DisplayName:  output.DisplayName,
		AccessToken:  output.AccessToken,
		RefreshToken: output.RefreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    output.ExpiresIn,
	}, nil
}

func optionalString(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}
