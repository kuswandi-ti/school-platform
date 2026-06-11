# 14 — Sprint 1 Task Prompts

Sprint: Sprint 1 — Identity & Access

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

# Task 1.1 — Create Identity Service Database Migrations

## Prompt

```text
You are working on `school-platform`.

Task:
Create Identity Service Database Migrations

Target:
services/identity-service

Goal:
Create identity_db schema foundation for users, sessions, roles, permissions, and assignments.

Scope:
- Create goose migrations for users, roles, permissions, role_permissions, user_role_assignments, user_sessions, user_devices optional, identity_audit_logs.
- Use UUID primary keys and timestamp fields.
- Add indexes for email, user_id, role_id, foundation_id, school_id, status.
- Add safe unique constraints.

Out of Scope:
- Login usecase.
- JWT generation.
- API Gateway auth middleware.
- UI.

Rules:
- Use goose only.
- Use snake_case plural table names.
- Do not store refresh token plain text.
- Prepare refresh_token_hash column.
- Do not add business logic in migrations.

Acceptance Criteria:
- Migrations run successfully on identity_db.
- Tables have required indexes and constraints.
- Down migrations exist.
- No production data or secrets.

Tests Required:
- Migration up/down test.
- Schema validation if available.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 1.2 — Implement Password Hashing and User Repository

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Password Hashing and User Repository

Target:
services/identity-service

Goal:
Implement secure password hashing and user repository with sqlc.

Scope:
- Add password hashing package using Argon2id or bcrypt.
- Create sqlc queries for users.
- Implement repository methods: create user, find by email, find by id, update last login, update status.
- Add unit tests.

Out of Scope:
- Login endpoint.
- Refresh token.
- Role management UI.

Rules:
- Never store plain password.
- Never log password.
- Use pgx + sqlc.
- Repository must not contain business authorization logic.

Acceptance Criteria:
- Password hash verifies correctly.
- Wrong password fails.
- User lookup by email works.
- Inactive user status can be read.

Tests Required:
- Hash verify unit tests.
- Repository integration tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 1.3 — Implement Login Usecase and REST Endpoint

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Login Usecase and REST Endpoint

Target:
identity-service + api-gateway

Goal:
Implement email/password login through API Gateway mapping to Identity Service.

Scope:
- Create Login usecase.
- Create gRPC method or internal handler for login.
- Add API Gateway REST endpoint POST /api/v1/auth/login.
- Return access token and refresh token according to API contract.
- Add standard error handling.

Out of Scope:
- Refresh token rotation.
- 2FA.
- OAuth/social login.
- User profile management.

Rules:
- Use standard response/error format.
- Do not expose password hash.
- Rate limit placeholder may be added but full rate limiting can be separate.
- Log with correlation_id only.

Acceptance Criteria:
- Active user can login.
- Wrong password returns standard error.
- Inactive/locked user rejected.
- Response format is standard.

Tests Required:
- Login success API test.
- Wrong password test.
- Inactive user test.
- Response format test.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 1.4 — Implement Refresh Token Rotation

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Refresh Token Rotation

Target:
services/identity-service + api-gateway

Goal:
Implement secure rotating refresh token flow.

Scope:
- Create session repository with refresh_token_hash.
- Implement POST /api/v1/auth/refresh.
- Rotate refresh token on every use.
- Revoke old token.
- Detect reused/revoked refresh token.
- Update last_used_at.

Out of Scope:
- OAuth.
- 2FA.
- Device management UI.

Rules:
- Store refresh token hash only.
- Do not log token.
- Use secure random token generation.
- Return standard errors.

Acceptance Criteria:
- Valid refresh token returns new access and refresh token.
- Old refresh token cannot be reused.
- Revoked session cannot refresh.
- Expired refresh token rejected.

Tests Required:
- Refresh success test.
- Refresh reuse detection test.
- Expired/revoked token tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 1.5 — Implement Logout and Session Revocation

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Logout and Session Revocation

Target:
services/identity-service + api-gateway

Goal:
Allow user logout and session revocation.

Scope:
- Implement POST /api/v1/auth/logout.
- Revoke current session.
- Add repository method for revoke session.
- Add audit log identity.session.revoked if needed.

Out of Scope:
- Logout all devices.
- Advanced session management UI.

Rules:
- Require valid actor context.
- Do not log token.
- Use standard response.

Acceptance Criteria:
- Authenticated user can logout.
- Logged-out refresh token cannot be used.
- Audit log recorded if implemented.

Tests Required:
- Logout API test.
- Refresh after logout rejected test.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 1.6 — Implement Role Permission Seed and Assignment Model

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Role Permission Seed and Assignment Model

Target:
services/identity-service

Goal:
Seed MVP roles and permissions and implement scoped role assignment.

Scope:
- Seed roles: admin_yayasan, kepala_sekolah, tu_staff, bendahara_sekolah, guru, orang_tua, siswa.
- Seed permission code format domain.resource.action.
- Implement user_role_assignments with foundation/school/class/student/subject scopes.
- Add queries to get user context.

Out of Scope:
- Role management UI.
- Advanced ABAC engine.

Rules:
- Role assignment is source of authz context.
- Wali Kelas is assignment, not main role.
- Aksi sensitif later requires approval/audit.

Acceptance Criteria:
- Seed can run idempotently.
- User context returns roles, permissions, and scope.
- School-scoped role assignment works.

Tests Required:
- Seed idempotency test.
- Get user context test.
- Scope assignment test.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 1.7 — Implement API Gateway Auth Middleware

## Prompt

```text
You are working on `school-platform`.

Task:
Implement API Gateway Auth Middleware

Target:
services/api-gateway

Goal:
Protect routes and extract actor context from access token.

Scope:
- Add JWT validation middleware.
- Extract user_id, foundation_id, school_id, roles, permissions, scope.
- Add actor context to request.
- Return standard UNAUTHORIZED/FORBIDDEN errors.
- Propagate request_id/correlation_id.

Out of Scope:
- Detailed module authorization.
- Refresh token cookie handling details.
- Business logic.

Rules:
- API Gateway performs basic guard only.
- Services still must perform authorization.
- No business logic in middleware.
- No token in logs.

Acceptance Criteria:
- Protected endpoint rejects missing token.
- Invalid token rejected.
- Valid token allows request.
- Actor context available to handlers.

Tests Required:
- Middleware tests.
- Missing/invalid/valid token tests.
- No token in log review.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 1.8 — Sprint 1 Final Verification

## Prompt

```text
You are working on `school-platform`.

Task:
Sprint 1 Final Verification

Target:
identity-service + api-gateway

Goal:
Verify Identity & Access sprint readiness.

Scope:
- Run migrations.
- Run tests.
- Verify login/refresh/logout/me flow.
- Verify role/scope context.
- Produce verification report.

Out of Scope:
- New features beyond Sprint 1.

Rules:
- Do not hide failing tests.
- List missing items honestly.

Acceptance Criteria:
- All Sprint 1 acceptance criteria pass or issues documented.
- No Critical/High blocker before Sprint 2.

Tests Required:
- Full Sprint 1 test suite.
- Manual API smoke test.

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
docs/30-sprint-1-plan.md if it exists
docs/14-sprint-1-task-prompts.md
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
docs/30-sprint-1-plan.md
docs/40-github-repository-setup-labels-project.md
docs/25-github-project-management.md
docs/27-workflow.md
```

Use `docs/30-sprint-1-plan.md` for sprint-level issue plan, dependencies, risks, and exit criteria.

Use `docs/40-github-repository-setup-labels-project.md` for labels, milestones, project fields, and repository workflow alignment.
