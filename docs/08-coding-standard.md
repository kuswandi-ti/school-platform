# 08 — Coding Standard and Project Convention

Project: `school-platform`  
Status: Final decision for MVP  
Scope: Monorepo, Go backend, Next.js web, Flutter mobile, naming, database, API, logging, testing, and AI Agent coding guardrails.

---

## 1. Purpose

This document defines the coding standard and project convention for the MVP.

It exists to keep implementation consistent across:

- developers
- AI Agent tasks
- microservices
- frontend apps
- API contracts
- database migrations
- tests

---

## 2. Core Decision

The project uses centralized coding standard and conventions.

Required docs:

```text
docs/CODING_STANDARD.md
docs/PROJECT_STRUCTURE.md
docs/API_CONVENTION.md
docs/AI_AGENT_RULES.md
```

For this repository, this file is the main coding standard reference.

---

## 3. Language and Naming

Internal code naming:

```text
English
```

UI labels:

```text
Bahasa Indonesia
```

Examples:

Internal code:

```text
student
teacher
guardian
academic_year
semester
bill
payment
report_card
announcement
approval_request
audit_log
```

UI labels:

```text
Siswa
Guru
Orang Tua/Wali
Tahun Ajaran
Semester
Tagihan
Pembayaran
Rapor
Pengumuman
```

Rule:

```text
Do not use Indonesian words in internal package, endpoint, database table, event type, or permission code.
```

---

## 4. Monorepo Structure

Recommended repository structure:

```text
school-platform/
├── apps/
│   ├── web-admin/
│   └── mobile-app/
│
├── services/
│   ├── api-gateway/
│   ├── identity-service/
│   ├── school-core-service/
│   ├── admission-service/
│   ├── academic-service/
│   ├── finance-service/
│   ├── communication-service/
│   └── reporting-service/
│
├── packages/
│   ├── proto/
│   ├── openapi/
│   ├── events/
│   ├── shared-go/
│   └── docs/
│
├── infra/
│   ├── docker/
│   ├── nginx/
│   ├── postgres/
│   ├── redis/
│   ├── rabbitmq/
│   ├── minio/
│   └── observability/
│
├── deploy/
│   ├── staging/
│   └── production/
│
├── docs/
├── scripts/
├── docker-compose.yml
├── Makefile
└── README.md
```

---

## 5. Go Service Structure

Every Go service must follow this structure:

```text
services/<service-name>/
├── cmd/
│   ├── server/
│   │   └── main.go
│   └── worker/
│       └── main.go
│
├── internal/
│   ├── app/
│   ├── config/
│   ├── domain/
│   ├── usecase/
│   ├── repository/
│   ├── transport/
│   │   ├── grpc/
│   │   └── http/
│   ├── event/
│   ├── authz/
│   ├── audit/
│   └── db/
│       ├── queries/
│       ├── sqlc/
│       └── migrations/
│
├── tests/
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

Rules:

```text
Business logic belongs in domain/usecase.
Transport layer must not contain business logic.
API Gateway must not contain domain business logic.
Repository layer must not perform authorization decisions.
Usecase layer must perform business rules and authorization coordination.
```

---

## 6. Backend Go Stack

Final MVP stack:

```text
HTTP router/API Gateway : Chi
Internal RPC            : gRPC + protobuf
PostgreSQL driver       : pgx
SQL generator           : sqlc
Migration tool          : goose
Logger                  : slog
Config                  : environment variable + envconfig/simple loader
Request validation      : go-playground/validator
Testing                 : Go testing package + testify
Mocking                 : manual mock / mockery when needed
Redis                   : go-redis
RabbitMQ                : amqp091-go
UUID                    : google/uuid
Decimal/money           : shopspring/decimal
Quality gate            : golangci-lint, gofmt, go vet, go test
Task runner             : Makefile
```

Not used in MVP:

```text
Gin
mixed migration tools
float for finance calculation
```

---

## 7. Go Style and Tooling

Required commands:

```bash
go fmt ./...
go vet ./...
go test ./...
golangci-lint run
```

Rules:

```text
All Go code must be gofmt-formatted.
No unused code.
No panic for normal business errors.
context.Context must be passed to database and external calls.
Errors must preserve useful context.
```

---

## 8. Database Access

Rules:

```text
Use pgx + sqlc.
All queries must be parameterized.
Do not concatenate SQL using user input.
Each service accesses only its own database.
No cross-service database query.
All resource queries must filter by foundation_id and school_id when relevant.
```

Example sqlc query:

```sql
-- name: GetBillByID :one
SELECT *
FROM student_bills
WHERE id = $1
  AND foundation_id = $2
  AND school_id = $3;
