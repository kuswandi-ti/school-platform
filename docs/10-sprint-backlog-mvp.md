# 10 — Sprint Backlog MVP

Project: `school-platform`  
Status: Final decision for MVP planning  
Scope: Sprint order, objective, scope, out of scope, acceptance criteria, and AI Agent task direction.

---

## 1. Purpose

This document defines MVP sprint priorities based on technical and business dependencies.

The sprint order is designed to avoid building features before their required foundation exists.

---

## 2. Core Sprint Order

MVP sprint sequence:

```text
Sprint 0  : Project Foundation
Sprint 1  : Identity & Access
Sprint 2  : School Core
Sprint 3  : File Management + Import Excel
Sprint 4  : PPDB
Sprint 5  : Finance / SPP
Sprint 6  : Academic Dasar
Sprint 7  : Report Card / E-Rapor Dasar
Sprint 8  : Communication / Notification
Sprint 9  : Reporting Dashboard
Sprint 10 : Security, Observability, Backup, and UAT Hardening
```

Milestones:

```text
Milestone 1: Platform Foundation
Milestone 2: Admission & Finance
Milestone 3: Academic & Communication
Milestone 4: Reporting & Production Readiness
```

---

## 3. Sprint Principles

Rules:

```text
Build technical foundation first.
Identity before all protected modules.
School Core before PPDB, Finance, and Academic.
File Management before PPDB/payment proof/report card PDF.
PPDB before applicant-to-student conversion.
Finance after student data is stable.
Academic after student/teacher/class data is stable.
Reporting after important events exist.
Security/audit/permission/logging/test must start from early sprint.
```

Out of MVP:

```text
payment gateway
WhatsApp
payroll
full HR
asset inventory
library
BK/UKS detail
full LMS
global search
offline write mobile
Kubernetes
```

---

## 4. Sprint 0 — Project Foundation

### Objective

Prepare repository, local environment, service template, and CI foundation.

### Scope

```text
- Monorepo structure
- Docker Compose local
- PostgreSQL, Redis, RabbitMQ, MinIO
- Basic Go service template
- API Gateway skeleton
- Shared packages skeleton
- Proto folder
- OpenAPI folder
- Events folder
- GitHub Actions basic CI
- Makefile
- .env.example
- Health/readiness endpoint template
- Logging/correlation ID basic middleware
```

### Out of Scope

```text
full auth
domain modules
production deployment
Kubernetes
```

### Acceptance Criteria

```text
docker compose up starts core dependencies.
API Gateway health check works.
Sample service health check works.
GitHub Actions runs basic lint/test.
Repository contains apps/services/packages/infra/deploy/docs/scripts.
Makefile has setup/up/down/test/lint commands.
```

### AI Agent Task Examples

```text
Task 0.1: Create monorepo folder structure.
Task 0.2: Create docker-compose with PostgreSQL, Redis, RabbitMQ, MinIO.
Task 0.3: Create Go service template with healthz/readyz.
Task 0.4: Create API Gateway skeleton with Chi.
Task 0.5: Add Makefile commands.
Task 0.6: Add GitHub Actions basic CI.
```

---

## 5. Sprint 1 — Identity & Access

### Objective

Implement authentication, user, session, role, permission, and scope foundation.

### Scope

```text
- Identity Service
- users
- roles
- permissions
- role assignments
- login
- refresh token rotation
- logout
- password reset basic
- JWT access token
- hashed refresh token
- RBAC + scope context
- gRPC validate token/get user context
- API Gateway auth middleware
```

### Out of Scope

```text
2FA
full user self-service profile
advanced anomaly detection
```

### Acceptance Criteria

```text
User can login.
Access token is generated.
Refresh token is stored as hash and rotated.
Logout revokes session.
API Gateway rejects request without token.
User context includes role, permission, and scope.
Internal gRPC metadata carries actor context.
```

### Test Plan

```text
Login success/fail
Inactive user login rejected
Refresh token rotation
Refresh token reuse rejected/revoked
Logout revokes session
Protected endpoint without token rejected
Role/permission scope returned correctly
```

### AI Agent Task Examples

```text
Task 1.1: Create identity_db migrations for users, roles, permissions, sessions.
Task 1.2: Implement password hashing and login usecase.
Task 1.3: Implement refresh token rotation.
Task 1.4: Implement role/permission seed.
Task 1.5: Implement API Gateway auth middleware.
Task 1.6: Add identity tests.
```

---

## 6. Sprint 2 — School Core

### Objective

Implement core school master data.

### Scope

