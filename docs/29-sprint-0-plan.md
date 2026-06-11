# Sprint 0 Plan — Project Foundation

Project: `school-platform`  
Sprint: Sprint 0 — Project Foundation  
Target Output: `docs/29-sprint-0-plan.md`  
Document Type: Sprint Planning Document  
Audience: Product Owner, Developer, QA, DevOps, AI Agent  
Status: Draft for Sprint Execution  

---

## 1. Sprint Summary

Sprint 0 adalah sprint fondasi untuk menyiapkan struktur awal project `school-platform` sebelum modul bisnis MVP mulai dikembangkan.

Fokus Sprint 0 adalah memastikan repository, local development environment, service skeleton, API Gateway skeleton, shared packages, contract folders, CI baseline, GitHub workflow, dan dokumentasi dasar sudah siap. Sprint ini tidak mengimplementasikan business module seperti Identity, School Core, PPDB, Finance, Academic, Report Card, Communication, atau Reporting.

Output Sprint 0 harus memungkinkan Sprint 1 — Identity & Access dimulai dengan fondasi teknis yang konsisten, mudah dijalankan secara lokal, dan siap dipakai oleh Developer, QA, DevOps, Product Owner, serta AI Agent.

Sprint 0 adalah checkpoint penting untuk memastikan aturan utama project sudah tertanam sejak awal:

```text
- monorepo structure jelas
- service boundary jelas
- local-first development berjalan
- Docker Compose tersedia
- CI baseline aktif
- .env.example tersedia
- tidak ada production secrets di repository
- healthz/readyz baseline tersedia
- request_id/correlation_id baseline tersedia
- structured logging baseline tersedia
- GitHub Issues/PR/Project workflow siap digunakan
```

---

## 2. Sprint Objective

Objective utama Sprint 0:

> Menyiapkan fondasi repository, local development environment, service template, API Gateway skeleton, shared packages, contract folders, Makefile, GitHub Actions basic CI, dan dokumentasi dasar agar development Sprint 1 dapat dimulai dengan disiplin teknis yang konsisten.

Tujuan praktis:

1. Developer dapat clone repository dan menjalankan local dependencies.
2. Struktur monorepo sudah disepakati dan tersedia.
3. API Gateway skeleton tersedia dengan health/readiness endpoint.
4. Template service Go tersedia untuk service MVP berikutnya.
5. Folder shared contract untuk proto, OpenAPI, dan event schema tersedia.
6. CI baseline dapat mendeteksi masalah repository, YAML, Go, Web, Mobile, dan Docker Compose secara kondisional.
7. GitHub workflow dasar tersedia: CODEOWNERS, PR template, Issue templates, CI workflow.
8. Dokumentasi onboarding dan local development tersedia.
9. AI Agent dapat membaca struktur dan aturan awal project sebelum coding.

---

## 3. Business Context

`school-platform` adalah platform internal yayasan sekolah multi-unit untuk TK, SD, SMP, dan SMA. Sebelum modul operasional seperti Identity, School Core, PPDB, Finance/SPP, Academic, Report Card, Communication, dan Reporting dibangun, project membutuhkan fondasi kerja yang kuat.

Tanpa Sprint 0, risiko yang muncul adalah:

```text
- setiap developer membuat struktur service berbeda
- local setup sulit direplikasi
- CI/CD tidak konsisten
- kontrak API/proto/event tersebar
- AI Agent tidak memiliki aturan kerja yang jelas
- service boundary berpotensi dilanggar sejak awal
- branch/PR/review workflow tidak tertib
- environment secret berisiko bocor ke repository
```

Sprint 0 memberikan nilai bisnis tidak langsung tetapi fundamental: mempercepat delivery sprint berikutnya, mengurangi rework teknis, mempermudah onboarding, dan menjaga kualitas implementasi sejak awal.

---

## 4. Technical Context

Project menggunakan arsitektur:

```text
Monorepo
Go microservices
Custom Go API Gateway
gRPC + protobuf untuk komunikasi internal
PostgreSQL database per service
RabbitMQ domain events
Redis
MinIO / S3-compatible object storage
Next.js web admin
Flutter mobile app
Docker Compose untuk local development dan staging awal
GitHub Actions untuk CI/CD
GitHub Projects untuk tracking
```

Sprint 0 harus menyiapkan struktur awal tanpa mengunci implementasi domain terlalu dini.

Batasan teknis utama:

```text
- API Gateway tidak boleh berisi business logic.
- Setiap service memiliki database sendiri.
- Tidak boleh ada cross-service database query.
- Komunikasi antar service menggunakan gRPC atau domain events.
- Reporting menggunakan read model/projection, bukan query DB operasional service lain.
- File private by default untuk modul yang nanti menggunakan storage.
- Internal code menggunakan English.
- UI labels nantinya menggunakan Bahasa Indonesia.
```

Service/app yang disiapkan secara skeleton:

```text
services/api-gateway
services/identity-service
services/school-core-service
services/admission-service
services/academic-service
services/finance-service
services/communication-service
services/reporting-service
apps/web-admin
apps/mobile-app
packages/proto
packages/openapi
packages/events
packages/shared-go
infra
deploy
docs
scripts
.github
```

---

## 5. Scope

Scope Sprint 0 mencakup:

```text
- monorepo structure
- Docker Compose local dependencies
- Go service template
- API Gateway skeleton
- shared-go skeleton
- proto/openapi/events placeholder
- Makefile
- GitHub Actions basic CI
- repository support files
- .env.example
- docs index/local development guide
- healthz/readyz baseline
- request_id/correlation_id baseline
- structured logging baseline
```

Detail scope:

1. Membuat struktur folder monorepo standar.
2. Membuat skeleton API Gateway dengan Go + Chi.
3. Membuat skeleton service template untuk microservice Go.
4. Membuat placeholder untuk service MVP.
5. Membuat folder `packages/proto`, `packages/openapi`, dan `packages/events`.
6. Membuat folder `packages/shared-go` untuk utility umum.
7. Membuat Docker Compose untuk dependency lokal.
8. Menyediakan PostgreSQL, Redis, RabbitMQ, dan MinIO di local.
9. Membuat `.env.example` root dan/atau per service.
10. Membuat Makefile command dasar.
11. Membuat GitHub Actions CI baseline.
12. Menambahkan CODEOWNERS, PR template, issue templates jika belum ada.
13. Menambahkan health check dan readiness check baseline.
14. Menambahkan request_id dan correlation_id baseline.
15. Menambahkan structured JSON logging baseline.
16. Memastikan dokumentasi local development tersedia.
17. Memastikan tidak ada `.env`, private key, atau production secret di repository.

---

## 6. Out of Scope

Hal berikut tidak boleh dikerjakan di Sprint 0:

```text
- full authentication
- JWT implementation final
- RBAC/ABAC final implementation
- business modules
- domain database schema final
- PPDB process
- Finance/SPP process
- Academic process
- Report Card process
- Communication/Notification process
- Reporting Dashboard implementation
- production deployment
- Kubernetes
- full UI feature implementation
- payment gateway
- WhatsApp integration
- offline mobile write
```

Sprint 0 boleh menyediakan placeholder atau skeleton, tetapi tidak boleh mulai mengimplementasikan business logic modul MVP.

---

## 7. Target Users / Actors

| Actor | Role in Sprint 0 | Need |
|---|---|---|
| Product Owner / Reviewer | Memastikan scope Sprint 0 sesuai kebutuhan MVP | Fondasi project siap untuk sprint berikutnya |
| Backend Developer | Membuat monorepo, API Gateway skeleton, service template, shared packages | Struktur backend konsisten |
| Frontend Developer | Menyiapkan placeholder web admin jika diperlukan | Struktur frontend siap diintegrasikan |
| Mobile Developer | Menyiapkan placeholder mobile app jika diperlukan | Struktur mobile siap diintegrasikan |
| QA | Memvalidasi local setup, CI, dan basic acceptance | Dasar test dan workflow jelas |
| Infrastructure/DevOps | Menyiapkan Docker Compose, CI, env separation, repository governance | Local/staging foundation siap |
| AI Agent | Membantu generate scaffold/docs/CI/checklist | Prompt dan output dapat direview manusia |

---

## 8. User Stories

### Product / Process Stories

- As a Product Owner, I want the repository and workflow to be ready, so that all MVP tasks can be tracked consistently.
- As a Technical Lead, I want a clear monorepo structure, so that each service and app has predictable ownership.
- As a DevOps Lead, I want Docker Compose and CI baseline, so that local development and PR validation are repeatable.
- As a QA, I want baseline health checks and CI checks, so that I can validate early system readiness.
- As an AI Agent operator, I want clear AGENTS/SKILLS/docs references, so that AI Agent can work safely on small scoped tasks.

### Developer Stories

