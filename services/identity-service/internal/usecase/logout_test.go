package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"school-platform/services/identity-service/internal/repository"
	"school-platform/services/identity-service/internal/token"
)

func TestLogoutRevokesActorSession(t *testing.T) {
	now := time.Now().UTC().Truncate(time.Microsecond)
	actorID := uuid.New()
	sessions := &logoutSessionStoreStub{session: repository.Session{ID: uuid.New(), UserID: actorID}}
	logout := NewLogout(sessions, accessTokenValidatorStub{userID: actorID})
	logout.now = func() time.Time { return now }

	err := logout.Execute(context.Background(), LogoutInput{AccessToken: "access", RefreshToken: "refresh"})
	require.NoError(t, err)
	require.Equal(t, actorID, sessions.revoked.UserID)
	require.Equal(t, token.HashRefreshToken("refresh"), sessions.revoked.RefreshTokenHash)
	require.Equal(t, now, sessions.revoked.RevokedAt)
}

func TestLogoutRejectsInvalidActor(t *testing.T) {
	logout := NewLogout(&logoutSessionStoreStub{}, accessTokenValidatorStub{err: errors.New("invalid")})
	err := logout.Execute(context.Background(), LogoutInput{AccessToken: "invalid", RefreshToken: "refresh"})
	require.ErrorIs(t, err, ErrInvalidAccessToken)
}

func TestLogoutRejectsSessionOwnedByAnotherUser(t *testing.T) {
	logout := NewLogout(&logoutSessionStoreStub{session: repository.Session{UserID: uuid.New()}}, accessTokenValidatorStub{userID: uuid.New()})
	err := logout.Execute(context.Background(), LogoutInput{AccessToken: "access", RefreshToken: "refresh"})
	require.ErrorIs(t, err, ErrSessionForbidden)
}

func TestLogoutIsIdempotentForRevokedSession(t *testing.T) {
	revokedAt := time.Now().UTC()
	actorID := uuid.New()
	sessions := &logoutSessionStoreStub{session: repository.Session{UserID: actorID, RevokedAt: &revokedAt}}
	logout := NewLogout(sessions, accessTokenValidatorStub{userID: actorID})
	require.NoError(t, logout.Execute(context.Background(), LogoutInput{AccessToken: "access", RefreshToken: "refresh"}))
	require.Zero(t, sessions.revokeCalls)
}

type accessTokenValidatorStub struct {
	userID uuid.UUID
	err    error
}

func (s accessTokenValidatorStub) ValidateAccessToken(string) (uuid.UUID, error) {
	return s.userID, s.err
}

type logoutSessionStoreStub struct {
	session     repository.Session
	findErr     error
	revoked     repository.RevokeSessionParams
	revokeErr   error
	revokeCalls int
}

func (s *logoutSessionStoreStub) FindByRefreshTokenHash(context.Context, string) (repository.Session, error) {
	return s.session, s.findErr
}

func (s *logoutSessionStoreStub) RevokeSession(_ context.Context, params repository.RevokeSessionParams) error {
	s.revoked = params
	s.revokeCalls++
	return s.revokeErr
}
