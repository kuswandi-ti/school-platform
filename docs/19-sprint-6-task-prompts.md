# 19 — Sprint 6 Task Prompts

Sprint: Sprint 6 — Academic Dasar

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

# Task 6.1 — Create Academic Basic Migrations

## Prompt

```text
You are working on `school-platform`.

Task:
Create Academic Basic Migrations

Target:
academic-service

Goal:
Create academic_db schema for curriculum, subjects, schedules, and attendance.

Scope:
- Create curriculums, learning_phases, subjects, subject_groups, class_subjects, schedules, student_attendances, academic_audit_logs.
- Add indexes for school/year/semester/class/teacher/date.

Out of Scope:
- Grades/report cards.
- LMS.

Rules:
- Reference School Core IDs only.
- No FK to school_core_db.
- Use school scope.

Acceptance Criteria:
- Migrations run.
- Indexes exist.
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

# Task 6.2 — Implement Curriculum and Subject Management

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Curriculum and Subject Management

Target:
academic-service

Goal:
Manage curriculum baseline, subjects, and subject groups.

Scope:
- CRUD curriculums.
- CRUD subjects.
- CRUD subject_groups.
- school_level.
- Audit.
- Publish subject.created.

Out of Scope:
- Assessment scheme.
- Report card.

Rules:
- Academic owns subject.
- School scope enforced.
- No UI complexity.

Acceptance Criteria:
- Subject created.
- Duplicate code rejected.
- Scope enforced.
- Event/audit recorded.

Tests Required:
- CRUD tests.
- Duplicate tests.
- Scope tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 6.3 — Implement Class Subject Assignment

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Class Subject Assignment

Target:
academic-service

Goal:
Assign subjects and teachers to classes.

Scope:
- class_subjects.
- teacher_id/class_id reference.
- Validate reference via gRPC/projection if available.
- Status active/inactive.
- Audit.

Out of Scope:
- Schedule.
- Grade book.

Rules:
- No query school_core_db.
- Teacher/class assignment scope enforced.

Acceptance Criteria:
- Class subject created.
- Invalid scope rejected.
- Teacher assignment respected.

Tests Required:
- Assignment tests.
- Scope tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 6.4 — Implement Schedule Management

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Schedule Management

Target:
academic-service

Goal:
Create and manage class schedules.

Scope:
- schedules CRUD.
- day_of_week/start_time/end_time.
- teacher/class/subject/room references.
- Conflict check basic.
- Publish schedule.created.

Out of Scope:
- Advanced optimization.
- Room booking full module.

Rules:
- School scope.
- Basic conflict prevention.
- No direct DB query to School Core.

Acceptance Criteria:
- Schedule created.
- Invalid time rejected.
- Basic conflict rejected.
- Event published.

Tests Required:
- Schedule tests.
- Conflict tests.
- Scope tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 6.5 — Implement Teacher Schedule View

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Teacher Schedule View

Target:
academic-service + api-gateway

Goal:
Allow teacher to view assigned schedule.

Scope:
- GET teacher schedule endpoint.
- Filter today/week.
- Use actor teacher assignment/scope.
- Return class/subject references/snapshots.

Out of Scope:
- Mobile full UI.
- Offline cache.

Rules:
- Teacher only sees assigned schedule.
- Other teacher schedule forbidden.

Acceptance Criteria:
- Teacher sees own schedule.
- Unauthorized schedule access rejected.

Tests Required:
- Schedule view tests.
- Scope tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 6.6 — Implement Attendance Input

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Attendance Input

Target:
academic-service

Goal:
Allow teachers to mark student attendance.

Scope:
- Create attendance records.
- Statuses present/sick/excused/absent/late.
- Validate teacher assignment.
- Prevent duplicate per date/class/student/subject.
- Publish attendance.marked and absent event if needed.

Out of Scope:
- Report card.
- Discipline module.

Rules:
- Guru scope required.
- Parent cannot mark attendance.
- Audit correction later.

Acceptance Criteria:
- Teacher marks attendance.
- Unassigned teacher rejected.
- Duplicate rejected/updated by rule.
- Event published.

Tests Required:
- Attendance tests.
- Teacher scope tests.
- Duplicate tests.
- Event tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 6.7 — Implement Attendance Correction

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Attendance Correction

Target:
academic-service

Goal:
Allow correction of attendance with audit and optional approval.

Scope:
- Correction usecase.
- Reason field.
- Audit old/new values.
- Publish attendance.corrected.
- Permission check.

Out of Scope:
- Complex approval multi-level.

Rules:
- Reason required.
- Sensitive correction audited.
- Scope enforced.

Acceptance Criteria:
- Attendance corrected.
- Reason required.
- Audit/event recorded.

Tests Required:
- Correction tests.
- Reason tests.
- Audit/event tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 6.8 — Sprint 6 Final Verification

## Prompt

```text
You are working on `school-platform`.

Task:
Sprint 6 Final Verification

Target:
academic-service

Goal:
Verify Academic Basic sprint readiness.

Scope:
- Run curriculum/subject/schedule/attendance flow.
- Verify teacher scope.
- Verify events.
- Produce report.

Out of Scope:
- Report card.

Rules:
- No unassigned teacher access.
- No direct School Core DB query.

Acceptance Criteria:
- Academic basic works end-to-end.
- No Critical/High bug.

Tests Required:
- Academic basic test suite.
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
docs/35-sprint-6-plan.md if it exists
docs/19-sprint-6-task-prompts.md
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
docs/35-sprint-6-plan.md
docs/40-github-repository-setup-labels-project.md
docs/25-github-project-management.md
docs/27-workflow.md
```

Use `docs/35-sprint-6-plan.md` for sprint-level issue plan, dependencies, risks, and exit criteria.

Use `docs/40-github-repository-setup-labels-project.md` for labels, milestones, project fields, and repository workflow alignment.
