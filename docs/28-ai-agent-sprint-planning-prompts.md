# 28 — AI Agent Sprint Planning Prompts

Project: `school-platform`  
Purpose: Prompt pack for AI Agent to generate sprint planning documents  
Target output examples: `docs/29-sprint-0-plan.md` through `docs/39-sprint-10-plan.md`

---

## 1. Purpose

This document contains AI Agent prompts for creating **Sprint Planning Documents** for Sprint 0 through Sprint 10.

These prompts are different from task implementation prompts.

Existing task prompt documents:

```text
13-sprint-0-task-prompts.md
14-sprint-1-task-prompts.md
15-sprint-2-task-prompts.md
...
23-sprint-10-task-prompts.md
```

Those documents are used when AI Agent starts writing code.

This document is used earlier, when AI Agent must generate a structured sprint plan containing:

```text
- sprint objective
- business context
- technical context
- scope
- out of scope
- user stories
- task breakdown
- backend/frontend/mobile/QA/DevOps tasks
- dependencies
- API/proto/event impact
- data model impact
- permission/scope requirements
- audit requirements
- testing plan
- acceptance criteria
- Definition of Ready
- Definition of Done
- risks
- GitHub issue list
- labels/milestone/project fields
```

---

## 2. How to Use This Document

Use this document when you want AI Agent to create a sprint-level planning document.

Recommended output files:

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

Recommended flow:

```text
Sprint Planning Prompt
→ Sprint Plan Document
→ GitHub Issues
→ AI Agent Task Prompt
→ Implementation PR
```

---

## 3. Global Prompt Header

Use this header before each sprint-specific prompt.

```text
Kamu adalah Senior Technical Project Manager, Product-Oriented Engineering Lead, Software Architect, QA Planner, dan AI Agent Workflow Designer.

Tugasmu adalah membuat dokumen Sprint Plan untuk project `school-platform`.

Dokumen Sprint Plan ini harus praktis, detail, dan bisa langsung digunakan untuk:
- membuat GitHub Issues
- mengisi GitHub Project
- mengarahkan Developer
- mengarahkan QA
- mengarahkan DevOps
- memberi konteks kepada AI Agent sebelum coding

Gunakan Bahasa Indonesia.

Jangan menulis kode implementasi.
Jangan mengubah keputusan arsitektur yang sudah ada.
Jangan menambahkan fitur di luar scope MVP.
```

---

## 4. Source of Truth for All Sprint Planning Prompts

Always include this section in each AI Agent prompt.

```text
Gunakan dokumen berikut sebagai source of truth:

- AGENTS.md
- SKILLS.md
- docs/01-technical-architecture.md
- docs/02-service-boundary.md
- docs/03-data-model-mvp.md
- docs/04-api-contract.md
- docs/05-event-contract.md
- docs/06-ui-screen-user-flow.md
- docs/07-test-plan-acceptance-criteria.md
- docs/08-coding-standard.md
- docs/09-ai-agent-rules.md
- docs/10-sprint-backlog-mvp.md
- docs/11-github-repository-rules.md
- docs/12-ai-agent-sprint-prompts.md
- docs/13-sprint-0-task-prompts.md
- docs/14-sprint-1-task-prompts.md
- docs/15-sprint-2-task-prompts.md
- docs/16-sprint-3-task-prompts.md
- docs/17-sprint-4-task-prompts.md
- docs/18-sprint-5-task-prompts.md
- docs/19-sprint-6-task-prompts.md
- docs/20-sprint-7-task-prompts.md
- docs/21-sprint-8-task-prompts.md
- docs/22-sprint-9-task-prompts.md
- docs/23-sprint-10-task-prompts.md
- docs/24-local-development-guide.md
- docs/25-prd-prompt.md
- docs/26-development-plan-prompt.md
- docs/27-workflow-prompt.md

Jika ada konflik antar dokumen, prioritaskan:
1. AGENTS.md
2. SKILLS.md
3. docs/01-technical-architecture.md
4. docs/02-service-boundary.md
5. docs/03-data-model-mvp.md
6. docs/04-api-contract.md
7. docs/05-event-contract.md
8. docs/07-test-plan-acceptance-criteria.md
9. docs/08-coding-standard.md
10. docs/09-ai-agent-rules.md
11. docs/10-sprint-backlog-mvp.md
12. sprint task prompt terkait
```

---

## 5. Standard Sprint Plan Output Format

Every sprint planning prompt must ask AI Agent to use this structure.

