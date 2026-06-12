# 24 — Local Development Guide

Project: `school-platform`  
Status: Development setup guide  
Scope: Local environment preparation, required tools, installation checklist, repository setup, Docker Compose, service commands, frontend/mobile setup, troubleshooting, and daily workflow.

---

## 1. Purpose

This document explains how to prepare a local development environment for `school-platform`.

The goal is to make all development run locally first before changes are pushed to staging or production.

Local development must support:

```text
- Go microservices
- Custom Go API Gateway
- PostgreSQL per service database
- Redis
- RabbitMQ
- MinIO / S3-compatible local storage
- Next.js web admin
- Flutter mobile app
- GitHub workflow
- AI Agent-assisted development
```

---

## 2. Local Development Principles

All developers and AI Agent-generated work must follow these principles:

```text
1. Develop locally first.
2. Do not develop directly on staging.
3. Do not use production data locally unless anonymized and approved.
4. Do not use production secrets locally.
5. Run lint/test before Pull Request.
6. Use Docker Compose for local dependencies.
7. Keep service database ownership separated.
8. Use .env.example as reference, never commit .env.
```

---

## 3. Recommended Machine Specification

Minimum recommended local machine:

```text
CPU    : 4 cores
RAM    : 16 GB
Disk   : 50 GB free
OS     : Windows 10/11, macOS, or Linux
```

Recommended for smoother development:

```text
CPU    : 8 cores
RAM    : 32 GB
Disk   : 100 GB free SSD
```

Why 32 GB is recommended:

```text
- Docker containers
- Multiple Go services
- Next.js dev server
- Flutter tooling/emulator
- IDE
- AI coding tools
```

---

## 4. Required Tools

## 4.1 Core Tools

Install these tools first:

```text
Git
Docker Desktop / Docker Engine
Docker Compose
Go
Node.js
pnpm or npm
Flutter SDK
Make
```

Recommended versions:

```text
Go        : 1.23+ or current stable
Node.js   : 20 LTS+
pnpm      : 9+
Flutter   : 3.x stable
Docker    : latest stable
PostgreSQL client tools: optional but recommended
```

---

## 4.2 Backend Tools

Required for backend development:

```text
Go
goose
sqlc
golangci-lint
protoc
protoc-gen-go
protoc-gen-go-grpc
grpcurl optional
```

### Install Go tools

Example:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

Install `golangci-lint` using the official installation method for your OS.

Verify:

```bash
go version
goose -version
sqlc version
golangci-lint --version
protoc --version
protoc-gen-go --version
grpcurl --version
```

---

## 4.3 Frontend Web Tools

Required for Next.js web admin:

```text
Node.js 20 LTS+
pnpm recommended
```

Install pnpm:

```bash
npm install -g pnpm
```

Verify:

```bash
node -v
pnpm -v
```

---

## 4.4 Flutter Mobile Tools

Required for mobile app:

```text
Flutter SDK 3.x
Dart SDK included with Flutter
Android Studio or VS Code Flutter extension
Android SDK
Android emulator or physical device
Xcode for iOS development on macOS
```

Verify:

```bash
flutter --version
flutter doctor
```

For MVP, Android testing is enough if iOS setup is not available.

---

## 4.5 Optional but Recommended Tools

Recommended tools:

```text
VS Code / VSCodium
Go extension
Flutter extension
Docker extension
REST Client / Thunder Client / Postman / Insomnia
TablePlus / DBeaver / pgAdmin
RabbitMQ Management UI
MinIO Console
```

Useful CLI tools:

```text
jq
curl
httpie
psql
tree
```

---

## 5. Windows-Specific Notes

For Windows developers, recommended setup:

```text
Windows 10/11
WSL2 enabled
Ubuntu on WSL2
Docker Desktop with WSL2 backend
VS Code Remote WSL extension
```

Recommended workflow:

```text
Clone repository inside WSL filesystem, not C:\ drive.
```

Example:

```bash
cd ~/source
git clone git@github.com:<org>/school-platform.git
cd school-platform
```

Avoid:

```text
/mnt/c/... for heavy Node/Go/Docker development
```

Reason:

