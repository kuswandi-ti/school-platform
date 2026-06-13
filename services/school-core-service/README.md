# School Core Service

Owns foundation, school, academic year, semester, student master, guardian or parent master, teacher master, class, student-class assignment, teacher assignment, and homeroom assignment data.

## Database Migrations

The service owns Goose migrations under `internal/db/migrations` and may only run them against `school_core_db`.

```bash
goose -dir internal/db/migrations postgres "$SCHOOL_CORE_DATABASE_URL" up
goose -dir internal/db/migrations postgres "$SCHOOL_CORE_DATABASE_URL" down
```

The schema uses foreign keys only for data owned by School Core. Identity, Academic, and file-service identifiers remain UUID references without cross-service foreign keys.

Migration `000002_create_outbox_events.sql` provides the transactional outbox used for School Core domain events.

## Migration Tests

The integration test creates and removes an isolated temporary schema in the configured PostgreSQL database.

```bash
SCHOOL_CORE_TEST_DATABASE_URL="$SCHOOL_CORE_DATABASE_URL" go test ./... -count=1
```

## Foundation and School API

The service exposes internal gRPC operations for current-foundation lookup and foundation-scoped school list, create, and update. Public HTTP access remains behind API Gateway:

```text
GET   /api/v1/foundations/current
GET   /api/v1/schools
POST  /api/v1/schools
PATCH /api/v1/schools/{school_id}
```

All operations require an authenticated `admin_yayasan` actor in the target foundation. School creation writes its audit record and `school.school.created` outbox event in the same transaction.

## SQLC

Queries are defined under `internal/db/queries` and generated into `internal/db/sqlc`.

```bash
sqlc generate
```