```text
# Sprint N Plan — [Sprint Name]

## 1. Sprint Summary
Jelaskan ringkasan sprint.

## 2. Sprint Objective
Tuliskan objective utama sprint.

## 3. Business Context
Jelaskan alasan sprint ini penting dari sisi operasional yayasan/sekolah.

## 4. Technical Context
Jelaskan konteks teknis sprint, service/app yang terlibat, dan batasan arsitektur.

## 5. Scope
Tuliskan fitur/pekerjaan yang masuk sprint.

## 6. Out of Scope
Tuliskan hal-hal yang tidak boleh dikerjakan dalam sprint ini.

## 7. Target Users / Actors
Tuliskan role yang terlibat.

## 8. User Stories
Buat user stories dengan format:
- As a [role], I want [capability], so that [benefit].

## 9. Functional Breakdown
Pecah fitur menjadi bagian fungsional.

## 10. Technical Breakdown
Pecah pekerjaan teknis:
- Backend
- API Gateway
- Web Frontend
- Mobile
- QA
- DevOps
- Documentation

## 11. Service and Data Ownership
Jelaskan service owner dan data owner.

## 12. API / gRPC / Event Impact
Jelaskan:
- REST API impact
- gRPC/proto impact
- event impact
- OpenAPI impact
- event schema impact

## 13. Data Model Impact
Jelaskan tabel/entity yang mungkin dibutuhkan.

## 14. Permission and Scope Requirements
Jelaskan permission, RBAC, ABAC/scope, dan object-level authorization.

## 15. Audit Requirements
Tuliskan action yang perlu audit log.

## 16. File and Privacy Requirements
Tuliskan kebutuhan file, klasifikasi data, signed URL, dan privacy jika relevan.

## 17. Test Plan
Tuliskan:
- unit tests
- integration tests
- API tests
- permission/scope tests
- event tests
- audit tests
- frontend tests
- mobile tests
- UAT scenarios

## 18. Acceptance Criteria
Buat checklist acceptance criteria.

## 19. Definition of Ready
Tuliskan syarat task/issue siap dikerjakan.

## 20. Definition of Done
Tuliskan syarat task/sprint dianggap selesai.

## 21. Dependencies
Tuliskan dependency ke sprint/modul/service lain.

## 22. Risks and Mitigations
Buat tabel:
| Risk | Impact | Probability | Mitigation | Owner |
|---|---|---|---|---|

## 23. AI Agent Usage Guidance
Tuliskan:
- task yang cocok untuk AI Agent
- task yang butuh human review
- task yang tidak boleh dikerjakan AI Agent
- dokumen prompt coding yang harus dipakai

## 24. GitHub Issue Plan
Buat daftar issue yang disarankan.

Format:
| Issue Title | Type | Area | Priority | Estimate | Labels | Milestone |
|---|---|---|---|---|---|---|

## 25. GitHub Project Fields
Tuliskan field yang direkomendasikan untuk setiap issue:
- Status
- Sprint
- Priority
- Area
- Type
- Owner
- Estimate
- Risk
- Platform
- AI Agent
- Target Release

## 26. Sprint Exit Criteria
Tuliskan kriteria sprint boleh ditutup.

## 27. Handoff Notes
Tuliskan catatan handoff ke sprint berikutnya.
```

---

## 6. Sprint 0 Planning Prompt — Project Foundation

```text
[Use Global Prompt Header]

Buat dokumen Sprint Plan untuk:

Sprint 0 — Project Foundation

Target output:
docs/29-sprint-0-plan.md

Sprint objective:
Menyiapkan fondasi repository, local development environment, service template, API Gateway skeleton, shared packages, contract folders, Makefile, GitHub Actions basic CI, dan dokumentasi dasar agar development Sprint 1 bisa dimulai dengan disiplin.

Main scope:
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

Out of scope:
- full authentication
- business modules
- production deployment
- Kubernetes
- domain database schema
- UI feature implementation

Important docs:
- docs/13-sprint-0-task-prompts.md
- docs/24-local-development-guide.md
- docs/11-github-repository-rules.md
- AGENTS.md
- SKILLS.md

Use the Standard Sprint Plan Output Format.
```

---

## 7. Sprint 1 Planning Prompt — Identity & Access