```text
WSL filesystem is faster and more stable for development.
```

---

## 6. Repository Setup

Clone repository:

```bash
git clone git@github.com:<org>/school-platform.git
cd school-platform
```

Checkout development branch:

```bash
git checkout develop
git pull origin develop
```

Create feature branch:

```bash
git checkout -b feature/sprint-0-local-setup
```

Repository structure:

```text
school-platform/
├── apps/
│   ├── web-admin/
│   └── mobile-app/
├── services/
│   ├── api-gateway/
│   ├── identity-service/
│   ├── school-core-service/
│   ├── admission-service/
│   ├── academic-service/
│   ├── finance-service/
│   ├── communication-service/
│   └── reporting-service/
├── packages/
│   ├── proto/
│   ├── openapi/
│   ├── events/
│   └── shared-go/
├── infra/
├── deploy/
├── docs/
├── scripts/
├── docker-compose.yml
├── Makefile
└── README.md
```

---

## 7. Environment File Setup

Copy example environment files:

```bash
cp .env.example .env
cp services/api-gateway/.env.example services/api-gateway/.env
```

For each service that exists:

```bash
cp services/identity-service/.env.example services/identity-service/.env
cp services/school-core-service/.env.example services/school-core-service/.env
cp services/admission-service/.env.example services/admission-service/.env
cp services/academic-service/.env.example services/academic-service/.env
cp services/finance-service/.env.example services/finance-service/.env
cp services/communication-service/.env.example services/communication-service/.env
cp services/reporting-service/.env.example services/reporting-service/.env
```

Rules:

```text
- Do not commit .env files.
- Commit only .env.example.
- Do not use production secrets locally.
- Keep local credentials simple and safe.
```

---

## 8. Local Infrastructure with Docker Compose

