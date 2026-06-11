# school-platform

`school-platform` is an internal multi-unit school foundation management system for TK, SD, SMP, and SMA operations.

The MVP provides the technical and product foundation for identity, school core data, PPDB, finance/SPP manual payment, academic basics, communication, reporting, file handling, audit, security, observability, and local development workflows.

Internal code, folder names, API paths, database names, events, and permissions use English. User-facing UI labels use Bahasa Indonesia.

## MVP Purpose

The MVP is designed to help a school foundation and its school units run core operational flows:

- manage foundation and school data
- manage students, guardians, teachers, classes, and assignments
- run PPDB from applicant intake to student conversion
- generate SPP bills and process manual payments
- record attendance, grades, and basic report cards
- publish announcements and notifications
- provide foundation, school, teacher, and parent/student dashboards

Large modules such as payroll, full HR, asset inventory, library, BK/UKS detail, full LMS, payment gateway, WhatsApp, offline write mobile, global search, and Kubernetes are outside the MVP.

## Tech Stack Summary

Backend:

- Go
- Chi for the API Gateway HTTP router
- gRPC and protobuf for internal service communication
- PostgreSQL with database per service
- pgx, sqlc, and goose
- Redis
- RabbitMQ
- MinIO or S3-compatible object storage
- slog for structured logging

Frontend:

- Next.js 14 with TypeScript for `apps/web-admin`
- Tailwind CSS and shadcn/ui
- React Query, Zustand, React Hook Form, and Zod

Mobile:

- Flutter 3 and Dart for `apps/mobile-app`
- Riverpod
- Dio and Retrofit
- Flutter Secure Storage

Infrastructure and workflow:

- Docker Compose for local dependencies and early staging
- GitHub Actions for CI
- Makefile for common developer commands
- Manual approval for production deployment

## Monorepo Structure

```text
apps/
  web-admin/              Next.js web admin app
  mobile-app/             Flutter mobile app
services/
  api-gateway/            External REST API gateway
  identity-service/       Identity, credentials, sessions, roles, permissions
  school-core-service/    Foundation, schools, students, guardians, teachers, classes
  admission-service/      PPDB and applicant workflow
  academic-service/       Curriculum, schedules, attendance, grades, report cards
  finance-service/        Fees, bills, payments, receipts, finance approvals
  communication-service/  Announcements and notifications
  reporting-service/      Dashboard projections and read models
packages/
  proto/                  Internal gRPC protobuf contracts
  openapi/                External REST API contracts
  events/                 Domain event schemas
  shared-go/              Shared Go helpers and conventions
infra/                    Docker, Nginx, PostgreSQL, Redis, RabbitMQ, MinIO, observability
deploy/                   Staging and production deployment assets
docs/                     Architecture, workflow, sprint, testing, and setup docs
scripts/                  Project automation scripts
```

## Local Setup Quick Start

Local development is done before staging. Do not develop directly on staging, and do not use production data or production secrets locally.

```bash
git clone git@github.com:<org>/school-platform.git
cd school-platform
git checkout develop
```

Create local environment files from examples when they are available:

```bash
cp .env.example .env
```

Start local dependencies when `docker-compose.yml` is available:

```bash
make up
# or
docker compose up -d
```

Run checks before opening a Pull Request:

```bash
make fmt
make lint
make test
```

See [docs/LOCAL_DEVELOPMENT.md](docs/LOCAL_DEVELOPMENT.md) for the practical local setup guide.

## Important Docs

- [AGENTS.md](AGENTS.md) - mandatory AI Agent rules
- [SKILLS.md](SKILLS.md) - AI Agent workflow skills
- [docs/README.md](docs/README.md) - documentation index
- [docs/01-technical-architecture.md](docs/01-technical-architecture.md) - architecture baseline
- [docs/02-service-boundary.md](docs/02-service-boundary.md) - service ownership and boundaries
- [docs/04-api-contract.md](docs/04-api-contract.md) - external REST and internal gRPC rules
- [docs/05-event-contract.md](docs/05-event-contract.md) - RabbitMQ event contract
- [docs/08-coding-standard.md](docs/08-coding-standard.md) - coding standard and project convention
- [docs/09-ai-agent-rules.md](docs/09-ai-agent-rules.md) - implementation guardrails
- [docs/10-sprint-backlog-mvp.md](docs/10-sprint-backlog-mvp.md) - MVP sprint roadmap
- [docs/11-github-repository-rules.md](docs/11-github-repository-rules.md) - branch, PR, CI/CD, and release rules
- [docs/LOCAL_DEVELOPMENT.md](docs/LOCAL_DEVELOPMENT.md) - basic local setup guide

## Branch Workflow

Protected branch flow:

```text
feature/* -> develop -> staging -> main/production
```

Rules:

- All changes go through Pull Request.
- `develop` is daily integration.
- `staging` is for QA, UAT, and release candidate validation.
- `main` represents production-ready code.
- Production deployment requires manual approval.
- Commit and PR titles should follow `type(scope): short description`.

Example:

```text
chore(repository): add local development docs
```