```

---

## 9. Database Naming Convention

Use snake_case.

Table names are plural:

```text
students
student_guardians
student_bills
student_payments
report_cards
approval_requests
audit_logs
```

Column names:

```text
foundation_id
school_id
academic_year_id
created_at
updated_at
deleted_at
```

Status enums:

```text
pending_verification
revision_requested
partially_paid
```

Required fields for domain tables:

```text
id UUID PRIMARY KEY
foundation_id UUID NOT NULL
school_id UUID NULL/NOT NULL depending on domain
created_at TIMESTAMP
updated_at TIMESTAMP
status VARCHAR where applicable
```

---

## 10. Migration Convention

Migration tool:

```text
goose
```

Rules:

```text
Each service owns its own migrations.
Service only migrates its own database.
Migration files must be ordered.
Avoid destructive migration in production.
Large migrations require backup/snapshot first.
Do not mix goose with golang-migrate.
```

Example:

```text
000001_create_fee_types.sql
000002_create_student_bills.sql
000003_create_student_payments.sql
```

Migration format:

```sql
-- +goose Up
CREATE TABLE fee_types (
    id UUID PRIMARY KEY,
    foundation_id UUID NOT NULL,
    school_id UUID NULL,
    code VARCHAR(50) NOT NULL,
    name VARCHAR(150) NOT NULL,
    status VARCHAR(30) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE fee_types;
```

---

## 11. Config Management

Config source:

```text
Environment variables
```

Rules:

```text
Each service has .env.example.
Config is validated at startup.
Missing required config fails startup with clear error.
No hardcoded production config.
Secrets are not committed to Git.
```

Example:

```text
SERVICE_NAME=finance-service
APP_ENV=local
HTTP_PORT=8080
GRPC_PORT=9090
DATABASE_URL=postgres://...
REDIS_URL=redis://...
RABBITMQ_URL=amqp://...
LOG_LEVEL=info
```

---

## 12. Logging Standard

Logger:

```text
slog
```

Use structured JSON logs.

Required fields:

```text
timestamp
level
service
environment
request_id
correlation_id
actor_user_id
foundation_id
school_id
message
error
```

Rules:

```text
Do not log password.
Do not log access token.
Do not log refresh token.
Do not log Confidential detail.
Mask sensitive identifiers when needed.
Audit log is separate from application log.
```

---

## 13. Error Handling

Use application error type with standard codes.

Standard codes:

```text
UNAUTHORIZED
FORBIDDEN
VALIDATION_ERROR
NOT_FOUND
CONFLICT
BUSINESS_RULE_VIOLATION
APPROVAL_REQUIRED
RESOURCE_LOCKED
RATE_LIMITED
INTERNAL_ERROR
SERVICE_UNAVAILABLE
```

Rules:

```text
Domain/usecase returns application error.
Transport maps application error to HTTP/gRPC status.
Do not expose stack trace to frontend.
Log internal error with correlation_id.
```

---

## 14. API Response Convention

External REST API response:

```json
{
  "data": {},
  "meta": null,
  "error": null
}
```

Error response:

```json
{
  "data": null,
  "meta": null,
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Data yang dikirim tidak valid.",
    "details": {}
  }
}
```

Rule:

```text
API Gateway standardizes external response format.
```

---

## 15. Authorization Convention

Every usecase accessing protected data must receive actor context.

Suggested structure:

```go
type ActorContext struct {
    UserID        string
    FoundationID  string
    SchoolID      *string
    Roles         []string
    Permissions   []string
    Scope         map[string]any
    RequestID     string
    CorrelationID string
}
```

Rules:

```text
Authorization is not only frontend.
Authorization is not only API Gateway.
Internal services must check permission and scope.
Object-level authorization is required for resource by ID.
Queries must filter tenant/school scope.
```

---

## 16. Audit Convention

Audit action format:

```text
domain.resource.action
```

Examples:

```text
finance.payment.verified
academic.report_card.published
identity.role.assigned
school.student.updated
```

Required audit fields:

```text
actor_user_id
foundation_id
school_id
action
entity_type
entity_id
old_values_json
new_values_json
reason
request_id
correlation_id
occurred_at
```

Rules:

```text
Sensitive actions must create audit log.
Audit log is different from application log.
Sensitive data must be masked.
```

---

## 17. Event Convention

Rules:

```text
Use standard event envelope.
Use event type constants.
Use typed event payload struct.
Publish via event publisher abstraction.
Important services use outbox pattern.
Do not publish event directly from HTTP handler.
```

Example:

```go
const EventPaymentVerified = "finance.payment.verified"
```

---

## 18. Numbering Convention

Rules:

```text
Use shared numbering package.
Do not generate document numbers manually in usecase.
Number is scoped by foundation_id, school_id, system_key, period_key.
Used numbers must not be reused.
Use transaction/row lock to avoid duplicates.
```

Example system keys:

```text
invoice
payment
receipt
admission_registration
student_internal_id
outgoing_letter
```

---

## 19. File Convention

Rules:

```text
Files are private by default.
Storage path includes foundation_id and school_id when relevant.
File metadata is stored in owner service.
Signed URL is generated only after permission/scope check.
Official files must not be overwritten.
Revision creates new version.
```

---

## 20. Next.js Web Convention

Recommended structure:

```text
apps/web-admin/
├── app/
│   ├── (auth)/
│   ├── (dashboard)/
│   └── api/
├── components/
│   ├── ui/
│   ├── layout/
│   ├── forms/
│   ├── tables/
│   └── feedback/
├── features/
│   ├── auth/
│   ├── students/
│   ├── teachers/
│   ├── admission/
│   ├── finance/
│   ├── academic/
│   ├── communication/
│   └── reporting/
├── lib/
├── stores/
├── schemas/
└── types/
```

Tools:

```text
React Query
Zustand
React Hook Form
Zod
Tailwind CSS
shadcn/ui
```

Rules:

```text
Feature-based folders.
Reusable UI goes to components.
API calls go through central api-client.
Form validation uses Zod.
Server state uses React Query.
Global UI/context state uses Zustand.
Permission guard component is required for menus/actions.
```

---

## 21. Flutter Mobile Convention

Recommended structure:

```text
apps/mobile-app/lib/
├── app/
├── core/
│   ├── config/
│   ├── network/
│   ├── storage/
│   ├── auth/
│   ├── errors/
│   └── widgets/
├── features/
│   ├── auth/
│   ├── home/
│   ├── billing/
│   ├── attendance/
│   ├── report_card/
│   ├── announcement/
│   ├── notification/
│   └── profile/
└── shared/
```

Tools:

```text
Riverpod
Dio
Retrofit
Flutter Secure Storage
```

Rules:

```text
Feature-based structure.
No offline write in MVP.
Do not store sensitive data offline unless required and encrypted.
Token storage uses Flutter Secure Storage.
```

---

## 22. API Naming Convention

REST endpoints use English plural nouns.

Examples:

```text
/students
/teachers
/classes
/finance/bills
/finance/payments
/academic/report-cards
```

Action endpoints:

```text
POST /finance/payments/{id}/verify
POST /finance/payments/{id}/reject
POST /academic/report-cards/{id}/publish
POST /approvals/{id}/approve
```

Rules:

```text
Use nouns for resources.
Use verb only for workflow actions.
Do not use Indonesian words in endpoints.
```

---

## 23. Testing Convention

Test names should be explicit.

Examples:

```text
TestVerifyPayment_Success
TestVerifyPayment_RejectsUnauthorizedRole
TestGenerateBills_PreventsDuplicateWithIdempotencyKey
TestPublishReportCard_LocksAfterPublish
```

Rules:

```text
Tests cover success and negative cases.
Sensitive features require permission/scope tests.
Tests must not depend on execution order.
Test data must be created through factories/helpers.
```

---

## 24. AI Agent Coding Guardrails

AI Agent must not:

```text
make query to another service database
put business logic in API Gateway
remove foundation_id or school_id
create endpoint without permission/scope check
make private personal file public
write token/password/Confidential data to log
change API/proto/event contract without updating docs
create features outside task scope
ignore tests
```

AI Agent must:

```text
follow service boundary
add tests
add audit for sensitive actions
use standard response/error format
use shared package for logging, audit, numbering, event, file if available
explain changed files and how to test
```

---

## 25. Final Summary

Coding standard for MVP:

```text
Internal code in English.
UI labels in Indonesian.
Service structure is consistent.
Business logic belongs in domain/usecase.
API Gateway is not business logic layer.
Database access uses pgx + sqlc.
Migration uses goose.
Logging uses slog.
Finance calculation does not use float.
AI Agent must follow guardrails and add tests.
```

## Git Commit Convention

Project `school-platform` menggunakan format commit berbasis Conventional Commit agar riwayat perubahan mudah dibaca, mudah ditelusuri ke GitHub Issue/PR, dan dapat mendukung release notes.

### Format Commit

Gunakan format:

```text
type(scope): short description
```

Contoh:

```text
feat(identity): add refresh token rotation
fix(finance): prevent duplicate bill generation
docs(workflow): add git commit convention
chore(ci): add repository validation workflow
test(academic): add attendance scope tests
refactor(api-gateway): simplify request context middleware
```

### Commit Type

Gunakan `type` berikut:

| Type | Kapan Dipakai |
|---|---|
| `feat` | Penambahan fitur baru |
| `fix` | Perbaikan bug |
| `docs` | Perubahan dokumentasi |
| `chore` | Maintenance, dependency, housekeeping |
| `refactor` | Refactor tanpa perubahan behavior |
| `test` | Penambahan/perubahan test |
| `ci` | Perubahan GitHub Actions atau CI/CD |
| `build` | Build system, Docker, dependency packaging |
| `perf` | Peningkatan performa |
| `security` | Security/privacy hardening atau fix |
| `revert` | Revert commit sebelumnya |

### Commit Scope

Gunakan `scope` sesuai area project.

Recommended scopes:

```text
api-gateway
identity
school-core
admission
academic
finance
communication
reporting
web-admin
mobile
infra
ci
docs
security
observability
file-management
shared-go
proto
openapi
events
```

Contoh:

```text
feat(school-core): add student class assignment
fix(admission): prevent duplicate applicant conversion
security(file-management): enforce signed url authorization
ci(repository): block committed env files
docs(github): add label setup guide
```

### Subject / Short Description

Rules:

```text
- Gunakan Bahasa Inggris.
- Gunakan lowercase setelah titik dua.
- Gunakan kalimat pendek dan jelas.
- Gunakan imperative mood jika memungkinkan.
- Maksimal disarankan 72 karakter.
- Jangan akhiri dengan titik.
```

Contoh baik:

```text
feat(finance): add bill generation snapshot
```

Contoh kurang baik:

```text
feat(finance): Added the billing feature.
update file
fix bug
```

### Commit Body

Gunakan commit body jika perubahan butuh penjelasan tambahan.

Format:

```text
type(scope): short description

