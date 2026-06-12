package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"school-platform/services/identity-service/internal/db/sqlc"
)

type SessionRepository struct {
	db      dbsqlc.DBTX
	queries *dbsqlc.Queries
}

type CreateSessionParams struct {
	ID               uuid.UUID
	UserID           uuid.UUID
	RefreshTokenHash string
	DeviceID         *uuid.UUID
	IPAddress        *string
	UserAgent        *string
	ExpiresAt        time.Time
}

type Session struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	DeviceID   *uuid.UUID
	ExpiresAt  time.Time
	RevokedAt  *time.Time
	LastUsedAt *time.Time
}

type RotateSessionParams struct {
	CurrentSessionID    uuid.UUID
	NewSessionID        uuid.UUID
	UserID              uuid.UUID
	DeviceID            *uuid.UUID
	RefreshTokenHash    string
	IPAddress           *string
	UserAgent           *string
	CurrentUsedAt       time.Time
	NewSessionExpiresAt time.Time
}

func NewSessionRepository(db dbsqlc.DBTX) *SessionRepository {
	return &SessionRepository{db: db, queries: dbsqlc.New(db)}
}

func (r *SessionRepository) CreateSession(ctx context.Context, params CreateSessionParams) error {
	_, err := r.queries.CreateUserSession(ctx, dbsqlc.CreateUserSessionParams{
		ID:               params.ID,
		UserID:           params.UserID,
		RefreshTokenHash: params.RefreshTokenHash,
		DeviceID:         params.DeviceID,
		IpAddress:        params.IPAddress,
		UserAgent:        textValue(params.UserAgent),
		ExpiresAt:        pgtype.Timestamptz{Time: params.ExpiresAt, Valid: true},
	})
	if err != nil {
		return fmt.Errorf("create user session: %w", err)
	}
	return nil
}

func (r *SessionRepository) FindByRefreshTokenHash(ctx context.Context, refreshTokenHash string) (Session, error) {
	session, err := r.queries.GetUserSessionByRefreshHash(ctx, refreshTokenHash)
	if err != nil {
		return Session{}, fmt.Errorf("find user session by refresh token hash: %w", err)
	}
	return Session{
		ID:         session.ID,
		UserID:     session.UserID,
		DeviceID:   session.DeviceID,
		ExpiresAt:  session.ExpiresAt.Time,
		RevokedAt:  nullableTime(session.RevokedAt),
		LastUsedAt: nullableTime(session.LastUsedAt),
	}, nil
}

func (r *SessionRepository) RotateSession(ctx context.Context, params RotateSessionParams) error {
	beginner, ok := r.db.(interface {
		Begin(context.Context) (pgx.Tx, error)
	})
	if !ok {
		return fmt.Errorf("rotate user session: database does not support transactions")
	}
	tx, err := beginner.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin session rotation: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	queries := r.queries.WithTx(tx)
	rotated, err := queries.RevokeUserSessionForRotation(ctx, dbsqlc.RevokeUserSessionForRotationParams{
		ID:        params.CurrentSessionID,
		RevokedAt: pgtype.Timestamptz{Time: params.CurrentUsedAt, Valid: true},
	})
	if err != nil {
		return fmt.Errorf("revoke user session for rotation: %w", err)
	}
	if rotated.UserID != params.UserID || !sameUUID(rotated.DeviceID, params.DeviceID) {
		return fmt.Errorf("rotate user session: session ownership changed")
	}

	_, err = queries.CreateUserSession(ctx, dbsqlc.CreateUserSessionParams{
		ID:               params.NewSessionID,
		UserID:           params.UserID,
		RefreshTokenHash: params.RefreshTokenHash,
		DeviceID:         params.DeviceID,
		IpAddress:        params.IPAddress,
		UserAgent:        textValue(params.UserAgent),
		ExpiresAt:        pgtype.Timestamptz{Time: params.NewSessionExpiresAt, Valid: true},
	})
	if err != nil {
		return fmt.Errorf("create rotated user session: %w", err)
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit session rotation: %w", err)
	}
	return nil
}

func sameUUID(left, right *uuid.UUID) bool {
	if left == nil || right == nil {
		return left == nil && right == nil
	}
	return *left == *right
}

func textValue(value *string) pgtype.Text {
	if value == nil {
		return pgtype.Text{}
	}
	return pgtype.Text{String: *value, Valid: true}
}
