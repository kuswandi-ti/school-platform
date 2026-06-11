# 13 — Sprint 0 Task Prompts

Project: `school-platform`  
Sprint: Sprint 0 — Project Foundation  
Status: AI Agent task prompt pack  
Scope: Monorepo foundation, local development environment, service template, API Gateway skeleton, Makefile, CI, and documentation placeholders.

---

## 1. Sprint 0 Objective

Sprint 0 prepares the technical foundation before domain feature development starts.

The goal is to make the repository ready for a team and AI Agent to work consistently.

Sprint 0 must produce:

```text
- Monorepo structure
- Local Docker Compose environment
- Go service template
- API Gateway skeleton
- Shared packages folders
- Basic health/readiness endpoints
- Logging and correlation ID baseline
- Makefile commands
- GitHub Actions basic CI
- Documentation placeholders
```

---

## 2. Global Rules for Sprint 0

AI Agent must follow these rules for every Sprint 0 task:

```text
- Do not implement business domain features yet.
- Do not implement full authentication yet.
- Do not implement database schema for business modules yet.
- Do not introduce Kubernetes.
- Do not introduce payment gateway, WhatsApp, LMS, HR, payroll, or other out-of-MVP modules.
- Keep structure simple, consistent, and ready for Sprint 1.
- Use Go, Chi, slog, Makefile, Docker Compose.
- Prepare folders for proto, OpenAPI, and event schemas.
- Do not commit secrets or real credentials.
- Use .env.example only.
```

Architecture guardrails:

```text
- API Gateway is not a business logic layer.
- Service template must support request_id and correlation_id.
- Services must be database-owner isolated later.
- Future services must fit cmd/internal/domain/usecase/repository/transport/event/authz/audit/db structure.
```

---

## 3. Sprint 0 Definition of Done

Sprint 0 is done when:

```text
- Repository structure exists.
- Docker Compose starts PostgreSQL, Redis, RabbitMQ, MinIO, and optional Mailpit.
- API Gateway starts locally.
- Sample service starts locally.
- API Gateway and sample service expose /healthz and /readyz.
- Basic slog structured logging works.
- request_id and correlation_id middleware exists.
- Makefile has setup/up/down/logs/test/lint/build commands.
- GitHub Actions CI exists.
- .env.example files exist.
- README has local setup instructions.
- No production secrets are committed.
```

---

# Task 0.1 — Create Monorepo Structure

## Prompt

```text
You are working on `school-platform`, a Go microservice monorepo for an internal school foundation management system.

Task:
Create the initial monorepo folder structure for Sprint 0.

Relevant docs:
- docs/01-technical-architecture.md
- docs/08-coding-standard.md
- docs/09-ai-agent-rules.md
- docs/10-sprint-backlog-mvp.md
- docs/11-github-repository-rules.md

Goal:
Prepare the repository structure so all future services, apps, shared packages, infrastructure, deployment scripts, and docs have a consistent location.

Scope:
Create these top-level folders:
- apps/
- services/
- packages/
- infra/
- deploy/
- docs/
- scripts/

Create these app placeholders:
- apps/web-admin/
- apps/mobile-app/

Create these service placeholders:
- services/api-gateway/
- services/identity-service/
- services/school-core-service/
- services/admission-service/
- services/academic-service/
- services/finance-service/
- services/communication-service/
- services/reporting-service/

Create these packages placeholders:
- packages/proto/
- packages/openapi/
- packages/events/
- packages/shared-go/

Create these infra placeholders:
- infra/docker/
- infra/nginx/
- infra/postgres/
- infra/redis/
- infra/rabbitmq/
- infra/minio/
- infra/observability/

Create these deploy placeholders:
- deploy/staging/
- deploy/production/

Create placeholder README.md files where useful so empty directories are tracked.

Out of Scope:
- Do not implement business code.
- Do not implement actual services yet.
- Do not implement Docker Compose yet.
- Do not add production credentials.

Rules:
- Use English for internal naming.
- Use kebab-case for service folders.
- Keep names aligned with service boundary decisions.
- Do not add unrelated modules.

Acceptance Criteria:
- All required folders exist.
- Placeholder README.md files explain the purpose of each main folder.
- No secrets or environment-specific credentials are committed.
- Folder names match documented conventions.

Tests Required:
- No automated test required for this task.
- Provide a tree output or folder list in the final response.

Expected Output:
- Summary of created folders.
- List of created placeholder files.
```