```text
- Foundation
- Schools TK/SD/SMP/SMA
- Academic year
- Semester
- Student master
- Guardian/parent master
- Teacher master dasar
- Grade level
- Class/rombel
- Student-class assignment
- Teacher assignment
- Homeroom assignment
```

### Out of Scope

```text
full HR
payroll
advanced room management
academic grades
finance billing
```

### Acceptance Criteria

```text
Admin Yayasan can view/manage school units.
TU/Staff can create student within school scope.
Guardian can be linked to student.
Teacher can be created.
Class can be created per academic year.
Student can be assigned to class.
Homeroom teacher can be assigned.
Cross-school scope is enforced.
```

### Test Plan

```text
CRUD student/guardian/teacher/class
Search/filter student
Assign student to class
Assign homeroom teacher
Scope test: SD user cannot access SMP data
Parent cannot access other student
```

### AI Agent Task Examples

```text
Task 2.1: Create School Core migrations.
Task 2.2: Implement foundation/school CRUD.
Task 2.3: Implement academic year/semester CRUD.
Task 2.4: Implement student/guardian CRUD.
Task 2.5: Implement teacher/class CRUD.
Task 2.6: Implement class assignment.
```

---

## 7. Sprint 3 — File Management + Import Excel

### Objective

Implement private file handling and initial data import.

### Scope

```text
- File metadata
- Upload file private
- Signed URL
- MIME/extension/size validation
- File classification
- Import templates
- Upload Excel
- Validate and preview
- Confirm import
- Import report
- Error report
```

### Out of Scope

```text
virus scanning integration
global file service
import score/payment/payroll/library/BK/UKS
```

### Acceptance Criteria

```text
File upload stores metadata and private object.
Invalid file is rejected.
Signed URL is created only after authorization.
Import template can be downloaded.
Import validates rows and displays preview.
Import does not insert before confirmation.
Import report is generated.
```

### Test Plan

```text
Upload valid/invalid file
Signed URL authorization
Import valid students
Import invalid extension
Import missing required column
Import duplicate student number
Import unknown class_code
Download error report
```

### AI Agent Task Examples

```text
Task 3.1: Implement file metadata table and upload service.
Task 3.2: Implement signed URL endpoint with authorization.
Task 3.3: Implement import_batch tables.
Task 3.4: Implement student import template.
Task 3.5: Implement import validation preview.
Task 3.6: Implement confirm import and report.
```

---

## 8. Sprint 4 — PPDB

### Objective

Implement admission workflow until applicant becomes student.

### Scope

```text
- Admission period
- Applicant
- Applicant guardian
- Applicant document
- Document verification
- Accept/reject applicant
- Convert applicant to student
- Event publication
- Notification status basic
```

### Out of Scope

```text
complex scoring/selection
payment gateway
advanced public portal
```

### Acceptance Criteria

```text
Admission period can be created.
Applicant can be submitted.
Document can be uploaded.
TU/Staff can verify document.
Kepala Sekolah can accept/reject applicant.
Accepted applicant can be converted to student.
Double conversion is prevented.
School Core owns student after conversion.
```

### Test Plan

```text
Create admission period
Submit applicant
Upload document
Verify document
Accept/reject applicant
Convert to student
Prevent double conversion
School scope test
```

### AI Agent Task Examples

```text
Task 4.1: Create admission_db migrations.
Task 4.2: Implement admission period CRUD.
Task 4.3: Implement applicant submission.
Task 4.4: Implement document upload/verification.
Task 4.5: Implement accept/reject.
Task 4.6: Implement convert applicant to student via School Core gRPC.
```

---

## 9. Sprint 5 — Finance / SPP

### Objective

Implement manual SPP billing and payment flow.

### Scope

```text
- Fee type
- Fee scheme
- Fee policy
- Free SPP
- Percentage/fixed discount
- Sibling discount rule
- Generate bill
- Bill item snapshot
- Manual payment
- Upload proof
- Verify/reject payment
- Receipt
- Outstanding bills
- Void request + approval
- Finance events
```

### Out of Scope

```text
payment gateway
automatic bank reconciliation
payroll
advanced accounting
```

### Acceptance Criteria

```text
Fee policy approved is applied during bill generation.
Bill stores base_amount, discount_amount, final_amount, applied_policy snapshot.
Generate bill is idempotent.
Parent only sees own child's bill.
Payment proof can be uploaded.
Bendahara can verify/reject payment.
Receipt is created after verified payment.
Void payment requires approval.
```

### Test Plan

```text
Generate normal bill
Generate bill with free_spp
Generate bill with sibling discount
Generate duplicate bill with same Idempotency-Key
Upload payment proof
Verify payment
Reject payment
Generate receipt
Void payment approval
Parent cannot access other child bill
```

