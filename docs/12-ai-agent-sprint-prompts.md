# 12 — AI Agent Sprint Prompts

Project: `school-platform`  
Status: Working guide for AI Agent implementation  
Scope: Prompt pack for Sprint 0–10 MVP implementation

---

## 1. Purpose

This document contains reusable AI Agent prompts for implementing the MVP of `school-platform`.

The goal is to keep AI-assisted implementation:

- focused
- consistent
- secure
- testable
- aligned with service boundaries
- aligned with coding standards
- aligned with GitHub workflow
- aligned with MVP scope

Use this document together with:

```text
01-technical-architecture.md
02-service-boundary.md
03-data-model-mvp.md
04-api-contract.md
05-event-contract.md
06-ui-screen-user-flow.md
07-test-plan-acceptance-criteria.md
08-coding-standard.md
09-ai-agent-rules.md
10-sprint-backlog-mvp.md
11-github-repository-rules.md
```

---

## 2. Global AI Agent Instruction

Use this instruction at the start of every AI Agent session.

```text
You are an expert software engineer working on `school-platform`, a Go microservice monorepo for an internal school foundation management system.

The system is built with:
- Backend: Go microservices
- API Gateway: Custom Go Gateway using Chi
- Internal communication: gRPC + protobuf
- External API: REST/JSON via API Gateway
- Database: PostgreSQL, one database per service
- DB access: pgx + sqlc
- Migration: goose
- Events: RabbitMQ topic exchange `domain.events`
- Cache/infra: Redis, RabbitMQ, MinIO
- Web frontend: Next.js + TypeScript + Tailwind CSS + shadcn/ui
- Mobile frontend: Flutter + Riverpod + Dio + Retrofit
- Logging: slog structured JSON logging
- Testing: Go testing package + testify

You must follow:
- service boundary decisions
- data ownership rules
- no cross-service database query
- API contract standards
- event contract standards
- coding standard
- AI Agent rules
- GitHub workflow rules

You must keep tasks small and scoped. Do not implement features outside the assigned task.

Every protected usecase must enforce:
- authentication
- permission check
- scope check
- object-level authorization

Every sensitive action must include:
- audit log
- reason if required
- approval flow if required
- event publication if required

Every task must include relevant tests.

Do not:
- put business logic in API Gateway
- query another service database
- hardcode foundation_id or school_id
- log tokens/passwords/Confidential data
- use float for finance calculation
- create public URLs for private files
- change API/proto/event contracts without updating docs
```

---

## 3. Universal Task Prompt Template

Use this template for each AI Agent task.

```text
Task Title:
[Clear task title]

Context:
You are working on `school-platform`.
Relevant docs:
- docs/01-technical-architecture.md
- docs/02-service-boundary.md
- docs/03-data-model-mvp.md
- docs/04-api-contract.md
- docs/05-event-contract.md
- docs/08-coding-standard.md
- docs/09-ai-agent-rules.md
- docs/10-sprint-backlog-mvp.md

Goal:
[Describe what must be implemented.]

Scope:
- [Included item 1]
- [Included item 2]
- [Included item 3]

Out of Scope:
- [Excluded item 1]
- [Excluded item 2]

Target Service/App:
- [service/app name]

Expected Files/Folders:
- [files/folders to create or modify]

Architecture Rules:
- Do not query another service database.
- Do not put business logic in API Gateway.
- Use actor context for protected usecases.
- Enforce permission and scope.
- Use foundation_id and school_id correctly.
- Use standard response/error format.
- Use gRPC/protobuf for internal API.
- Use RabbitMQ event contract if event is required.

Database Rules:
- Use goose migration.
- Use pgx + sqlc.
- Use parameterized queries.
- Use UUID primary keys.
- Use snake_case table/column names.
- Add indexes/unique constraints where needed.

Security Rules:
- No secrets in code.
- No sensitive data in logs.
- Private files by default.
- Object-level authorization for resource by ID.

Audit/Event Rules:
- Add audit log for sensitive actions.
- Publish event through outbox if domain event is required.
- Use standard event envelope.
- Include request_id and correlation_id.

Acceptance Criteria:
- [Given/When/Then or checklist]

Tests Required:
- Unit test:
- Integration test:
- API test:
- Permission/scope test:
- Event test if relevant:
- Audit test if relevant:

Definition of Done:
- Code implemented within scope.
- Tests added and passing.
- Lint passes.
- Permission/scope checks implemented.
- Audit/event implemented if required.
- Docs/contracts updated if changed.
- Summary of changes provided.
- How to test provided.
```