```text
[Use Global Prompt Header]

Buat dokumen Sprint Plan untuk:

Sprint 1 — Identity & Access

Target output:
docs/30-sprint-1-plan.md

Sprint objective:
Membangun fondasi autentikasi dan otorisasi, termasuk users, password hashing, login, JWT access token, rotating refresh token, logout/session revocation, roles, permissions, role assignment, user context, dan API Gateway auth middleware.

Main scope:
- identity_db migrations
- users
- roles
- permissions
- role_permissions
- user_role_assignments
- user_sessions
- password hashing
- login
- refresh token rotation
- logout/revoke session
- role/permission seed
- actor context
- API Gateway auth middleware
- `/api/v1/auth/login`
- `/api/v1/auth/refresh`
- `/api/v1/auth/logout`
- `/api/v1/me`
- `/api/v1/me/permissions`
- `/api/v1/me/context`

Out of scope:
- OAuth/social login
- 2FA
- advanced anomaly detection
- full user profile management
- UI lengkap role management

Important security requirements:
- refresh token stored hashed
- access token short-lived
- no token/password in logs
- RBAC + ABAC/scope foundation
- service-side authorization remains mandatory

Important docs:
- docs/14-sprint-1-task-prompts.md
- docs/04-api-contract.md
- docs/08-coding-standard.md
- docs/09-ai-agent-rules.md

Use the Standard Sprint Plan Output Format.
```

---

## 8. Sprint 2 Planning Prompt — School Core

```text
[Use Global Prompt Header]

Buat dokumen Sprint Plan untuk:

Sprint 2 — School Core

Target output:
docs/31-sprint-2-plan.md

Sprint objective:
Membangun master data inti yayasan/sekolah: foundation, school, academic year, semester, students, guardians, teachers, grade levels, classes, student-class assignment, teacher assignment, dan homeroom assignment.

Main scope:
- school_core_db migrations
- foundations
- schools
- academic_years
- semesters
- students
- guardians
- student_guardians
- teachers
- grade_levels
- classes
- student_class_assignments
- teacher_assignments
- homeroom_assignments
- CRUD core data
- search/filter student and teacher
- school/foundation scope checks
- school core events
- audit for sensitive changes

Out of scope:
- PPDB process
- Excel import
- finance
- academic grade/report card
- payroll/HR lengkap

Important rules:
- School Core owns student/guardian/teacher/class master data
- no direct role credential creation here
- use reference user_id only when needed
- no cross-service DB query
- free SPP is not student status

Important docs:
- docs/15-sprint-2-task-prompts.md
- docs/03-data-model-mvp.md
- docs/02-service-boundary.md

Use the Standard Sprint Plan Output Format.
```

---

## 9. Sprint 3 Planning Prompt — File Management + Import Excel

```text
[Use Global Prompt Header]

Buat dokumen Sprint Plan untuk:

Sprint 3 — File Management + Import Excel

Target output:
docs/32-sprint-3-plan.md

Sprint objective:
Membangun fondasi private file management dan import Excel data awal untuk students, guardians, teachers, classes, dan assignments.

Main scope:
- file metadata structure
- MinIO/S3 storage abstraction
- private file upload
- signed URL authorization
- MIME/extension/size validation
- classification: public/internal/restricted/confidential
- import_batches
- import_batch_rows
- template download
- Excel upload
- validation preview
- confirm import
- import report
- error report

Supported import MVP:
- students
- guardians
- teachers
- classes
- student-class assignment via class_code
- optional homeroom/teacher assignment if data ready

Out of scope:
- import grades
- historical payments import
- payroll import
- asset/library/BK/UKS/alumni/koperasi import
- virus scanning integration
- central File Service

Important rules:
- file private by default
- import file is Restricted
- no insert before validation and preview
- no raw import data in logs
- every import has import_batch_id

Important docs:
- docs/16-sprint-3-task-prompts.md
- docs/03-data-model-mvp.md
- docs/24-local-development-guide.md

Use the Standard Sprint Plan Output Format.
```

---

## 10. Sprint 4 Planning Prompt — PPDB

```text
[Use Global Prompt Header]

Buat dokumen Sprint Plan untuk:

Sprint 4 — PPDB

Target output:
docs/33-sprint-4-plan.md

Sprint objective:
Membangun proses PPDB dari periode pendaftaran, applicant submission, upload dokumen, verifikasi dokumen, keputusan accept/reject, sampai konversi applicant menjadi student di School Core.

Main scope:
- admission_db migrations
- admission_periods
- applicants
- applicant_guardians
- applicant_documents
- applicant_verifications
- admission_decisions
- admission period CRUD
- applicant submission
- document upload
- document verification
- accept/reject decision
- conversion to student through School Core gRPC
- converted_student_id tracking
- idempotency for conversion
- PPDB events
- audit for decisions

Out of scope:
- complex scoring/selection
- public marketing website
- payment gateway
- advanced admission analytics
- direct write to school_core_db

Important rules:
- Admission owns applicant before conversion
- School Core owns student after conversion
- conversion must use gRPC
- double conversion prevented
- applicant documents are Restricted files

Important docs:
- docs/17-sprint-4-task-prompts.md
- docs/02-service-boundary.md
- docs/04-api-contract.md
- docs/05-event-contract.md

Use the Standard Sprint Plan Output Format.
```

