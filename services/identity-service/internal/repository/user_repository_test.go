package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sync"
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
	repo, _, ctx := newTestRepository(t)

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
	repo, _, ctx := newTestRepository(t)
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
	repo, _, ctx := newTestRepository(t)

	_, err := repo.FindByID(ctx, uuid.New())
	require.ErrorIs(t, err, pgx.ErrNoRows)

	_, err = repo.FindByEmail(ctx, "missing@example.com")
	require.ErrorIs(t, err, pgx.ErrNoRows)
}

func TestSessionRepositoryStoresRefreshTokenHash(t *testing.T) {
	users, pool, ctx := newTestRepository(t)
	passwordHash, err := password.Hash("repository test password")
	require.NoError(t, err)
	userID := uuid.New()
	_, err = users.CreateUser(ctx, repository.CreateUserParams{
		ID: userID, Email: "session@example.com", PasswordHash: passwordHash,
		DisplayName: "Session User", Status: "active",
	})
	require.NoError(t, err)

	refreshTokenHash := "stored-refresh-token-hash"
	err = repository.NewSessionRepository(pool).CreateSession(ctx, repository.CreateSessionParams{
		ID: uuid.New(), UserID: userID, RefreshTokenHash: refreshTokenHash,
		ExpiresAt: time.Now().UTC().Add(24 * time.Hour),
	})
	require.NoError(t, err)

	var storedHash string
	require.NoError(t, pool.QueryRow(ctx, "SELECT refresh_token_hash FROM user_sessions WHERE user_id = $1", userID).Scan(&storedHash))
	require.Equal(t, refreshTokenHash, storedHash)
}

func TestSessionRepositoryRotatesTokenOnce(t *testing.T) {
	users, pool, ctx := newTestRepository(t)
	passwordHash, err := password.Hash("repository test password")
	require.NoError(t, err)
	userID := uuid.New()
	_, err = users.CreateUser(ctx, repository.CreateUserParams{
		ID: userID, Email: "rotate@example.com", PasswordHash: passwordHash,
		DisplayName: "Rotate User", Status: "active",
	})
	require.NoError(t, err)

	sessions := repository.NewSessionRepository(pool)
	currentSessionID := uuid.New()
	oldHash := "old-refresh-token-hash"
	err = sessions.CreateSession(ctx, repository.CreateSessionParams{
		ID: currentSessionID, UserID: userID, RefreshTokenHash: oldHash,
		ExpiresAt: time.Now().UTC().Add(time.Hour),
	})
	require.NoError(t, err)

	usedAt := time.Now().UTC().Truncate(time.Microsecond)
	rotate := func(newHash string) error {
		return sessions.RotateSession(ctx, repository.RotateSessionParams{
			CurrentSessionID: currentSessionID,
			NewSessionID:     uuid.New(), UserID: userID, RefreshTokenHash: newHash,
			CurrentUsedAt: usedAt, NewSessionExpiresAt: usedAt.Add(24 * time.Hour),
		})
	}

	errorsCh := make(chan error, 2)
	var waitGroup sync.WaitGroup
	for _, newHash := range []string{"new-refresh-hash-one", "new-refresh-hash-two"} {
		waitGroup.Add(1)
		go func(hash string) {
			defer waitGroup.Done()
			errorsCh <- rotate(hash)
		}(newHash)
	}
	waitGroup.Wait()
	close(errorsCh)

	var successCount, failureCount int
	for rotateErr := range errorsCh {
		if rotateErr == nil {
			successCount++
		} else {
			failureCount++
			require.ErrorIs(t, rotateErr, pgx.ErrNoRows)
		}
	}
	require.Equal(t, 1, successCount)
	require.Equal(t, 1, failureCount)

	oldSession, err := sessions.FindByRefreshTokenHash(ctx, oldHash)
	require.NoError(t, err)
	require.NotNil(t, oldSession.RevokedAt)
	require.NotNil(t, oldSession.LastUsedAt)
	require.WithinDuration(t, usedAt, *oldSession.LastUsedAt, time.Microsecond)

	var activeSessions int
	require.NoError(t, pool.QueryRow(ctx, `SELECT COUNT(*) FROM user_sessions WHERE user_id = $1 AND revoked_at IS NULL`, userID).Scan(&activeSessions))
	require.Equal(t, 1, activeSessions)
}

func newTestRepository(t *testing.T) (*repository.UserRepository, *pgxpool.Pool, context.Context) {
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

	return repository.NewUserRepository(pool), pool, ctx
}
