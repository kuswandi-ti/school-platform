package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"school-platform/services/identity-service/internal/repository"
	"school-platform/services/identity-service/internal/token"
)

var (
	ErrInvalidAccessToken = errors.New("invalid access token")
	ErrSessionForbidden   = errors.New("session does not belong to actor")
)

type AccessTokenValidator interface {
	ValidateAccessToken(accessToken string) (uuid.UUID, error)
}

type LogoutSessionStore interface {
	FindByRefreshTokenHash(ctx context.Context, refreshTokenHash string) (repository.Session, error)
	RevokeSession(ctx context.Context, params repository.RevokeSessionParams) error
}

type Logout struct {
	sessions LogoutSessionStore
	tokens   AccessTokenValidator
	now      func() time.Time
}

type LogoutInput struct {
	AccessToken  string
	RefreshToken string
}

func NewLogout(sessions LogoutSessionStore, tokens AccessTokenValidator) *Logout {
	return &Logout{sessions: sessions, tokens: tokens, now: time.Now}
}

func (u *Logout) Execute(ctx context.Context, input LogoutInput) error {
	actorID, err := u.tokens.ValidateAccessToken(input.AccessToken)
	if err != nil {
		return ErrInvalidAccessToken
	}

	refreshTokenHash := token.HashRefreshToken(input.RefreshToken)
	session, err := u.sessions.FindByRefreshTokenHash(ctx, refreshTokenHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrInvalidRefreshToken
		}
		return fmt.Errorf("find logout session: %w", err)
	}
	if session.UserID != actorID {
		return ErrSessionForbidden
	}
	if session.RevokedAt != nil {
		return nil
	}

	err = u.sessions.RevokeSession(ctx, repository.RevokeSessionParams{
		RefreshTokenHash: refreshTokenHash,
		UserID:           actorID,
		RevokedAt:        u.now().UTC(),
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("revoke logout session: %w", err)
	}
	return nil
}
