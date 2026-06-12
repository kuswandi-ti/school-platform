package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"

	"school-platform/services/identity-service/internal/domain"
	"school-platform/services/identity-service/internal/repository"
	"school-platform/services/identity-service/internal/token"
)

func TestRefreshRotatesToken(t *testing.T) {
	now := time.Now().UTC().Truncate(time.Microsecond)
	userID := uuid.New()
	sessionID := uuid.New()
	sessions := &refreshSessionStoreStub{session: repository.Session{
		ID: sessionID, UserID: userID, ExpiresAt: now.Add(time.Hour),
	}}
	issuer := tokenIssuerStub{tokens: token.Tokens{
		AccessToken: "new-access-token", RefreshToken: "new-refresh-token",
		RefreshTokenHash: token.HashRefreshToken("new-refresh-token"),
		AccessExpiresAt:  now.Add(15 * time.Minute), RefreshExpiresAt: now.Add(30 * 24 * time.Hour),
	}}
	refresh := NewRefresh(refreshUserStoreStub{user: domain.User{ID: userID, Status: "active"}}, sessions, issuer)
	refresh.now = func() time.Time { return now }

	result, err := refresh.Execute(context.Background(), RefreshInput{RefreshToken: "old-refresh-token"})
	require.NoError(t, err)
	require.Equal(t, "new-access-token", result.AccessToken)
	require.Equal(t, "new-refresh-token", result.RefreshToken)
	require.Equal(t, token.HashRefreshToken("old-refresh-token"), sessions.searchedHash)
	require.Equal(t, sessionID, sessions.rotated.CurrentSessionID)
	require.Equal(t, issuer.tokens.RefreshTokenHash, sessions.rotated.RefreshTokenHash)
	require.Equal(t, now, sessions.rotated.CurrentUsedAt)
}

func TestRefreshRejectsReusedOrRevokedToken(t *testing.T) {
	revokedAt := time.Now().UTC()
	sessions := &refreshSessionStoreStub{session: repository.Session{
		ID: uuid.New(), UserID: uuid.New(), ExpiresAt: revokedAt.Add(time.Hour), RevokedAt: &revokedAt,
	}}
	refresh := NewRefresh(refreshUserStoreStub{}, sessions, tokenIssuerStub{})

	_, err := refresh.Execute(context.Background(), RefreshInput{RefreshToken: "reused-token"})
	require.ErrorIs(t, err, ErrRefreshTokenReused)
	require.Zero(t, sessions.rotateCalls)
}

func TestRefreshRejectsExpiredToken(t *testing.T) {
	now := time.Now().UTC()
	sessions := &refreshSessionStoreStub{session: repository.Session{
		ID: uuid.New(), UserID: uuid.New(), ExpiresAt: now.Add(-time.Second),
	}}
	refresh := NewRefresh(refreshUserStoreStub{}, sessions, tokenIssuerStub{})
	refresh.now = func() time.Time { return now }

	_, err := refresh.Execute(context.Background(), RefreshInput{RefreshToken: "expired-token"})
	require.ErrorIs(t, err, ErrRefreshTokenExpired)
}

func TestRefreshRejectsUnknownToken(t *testing.T) {
	refresh := NewRefresh(refreshUserStoreStub{}, &refreshSessionStoreStub{findErr: pgx.ErrNoRows}, tokenIssuerStub{})
	_, err := refresh.Execute(context.Background(), RefreshInput{RefreshToken: "unknown-token"})
	require.ErrorIs(t, err, ErrInvalidRefreshToken)
}

func TestRefreshDetectsConcurrentReuse(t *testing.T) {
	now := time.Now().UTC()
	userID := uuid.New()
	sessions := &refreshSessionStoreStub{
		session:   repository.Session{ID: uuid.New(), UserID: userID, ExpiresAt: now.Add(time.Hour)},
		rotateErr: pgx.ErrNoRows,
	}
	issuer := tokenIssuerStub{tokens: token.Tokens{
		AccessToken: "discarded-access-token", RefreshToken: "discarded-refresh-token",
		RefreshTokenHash: token.HashRefreshToken("discarded-refresh-token"),
		AccessExpiresAt:  now.Add(time.Minute), RefreshExpiresAt: now.Add(time.Hour),
	}}
	refresh := NewRefresh(refreshUserStoreStub{user: domain.User{ID: userID, Status: "active"}}, sessions, issuer)
	refresh.now = func() time.Time { return now }

	_, err := refresh.Execute(context.Background(), RefreshInput{RefreshToken: "already-consumed-concurrently"})
	require.ErrorIs(t, err, ErrRefreshTokenReused)
}

type refreshUserStoreStub struct {
	user domain.User
	err  error
}

func (s refreshUserStoreStub) FindByID(context.Context, uuid.UUID) (domain.User, error) {
	return s.user, s.err
}

type refreshSessionStoreStub struct {
	session      repository.Session
	findErr      error
	rotateErr    error
	searchedHash string
	rotated      repository.RotateSessionParams
	rotateCalls  int
}

func (s *refreshSessionStoreStub) FindByRefreshTokenHash(_ context.Context, hash string) (repository.Session, error) {
	s.searchedHash = hash
	return s.session, s.findErr
}

func (s *refreshSessionStoreStub) RotateSession(_ context.Context, params repository.RotateSessionParams) error {
	s.rotated = params
	s.rotateCalls++
	return s.rotateErr
}