---

## 4. Sprint 0 — Project Foundation Prompt

```text
You are implementing Sprint 0 — Project Foundation for `school-platform`.

Objective:
Create the initial monorepo foundation, local development environment, service template, shared package structure, and basic CI.

Scope:
- Create monorepo structure:
  - apps/
  - services/
  - packages/
  - infra/
  - deploy/
  - docs/
  - scripts/
- Create Docker Compose for local dependencies:
  - PostgreSQL
  - Redis
  - RabbitMQ
  - MinIO
  - optional Mailpit
- Create initial Go service template.
- Create API Gateway skeleton using Chi.
- Add healthz and readyz endpoints.
- Add basic structured logging with slog.
- Add request_id and correlation_id middleware.
- Create packages/proto, packages/openapi, packages/events.
- Add .env.example files.
- Add Makefile commands.
- Add basic GitHub Actions CI.

Out of Scope:
- Full authentication implementation.
- Domain business modules.
- Production deployment.
- Kubernetes.
- Advanced observability.

Architecture Rules:
- Follow docs/08-coding-standard.md.
- Follow docs/09-ai-agent-rules.md.
- Keep business logic out of API Gateway.
- Keep service template generic and reusable.

Expected Output:
- Repository structure is created.
- Docker Compose runs local infrastructure.
- API Gateway starts and exposes /healthz and /readyz.
- Sample service starts and exposes /healthz and /readyz.
- Makefile includes setup/up/down/test/lint commands.
- CI runs basic lint/test.

Acceptance Criteria:
- `docker compose up` starts PostgreSQL, Redis, RabbitMQ, MinIO.
- API Gateway health check returns success.
- Sample service health check returns success.
- GitHub Actions workflow exists.
- .env.example exists.
- No production secrets are included.

Tests Required:
- Basic Go test for sample health handler.
- CI validation.
```

---

## 5. Sprint 1 — Identity & Access Prompt

```text
You are implementing Sprint 1 — Identity & Access for `school-platform`.

Objective:
Implement authentication, users, sessions, refresh token rotation, roles, permissions, role assignments, and actor context.

Target Service:
- identity-service
- api-gateway for REST mapping/auth middleware

Scope:
- identity_db migrations:
  - users
  - roles
  - permissions
  - role_permissions
  - user_role_assignments
  - user_sessions
  - user_devices if needed
- Password hashing using Argon2id or bcrypt.
- Login usecase.
- JWT access token.
- Hashed rotating refresh token.
- Logout/revoke session.
- Password reset basic.
- Role and permission seed for MVP roles.
- gRPC endpoint for validate token/get user context.
- API Gateway auth middleware.
- Standard REST endpoints:
  - POST /api/v1/auth/login
  - POST /api/v1/auth/refresh
  - POST /api/v1/auth/logout
  - GET /api/v1/me
  - GET /api/v1/me/permissions
  - GET /api/v1/me/context

Out of Scope:
- OAuth/social login.
- 2FA.
- Advanced anomaly detection.
- Full user profile management.

Rules:
- Refresh token must be stored hashed.
- Refresh token must rotate on use.
- Access token must be short-lived.
- API Gateway validates token and extracts actor context.
- Services must still perform authorization checks.
- Use standard response/error format.
- Use correlation_id and request_id.
- Do not log token/password.

Acceptance Criteria:
- Active user can login.
- Wrong password is rejected.
- Inactive/locked user is rejected.
- Access token is issued.
- Refresh token is stored as hash.
- Refresh token rotation works.
- Reuse of old refresh token is rejected/revoked.
- Logout revokes session.
- /me returns user context with roles, permissions, and scope.
- Protected endpoint rejects missing/invalid token.

Tests Required:
- Login success/fail.
- Refresh token rotation.
- Refresh token reuse detection.
- Logout revoke.
- Permission/scope context.
- API Gateway protected endpoint test.
- No token/password appears in logs.
```

