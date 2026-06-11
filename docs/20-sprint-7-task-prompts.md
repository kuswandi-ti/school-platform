# 20 — Sprint 7 Task Prompts

Sprint: Sprint 7 — Report Card / E-Rapor Dasar

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

# Task 7.1 — Create Report Card Database Migrations

## Prompt

```text
You are working on `school-platform`.

Task:
Create Report Card Database Migrations

Target:
academic-service

Goal:
Create academic schema for assessment, grade book, scores, report cards.

Scope:
- Create assessment_components, assessment_schemes, grade_books, student_scores, report_templates, report_cards, report_card_items, academic_approval_requests.
- Add indexes for student/class/subject/year/semester/status.

Out of Scope:
- Full LMS.
- Advanced analytics.

Rules:
- Academic owns grades/report cards.
- No FK to School Core DB.
- Report card snapshot required.

Acceptance Criteria:
- Migrations run.
- Status enums supported.
- Indexes exist.

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

# Task 7.2 — Implement Assessment Components and Schemes

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Assessment Components and Schemes

Target:
academic-service

Goal:
Define score components and weighting per subject/class.

Scope:
- CRUD assessment_components.
- CRUD assessment_schemes.
- weight_percentage validation.
- Audit.

Out of Scope:
- Complex curriculum rule engine.

Rules:
- Use decimal if needed.
- No float if precision required.
- School scope enforced.

Acceptance Criteria:
- Components created.
- Scheme weights valid.
- Invalid weight rejected.

Tests Required:
- CRUD tests.
- Validation tests.
- Scope tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 7.3 — Implement Grade Book

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Grade Book

Target:
academic-service

Goal:
Create and manage grade books for class/subject/teacher.

Scope:
- Create grade_books.
- Status draft/submitted/approved/published/locked.
- Teacher assignment validation.
- Audit.
- Publish grade_book.created.

Out of Scope:
- Report generation.
- PDF.

Rules:
- Teacher can only manage assigned grade book.
- Status transitions controlled.

Acceptance Criteria:
- Grade book created.
- Unassigned teacher rejected.
- Status valid.

Tests Required:
- Grade book tests.
- Teacher scope tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 7.4 — Implement Score Input

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Score Input

Target:
academic-service

Goal:
Allow teacher to input student scores and descriptions.

Scope:
- student_scores CRUD within grade book.
- Validate score range.
- Teacher assignment check.
- Prevent edit after locked/published.
- Audit if needed.

Out of Scope:
- Publish report card.

Rules:
- Unassigned teacher cannot input.
- Published/locked cannot edit.
- Score validation.

Acceptance Criteria:
- Scores saved.
- Invalid score rejected.
- Locked edit rejected.

Tests Required:
- Score input tests.
- Scope tests.
- Locked tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 7.5 — Implement Grade Book Submit and Review

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Grade Book Submit and Review

Target:
academic-service

Goal:
Submit grade book and allow review flow.

Scope:
- Submit grade book.
- Wali Kelas/Kepala Sekolah review as applicable.
- Status submitted/approved/revision_requested.
- Reason for revision.
- Audit.
- Publish submitted/approved events.

Out of Scope:
- Report card publish.

Rules:
- Submit requires complete/valid data if rules require.
- Reviewer scope enforced.

Acceptance Criteria:
- Grade book submitted.
- Reviewer can approve/request revision.
- Unauthorized rejected.

Tests Required:
- Submit tests.
- Review tests.
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

# Task 7.6 — Implement Report Card Generation

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Report Card Generation

Target:
academic-service

Goal:
Generate report cards from grade book data.

Scope:
- Create report_cards.
- Create report_card_items.
- Store student_snapshot_json.
- Use report_template.
- Status draft/reviewed.
- Publish report_card.generated.

Out of Scope:
- Final PDF design.
- National integration.

Rules:
- Report card stores snapshot.
- Only valid grade data used.
- No direct School Core DB query.

Acceptance Criteria:
- Report card generated.
- Items created.
- Snapshot stored.
- Event published.

Tests Required:
- Generation tests.
- Snapshot tests.
- Event tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 7.7 — Implement Publish and Lock Report Card

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Publish and Lock Report Card

Target:
academic-service

Goal:
Allow Kepala Sekolah to publish report cards and lock them.

Scope:
- Publish endpoint.
- Set published_at/by.
- Set locked status.
- Prevent edit after publish.
- Publish report_card.published.
- Audit.

Out of Scope:
- Revision flow can be next task.
- Advanced PDF.

Rules:
- Kepala Sekolah scope required.
- Parent/student only see published.
- Audit required.

Acceptance Criteria:
- Report card published.
- Locked after publish.
- Parent can view.
- Unpublished not visible to parent.

Tests Required:
- Publish tests.
- Lock tests.
- Parent visibility tests.
- Audit/event tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 7.8 — Implement Report Card Revision Approval

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Report Card Revision Approval

Target:
academic-service

Goal:
Allow revision request after report card publication.

Scope:
- Create revision request.
- Approval flow.
- Unlock limited edit after approval.
- Re-lock after revised.
- Audit old/new values.
- Publish revision events.

Out of Scope:
- Multi-level complex approval.

Rules:
- Revision after publish always approval.
- Reason required.
- Parent sees latest published/revised according to rule.

Acceptance Criteria:
- Revision request created.
- Approval enables revision.
- Revised report locked.
- Audit/event recorded.

Tests Required:
- Revision tests.
- Approval tests.
- Audit/event tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 7.9 — Sprint 7 Final Verification

## Prompt

```text
You are working on `school-platform`.

Task:
Sprint 7 Final Verification

Target:
academic-service

Goal:
Verify grade and report card flow.

Scope:
- Run score → submit → review → generate → publish → parent view → revision flow.
- Verify teacher scope.
- Produce report.

Out of Scope:
- Full LMS.

Rules:
- No edits after publish without approval.
- No unpublished parent access.

Acceptance Criteria:
- E-Rapor dasar works end-to-end.
- No Critical/High bug.

Tests Required:
- Report card test suite.
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
docs/36-sprint-7-plan.md if it exists
docs/20-sprint-7-task-prompts.md
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
docs/36-sprint-7-plan.md
docs/40-github-repository-setup-labels-project.md
docs/25-github-project-management.md
docs/27-workflow.md
```

Use `docs/36-sprint-7-plan.md` for sprint-level issue plan, dependencies, risks, and exit criteria.

Use `docs/40-github-repository-setup-labels-project.md` for labels, milestones, project fields, and repository workflow alignment.