- As a Backend Developer, I want a Go service template, so that new services can be created consistently.
- As a Backend Developer, I want API Gateway skeleton, so that Sprint 1 can add authentication routing cleanly.
- As a Frontend Developer, I want web app structure prepared, so that later UI modules have a standard location.
- As a Mobile Developer, I want mobile app structure prepared, so that mobile MVP can be developed without restructuring later.
- As a Developer, I want Makefile commands, so that common local commands are standardized.

### DevOps Stories

- As a DevOps engineer, I want environment examples and no committed secrets, so that environment separation is safe.
- As a DevOps engineer, I want CI baseline, so that PRs are checked automatically.
- As a DevOps engineer, I want Docker Compose dependency setup, so that local development is reproducible.

---

## 9. Functional Breakdown

Sprint 0 functional breakdown:

### 9.1 Repository Foundation

- Root repository structure.
- Standard app/service/package folders.
- Basic README/document index.
- `.gitignore`.
- `.env.example`.
- Makefile.

### 9.2 API Gateway Foundation

- API Gateway folder.
- Go module.
- Chi router skeleton.
- `/healthz`.
- `/readyz`.
- request_id middleware.
- correlation_id middleware.
- structured logging baseline.
- route grouping placeholder.

### 9.3 Go Service Template

- Base Go service layout.
- `cmd/server`.
- `internal/config`.
- `internal/http`.
- `internal/grpc` placeholder.
- `internal/repository` placeholder.
- `internal/service` placeholder.
- `internal/platform` placeholder.
- health/readiness baseline.
- structured logging baseline.

### 9.4 Shared Packages

- `packages/shared-go`.
- logging utility placeholder.
- request/correlation context utility placeholder.
- error response convention placeholder.
- common validation helper placeholder if needed.

### 9.5 Contract Folders

- `packages/proto`.
- `packages/openapi`.
- `packages/events`.
- README placeholder in each contract folder.
- Contract ownership rules documented.

### 9.6 Local Dependencies

- PostgreSQL.
- Redis.
- RabbitMQ.
- MinIO.
- Optional Mailpit.
- Optional observability placeholders.

### 9.7 GitHub Governance

- CODEOWNERS.
- Pull Request template.
- Issue templates.
- CI workflow.
- GitHub Project/labels/milestones documentation alignment.

### 9.8 Documentation

- Local development guide.
- Repository rules.
- AI Agent rules.
- Sprint 0 plan.
- README/doc index update.

---

## 10. Technical Breakdown

### 10.1 Backend

- Create service skeletons.
- Create API Gateway skeleton.
- Add Go module baseline.
- Add healthz/readyz.
- Add logging baseline.
- Add request_id/correlation_id middleware.
- Add configuration loading pattern.
- Add graceful shutdown pattern if included in template.
- Add base test examples.

### 10.2 API Gateway

- Use Go + Chi.
- Provide REST entrypoint skeleton.
- Add middleware chain placeholder.
- Add route group placeholder for `/api/v1`.
- Add health/readiness endpoint.
- Prepare auth middleware placeholder for Sprint 1.
- Do not add business logic.

### 10.3 Web Frontend

- Prepare `apps/web-admin` folder.
- If Next.js initialized in Sprint 0, keep it minimal.
- Add README or placeholder for web admin.
- No business UI implementation.

### 10.4 Mobile

- Prepare `apps/mobile-app` folder.
- If Flutter initialized in Sprint 0, keep it minimal.
- Add README or placeholder for mobile app.
- No business UI implementation.

### 10.5 QA

- Validate repository structure.
- Validate Docker Compose can start.
- Validate CI workflow syntax.
- Validate health/readiness endpoints if runnable.
- Validate no `.env` or private keys committed.
- Prepare Sprint 0 QA checklist.

### 10.6 DevOps

- Create Docker Compose file.
- Configure local service dependency ports.
- Add `.env.example`.
- Add GitHub Actions CI.
- Add CI checks for docs, YAML, Go, Web, Mobile, Docker Compose.
- Add branch/repository support files.

### 10.7 Documentation

- Update docs index.
- Add Sprint 0 plan.
- Ensure local development guide links to Docker Compose and Makefile commands.
- Ensure GitHub workflow docs align with current files.

---

## 11. Service and Data Ownership