---

## 6. Sprint 2 — School Core Prompt

```text
You are implementing Sprint 2 — School Core for `school-platform`.

Objective:
Implement core master data for foundation, school, academic year, semester, students, guardians, teachers, classes, and assignments.

Target Service:
- school-core-service
- api-gateway REST mapping as needed

Scope:
- school_core_db migrations:
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
- CRUD usecases for core entities.
- Scope-based access by foundation_id and school_id.
- Basic search/filter for students and teachers.
- Student-class assignment.
- Teacher assignment.
- Homeroom assignment.
- Domain events:
  - school.student.created
  - school.student.updated
  - school.teacher.created
  - school.class.created
  - school.student_class.assigned
  - school.homeroom.assigned

Out of Scope:
- HR complete.
- Payroll.
- Finance.
- Grade/report card.
- Import Excel.
- PPDB conversion.

Rules:
- School Core owns student/guardian/teacher/class master data.
- Do not create user credentials here; use Identity reference user_id only.
- Use UUID primary keys.
- Use foundation_id and school_id.
- No cross-service DB query.
- Use gRPC or events for cross-service interaction.
- Sensitive changes need audit log.
- Teacher subject_id may reference Academic Service, but do not query academic_db.

Acceptance Criteria:
- Admin Yayasan can view/manage school units within foundation.
- Kepala Sekolah can only see own school.
- TU/Staff can create student within school.
- Guardian can be linked to student.
- Teacher can be created.
- Class can be created per academic year.
- Student can be assigned to class.
- Homeroom teacher can be assigned.
- Cross-school access is rejected.

Tests Required:
- CRUD foundation/school/student/guardian/teacher/class.
- Search/filter.
- Scope test between schools.
- Student-class assignment test.
- Homeroom assignment test.
- Event publication test.
- Audit test for sensitive update.
```

---

## 7. Sprint 3 — File Management + Import Excel Prompt

```text
You are implementing Sprint 3 — File Management + Import Excel for `school-platform`.

Objective:
Implement private file handling and initial Excel import for students, guardians, teachers, classes, and assignments.

Target Service:
- school-core-service for initial import
- domain owner services for file metadata where applicable
- api-gateway REST mapping

Scope:
- File metadata structure for owner service.
- Upload private files to MinIO/S3-compatible storage.
- Validate MIME type, extension, file size.
- Generate signed URL after authorization.
- File classification:
  - public
  - internal
  - restricted
  - confidential
- Import tables:
  - import_batches
  - import_batch_rows
- Download import template.
- Upload Excel file.
- Validate column structure.
- Validate row data.
- Preview import result.
- Confirm import.
- Generate import report.
- Generate error report.

Supported import MVP:
- students
- guardians
- teachers
- classes
- student-class assignment via class_code
- homeroom/teacher assignment optional if data is ready

Out of Scope:
- Import grades.
- Import historical payments.
- Import payroll.
- Import asset/library/BK/UKS/alumni/cooperative.
- Virus scanning integration.
- Central File Service.

Rules:
- File private by default.
- Import file is Restricted data.
- Import must not directly insert before validation and preview.
- Raw import data must not be logged.
- Every import has import_batch_id.
- Every import row result is tracked.
- Signed URL requires permission/scope check.
- Download of Restricted/Confidential file requires audit log when applicable.

Acceptance Criteria:
- User can download template.
- User can upload valid Excel.
- Invalid extension is rejected.
- Missing required column is detected.
- Duplicate student number is detected.
- Unknown class_code is detected.
- Preview displays valid/warning/error rows.
- Confirm import creates valid records.
- Import report is available.
- Error report is downloadable.
- Import file is not public.

Tests Required:
- File upload valid/invalid.
- Signed URL authorization.
- Import valid students.
- Import missing column.
- Import duplicate student number.
- Import invalid birth date.
- Import unknown class_code.
- Confirm import.
- Permission test: Guru cannot import.
- Audit test for import completion.
```