Start local dependencies:

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
```

or:

```bash
docker compose logs -f
```

Stop containers:

```bash
make down
```

Restart:

```bash
make restart
```

Optional local observability profile:

```bash
docker compose --profile observability up -d
docker compose --profile observability ps
```

This profile is intentionally optional in Sprint 0 so normal local setup stays lightweight.

---

## 9. Local Services and Ports

Recommended local ports:

```text
API Gateway         : http://localhost:8080
Web Admin           : http://localhost:3000
PostgreSQL          : localhost:5432
Redis               : localhost:6379
RabbitMQ AMQP       : localhost:5672
RabbitMQ Management : http://localhost:15672
MinIO API           : http://localhost:9000
MinIO Console       : http://localhost:9001
Mailpit SMTP        : localhost:1025
Mailpit UI          : http://localhost:8025
Prometheus optional : http://localhost:9090
Grafana optional    : http://localhost:3001
Loki optional       : http://localhost:3100
```

RabbitMQ local login example:

```text
username: guest
password: guest
```

MinIO local login example:

```text
username: minioadmin
password: minioadmin
```

These are local-only credentials and must not be used in production.

---

## 10. Local PostgreSQL Databases

MVP uses one PostgreSQL instance with separate databases per service:

```text
identity_db
school_core_db
admission_db
academic_db
finance_db
communication_db
reporting_db
```

Each service must only connect to its own database.

Example:

```text
identity-service      → identity_db
school-core-service   → school_core_db
admission-service     → admission_db
academic-service      → academic_db
finance-service       → finance_db
communication-service → communication_db
reporting-service     → reporting_db
```

Rule:

```text
No service may query another service database.
```

---

## 11. Run Database Migrations

Each service owns its migration.

Example for Identity Service:

```bash
cd services/identity-service
goose -dir internal/db/migrations postgres "$DATABASE_URL" up
```

Recommended Makefile pattern:

```bash
make migrate-up service=identity-service
make migrate-down service=identity-service
make migrate-status service=identity-service
```

Rules:

```text
- Migration tool: goose
- Do not mix goose with golang-migrate
- Migration belongs to the service owner
- Do not run finance migration against school_core_db
```

---

## 12. Generate SQLC Code

Each service owns its sqlc config.

Example:

```bash
cd services/finance-service
sqlc generate
```

Recommended Makefile pattern:

```bash
make sqlc service=finance-service
```

Rules:

```text
- All main database queries must go through sqlc.
- Do not create unsafe raw SQL string concatenation.
- Query must include foundation_id and school_id filters when relevant.
```

---

## 13. Generate Protobuf Code

Proto files are stored in:

```text
packages/proto/
```

Recommended command:

```bash
make proto
```

Expected generated code location:

```text
packages/proto/gen/go/
```

Rules:

```text
- Do not manually edit generated files.
- Do not reuse deleted protobuf field numbers.
- Use reserved fields for removed proto fields.
```

---

## 14. Run Backend Services

Run API Gateway:

```bash
cd services/api-gateway
go run ./cmd/server
```

Run Identity Service:

```bash
cd services/identity-service
go run ./cmd/server
```

Run a service worker if available:

```bash
cd services/finance-service
go run ./cmd/worker
```

Recommended Makefile pattern:

```bash
make run service=api-gateway
make run service=identity-service
make worker service=finance-service
```

Health checks:

```bash
curl http://localhost:8080/healthz
curl http://localhost:8080/readyz
```

API ping:

```bash
curl http://localhost:8080/api/v1/ping
```

---

## 15. Run Web Admin

Go to web app folder:

```bash
cd apps/web-admin
pnpm install
pnpm dev
```

Default URL:

```text
http://localhost:3000
```

Recommended commands:

```bash
pnpm lint
pnpm typecheck
pnpm build
```

Rules:

```text
- Use API Gateway as backend API.
- Do not call microservice directly from frontend.
- Use React Query for server state.
- Use Zustand for global UI/context state.
- Use Zod for form validation.
```

---

## 16. Run Mobile App

Go to mobile app folder:

```bash
cd apps/mobile-app
flutter pub get
flutter run
```

Check Flutter environment:

```bash
flutter doctor
```

Recommended commands:

```bash
flutter analyze
flutter test
```

Rules:

```text
- Mobile calls API Gateway only.
- Store tokens using Flutter Secure Storage.
- Offline write is not part of MVP.
- Do not store sensitive data locally unless encrypted and explicitly required.
```

---

## 17. Common Makefile Commands

Recommended commands:

```bash
make help
make setup
make up
make down
make restart
make ps
make logs
make test
make lint
make fmt
make build
make proto
make openapi-check
make event-schema-check
```

Service-specific examples:

```bash
make test service=identity-service
make lint service=finance-service
make build service=api-gateway
make migrate-up service=school-core-service
make sqlc service=academic-service
```

---

## 18. Daily Development Workflow

Recommended daily workflow:

```bash
git checkout develop
git pull origin develop
git checkout -b feature/<area>-<short-description>

make up

# work on one small task
make fmt
make lint
make test

