package tests

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"
)

func TestIdentityMigrationsUpAndDown(t *testing.T) {
	databaseURL := os.Getenv("IDENTITY_TEST_DATABASE_URL")
	if databaseURL == "" {
		t.Skip("IDENTITY_TEST_DATABASE_URL is required for migration integration tests")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	adminDB, err := sql.Open("pgx", databaseURL)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, adminDB.Close())
	})
	require.NoError(t, adminDB.PingContext(ctx))

	schemaName := fmt.Sprintf("identity_migration_test_%d", time.Now().UnixNano())
	_, err = adminDB.ExecContext(ctx, "CREATE SCHEMA "+schemaName)
	require.NoError(t, err)
	t.Cleanup(func() {
		_, _ = adminDB.ExecContext(context.Background(), "DROP SCHEMA IF EXISTS "+schemaName+" CASCADE")
	})

	config, err := pgx.ParseConfig(databaseURL)
	require.NoError(t, err)
	config.RuntimeParams["search_path"] = schemaName
	testDB := stdlib.OpenDB(*config)
	defer testDB.Close()
	require.NoError(t, testDB.PingContext(ctx))

	require.NoError(t, goose.SetDialect("postgres"))
	migrationsDir := filepath.Join("..", "internal", "db", "migrations")
	require.NoError(t, goose.Up(testDB, migrationsDir))

	assertTables(t, ctx, testDB, []string{
		"identity_audit_logs",
		"permissions",
		"role_permissions",
		"roles",
		"user_devices",
		"user_role_assignments",
		"user_sessions",
		"users",
	})
	assertColumn(t, ctx, testDB, "user_sessions", "refresh_token_hash", false)
	assertColumnMissing(t, ctx, testDB, "user_sessions", "refresh_token")
	assertIndexes(t, ctx, testDB, []string{
		"identity_audit_logs_foundation_id_idx",
		"identity_audit_logs_school_id_idx",
		"role_permissions_role_id_idx",
		"user_role_assignments_foundation_id_idx",
		"user_role_assignments_role_id_idx",
		"user_role_assignments_school_id_idx",
		"user_role_assignments_user_id_idx",
		"user_sessions_user_id_idx",
		"users_email_unique_idx",
		"users_status_idx",
	})
	assertUniqueIndex(t, ctx, testDB, "users_email_unique_idx", "")
	assertUniqueIndex(t, ctx, testDB, "user_role_assignments_active_scope_unique_idx", "NULLS NOT DISTINCT")
	assertConstraint(t, ctx, testDB, "role_permissions", "role_permissions_role_permission_unique")
	assertConstraint(t, ctx, testDB, "user_sessions", "user_sessions_refresh_token_hash_unique")
	assertConstraint(t, ctx, testDB, "user_devices", "user_devices_user_device_unique")

	require.NoError(t, goose.DownTo(testDB, migrationsDir, 0))
	assertTablesMissing(t, ctx, testDB, []string{
		"identity_audit_logs",
		"permissions",
		"role_permissions",
		"roles",
		"user_devices",
		"user_role_assignments",
		"user_sessions",
		"users",
	})
}

func assertUniqueIndex(t *testing.T, ctx context.Context, db *sql.DB, indexName, expectedDefinition string) {
	t.Helper()
	var isUnique bool
	var definition string
	err := db.QueryRowContext(ctx, `
		SELECT i.indisunique, pg_catalog.pg_get_indexdef(i.indexrelid)
		FROM pg_catalog.pg_index i
		JOIN pg_catalog.pg_class c ON c.oid = i.indexrelid
		JOIN pg_catalog.pg_namespace n ON n.oid = c.relnamespace
		WHERE n.nspname = current_schema() AND c.relname = $1`, indexName).Scan(&isUnique, &definition)
	require.NoError(t, err)
	require.True(t, isUnique, "index %s is not unique", indexName)
	if expectedDefinition != "" {
		require.Contains(t, definition, expectedDefinition)
	}
}

func assertTables(t *testing.T, ctx context.Context, db *sql.DB, expected []string) {
	t.Helper()
	rows, err := db.QueryContext(ctx, `
		SELECT tablename
		FROM pg_catalog.pg_tables
		WHERE schemaname = current_schema()
		  AND tablename <> 'goose_db_version'
		ORDER BY tablename`)
	require.NoError(t, err)
	defer rows.Close()

	var actual []string
	for rows.Next() {
		var tableName string
		require.NoError(t, rows.Scan(&tableName))
		actual = append(actual, tableName)
	}
	require.NoError(t, rows.Err())
	sort.Strings(expected)
	require.Equal(t, expected, actual)
}

func assertTablesMissing(t *testing.T, ctx context.Context, db *sql.DB, tables []string) {
	t.Helper()
	for _, tableName := range tables {
		var exists bool
		err := db.QueryRowContext(ctx, `
			SELECT EXISTS (
				SELECT 1 FROM pg_catalog.pg_tables
				WHERE schemaname = current_schema() AND tablename = $1
			)`, tableName).Scan(&exists)
		require.NoError(t, err)
		require.False(t, exists, "table %s still exists after down migration", tableName)
	}
}

func assertColumn(t *testing.T, ctx context.Context, db *sql.DB, tableName, columnName string, nullable bool) {
	t.Helper()
	var isNullable string
	err := db.QueryRowContext(ctx, `
		SELECT is_nullable
		FROM information_schema.columns
		WHERE table_schema = current_schema()
		  AND table_name = $1
		  AND column_name = $2`, tableName, columnName).Scan(&isNullable)
	require.NoError(t, err)
	require.Equal(t, nullable, strings.EqualFold(isNullable, "YES"))
}

func assertColumnMissing(t *testing.T, ctx context.Context, db *sql.DB, tableName, columnName string) {
	t.Helper()
	var count int
	err := db.QueryRowContext(ctx, `
		SELECT COUNT(*)
		FROM information_schema.columns
		WHERE table_schema = current_schema()
		  AND table_name = $1
		  AND column_name = $2`, tableName, columnName).Scan(&count)
	require.NoError(t, err)
	require.Zero(t, count)
}

func assertIndexes(t *testing.T, ctx context.Context, db *sql.DB, indexNames []string) {
	t.Helper()
	for _, indexName := range indexNames {
		var exists bool
		err := db.QueryRowContext(ctx, `
			SELECT EXISTS (
				SELECT 1 FROM pg_catalog.pg_indexes
				WHERE schemaname = current_schema() AND indexname = $1
			)`, indexName).Scan(&exists)
		require.NoError(t, err)
		require.True(t, exists, "index %s does not exist", indexName)
	}
}

func assertConstraint(t *testing.T, ctx context.Context, db *sql.DB, tableName, constraintName string) {
	t.Helper()
	var exists bool
	err := db.QueryRowContext(ctx, `
		SELECT EXISTS (
			SELECT 1
			FROM pg_catalog.pg_constraint c
			JOIN pg_catalog.pg_class r ON r.oid = c.conrelid
			JOIN pg_catalog.pg_namespace n ON n.oid = r.relnamespace
			WHERE n.nspname = current_schema()
			  AND r.relname = $1
			  AND c.conname = $2
		)`, tableName, constraintName).Scan(&exists)
	require.NoError(t, err)
	require.True(t, exists, "constraint %s does not exist", constraintName)
}