---

## 8. Sprint 4 — PPDB Prompt

```text
You are implementing Sprint 4 — PPDB for `school-platform`.

Objective:
Implement admission process from applicant registration to accepted applicant conversion into student.

Target Service:
- admission-service
- school-core-service gRPC integration for conversion
- communication-service integration via event where needed
- api-gateway REST mapping

Scope:
- admission_db migrations:
  - admission_periods
  - applicants
  - applicant_guardians
  - applicant_documents
  - applicant_verifications
  - admission_decisions
- Admission period CRUD.
- Applicant submission.
- Applicant guardian data.
- Document upload reference.
- Document verification.
- Accept/reject applicant.
- Convert accepted applicant to student via School Core gRPC.
- Store converted_student_id.
- Prevent double conversion.
- Publish events:
  - admission.applicant.submitted
  - admission.applicant.document_uploaded
  - admission.applicant.verified
  - admission.applicant.accepted
  - admission.applicant.rejected
  - admission.applicant.converted_to_student

Out of Scope:
- Complex scoring/selection.
- Public marketing site.
- Payment gateway.
- Advanced admission analytics.

Rules:
- Admission owns applicant before conversion.
- School Core owns student after conversion.
- Admission must not insert directly into school_core_db.
- Conversion must use gRPC.
- Accept/reject must respect permission/scope.
- Sensitive decisions require audit log.
- Applicant document files are Restricted.
- Convert operation should use Idempotency-Key.

Acceptance Criteria:
- Admission period can be created.
- Applicant can be submitted.
- Applicant document can be uploaded.
- TU/Staff can verify documents.
- Kepala Sekolah can accept/reject applicant within school scope.
- Accepted applicant can be converted to student.
- Double conversion is prevented.
- School Core creates student and guardian.
- Events are published.
- Notifications can be triggered through events.

Tests Required:
- Create admission period.
- Submit applicant.
- Upload document.
- Verify document.
- Accept applicant.
- Reject applicant.
- Convert to student.
- Prevent double conversion.
- Scope test between schools.
- Audit test for accept/reject.
- Event publication test.
```

---

## 9. Sprint 5 — Finance / SPP Prompt

```text
You are implementing Sprint 5 — Finance / SPP for `school-platform`.

Objective:
Implement manual SPP billing, fee policy, discount, payment proof upload, verification, receipt, and void approval.

Target Service:
- finance-service
- school-core-service gRPC/reference validation
- communication-service via events
- reporting-service via events
- api-gateway REST mapping

Scope:
- finance_db migrations:
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
  - finance_approval_requests if local approval
- Fee type CRUD.
- Fee scheme CRUD.
- Fee policy per student:
  - normal
  - free_spp
  - percentage_discount
  - fixed_amount_discount
  - sibling_discount
  - scholarship
  - custom_fee
- Sibling discount configurable rule.
- Fee policy approval.
- Generate bill with snapshot.
- Manual payment creation.
- Upload payment proof.
- Verify/reject payment.
- Generate receipt.
- Outstanding/tunggakan list.
- Void payment request + approval.
- Publish events:
  - finance.fee_policy.approved
  - finance.bill.generated
  - finance.payment.proof_uploaded
  - finance.payment.verified
  - finance.payment.rejected
  - finance.payment.void_requested
  - finance.payment.voided
  - finance.receipt.generated

Out of Scope:
- Payment gateway.
- Automatic bank reconciliation.
- Full accounting ledger.
- Payroll.
- Tax.

Rules:
- Finance owns fee/bill/payment/receipt.
- Free SPP/discount is fee policy, not student status.
- Use decimal library. Do not use float.
- Bill item must store base_amount, discount_amount, final_amount, applied_policy_snapshot_json.
- Generate bill must be idempotent.
- Parent can only view linked child bills.
- Payment proof file is Restricted.
- Verify/reject payment requires audit log.
- Void payment requires approval.
- Use outbox pattern for important events.
- Do not query school_core_db directly.

Acceptance Criteria:
- Bendahara can create fee type/scheme.
- Fee policy can be submitted and approved.
- Bill generation applies approved fee policy.
- Bill snapshot is stored.
- Duplicate bill generation is prevented by Idempotency-Key.
- Parent can view own child bill.
- Parent can upload payment proof.
- Bendahara can verify/reject payment.
- Receipt is generated after verified payment.
- Void request requires approval and audit log.
- Reporting/notification events are published.

Tests Required:
- Create fee type.
- Create fee scheme.
- Create free_spp policy.
- Create sibling discount policy.
- Approve fee policy.
- Generate normal bill.
- Generate bill with discount/free_spp.
- Duplicate generation idempotency.
- Upload payment proof.
- Verify payment.
- Reject payment.
- Receipt generation.
- Void approval flow.
- Parent cannot access other child bill.
- Finance calculation does not use float.
- Event publication test.
- Audit test.
```

