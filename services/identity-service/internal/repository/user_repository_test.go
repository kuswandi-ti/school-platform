package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"

	"school-platform/services/identity-service/internal/password"
	"school-platform/services/identity-service/internal/repository"
)

func TestUserRepository(t *testing.T) {
	repo, ctx := newTestRepository(t)

	passwordHash, err := password.Hash("repository test password")
	require.NoError(t, err)

	userID := uuid.New()
	created, err := repo.CreateUser(ctx, repository.CreateUserParams{
		ID:           userID,
		Email:        " User.One@Example.com ",
		PasswordHash: passwordHash,
		DisplayName:  "User One",
		Status:       "active",
	})
	require.NoError(t, err)
	require.Equal(t, userID, created.ID)
	require.Equal(t, "User.One@Example.com", created.Email)
	require.Equal(t, passwordHash, created.PasswordHash)
	require.Nil(t, created.LastLoginAt)

	byEmail, err := repo.FindByEmail(ctx, "user.one@example.com")
	require.NoError(t, err)
	require.Equal(t, userID, byEmail.ID)

	byID, err := repo.FindByID(ctx, userID)
	require.NoError(t, err)
	require.Equal(t, "active", byID.Status)

	lastLoginAt := time.Now().UTC().Truncate(time.Microsecond)
	updatedLogin, err := repo.UpdateLastLogin(ctx, userID, lastLoginAt)
	require.NoError(t, err)
	require.NotNil(t, updatedLogin.LastLoginAt)
	require.WithinDuration(t, lastLoginAt, *updatedLogin.LastLoginAt, time.Microsecond)

	inactive, err := repo.UpdateStatus(ctx, userID, "inactive")
	require.NoError(t, err)
	require.Equal(t, "inactive", inactive.Status)

	storedInactive, err := repo.FindByID(ctx, userID)
	require.NoError(t, err)
	require.Equal(t, "inactive", storedInactive.Status)
}

func TestUserRepositoryRejectsDuplicateEmailIgnoringCase(t *testing.T) {
	repo, ctx := newTestRepository(t)
	passwordHash, err := password.Hash("repository test password")
	require.NoError(t, err)

	_, err = repo.CreateUser(ctx, repository.CreateUserParams{
		ID:           uuid.New(),
		Email:        "duplicate@example.com",
		PasswordHash: passwordHash,
		DisplayName:  "First User",
		Status:       "active",
	})
	require.NoError(t, err)

	_, err = repo.CreateUser(ctx, repository.CreateUserParams{
		ID:           uuid.New(),
		Email:        "DUPLICATE@example.com",
		PasswordHash: passwordHash,
		DisplayName:  "Second User",
		Status:       "active",
	})
	require.Error(t, err)
}

func TestUserRepositoryReturnsNotFound(t *testing.T) {
	repo, ctx := newTestRepository(t)

	_, err := repo.FindByID(ctx, uuid.New())
	require.ErrorIs(t, err, pgx.ErrNoRows)

	_, err = repo.FindByEmail(ctx, "missing@example.com")
	require.ErrorIs(t, err, pgx.ErrNoRows)
}

func newTestRepository(t *testing.T) (*repository.UserRepository, context.Context) {
	t.Helper()
	databaseURL := os.Getenv("IDENTITY_TEST_DATABASE_URL")
	if databaseURL == "" {
		t.Skip("IDENTITY_TEST_DATABASE_URL is required for repository integration tests")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	t.Cleanup(cancel)

	adminDB, err := sql.Open("pgx", databaseURL)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, adminDB.Close())
	})
	require.NoError(t, adminDB.PingContext(ctx))

	schemaName := fmt.Sprintf("identity_repository_test_%d", time.Now().UnixNano())
	_, err = adminDB.ExecContext(ctx, "CREATE SCHEMA "+schemaName)
	require.NoError(t, err)
	t.Cleanup(func() {
		_, dropErr := adminDB.ExecContext(context.Background(), "DROP SCHEMA IF EXISTS "+schemaName+" CASCADE")
		require.NoError(t, dropErr)
	})

	config, err := pgx.ParseConfig(databaseURL)
	require.NoError(t, err)
	config.RuntimeParams["search_path"] = schemaName
	migrationDB := stdlib.OpenDB(*config)
	t.Cleanup(func() {
		require.NoError(t, migrationDB.Close())
	})
	require.NoError(t, migrationDB.PingContext(ctx))

	require.NoError(t, goose.SetDialect("postgres"))
	migrationsDir := filepath.Join("..", "db", "migrations")
	require.NoError(t, goose.Up(migrationDB, migrationsDir))

	poolConfig, err := pgxpool.ParseConfig(databaseURL)
	require.NoError(t, err)
	poolConfig.ConnConfig.RuntimeParams["search_path"] = schemaName
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	require.NoError(t, err)
	t.Cleanup(pool.Close)
	require.NoError(t, pool.Ping(ctx))

	return repository.NewUserRepository(pool), ctx
}
