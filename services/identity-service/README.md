# Identity Service

Owns user accounts, credentials, sessions, refresh tokens, roles, permissions, role assignments, devices, and identity audit logs.

## Database Migrations

The service owns Goose migrations under `internal/db/migrations` and may only run them against `identity_db`.

```bash
goose -dir internal/db/migrations postgres "$IDENTITY_DATABASE_URL" up
goose -dir internal/db/migrations postgres "$IDENTITY_DATABASE_URL" down
```

Refresh tokens must never be stored in plain text. The schema stores only `refresh_token_hash`.

## Migration Tests

The integration test creates and removes an isolated temporary schema in the configured PostgreSQL database.

```bash
IDENTITY_TEST_DATABASE_URL="$IDENTITY_DATABASE_URL" go test ./...
```

## SQLC

User persistence queries are defined under `internal/db/queries` and generated into `internal/db/sqlc`.

```bash
sqlc generate
```

Repository integration tests use the same `IDENTITY_TEST_DATABASE_URL` and create isolated temporary schemas.
