# Development Plan — School Platform MVP

Project: `school-platform`  
Document Type: Development Plan  
Target Audience: Developer, QA, DevOps, Product Owner, AI Agent  
Status: Draft for MVP Implementation  
Repository Target Path: `docs/26-development-plan.md`

---

## 1. Executive Summary

Development Plan ini menjelaskan strategi implementasi MVP `school-platform` secara teknis, bertahap, dan praktis.

`school-platform` adalah platform internal yayasan sekolah multi-unit untuk TK, SD, SMP, dan SMA. MVP dirancang untuk mendukung proses inti yayasan/sekolah, yaitu Identity & Access, School Core, File Management + Import Excel, PPDB, Finance/SPP manual, Academic Basic, Report Card/E-Rapor Basic, Communication/Notification, Reporting Dashboard, serta Security, Observability, Backup, dan UAT Hardening.

Strategi delivery menggunakan pendekatan:

```text
Monorepo
→ Go microservices
→ Custom Go API Gateway
→ PostgreSQL database per service
→ gRPC + protobuf internal communication
→ RabbitMQ domain events
→ Next.js web admin
→ Flutter mobile app
→ Docker Compose local/staging awal
→ GitHub Actions CI/CD
→ branch workflow feature/* → develop → staging → main
```

Development dilakukan melalui sprint bertahap dari Sprint 0 sampai Sprint 10. Setiap sprint memiliki objective, scope, deliverables, acceptance criteria, Definition of Done, dan quality gate. Semua pekerjaan harus dilacak melalui GitHub Issues, GitHub Project, Pull Request, CI, review, QA, dan release workflow.

AI Agent dapat digunakan sebagai assistant untuk membantu task kecil yang sudah jelas scope-nya. Namun AI Agent bukan final authority. Output AI Agent wajib direview manusia, tidak boleh menangani production secrets, tidak boleh membuat keputusan final security/legal/compliance, dan tidak boleh mengubah arsitektur tanpa instruksi eksplisit.

MVP dianggap siap pilot/production jika core flow berjalan, tidak ada Critical/High bug pada alur utama, permission/scope dan object-level authorization sudah diuji, audit log tersedia untuk aksi sensitif, backup/restore sudah diuji, observability dasar tersedia, QA/UAT pass, dan deployment production melalui approval.

---

## 2. Development Principles

Prinsip development yang wajib diikuti:

### 2.1 Local-First Development

Semua developer harus dapat menjalankan lingkungan lokal menggunakan Docker Compose. Local development menjadi tempat utama untuk menjalankan service, database, Redis, RabbitMQ, MinIO, dan dependency pendukung sebelum perubahan masuk ke staging.

Rules:

```text
- Setiap service harus memiliki .env.example.
- Local setup tidak boleh membutuhkan production secret.
- External provider harus dimock/log-only di local.
- Production data tidak boleh digunakan di local kecuali sudah dianonimkan dan disetujui.
```

### 2.2 Small Task Delivery

Pekerjaan harus dipecah menjadi issue kecil yang jelas scope-nya.

Setiap issue harus memiliki:

```text
- objective
- scope
- out of scope
- acceptance criteria
- area
- priority
- sprint/milestone
- risk
- owner
- AI Agent status jika relevan
```

### 2.3 Service Boundary Discipline

Setiap service hanya boleh mengakses database miliknya sendiri.

Non-negotiable rules:

```text
- Tidak boleh query database service lain.
- Cross-service communication harus melalui gRPC atau domain events.
- API Gateway tidak boleh berisi business logic.
- Reporting tidak boleh query operational database secara langsung.
```

### 2.4 Test-Driven Quality Gate

Setiap perubahan harus memiliki test sesuai risiko dan scope.

Minimum:

```text
- unit test untuk business logic
- integration test untuk repository/service boundary
- API test untuk endpoint penting
- permission/scope test untuk akses data
- event test untuk publish/consume event
- audit test untuk aksi sensitif
```

### 2.5 Security From Start

Security tidak boleh hanya dikerjakan di akhir MVP. Setiap sprint harus memperhatikan:

```text
- authentication
- authorization
- object-level authorization
- field-level access jika diperlukan
- audit log
- privacy
- file access
- no sensitive logs
- no secrets in repo
```

### 2.6 Audit and Permission From Early Sprint

Permission/scope dan audit harus dimasukkan sejak Sprint 1 dan diperluas di sprint berikutnya.

Aksi sensitif wajib diaudit, terutama:

```text
- role assignment
- student sensitive update
- PPDB decision
- import confirm
- payment verification
- payment void
- fee policy approval
- report card publish
- report revision after publish
- restricted/confidential file download/export
```

### 2.7 Event-Driven Reporting and Notification

Operational service menerbitkan domain event. Reporting dan Communication mengonsumsi event tersebut.

Rules:

```text
- Business service tidak boleh memanggil notification provider langsung.
- Reporting Service hanya membaca reporting_db.
- Event consumer harus idempotent.
- Event payload tidak boleh membawa Confidential data mentah.
```

### 2.8 No Cross-Service Database Query

Cross-service query langsung dilarang.

Jika service membutuhkan data service lain:

```text
- gunakan gRPC untuk request/response synchronous;
- gunakan domain event untuk projection/asynchronous read model;
- gunakan ID reference dan snapshot jika diperlukan.
```

### 2.9 Documentation as Source of Truth

Dokumen menjadi acuan utama untuk developer, QA, DevOps, dan AI Agent.

Jika implementasi mengubah kontrak, workflow, API, event, data model, atau sprint scope, dokumentasi harus diupdate.

### 2.10 AI Agent With Human Review

AI Agent boleh membantu:

```text
- membuat draft kode untuk task kecil
- membuat test
- membuat dokumentasi
- membuat migration draft
- membuat prompt/task breakdown
```

AI Agent tidak boleh:

```text
- menangani production secrets
- melakukan final security approval
- membuat keputusan legal/compliance
- melakukan production deployment approval
- mengubah arsitektur besar tanpa instruksi
- mengakses data asli sensitif
```

---

## 3. Team and Responsibilities

| Role | Main Responsibilities | Deliverables | Review Responsibility |
|---|---|---|---|
| Backend Developer | Membangun Go services, API Gateway, gRPC, database, event, business logic, tests | Services, migrations, sqlc queries, gRPC/proto, event producer/consumer, backend tests | Service boundary, business logic, permission/scope, audit, event correctness |
| Frontend Developer | Membangun Next.js web admin | Web screens, forms, API integration, state management, validation, frontend tests | UI flow, permission-based menu, usability, API integration |
| QA | Membuat test plan, test case, regression, UAT, bug validation | QA scenarios, bug reports, UAT checklist, sign-off | Acceptance criteria, regression risk, release readiness |
| Infrastructure/DevOps | Menyiapkan Docker Compose, CI/CD, environment, observability, backup, deployment | Docker Compose, GitHub Actions, environment docs, backup/restore scripts, monitoring setup | CI/CD, deployment safety, secrets, rollback, backup |
| AI Agent | Membantu task kecil yang jelas scope-nya | Draft code/docs/tests sesuai prompt | Tidak menjadi final authority; output wajib human review |
| Product Owner/Reviewer | Menjaga scope, prioritas, acceptance criteria, UAT direction | PRD, backlog, sprint acceptance, release sign-off | Product fit, business flow, MVP boundary |