---

## 10. Sprint 6 — Academic Dasar Prompt

```text
You are implementing Sprint 6 — Academic Dasar for `school-platform`.

Objective:
Implement curriculum baseline, subject, schedule, class subject, and attendance.

Target Service:
- academic-service
- school-core-service gRPC/reference validation
- communication-service via events
- reporting-service via events
- api-gateway REST mapping

Scope:
- academic_db migrations:
  - curriculums
  - learning_phases
  - subjects
  - subject_groups
  - class_subjects
  - schedules
  - student_attendances
- Curriculum baseline.
- Subject CRUD.
- Class subject assignment.
- Schedule CRUD.
- Teacher schedule view.
- Attendance input.
- Attendance correction if allowed.
- Publish events:
  - academic.subject.created
  - academic.schedule.created
  - academic.attendance.marked
  - academic.attendance.corrected
  - academic.attendance.absent_recorded if needed

Out of Scope:
- Full LMS.
- Report card.
- Advanced timetable optimization.
- BK/UKS detail.

Rules:
- Academic owns subject/schedule/attendance.
- School Core owns student/teacher/class.
- Use reference IDs only.
- Do not query school_core_db directly.
- Validate teacher assignment/scope.
- Guru can only input attendance for assigned class/subject.
- Absence notification must not include Confidential detail.
- Attendance changes should be audited if sensitive/corrected.

Acceptance Criteria:
- Subject can be created.
- Schedule can be created.
- Guru sees assigned schedule.
- Guru can mark attendance for assigned class.
- Guru cannot mark attendance for unassigned class.
- Attendance event is published.
- Reporting projection can consume attendance event.

Tests Required:
- Create subject.
- Create schedule.
- Teacher assigned schedule view.
- Mark attendance success.
- Mark attendance unauthorized teacher rejected.
- Scope test between schools.
- Attendance event publication.
- Audit test for correction if implemented.
```

---

## 11. Sprint 7 — Report Card / E-Rapor Dasar Prompt