### AI Agent Task Examples

```text
Task 5.1: Implement fee type and fee scheme.
Task 5.2: Implement fee policy approval flow.
Task 5.3: Implement sibling discount rules.
Task 5.4: Implement bill generation with snapshots.
Task 5.5: Implement payment proof upload.
Task 5.6: Implement payment verification/rejection.
Task 5.7: Implement receipt generation.
Task 5.8: Implement void payment approval.
```

---

## 10. Sprint 6 — Academic Dasar

### Objective

Implement academic master, schedule, and attendance.

### Scope

```text
- Curriculum baseline
- Subject
- Subject group
- Class subject
- Schedule
- Attendance
- Teacher assignment validation
- Attendance event
```

### Out of Scope

```text
full LMS
full report card
advanced timetable optimization
```

### Acceptance Criteria

```text
Subject can be created.
Schedule can be created for class and teacher.
Guru only sees assigned class/subject.
Guru can input attendance for assigned class.
Attendance is stored per date.
Attendance event is published.
Absent status can trigger notification if enabled.
```

### Test Plan

```text
Create subject
Create schedule
Mark attendance
Teacher assignment scope
Attendance event publication
Parent notification for absent if enabled
```

### AI Agent Task Examples

```text
Task 6.1: Create academic_db migrations for subjects/schedules/attendance.
Task 6.2: Implement subject CRUD.
Task 6.3: Implement schedule CRUD.
Task 6.4: Implement teacher schedule view.
Task 6.5: Implement attendance marking.
Task 6.6: Publish attendance event.
```

---

## 11. Sprint 7 — Report Card / E-Rapor Dasar

### Objective

Implement grade input and basic report card publishing.

### Scope

```text
- Assessment component
- Assessment scheme
- Grade book
- Student score
- Report template per level
- Report card
- Report card item
- Publish report card
- Lock after publish
- Revision request
- Basic PDF generation
```

### Out of Scope

```text
full LMS
complex analytics
advanced curriculum engine
```

### Acceptance Criteria

```text
Guru can input scores for assigned class/subject.
Grade book can be submitted.
Wali Kelas can review class report.
Kepala Sekolah can publish report card.
Published report card is locked.
Parent/student can view only published report card.
Revision after publish requires approval and audit log.
```

### Test Plan

```text
Input score
Submit grade book
Generate report card
Publish report card
Lock after publish
Parent view published report
Revision approval
Teacher cannot edit after publish
```

### AI Agent Task Examples

```text
Task 7.1: Implement assessment components and schemes.
Task 7.2: Implement grade book.
Task 7.3: Implement score input.
Task 7.4: Implement grade book submit.
Task 7.5: Implement report card generation.
Task 7.6: Implement publish and lock.
Task 7.7: Implement revision approval flow.
```

---

## 12. Sprint 8 — Communication / Notification

### Objective

Implement announcement and event-driven notification.

### Scope

```text
- Announcement
- Announcement target
- Notification template
- In-app notification
- Notification delivery log
- FCM token integration structure
- Email auth/reset/invitation/status important
- Notification preference
- Event consumers for Finance/Academic/PPDB
```

### Out of Scope

```text
WhatsApp
SMS
advanced campaign system
```

### Acceptance Criteria

```text
Announcement can be created and published by authorized role.
Announcement targets school/role/class if applicable.
Notification is created from event.
In-app notification is stored.
FCM is mocked/active according to environment.
Email is sent only for selected events.
Confidential data is not included in notification body.
Critical notification cannot be fully disabled.
```

### Test Plan

```text
Publish announcement
Target by school/role/class
Payment verified notification
Report card published notification
Notification read/unread
FCM mocked
Email mocked
Notification preference respected
```

### AI Agent Task Examples

```text
Task 8.1: Implement announcement tables and CRUD.
Task 8.2: Implement announcement target.
Task 8.3: Implement notification templates.
Task 8.4: Implement in-app notification.
Task 8.5: Implement event consumer for finance and academic events.
Task 8.6: Implement delivery log and retry metadata.
```

---

## 13. Sprint 9 — Reporting Dashboard

### Objective

Implement dashboard projection and near real-time reporting.

### Scope

```text
- Reporting Service
- Event consumers
- Foundation dashboard
- School dashboard
- Teacher dashboard simple
- Parent/student dashboard simple
- Scheduled rebuild/sync
- Projection idempotency
```

### Out of Scope

```text
advanced analytics
global search
direct query to operational DB
BI warehouse
```

