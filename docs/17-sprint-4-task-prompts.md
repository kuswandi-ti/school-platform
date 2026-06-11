# 17 — Sprint 4 Task Prompts

Sprint: Sprint 4 — PPDB

Project: `school-platform`  
Status: AI Agent task prompt pack  
Usage: Copy one task prompt at a time into AI Agent. Keep tasks small and do not combine unrelated tasks.

---

## Global Rules for This Sprint

AI Agent must follow these rules for every task in this document:

```text
- Follow docs/01-technical-architecture.md.
- Follow docs/02-service-boundary.md.
- Follow docs/03-data-model-mvp.md.
- Follow docs/04-api-contract.md.
- Follow docs/05-event-contract.md.
- Follow docs/07-test-plan-acceptance-criteria.md.
- Follow docs/08-coding-standard.md.
- Follow docs/09-ai-agent-rules.md.
- Follow docs/10-sprint-backlog-mvp.md.
- Follow docs/11-github-repository-rules.md.
- Do not query another service database.
- Do not put business logic in API Gateway.
- Use Go, Chi, gRPC, pgx, sqlc, goose, slog, validator, testify.
- Use UUID primary keys.
- Use foundation_id and school_id correctly.
- Enforce authentication, permission, scope, and object-level authorization.
- Add audit log for sensitive actions.
- Publish events through the standard outbox/event mechanism when required.
- Use standard API response and error format.
- Add tests for success and negative cases.
- Do not log tokens, passwords, or Confidential data.
- Do not implement out-of-scope features.
```

---

# Task 4.1 — Create Admission Database Migrations

## Prompt

```text
You are working on `school-platform`.

Task:
Create Admission Database Migrations

Target:
services/admission-service

Goal:
Create admission_db schema for PPDB process.

Scope:
- Create admission_periods, applicants, applicant_guardians, applicant_documents, applicant_verifications, admission_decisions, admission_audit_logs.
- Add indexes for school, period, status, registration_number.

Out of Scope:
- Student master tables.
- Finance bills.

Rules:
- Applicant belongs to Admission until conversion.
- No FK to school_core_db.
- Use reference IDs.

Acceptance Criteria:
- Migrations run up/down.
- Registration number unique per scope.
- Status enums supported.

Tests Required:
- Migration tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 4.2 — Implement Admission Period CRUD

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Admission Period CRUD

Target:
services/admission-service + api-gateway

Goal:
Manage PPDB admission periods per school and academic year.

Scope:
- CRUD period.
- Open/close period.
- Status draft/open/closed/archived.
- Audit status changes.
- Publish period events.

Out of Scope:
- Public registration page.
- Complex quota engine.

Rules:
- School scope enforced.
- Only authorized roles can open/close.
- No cross-school access.

Acceptance Criteria:
- Period can be created/opened/closed.
- Unauthorized role rejected.
- Event/audit recorded.

Tests Required:
- CRUD tests.
- Permission/scope tests.
- Event tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 4.3 — Implement Applicant Submission

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Applicant Submission

Target:
services/admission-service

Goal:
Create applicant and guardian data in Admission Service.

Scope:
- Create applicant.
- Create applicant_guardians.
- Generate registration_number.
- Status draft/submitted.
- Publish applicant submitted event.

Out of Scope:
- Convert to student.
- Document verification.

Rules:
- Admission owns applicant.
- Use numbering standard.
- No insert to school_core_db.

Acceptance Criteria:
- Applicant submitted.
- Registration number generated unique.
- Guardian data stored.
- Event published.

Tests Required:
- Submission tests.
- Number uniqueness test.
- Scope tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 4.4 — Implement Applicant Document Upload

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Applicant Document Upload

Target:
admission-service

Goal:
Upload and track PPDB documents.

Scope:
- Upload document metadata.
- Document type configurable enough for MVP.
- Status uploaded/verified/rejected.
- File private/restricted.
- Publish document_uploaded event.

Out of Scope:
- Central file service.
- Virus scanning.

Rules:
- File private.
- Applicant scope validated.
- No raw file logs.

Acceptance Criteria:
- Document upload succeeds.
- Invalid file rejected.
- Out-of-scope access rejected.
- Event published.

Tests Required:
- Upload tests.
- File validation tests.
- Scope tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 4.5 — Implement Document Verification

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Document Verification

Target:
admission-service

Goal:
Allow TU/Staff to verify or reject applicant documents.

Scope:
- Verification usecase.
- Set status verified/rejected.
- Notes.
- verified_by/verified_at.
- Audit log.
- Publish applicant.verified if all required docs verified.

Out of Scope:
- Final applicant acceptance.

Rules:
- TU/Staff scope required.
- Reason required for rejection.
- Audit sensitive action.

Acceptance Criteria:
- Document can be verified.
- Document can be rejected with reason.
- Applicant verification state updates.

Tests Required:
- Verification tests.
- Rejection reason tests.
- Audit tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 4.6 — Implement Applicant Accept/Reject Decision

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Applicant Accept/Reject Decision

Target:
admission-service

Goal:
Allow Kepala Sekolah to accept/reject PPDB applicant.

Scope:
- Decision endpoint.
- Status accepted/rejected.
- Reason for rejection.
- approval_request_id if needed.
- Audit.
- Publish accepted/rejected event.

Out of Scope:
- Student conversion.
- Payment/registration bills.

Rules:
- Kepala Sekolah school scope required.
- TU cannot final accept if policy says Kepala Sekolah.
- Cross-school rejected.

Acceptance Criteria:
- Applicant accepted.
- Applicant rejected.
- Unauthorized role rejected.
- Event/audit recorded.

Tests Required:
- Accept tests.
- Reject tests.
- Permission/scope tests.
- Event/audit tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 4.7 — Implement Convert Applicant to Student

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Convert Applicant to Student

Target:
admission-service + school-core-service

Goal:
Convert accepted applicant into School Core student through gRPC.

Scope:
- Add gRPC call to School Core.
- Create student and guardian in School Core.
- Store converted_student_id.
- Prevent double conversion.
- Use Idempotency-Key.
- Publish converted event.

Out of Scope:
- Direct DB insert.
- Complex enrollment fees.

Rules:
- Admission must not write school_core_db.
- Only accepted applicant can convert.
- Conversion idempotent.

Acceptance Criteria:
- Accepted applicant converted.
- School Core owns new student.
- Double conversion prevented.
- Events published.

Tests Required:
- Conversion tests.
- Double conversion tests.
- gRPC integration tests.
- Event tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 4.8 — Sprint 4 Final Verification

## Prompt

```text
You are working on `school-platform`.

Task:
Sprint 4 Final Verification

Target:
admission-service

Goal:
Verify PPDB end-to-end flow.

Scope:
- Run PPDB flow from period to conversion.
- Verify file privacy.
- Verify scope.
- Verify events.
- Produce report.

Out of Scope:
- Finance integration beyond references.

Rules:
- Document missing items honestly.

Acceptance Criteria:
- PPDB flow works end-to-end.
- No direct write to school_core_db from Admission.

Tests Required:
- Full PPDB tests.
- Manual smoke test.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

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
docs/33-sprint-4-plan.md if it exists
docs/17-sprint-4-task-prompts.md
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
docs/33-sprint-4-plan.md
docs/40-github-repository-setup-labels-project.md
docs/25-github-project-management.md
docs/27-workflow.md
```

Use `docs/33-sprint-4-plan.md` for sprint-level issue plan, dependencies, risks, and exit criteria.

Use `docs/40-github-repository-setup-labels-project.md` for labels, milestones, project fields, and repository workflow alignment.