| Component | Owner | Data Ownership | Notes |
|---|---|---|---|
| API Gateway | API Gateway service | No business data ownership | External REST entrypoint only |
| Identity Service | Identity service | Future identity_db | Sprint 0 skeleton only |
| School Core Service | School Core service | Future school_core_db | Sprint 0 skeleton only |
| Admission Service | Admission service | Future admission_db | Sprint 0 skeleton only |
| Academic Service | Academic service | Future academic_db | Sprint 0 skeleton only |
| Finance Service | Finance service | Future finance_db | Sprint 0 skeleton only |
| Communication Service | Communication service | Future communication_db | Sprint 0 skeleton only |
| Reporting Service | Reporting service | Future reporting_db | Sprint 0 skeleton only |
| Web Admin | Web app | No database ownership | Consumes API Gateway |
| Mobile App | Mobile app | No database ownership | Consumes API Gateway |
| packages/proto | Shared contract | gRPC contract ownership by service | No generated business logic in Sprint 0 |
| packages/openapi | Shared contract | REST contract documentation | API Gateway/external API |
| packages/events | Shared contract | Event schema documentation | RabbitMQ `domain.events` |
| packages/shared-go | Shared utilities | No domain ownership | Must avoid business logic |

Rules:

```text
- Shared packages must not contain domain business rules.
- Each service owns its future database.
- API Gateway owns routing and edge concerns only.
- Reporting Service owns read model/projection only.
```

---

## 12. API / gRPC / Event Impact

### 12.1 REST API Impact

Sprint 0 may introduce only infrastructure endpoints:

| Endpoint | Owner | Purpose |
|---|---|---|
| `GET /healthz` | API Gateway / service template | Liveness check |
| `GET /readyz` | API Gateway / service template | Readiness check |

No business REST API should be implemented in Sprint 0.

### 12.2 gRPC / Proto Impact

Sprint 0 only prepares folder structure:

```text
packages/proto/
```

Allowed:

```text
- README
- placeholder directory
- buf/proto tooling placeholder if planned
```

Not allowed:

```text
- final business proto definitions without sprint scope
```

### 12.3 Event Impact

Sprint 0 only prepares event schema folder:

```text
packages/events/
```

Allowed:

```text
- README
- event naming convention placeholder
- base metadata convention placeholder
```

Not allowed:

```text
- domain event implementation
- event consumer business logic
```

### 12.4 OpenAPI Impact

Sprint 0 may prepare:

```text
packages/openapi/
```

Allowed:

```text
- README
- placeholder OpenAPI root file if needed
```

Not allowed:

```text
- complete business API contract beyond health/readiness
```

### 12.5 Event Schema Impact

No final event schema required in Sprint 0. However, Sprint 0 should document future event metadata convention:

```text
event_id
event_name
event_version
occurred_at
producer
correlation_id
foundation_id if applicable
school_id if applicable
payload
```

---

## 13. Data Model Impact

Sprint 0 should not create final domain tables for MVP modules.

Allowed data/model impact:

| Area | Allowed in Sprint 0 |
|---|---|
| Database | Local PostgreSQL containers and empty DB placeholders |
| Migrations | Migration folder structure per service |
| SQLC | `sqlc.yaml` placeholder if service initialized |
| Seeds | No business seeds |
| Domain Tables | Not included |

Recommended future database names prepared in local config:

```text
identity_db
school_core_db
admission_db
academic_db
finance_db
communication_db
reporting_db
```

Sprint 0 may include placeholders for database URL naming convention but should not implement domain schema.

---

## 14. Permission and Scope Requirements

Sprint 0 does not implement RBAC/ABAC final logic, but must prepare conventions and guardrails.

Required principles to document and preserve:

```text
- authentication implemented in Sprint 1
- authorization enforced in backend services
- API Gateway validates token and routes request, but business authorization remains service-side
- object-level authorization mandatory for resource by ID
- foundation_id and school_id are required for main scoped data in later sprint
- no cross-service database query
```

Sprint 0 should prepare placeholders for:

```text
- actor context propagation
- request_id propagation
- correlation_id propagation
- auth middleware placeholder
```

Acceptance expectation:

```text
No Sprint 0 code should bypass or pre-implement insecure authorization shortcuts.
```

---

## 15. Audit Requirements

Sprint 0 does not implement full audit log, but must prepare audit-aware conventions.

Required decisions to document:

```text
- audit log implementation starts from sensitive actions in later sprints
- request_id and correlation_id must be available to audit events
- audit payload must not store raw Confidential data
- audit logs are separate from application logs
```

Sprint 0 tasks that may require audit awareness:

| Action | Audit Requirement |
|---|---|
| Repository setup | No runtime audit |
| CI workflow setup | No runtime audit |
| Logging baseline | Must avoid sensitive data |
| Request/correlation ID baseline | Must support later audit |
| GitHub workflow setup | Tracked through Git history/PR |