Longer explanation of what changed and why.
Mention constraints, risk, or migration notes if needed.
```

Contoh:

```text
feat(finance): add bill generation snapshot

Store applied fee policy snapshot on bill items to prevent
future fee policy changes from changing historical bills.
```

### Issue Reference

Jika commit terkait issue tertentu, tambahkan reference di body atau footer.

```text
Refs #123
```

Untuk menutup issue, gunakan di PR description, bukan wajib di commit:

```text
Closes #123
```

### Breaking Change

Jika ada breaking change, gunakan footer:

```text
BREAKING CHANGE: explain what changed and migration required
```

Contoh:

```text
refactor(identity): change actor context payload

BREAKING CHANGE: API Gateway must forward actor context using X-Actor-Context header.
```

Breaking change harus diberi label:

```text
risk: breaking-change
```

dan wajib human review.

### Migration Commit

Jika commit menambahkan migration, gunakan scope service terkait.

Contoh:

```text
feat(identity): add users and sessions migrations
feat(finance): add student bills migrations
```

Migration commit wajib mencantumkan risiko jika menyentuh data existing.

### Security Commit

Untuk perubahan security/privacy:

```text
security(identity): hash refresh tokens at rest
security(file-management): restrict signed url access by scope
```

Security commit wajib review dengan label:

```text
review: security
```

### Commit Granularity

Rules:

```text
- Satu commit sebaiknya merepresentasikan satu perubahan logis.
- Jangan campur refactor besar dengan feature.
- Jangan campur formatting massal dengan business logic.
- Jangan commit file hasil build/cache kecuali memang diperlukan.
- Jangan commit .env, private key, secret, dump database, atau file sensitif.
```

### Commit Before PR

Sebelum push:

```bash
git status
git diff
git add .
git commit -m "feat(scope): short description"
git push origin feature/branch-name
```

Jika perlu body:

```bash
git commit
```

lalu tulis:

```text
feat(scope): short description

Detailed explanation.
```

### Relationship with PR Title

PR title sebaiknya mengikuti format yang sama:

```text
type(scope): short description
```

Jika branch memiliki banyak commit kecil, PR tetap harus memiliki title yang mewakili hasil akhir.

### Squash Merge

Jika menggunakan squash merge, pastikan squash commit mengikuti format:

```text
type(scope): short description
```

Contoh:

```text
feat(identity): implement login and refresh token flow
```

### Commit Review Checklist

Sebelum push/PR:

```text
- [ ] Commit message mengikuti format `type(scope): description`
- [ ] Scope sesuai area yang berubah
- [ ] Subject jelas dan singkat
- [ ] Tidak ada secret atau .env
- [ ] Tidak ada file build/cache yang tidak perlu
- [ ] Commit tidak mencampur perubahan tidak terkait
- [ ] Breaking change diberi catatan jika ada
- [ ] Migration/security change diberi konteks jika perlu
```


### Coding Standard Relationship

Commit convention is part of engineering standard because it affects:

```text
- traceability
- release notes
- code review clarity
- auditability of sensitive changes
- AI Agent output review
```

Any AI Agent-generated commit suggestion must also follow this convention.
