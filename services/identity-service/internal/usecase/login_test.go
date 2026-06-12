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
	contexts := loginContextStoreStub{context: domain.UserContext{
		UserID: userID, Roles: []string{"guru"}, Permissions: []string{"academic.grade.manage"},
		Assignments: []domain.RoleAssignment{{FoundationID: uuid.MustParse("11111111-1111-1111-1111-111111111111")}},
	}}
	issuer := &tokenIssuerStub{tokens: token.Tokens{
		AccessToken: "access-token", RefreshToken: "refresh-token",
		RefreshTokenHash: "refresh-token-hash",
		AccessExpiresAt:  time.Now().Add(15 * time.Minute),
		RefreshExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	}}

	login, err := NewLogin(users, sessions, contexts, issuer)
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
	require.Equal(t, "11111111-1111-1111-1111-111111111111", issuer.actorClaims.FoundationID)
	require.Equal(t, []string{"guru"}, issuer.actorClaims.Roles)
}

func TestLoginRejectsWrongPassword(t *testing.T) {
	passwordHash, err := password.Hash("valid password")
	require.NoError(t, err)
	sessions := &sessionStoreStub{}
	contexts := loginContextStoreStub{}
	login, err := NewLogin(&userStoreStub{user: domain.User{
		ID: uuid.New(), PasswordHash: passwordHash, Status: "active",
	}}, sessions, contexts, &tokenIssuerStub{})
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
			contexts := loginContextStoreStub{}
			login, err := NewLogin(&userStoreStub{user: domain.User{
				ID: uuid.New(), PasswordHash: passwordHash, Status: userStatus,
			}}, sessions, contexts, &tokenIssuerStub{})
			require.NoError(t, err)
			_, err = login.Execute(context.Background(), LoginInput{Password: "valid password"})
			require.ErrorIs(t, err, ErrUserInactive)
			require.Equal(t, 0, sessions.calls)
		})
	}
}

func TestLoginRejectsUnknownEmail(t *testing.T) {
	login, err := NewLogin(&userStoreStub{err: pgx.ErrNoRows}, &sessionStoreStub{}, loginContextStoreStub{}, &tokenIssuerStub{})
	require.NoError(t, err)
	_, err = login.Execute(context.Background(), LoginInput{Email: "missing@example.com", Password: "password"})
	require.ErrorIs(t, err, ErrInvalidCredentials)
}

func TestLoginRejectsMalformedStoredHashAsInternalError(t *testing.T) {
	login, err := NewLogin(&userStoreStub{user: domain.User{
		ID: uuid.New(), PasswordHash: "malformed", Status: "active",
	}}, &sessionStoreStub{}, loginContextStoreStub{}, &tokenIssuerStub{})
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
	tokens      token.Tokens
	err         error
	actorClaims token.ActorClaims
}

func (s *tokenIssuerStub) Issue(_ uuid.UUID, actorClaims token.ActorClaims) (token.Tokens, error) {
	s.actorClaims = actorClaims
	if s.err != nil {
		return token.Tokens{}, s.err
	}
	if s.tokens.AccessToken == "" {
		return token.Tokens{}, errors.New("token issuer should not be called")
	}
	return s.tokens, nil
}

type loginContextStoreStub struct {
	context domain.UserContext
	err     error
}

func (s loginContextStoreStub) GetUserContext(context.Context, uuid.UUID, time.Time) (domain.UserContext, error) {
	return s.context, s.err
}