---

## 16. File and Privacy Requirements

Sprint 0 may set up MinIO/S3-compatible local dependency but must not implement final file management.

Required privacy baseline:

```text
- file private by default in future modules
- no public bucket by default unless explicitly configured for safe public assets
- local MinIO credentials must be development-only
- .env and secrets must not be committed
- signed URL implementation deferred to Sprint 3
```

File classification principles to document:

```text
Public
Internal
Restricted
Confidential
```

Sprint 0 CI must help prevent:

```text
- committed .env files
- committed private keys
- committed certificates/secrets
```

---

## 17. Test Plan

### 17.1 Unit Tests

| Target | Test |
|---|---|
| Go service template | Basic handler/service test if implemented |
| API Gateway health handler | Returns expected status |
| Request ID middleware | Adds/propagates request_id |
| Correlation ID middleware | Adds/propagates correlation_id |
| Config loader | Loads safe local config if implemented |

### 17.2 Integration Tests

| Target | Test |
|---|---|
| Docker Compose | Config validates |
| API Gateway | Starts locally and responds to healthz/readyz if runnable |
| Service template | Starts with local env if runnable |

### 17.3 API Tests

| Endpoint | Expected |
|---|---|
| `GET /healthz` | 200 OK |
| `GET /readyz` | 200 OK or dependency-aware response |

### 17.4 Permission/Scope Tests

Not required for Sprint 0 final logic. However, ensure no code introduces unsafe shortcuts.

### 17.5 Event Tests

Not required for domain events. Optional validation for event schema placeholder docs.

### 17.6 Audit Tests

Not required for runtime audit. Ensure logging tests do not include sensitive data if applicable.

### 17.7 Frontend Tests

If Next.js app is initialized:

```text
- lint pass
- build pass
- smoke test placeholder if configured
```

### 17.8 Mobile Tests

If Flutter app is initialized:

```text
- flutter analyze pass
- flutter test pass if baseline test exists
```

### 17.9 UAT Scenarios

Sprint 0 UAT is technical readiness, not business UAT.

Checklist:

```text
- Developer can clone repository.
- Developer can run Docker Compose.
- Developer can run CI-equivalent local commands.
- API Gateway health endpoint works if implemented.
- Documentation explains how to start local development.
- GitHub PR/Issue workflow files exist.
```

---

## 18. Acceptance Criteria

Sprint 0 acceptance criteria:

```text
- [ ] Monorepo root structure exists.
- [ ] `services/` folder exists with MVP service placeholders or initialized skeletons.
- [ ] `services/api-gateway/` exists with Go + Chi skeleton.
- [ ] API Gateway has `/healthz` endpoint.
- [ ] API Gateway has `/readyz` endpoint.
- [ ] Base request_id middleware or convention exists.
- [ ] Base correlation_id middleware or convention exists.
- [ ] Structured JSON logging baseline exists.
- [ ] `packages/proto/` exists.
- [ ] `packages/openapi/` exists.
- [ ] `packages/events/` exists.
- [ ] `packages/shared-go/` exists.
- [ ] `apps/web-admin/` exists or has explicit placeholder.
- [ ] `apps/mobile-app/` exists or has explicit placeholder.
- [ ] Docker Compose file exists and validates.
- [ ] Docker Compose includes PostgreSQL, Redis, RabbitMQ, and MinIO.
- [ ] `.env.example` exists.
- [ ] `.env` is not committed.
- [ ] Private keys are not committed.
- [ ] Makefile exists with common commands.
- [ ] GitHub Actions CI exists.
- [ ] CI validates repository checks and YAML.
- [ ] CI conditionally checks Go/Web/Mobile when files exist.
- [ ] CODEOWNERS exists.
- [ ] Pull Request template exists.
- [ ] Issue templates exist.
- [ ] Local development documentation exists or is updated.
- [ ] Sprint 0 plan is committed as `docs/29-sprint-0-plan.md`.
```

---

## 19. Definition of Ready

Sprint 0 task/issue is ready when:

```text
- Objective is clear.
- Scope is small and specific.
- Out of scope is explicit.
- Acceptance criteria are checklist-based.
- Target folder/file is identified.
- Owner is assigned.
- Labels are assigned.
- Milestone is Sprint 0.
- Project field Status is Ready.
- AI Agent status is set.
- Required docs are listed.
```

Checklist:

```text
- [ ] Issue uses `feature_task.yml`, `ai_agent_task.yml`, or relevant template.
- [ ] Label `sprint: 0` assigned.
- [ ] Area label assigned.
- [ ] Priority assigned.
- [ ] Risk assigned.
- [ ] Milestone `Sprint 0 — Project Foundation` assigned.
- [ ] GitHub Project fields filled.
- [ ] Acceptance criteria defined.
```

---

## 20. Definition of Done

Sprint 0 task is done when:

```text
- implementation matches issue scope
- no out-of-scope business logic added
- relevant files created/updated
- local validation done
- tests/checks pass
- CI pass
- docs updated if relevant
- PR approved
- issue moved to Done
```

Sprint 0 is done when:

```text
- all Sprint 0 blocking issues are Done
- local setup can be run by developer
- CI baseline passes
- repository support files exist
- API Gateway skeleton is ready for Sprint 1
- documentation is updated
- handoff notes to Sprint 1 are written
```

Checklist:

```text
- [ ] All Sprint 0 Critical/High issues closed.
- [ ] CI pass on develop.
- [ ] Local setup verified.
- [ ] API Gateway skeleton verified.
- [ ] No secrets committed.
- [ ] Sprint 1 can start.
```

---

## 21. Dependencies

| Dependency | Required For | Notes |
|---|---|---|
| GitHub repository | All Sprint 0 work | Repository must exist before branch/PR workflow |
| Branches develop/staging/main | Workflow | Required for protected branch setup |
| Docker installed locally | Local dev | Developer machine prerequisite |
| Go toolchain | API Gateway/service skeleton | Required for backend checks |
| Node.js | Web admin if initialized | Required only if app exists |
| Flutter SDK | Mobile app if initialized | Required only if app exists |
| GitHub Actions | CI | Required for PR validation |
| Docker Compose | Local dependency orchestration | Required for PostgreSQL/Redis/RabbitMQ/MinIO |

Sprint dependencies:

| Sprint | Depends on Sprint 0 because |
|---|---|
| Sprint 1 | Needs API Gateway skeleton, service template, CI, local env |
| Sprint 2 | Needs service template, DB/migration conventions |
| Sprint 3 | Needs MinIO/local file dependency baseline |
| Sprint 4–10 | Need branch, CI, issue/PR workflow discipline |

---

## 22. Risks and Mitigations

| Risk | Impact | Probability | Mitigation | Owner |
|---|---|---|---|---|
| Monorepo structure tidak konsisten | Rework besar di sprint berikutnya | Medium | Review struktur terhadap architecture docs sebelum merge | Backend Lead |
| CI terlalu ketat saat project belum lengkap | PR awal sering gagal tidak perlu | Medium | Gunakan conditional checks untuk Go/Web/Mobile | DevOps |
| CI terlalu longgar | Bug dasar lolos ke develop | Medium | Minimal repository/YAML/secret/Docker checks wajib | DevOps |
| Secret atau `.env` tercommit | Risiko keamanan | Medium | `.gitignore`, CI secret check, PR checklist | DevOps |
| API Gateway mulai berisi business logic | Architecture violation | Low-Medium | Review PR, dokumentasikan gateway responsibility | Backend Lead |
| Shared package berisi domain logic | Coupling antar service | Medium | Shared-go hanya utility non-domain | Backend Lead |
| Docker Compose sulit dijalankan | Developer onboarding lambat | Medium | Local development guide dan Makefile | DevOps |
| AI Agent membuat scaffold terlalu luas | Scope creep | Medium | Issue kecil, prompt jelas, human review | Tech Lead |
| Placeholder dianggap final design | Kebingungan sprint berikutnya | Low | Dokumentasikan placeholder vs final implementation | Tech Lead |
| Service skeleton tidak siap untuk Sprint 1 | Sprint 1 tertunda | Medium | Exit criteria wajib sebelum Sprint 0 ditutup | Backend Lead |

---

## 23. AI Agent Usage Guidance

### 23.1 Task yang Cocok untuk AI Agent

AI Agent cocok membantu:

```text
- membuat struktur folder awal
- membuat README placeholder
- membuat Makefile draft
- membuat Docker Compose draft
- membuat GitHub Actions draft
- membuat service skeleton
- membuat API Gateway health endpoint skeleton
- membuat request_id/correlation_id middleware draft
- membuat structured logging utility draft
- membuat issue/task breakdown
- membuat dokumentasi Sprint 0
```

### 23.2 Task yang Butuh Human Review

Wajib human review untuk:

```text
- repository structure final
- Docker Compose final
- CI workflow final
- API Gateway middleware pattern
- logging fields
- environment variable naming
- branch protection/GitHub workflow
- security checks
```