git status
git add .
git commit -m "feat(area): clear message"
git push origin feature/<area>-<short-description>
```

Then create Pull Request:

```text
feature/* → develop
```

Before PR:

```text
- Tests pass
- Lint passes
- No secrets committed
- Docs updated if contract changed
- Migration checked if added
- OpenAPI/proto/event updated if changed
```

---

## 19. Working with AI Agent Locally

When using AI Agent, always provide:

```text
- target sprint
- target task
- target service/app
- relevant docs
- scope
- out of scope
- required tests
- expected files
```

Use task prompt docs:

```text
13-sprint-0-task-prompts.md
14-sprint-1-task-prompts.md
...
23-sprint-10-task-prompts.md
```

AI Agent must not:

```text
- work on too many modules in one task
- change architecture without instruction
- query another service database
- put business logic in API Gateway
- skip tests
- log sensitive data
```

Good prompt style:

```text
Implement Task 5.5 — Bill Generation with Snapshots from docs/18-sprint-5-task-prompts.md.
Only work inside finance-service and required shared packages.
Do not implement payment gateway.
Add migration, sqlc query, usecase, tests, and event publishing.
```

Bad prompt style:

```text
Build the whole finance system.
```

---

## 20. Local Testing Checklist

Before pushing a branch, run:

```bash
make fmt
make lint
make test
```

If backend service changed:

```bash
make test service=<service-name>
make lint service=<service-name>
```

If migration changed:

```bash
make migrate-up service=<service-name>
make migrate-down service=<service-name>
make migrate-up service=<service-name>
```

If sqlc query changed:

```bash
make sqlc service=<service-name>
make test service=<service-name>
```

If proto changed:

```bash
make proto
make test
```

If OpenAPI changed:

```bash
make openapi-check
```

If event schema changed:

```bash
make event-schema-check
```

If web changed:

```bash
cd apps/web-admin
pnpm lint
pnpm typecheck
pnpm build
```

If mobile changed:

```bash
cd apps/mobile-app
flutter analyze
flutter test
```

---

## 21. Local Seed Data

Local seed data should include:

```text
1 foundation
TK, SD, SMP, SMA schools
MVP roles and permissions
dummy users for each role
active academic year
active semester
grade levels
classes
students
guardians
teachers
PPDB sample
fee type and SPP sample
sample bills
sample attendance
```

Rules:

```text
- Use synthetic/dummy data only.
- Do not use real student/parent data.
- Do not use production dump.
- Seed data must be safe to commit if it contains no secrets and no real PII.
```

Recommended seed accounts:

```text
admin.yayasan@example.test
kepsek.sd@example.test
tu.sd@example.test
bendahara.sd@example.test
guru.sd@example.test
parent.sd@example.test
siswa.sd@example.test
```

Use obvious local-only passwords in `.env.example` or seed docs, not production.

---

## 22. Local Object Storage

Use MinIO for local S3-compatible storage.

Recommended local buckets:

```text
school-platform-local
school-platform-private
school-platform-public
```

File rules:

```text
- Private files by default.
- Restricted/Confidential files must not be public.
- Signed URL only after authorization.
- Do not log raw file content.
```

---

## 23. Local RabbitMQ

RabbitMQ is used for domain events.

Exchange:

```text
domain.events
```

Exchange type:

```text
topic
```

Local queues may include:

```text
reporting-service.events
communication-service.events
reporting-service.events.retry
reporting-service.events.dlq
communication-service.events.retry
communication-service.events.dlq
```

Rules:

```text
- Consumers must be idempotent.
- Use processed_events table where needed.
- Retry and DLQ must be prepared for important consumers.
```

---

## 24. Local Redis

Redis may be used for:

```text
- rate limiting
- cache
- temporary session-related storage if needed
- distributed lock if needed
```

Rules:

```text
- Do not store sensitive token raw unless explicitly designed and protected.
- Prefer hashed refresh token in database.
- Keep caching minimal in MVP.
```

---

## 25. Environment Variables Reference

Common service variables:

```text
SERVICE_NAME=
APP_ENV=local
HTTP_PORT=
GRPC_PORT=
DATABASE_URL=
REDIS_URL=
RABBITMQ_URL=
LOG_LEVEL=debug
JWT_PUBLIC_KEY_PATH=
JWT_PRIVATE_KEY_PATH=
MINIO_ENDPOINT=
MINIO_ACCESS_KEY=
MINIO_SECRET_KEY=
MINIO_BUCKET=
```

API Gateway variables:

```text
SERVICE_NAME=api-gateway
APP_ENV=local
HTTP_PORT=8080
LOG_LEVEL=debug
CORS_ALLOWED_ORIGINS=http://localhost:3000
IDENTITY_GRPC_ADDR=localhost:9101
SCHOOL_CORE_GRPC_ADDR=localhost:9102
ADMISSION_GRPC_ADDR=localhost:9103
ACADEMIC_GRPC_ADDR=localhost:9104
FINANCE_GRPC_ADDR=localhost:9105
COMMUNICATION_GRPC_ADDR=localhost:9106
REPORTING_GRPC_ADDR=localhost:9107
```

Rules:

```text
- Environment variable names should be consistent.
- Production secrets must be managed outside Git.
- .env files are ignored by Git.
```

---

## 26. Troubleshooting

## 26.1 Docker Port Already Used

Check ports:

```bash
docker compose ps
lsof -i :5432
lsof -i :8080
```

Fix:

```text
- Stop conflicting service
- Change local port mapping
- Restart Docker
```

---

## 26.2 PostgreSQL Connection Failed

Check:

```bash
docker compose ps postgres
docker compose logs postgres
```

Verify DATABASE_URL:

```text
postgres://user:password@localhost:5432/identity_db?sslmode=disable
```

Common causes:

```text
- container not running
- wrong port
- wrong database name
- migration not run
```

---

## 26.3 RabbitMQ Not Available

Check:

```bash
docker compose ps rabbitmq
docker compose logs rabbitmq
```

Open management UI:

```text
http://localhost:15672
```

Common local credentials:

```text
guest / guest
```

---

## 26.4 MinIO Upload Failed

Check:

```bash
docker compose ps minio
docker compose logs minio
```

Open console:

```text
http://localhost:9001
```

Check:

```text
- bucket exists
- access key/secret correct
- endpoint correct
```

---

## 26.5 Go Tool Not Found

Check:

```bash
which goose
which sqlc
which golangci-lint
echo $PATH
```

Make sure Go bin is in PATH:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Add to shell config:

```bash
~/.bashrc
~/.zshrc
```

---

## 26.6 Node Dependencies Failed

Try:

```bash
cd apps/web-admin
rm -rf node_modules
pnpm install
```

If pnpm is missing:

```bash
npm install -g pnpm
```

---

## 26.7 Flutter Doctor Issues

Run:

```bash
flutter doctor
```

Common fixes:

```text
- Install Android Studio
- Install Android SDK
- Accept Android licenses
- Configure emulator/device
```

Accept Android licenses:

```bash
flutter doctor --android-licenses
```

---

## 27. Local Security Rules

Do not store in repository:

```text
.env
private keys
production credentials
database dumps
backup files
real student data
real parent data
payment proof files
tokens
```

Do not log:

```text
password
access token
refresh token
raw payment proof
raw document content
NIK full
BK/UKS detail
payroll data
Confidential data
```

Local development must use:

```text
dummy data
synthetic data
local-only credentials
```

---

## 28. Recommended IDE Setup

VS Code / VSCodium extensions:

```text
Go
Docker
ESLint
Prettier
Tailwind CSS IntelliSense
Flutter
Dart
YAML
REST Client
GitLens optional
```

Recommended settings:

```text
format on save
go test on save optional
eslint validate TypeScript
editorconfig if added
```

---

## 29. Pre-PR Checklist

Before opening Pull Request:

```text
- [ ] Branch is based on latest develop
- [ ] Task scope is respected
- [ ] No unrelated files changed
- [ ] make fmt passed
- [ ] make lint passed
- [ ] make test passed
- [ ] Migration up/down tested if migration exists
- [ ] sqlc generated if queries changed
- [ ] proto generated if proto changed
- [ ] OpenAPI updated if API changed
- [ ] Event schema updated if event changed
- [ ] No secrets committed
- [ ] No sensitive data logged
- [ ] README/docs updated if setup changed
```

---

## 30. Local Development Readiness Checklist

A developer is ready to start implementation when:

```text
- [ ] Git is installed
- [ ] Docker is installed and running
- [ ] Go is installed
- [ ] Node.js is installed
- [ ] pnpm is installed
- [ ] Flutter is installed if working on mobile
- [ ] goose is installed
- [ ] sqlc is installed
- [ ] golangci-lint is installed
- [ ] protoc and Go plugins are installed
- [ ] Repository is cloned
- [ ] .env files are created from .env.example
- [ ] docker compose up works
- [ ] PostgreSQL is reachable
- [ ] Redis is reachable
- [ ] RabbitMQ Management is reachable
- [ ] MinIO Console is reachable
- [ ] API Gateway /healthz works
- [ ] make test works
- [ ] make lint works
```

---

## 31. Final Summary

Local development uses:

```text
Docker Compose for dependencies
Go services locally or in containers
Next.js dev server for web
Flutter tooling for mobile
PostgreSQL database per service
RabbitMQ for events
MinIO for file storage
Redis for cache/rate limit
Makefile for common commands
```

Development flow:

```text
develop locally
→ test locally
→ push feature branch
→ Pull Request to develop
→ CI
→ review
→ merge
→ staging later
→ production only from main with approval
```