```text
You are implementing Sprint 7 — Report Card / E-Rapor Dasar for `school-platform`.

Objective:
Implement grade book, score input, report card generation, review, publish, lock, and revision approval.

Target Service:
- academic-service
- school-core-service reference validation
- communication-service via events
- reporting-service via events
- api-gateway REST mapping

Scope:
- academic_db migrations:
  - assessment_components
  - assessment_schemes
  - grade_books
  - student_scores
  - report_templates
  - report_cards
  - report_card_items
  - academic_approval_requests if local approval
- Assessment component.
- Assessment scheme.
- Grade book creation.
- Score input.
- Submit grade book.
- Wali Kelas review.
- Generate report card.
- Publish report card.
- Lock after publish.
- Revision request after publish.
- Basic report card PDF metadata.
- Publish events:
  - academic.grade_book.created
  - academic.grade_book.submitted
  - academic.grade_book.approved
  - academic.report_card.generated
  - academic.report_card.published
  - academic.report_card.revision_requested
  - academic.report_card.revised

Out of Scope:
- Full LMS.
- Advanced analytics.
- Complex national e-rapor integration.
- Offline score input.

Rules:
- Guru can input score only for assigned class/subject.
- Wali Kelas can review assigned class.
- Kepala Sekolah can publish report card.
- Published report card is locked.
- Revision after publish requires approval and audit log.
- Parent/student can only view published report card.
- Report card stores student snapshot.
- Do not use float for score weighting if decimal precision is needed.
- Do not query school_core_db directly.

Acceptance Criteria:
- Guru can input score for assigned class/subject.
- Guru cannot input score for unassigned class.
- Grade book can be submitted.
- Wali Kelas can review report card summary.
- Kepala Sekolah can publish report card.
- Published report card becomes locked.
- Parent/student can view published report card only.
- Revision after publish requires approval.
- Events are published.

Tests Required:
- Input score success.
- Input score unauthorized teacher rejected.
- Submit grade book.
- Generate report card.
- Publish report card.
- Lock after publish.
- Parent view published report.
- Parent cannot view unpublished report.
- Revision approval.
- Event publication.
- Audit test.
```

---

## 12. Sprint 8 — Communication / Notification Prompt

```text
You are implementing Sprint 8 — Communication / Notification for `school-platform`.

Objective:
Implement announcement, notification template, in-app notification, delivery log, FCM structure, email limited use, and event-driven notification.

Target Service:
- communication-service
- identity-service for user/device target through gRPC/event projection
- api-gateway REST mapping

Scope:
- communication_db migrations:
  - announcements
  - announcement_targets
  - notifications
  - notification_templates
  - notification_deliveries
  - notification_preferences
  - letters if basic letters included
- Announcement CRUD.
- Announcement publish.
- Announcement target by foundation/school/class/role/user.
- Notification template.
- In-app notification.
- Notification read/unread.
- Delivery log.
- FCM provider abstraction/mock.
- Email provider abstraction/mock.
- Event consumers for:
  - finance.bill.generated
  - finance.payment.verified
  - finance.payment.rejected
  - academic.report_card.published
  - admission.applicant.accepted
  - admission.applicant.rejected
  - approval.request.created

Out of Scope:
- WhatsApp.
- SMS.
- Advanced campaign builder.
- Marketing automation.

Rules:
- Notification must be event-driven.
- Business services must not call FCM/email directly.
- Confidential data must not be included in notification body.
- Critical notifications cannot be fully disabled.
- Delivery attempt must be logged.
- User preference respected where allowed.
- Use retry/DLQ for event consumer.

Acceptance Criteria:
- Authorized user can create and publish announcement.
- Announcement target works by school/role/class where applicable.
- Event creates notification.
- In-app notification is stored.
- User can mark notification as read.
- FCM delivery can be mocked.
- Email delivery can be mocked.
- Delivery log records status.
- Confidential data is not exposed.

Tests Required:
- Create announcement.
- Publish announcement.
- Target school/role/class.
- Notification created from PaymentVerified.
- Notification created from ReportCardPublished.
- Read/unread notification.
- Preference respected.
- Critical notification cannot be disabled.
- No Confidential detail in notification.
- Event consumer idempotency.
```

---

## 13. Sprint 9 — Reporting Dashboard Prompt

