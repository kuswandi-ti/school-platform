# 15 — Sprint 2 Task Prompts

Sprint: Sprint 2 — School Core

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

# Task 2.1 — Create School Core Database Migrations

## Prompt

```text
You are working on `school-platform`.

Task:
Create School Core Database Migrations

Target:
services/school-core-service

Goal:
Create school_core_db schema for foundation, schools, academic period, students, guardians, teachers, classes, and assignments.

Scope:
- Create goose migrations for foundations, schools, academic_years, semesters, students, guardians, student_guardians, teachers, grade_levels, classes, student_class_assignments, teacher_assignments, homeroom_assignments, rooms optional.
- Add indexes and unique constraints.
- Add audit_logs table if local audit needed.

Out of Scope:
- Import Excel.
- PPDB.
- Finance.
- Academic grading.

Rules:
- Use UUID.
- Use foundation_id and school_id.
- No foreign keys to other service databases.
- Use goose.

Acceptance Criteria:
- Migrations run up/down.
- Unique constraints exist for school_code, student_number, class_code per scope.
- Indexes exist for common filters.

Tests Required:
- Migration tests.
- Schema review.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 2.2 — Implement Foundation and School CRUD

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Foundation and School CRUD

Target:
services/school-core-service + api-gateway

Goal:
Implement foundation and school management for Admin Yayasan.

Scope:
- Repository/usecase/API for foundations and schools.
- GET current foundation.
- GET/POST/PATCH schools.
- Scope checks.
- Audit for school changes.
- Publish school.school.created event.

Out of Scope:
- Multi-tenant SaaS billing.
- School deletion complex flow.

Rules:
- Admin Yayasan foundation scope required.
- No direct access from frontend to service.
- Use standard response.

Acceptance Criteria:
- Admin can list schools.
- Admin can create/update school.
- Non-admin forbidden.
- Event/audit recorded.

Tests Required:
- CRUD tests.
- Permission tests.
- Event/audit tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 2.3 — Implement Academic Year and Semester

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Academic Year and Semester

Target:
services/school-core-service

Goal:
Implement global foundation-level academic year and semester management.

Scope:
- CRUD academic_years and semesters.
- Activate academic year/semester.
- Close semester.
- Status enum draft/active/closed/archived.
- Publish events for activated/closed.

Out of Scope:
- Per-school academic year override.
- Calendar event detail.

Rules:
- Academic year/semester global at foundation level.
- Closing should be auditable.
- Only authorized roles can activate/close.

Acceptance Criteria:
- Admin Yayasan can activate academic year.
- Active academic year is unique per foundation.
- Closed semester cannot be edited freely.
- Events published.

Tests Required:
- CRUD/status tests.
- Uniqueness active tests.
- Permission tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 2.4 — Implement Student and Guardian CRUD

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Student and Guardian CRUD

Target:
services/school-core-service + api-gateway

Goal:
Implement student and guardian master data.

Scope:
- CRUD students.
- CRUD guardians.
- Link guardian to student.
- Search/filter students.
- Status active/inactive/transferred/graduated/dropped_out.
- Audit sensitive changes.
- Publish student events.

Out of Scope:
- Import Excel.
- PPDB conversion.
- Finance policy.
- Parent user creation.

Rules:
- School Core owns student/guardian master.
- Free SPP is not student status.
- Scope by foundation_id/school_id.
- Parent access must be scoped to linked student.

Acceptance Criteria:
- TU can create student in school.
- Kepala Sekolah can view school students.
- Cross-school access rejected.
- Guardian linked successfully.
- Event published.

Tests Required:
- Student CRUD tests.
- Guardian link tests.
- Search/filter tests.
- Scope tests.
- Event/audit tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 2.5 — Implement Teacher, Grade Level, and Class CRUD

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Teacher, Grade Level, and Class CRUD

Target:
services/school-core-service

Goal:
Implement teacher master, grade levels, and class/rombel.

Scope:
- CRUD teachers.
- CRUD grade_levels.
- CRUD classes.
- Class unique per school + academic_year + class_code.
- Audit changes.
- Publish teacher/class events.

Out of Scope:
- HR complete.
- Payroll.
- Academic subject management.

Rules:
- Teacher master dasar only.
- Subject belongs to Academic Service.
- No query to academic_db.

Acceptance Criteria:
- TU can create teachers/classes.
- Class code uniqueness enforced.
- School scope enforced.
- Events published.

Tests Required:
- Teacher CRUD tests.
- Class CRUD tests.
- Uniqueness tests.
- Scope tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 2.6 — Implement Student-Class Assignment

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Student-Class Assignment

Target:
services/school-core-service

Goal:
Assign students to classes per academic year/semester.

Scope:
- Create student_class_assignments.
- Move student between classes.
- Status active/moved/completed.
- Prevent duplicate active assignment.
- Publish school.student_class.assigned event.

Out of Scope:
- Promotion automation.
- Report card generation.

Rules:
- Validate student/class same school.
- Actor must have school scope.
- Audit assignment changes.

Acceptance Criteria:
- Student assigned to one active class.
- Duplicate active assignment rejected.
- Cross-school assignment rejected.
- Event/audit recorded.

Tests Required:
- Assignment tests.
- Duplicate tests.
- Scope tests.
- Event/audit tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 2.7 — Implement Teacher and Homeroom Assignment

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Teacher and Homeroom Assignment

Target:
services/school-core-service

Goal:
Assign teachers and homeroom teachers.

Scope:
- Create teacher_assignments.
- Create homeroom_assignments.
- Validate teacher/class same school.
- Assignment status active/inactive.
- Publish homeroom event.

Out of Scope:
- Academic schedule.
- Subject teaching detail beyond reference ID.

Rules:
- Wali Kelas is assignment, not main role.
- Sensitive assignment should be auditable.
- subject_id is reference only.

Acceptance Criteria:
- Teacher can be assigned to class.
- Homeroom teacher can be assigned.
- Duplicate active homeroom for same class rejected.
- Scope enforced.

Tests Required:
- Teacher assignment tests.
- Homeroom uniqueness tests.
- Scope tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 2.8 — Sprint 2 Final Verification

## Prompt

```text
You are working on `school-platform`.

Task:
Sprint 2 Final Verification

Target:
school-core-service

Goal:
Verify School Core sprint readiness.

Scope:
- Run migrations and tests.
- Smoke test core CRUD.
- Verify scope checks.
- Verify event/audit creation.
- Produce verification report.

Out of Scope:
- Sprint 3 work.

Rules:
- Document any missing items.
- Do not proceed if core scope is broken.

Acceptance Criteria:
- All Sprint 2 criteria pass or documented.
- No Critical/High blocker for Import/PPDB/Finance/Academic dependency.

Tests Required:
- Full Sprint 2 test suite.
- Manual API smoke tests.

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
docs/31-sprint-2-plan.md if it exists
docs/15-sprint-2-task-prompts.md
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
docs/31-sprint-2-plan.md
docs/40-github-repository-setup-labels-project.md
docs/25-github-project-management.md
docs/27-workflow.md
```

Use `docs/31-sprint-2-plan.md` for sprint-level issue plan, dependencies, risks, and exit criteria.

Use `docs/40-github-repository-setup-labels-project.md` for labels, milestones, project fields, and repository workflow alignment.