### 23.3 Task yang Tidak Boleh Dikerjakan AI Agent

AI Agent tidak boleh:

```text
- membuat atau mengakses production secrets
- membuat final production deployment approval
- mengubah arsitektur besar tanpa instruksi
- menambahkan business logic modul MVP di Sprint 0
- membuat credential asli
- menggunakan data production
- memutuskan final security/compliance approval
```

### 23.4 Dokumen Prompt Coding yang Harus Dipakai

Untuk Sprint 0 coding task, AI Agent harus membaca:

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
docs/13-sprint-0-task-prompts.md
docs/24-local-development-guide.md
docs/29-sprint-0-plan.md
```

### 23.5 Required AI Agent Output

AI Agent final response untuk task Sprint 0 harus berisi:

```text
Summary
Changed files
Tests/checks run
Assumptions
Risks/notes
Next steps if any
```

---

## 24. GitHub Issue Plan

| Issue Title | Type | Area | Priority | Estimate | Labels | Milestone |
|---|---|---|---|---|---|---|
| Sprint 0 Task 0.1 — Create Monorepo Structure | feature | infra | critical | 3 | `type: feature`, `area: infra`, `sprint: 0`, `priority: critical`, `status: ready`, `ai: ready`, `risk: medium`, `review: backend` | Sprint 0 — Project Foundation |
| Sprint 0 Task 0.2 — Add Root README and Docs Index | docs | docs | high | 2 | `type: docs`, `area: docs`, `sprint: 0`, `priority: high`, `status: ready`, `ai: ready`, `risk: low`, `review: product` | Sprint 0 — Project Foundation |
| Sprint 0 Task 0.3 — Add Docker Compose Local Dependencies | infra | infra | critical | 5 | `type: infra`, `area: infra`, `sprint: 0`, `priority: critical`, `status: ready`, `ai: ready`, `risk: medium`, `review: infra` | Sprint 0 — Project Foundation |
| Sprint 0 Task 0.4 — Add Environment Examples and Gitignore Rules | chore | infra | high | 2 | `type: chore`, `area: infra`, `sprint: 0`, `priority: high`, `status: ready`, `ai: ready`, `risk: data-sensitive`, `review: security` | Sprint 0 — Project Foundation |
| Sprint 0 Task 0.5 — Create API Gateway Skeleton | feature | api-gateway | critical | 5 | `type: feature`, `area: api-gateway`, `sprint: 0`, `priority: critical`, `status: ready`, `ai: ready`, `risk: medium`, `review: backend` | Sprint 0 — Project Foundation |
| Sprint 0 Task 0.6 — Create Go Service Template | feature | infra | critical | 5 | `type: feature`, `area: infra`, `sprint: 0`, `priority: critical`, `status: ready`, `ai: ready`, `risk: medium`, `review: backend` | Sprint 0 — Project Foundation |
| Sprint 0 Task 0.7 — Add Shared Contract Folders | feature | infra | high | 2 | `type: feature`, `area: infra`, `sprint: 0`, `priority: high`, `status: ready`, `ai: ready`, `risk: low`, `review: backend` | Sprint 0 — Project Foundation |
| Sprint 0 Task 0.8 — Add Shared Go Package Skeleton | feature | infra | high | 3 | `type: feature`, `area: infra`, `sprint: 0`, `priority: high`, `status: ready`, `ai: ready`, `risk: medium`, `review: backend` | Sprint 0 — Project Foundation |
| Sprint 0 Task 0.9 — Add Request ID and Correlation ID Baseline | feature | api-gateway | high | 3 | `type: feature`, `area: api-gateway`, `sprint: 0`, `priority: high`, `status: ready`, `ai: ready`, `risk: medium`, `review: backend` | Sprint 0 — Project Foundation |
| Sprint 0 Task 0.10 — Add Structured JSON Logging Baseline | feature | observability | high | 3 | `type: feature`, `area: observability`, `sprint: 0`, `priority: high`, `status: ready`, `ai: ready`, `risk: medium`, `review: backend` | Sprint 0 — Project Foundation |
| Sprint 0 Task 0.11 — Add Makefile Common Commands | chore | infra | medium | 2 | `type: chore`, `area: infra`, `sprint: 0`, `priority: medium`, `status: ready`, `ai: ready`, `risk: low`, `review: infra` | Sprint 0 — Project Foundation |
| Sprint 0 Task 0.12 — Add GitHub Actions CI Baseline | infra | ci-cd | critical | 5 | `type: infra`, `area: ci-cd`, `sprint: 0`, `priority: critical`, `status: ready`, `ai: ready`, `risk: medium`, `review: infra` | Sprint 0 — Project Foundation |
| Sprint 0 Task 0.13 — Add GitHub Repository Support Files | chore | ci-cd | high | 3 | `type: chore`, `area: ci-cd`, `sprint: 0`, `priority: high`, `status: ready`, `ai: ready`, `risk: low`, `review: infra` | Sprint 0 — Project Foundation |
| Sprint 0 Task 0.14 — Add Local Development Guide Validation | docs | docs | high | 3 | `type: docs`, `area: docs`, `sprint: 0`, `priority: high`, `status: ready`, `ai: ready`, `risk: low`, `review: qa` | Sprint 0 — Project Foundation |
| Sprint 0 Task 0.15 — QA Validate Sprint 0 Foundation | test | qa | high | 3 | `type: test`, `area: ci-cd`, `sprint: 0`, `priority: high`, `status: ready`, `ai: needs-human-review`, `risk: medium`, `review: qa` | Sprint 0 — Project Foundation |

---

## 25. GitHub Project Fields

Recommended GitHub Project fields for Sprint 0 issues:

| Field | Recommended Value / Options |
|---|---|
| Status | Backlog, Ready, In Progress, In Review, QA, Blocked, Done |
| Sprint | Sprint 0 |
| Priority | Critical, High, Medium, Low |
| Area | infra, api-gateway, docs, ci-cd, observability |
| Type | feature, chore, docs, infra, test |
| Owner | Backend Developer, DevOps, QA, Product Owner, AI Agent operator |
| Estimate | 1, 2, 3, 5 |
| Risk | Low, Medium, Data Sensitive |
| Platform | Backend, Infra, Docs, QA |
| AI Agent | Ready, Needs Human Review, Needs Context |
| Target Release | MVP |

Example field values for Task 0.5:

```text
Status: Ready
Sprint: Sprint 0
Priority: Critical
Area: api-gateway
Type: feature
Owner: Backend Developer
Estimate: 5
Risk: Medium
Platform: Backend
AI Agent: Ready
Target Release: MVP
```

---

## 26. Sprint Exit Criteria

Sprint 0 boleh ditutup jika:

```text
- All Critical Sprint 0 issues are Done.
- API Gateway skeleton is ready.
- Go service template is available.
- Docker Compose local dependencies are available and validated.
- CI baseline passes on develop.
- Repository support files are available.
- Local development documentation is available.
- No .env or secret files are committed.
- Request ID/correlation ID baseline exists.
- Structured logging baseline exists.
- Contract folders exist.
- Sprint 1 can start without repository restructuring.
```

Exit checklist:

```text
- [ ] Sprint 0 milestone has no open Critical/High blocking issue.
- [ ] CI pass.
- [ ] Local setup verified by at least one developer other than implementer.
- [ ] QA validates local startup and repository checks.
- [ ] DevOps validates Docker Compose and CI.
- [ ] Product/Reviewer accepts Sprint 0 deliverables.
- [ ] Handoff notes for Sprint 1 are documented.
```

---

## 27. Handoff Notes

Handoff to Sprint 1 — Identity & Access:

### 27.1 What Sprint 1 Can Use

Sprint 1 should be able to use:

```text
- services/api-gateway skeleton
- Go service template
- identity-service skeleton
- packages/shared-go utilities
- request_id/correlation_id baseline
- structured logging baseline
- Docker Compose PostgreSQL/Redis/RabbitMQ/MinIO
- CI baseline
- Makefile commands
- GitHub issue/PR workflow
```

### 27.2 What Sprint 1 Must Build

Sprint 1 must build:

```text
- identity_db migrations
- users table
- password hashing
- login endpoint
- JWT access token
- rotating refresh token
- logout/revoke session
- roles and permissions
- actor context
- API Gateway auth middleware
```

### 27.3 Known Constraints for Sprint 1

```text
- API Gateway must not contain business authorization logic.
- Identity service owns identity data.
- Token/password must never be logged.
- Refresh token must be stored hashed.
- RBAC + ABAC/scope foundation starts in Sprint 1.
```

### 27.4 Recommended First Sprint 1 Issue

```text
Sprint 1 Task 1.1 — Create Identity Service Database Migrations
```

Suggested labels:

```text
type: feature
area: identity
sprint: 1
priority: critical
risk: data-sensitive
review: backend
review: security
ai: ready
```
