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

## Login RPC

Identity Service exposes `schoolplatform.identity.v1.IdentityService/Login` over gRPC. It verifies credentials, rejects inactive accounts, issues Ed25519 access tokens, and stores only the SHA-256 hash of the opaque refresh token.

JWT private keys must be supplied through `JWT_PRIVATE_KEY_PATH` and must never be committed.

Refresh token rotation revokes the consumed session row, updates `last_used_at`, and creates a new session containing only the new token hash. Reused, revoked, and expired tokens are rejected.

Logout validates the actor's Ed25519 access token, verifies that the submitted refresh token belongs to that actor, and revokes the current session. A revoked refresh token cannot be used again.

## Authorization Baseline

Migration `000002_seed_roles_permissions.sql` idempotently seeds the seven MVP roles and their `domain.resource.action` permission baseline. `wali_kelas` is represented by a scoped teacher assignment rather than a separate role.

`AuthorizationRepository` creates foundation, school, class, subject, and student-scoped role assignments. Assignment creation and the `identity.role.assigned` audit record are committed atomically. User context queries return active assignments together with deduplicated roles and permissions.
