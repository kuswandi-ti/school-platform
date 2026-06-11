# 22 — Sprint 9 Task Prompts

Sprint: Sprint 9 — Reporting Dashboard

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

# Task 9.1 — Create Reporting Database Migrations

## Prompt

```text
You are working on `school-platform`.

Task:
Create Reporting Database Migrations

Target:
reporting-service

Goal:
Create reporting_db projection tables and processed event tracking.

Scope:
- Create foundation_dashboard_metrics, school_dashboard_metrics, student_summary_metrics, teacher_summary_metrics, admission_summary_metrics, finance_summary_metrics, attendance_summary_metrics, academic_progress_metrics, approval_pending_metrics, notification_summary_metrics, processed_events, reporting_projection_offsets.

Out of Scope:
- Operational source-of-truth tables.
- BI warehouse.

Rules:
- Reporting is projection only.
- No direct operational DB query.
- Idempotency table required.

Acceptance Criteria:
- Migrations run.
- Projection tables created.
- processed_events exists.

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

# Task 9.2 — Implement Event Consumer Infrastructure

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Event Consumer Infrastructure

Target:
reporting-service

Goal:
Build RabbitMQ consumer skeleton with idempotency.

Scope:
- RabbitMQ subscriber.
- Event envelope validation.
- processed_events check.
- Transaction per event.
- Retry/DLQ ready.
- slog with correlation_id.

Out of Scope:
- All projections.
- Rebuild job.

Rules:
- Idempotent consumer.
- No sensitive logs.
- No operational DB access.

Acceptance Criteria:
- Event consumed.
- Duplicate skipped.
- Invalid event handled safely.

Tests Required:
- Consumer tests.
- Duplicate event tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 9.3 — Implement Student and Teacher Summary Projection

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Student and Teacher Summary Projection

Target:
reporting-service

Goal:
Update student/teacher summary from School Core events.

Scope:
- Handle school.student.created/status_changed, school.teacher.created/updated.
- Update summary counts by foundation/school/status.
- Idempotent.

Out of Scope:
- Detailed student list.

Rules:
- Projection from events only.
- Safe payload usage.

Acceptance Criteria:
- Student count updates.
- Teacher count updates.
- Duplicate event skipped.

Tests Required:
- Projection tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 9.4 — Implement Admission Summary Projection

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Admission Summary Projection

Target:
reporting-service

Goal:
Update PPDB summary metrics from Admission events.

Scope:
- Handle applicant.submitted/verified/accepted/rejected/converted.
- Aggregate by school/period/status.

Out of Scope:
- Admission operational list.

Rules:
- Event-driven only.
- Idempotent.

Acceptance Criteria:
- PPDB metrics update.
- Duplicate skipped.

Tests Required:
- Projection tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 9.5 — Implement Finance Summary Projection

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Finance Summary Projection

Target:
reporting-service

Goal:
Update finance metrics from Finance events.

Scope:
- Handle bill.generated, payment.verified, payment.rejected, payment.voided, fee_policy.approved.
- Update total_billed, total_paid, outstanding, collection_rate.

Out of Scope:
- Accounting ledger.
- Direct finance_db query.

Rules:
- Use decimal safe handling.
- Projection only.
- Idempotent.

Acceptance Criteria:
- Finance metrics update correctly.
- Duplicate skipped.
- Void recalculates projection.

Tests Required:
- Finance projection tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 9.6 — Implement Attendance and Academic Progress Projection

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Attendance and Academic Progress Projection

Target:
reporting-service

Goal:
Update attendance and report card progress summaries.

Scope:
- Handle attendance.marked, report_card.published, grade_book.submitted/approved.
- Aggregate per school/class/date/semester.

Out of Scope:
- Detailed grade report.

Rules:
- Projection only.
- No Confidential detail.

Acceptance Criteria:
- Attendance counts update.
- Academic progress updates.
- Duplicate skipped.

Tests Required:
- Projection tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 9.7 — Implement Dashboard APIs

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Dashboard APIs

Target:
reporting-service + api-gateway

Goal:
Expose foundation, school, teacher, and parent dashboard endpoints.

Scope:
- GET /api/v1/dashboard/foundation.
- GET /api/v1/dashboard/school.
- GET /api/v1/dashboard/teacher.
- GET /api/v1/dashboard/parent.
- Scope checks.

Out of Scope:
- Advanced analytics.
- Charts rendering.

Rules:
- Dashboard reads reporting_db only.
- Role/scope enforced.
- Standard response.

Acceptance Criteria:
- Admin sees foundation metrics.
- Kepala Sekolah sees school metrics.
- Parent sees child summary only.
- Unauthorized rejected.

Tests Required:
- API tests.
- Scope tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 9.8 — Implement Scheduled Rebuild Skeleton

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Scheduled Rebuild Skeleton

Target:
reporting-service

Goal:
Add controlled rebuild/sync skeleton for projection repair.

Scope:
- Command/job skeleton.
- Rebuild interfaces.
- Logging.
- Safety notes.
- Initial rebuild for selected metrics if practical.

Out of Scope:
- Full data warehouse rebuild.
- Direct DB access to all operational DBs.

Rules:
- Do not query operational DB directly unless controlled approved gRPC endpoint exists.
- Document limitations.

Acceptance Criteria:
- Rebuild command exists.
- Can be run safely.
- Logs correlation/job id.

Tests Required:
- Command tests if practical.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 9.9 — Sprint 9 Final Verification

## Prompt

```text
You are working on `school-platform`.

Task:
Sprint 9 Final Verification

Target:
reporting-service

Goal:
Verify dashboard and projection readiness.

Scope:
- Run event simulation.
- Verify dashboards.
- Verify idempotency.
- Produce report.

Out of Scope:
- Advanced BI.

Rules:
- No operational DB direct query.
- Scope enforced.

Acceptance Criteria:
- Dashboard works near real-time.
- No Critical/High bug.

Tests Required:
- Reporting test suite.
- Manual dashboard smoke test.

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
docs/38-sprint-9-plan.md if it exists
docs/22-sprint-9-task-prompts.md
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
docs/38-sprint-9-plan.md
docs/40-github-repository-setup-labels-project.md
docs/25-github-project-management.md
docs/27-workflow.md
```

Use `docs/38-sprint-9-plan.md` for sprint-level issue plan, dependencies, risks, and exit criteria.

Use `docs/40-github-repository-setup-labels-project.md` for labels, milestones, project fields, and repository workflow alignment.