---

# Task 0.2 — Add Base README and Local Setup Guide

## Prompt

```text
You are working on Sprint 0 for `school-platform`.

Task:
Create the root README.md and a basic local setup guide.

Goal:
Provide developers and AI Agent with a clear entry point for setting up and understanding the repository.

Scope:
Create/update:
- README.md
- docs/LOCAL_DEVELOPMENT.md if not already present

README.md should include:
- project name: school-platform
- short project description
- MVP purpose
- tech stack summary
- monorepo structure overview
- local setup quick start
- important docs list
- branch workflow summary

docs/LOCAL_DEVELOPMENT.md should include:
- prerequisites
- Docker Compose startup
- environment file setup
- common Makefile commands
- local service ports
- how to run tests
- notes about not using production data/secrets

Out of Scope:
- Do not create domain-specific setup details.
- Do not document production deployment deeply.
- Do not add secrets.

Rules:
- Keep the language clear and practical.
- Use English for technical terms.
- Mention that UI labels use Bahasa Indonesia.
- Mention local development is done before staging.

Acceptance Criteria:
- README.md exists and is useful for first-time developer.
- docs/LOCAL_DEVELOPMENT.md exists.
- Local setup steps are clear.
- No production secrets are included.

Tests Required:
- No automated test required.
- Review Markdown formatting.
```

---

# Task 0.3 — Create Docker Compose for Local Dependencies

## Prompt

```text
You are working on Sprint 0 for `school-platform`.

Task:
Create Docker Compose configuration for local development dependencies.

Goal:
Allow developers to start core local dependencies consistently.

Scope:
Create/update:
- docker-compose.yml
- .env.example
- infra/postgres/README.md
- infra/rabbitmq/README.md
- infra/minio/README.md
- infra/redis/README.md

Docker Compose services:
- postgres
- redis
- rabbitmq with management UI
- minio
- optional mailpit

PostgreSQL should prepare the following databases:
- identity_db
- school_core_db
- admission_db
- academic_db
- finance_db
- communication_db
- reporting_db

Recommended local ports:
- PostgreSQL: 5432
- Redis: 6379
- RabbitMQ: 5672
- RabbitMQ Management: 15672
- MinIO API: 9000
- MinIO Console: 9001
- Mailpit UI: 8025
- Mailpit SMTP: 1025

Out of Scope:
- Do not create business database migrations yet.
- Do not add production credentials.
- Do not add Kubernetes.
- Do not add cloud provider configs.

Rules:
- Use .env.example for configurable values.
- Use safe local-only credentials.
- Make clear these credentials are not for production.
- Use named Docker volumes.
- Use healthchecks where practical.

Acceptance Criteria:
- `docker compose up -d` starts all dependency containers.
- PostgreSQL container is available.
- RabbitMQ management UI is available.
- MinIO console is available.
- Redis is available.
- Mailpit is available if included.
- Databases for all MVP services are created or documented.
- No production secrets are included.

Tests Required:
- Provide commands to verify services:
  - docker compose ps
  - docker compose logs
  - psql connection example if available
```

---

# Task 0.4 — Create Go Service Template

## Prompt

```text
You are working on Sprint 0 for `school-platform`.

Task:
Create a reusable Go service template that future microservices can follow.

Goal:
Establish consistent service structure, health endpoints, config loading, logging, and graceful shutdown.

Target:
Create a sample service or template under:
- services/_template-service/
or use:
- services/identity-service/ as the first skeleton if preferred

Scope:
Create structure:
- cmd/server/main.go
- internal/app/app.go
- internal/config/config.go
- internal/transport/http/health.go
- internal/logger/logger.go or shared logger usage
- Dockerfile
- go.mod
- README.md
- .env.example

Features:
- Load config from environment variables.
- Start HTTP server.
- Expose:
  - GET /healthz
  - GET /readyz
- Use slog structured logging.
- Include request_id/correlation_id support if HTTP middleware is included.
- Graceful shutdown on SIGINT/SIGTERM.

Out of Scope:
- No business logic.
- No database repository.
- No gRPC yet unless needed for template.
- No authentication.
- No domain modules.

Rules:
- Use Go standard library where possible.
- Use Chi only if routing is needed.
- Use slog.
- Keep template minimal.
- Do not add unnecessary framework.

Acceptance Criteria:
- Service can run locally.
- /healthz returns 200.
- /readyz returns 200.
- Startup logs include service name and environment.
- Service shuts down gracefully.
- Dockerfile builds service image.

Tests Required:
- Unit test for health handler.
- `go test ./...` passes.
```

