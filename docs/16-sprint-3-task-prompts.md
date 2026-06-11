# 16 — Sprint 3 Task Prompts

Sprint: Sprint 3 — File Management + Import Excel

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

# Task 3.1 — Implement File Metadata and Storage Abstraction

## Prompt

```text
You are working on `school-platform`.

Task:
Implement File Metadata and Storage Abstraction

Target:
shared-go + school-core-service

Goal:
Create private file metadata and MinIO/S3 storage abstraction for owner services.

Scope:
- Create file metadata structure.
- Implement storage interface.
- Implement MinIO local adapter.
- Validate MIME/extension/size.
- Support classification public/internal/restricted/confidential.

Out of Scope:
- Central File Service.
- Virus scanning.
- Cloud production config.

Rules:
- Files private by default.
- Metadata stored in owner service.
- Do not return permanent public URL.
- No raw file logs.

Acceptance Criteria:
- Valid file stored private.
- Invalid type rejected.
- Metadata saved.
- Classification stored.

Tests Required:
- File upload tests.
- Validation tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 3.2 — Implement Signed URL Authorization

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Signed URL Authorization

Target:
school-core-service + api-gateway

Goal:
Generate signed URL only after permission/scope validation.

Scope:
- Add signed URL endpoint/domain-specific helper.
- Check actor scope.
- Short expiry based on classification.
- Audit restricted/confidential downloads.

Out of Scope:
- Public CDN.
- Permanent URLs.

Rules:
- Object-level authorization required.
- Restricted/Confidential access audited.
- No signed URL for out-of-scope user.

Acceptance Criteria:
- Authorized user gets signed URL.
- Unauthorized user rejected.
- Audit log recorded for restricted file.

Tests Required:
- Authorization tests.
- Audit tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 3.3 — Create Import Batch Tables and Repository

## Prompt

```text
You are working on `school-platform`.

Task:
Create Import Batch Tables and Repository

Target:
services/school-core-service

Goal:
Create import_batches and import_batch_rows foundation.

Scope:
- Migrations.
- sqlc queries.
- Repository/usecase skeleton.
- Status uploaded/validated/confirmed/processing/completed/failed.
- Track row result.

Out of Scope:
- Actual Excel parsing.
- All import types.

Rules:
- Import file Restricted.
- Raw data stored only as needed.
- No raw import data in logs.

Acceptance Criteria:
- Import batch can be created.
- Rows can be stored.
- Status can be updated.
- Error counts tracked.

Tests Required:
- Repository integration tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 3.4 — Implement Import Template Download

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Import Template Download

Target:
school-core-service + api-gateway

Goal:
Provide official Excel templates for initial data import.

Scope:
- Template for students.
- Template for teachers.
- Template for classes.
- Template docs/field descriptions.
- Download endpoint.

Out of Scope:
- Advanced custom template builder.

Rules:
- Templates must match validation fields.
- Do not include real personal data.
- Use synthetic examples.

Acceptance Criteria:
- User can download template.
- Template includes required columns.
- Only authorized roles can download if protected.

Tests Required:
- Template generation test.
- Permission test.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 3.5 — Implement Excel Upload and Validation Preview

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Excel Upload and Validation Preview

Target:
school-core-service

Goal:
Parse uploaded Excel and produce validation preview without inserting data.

Scope:
- Parse Excel.
- Validate required columns.
- Validate field format.
- Detect duplicate student number.
- Detect unknown class_code.
- Store import row statuses.
- Return preview summary.

Out of Scope:
- Confirm import.
- Rollback full import.
- Import scores/payments.

Rules:
- No DB insert into domain tables before confirm.
- No raw data logs.
- File private.

Acceptance Criteria:
- Valid rows marked valid.
- Invalid rows show errors.
- Warnings tracked.
- Preview available.

Tests Required:
- Valid file test.
- Missing column test.
- Duplicate test.
- Invalid date test.
- Unknown class test.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 3.6 — Implement Confirm Import

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Confirm Import

Target:
school-core-service

Goal:
Commit validated import rows into School Core domain tables.

Scope:
- Confirm import endpoint/usecase.
- Create students/guardians/teachers/classes based on import type.
- Track success/failed rows.
- Audit import completion.
- Publish events for created entities where needed.

Out of Scope:
- Automatic full rollback.
- Import payment/grade/payroll.

Rules:
- Only valid preview can be confirmed.
- Idempotency or status guard prevents double confirm.
- Actor scope enforced.

Acceptance Criteria:
- Confirm creates valid records.
- Double confirm rejected.
- Import report generated.
- Audit recorded.

Tests Required:
- Confirm import tests.
- Double confirm test.
- Scope test.
- Event/audit tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 3.7 — Implement Import Error Report

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Import Error Report

Target:
school-core-service + api-gateway

Goal:
Allow user to download import error report.

Scope:
- Generate error report from import rows.
- Domain-specific download endpoint.
- Private signed URL or direct generated file.
- Audit if restricted.

Out of Scope:
- Advanced analytics.

Rules:
- Only uploader/authorized role can download.
- No public file.
- Report contains row number and error message.

Acceptance Criteria:
- Error report downloadable.
- Out-of-scope user rejected.
- Report content correct.

Tests Required:
- Report generation test.
- Permission test.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 3.8 — Sprint 3 Final Verification

## Prompt

```text
You are working on `school-platform`.

Task:
Sprint 3 Final Verification

Target:
school-core-service

Goal:
Verify file management and import readiness.

Scope:
- Run full import flow.
- Verify file privacy.
- Verify signed URL auth.
- Verify import report.
- Produce verification report.

Out of Scope:
- Sprint 4 PPDB.

Rules:
- Do not proceed if import inserts without preview.
- Document missing items.

Acceptance Criteria:
- All Sprint 3 criteria pass or documented.
- No file access leak.

Tests Required:
- File/import test suite.
- Manual import smoke test.

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
docs/32-sprint-3-plan.md if it exists
docs/16-sprint-3-task-prompts.md
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
docs/32-sprint-3-plan.md
docs/40-github-repository-setup-labels-project.md
docs/25-github-project-management.md
docs/27-workflow.md
```

Use `docs/32-sprint-3-plan.md` for sprint-level issue plan, dependencies, risks, and exit criteria.

Use `docs/40-github-repository-setup-labels-project.md` for labels, milestones, project fields, and repository workflow alignment.
