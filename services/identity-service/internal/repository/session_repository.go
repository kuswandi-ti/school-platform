package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"school-platform/services/identity-service/internal/db/sqlc"
)

type SessionRepository struct {
	queries *dbsqlc.Queries
}

type CreateSessionParams struct {
	ID               uuid.UUID
	UserID           uuid.UUID
	RefreshTokenHash string
	IPAddress        *string
	UserAgent        *string
	ExpiresAt        time.Time
}

func NewSessionRepository(db dbsqlc.DBTX) *SessionRepository {
	return &SessionRepository{queries: dbsqlc.New(db)}
}

func (r *SessionRepository) CreateSession(ctx context.Context, params CreateSessionParams) error {
	_, err := r.queries.CreateUserSession(ctx, dbsqlc.CreateUserSessionParams{
		ID:               params.ID,
		UserID:           params.UserID,
		RefreshTokenHash: params.RefreshTokenHash,
		IpAddress:        params.IPAddress,
		UserAgent:        textValue(params.UserAgent),
		ExpiresAt:        pgtype.Timestamptz{Time: params.ExpiresAt, Valid: true},
	})
	if err != nil {
		return fmt.Errorf("create user session: %w", err)
	}
	return nil
}

func textValue(value *string) pgtype.Text {
	if value == nil {
		return pgtype.Text{}
	}
	return pgtype.Text{String: *value, Valid: true}
}
