package tests

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

	"school-platform/services/school-core-service/internal/domain"
	"school-platform/services/school-core-service/internal/repository"
)

func TestCreateSchoolPersistsAuditAndOutboxAtomically(t *testing.T) {
	databaseURL := os.Getenv("SCHOOL_CORE_TEST_DATABASE_URL")
	if databaseURL == "" {
		t.Skip("SCHOOL_CORE_TEST_DATABASE_URL is required for repository integration tests")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	adminDB, err := sql.Open("pgx", databaseURL)
	require.NoError(t, err)
	defer adminDB.Close()
	schemaName := fmt.Sprintf("school_repository_test_%d", time.Now().UnixNano())
	_, err = adminDB.ExecContext(ctx, "CREATE SCHEMA "+schemaName)
	require.NoError(t, err)
	defer func() {
		_, _ = adminDB.ExecContext(context.Background(), "DROP SCHEMA IF EXISTS "+schemaName+" CASCADE")
	}()

	config, err := pgx.ParseConfig(databaseURL)
	require.NoError(t, err)
	config.RuntimeParams["search_path"] = schemaName
	testDB := stdlib.OpenDB(*config)
	defer testDB.Close()
	require.NoError(t, goose.SetDialect("postgres"))
	require.NoError(t, goose.Up(testDB, filepath.Join("..", "internal", "db", "migrations")))

	foundationID := uuid.New()
	_, err = testDB.ExecContext(ctx, `INSERT INTO foundations (id, foundation_code, name, status) VALUES ($1, 'TEST', 'Test Foundation', 'active')`, foundationID)
	require.NoError(t, err)

	poolConfig, err := pgxpool.ParseConfig(databaseURL)
	require.NoError(t, err)
	poolConfig.ConnConfig.RuntimeParams["search_path"] = schemaName
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	require.NoError(t, err)
	defer pool.Close()

	repo := repository.NewSchoolRepository(pool)
	schoolID := uuid.New()
	actor := domain.Actor{UserID: uuid.New(), FoundationID: foundationID, Roles: []string{"admin_yayasan"}, RequestID: "request-test", CorrelationID: "correlation-test"}
	school, err := repo.CreateSchool(ctx, repository.SchoolWrite{ID: schoolID, FoundationID: foundationID, SchoolCode: "SD", Name: "SD Test", SchoolLevel: "elementary", Status: "active"}, repository.AuditContext{Actor: actor, NewValues: map[string]any{"school_code": "SD"}, Action: "school.school.created"})
	require.NoError(t, err)
	require.Equal(t, schoolID, school.ID)

	var auditCount, outboxCount int
	var eventType, payloadSchoolCode string
	require.NoError(t, testDB.QueryRowContext(ctx, `SELECT COUNT(*) FROM audit_logs WHERE entity_id = $1 AND action = 'school.school.created'`, schoolID).Scan(&auditCount))
	require.NoError(t, testDB.QueryRowContext(ctx, `SELECT COUNT(*) FROM outbox_events WHERE aggregate_id = $1 AND event_type = 'school.school.created' AND status = 'pending'`, schoolID).Scan(&outboxCount))
	require.NoError(t, testDB.QueryRowContext(ctx, `SELECT payload_json->>'event_type', payload_json#>>'{payload,school_code}' FROM outbox_events WHERE aggregate_id = $1`, schoolID).Scan(&eventType, &payloadSchoolCode))
	require.Equal(t, 1, auditCount)
	require.Equal(t, 1, outboxCount)
	require.Equal(t, "school.school.created", eventType)
	require.Equal(t, "SD", payloadSchoolCode)
}
