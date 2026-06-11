# Local Development

This guide is the practical entry point for local setup in `school-platform`.

Local development happens before staging. Do not use staging as a development environment, and never use production data, production credentials, private keys, or real student/parent data locally.

Internal technical naming uses English. User-facing UI labels use Bahasa Indonesia.

## Prerequisites

Install the tools needed for the area you work on.

Core tools:

- Git
- Docker Desktop or Docker Engine
- Docker Compose
- Make

Backend:

- Go 1.23+ or current stable
- goose
- sqlc
- golangci-lint
- protoc
- protoc-gen-go
- protoc-gen-go-grpc

Web admin:

- Node.js 20 LTS+
- pnpm 9+ or npm

Mobile:

- Flutter 3 stable
- Android Studio or equivalent Android SDK setup
- Xcode for iOS development on macOS, when needed

Optional but useful:

- curl
- jq
- psql
- grpcurl
- DBeaver, TablePlus, or pgAdmin
- Postman, Insomnia, Thunder Client, or REST Client

## Repository Setup

```bash
git clone git@github.com:<org>/school-platform.git
cd school-platform
git checkout develop
git pull origin develop
```

Create a task branch:

```bash
git checkout -b feature/<area>-<short-description>
```

Use lowercase English branch names with hyphen separators.

## Environment File Setup

Create local `.env` files from examples when the examples are available:

```bash
cp .env.example .env
cp services/api-gateway/.env.example services/api-gateway/.env
cp services/_template-service/.env.example services/_template-service/.env
```

Repeat the service-level copy for each service that has its own `.env.example`.

Rules:

- Commit `.env.example`, not `.env`.
- Do not commit production credentials.
- Do not commit private keys.
- Do not use production secrets locally.
- Use synthetic or dummy local data only.

## Docker Compose Startup

Docker Compose is used for local dependencies such as PostgreSQL, Redis, RabbitMQ, MinIO, and Mailpit.

Start dependencies with:

```bash
make up
```

or:

```bash
docker compose up -d
```

Check containers:

```bash
docker compose ps
```

View logs:

```bash
make logs
# or
docker compose logs -f
```

Stop dependencies:

```bash
make down
# or
docker compose down
```

## Common Makefile Commands

The project standardizes common commands through `Makefile`.

Expected commands:

```bash
make help
make setup
make up
make down
make ps
make logs
make fmt
make lint
make test
make build
make proto
make openapi-check
make event-schema-check
```

Service-specific commands may follow this pattern:

```bash
make test service=identity-service
make lint service=finance-service
make build service=api-gateway
make migrate-up service=school-core-service
make sqlc service=academic-service
```

If a command is not available yet, check the current Sprint 0 task status before adding new behavior.

## Local Service Ports

Recommended local ports:

| Component | URL / Port |
|---|---|
| API Gateway | `http://localhost:8080` |
| Web Admin | `http://localhost:3000` |
| PostgreSQL | `localhost:5432` |
| Redis | `localhost:6379` |
| RabbitMQ AMQP | `localhost:5672` |
| RabbitMQ Management | `http://localhost:15672` |
| MinIO API | `http://localhost:9000` |
| MinIO Console | `http://localhost:9001` |
| Mailpit SMTP | `localhost:1025` |
| Mailpit UI | `http://localhost:8025` |

Local-only example credentials may be used in `.env.example`, but production credentials must never be stored in the repository.

## Local Databases

The MVP uses one PostgreSQL instance with one database per service:

```text
identity_db
school_core_db
admission_db
academic_db
finance_db
communication_db
reporting_db
```

Each service may access only its own database. Do not query another service database.

## Running Services

Run API Gateway:

```bash
cd services/api-gateway
go run ./cmd/server
```

Run a backend service:

```bash
cd services/identity-service
go run ./cmd/server
```

Run web admin:

```bash
cd apps/web-admin
pnpm install
pnpm dev
```

Run mobile app:

```bash
cd apps/mobile-app
flutter pub get
flutter run
```

Frontend and mobile clients must call the API Gateway, not internal microservices directly.

## Running Tests

Run all available checks from the repository root:

```bash
make fmt
make lint
make test
```

Backend service checks:

```bash
make test service=<service-name>
make lint service=<service-name>
```

Web checks:

```bash
cd apps/web-admin
pnpm lint
pnpm typecheck
pnpm build
```

Mobile checks:

```bash
cd apps/mobile-app
flutter analyze
flutter test
```

Contract checks:

```bash
make proto
make openapi-check
make event-schema-check
```

No automated test is required for documentation-only changes unless the task asks for it.

## Data and Secret Safety

Never commit:

- `.env`
- production credentials
- private keys
- access tokens or refresh tokens
- database dumps
- backup files
- real student, guardian, teacher, or payment data
- raw payment proof files
- Confidential documents

Use:

- dummy data
- synthetic test accounts
- local-only credentials
- `.env.example` files without real secrets

Application logs must not contain passwords, tokens, raw document content, raw payment proof content, or Confidential data.

## Daily Workflow

```bash
git checkout develop
git pull origin develop
git checkout -b feature/<area>-<short-description>

make up

# work on a small scoped task
make fmt
make lint
make test

git status
git add .
git commit -m "type(scope): short description"
git push origin feature/<area>-<short-description>
```

Open a Pull Request into `develop`.

Before PR:

- Keep the task scope small.
- Avoid unrelated file changes.
- Run relevant checks.
- Update docs when contracts, setup, or workflow change.
- Confirm no secrets or production data are included.

## Related Docs

- [README.md](../README.md)
- [docs/24-local-development-guide.md](24-local-development-guide.md)
- [docs/08-coding-standard.md](08-coding-standard.md)
- [docs/09-ai-agent-rules.md](09-ai-agent-rules.md)
- [docs/10-sprint-backlog-mvp.md](10-sprint-backlog-mvp.md)
- [docs/11-github-repository-rules.md](11-github-repository-rules.md)