---

## 4. Repository and Branching Strategy

### 4.1 Repository

Repository utama:

```text
school-platform
```

Struktur monorepo yang diharapkan:

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
├── .github/
├── AGENTS.md
├── SKILLS.md
├── README.md
└── Makefile
```

### 4.2 Branches

| Branch | Purpose |
|---|---|
| `develop` | Development integration branch |
| `staging` | QA/UAT branch and release candidate |
| `main` | Production branch |
| `feature/*` | Feature/task branch |
| `fix/*` | Bug fix branch |
| `docs/*` | Documentation branch |
| `chore/*` | Maintenance branch |
| `refactor/*` | Refactor branch |
| `test/*` | Test-only branch |
| `hotfix/*` | Urgent production fix from main |

### 4.3 Branch Flow

```text
feature/* → develop
fix/*     → develop
docs/*    → develop
chore/*   → develop
develop   → staging
staging   → main
hotfix/*  → main → staging/develop
```

### 4.4 Pull Request Rule

Semua perubahan wajib melalui Pull Request.

Minimum PR rule:

```text
- linked issue
- PR template diisi
- CI pass
- at least 1 approval
- no unresolved conversation
- no direct push to protected branches
```

### 4.5 CI

CI dijalankan pada:

```text
pull_request to develop/staging/main
push to develop/staging/main
```

CI minimum:

```text
- repository check
- required docs check
- no .env/secrets committed
- YAML validation
- Go fmt/vet/test if Go modules exist
- Web lint/test/build if web app exists
- Flutter analyze/test if mobile app exists
- Docker Compose config validation if compose files exist
```

### 4.6 Branch Protection

| Branch | Protection |
|---|---|
| `develop` | PR required, 1 approval, CI required, no force push, no deletion |
| `staging` | PR required, CI required, QA/UAT required, no force push, no deletion |
| `main` | PR required, approval required, CI required, production environment approval, no force push, no deletion |

### 4.7 Release Tag

Production release tag format:

```text
v0.1.0
v0.1.1
v1.0.0
```

Pre-release/release candidate tag format:

```text
v0.1.0-rc.1
```

### 4.8 Hotfix Workflow

Hotfix digunakan hanya untuk urgent production issue.

Flow:

```text
main
→ hotfix/fix-critical-issue
→ PR to main
→ production deploy
→ tag patch release
→ back-merge main to staging
→ back-merge main to develop
```

---

## 5. GitHub Project Management

### 5.1 GitHub Project

Project name:

```text
School Platform MVP
```

Purpose:

```text
- backlog management
- sprint tracking
- AI Agent task coordination
- QA/UAT tracking
- release readiness
```

### 5.2 Labels

Label categories:

```text
type
area
sprint
priority
status
ai
risk
review
```

Recommended examples:

```text
type: feature
type: bug
area: identity
area: finance
sprint: 0
priority: high
status: ready
ai: ready
risk: data-sensitive
review: security
```

### 5.3 Milestones

Milestones:

```text
Sprint 0 — Project Foundation
Sprint 1 — Identity & Access
Sprint 2 — School Core
Sprint 3 — File Management + Import Excel
Sprint 4 — PPDB
Sprint 5 — Finance / SPP
Sprint 6 — Academic Basic
Sprint 7 — Report Card / E-Rapor Basic
Sprint 8 — Communication / Notification
Sprint 9 — Reporting Dashboard
Sprint 10 — Security, Observability, Backup, UAT Hardening
MVP Release
```

### 5.4 Issue Lifecycle

```text
Backlog
→ Ready
→ In Progress
→ In Review
→ QA
→ Done
```

Exception:

```text
Blocked
```

### 5.5 PR Relation to Issue

Every PR must link to an issue using:

```text
Closes #123
Fixes #123
Related #123
```

Rules:

```text
- One PR should ideally close one small issue.
- Large PR must be split unless explicitly approved.
- PR must use `.github/pull_request_template.md`.
```

### 5.6 AI Agent Labels

AI Agent task status:

| Label | Meaning |
|---|---|
| `ai: ready` | AI Agent can work on this task |
| `ai: needs-context` | Task needs more detail |
| `ai: generated` | Output generated/assisted by AI Agent |
| `ai: needs-human-review` | Human review required |
| `ai: do-not-use-agent` | AI Agent must not be used |

### 5.7 GitHub Project Fields

| Field | Type | Values |
|---|---|---|
| Status | Single select | Backlog, Ready, In Progress, In Review, QA, Blocked, Done |
| Sprint | Single select | Sprint 0–Sprint 10, MVP Release |
| Priority | Single select | Critical, High, Medium, Low |
| Area | Single select | api-gateway, identity, school-core, admission, academic, finance, communication, reporting, web-admin, mobile, infra, docs, security, observability, ci-cd, file-management |
| Type | Single select | feature, bug, chore, docs, refactor, test, security, infra, spike, hotfix |
| Owner | User | GitHub assignee |
| Estimate | Number | 1, 2, 3, 5, 8 |
| Risk | Single select | Low, Medium, High, Breaking Change, Migration, Data Sensitive |
| Platform | Single select | Backend, Web, Mobile, Infra, Docs, QA, Product |
| AI Agent | Single select | Ready, Needs Context, Generated, Needs Human Review, Do Not Use |
| Target Release | Text / Single select | MVP, v0.1, v1.0 |

---

## 6. Environment Strategy

### 6.1 Local Environment

Purpose:

```text
developer local development
```

Includes:

```text
PostgreSQL
Redis
RabbitMQ
MinIO
Mailpit optional
Grafana/Loki optional
```

Rules:

```text
- Uses Docker Compose.
- No production secrets.
- No production data unless anonymized and approved.
- External providers mocked/log-only.
```

### 6.2 Staging Environment

Purpose:

```text
QA/UAT and release candidate validation
```

Rules:

```text
- Deploy from staging branch.
- CI must pass.
- QA/UAT runs here.
- Secrets separated from production.
- Data should be test/staging data.
```

### 6.3 Production Environment

Purpose:

```text
pilot/real operation
```

Rules:

```text
- Deploy from main branch only.
- Manual approval required.
- Production secrets stored in GitHub Environment secrets or secure secret manager.
- Backup before high-risk release/migration.
- Rollback plan required.
```

### 6.4 Docker Compose

Docker Compose is used for:

```text
- local development
- early staging if needed
- dependency orchestration
```

Compose should include:

```text
postgres
redis
rabbitmq
minio
mailpit optional
observability optional
```

### 6.5 GitHub Environments

Required environments:

```text
staging
production
```

Production environment must require reviewer approval.

### 6.6 Secrets Separation

Secrets must be separated by environment:

```text
LOCAL_*
STAGING_*
PRODUCTION_*
```

Rules:

```text
- Never commit .env.
- Never commit keys/certificates.
- Never expose secrets in logs.
```

### 6.7 Local Data Policy

```text
- Use seed data.
- Do not use real production data.
- If production-like data is needed, anonymize it and get approval.
```

### 6.8 Production Deployment Approval

Production deployment requires:

```text
- PR staging → main approved
- CI pass
- QA sign-off
- release notes
- rollback plan
- manual approval through GitHub Environment
```

---

## 7. Development Phases

| Phase | Objective | Related Sprint | Key Deliverables | Risk |
|---|---|---|---|---|
| Phase 1 — Platform Foundation | Menyiapkan fondasi teknis, repo, CI, service template, auth, school core | Sprint 0–2 | Monorepo, local dev, API Gateway, Identity, School Core | Service boundary dan auth harus benar sejak awal |
| Phase 2 — Admission & Finance | Mendukung PPDB dan SPP manual | Sprint 3–5 | File/import, PPDB, finance billing/payment manual | Data migration dan finance calculation risk |
| Phase 3 — Academic & Communication | Mendukung proses akademik dan komunikasi | Sprint 6–8 | Schedule, attendance, grade, report card, notification | Workflow akademik dan publish/lock harus konsisten |
| Phase 4 — Reporting & Production Readiness | Dashboard, security hardening, observability, backup, UAT | Sprint 9–10 | Reporting projections, monitoring, backup/restore, QA/UAT | Release readiness, privacy, and operational risk |

---

## 8. Sprint Plan Overview

| Sprint | Objective | Main Modules | Key Deliverables | Dependencies | Owner | Exit Criteria |
|---|---|---|---|---|---|---|
| Sprint 0 | Menyiapkan fondasi project | Repo, CI, Docker, service skeleton | Monorepo, Docker Compose, CI, API Gateway skeleton, service template | None | Backend, DevOps | Local dev jalan, CI basic pass |
| Sprint 1 | Membangun Identity & Access | Identity, API Gateway auth | Login, JWT, refresh rotation, RBAC, context | Sprint 0 | Backend | Auth flow tested, permission baseline ready |
| Sprint 2 | Membangun School Core | Foundation, school, student, teacher, class | Master data dan assignment | Sprint 1 | Backend, Frontend | Core data CRUD dan scope pass |
| Sprint 3 | Membangun File + Import | File, MinIO, import Excel | Private file, signed URL, import validation-preview-confirm | Sprint 2 | Backend, DevOps, QA | Import data awal bisa digunakan |
| Sprint 4 | Membangun PPDB | Admission | Period, applicant, docs, decision, conversion | Sprint 2–3 | Backend, Frontend | Applicant accepted dapat menjadi student |
| Sprint 5 | Membangun Finance/SPP | Finance | Fee, policy, bill, payment proof, verification | Sprint 2–3 | Backend, Frontend, QA | Bill/payment manual core flow pass |
| Sprint 6 | Membangun Academic Basic | Academic | Curriculum, subject, schedule, attendance | Sprint 2 | Backend, Frontend | Guru dapat input absensi sesuai assignment |
| Sprint 7 | Membangun E-Rapor Basic | Academic/report card | Score, grade book, report card publish/lock | Sprint 6 | Backend, Frontend, QA | Rapor published dapat dilihat parent/student |
| Sprint 8 | Membangun Communication/Notification | Communication | Announcement, in-app notification, event consumers | Sprint 1, 2, 5, 7 | Backend, Frontend, Mobile | Notification event-driven berjalan |
| Sprint 9 | Membangun Reporting Dashboard | Reporting | Projections, dashboard APIs, event consumers | Sprint 4–8 | Backend, Frontend | Dashboard sesuai scope tersedia |
| Sprint 10 | Hardening MVP | Security, observability, backup, UAT | Security review, logs, metrics, backup/restore, UAT | Sprint 0–9 | QA, DevOps, Backend | No Critical/High core bug, MVP ready |

---

## 9. Detailed Sprint Plan

### Sprint 0 — Project Foundation

#### Objective

Menyiapkan fondasi repository, local development, CI, service template, API Gateway skeleton, contract folders, dan dokumentasi awal.

#### Scope

- monorepo structure
- Docker Compose local dependencies
- Go service template
- API Gateway skeleton
- shared-go skeleton
- packages/proto, packages/openapi, packages/events
- Makefile
- GitHub Actions basic CI
- repository support files
- `.env.example`
- healthz/readyz baseline
- structured logging baseline

#### Out of Scope

- full authentication
- business modules
- production deployment
- Kubernetes
- domain database schema
- UI feature implementation

#### Backend Tasks

- Create base Go service layout.
- Create API Gateway skeleton.
- Add shared logging/request_id utilities.
- Add healthz/readyz endpoints.

#### Frontend Tasks

- Prepare placeholder app structure if needed.
- No feature UI yet.

#### Mobile Tasks

- Prepare placeholder app structure if needed.
- No feature UI yet.

#### QA Tasks

- Validate repo structure.
- Validate CI runs.
- Validate Docker Compose config.

#### DevOps Tasks

- Add Docker Compose.
- Add GitHub Actions CI.
- Add branch workflow support files.

#### AI Agent Suitable Tasks

- Generate scaffold files.
- Generate documentation.
- Generate basic CI config.
- Generate Makefile commands.

#### Human Review Required Tasks

- CI final validation.
- Repo structure approval.
- Security check for secrets.

#### Dependencies

None.

#### Deliverables

- runnable local foundation
- CI baseline
- docs baseline
- service skeletons

#### Acceptance Criteria

- Repository structure exists.
- Docker Compose config validates.
- CI passes repository checks.
- API Gateway skeleton can run healthz/readyz.
- No `.env` or secrets committed.

#### Definition of Done

- PR merged to develop.
- CI pass.
- docs updated.
- local run instructions available.

#### Risks

| Risk | Mitigation |
|---|---|
| Repo structure wrong early | Review against architecture docs before Sprint 1 |
| CI too strict before code exists | Use conditional checks for Go/Web/Mobile |
| Secrets accidentally committed | Add CI secret/file check |

---

### Sprint 1 — Identity & Access

#### Objective

Membangun autentikasi dan otorisasi dasar.

#### Scope

- identity_db
- users
- password hashing
- login
- JWT access token
- rotating refresh token
- logout/revoke session
- roles
- permissions
- role assignment
- actor context
- API Gateway auth middleware

#### Out of Scope

- OAuth/social login
- 2FA
- advanced anomaly detection
- full profile management

#### Backend Tasks

- Create identity migrations.
- Implement login/refresh/logout.
- Implement RBAC seed.
- Implement actor context.
- Implement auth middleware in API Gateway.

#### Frontend Tasks

- Login screen.
- Auth state.
- Basic protected route handling.

#### Mobile Tasks

- Login flow baseline if mobile app initialized.
- Secure storage integration placeholder.

#### QA Tasks

- Login success/failure tests.
- Refresh token rotation tests.
- Permission access tests.

#### DevOps Tasks

- Add identity env vars.
- Ensure DB migration command works.

#### AI Agent Suitable Tasks

- Draft migrations.
- Draft DTO/handler/service/repository.
- Generate tests.

#### Human Review Required Tasks

- Password/token security.
- Refresh token rotation.
- RBAC/scope model.

#### Dependencies

Sprint 0.

#### Deliverables

- Identity service
- Auth API
- RBAC foundation
- Gateway auth

#### Acceptance Criteria

- User can login.
- Refresh token rotates.
- Token reuse rejected.
- User context includes scope.
- Protected endpoint rejects invalid token.

#### Definition of Done

- Auth tests pass.
- Permission baseline tested.
- No token/password logged.

#### Risks

| Risk | Mitigation |
|---|---|
| Token leakage | Ensure no sensitive logs |
| Incorrect scope model | Review with architecture docs |
| Refresh reuse vulnerability | Add explicit token reuse tests |

---

### Sprint 2 — School Core

#### Objective

Membangun data master yayasan/sekolah/siswa/orang tua/guru/kelas dan assignment.

#### Scope

- foundation
- school
- academic year
- semester
- student
- guardian
- teacher
- grade level
- class
- student-class assignment
- teacher assignment
- homeroom assignment

#### Out of Scope

- PPDB
- Import Excel
- Finance
- Academic grade/report card

#### Backend Tasks

- Create school_core_db migrations.
- Implement CRUD and filters.
- Implement scope checks.
- Publish core events.
- Add audit for sensitive changes.

#### Frontend Tasks

- Master data screens.
- Student/teacher/class list and forms.

#### Mobile Tasks

- No major mobile scope except basic profile context if needed.

#### QA Tasks

- CRUD tests.
- Scope tests.
- Search/filter tests.

#### DevOps Tasks

- Migration support for school-core-service.

#### AI Agent Suitable Tasks

- CRUD scaffolding.
- sqlc query drafts.
- table-driven tests.

#### Human Review Required Tasks

- Data model relationships.
- Scope and object-level auth.

#### Dependencies

Sprint 1.

#### Deliverables

- School Core service
- Web admin master data baseline
- Scope-protected core data

#### Acceptance Criteria

- Data master can be created/updated.
- Access is limited by foundation/school.
- Sensitive changes are audited.

#### Definition of Done

- CRUD tests pass.
- Scope tests pass.
- docs updated if data model changes.

#### Risks

| Risk | Mitigation |
|---|---|
| Data model mismatch across TK/SD/SMP/SMA | Keep flexible grade/class structures |
| Cross-school data leakage | Add object-level auth tests |

---

### Sprint 3 — File Management + Import Excel

#### Objective

Membangun private file management dan import Excel data awal.

#### Scope

- file metadata
- MinIO/S3 abstraction
- private upload
- signed URL
- MIME/extension/size validation
- import batches
- import row validation
- preview
- confirm import
- import report

#### Out of Scope

- import grades
- historical payments import
- virus scanning integration
- central File Service

#### Backend Tasks

- Implement storage abstraction.
- Implement metadata tables.
- Implement import batch tables.
- Implement validation-preview-confirm.
- Implement import report.

#### Frontend Tasks

- Upload file UI.
- Import preview UI.
- Error report UI.

#### Mobile Tasks

- No major mobile scope.

#### QA Tasks

- File validation tests.
- Import success/failure tests.
- Permission tests for file access.

#### DevOps Tasks

- MinIO local setup.
- Bucket/env configuration.

#### AI Agent Suitable Tasks

- Generate import template docs.
- Draft validation test cases.
- Draft file metadata implementation.

#### Human Review Required Tasks

- File privacy rules.
- Import data integrity.

#### Dependencies

Sprint 2.

#### Deliverables

- Private file upload
- Signed URL
- Excel import flow

#### Acceptance Criteria

- File private by default.
- Import validates before writing.
- Error report available.
- No raw sensitive data in logs.

#### Definition of Done

- File tests pass.
- Import tests pass.
- Security review for signed URL.

#### Risks

| Risk | Mitigation |
|---|---|
| Invalid import corrupts data | Use validation-preview-confirm |
| File access leak | Backend authorization before signed URL |

---

### Sprint 4 — PPDB

#### Objective

Membangun proses PPDB dari admission period hingga conversion applicant menjadi student.

#### Scope

- admission period
- applicant
- applicant guardian
- applicant document
- verification
- accept/reject decision
- conversion to student through School Core gRPC
- PPDB events
- audit

#### Out of Scope

- advanced scoring
- public marketing website
- payment gateway
- direct write to school_core_db

#### Backend Tasks

- Create admission_db migrations.
- Implement applicant workflow.
- Implement document integration.
- Implement accept/reject.
- Implement conversion via gRPC.

#### Frontend Tasks

- PPDB admin screens.
- Applicant status/detail screens.
- Document verification UI.

#### Mobile Tasks

- Optional parent applicant view if included.

#### QA Tasks

- Applicant workflow tests.
- Document verification tests.
- Conversion idempotency tests.

#### DevOps Tasks

- Admission service env/migration setup.

#### AI Agent Suitable Tasks

- Draft status enum/workflow tests.
- Draft CRUD handlers.
- Draft event publisher.

#### Human Review Required Tasks

- Conversion flow.
- Decision audit.
- Document privacy.

#### Dependencies

Sprint 2 and Sprint 3.

#### Deliverables

- PPDB workflow
- Conversion to student
- PPDB dashboard data events

#### Acceptance Criteria

- Applicant can be submitted.
- Documents can be verified.
- Accepted applicant can be converted once.
- Rejected applicant has reason.
- Decision is audited.

#### Definition of Done

- PPDB tests pass.
- Conversion idempotency verified.
- Event contract updated.

#### Risks

| Risk | Mitigation |
|---|---|
| Duplicate student conversion | Idempotency key and converted_student_id |
| Document access leak | Reuse private file authorization |

---

### Sprint 5 — Finance / SPP

#### Objective

Membangun Finance/SPP manual payment flow.

#### Scope

- fee type
- fee scheme
- fee policy
- sibling discount
- bill generation
- payment proof upload
- payment verification/rejection
- receipt
- outstanding
- void payment with approval
- finance events
- audit

#### Out of Scope

- payment gateway
- bank reconciliation
- full accounting ledger
- payroll

#### Backend Tasks

- Create finance_db migrations.
- Implement decimal amount handling.
- Implement fee/policy/bill/payment services.
- Implement bill snapshot.
- Implement verification/rejection.
- Implement events/audit.

#### Frontend Tasks

- Finance admin screens.
- Bill/payment verification UI.
- Outstanding report UI.

#### Mobile Tasks

- Parent bill list.
- Upload proof.
- Payment status view.

#### QA Tasks

- Bill generation tests.
- Idempotency tests.
- Payment verification tests.
- Parent-child scope tests.

#### DevOps Tasks

- Finance service migration and env setup.

#### AI Agent Suitable Tasks

- Draft migrations and queries.
- Generate unit tests.
- Draft API handlers.

#### Human Review Required Tasks

- Money calculation.
- Fee policy approval.
- Payment void.
- Privacy/security.

#### Dependencies

Sprint 2 and Sprint 3.

#### Deliverables

- Finance service
- Manual SPP workflow
- Parent payment proof flow

#### Acceptance Criteria

- Bill generated idempotently.
- Bill stores snapshot.
- Amount uses decimal/NUMERIC.
- Parent can upload proof.
- Bendahara can verify/reject.
- Verification audited.

#### Definition of Done

- Finance core tests pass.
- Scope/security tests pass.
- No float for money.

#### Risks

| Risk | Mitigation |
|---|---|
| Wrong billing amount | Decimal, snapshots, table-driven tests |
| Duplicate bills | Unique constraint and idempotency |
| Payment proof leak | Restricted file and signed URL |

---

### Sprint 6 — Academic Basic

#### Objective

Membangun curriculum, subject, schedule, dan attendance dasar.

#### Scope

- curriculum
- subject
- subject group
- class subject
- schedule
- teacher schedule view
- attendance input
- attendance correction

#### Out of Scope

- report card
- full grade book
- LMS
- advanced timetable optimization

#### Backend Tasks

- Create academic_db migrations.
- Implement subject/schedule/attendance.
- Enforce assignment scope.
- Add attendance correction audit.

#### Frontend Tasks

- Academic setup screens.
- Schedule UI.
- Attendance input UI.

#### Mobile Tasks

- Teacher quick attendance optional.
- Student schedule view if included.

#### QA Tasks

- Teacher assignment tests.
- Attendance tests.
- Correction audit tests.

#### DevOps Tasks

- Academic service migration/env setup.

#### AI Agent Suitable Tasks

- Draft migrations.
- Draft CRUD handlers.
- Generate attendance tests.

#### Human Review Required Tasks

- Assignment-based access.
- Correction audit.

#### Dependencies

Sprint 2.

#### Deliverables

- Academic Basic service
- Schedule
- Attendance

#### Acceptance Criteria

- Guru can view assigned schedule.
- Guru can input attendance for assigned class/subject.
- Correction requires reason and audit.

#### Definition of Done

- Assignment scope tests pass.
- Attendance tests pass.

#### Risks

| Risk | Mitigation |
|---|---|
| Teacher sees wrong class | Scope tests |
| Attendance duplicate | Unique constraints |

---

### Sprint 7 — Report Card / E-Rapor Basic

#### Objective

Membangun nilai dan rapor dasar.

#### Scope

- assessment components
- assessment scheme
- grade book
- student score
- report card
- report card item
- review
- publish/lock
- revision after publish
- parent/student view

#### Out of Scope

- national e-rapor integration
- advanced analytics
- LMS
- complex PDF design

#### Backend Tasks

- Implement grade book and score flow.
- Implement report generation.
- Implement publish/lock.
- Implement revision approval.
- Publish report events.

#### Frontend Tasks

- Score input UI.
- Grade book review UI.
- Report publish UI.
- Parent/student report view.

#### Mobile Tasks

- Parent/student published report view.

#### QA Tasks

- Score input tests.
- Publish/lock tests.
- Revision approval tests.
- Parent-child scope tests.

#### DevOps Tasks

- Ensure file/report storage if PDF placeholder used.

#### AI Agent Suitable Tasks

- Draft scoring tests.
- Draft report view models.
- Generate workflow documentation.

#### Human Review Required Tasks

- Publish/lock workflow.
- Revision approval.
- Parent/student privacy.

#### Dependencies

Sprint 6.

#### Deliverables

- Grade book flow
- Report card flow
- Published report view

#### Acceptance Criteria

- Guru inputs score by assignment.
- Wali Kelas reviews.
- Kepala Sekolah publishes.
- Published report locked.
- Parent/student sees only published report.

#### Definition of Done

- Report workflow tests pass.
- Scope tests pass.
- Audit tests pass.

#### Risks

| Risk | Mitigation |
|---|---|
| Published report changed without approval | Lock status and approval tests |
| Parent sees other child report | Parent-child scope tests |

---

### Sprint 8 — Communication / Notification

#### Objective

Membangun announcement dan notification system berbasis event.

#### Scope

- announcements
- announcement targets
- notifications
- templates
- deliveries
- preferences
- event consumers
- in-app notification
- FCM/email abstraction

#### Out of Scope

- WhatsApp
- SMS
- advanced campaign builder

#### Backend Tasks

- Implement communication_db.
- Implement announcement workflow.
- Implement event consumers.
- Implement notification delivery log.
- Implement provider abstraction.

#### Frontend Tasks

- Announcement UI.
- Notification center.
- Read/unread UI.

#### Mobile Tasks

- Notification list.
- Push provider integration placeholder.
- Announcement view.

#### QA Tasks

- Announcement target tests.
- Notification consumer tests.
- Confidential body tests.

#### DevOps Tasks

- Provider env placeholders.
- Queue config.

#### AI Agent Suitable Tasks

- Draft event consumer.
- Draft notification templates.
- Generate tests.

#### Human Review Required Tasks

- Confidential data handling.
- Target audience logic.

#### Dependencies

Sprint 1, Sprint 2, Sprint 5, Sprint 7.

#### Deliverables

- Announcement system
- Notification center
- Event-driven notification

#### Acceptance Criteria

- Announcement can target audience.
- Notification can be read/unread.
- Confidential detail not in body.
- Consumer idempotent.

#### Definition of Done

- Notification tests pass.
- Event consumer tests pass.
- Privacy review done.

#### Risks

| Risk | Mitigation |
|---|---|
| Sensitive data in notification | Template review and tests |
| Duplicate notifications | Idempotent consumer |

---

### Sprint 9 — Reporting Dashboard

#### Objective

Membangun Reporting Service dan dashboard berbasis projection.

#### Scope

- reporting_db
- processed_events
- projection tables
- event consumers
- dashboard APIs
- scheduled rebuild/sync skeleton
- dashboard UI

#### Out of Scope

- advanced BI
- global search
- data warehouse
- direct query to operational DBs

#### Backend Tasks

- Implement reporting migrations.
- Implement event consumers.
- Implement projections.
- Implement dashboard APIs.
- Implement idempotency.

#### Frontend Tasks

- Dashboard Yayasan.
- Dashboard Sekolah.
- Dashboard Guru.
- Parent/student summary if included.

#### Mobile Tasks

- Parent/student summary dashboard if included.

#### QA Tasks

- Projection tests.
- Dashboard scope tests.
- Event idempotency tests.

#### DevOps Tasks

- Reporting consumer/worker setup.
- Monitoring queue health.

#### AI Agent Suitable Tasks

- Draft projection tables.
- Draft dashboard DTOs.
- Generate tests.

#### Human Review Required Tasks

- Reporting boundary.
- Dashboard privacy/scope.
- Event projection correctness.

#### Dependencies

Sprint 4–8.

#### Deliverables

- Reporting Service
- Dashboard APIs
- Web dashboard

#### Acceptance Criteria

- Reporting reads reporting_db only.
- Projections update from events.
- Duplicate events skipped.
- Dashboard follows role/scope.

#### Definition of Done

- Projection tests pass.
- Dashboard scope tests pass.

#### Risks

| Risk | Mitigation |
|---|---|
| Reporting uses operational DB | Code review and architecture check |
| Dashboard stale | Scheduled rebuild/sync skeleton |

---

### Sprint 10 — Security, Observability, Backup, UAT Hardening

#### Objective

Melakukan hardening MVP sebelum pilot/production.

#### Scope

- security review
- permission/scope regression
- object-level authorization test
- audit log review
- structured logging review
- Prometheus metrics
- Grafana/Loki setup
- RabbitMQ monitoring
- backup script
- restore procedure
- restore test
- UAT checklist
- regression checklist
- release readiness
- rollback documentation
- bug fixing

#### Out of Scope

- Kubernetes
- advanced SIEM
- full penetration test unless separately planned
- new feature development

#### Backend Tasks

- Fix security/test gaps.
- Add missing audit logs.
- Add missing metrics.
- Validate permission/scope.

#### Frontend Tasks

- Fix UAT findings.
- Improve validation/UX issues.
- Finalize release-critical screens.

#### Mobile Tasks

- Fix UAT findings.
- Validate parent/student flows.

#### QA Tasks

- Regression test.
- UAT execution.
- Release sign-off.
- Bug verification.

#### DevOps Tasks

- Backup/restore test.
- Monitoring setup.
- Production checklist.
- Rollback plan.

#### AI Agent Suitable Tasks

- Generate checklist.
- Draft docs.
- Generate missing test scaffolds.

#### Human Review Required Tasks

- Security sign-off.
- Production approval.
- Backup/restore verification.
- Release decision.

#### Dependencies

Sprint 0–9.

#### Deliverables

- Release readiness package
- QA/UAT sign-off
- Backup/restore proof
- Observability baseline

#### Acceptance Criteria

- No Critical/High bug in MVP core flow.
- Restore test performed.
- Logs do not expose sensitive data.
- Production release checklist pass.

#### Definition of Done

- MVP readiness criteria met.
- Production approval ready.
- Release notes prepared.

#### Risks

| Risk | Mitigation |
|---|---|
| Late security findings | Security review before Sprint 10 where possible |
| Backup restore fails | Test restore before release |
| UAT discovers blocker | Prioritize Critical/High only for release |

---

## 10. Dependency Map

| Dependency | Reason | Consuming Sprint/Module | Risk if Missing |
|---|---|---|---|
| Sprint 0 → All | Repo, CI, local dev, service skeleton | Sprint 1–10 | Development blocked |
| Sprint 1 → All Protected Modules | Auth, role, context, permission | Sprint 2–10 | Access control incomplete |
| Sprint 2 → PPDB | Conversion requires School Core student | Sprint 4 | Applicant cannot become student |
| Sprint 2 → Finance | Bills require student/school/class data | Sprint 5 | Billing target invalid |
| Sprint 2 → Academic | Schedule/attendance require student/teacher/class | Sprint 6 | Academic assignment invalid |
| Sprint 3 → PPDB | Applicant documents use file management | Sprint 4 | Document upload unavailable |
| Sprint 3 → Finance | Payment proof uses file management | Sprint 5 | Payment proof unavailable |
| Sprint 6 → Report Card | Report card requires subject/attendance/score foundation | Sprint 7 | Rapor cannot be generated |
| Sprint 4–8 → Reporting | Reporting needs events from operational services | Sprint 9 | Dashboard incomplete |
| Sprint 5/7 → Communication | Payment/report events trigger notification | Sprint 8 | Notifications incomplete |
| Sprint 0–9 → Sprint 10 | Hardening needs completed core modules | Sprint 10 | UAT/production readiness blocked |

---

## 11. AI Agent Usage Plan

### 11.1 Tasks AI Agent May Work On

AI Agent may work on:

```text
- scaffold code
- boilerplate handlers/services/repositories
- DTO/request/response structs
- migration drafts
- sqlc query drafts
- unit tests
- table-driven tests
- documentation drafts
- prompt refinement
- CI config drafts
- local setup docs
```

### 11.2 Tasks Requiring Human Review

Human review is mandatory for:

```text
- authentication/token logic
- authorization and scope logic
- object-level authorization
- finance calculations
- migration design
- event contract changes
- file privacy/signed URL logic
- audit log design
- report card publish/lock logic
- production deployment
```

### 11.3 Tasks AI Agent Must Not Do

AI Agent must not:

```text
- handle production secrets
- make final security approval
- make legal/compliance decisions
- approve production deployment
- access real sensitive data
- change architecture without explicit instruction
- merge PR without human review
```

### 11.4 How to Use AGENTS.md

AI Agent must read `AGENTS.md` first for:

```text
- project rules
- non-negotiable architecture constraints
- source of truth priority
- final response format
- stop conditions
```

### 11.5 How to Use SKILLS.md

AI Agent must use `SKILLS.md` to select the right workflow by task type:

```text
- Backend Go Service Skill
- Database Migration Skill
- API Gateway Skill
- Event/RabbitMQ Skill
- File Management Skill
- Finance Skill
- Testing Skill
- GitHub Project Management Skill
- Sprint Planning Skill
- Coding Prompt Usage Skill
```

### 11.6 How to Use Sprint Task Prompt

For implementation:

```text
1. Select one GitHub issue.
2. Read active sprint plan if available.
3. Open corresponding task prompt from docs/13 through docs/23.
4. Use only one task prompt.
5. Implement only selected scope.
6. Run tests.
7. Report changed files, tests, and risks.
```

### 11.7 Required AI Agent Final Output

For coding tasks, AI Agent final response must include:

```text
Summary
Changed files
Tests run
Notes/Risks
Next step if any
```

### 11.8 Quality Control for AI Agent Output

Human reviewer must check:

```text
- scope compliance
- architecture compliance
- service boundary
- no cross-service DB query
- permission/scope
- object-level auth
- audit/event/file/privacy
- test coverage
- no sensitive logs
```

---

## 12. Quality Gates

### 12.1 Quality Gate List

| Gate | Purpose |
|---|---|
| Lint | Ensure style and static checks |
| Formatting | Ensure consistent code formatting |
| Unit Test | Validate business logic |
| Integration Test | Validate service/repository/API integration |
| API Test | Validate external API behavior |
| Permission/Scope Test | Validate access control |
| Event Test | Validate event publish/consume/idempotency |
| Audit Test | Validate sensitive action logging |
| Build | Validate app/service build |
| Review | Human code review |
| QA Sign-off | Functional validation before release |

### 12.2 Quality Gate Per Branch

| Branch | Required Gate |
|---|---|
| feature/* | Local test relevant to change |
| develop | PR review, CI pass, unit/integration tests relevant |
| staging | CI pass, QA/UAT, regression for affected module |
| main | CI pass, QA sign-off, release notes, rollback plan, production approval |
| hotfix/* | Targeted test, CI pass, urgent review, post-fix back-merge |

---

## 13. Testing Strategy

### 13.1 Testing by Type

| Test Type | Scope | Owner |
|---|---|---|
| Unit Test | Services, domain logic, validators | Developer |
| Integration Test | Database, repository, service integration | Developer |
| API Test | REST endpoints and status codes | Developer/QA |
| Permission/Scope Test | RBAC, ABAC, object access | Developer/QA |
| Event Test | Publish/consume/idempotency | Developer |
| Audit Test | Sensitive action audit log | Developer/QA |
| Frontend Test | UI components, forms, API integration | Frontend Developer |
| Mobile Test | Widget/unit/API integration | Mobile Developer |
| Regression Test | Existing core flows | QA |
| UAT | User/business validation | QA/Product Owner |

### 13.2 Test Matrix Per Module

| Module | Unit | Integration | API | Scope | Event | Audit | Frontend/Mobile | UAT |
|---|---|---|---|---|---|---|---|---|
| Identity | Yes | Yes | Yes | Yes | Optional | Yes | Yes | Yes |
| School Core | Yes | Yes | Yes | Yes | Yes | Yes | Yes | Yes |
| File/Import | Yes | Yes | Yes | Yes | Optional | Yes | Yes | Yes |
| PPDB | Yes | Yes | Yes | Yes | Yes | Yes | Yes | Yes |
| Finance | Yes | Yes | Yes | Yes | Yes | Yes | Yes | Yes |
| Academic | Yes | Yes | Yes | Yes | Yes | Yes | Yes | Yes |
| Report Card | Yes | Yes | Yes | Yes | Yes | Yes | Yes | Yes |
| Communication | Yes | Yes | Yes | Yes | Yes | Optional | Yes | Yes |
| Reporting | Yes | Yes | Yes | Yes | Yes | Optional | Yes | Yes |
| Security/Hardening | N/A | Yes | Yes | Yes | Yes | Yes | N/A | Yes |

### 13.3 Sprint Testing Focus

| Sprint | Main Testing Focus |
|---|---|
| Sprint 0 | CI, local dev, Docker Compose, health checks |
| Sprint 1 | Auth, token rotation, permission baseline |
| Sprint 2 | CRUD, scope, object-level authorization |
| Sprint 3 | File privacy, import validation, signed URL |
| Sprint 4 | PPDB workflow, document, conversion idempotency |
| Sprint 5 | Bill idempotency, decimal amount, payment verification |
| Sprint 6 | Assignment-based access, attendance |
| Sprint 7 | Score input, publish/lock, revision approval |
| Sprint 8 | Event-driven notification, privacy payload |
| Sprint 9 | Projection idempotency, dashboard scope |
| Sprint 10 | Regression, UAT, security, backup/restore |

---

## 14. Security and Compliance Plan

### 14.1 Authentication

- JWT access token.
- Rotating refresh token.
- Refresh token stored hashed.
- Logout/revoke session.
- Token reuse detection.

### 14.2 Authorization

- RBAC for role/permission.
- ABAC/scope for foundation/school/class/subject/student/child.
- Service-side authorization mandatory.
- Frontend hiding is not sufficient.

### 14.3 Object-Level Authorization

Every resource by ID must be checked against actor scope.

Examples:

```text
- student belongs to actor school
- bill belongs to linked child
- report card belongs to linked child and is published
- file belongs to accessible entity
```

### 14.4 Field-Level Access

Restricted/Confidential fields may require masking or omission depending on role.

### 14.5 Audit

Audit log required for sensitive actions.

Audit must include:

```text
actor
action
resource
timestamp
request_id
correlation_id
metadata/safe payload
```

### 14.6 File Privacy

- Private by default.
- Signed URL only after backend authorization.
- Expiry depends on classification.
- Official files versioned.

### 14.7 Signed URL

Recommended expiry:

```text
Internal: ~30 minutes
Restricted: ~10 minutes
Confidential: ~3 minutes
```

### 14.8 Secrets

- No secrets in repository.
- `.env` not committed.
- GitHub Environment secrets per environment.
- Production secrets not handled by AI Agent.

### 14.9 Backup

- Daily backup for PostgreSQL DBs.
- Object storage backup.
- Encrypted backups.
- Restore test required.

### 14.10 Data Anak, Orang Tua, dan Keuangan

Classify as Restricted by default. Some data may be Confidential depending on sensitivity.

Rules:

```text
- no raw sensitive logs
- no Confidential data in notification body
- audit download/export
- enforce parent-child scope
```

---

## 15. Observability Plan

### 15.1 Structured JSON Logging

All services should use structured JSON logs.

Minimum fields:

```text
timestamp
level
service
environment
request_id
correlation_id
actor_id if available
route/operation
duration_ms
error if any
```

### 15.2 Request ID

Every external request should have `request_id`.

If missing, API Gateway generates it.

### 15.3 Correlation ID

Correlation ID should propagate across:

```text
API Gateway
service calls
gRPC calls
events
workers/consumers
```

### 15.4 Healthz

`/healthz` indicates service is alive.

### 15.5 Readyz

`/readyz` indicates service dependencies are ready.

### 15.6 Metrics

Minimum metrics:

```text
request count
request latency
error rate
database connection status
event publish count
event consume count
queue depth
DLQ count
```

### 15.7 Prometheus

Prometheus scrapes service metrics where available.

### 15.8 Grafana

Grafana dashboards for:

```text
service health
error rate
latency
RabbitMQ
PostgreSQL
Redis
MinIO
```

### 15.9 Loki

Loki stores centralized logs.

### 15.10 RabbitMQ Monitoring

Track:

```text
queue depth
consumer lag
retry count
DLQ count
failed events
```

### 15.11 DLQ Monitoring

DLQ must be visible before production readiness.

---

## 16. Backup and Restore Plan

### 16.1 Backup Scope

Backup includes:

```text
all PostgreSQL service DBs
object storage files
deployment config
encrypted secrets backup if applicable
RabbitMQ definitions if needed
```

### 16.2 Database Backup

Use daily `pg_dump` per service database:

```text
identity_db
school_core_db
admission_db
academic_db
finance_db
communication_db
reporting_db
```

### 16.3 Object Storage Backup

Sync object storage daily to backup destination.

### 16.4 Retention

Recommended retention:

```text
daily: 30 days
weekly: 12 weeks
monthly: 12 months
```

### 16.5 Encryption

Backups are Confidential and must be encrypted.

### 16.6 Restore Test

Restore test must be performed before MVP release.

Minimum:

```text
restore one database backup
restore object storage sample
verify app can read restored data
document result
```

### 16.7 RPO

Target RPO:

```text
<= 24 hours
```

### 16.8 RTO

Target RTO:

```text
4–8 hours
```

### 16.9 Backup Before Risky Migration

Backup must be taken before:

```text
large migration
mass import
mass billing generation
school year closing
production deployment with schema changes
```

---

## 17. Release Plan

### 17.1 Develop to Staging

Flow:

```text
develop → PR → staging
```

Requirements:

```text
CI pass
review complete
release candidate notes
no known blocker
```

### 17.2 Staging QA/UAT

Run:

```text
smoke test
regression test
module test
UAT scenarios
security checklist
backup/restore if relevant
```

### 17.3 Staging to Main

Flow:

```text
staging → PR → main
```

Requirements:

```text
QA sign-off
no Critical/High core bug
release notes
rollback plan
production approval ready
```

### 17.4 Production Approval

Production deploy requires GitHub Environment manual approval.

### 17.5 Release Tag

After production deploy:

```text
git tag v0.1.0
git push origin v0.1.0
```

### 17.6 Post-Release Verification

Verify:

```text
login
core APIs
dashboard
background workers
RabbitMQ queues
database connectivity
file access
logs/metrics
```

### 17.7 Rollback

Rollback options:

```text
redeploy previous image tag
disable feature via config if available
restore backup if data corruption occurs
run compensating migration if needed
```

---

## 18. Hotfix Plan

### 18.1 When to Use Hotfix

Use hotfix only for:

```text
production outage
security issue
data corruption risk
critical finance/auth bug
release blocker discovered after production
```

### 18.2 Hotfix Branch

Create from main:

```text
git checkout main
git pull
git checkout -b hotfix/fix-critical-issue
```

### 18.3 PR to Main

Hotfix PR must include:

```text
root cause
scope
tests
risk
rollback
```

### 18.4 Deploy Production

After approval, deploy to production.

### 18.5 Release Tag

Patch tag:

```text
v0.1.1
```

### 18.6 Back-Merge

After production deploy:

```text
main → staging
main → develop
```

---

## 19. Risk Management

| Risk | Impact | Probability | Mitigation | Owner |
|---|---|---|---|---|
| Scope creep | MVP terlambat | High | Lock MVP scope and non-goals | Product Owner |
| AI Agent inconsistent output | Kode tidak sesuai standar | Medium | Use AGENTS/SKILLS/task prompt and human review | Tech Lead |
| Service boundary violation | Arsitektur rusak | Medium | Code review and no cross-service DB rule | Backend Lead |
| Cross-service DB query | Coupling tinggi dan data risk | Medium | Enforce gRPC/events only | Backend Lead |
| Permission/scope leak | Data lintas sekolah bocor | High | Scope tests, object-level auth, security review | Backend/QA |
| Finance calculation bug | Tagihan salah | Medium | Decimal, snapshot, idempotency tests | Backend/QA |
| Data migration error | Data awal rusak | Medium | Validation-preview-confirm-report | Backend/QA |
| Reporting delay | Dashboard tidak update | Medium | Event projection and scheduled rebuild | Backend |
| Deployment failure | Staging/production down | Medium | CI, rollback, health checks | DevOps |
| Backup restore failure | Data loss risk | Low-Medium | Monthly restore test and release restore test | DevOps |
| Adoption issue | User sulit memakai sistem | Medium | UAT, simple UI labels, training | Product Owner |
| Privacy leak | Data anak/keuangan terekspos | Medium | Classification, masking, signed URL, audit | Security Reviewer |

---

## 20. MVP Readiness Criteria

MVP siap pilot/production jika:

```text
- login and refresh token flow works
- RBAC and ABAC/scope enforced
- object-level authorization tested
- School Core data can be managed
- Excel import works with validation-preview-confirm
- PPDB applicant to student conversion works
- Finance bill generation and manual payment verification work
- Academic schedule and attendance work
- Grade and report card publish flow works
- Parent/student can view published report card
- Communication/notification basic works
- Reporting dashboard works from reporting_db projection
- audit log exists for sensitive actions
- private file and signed URL work
- CI passes
- staging QA/UAT passes
- no Critical/High bug remains in core flow
- backup and restore test completed
- observability baseline available
- production rollback plan documented
```

---

## 21. Appendix

### 21.1 Recommended GitHub Labels

```text
type: feature
type: bug
type: chore
type: docs
type: refactor
type: test
type: security
type: infra
type: spike
type: hotfix

area: api-gateway
area: identity
area: school-core
area: admission
area: academic
area: finance
area: communication
area: reporting
area: web-admin
area: mobile
area: infra
area: docs
area: security
area: observability
area: ci-cd
area: file-management

sprint: 0
...
sprint: 10

priority: critical
priority: high
priority: medium
priority: low

status: ready
status: in-progress
status: blocked
status: needs-review
status: needs-qa
status: qa-passed
status: done

ai: ready
ai: needs-context
ai: generated
ai: needs-human-review
ai: do-not-use-agent

risk: low
risk: medium
risk: high
risk: breaking-change
risk: migration
risk: data-sensitive

review: backend
review: frontend
review: mobile
review: infra
review: qa
review: security
review: product
```

### 21.2 GitHub Project Fields

```text
Status
Sprint
Priority
Area
Type
Owner
Estimate
Risk
Platform
AI Agent
Target Release
```

### 21.3 Sprint Milestone List

```text
Sprint 0 — Project Foundation
Sprint 1 — Identity & Access
Sprint 2 — School Core
Sprint 3 — File Management + Import Excel
Sprint 4 — PPDB
Sprint 5 — Finance/SPP
Sprint 6 — Academic Basic
Sprint 7 — Report Card/E-Rapor Basic
Sprint 8 — Communication/Notification
Sprint 9 — Reporting Dashboard
Sprint 10 — Security, Observability, Backup, UAT Hardening
MVP Release
```

### 21.4 Useful Local Commands

```bash
# clone
git clone git@github.com:<org-or-user>/school-platform.git
cd school-platform

# branch
git checkout develop
git checkout -b feature/sprint-0-monorepo-structure

# docker compose
docker compose up -d
docker compose ps
docker compose logs -f

# go
go test ./...
go vet ./...
gofmt -w .

# web
npm install
npm run lint
npm run test
npm run build

# flutter
flutter pub get
flutter analyze
flutter test
```

### 21.5 Useful CI Commands

```bash
# validate compose
docker compose config

# check no env files
find . -name ".env" -o -name ".env.local"

# check markdown docs
find docs -name "*.md" | sort
```

### 21.6 AI Agent Prompt Usage Example

```text
You are working on project school-platform.

Read:
- AGENTS.md
- SKILLS.md
- docs/README.md
- docs/09-ai-agent-rules.md
- docs/08-coding-standard.md
- docs/29-sprint-0-plan.md if available
- docs/13-sprint-0-task-prompts.md

Task:
Implement Sprint 0 Task 0.1 — Create Monorepo Structure.

Rules:
- Implement only this task.
- Do not add business module code.
- Do not add production secrets.
- Update docs if needed.
- Add tests/checks if relevant.

Final response:
- Summary
- Changed files
- Tests run
- Risks/notes
```

## Sprint Plan and GitHub Setup References

Sprint execution must use the final sprint plan documents:

```text
docs/29-sprint-0-plan.md
docs/30-sprint-1-plan.md
docs/31-sprint-2-plan.md
docs/32-sprint-3-plan.md
docs/33-sprint-4-plan.md
docs/34-sprint-5-plan.md
docs/35-sprint-6-plan.md
docs/36-sprint-7-plan.md
docs/37-sprint-8-plan.md
docs/38-sprint-9-plan.md
docs/39-sprint-10-plan.md
```

Repository setup, labels, milestones, GitHub Project fields/views, branch protection, and GitHub Environments must use:

```text
docs/40-github-repository-setup-labels-project.md
```

## Git Commit and PR Naming Convention

Development execution must use the project commit and PR title convention:

```text
type(scope): short description
```

This convention applies to:

```text
- local commits
- PR titles
- squash merge commits
- release-related commits
- AI Agent-generated commit suggestions
```

References:

```text
docs/08-coding-standard.md
docs/27-workflow.md
docs/40-github-repository-setup-labels-project.md
```

## Repository Setup Automation

Sprint 0 and repository foundation work should use the GitHub setup guide and helper scripts:

```text
docs/40-github-repository-setup-labels-project.md
scripts/github/
```

Automation coverage:

```text
- repository creation
- repository support files
- branch creation
- branch protection
- environments
- labels
- milestones
- GitHub Project
- project fields
```

Manual review is still required for:

```text
- production environment reviewer approval
- GitHub Project views
- branch protection verification
- production secrets
```