---

## 11. Sprint 5 Planning Prompt — Finance / SPP

```text
[Use Global Prompt Header]

Buat dokumen Sprint Plan untuk:

Sprint 5 — Finance / SPP

Target output:
docs/34-sprint-5-plan.md

Sprint objective:
Membangun proses Finance/SPP MVP berbasis manual payment: fee type, fee scheme, fee policy, sibling discount, bill generation, payment proof upload, payment verification/rejection, receipt, dan void payment approval.

Main scope:
- finance_db migrations
- fee_types
- fee_schemes
- fee_scheme_items
- student_fee_policies
- sibling_discount_rules
- student_bills
- student_bill_items
- student_payments
- payment_proofs
- payment_receipts
- payment_reconciliations
- finance approval requests
- fee policy approval
- bill generation with snapshots
- manual payment proof
- verify/reject payment
- receipt generation
- outstanding/tunggakan
- void payment with approval
- finance events
- audit sensitive actions

Out of scope:
- payment gateway
- automatic bank reconciliation
- full accounting ledger
- payroll
- tax

Important rules:
- use decimal, never float
- free SPP/discount is fee policy, not student status
- bill item stores base_amount, discount_amount, final_amount, applied_policy_snapshot
- bill generation idempotent
- parent can only access linked child bills
- payment proof is Restricted file
- void requires approval

Important docs:
- docs/18-sprint-5-task-prompts.md
- docs/03-data-model-mvp.md
- docs/05-event-contract.md
- SKILLS.md Finance Skill

Use the Standard Sprint Plan Output Format.
```

---

## 12. Sprint 6 Planning Prompt — Academic Basic

```text
[Use Global Prompt Header]

Buat dokumen Sprint Plan untuk:

Sprint 6 — Academic Basic

Target output:
docs/35-sprint-6-plan.md

Sprint objective:
Membangun fondasi akademik dasar: curriculum, subject, subject group, class subject, schedule, teacher schedule view, attendance input, dan attendance correction.

Main scope:
- academic_db migrations
- curriculums
- learning_phases
- subjects
- subject_groups
- class_subjects
- schedules
- student_attendances
- curriculum baseline
- subject CRUD
- class subject assignment
- schedule CRUD
- teacher schedule view
- attendance input
- attendance correction
- academic basic events
- audit correction

Out of scope:
- report card
- grade book
- full LMS
- advanced timetable optimization
- BK/UKS detail

Important rules:
- Academic owns subject/schedule/attendance
- School Core owns student/teacher/class
- use reference IDs only
- do not query school_core_db directly
- Guru can only input attendance for assigned class/subject
- attendance correction requires reason and audit

Important docs:
- docs/19-sprint-6-task-prompts.md
- docs/02-service-boundary.md
- docs/06-ui-screen-user-flow.md

Use the Standard Sprint Plan Output Format.
```

---

## 13. Sprint 7 Planning Prompt — Report Card / E-Rapor Basic

```text
[Use Global Prompt Header]

Buat dokumen Sprint Plan untuk:

Sprint 7 — Report Card / E-Rapor Basic

Target output:
docs/36-sprint-7-plan.md

Sprint objective:
Membangun proses nilai dan rapor dasar: assessment components, schemes, grade book, score input, grade book submit/review, report card generation, publish/lock, dan revision approval.

Main scope:
- assessment_components
- assessment_schemes
- grade_books
- student_scores
- report_templates
- report_cards
- report_card_items
- academic approval requests
- score input
- grade book submit
- Wali Kelas review
- Kepala Sekolah publish
- report card lock
- parent/student published view
- revision after publish with approval
- report card events
- audit sensitive actions

Out of scope:
- full LMS
- advanced analytics
- national e-rapor integration
- offline score input
- final PDF design beyond MVP placeholder

Important rules:
- Guru can input score only for assigned class/subject
- Wali Kelas reviews assigned class
- Kepala Sekolah publishes report card
- published report card locked
- revision after publish requires approval and audit
- parent/student can only view published report card
- report card stores student snapshot

Important docs:
- docs/20-sprint-7-task-prompts.md
- docs/06-ui-screen-user-flow.md
- docs/07-test-plan-acceptance-criteria.md

Use the Standard Sprint Plan Output Format.
```

