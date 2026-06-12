package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"school-platform/services/identity-service/internal/domain"
	"school-platform/services/identity-service/internal/password"
	"school-platform/services/identity-service/internal/repository"
	"school-platform/services/identity-service/internal/token"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserInactive       = errors.New("user is not active")
)

type UserStore interface {
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	UpdateLastLogin(ctx context.Context, id uuid.UUID, lastLoginAt time.Time) (domain.User, error)
}

type SessionStore interface {
	CreateSession(ctx context.Context, params repository.CreateSessionParams) error
}

type TokenIssuer interface {
	Issue(userID uuid.UUID) (token.Tokens, error)
}

type Login struct {
	users             UserStore
	sessions          SessionStore
	tokens            TokenIssuer
	dummyPasswordHash string
	now               func() time.Time
}

type LoginInput struct {
	Email     string
	Password  string
	IPAddress *string
	UserAgent *string
}

type LoginOutput struct {
	UserID       uuid.UUID
	DisplayName  string
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
}

func NewLogin(users UserStore, sessions SessionStore, tokens TokenIssuer) (*Login, error) {
	dummyPasswordHash, err := password.Hash(uuid.NewString())
	if err != nil {
		return nil, fmt.Errorf("create dummy password hash: %w", err)
	}
	return &Login{
		users:             users,
		sessions:          sessions,
		tokens:            tokens,
		dummyPasswordHash: dummyPasswordHash,
		now:               time.Now,
	}, nil
}

func (u *Login) Execute(ctx context.Context, input LoginInput) (LoginOutput, error) {
	user, err := u.users.FindByEmail(ctx, strings.TrimSpace(input.Email))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			_, _ = password.Verify(input.Password, u.dummyPasswordHash)
			return LoginOutput{}, ErrInvalidCredentials
		}
		return LoginOutput{}, fmt.Errorf("find login user: %w", err)
	}

	verified, err := password.Verify(input.Password, user.PasswordHash)
	if err != nil {
		return LoginOutput{}, fmt.Errorf("verify stored password hash: %w", err)
	}
	if !verified {
		return LoginOutput{}, ErrInvalidCredentials
	}
	if user.Status != "active" {
		return LoginOutput{}, ErrUserInactive
	}

	issued, err := u.tokens.Issue(user.ID)
	if err != nil {
		return LoginOutput{}, fmt.Errorf("issue login tokens: %w", err)
	}
	if err := u.sessions.CreateSession(ctx, repository.CreateSessionParams{
		ID:               uuid.New(),
		UserID:           user.ID,
		RefreshTokenHash: issued.RefreshTokenHash,
		IPAddress:        input.IPAddress,
		UserAgent:        input.UserAgent,
		ExpiresAt:        issued.RefreshExpiresAt,
	}); err != nil {
		return LoginOutput{}, fmt.Errorf("persist login session: %w", err)
	}

	now := u.now().UTC()
	if _, err := u.users.UpdateLastLogin(ctx, user.ID, now); err != nil {
		return LoginOutput{}, fmt.Errorf("update login timestamp: %w", err)
	}

	return LoginOutput{
		UserID:       user.ID,
		DisplayName:  user.DisplayName,
		AccessToken:  issued.AccessToken,
		RefreshToken: issued.RefreshToken,
		ExpiresIn:    int64(issued.AccessExpiresAt.Sub(now).Seconds()),
	}, nil
}