---

# Task 0.5 — Create API Gateway Skeleton

## Prompt

```text
You are working on Sprint 0 for `school-platform`.

Task:
Create the initial API Gateway skeleton using Go and Chi.

Goal:
Prepare the API Gateway as the single external REST/JSON entrypoint for Next.js and Flutter.

Target:
- services/api-gateway/

Scope:
Create:
- cmd/server/main.go
- internal/app/
- internal/config/
- internal/transport/http/
- internal/middleware/
- internal/response/
- Dockerfile
- go.mod
- .env.example
- README.md

Features:
- Start HTTP server.
- Expose:
  - GET /healthz
  - GET /readyz
  - GET /api/v1/ping
- Use Chi router.
- Add middleware:
  - recover
  - request_id
  - correlation_id
  - structured logging
  - basic CORS placeholder
- Add standard response helper:
  - data
  - meta
  - error
- Add standard error response helper.
- Add placeholder route group /api/v1.
- Add placeholder for future REST-to-gRPC clients.

Out of Scope:
- Do not implement login yet.
- Do not implement JWT validation yet.
- Do not implement service-to-service gRPC calls yet.
- Do not implement domain endpoints yet.
- Do not add business logic.

Rules:
- API Gateway must not contain business logic.
- API Gateway standardizes external API response format.
- API Gateway must propagate request_id/correlation_id in future.
- Use slog structured logging.
- Use environment variable config.

Acceptance Criteria:
- API Gateway runs locally.
- /healthz returns success.
- /readyz returns success.
- /api/v1/ping returns standard response format.
- Logs include request_id and correlation_id.
- Error response follows standard format.
- go test ./... passes.

Tests Required:
- Health handler test.
- Ping handler test.
- Response format test.
- Middleware request_id/correlation_id test if practical.
```

---

# Task 0.6 — Add Shared Go Packages Skeleton

## Prompt

```text
You are working on Sprint 0 for `school-platform`.

Task:
Create skeleton shared Go packages for common cross-service utilities.

Goal:
Prepare shared utilities without over-implementing business logic.

Target:
- packages/shared-go/

Scope:
Create package folders:
- packages/shared-go/logger
- packages/shared-go/config
- packages/shared-go/response
- packages/shared-go/errors
- packages/shared-go/context
- packages/shared-go/audit
- packages/shared-go/events
- packages/shared-go/messaging
- packages/shared-go/numbering
- packages/shared-go/files

Each package should include:
- README.md describing intended responsibility
- minimal Go files only if useful
- no domain-specific logic

Suggested minimal implementations:
- response: standard response/error structs
- errors: standard error code constants
- context: ActorContext struct placeholder
- logger: slog setup helper
- events: event envelope struct placeholder
- messaging: RabbitMQ abstraction placeholder

Out of Scope:
- Do not implement full audit system.
- Do not implement full event publisher yet.
- Do not implement numbering logic yet.
- Do not implement file storage integration yet.
- Do not implement domain-specific code.

Rules:
- Shared packages must not depend on specific service database.
- Shared packages must remain generic.
- Avoid circular dependencies.
- Keep minimal and extend later when needed.

Acceptance Criteria:
- Shared package folders exist.
- README explains each package.
- Minimal common structs/constants exist where useful.
- No business logic included.
- go test ./... passes if Go module is initialized.

Tests Required:
- Basic compile test if Go files are added.
```

---

# Task 0.7 — Create Proto, OpenAPI, and Events Placeholder Structure

## Prompt

