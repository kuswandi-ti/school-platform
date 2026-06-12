# PostgreSQL

Local PostgreSQL configuration for `school-platform`.

The MVP uses one PostgreSQL instance with one database per service. The local Docker Compose setup mounts `infra/postgres/init` into the PostgreSQL container so service databases are created when the `postgres_data` volume is initialized for the first time.

## Local Databases

The init script creates:

```text
identity_db
school_core_db
admission_db
academic_db
finance_db
communication_db
reporting_db
```

The default maintenance database is:

```text
school_platform_local
```

## Local Credentials

Defaults are defined in `.env.example`:

```text
POSTGRES_USER=school_local
POSTGRES_PASSWORD=school_local_password
```

These credentials are for local development only. Do not use them in staging or production.

## Verify Locally

When Docker Compose is running:

```bash
docker compose ps postgres
docker compose logs postgres
```

Example `psql` connection:

```bash
psql "postgres://school_local:school_local_password@localhost:5432/identity_db?sslmode=disable"
```

If `psql` is not installed on the host, use the container:

```bash
docker compose exec postgres psql -U school_local -d identity_db -c "SELECT current_database();"
```

List databases:

```bash
psql "postgres://school_local:school_local_password@localhost:5432/school_platform_local?sslmode=disable" -c "\l"
```

Container-only equivalent:

```bash
docker compose exec postgres psql -U school_local -d school_platform_local -c "\l"
```

If databases are not created after changing the init script, remove the local `postgres_data` Docker volume and start again. Do not do this if the local volume contains data you still need.
