# 21 — Sprint 8 Task Prompts

Sprint: Sprint 8 — Communication / Notification

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

# Task 8.1 — Create Communication Database Migrations

## Prompt

```text
You are working on `school-platform`.

Task:
Create Communication Database Migrations

Target:
communication-service

Goal:
Create communication_db schema for announcements, notifications, templates, deliveries, preferences.

Scope:
- Create announcements, announcement_targets, notifications, notification_templates, notification_deliveries, notification_preferences, communication_audit_logs.
- Add indexes.

Out of Scope:
- WhatsApp.
- SMS.
- Marketing campaigns.

Rules:
- No Confidential detail in notifications.
- Use user_id references.
- No direct business DB queries.

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

# Task 8.2 — Implement Announcement CRUD and Publish

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Announcement CRUD and Publish

Target:
communication-service

Goal:
Allow authorized users to create and publish announcements.

Scope:
- CRUD announcements.
- Publish announcement.
- Status draft/submitted/published/rejected/archived.
- Target scope.
- Audit.
- Publish announcement.published event.

Out of Scope:
- Complex approval if not required.
- Rich media.

Rules:
- Scope by foundation/school.
- Authorized roles only.
- Confidential data not in body.

Acceptance Criteria:
- Announcement created/published.
- Unauthorized rejected.
- Target scope stored.
- Event/audit recorded.

Tests Required:
- CRUD tests.
- Publish tests.
- Scope tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 8.3 — Implement Announcement Targets

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Announcement Targets

Target:
communication-service

Goal:
Target announcements to foundation/school/class/role/user/student/parent/teacher.

Scope:
- announcement_targets.
- Resolve recipients through Identity/School Core gRPC or read model.
- Target validation.
- Recipient creation for notification.

Out of Scope:
- Global search target.
- Complex segmentation.

Rules:
- Do not query other DBs.
- Use gRPC/events/read model.
- Respect scope.

Acceptance Criteria:
- Targets stored.
- Recipients resolved.
- Invalid target rejected.

Tests Required:
- Target tests.
- Recipient resolution tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 8.4 — Implement Notification Template and In-App Notification

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Notification Template and In-App Notification

Target:
communication-service

Goal:
Create notification templates and in-app notification records.

Scope:
- notification_templates CRUD.
- notifications table.
- Create notification from event/template.
- Read/unread endpoint.
- Preference check basic.

Out of Scope:
- FCM actual provider.
- Email delivery.

Rules:
- Critical notifications cannot be fully disabled.
- No Confidential detail.

Acceptance Criteria:
- Notification created.
- User can mark read.
- Template applied.
- Preference respected.

Tests Required:
- Template tests.
- Notification tests.
- Read/unread tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 8.5 — Implement Event Consumers for Notifications

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Event Consumers for Notifications

Target:
communication-service

Goal:
Consume key domain events and create notifications.

Scope:
- Consume finance.bill.generated, finance.payment.verified/rejected, academic.report_card.published, admission.applicant.accepted/rejected, approval.request.created.
- Idempotent processed_events.
- Retry/DLQ ready.

Out of Scope:
- All possible events.
- Advanced routing UI.

Rules:
- Consumers idempotent.
- No duplicate notifications.
- No sensitive payload leakage.

Acceptance Criteria:
- Events create correct notifications.
- Duplicate event skipped.
- Failed event can retry/DLQ.

Tests Required:
- Consumer tests.
- Idempotency tests.
- Payload safety tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 8.6 — Implement Delivery Log and Provider Abstractions

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Delivery Log and Provider Abstractions

Target:
communication-service

Goal:
Prepare FCM/email provider abstraction with mock/local behavior.

Scope:
- notification_deliveries.
- FCM mock provider.
- Email mock/Mailpit adapter.
- Delivery status pending/sent/failed/retrying.
- Retry metadata.

Out of Scope:
- Production FCM setup.
- WhatsApp.

Rules:
- Do not send real external messages in local.
- Log delivery safely.
- Provider errors handled.

Acceptance Criteria:
- Delivery log created.
- Mock provider records sent.
- Failure recorded.

Tests Required:
- Provider mock tests.
- Delivery status tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 8.7 — Implement Notification Preferences

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Notification Preferences

Target:
communication-service

Goal:
Allow users to manage non-critical notification preferences.

Scope:
- notification_preferences.
- Enable/disable category/channel.
- Critical notification protection.
- Preference check in delivery.

Out of Scope:
- Complex preference UI.

Rules:
- Critical/security/finance important cannot be fully disabled.
- User can manage own preferences only.

Acceptance Criteria:
- Preference saved.
- Delivery respects preference.
- Critical still delivered.

Tests Required:
- Preference tests.
- Critical override tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 8.8 — Sprint 8 Final Verification

## Prompt

```text
You are working on `school-platform`.

Task:
Sprint 8 Final Verification

Target:
communication-service

Goal:
Verify announcement and notification flow.

Scope:
- Run announcement publish.
- Run event-to-notification.
- Verify in-app read.
- Verify delivery log.
- Produce report.

Out of Scope:
- WhatsApp/SMS.

Rules:
- No Confidential detail in notifications.
- No duplicate notifications.

Acceptance Criteria:
- Communication flow works.
- No Critical/High bug.

Tests Required:
- Communication test suite.
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
docs/37-sprint-8-plan.md if it exists
docs/21-sprint-8-task-prompts.md
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
docs/37-sprint-8-plan.md
docs/40-github-repository-setup-labels-project.md
docs/25-github-project-management.md
docs/27-workflow.md
```

Use `docs/37-sprint-8-plan.md` for sprint-level issue plan, dependencies, risks, and exit criteria.

Use `docs/40-github-repository-setup-labels-project.md` for labels, milestones, project fields, and repository workflow alignment.