---

## 14. Sprint 8 Planning Prompt — Communication / Notification

```text
[Use Global Prompt Header]

Buat dokumen Sprint Plan untuk:

Sprint 8 — Communication / Notification

Target output:
docs/37-sprint-8-plan.md

Sprint objective:
Membangun announcement dan notification system berbasis event: announcement CRUD/publish, target audience, notification templates, in-app notifications, delivery logs, provider abstraction untuk FCM/email, preferences, dan event consumers.

Main scope:
- communication_db migrations
- announcements
- announcement_targets
- notifications
- notification_templates
- notification_deliveries
- notification_preferences
- announcement publish
- target by foundation/school/class/role/user
- event consumers
- in-app notification
- read/unread
- FCM mock/provider abstraction
- email mock/provider abstraction
- delivery log
- notification preferences

Out of scope:
- WhatsApp
- SMS
- marketing automation
- advanced campaign builder

Important rules:
- notification must be event-driven
- business services must not call FCM/email directly
- Confidential data must not be included in notification body
- critical notifications cannot be fully disabled
- event consumers must be idempotent

Important docs:
- docs/21-sprint-8-task-prompts.md
- docs/05-event-contract.md
- docs/06-ui-screen-user-flow.md

Use the Standard Sprint Plan Output Format.
```

---

## 15. Sprint 9 Planning Prompt — Reporting Dashboard

```text
[Use Global Prompt Header]

Buat dokumen Sprint Plan untuk:

Sprint 9 — Reporting Dashboard

Target output:
docs/38-sprint-9-plan.md

Sprint objective:
Membangun Reporting Service sebagai read model/projection untuk dashboard Yayasan, Sekolah, Guru, dan Orang Tua/Siswa berbasis event dari service operasional.

Main scope:
- reporting_db migrations
- dashboard projection tables
- processed_events
- projection offsets
- RabbitMQ consumer infrastructure
- student/teacher summary projection
- admission summary projection
- finance summary projection
- attendance summary projection
- academic progress projection
- dashboard APIs
- scheduled rebuild/sync skeleton
- idempotent consumers

Out of scope:
- advanced BI
- global search
- data warehouse
- direct query to operational DBs

Important rules:
- Reporting reads reporting_db only
- Reporting does not query operational service DBs directly
- projections built from events
- duplicate events skipped by processed_events
- dashboard endpoints enforce actor role/scope
- Reporting is not source of truth

Important docs:
- docs/22-sprint-9-task-prompts.md
- docs/05-event-contract.md
- docs/07-test-plan-acceptance-criteria.md

Use the Standard Sprint Plan Output Format.
```

---

## 16. Sprint 10 Planning Prompt — Security, Observability, Backup, UAT Hardening

```text
[Use Global Prompt Header]

Buat dokumen Sprint Plan untuk:

Sprint 10 — Security, Observability, Backup, and UAT Hardening

Target output:
docs/39-sprint-10-plan.md

Sprint objective:
Melakukan hardening MVP sebelum pilot/production: security review, permission/scope regression, audit log review, metrics/logging, backup/restore, UAT checklist, release readiness, dan rollback documentation.

Main scope:
- security baseline review
- permission/scope audit
- object-level authorization tests
- audit log consistency
- structured JSON logging review
- Prometheus metrics
- Loki/Grafana setup
- RabbitMQ queue/DLQ monitoring
- backup script
- restore procedure
- restore test documentation
- UAT checklist
- regression checklist
- production deployment readiness
- rollback documentation
- bug fixing

Out of scope:
- Kubernetes
- advanced SIEM
- full penetration test unless separately planned
- WAF unless needed
- new feature development

Important rules:
- no Critical/High bugs in MVP core flows before production
- production deployment requires manual approval
- backup classified Confidential
- restore test must be performed
- logs must not contain sensitive raw data
- audit logs separate from application logs
- metrics must not expose sensitive data

Important docs:
- docs/23-sprint-10-task-prompts.md
- docs/07-test-plan-acceptance-criteria.md
- docs/11-github-repository-rules.md
- docs/24-local-development-guide.md

Use the Standard Sprint Plan Output Format.
```

