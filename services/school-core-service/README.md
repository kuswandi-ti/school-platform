# School Core Service

Owns foundation, school, academic year, semester, student master, guardian or parent master, teacher master, class, student-class assignment, teacher assignment, and homeroom assignment data.

## Database Migrations

The service owns Goose migrations under `internal/db/migrations` and may only run them against `school_core_db`.

```bash
goose -dir internal/db/migrations postgres "$SCHOOL_CORE_DATABASE_URL" up
goose -dir internal/db/migrations postgres "$SCHOOL_CORE_DATABASE_URL" down
```

The schema uses foreign keys only for data owned by School Core. Identity, Academic, and file-service identifiers remain UUID references without cross-service foreign keys.

## Migration Tests

The integration test creates and removes an isolated temporary schema in the configured PostgreSQL database.

```bash
SCHOOL_CORE_TEST_DATABASE_URL="$SCHOOL_CORE_DATABASE_URL" go test ./... -count=1
```