### Acceptance Criteria

```text
Dashboard reads from reporting_db.
Reporting consumes events.
Event processing is idempotent.
StudentCreated updates student summary.
PaymentVerified updates finance summary.
AttendanceMarked updates attendance summary.
ReportCardPublished updates academic progress.
Dashboard is scoped by role.
Scheduled rebuild can run.
```

### Test Plan

```text
Projection from StudentCreated
Projection from PaymentVerified
Projection from AttendanceMarked
Projection from ReportCardPublished
Duplicate event skipped
Foundation dashboard scope
School dashboard scope
Parent dashboard child scope
```

### AI Agent Task Examples

```text
Task 9.1: Create reporting_db migrations.
Task 9.2: Implement processed_events.
Task 9.3: Implement finance summary projection.
Task 9.4: Implement attendance summary projection.
Task 9.5: Implement dashboard API.
Task 9.6: Implement scheduled rebuild skeleton.
```

---

## 14. Sprint 10 — Security, Observability, Backup, and UAT Hardening

### Objective

Prepare MVP for pilot/production readiness.

### Scope

```text
- Security baseline review
- Permission/scope audit
- Object-level authorization test
- Audit log review
- Structured logging
- Metrics
- Loki/Grafana/Prometheus
- Backup script
- Restore test
- UAT checklist
- Bug fixing
- Basic performance review
- Production deployment preparation
```

### Out of Scope

```text
Kubernetes
advanced SIEM
advanced autoscaling
WAF unless needed
full penetration test unless scheduled separately
```

### Acceptance Criteria

```text
All MVP core flows pass UAT.
No open Critical/High bugs.
Daily backup works.
Restore test succeeds at least once.
Health/readiness endpoints work for all services.
Structured log and correlation_id work.
Metrics dashboard exists.
Production deploy manual approval is ready.
```

### Test Plan

```text
Full regression test
Security/scope test
Backup restore test
Deployment rollback test
Observability check
UAT checklist
```

### AI Agent Task Examples

```text
Task 10.1: Add missing authorization tests.
Task 10.2: Add audit log consistency checks.
Task 10.3: Implement Prometheus metrics endpoints.
Task 10.4: Implement backup script.
Task 10.5: Document restore procedure.
Task 10.6: Prepare UAT checklist.
```

---

## 15. Parallel Work Strategy

Backend:

```text
Identity → School Core → PPDB/Finance/Academic
```

Frontend:

```text
Layout → Auth UI → School Core UI → PPDB/Finance/Academic UI
```

QA:

```text
Prepare test plan before feature completion.
Start UAT checklist early.
Verify staging continuously.
```

Infrastructure:

```text
Docker/CI/CD → Staging → Observability → Backup → Production readiness
```

Rule:

```text
Do not parallelize too far before API contract and data model are stable.
```

---

## 16. Minimum Priority if Resources Are Limited

If team capacity is constrained, prioritize:

```text
1. Sprint 0 — Foundation
2. Sprint 1 — Identity
3. Sprint 2 — School Core
4. Sprint 3 — Import Excel
5. Sprint 5 — Finance/SPP
6. Sprint 6 — Academic Attendance
7. Sprint 7 — Report Card
8. Sprint 8 — Notification
9. Sprint 9 — Reporting
```

PPDB may move slightly later only if pilot starts from existing student data, but for full MVP, PPDB remains Sprint 4.

---

## 17. AI Agent Task Template

Every task should use this template:

```text
Task Title:

Context:
Service/module and architecture decision references.

Goal:
What must be implemented.

Scope:
What is included.

Out of Scope:
What must not be implemented.

Files to Create/Modify:
Expected files.

Rules:
Permission, scope, audit, event, database, coding conventions.

Acceptance Criteria:
Given-When-Then or checklist.

Tests:
Unit/integration/API/event/scope tests required.

Expected Output:
Summary of changes and how to run tests.
```

---

## 18. Definition of Done per Sprint

Every sprint must finish with:

```text
Scope implemented
Out-of-scope not implemented
Migrations added and tested
Tests added
Permission/scope checks verified
Audit log implemented for sensitive actions
Events published if required
Documentation updated if contracts changed
CI pass
QA can verify in staging
```

---

## 19. Final Summary

MVP implementation must follow dependency order:

```text
Foundation
→ Identity
→ School Core
→ File/Import
→ PPDB
→ Finance
→ Academic
→ Report Card
→ Notification
→ Reporting
→ Hardening
```

Security, audit, permission/scope, logging, correlation_id, and tests are not postponed to the end. Their foundation starts from early sprints.
