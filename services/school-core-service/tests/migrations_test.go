package tests

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"
)

func TestSchoolCoreMigrationsUpAndDown(t *testing.T) {
	databaseURL := os.Getenv("SCHOOL_CORE_TEST_DATABASE_URL")
	if databaseURL == "" {
		t.Skip("SCHOOL_CORE_TEST_DATABASE_URL is required for migration integration tests")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	adminDB, err := sql.Open("pgx", databaseURL)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, adminDB.Close()) })
	require.NoError(t, adminDB.PingContext(ctx))

	schemaName := fmt.Sprintf("school_core_migration_test_%d", time.Now().UnixNano())
	_, err = adminDB.ExecContext(ctx, "CREATE SCHEMA "+schemaName)
	require.NoError(t, err)
	t.Cleanup(func() {
		_, _ = adminDB.ExecContext(context.Background(), "DROP SCHEMA IF EXISTS "+schemaName+" CASCADE")
	})

	config, err := pgx.ParseConfig(databaseURL)
	require.NoError(t, err)
	config.RuntimeParams["search_path"] = schemaName
	testDB := stdlib.OpenDB(*config)
	t.Cleanup(func() { require.NoError(t, testDB.Close()) })
	require.NoError(t, testDB.PingContext(ctx))

	require.NoError(t, goose.SetDialect("postgres"))
	migrationsDir := filepath.Join("..", "internal", "db", "migrations")
	require.NoError(t, goose.Up(testDB, migrationsDir))

	expectedTables := []string{
		"academic_years", "audit_logs", "classes", "foundations", "grade_levels",
		"guardians", "homeroom_assignments", "rooms", "schools", "semesters",
		"student_class_assignments", "student_guardians", "students",
		"teacher_assignments", "teachers",
	}
	assertTables(t, ctx, testDB, expectedTables)
	assertTenantColumns(t, ctx, testDB, expectedTables)

	assertConstraints(t, ctx, testDB, []string{
		"foundations_foundation_code_unique",
		"schools_foundation_school_code_unique",
		"classes_scope_class_code_unique",
		"student_guardians_student_guardian_unique",
	})
	assertIndexes(t, ctx, testDB, []string{
		"academic_years_one_active_per_foundation_idx",
		"classes_foundation_school_academic_year_idx",
		"guardians_email_idx",
		"guardians_phone_idx",
		"homeroom_assignments_one_active_per_class_idx",
		"schools_foundation_id_idx",
		"student_class_assignments_class_id_idx",
		"student_class_assignments_one_active_idx",
		"student_class_assignments_student_id_idx",
		"students_foundation_school_idx",
		"students_full_name_idx",
		"students_nisn_idx",
		"students_status_idx",
		"students_student_number_idx",
		"students_student_number_unique_idx",
		"teachers_foundation_school_idx",
		"teachers_full_name_idx",
	})
	assertUniqueIndexes(t, ctx, testDB, []string{
		"academic_years_one_active_per_foundation_idx",
		"homeroom_assignments_one_active_per_class_idx",
		"student_class_assignments_one_active_idx",
		"students_nisn_unique_idx",
		"students_student_number_unique_idx",
	})
	assertNoCrossServiceForeignKeys(t, ctx, testDB)

	require.NoError(t, goose.DownTo(testDB, migrationsDir, 0))
	assertTablesMissing(t, ctx, testDB, expectedTables)
}

func assertTables(t *testing.T, ctx context.Context, db *sql.DB, expected []string) {
	t.Helper()
	rows, err := db.QueryContext(ctx, `
		SELECT tablename
		FROM pg_catalog.pg_tables
		WHERE schemaname = current_schema() AND tablename <> 'goose_db_version'
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

func assertTenantColumns(t *testing.T, ctx context.Context, db *sql.DB, tables []string) {
	t.Helper()
	for _, tableName := range tables {
		if tableName == "foundations" {
			continue
		}
		var count int
		err := db.QueryRowContext(ctx, `
			SELECT COUNT(*)
			FROM information_schema.columns
			WHERE table_schema = current_schema()
			  AND table_name = $1
			  AND column_name = 'foundation_id'
			  AND is_nullable = 'NO'`, tableName).Scan(&count)
		require.NoError(t, err)
		require.Equal(t, 1, count, "%s must have a non-null foundation_id", tableName)
	}

	for _, tableName := range []string{
		"audit_logs", "classes", "grade_levels", "homeroom_assignments", "rooms",
		"student_class_assignments", "students", "teacher_assignments", "teachers",
	} {
		var count int
		err := db.QueryRowContext(ctx, `
			SELECT COUNT(*)
			FROM information_schema.columns
			WHERE table_schema = current_schema()
			  AND table_name = $1
			  AND column_name = 'school_id'`, tableName).Scan(&count)
		require.NoError(t, err)
		require.Equal(t, 1, count, "%s must have school_id", tableName)
	}
}

func assertConstraints(t *testing.T, ctx context.Context, db *sql.DB, names []string) {
	t.Helper()
	for _, name := range names {
		var exists bool
		err := db.QueryRowContext(ctx, `
			SELECT EXISTS (
				SELECT 1
				FROM pg_catalog.pg_constraint c
				JOIN pg_catalog.pg_namespace n ON n.oid = c.connamespace
				WHERE n.nspname = current_schema() AND c.conname = $1
			)`, name).Scan(&exists)
		require.NoError(t, err)
		require.True(t, exists, "constraint %s does not exist", name)
	}
}

func assertIndexes(t *testing.T, ctx context.Context, db *sql.DB, names []string) {
	t.Helper()
	for _, name := range names {
		var exists bool
		err := db.QueryRowContext(ctx, `
			SELECT EXISTS (
				SELECT 1 FROM pg_catalog.pg_indexes
				WHERE schemaname = current_schema() AND indexname = $1
			)`, name).Scan(&exists)
		require.NoError(t, err)
		require.True(t, exists, "index %s does not exist", name)
	}
}

func assertUniqueIndexes(t *testing.T, ctx context.Context, db *sql.DB, names []string) {
	t.Helper()
	for _, name := range names {
		var isUnique bool
		err := db.QueryRowContext(ctx, `
			SELECT i.indisunique
			FROM pg_catalog.pg_index i
			JOIN pg_catalog.pg_class c ON c.oid = i.indexrelid
			JOIN pg_catalog.pg_namespace n ON n.oid = c.relnamespace
			WHERE n.nspname = current_schema() AND c.relname = $1`, name).Scan(&isUnique)
		require.NoError(t, err)
		require.True(t, isUnique, "index %s is not unique", name)
	}
}

func assertNoCrossServiceForeignKeys(t *testing.T, ctx context.Context, db *sql.DB) {
	t.Helper()
	for _, columnName := range []string{
		"actor_user_id", "approved_by", "logo_file_id", "photo_file_id", "subject_id", "user_id",
	} {
		var count int
		err := db.QueryRowContext(ctx, `
			SELECT COUNT(*)
			FROM pg_catalog.pg_constraint c
			JOIN pg_catalog.pg_class r ON r.oid = c.conrelid
			JOIN pg_catalog.pg_namespace n ON n.oid = r.relnamespace
			JOIN unnest(c.conkey) AS key(attnum) ON TRUE
			JOIN pg_catalog.pg_attribute a ON a.attrelid = r.oid AND a.attnum = key.attnum
			WHERE n.nspname = current_schema()
			  AND c.contype = 'f'
			  AND a.attname = $1`, columnName).Scan(&count)
		require.NoError(t, err)
		require.Zero(t, count, "%s must remain a cross-service reference without a foreign key", columnName)
	}
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