---

## 17. Prompt for Generating All Sprint Plans at Once

Use this only if the AI Agent has enough context and token capacity.

```text
[Use Global Prompt Header]

Buat dokumen sprint plan untuk semua sprint berikut:

- Sprint 0 — Project Foundation
- Sprint 1 — Identity & Access
- Sprint 2 — School Core
- Sprint 3 — File Management + Import Excel
- Sprint 4 — PPDB
- Sprint 5 — Finance/SPP
- Sprint 6 — Academic Basic
- Sprint 7 — Report Card/E-Rapor Basic
- Sprint 8 — Communication/Notification
- Sprint 9 — Reporting Dashboard
- Sprint 10 — Security, Observability, Backup, UAT Hardening

Output:
Buat 11 dokumen terpisah:

- docs/29-sprint-0-plan.md
- docs/30-sprint-1-plan.md
- docs/31-sprint-2-plan.md
- docs/32-sprint-3-plan.md
- docs/33-sprint-4-plan.md
- docs/34-sprint-5-plan.md
- docs/35-sprint-6-plan.md
- docs/36-sprint-7-plan.md
- docs/37-sprint-8-plan.md
- docs/38-sprint-9-plan.md
- docs/39-sprint-10-plan.md

Gunakan Standard Sprint Plan Output Format untuk setiap dokumen.

Jika output terlalu panjang, buat satu sprint per respons dan mulai dari Sprint 0.
```

---

## 18. Recommended GitHub Issue Format Generated from Sprint Plan

When the AI Agent creates GitHub issue list, use this format:

```text
Title:
Sprint N Task X.Y — [Task Name]

Description:
## Objective
[Objective]

## Scope
- [item]
- [item]

## Out of Scope
- [item]
- [item]

## Acceptance Criteria
- [ ] [criteria]
- [ ] [criteria]

## Technical Notes
- Target service/app:
- Data model impact:
- API/proto/event impact:
- Permission/scope:
- Audit:
- Tests:

## AI Agent Instruction
Use:
- AGENTS.md
- SKILLS.md
- docs/[sprint-task-prompt].md

## Definition of Done
- [ ] Code/docs completed
- [ ] Tests added/updated
- [ ] Tests pass
- [ ] Lint pass
- [ ] Permission/scope checked
- [ ] Audit/event added if required
- [ ] Docs updated if required
```

Recommended labels:

```text
type: feature
area: [module]
sprint: N
priority: [critical/high/medium/low]
status: ready
ai: ready / ai: needs-context
risk: [low/medium/high]
review: [backend/frontend/mobile/infra/qa/security]
```

---

## 19. Final Notes

Use this document to generate sprint plans before creating GitHub issues.

Do not use this document directly for coding.

For coding, use:

```text
13-sprint-0-task-prompts.md
14-sprint-1-task-prompts.md
15-sprint-2-task-prompts.md
...
23-sprint-10-task-prompts.md
```

## GitHub Project Management Context

When generating or updating this document, also consider:

```text
docs/25-github-project-management.md
.github/CODEOWNERS
.github/pull_request_template.md
.github/workflows/ci.yml
.github/ISSUE_TEMPLATE/feature_task.yml
.github/ISSUE_TEMPLATE/bug_report.yml
.github/ISSUE_TEMPLATE/ai_agent_task.yml
.github/ISSUE_TEMPLATE/security_review.yml
.github/ISSUE_TEMPLATE/qa_uat.yml
```

Ensure GitHub labels, milestones, project fields, issue lifecycle, PR workflow, AI Agent task workflow, QA/UAT workflow, and release readiness are aligned with `docs/25-github-project-management.md`.

## Required Inputs After Final Planning Docs

When generating Sprint Plan documents, AI Agent must read:

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/25-product-requirement-document.md
docs/26-development-plan.md
docs/27-workflow.md
docs/25-github-project-management.md
docs/10-sprint-backlog-mvp.md
active sprint task prompt from docs/13 through docs/23
```

Sprint plans must align with:

```text
- PRD MVP scope and non-goals
- Development Plan sprint strategy
- Workflow/SOP rules
- GitHub Project Management fields, labels, milestones, issue lifecycle, and QA/UAT workflow
```

## Generated Sprint Plan Status

The following Sprint Plan documents have already been generated:

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

Use this prompt pack only when regenerating or updating sprint plan documents.

When executing sprint tasks, read the final sprint plan first, then read the sprint task prompt.