```text
You are working on Sprint 0 for `school-platform`.

Task:
Create placeholder structure for protobuf, OpenAPI, and event schemas.

Goal:
Prepare contract-first folders for future API, gRPC, and event development.

Target:
- packages/proto/
- packages/openapi/
- packages/events/

Scope:
Create proto structure:
- packages/proto/identity/v1/
- packages/proto/schoolcore/v1/
- packages/proto/admission/v1/
- packages/proto/academic/v1/
- packages/proto/finance/v1/
- packages/proto/communication/v1/
- packages/proto/reporting/v1/
- packages/proto/common/v1/

Create OpenAPI structure:
- packages/openapi/api-gateway.v1.yaml
- packages/openapi/README.md

Create event schema structure:
- packages/events/envelope.schema.json
- packages/events/identity/
- packages/events/school/
- packages/events/admission/
- packages/events/academic/
- packages/events/finance/
- packages/events/communication/
- packages/events/reporting/
- packages/events/README.md

Add README files explaining:
- proto naming convention
- OpenAPI update rules
- event schema update rules

Out of Scope:
- Do not fully define all service proto methods yet.
- Do not fully define all OpenAPI endpoints yet.
- Do not fully define all event payloads yet.

Rules:
- Use package naming:
  - schoolplatform.identity.v1
  - schoolplatform.finance.v1
- Preserve future contract compatibility.
- Do not invent domain APIs beyond placeholder examples.

Acceptance Criteria:
- Contract folders exist.
- README explains update rules.
- envelope.schema.json has basic event envelope schema.
- api-gateway.v1.yaml has minimal valid OpenAPI skeleton.
- Placeholder proto files compile only if build tooling is added; otherwise keep README placeholders.

Tests Required:
- If OpenAPI lint exists, it passes.
- If proto-check exists, it passes.
```

---

# Task 0.8 — Add Makefile Commands

## Prompt

```text
You are working on Sprint 0 for `school-platform`.

Task:
Create the root Makefile for local development and CI-friendly commands.

Goal:
Standardize common development commands for developers and AI Agent.

Scope:
Create/update:
- Makefile

Required commands:
- make setup
- make up
- make down
- make restart
- make logs
- make ps
- make test
- make lint
- make fmt
- make build
- make clean
- make proto
- make openapi-check
- make event-schema-check

Optional service-specific pattern:
- make test service=finance-service
- make lint service=identity-service
- make build service=api-gateway
- make logs service=api-gateway

Out of Scope:
- Do not create production deploy command yet.
- Do not add destructive commands without warning.
- Do not hardcode secrets.

Rules:
- Commands should be safe for local development.
- Destructive commands must be clearly named and documented.
- Use Docker Compose commands where relevant.
- Commands must work on typical Unix shell environment.

Acceptance Criteria:
- make up starts local stack.
- make down stops local stack.
- make logs shows logs.
- make test runs Go tests where modules exist.
- make lint has placeholder or actual lint command.
- README references Makefile commands.

Tests Required:
- Run make help or make command list if implemented.
- Provide command examples in final response.
```

---

# Task 0.9 — Add GitHub Actions Basic CI

## Prompt

```text
You are working on Sprint 0 for `school-platform`.

Task:
Add GitHub Actions basic CI workflow.

Goal:
Ensure every Pull Request has automated checks before merge.

Target:
- .github/workflows/ci.yml

Scope:
CI should support:
- checkout repository
- setup Go
- run gofmt check
- run go vet
- run go test
- optional golangci-lint
- setup Node if web-admin exists
- run frontend lint/typecheck/build if package exists
- setup Flutter placeholder if mobile app is initialized later
- run proto/openapi/event schema placeholder checks if scripts exist

Out of Scope:
- No deployment workflow yet.
- No production secrets.
- No staging deploy yet.
- No Docker image push yet unless already ready.

Rules:
- CI must not require production secrets.
- CI should be path-aware eventually, but simple initial CI is acceptable.
- CI must be safe for PR from feature branches.
- Use branch protection required checks later.

Acceptance Criteria:
- .github/workflows/ci.yml exists.
- CI runs on pull_request and push to develop/staging/main.
- Go test/lint steps exist.
- Frontend steps are conditional or documented if app not initialized.
- No secrets required.

Tests Required:
- CI syntax should be valid.
- Provide expected checks in final response.
```

---

# Task 0.10 — Add Dockerfiles for API Gateway and Service Template

## Prompt

