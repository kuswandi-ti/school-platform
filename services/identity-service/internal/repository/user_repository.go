package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"school-platform/services/identity-service/internal/db/sqlc"
	"school-platform/services/identity-service/internal/domain"
)

type UserRepository struct {
	queries *dbsqlc.Queries
}

type CreateUserParams struct {
	ID           uuid.UUID
	Email        string
	Phone        *string
	PasswordHash string
	DisplayName  string
	AvatarFileID *uuid.UUID
	Status       string
}

func NewUserRepository(db dbsqlc.DBTX) *UserRepository {
	return &UserRepository{queries: dbsqlc.New(db)}
}

func (r *UserRepository) CreateUser(ctx context.Context, params CreateUserParams) (domain.User, error) {
	user, err := r.queries.CreateUser(ctx, dbsqlc.CreateUserParams{
		ID:           params.ID,
		Email:        strings.TrimSpace(params.Email),
		Phone:        params.Phone,
		PasswordHash: params.PasswordHash,
		DisplayName:  params.DisplayName,
		AvatarFileID: params.AvatarFileID,
		Status:       params.Status,
	})
	if err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}
	return mapUser(user), nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	user, err := r.queries.GetUserByEmail(ctx, strings.TrimSpace(email))
	if err != nil {
		return domain.User{}, fmt.Errorf("find user by email: %w", err)
	}
	return mapUser(user), nil
}

func (r *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	user, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("find user by id: %w", err)
	}
	return mapUser(user), nil
}

func (r *UserRepository) UpdateLastLogin(ctx context.Context, id uuid.UUID, lastLoginAt time.Time) (domain.User, error) {
	user, err := r.queries.UpdateUserLastLogin(ctx, dbsqlc.UpdateUserLastLoginParams{
		ID:          id,
		LastLoginAt: pgtype.Timestamptz{Time: lastLoginAt, Valid: true},
	})
	if err != nil {
		return domain.User{}, fmt.Errorf("update user last login: %w", err)
	}
	return mapUser(user), nil
}

func (r *UserRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status string) (domain.User, error) {
	user, err := r.queries.UpdateUserStatus(ctx, dbsqlc.UpdateUserStatusParams{
		ID:     id,
		Status: status,
	})
	if err != nil {
		return domain.User{}, fmt.Errorf("update user status: %w", err)
	}
	return mapUser(user), nil
}

func mapUser(user dbsqlc.User) domain.User {
	return domain.User{
		ID:           user.ID,
		Email:        user.Email,
		Phone:        user.Phone,
		PasswordHash: user.PasswordHash,
		DisplayName:  user.DisplayName,
		AvatarFileID: user.AvatarFileID,
		Status:       user.Status,
		LastLoginAt:  nullableTime(user.LastLoginAt),
		CreatedAt:    user.CreatedAt.Time,
		UpdatedAt:    user.UpdatedAt.Time,
	}
}

func nullableTime(value pgtype.Timestamptz) *time.Time {
	if !value.Valid {
		return nil
	}
	return &value.Time
}
