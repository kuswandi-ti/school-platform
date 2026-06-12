package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"

	"school-platform/services/identity-service/internal/domain"
	"school-platform/services/identity-service/internal/password"
	"school-platform/services/identity-service/internal/repository"
	"school-platform/services/identity-service/internal/token"
)

func TestLoginSuccess(t *testing.T) {
	passwordHash, err := password.Hash("valid password")
	require.NoError(t, err)
	userID := uuid.New()
	users := &userStoreStub{user: domain.User{
		ID: userID, Email: "user@example.com", PasswordHash: passwordHash,
		DisplayName: "Test User", Status: "active",
	}}
	sessions := &sessionStoreStub{}
	issuer := tokenIssuerStub{tokens: token.Tokens{
		AccessToken: "access-token", RefreshToken: "refresh-token",
		RefreshTokenHash: "refresh-token-hash",
		AccessExpiresAt:  time.Now().Add(15 * time.Minute),
		RefreshExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	}}

	login, err := NewLogin(users, sessions, issuer)
	require.NoError(t, err)
	result, err := login.Execute(context.Background(), LoginInput{
		Email: "user@example.com", Password: "valid password",
	})
	require.NoError(t, err)
	require.Equal(t, userID, result.UserID)
	require.Equal(t, "access-token", result.AccessToken)
	require.Equal(t, "refresh-token", result.RefreshToken)
	require.Equal(t, "refresh-token-hash", sessions.created.RefreshTokenHash)
	require.NotEqual(t, result.RefreshToken, sessions.created.RefreshTokenHash)
	require.True(t, users.lastLoginUpdated)
}

func TestLoginRejectsWrongPassword(t *testing.T) {
	passwordHash, err := password.Hash("valid password")
	require.NoError(t, err)
	sessions := &sessionStoreStub{}
	login, err := NewLogin(&userStoreStub{user: domain.User{
		ID: uuid.New(), PasswordHash: passwordHash, Status: "active",
	}}, sessions, tokenIssuerStub{})
	require.NoError(t, err)
	_, err = login.Execute(context.Background(), LoginInput{Password: "wrong password"})
	require.ErrorIs(t, err, ErrInvalidCredentials)
	require.Equal(t, 0, sessions.calls)
}

func TestLoginRejectsInactiveAndLockedUsers(t *testing.T) {
	passwordHash, err := password.Hash("valid password")
	require.NoError(t, err)
	for _, userStatus := range []string{"inactive", "locked"} {
		t.Run(userStatus, func(t *testing.T) {
			sessions := &sessionStoreStub{}
			login, err := NewLogin(&userStoreStub{user: domain.User{
				ID: uuid.New(), PasswordHash: passwordHash, Status: userStatus,
			}}, sessions, tokenIssuerStub{})
			require.NoError(t, err)
			_, err = login.Execute(context.Background(), LoginInput{Password: "valid password"})
			require.ErrorIs(t, err, ErrUserInactive)
			require.Equal(t, 0, sessions.calls)
		})
	}
}

func TestLoginRejectsUnknownEmail(t *testing.T) {
	login, err := NewLogin(&userStoreStub{err: pgx.ErrNoRows}, &sessionStoreStub{}, tokenIssuerStub{})
	require.NoError(t, err)
	_, err = login.Execute(context.Background(), LoginInput{Email: "missing@example.com", Password: "password"})
	require.ErrorIs(t, err, ErrInvalidCredentials)
}

func TestLoginRejectsMalformedStoredHashAsInternalError(t *testing.T) {
	login, err := NewLogin(&userStoreStub{user: domain.User{
		ID: uuid.New(), PasswordHash: "malformed", Status: "active",
	}}, &sessionStoreStub{}, tokenIssuerStub{})
	require.NoError(t, err)
	_, err = login.Execute(context.Background(), LoginInput{Password: "password"})
	require.Error(t, err)
	require.NotErrorIs(t, err, ErrInvalidCredentials)
}

type userStoreStub struct {
	user             domain.User
	err              error
	lastLoginUpdated bool
}

func (s *userStoreStub) FindByEmail(context.Context, string) (domain.User, error) {
	return s.user, s.err
}

func (s *userStoreStub) UpdateLastLogin(_ context.Context, _ uuid.UUID, _ time.Time) (domain.User, error) {
	s.lastLoginUpdated = true
	return s.user, nil
}

type sessionStoreStub struct {
	created repository.CreateSessionParams
	calls   int
}

func (s *sessionStoreStub) CreateSession(_ context.Context, params repository.CreateSessionParams) error {
	s.created = params
	s.calls++
	return nil
}

type tokenIssuerStub struct {
	tokens token.Tokens
	err    error
}

func (s tokenIssuerStub) Issue(uuid.UUID) (token.Tokens, error) {
	if s.err != nil {
		return token.Tokens{}, s.err
	}
	if s.tokens.AccessToken == "" {
		return token.Tokens{}, errors.New("token issuer should not be called")
	}
	return s.tokens, nil
}