```text
You are working on Sprint 0 for `school-platform`.

Task:
Add Dockerfiles for API Gateway and Go service template.

Goal:
Prepare containerized build for local/staging CI usage.

Targets:
- services/api-gateway/Dockerfile
- services/_template-service/Dockerfile or first service Dockerfile

Scope:
- Multi-stage Dockerfile for Go service.
- Build binary.
- Use minimal runtime image.
- Expose configured port.
- Run as non-root if practical.
- Include healthcheck if appropriate or document health endpoint.

Out of Scope:
- No production hardening beyond reasonable MVP baseline.
- No Kubernetes manifests.
- No image registry workflow yet unless simple.

Rules:
- Do not include .env or secrets in image.
- Use build args only if needed.
- Keep Dockerfile generic and consistent.
- Use static binary if practical.

Acceptance Criteria:
- Docker image builds for API Gateway.
- Docker image builds for template/sample service.
- Container starts locally.
- /healthz works in container.
- No secrets baked into image.

Tests Required:
- docker build command.
- docker run or docker compose verification.
```

---

# Task 0.11 — Add Basic Local Observability Skeleton

## Prompt

```text
You are working on Sprint 0 for `school-platform`.

Task:
Add basic local observability skeleton.

Goal:
Prepare local logging/metrics structure without overbuilding full observability.

Scope:
- Structured JSON logging with slog in API Gateway and sample service.
- request_id and correlation_id in logs.
- Placeholder /metrics endpoint or documented future metrics path.
- Optional Docker Compose placeholders for:
  - Prometheus
  - Grafana
  - Loki
- infra/observability/README.md

Out of Scope:
- Full distributed tracing.
- Jaeger/Tempo integration.
- Advanced dashboards.
- Production alerting.

Rules:
- No sensitive data in logs.
- Log must include service name and environment.
- Correlation ID must be propagated within request handling.
- Keep observability optional in local compose if heavy.

Acceptance Criteria:
- API Gateway logs structured JSON.
- Logs include request_id and correlation_id.
- Observability README explains future Prometheus/Grafana/Loki usage.
- Optional services are documented and can be enabled later.

Tests Required:
- Run API Gateway and verify log fields.
- Handler test for correlation_id middleware if practical.
```

---

# Task 0.12 — Add Documentation Placeholders

## Prompt

```text
You are working on Sprint 0 for `school-platform`.

Task:
Add documentation placeholders and index for project docs.

Goal:
Make all key architecture docs discoverable from the repository.

Target:
- docs/README.md
- docs/INDEX.md or update docs/README.md
- docs/adr/README.md if ADR folder is used

Scope:
Add links/placeholders for:
- 01-technical-architecture.md
- 02-service-boundary.md
- 03-data-model-mvp.md
- 04-api-contract.md
- 05-event-contract.md
- 06-ui-screen-user-flow.md
- 07-test-plan-acceptance-criteria.md
- 08-coding-standard.md
- 09-ai-agent-rules.md
- 10-sprint-backlog-mvp.md
- 11-github-repository-rules.md
- 12-ai-agent-sprint-prompts.md
- 13-sprint-0-task-prompts.md

Out of Scope:
- Do not rewrite all existing docs.
- Do not create unrelated documents.

Rules:
- Docs must be easy for AI Agent and developers to find.
- Use clear titles.
- Mention which docs are mandatory before coding.

Acceptance Criteria:
- docs/README.md lists all key docs.
- AI Agent Rules are clearly marked mandatory.
- Sprint prompt docs are linked.
- GitHub rules are linked.
```

---

# Task 0.13 — Add Environment Examples

## Prompt

```text
You are working on Sprint 0 for `school-platform`.

Task:
Create .env.example files for root, API Gateway, and service template.

Goal:
Standardize configuration names without exposing secrets.

Targets:
- .env.example
- services/api-gateway/.env.example
- services/_template-service/.env.example or first sample service .env.example

Scope:
Root .env.example:
- PostgreSQL local config
- Redis local config
- RabbitMQ local config
- MinIO local config
- Mailpit local config

API Gateway .env.example:
- SERVICE_NAME
- APP_ENV
- HTTP_PORT
- LOG_LEVEL
- CORS_ALLOWED_ORIGINS
- JWT_PUBLIC_KEY_PATH placeholder
- GRPC target placeholders

Service template .env.example:
- SERVICE_NAME
- APP_ENV
- HTTP_PORT
- GRPC_PORT if used
- DATABASE_URL placeholder
- REDIS_URL
- RABBITMQ_URL
- LOG_LEVEL

Out of Scope:
- No real secrets.
- No production config.
- No private keys.

Rules:
- Use local-safe placeholder values.
- Add comments where helpful.
- Make it clear production secrets must be configured through environment/secrets manager.

Acceptance Criteria:
- .env.example files exist.
- No real credentials.
- Variables align with config loader.
- README references .env setup.
```

