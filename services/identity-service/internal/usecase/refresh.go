package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"school-platform/services/identity-service/internal/domain"
	"school-platform/services/identity-service/internal/repository"
	"school-platform/services/identity-service/internal/token"
)

var (
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
	ErrRefreshTokenReused  = errors.New("refresh token was reused or revoked")
	ErrRefreshTokenExpired = errors.New("refresh token expired")
)

type RefreshUserStore interface {
	FindByID(ctx context.Context, id uuid.UUID) (domain.User, error)
}

type RefreshSessionStore interface {
	FindByRefreshTokenHash(ctx context.Context, refreshTokenHash string) (repository.Session, error)
	RotateSession(ctx context.Context, params repository.RotateSessionParams) error
}

type Refresh struct {
	users    RefreshUserStore
	contexts LoginContextStore
	sessions RefreshSessionStore
	tokens   TokenIssuer
	now      func() time.Time
}

type RefreshInput struct {
	RefreshToken string
	IPAddress    *string
	UserAgent    *string
}

type RefreshOutput struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
}

func NewRefresh(users RefreshUserStore, contexts LoginContextStore, sessions RefreshSessionStore, tokens TokenIssuer) *Refresh {
	return &Refresh{users: users, contexts: contexts, sessions: sessions, tokens: tokens, now: time.Now}
}

func (u *Refresh) Execute(ctx context.Context, input RefreshInput) (RefreshOutput, error) {
	now := u.now().UTC()
	refreshTokenHash := token.HashRefreshToken(input.RefreshToken)
	session, err := u.sessions.FindByRefreshTokenHash(ctx, refreshTokenHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return RefreshOutput{}, ErrInvalidRefreshToken
		}
		return RefreshOutput{}, fmt.Errorf("find refresh session: %w", err)
	}
	if session.RevokedAt != nil {
		return RefreshOutput{}, ErrRefreshTokenReused
	}
	if !session.ExpiresAt.After(now) {
		return RefreshOutput{}, ErrRefreshTokenExpired
	}

	user, err := u.users.FindByID(ctx, session.UserID)
	if err != nil {
		return RefreshOutput{}, fmt.Errorf("find refresh user: %w", err)
	}
	if user.Status != "active" {
		return RefreshOutput{}, ErrUserInactive
	}

	userContext, err := u.contexts.GetUserContext(ctx, session.UserID, now)
	if err != nil {
		return RefreshOutput{}, fmt.Errorf("load refresh user context: %w", err)
	}

	issued, err := u.tokens.Issue(session.UserID, actorClaimsFromContext(userContext))
	if err != nil {
		return RefreshOutput{}, fmt.Errorf("issue refreshed tokens: %w", err)
	}
	if err := u.sessions.RotateSession(ctx, repository.RotateSessionParams{
		CurrentSessionID:    session.ID,
		NewSessionID:        uuid.New(),
		UserID:              session.UserID,
		DeviceID:            session.DeviceID,
		RefreshTokenHash:    issued.RefreshTokenHash,
		IPAddress:           input.IPAddress,
		UserAgent:           input.UserAgent,
		CurrentUsedAt:       now,
		NewSessionExpiresAt: issued.RefreshExpiresAt,
	}); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return RefreshOutput{}, ErrRefreshTokenReused
		}
		return RefreshOutput{}, fmt.Errorf("rotate refresh session: %w", err)
	}

	return RefreshOutput{
		AccessToken:  issued.AccessToken,
		RefreshToken: issued.RefreshToken,
		ExpiresIn:    int64(issued.AccessExpiresAt.Sub(now).Seconds()),
	}, nil
}