```text
You are implementing Sprint 9 — Reporting Dashboard for `school-platform`.

Objective:
Implement Reporting Service projection/read model and MVP dashboards.

Target Service:
- reporting-service
- api-gateway REST mapping

Scope:
- reporting_db migrations:
  - foundation_dashboard_metrics
  - school_dashboard_metrics
  - student_summary_metrics
  - teacher_summary_metrics
  - admission_summary_metrics
  - finance_summary_metrics
  - attendance_summary_metrics
  - academic_progress_metrics
  - approval_pending_metrics
  - notification_summary_metrics
  - processed_events
  - reporting_projection_offsets
- Event consumers for school/admission/finance/academic/approval events.
- Foundation dashboard API.
- School dashboard API.
- Teacher dashboard simple API.
- Parent/student dashboard simple API.
- Scheduled rebuild/sync skeleton.
- Idempotent projection.

Out of Scope:
- Advanced BI.
- Global search.
- Direct query to operational databases.
- Data warehouse.

Rules:
- Reporting reads reporting_db only.
- Reporting does not query operational service databases directly.
- Projection is built from events.
- Scheduled rebuild may use controlled APIs if needed, not raw DB access.
- Every event consumer must be idempotent.
- Dashboard data must be scoped by actor role/scope.
- Reporting is not source of truth.

Acceptance Criteria:
- Dashboard reads from reporting_db.
- StudentCreated updates student summary.
- PaymentVerified updates finance summary.
- AttendanceMarked updates attendance summary.
- ReportCardPublished updates academic progress.
- Duplicate event is skipped.
- Admin Yayasan sees foundation-level metrics.
- Kepala Sekolah sees school metrics.
- Parent sees only linked child summary.
- Scheduled rebuild command exists.

Tests Required:
- Projection from StudentCreated.
- Projection from PaymentVerified.
- Projection from AttendanceMarked.
- Projection from ReportCardPublished.
- Idempotency duplicate event test.
- Dashboard scope test.
- Reporting does not access operational DB config.
```

---

## 14. Sprint 10 — Security, Observability, Backup, UAT Hardening Prompt

```text
You are implementing Sprint 10 — Security, Observability, Backup, and UAT Hardening for `school-platform`.

Objective:
Harden MVP before pilot/production.

Target Areas:
- all services
- infrastructure
- CI/CD
- staging/production deployment
- QA/UAT

Scope:
- Security baseline review.
- Permission/scope audit.
- Object-level authorization tests.
- Audit log consistency check.
- Structured JSON logging review.
- Prometheus metrics endpoints.
- Loki/Grafana/Prometheus setup.
- RabbitMQ queue/DLQ monitoring.
- Backup script.
- Restore procedure.
- Restore test documentation.
- UAT checklist.
- Regression checklist.
- Production deployment readiness.
- Rollback documentation.
- Bug fixing.

Out of Scope:
- Kubernetes.
- Advanced SIEM.
- Advanced anomaly detection.
- Full penetration test unless separately planned.
- WAF unless needed.

Rules:
- No Critical/High bugs in MVP core flows before production.
- Production deployment requires manual approval.
- Backup must be encrypted/treated as Confidential.
- Restore test must be performed at least once.
- Logs must not contain sensitive raw data.
- Audit log remains separate from application log.
- Metrics must not expose sensitive data.

Acceptance Criteria:
- All MVP core flows pass UAT.
- No open Critical/High bug.
- Health/readiness endpoints work for all services.
- Structured logs include correlation_id.
- Basic metrics available.
- Backup script works.
- Restore test succeeds.
- Production deployment workflow is approval-gated.
- Regression checklist is documented.
- Rollback guide is documented.

Tests Required:
- Full regression test.
- Security/scope test.
- Object-level authorization test.
- Backup/restore test.
- Deployment rollback test if possible.
- Observability check.
- DLQ alert check.
```

---

## 15. Sprint-to-Task Breakdown Guidance

When creating detailed task prompts from this document, use this structure:

```text
Sprint N Task NN — [Task Name]

Context:
[Relevant service and docs]

Goal:
[One clear feature/subfeature]

Scope:
[Small list]

Out of Scope:
[Explicit exclusions]

Implementation Notes:
[Architecture/coding requirements]

Acceptance Criteria:
[Given/When/Then]

Tests:
[Specific tests]

Output:
[Expected AI Agent final response]
```

Good task size:

```text
1 service
1 feature
1 migration set
1 API/usecase flow
1 test set
```

Bad task size:

```text
Build the entire module.
Implement all finance features.
Create full school platform.
```

---

## 16. Suggested Detailed Task Split

### Sprint 0