---

# Task 0.14 — Add Initial GitHub Repository Support Files

## Prompt

```text
You are working on Sprint 0 for `school-platform`.

Task:
Add initial GitHub repository support files.

Goal:
Prepare repository for PR-based workflow, review, and branch protection.

Targets:
- .github/pull_request_template.md
- .github/CODEOWNERS
- .github/workflows/ci.yml if not already created
- .gitignore

Scope:
PR template should include:
- Summary
- Affected Area
- Changes
- How to Test
- Checklist
- Screenshots
- Risk
- Rollback Plan

CODEOWNERS should include:
- services/
- services/api-gateway/
- packages/proto/
- packages/openapi/
- packages/events/
- apps/web-admin/
- apps/mobile-app/
- infra/
- deploy/
- .github/workflows/
- docs/

.gitignore should exclude:
- .env
- .env.*
- node_modules
- dist
- build
- coverage
- *.log
- tmp
- .DS_Store
- private keys
- backup files

But allow:
- .env.example

Out of Scope:
- No branch protection setup through API.
- No deployment workflow.

Rules:
- Align with docs/11-github-repository-rules.md.
- Do not reference unknown GitHub usernames unless placeholders are used.
- No secrets.

Acceptance Criteria:
- PR template exists.
- CODEOWNERS exists with placeholders.
- .gitignore protects secrets and build artifacts.
- CI file exists or is referenced if already created.
```

---

# Task 0.15 — Sprint 0 Final Verification

## Prompt

```text
You are completing Sprint 0 final verification for `school-platform`.

Task:
Review Sprint 0 deliverables and produce a verification report.

Goal:
Confirm that Sprint 0 is ready before starting Sprint 1.

Scope:
Check:
- Monorepo structure
- Docker Compose dependencies
- API Gateway skeleton
- Service template
- Shared packages skeleton
- Proto/OpenAPI/Event placeholders
- Makefile
- GitHub Actions CI
- .env.example files
- README/local development docs
- GitHub PR/CODEOWNERS/.gitignore support files
- Basic logging/correlation ID
- Health/readiness endpoints

Out of Scope:
- Do not implement new features unless very small missing fixes are necessary.
- Do not start Sprint 1 Identity work.

Acceptance Criteria:
- `docker compose up` works.
- API Gateway /healthz works.
- API Gateway /readyz works.
- Sample service /healthz works if sample service exists.
- `make test` works or clearly reports no test modules yet.
- `make lint` works or has documented placeholder.
- CI config exists.
- No secrets committed.
- Docs are discoverable.

Output:
Create a Sprint 0 verification summary containing:
- What passed
- What failed
- Missing items
- Recommended fixes before Sprint 1
- Commands used
```

---

## Final Notes for Sprint 0

Sprint 0 should stay focused on foundation.

Do not implement:

```text
auth
students
finance
PPDB
academic
reporting
notification
```

Those belong to later sprints.

Sprint 0 success means:

```text
The repository is ready for disciplined development.
The local environment is runnable.
AI Agent has structure and guardrails.
The team can start Sprint 1 without reorganizing the project.
```

## Final Planning Context Before Implementation

Before using this task prompt for coding, read:

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/25-product-requirement-document.md
docs/26-development-plan.md
docs/27-workflow.md
docs/25-github-project-management.md
docs/09-ai-agent-rules.md
docs/08-coding-standard.md
docs/29-sprint-0-plan.md if it exists
docs/13-sprint-0-task-prompts.md
```

Rules:

- Use final PRD as product scope reference.
- Use Development Plan as sprint and delivery reference.
- Use Workflow as daily SOP reference.
- Use GitHub Project Management guide for issue/PR/QA/release tracking.
- Implement only one selected issue/task at a time.

## Sprint Plan and GitHub Setup Context

Before using this task prompt for implementation, read:

```text
docs/29-sprint-0-plan.md
docs/40-github-repository-setup-labels-project.md
docs/25-github-project-management.md
docs/27-workflow.md
```

Use `docs/29-sprint-0-plan.md` for sprint-level issue plan, dependencies, risks, and exit criteria.

Use `docs/40-github-repository-setup-labels-project.md` for labels, milestones, project fields, and repository workflow alignment.
