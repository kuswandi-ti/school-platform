# 23 — Sprint 10 Task Prompts

Sprint: Sprint 10 — Security, Observability, Backup, and UAT Hardening

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

# Task 10.1 — Run Security Baseline Review

## Prompt

```text
You are working on `school-platform`.

Task:
Run Security Baseline Review

Target:
all services

Goal:
Review and fix MVP security baseline gaps.

Scope:
- Review auth, authorization, object-level checks, rate limiting, CORS, headers, input validation, secrets, private files.
- Create checklist results.

Out of Scope:
- Full penetration test.
- WAF setup unless needed.

Rules:
- No sensitive logs.
- No secrets committed.
- No public private files.

Acceptance Criteria:
- Security checklist completed.
- Critical gaps fixed or documented.
- No Critical/High open.

Tests Required:
- Security tests.
- Manual checklist.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 10.2 — Add Permission and Scope Regression Tests

## Prompt

```text
You are working on `school-platform`.

Task:
Add Permission and Scope Regression Tests

Target:
all services

Goal:
Add/complete cross-role and cross-school permission tests.

Scope:
- Admin/Kepsek/TU/Bendahara/Guru/Parent/Student cases.
- Cross-school denial.
- Object-level authorization.

Out of Scope:
- New features.

Rules:
- Permission/scope tests are priority.
- Use realistic test fixtures.

Acceptance Criteria:
- Critical access paths covered.
- Cross-school leak tests pass.

Tests Required:
- Regression permission tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 10.3 — Review Audit Log Coverage

## Prompt

```text
You are working on `school-platform`.

Task:
Review Audit Log Coverage

Target:
all services

Goal:
Ensure sensitive actions create audit logs.

Scope:
- Review role changes, student update, PPDB decision, payment verify/void, fee policy, report publish/revision, file download.
- Add missing audit tests.

Out of Scope:
- Central Audit Service.

Rules:
- Audit separate from app logs.
- Mask sensitive values.

Acceptance Criteria:
- Sensitive actions audited.
- Audit includes correlation_id.
- Masking verified.

Tests Required:
- Audit tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 10.4 — Implement Metrics and Health Readiness Review

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Metrics and Health Readiness Review

Target:
all services + infra

Goal:
Ensure health/readiness/metrics are production-ready enough for MVP.

Scope:
- Verify /healthz, /readyz, /metrics.
- Prometheus scrape config.
- Basic service metrics.
- RabbitMQ/DB readiness checks.

Out of Scope:
- Advanced tracing.

Rules:
- No sensitive metric labels.
- Readiness reflects dependencies.

Acceptance Criteria:
- All services expose health/readiness.
- Metrics available.
- Prometheus can scrape.

Tests Required:
- Health/ready tests.
- Metrics smoke test.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 10.5 — Set Up Logging Dashboard Skeleton

## Prompt

```text
You are working on `school-platform`.

Task:
Set Up Logging Dashboard Skeleton

Target:
infra/observability

Goal:
Prepare Loki/Grafana logging view for MVP.

Scope:
- Docker Compose/staging config for Loki/Grafana.
- Basic dashboard docs.
- Correlation ID search guide.

Out of Scope:
- Full SIEM.
- Advanced alert rules.

Rules:
- No sensitive logs.
- Logs include service/env/correlation_id.

Acceptance Criteria:
- Logs visible in Grafana/Loki.
- Can search by correlation_id.

Tests Required:
- Manual observability test.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 10.6 — Create Backup Script and Schedule

## Prompt

```text
You are working on `school-platform`.

Task:
Create Backup Script and Schedule

Target:
infra/deploy

Goal:
Implement database/object storage backup scripts for MVP.

Scope:
- PostgreSQL backup per service database.
- MinIO backup approach/documentation.
- Encrypted backup note.
- Retention config.
- Failure logging.

Out of Scope:
- Cloud provider-specific full DR automation.

Rules:
- Backup treated as Confidential.
- No backup in Git.
- Secrets from env.

Acceptance Criteria:
- Backup script runs.
- Outputs stored outside repo.
- Failure visible.

Tests Required:
- Backup dry-run/manual test.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 10.7 — Perform Restore Test and Document Procedure

## Prompt

```text
You are working on `school-platform`.

Task:
Perform Restore Test and Document Procedure

Target:
infra/docs

Goal:
Test and document restore process.

Scope:
- Restore test for at least one database.
- Document restore steps.
- Document RPO/RTO target.
- Verification checklist.

Out of Scope:
- Automated full DR drill.

Rules:
- Use non-production/safe test data.
- Do not overwrite production.

Acceptance Criteria:
- Restore test succeeds.
- Procedure documented.
- Verification steps clear.

Tests Required:
- Restore manual test.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 10.8 — Create UAT and Regression Checklist

## Prompt

```text
You are working on `school-platform`.

Task:
Create UAT and Regression Checklist

Target:
docs/uat

Goal:
Prepare MVP UAT checklist for school pilot.

Scope:
- Auth flow.
- School Core.
- Import.
- PPDB.
- Finance.
- Academic.
- Report card.
- Notification.
- Dashboard.
- Audit checks.

Out of Scope:
- Automated E2E framework full build.

Rules:
- Checklist role-based.
- Critical/High bug policy included.

Acceptance Criteria:
- UAT checklist created.
- Regression checklist created.
- QA sign-off section exists.

Tests Required:
- Document review.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 10.9 — Prepare Production Deployment Readiness

## Prompt

```text
You are working on `school-platform`.

Task:
Prepare Production Deployment Readiness

Target:
deploy/production + .github

Goal:
Finalize production deploy readiness with manual approval.

Scope:
- Review GitHub Environment production.
- Secrets checklist.
- Deploy workflow review.
- Rollback guide.
- Release notes template.
- Image tag strategy.

Out of Scope:
- Actual production server provisioning if not available.

Rules:
- Production deploy from main only.
- Manual approval required.
- No secrets in repo.

Acceptance Criteria:
- Readiness checklist complete.
- Rollback documented.
- Release tag policy clear.

Tests Required:
- Workflow dry-run if possible.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 10.10 — Sprint 10 Final Verification

## Prompt

```text
You are working on `school-platform`.

Task:
Sprint 10 Final Verification

Target:
all services

Goal:
Perform final MVP hardening verification.

Scope:
- Run regression.
- Run security/scope tests.
- Run backup/restore test.
- Check observability.
- Review open bugs.
- Produce final readiness report.

Out of Scope:
- New feature development.

Rules:
- No Critical/High bugs for release.
- Be honest about gaps.

Acceptance Criteria:
- MVP ready for pilot or gaps documented.
- QA sign-off possible.

Tests Required:
- Full regression.
- UAT checklist.
- Security checklist.

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
docs/39-sprint-10-plan.md if it exists
docs/23-sprint-10-task-prompts.md
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
docs/39-sprint-10-plan.md
docs/40-github-repository-setup-labels-project.md
docs/25-github-project-management.md
docs/27-workflow.md
```

Use `docs/39-sprint-10-plan.md` for sprint-level issue plan, dependencies, risks, and exit criteria.

Use `docs/40-github-repository-setup-labels-project.md` for labels, milestones, project fields, and repository workflow alignment.