```text
0.1 Create monorepo structure
0.2 Create Docker Compose dependencies
0.3 Create Go service template
0.4 Create API Gateway skeleton
0.5 Add Makefile
0.6 Add CI workflow
0.7 Add docs placeholders
```

### Sprint 1

```text
1.1 Identity migrations
1.2 Password hashing and user repository
1.3 Login usecase
1.4 Refresh token rotation
1.5 Logout/revoke session
1.6 Role/permission seed
1.7 API Gateway auth middleware
1.8 Identity tests
```

### Sprint 2

```text
2.1 Foundation/school model
2.2 Academic year/semester model
2.3 Student/guardian model
2.4 Teacher model
2.5 Class/grade level model
2.6 Student-class assignment
2.7 Homeroom assignment
2.8 School Core tests
```

### Sprint 3

```text
3.1 File metadata and storage abstraction
3.2 Signed URL authorization
3.3 Import batch tables
3.4 Student import template
3.5 Import validation preview
3.6 Confirm import
3.7 Error report
3.8 File/import tests
```

### Sprint 4

```text
4.1 Admission period
4.2 Applicant registration
4.3 Applicant guardian
4.4 Applicant document upload
4.5 Document verification
4.6 Applicant accept/reject
4.7 Convert applicant to student
4.8 PPDB tests
```

### Sprint 5

```text
5.1 Fee type
5.2 Fee scheme
5.3 Fee policy
5.4 Sibling discount
5.5 Bill generation
5.6 Payment proof upload
5.7 Payment verification/rejection
5.8 Receipt generation
5.9 Void payment approval
5.10 Finance tests
```

### Sprint 6

```text
6.1 Curriculum/subject
6.2 Class subject
6.3 Schedule
6.4 Teacher schedule view
6.5 Attendance input
6.6 Attendance correction
6.7 Academic basic tests
```

### Sprint 7

```text
7.1 Assessment components/schemes
7.2 Grade book
7.3 Score input
7.4 Grade book submit
7.5 Report card generation
7.6 Publish and lock report card
7.7 Revision approval
7.8 Report card tests
```

### Sprint 8

```text
8.1 Announcement
8.2 Announcement target
8.3 Notification template
8.4 In-app notification
8.5 Notification event consumer
8.6 Delivery log
8.7 Notification tests
```

### Sprint 9

```text
9.1 Reporting schema
9.2 Processed events
9.3 Student summary projection
9.4 Finance summary projection
9.5 Attendance summary projection
9.6 Academic progress projection
9.7 Dashboard APIs
9.8 Reporting tests
```

### Sprint 10

```text
10.1 Security review
10.2 Permission/scope test coverage
10.3 Audit log review
10.4 Metrics and logging check
10.5 Backup script
10.6 Restore procedure
10.7 UAT checklist
10.8 Release readiness
```

---

## 17. Required AI Agent Final Response Format

Every AI Agent implementation response should end with:

```text
Summary:
- [what changed]

Files changed:
- [file 1]
- [file 2]

Tests:
- [test added]
- [test command]

How to run:
```bash
[commands]
```

Notes:
- [limitations]
- [follow-up]
```

---

## 18. Final Summary

Use this prompt pack to guide AI Agent implementation.

Rules:

```text
Keep tasks small.
Keep scope explicit.
Always enforce permission/scope.
Always add tests.
Always preserve service boundary.
Always update contracts when changed.
Never query another service database.
Never put business logic in API Gateway.
Never log sensitive data.
```

## Required Final Planning Context

Before using any sprint prompt for implementation, AI Agent must also read the final planning documents:

```text
docs/25-product-requirement-document.md
docs/26-development-plan.md
docs/27-workflow.md
docs/25-github-project-management.md
```

Use these documents to understand:

```text
- product scope and MVP boundaries
- development and sprint execution strategy
- daily workflow/SOP
- GitHub issue/PR/QA/release tracking
```

The prompt files `docs/25-prd-prompt.md`, `docs/26-development-plan-prompt.md`, and `docs/27-workflow-prompt.md` are used for regeneration/update, not for day-to-day coding.
