# Sprint 1 Plan — Identity & Access

Project: `school-platform`  
Sprint: Sprint 1 — Identity & Access  
Target Output: `docs/30-sprint-1-plan.md`  
Document Type: Sprint Planning Document  
Audience: Product Owner, Developer, QA, DevOps, AI Agent  
Status: Draft for Sprint Execution  

---

## 1. Sprint Summary

Sprint 1 membangun fondasi autentikasi dan otorisasi: users, password hashing, JWT access token, rotating refresh token, logout/session revocation, roles, permissions, actor context, dan API Gateway auth middleware.

Dokumen ini disusun berdasarkan prompt pack Sprint Planning `school-platform` dan harus dibaca bersama:

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/25-product-requirement-document.md
docs/26-development-plan.md
docs/27-workflow.md
docs/25-github-project-management.md
docs/10-sprint-backlog-mvp.md
docs/13-sprint-0-task-prompts.md sampai docs/23-sprint-10-task-prompts.md
```


---

## 2. Sprint Objective

Membangun fondasi autentikasi dan otorisasi agar semua modul MVP berikutnya dapat menerapkan RBAC, ABAC/scope, dan service-side authorization secara konsisten.

---

## 3. Business Context

Yayasan dan sekolah membutuhkan akses sistem yang aman untuk Admin Yayasan, Kepala Sekolah, TU/Staff, Bendahara, Guru, Orang Tua, dan Siswa. Sprint ini memastikan hanya user yang sah dapat masuk dan setiap user membawa konteks akses yang dapat dipakai modul berikutnya.

---

## 4. Technical Context

Identity Service menjadi owner data user, role, permission, session, dan refresh token. API Gateway menangani validasi token dan propagasi actor context, tetapi business authorization tetap berada di service masing-masing.

Key technical constraints:

```text
- Backend menggunakan Go microservices.
- API Gateway custom menggunakan Go + Chi.
- Komunikasi internal antar-service menggunakan gRPC + protobuf.
- Database menggunakan PostgreSQL, satu database per service.
- Event menggunakan RabbitMQ topic exchange domain.events.
- Tidak boleh ada cross-service database query.
- API Gateway tidak boleh berisi business logic.
- Reporting hanya menggunakan reporting_db/read model.
- File private by default.
- Internal code menggunakan English.
- UI label menggunakan Bahasa Indonesia.
```

---

## 5. Scope

- identity_db migrations
- users
- roles
- permissions
- role_permissions
- user_role_assignments
- user_sessions
- password hashing
- login
- refresh token rotation
- logout/revoke session
- role/permission seed
- actor context
- API Gateway auth middleware
- `/api/v1/auth/login`
- `/api/v1/auth/refresh`
- `/api/v1/auth/logout`
- `/api/v1/me`
- `/api/v1/me/permissions`
- `/api/v1/me/context`

---

## 6. Out of Scope

- OAuth/social login
- 2FA
- advanced anomaly detection
- full user profile management
- UI lengkap role management

---

## 7. Target Users / Actors

- Admin Yayasan
- Kepala Sekolah
- TU/Staff
- Bendahara
- Guru
- Wali Kelas
- Orang Tua/Wali Murid
- Siswa
- API Gateway
- Identity Service

---

## 8. User Stories

- As a Semua user, I want login menggunakan credential yang valid, so that dapat mengakses sistem sesuai role.
- As a Admin Yayasan, I want memiliki role dan permission dasar, so that dapat mengelola akses awal sistem.
- As a API Gateway, I want memvalidasi access token, so that request ke service membawa actor context.
- As a Backend Service, I want menerima actor context, so that dapat melakukan service-side authorization.
- As a QA, I want menguji refresh token rotation, so that risiko token reuse dapat dicegah.

---

## 9. Functional Breakdown

- Login dengan email/password.
- Access token short-lived.
- Refresh token rotation dan revoke.
- Endpoint current user dan permissions.
- Seed role/permission baseline.
- Actor context propagation.
- Auth middleware di API Gateway.

---

## 10. Technical Breakdown

### Backend

- identity migrations
- auth service
- token service
- password hashing
- session repository
- RBAC seed
- actor context

### API Gateway

- auth middleware
- token validation
- request_id/correlation_id propagation
- protected route placeholder

### Web Frontend

- login page
- auth state
- protected route shell
- logout action

### Mobile

- login screen baseline
- secure token storage placeholder
- logout flow

### QA

- login tests
- invalid login tests
- refresh rotation tests
- permission baseline tests

### DevOps

- identity DB env
- JWT secret separation
- CI secret check

### Documentation

- auth API notes
- permission baseline notes

---

## 11. Service and Data Ownership

| Service/App | Ownership |
| --- | --- |
| Identity Service | users, roles, permissions, sessions, refresh tokens |
| API Gateway | token validation and request context propagation |
| Web Admin | login UI and auth state |
| Mobile App | login UI and secure token storage |

Ownership rules:

```text
- Service hanya boleh mengakses database miliknya sendiri.
- Cross-service data access menggunakan gRPC atau domain events.
- Shared package tidak boleh berisi domain business logic.
- API Gateway bertanggung jawab pada routing/edge concern, bukan business logic.
```

---

## 12. API / gRPC / Event Impact

| Area | Impact |
| --- | --- |
| REST | `POST /api/v1/auth/login`, `POST /api/v1/auth/refresh`, `POST /api/v1/auth/logout`, `GET /api/v1/me`, `GET /api/v1/me/permissions`, `GET /api/v1/me/context` |
| gRPC/proto | Optional Identity internal endpoints for validating user context if needed later |
| Event | Optional `identity.user_logged_in` / security events if defined; not required for business modules |
| OpenAPI | Auth endpoint contracts must be documented |
| Event schema | No Confidential token/password data in event payload |

---

## 13. Data Model Impact

Potential entities/tables:

- users
- roles
- permissions
- role_permissions
- user_role_assignments
- user_sessions
- refresh_tokens/security session table if separate

Migration rules:

```text
- Migration harus reviewable dan rollback-aware.
- Data sensitif harus diklasifikasi.
- Finance amount wajib decimal/NUMERIC jika ada.
- Tidak boleh membuat foreign key lintas database service.
```

---

## 14. Permission and Scope Requirements

- RBAC seed baseline
- foundation/school scope placeholders
- actor context includes roles, permissions, foundation_id, school_id if available
- service-side authorization remains mandatory

General rules:

```text
- RBAC menentukan role/permission.
- ABAC/scope menentukan foundation/school/class/subject/student/parent-child access.
- Object-level authorization wajib untuk resource by ID.
- Frontend visibility bukan pengganti backend authorization.
```

---

## 15. Audit Requirements

Actions requiring audit or audit awareness:

- login failure threshold/security event if implemented
- logout/revoke session
- role assignment
- permission assignment

Audit log should include:

```text
actor
action
resource
timestamp
request_id
correlation_id
safe metadata
```

---

## 16. File and Privacy Requirements

Identity & Access privacy requirements:

- password hashed
- refresh token hashed
- no password/token in logs
- no production secrets
- token storage secure in clients

Data classification reference:

```text
Public
Internal
Restricted
Confidential
```

---

## 17. Test Plan

### Unit Tests

- Service/business logic tests.
- Validation tests.
- Error handling tests.

### Integration Tests

- Repository/database tests.
- Service integration tests.
- gRPC/event integration tests if relevant.

### API Tests

- Success response.
- Validation error.
- Unauthorized access.
- Forbidden access.
- Not found with object-level auth.

### Permission/Scope Tests

- Allowed role/scope can access.
- Disallowed role/scope is rejected.
- Cross-school/cross-foundation access rejected.
- Object by ID outside scope rejected.

### Event Tests

- Event published when required.
- Event payload safe.
- Consumer idempotency if consumer exists.

### Audit Tests

- Sensitive action creates audit log.
- Audit payload does not include raw Confidential data.

### Frontend Tests

- Form validation.
- Loading/empty/error state.
- Permission-based UI visibility.
- API error handling.

### Mobile Tests

- Auth/session behavior if mobile affected.
- Main mobile flow if included.
- Secure token/file behavior if relevant.

### UAT Scenarios

- Role-based user validates core flow.
- Negative case validated.
- QA verifies acceptance criteria.
- Product Owner validates business expectation.

---

## 18. Acceptance Criteria

Sprint 1 acceptance criteria:

- [ ] All Critical scope items completed.
- [ ] No out-of-scope MVP module added.
- [ ] Relevant APIs/contracts documented or updated.
- [ ] Relevant migrations added and reviewed.
- [ ] Permission/scope rules implemented where applicable.
- [ ] Object-level authorization tested where applicable.
- [ ] Sensitive actions audited where applicable.
- [ ] Events added and tested where applicable.
- [ ] Frontend/mobile flow implemented where in scope.
- [ ] Unit/integration/API tests pass.
- [ ] CI pass.
- [ ] QA validates core flow.
- [ ] Documentation updated.

---

## 19. Definition of Ready

Task/issue is ready when:

```text
- objective is clear
- scope is clear
- out of scope is clear
- acceptance criteria are checklist-based
- target service/app is identified
- data model impact is known
- API/proto/event impact is known
- permission/scope requirement is known
- audit/file/privacy requirement is known
- owner is assigned
- labels are assigned
- milestone is Sprint 1
- AI Agent status is set
```

---

## 20. Definition of Done

Task is done when:

```text
- implementation matches issue scope
- no out-of-scope feature added
- tests added or updated
- tests pass locally
- CI pass
- reviewer approved
- QA pass if required
- documentation updated if required
- issue moved to Done
```

Sprint 1 is done when:

```text
- all blocking Critical/High Sprint 1 issues are Done
- core flow acceptance criteria pass
- no Critical/High bug remains for sprint scope
- handoff notes for next sprint are documented
```

---

## 21. Dependencies

- Sprint 0 API Gateway skeleton
- Sprint 0 service template
- PostgreSQL local dependency
- CI baseline

Dependency notes:

```text
- If dependency is not ready, issue status should be Blocked.
- Blocked issue must explain blocker and owner.
- Do not bypass service boundary to avoid dependency.
```

---

## 22. Risks and Mitigations

| Risk | Impact | Probability | Mitigation | Owner |
| --- | --- | --- | --- | --- |
| Refresh token reuse vulnerability | Credential/session compromise | Medium | Hash refresh token, rotate every refresh, test reuse rejection | Backend Developer |
| Token/password leaked to logs | Security breach | Medium | Log review and no sensitive log tests | Security Reviewer |
| RBAC model too narrow | Rework in later sprints | Medium | Use extensible role/permission model with scope placeholders | Backend Lead |

---

## 23. AI Agent Usage Guidance

### Suitable for AI Agent

- Draft scaffolding.
- Generate boilerplate handlers/services/repositories.
- Draft migrations and sqlc queries for human review.
- Generate table-driven tests.
- Draft documentation updates.
- Create QA checklist drafts.

### Requires Human Review

- Authentication/security-sensitive logic.
- Authorization/scope/object-level access.
- Database migration final shape.
- Event contract final shape.
- Finance calculations.
- File privacy/signed URL logic.
- Audit behavior.
- Release/production decision.

### Must Not Be Done by AI Agent

```text
- production secrets
- final security approval
- legal/compliance decision
- production deployment approval
- access to real sensitive data
- major architecture change without explicit instruction
- hotfix production without human review
```

### Required Coding Prompt

Use:

```text
AGENTS.md
SKILLS.md
docs/25-product-requirement-document.md
docs/26-development-plan.md
docs/27-workflow.md
docs/25-github-project-management.md
docs/09-ai-agent-rules.md
docs/08-coding-standard.md
docs/14-sprint-1-task-prompts.md
docs/30-sprint-1-plan.md
```

---

## 24. GitHub Issue Plan

| Issue Title | Type | Area | Priority | Estimate | Labels | Milestone |
| --- | --- | --- | --- | --- | --- | --- |
| Sprint 1 Task 1.1 — Create Identity Database Migrations | feature | identity | critical | 5 | `type: feature`, `area: identity`, `sprint: 1`, `priority: critical`, `status: ready`, `ai: ready`, `risk: data-sensitive, review: backend, review: security` | Sprint 1 — Identity & Access |
| Sprint 1 Task 1.2 — Implement Password Hashing and User Repository | feature | identity | critical | 3 | `type: feature`, `area: identity`, `sprint: 1`, `priority: critical`, `status: ready`, `ai: ready`, `risk: data-sensitive, review: security` | Sprint 1 — Identity & Access |
| Sprint 1 Task 1.3 — Implement Login Endpoint | feature | identity | critical | 5 | `type: feature`, `area: identity`, `sprint: 1`, `priority: critical`, `status: ready`, `ai: ready`, `risk: data-sensitive, review: backend` | Sprint 1 — Identity & Access |
| Sprint 1 Task 1.4 — Implement Rotating Refresh Token | feature | identity | critical | 5 | `type: feature`, `area: identity`, `sprint: 1`, `priority: critical`, `status: ready`, `ai: ready`, `risk: data-sensitive, review: security` | Sprint 1 — Identity & Access |
| Sprint 1 Task 1.5 — Implement Logout and Session Revocation | feature | identity | high | 3 | `type: feature`, `area: identity`, `sprint: 1`, `priority: high`, `status: ready`, `ai: ready`, `review: backend` | Sprint 1 — Identity & Access |
| Sprint 1 Task 1.6 — Seed Roles and Permissions Baseline | feature | identity | critical | 3 | `type: feature`, `area: identity`, `sprint: 1`, `priority: critical`, `status: ready`, `ai: ready`, `review: backend, review: product` | Sprint 1 — Identity & Access |
| Sprint 1 Task 1.7 — Implement Actor Context Endpoints | feature | identity | critical | 3 | `type: feature`, `area: identity`, `sprint: 1`, `priority: critical`, `status: ready`, `ai: ready`, `review: backend` | Sprint 1 — Identity & Access |
| Sprint 1 Task 1.8 — Add API Gateway Auth Middleware | feature | api-gateway | critical | 5 | `type: feature`, `area: api-gateway`, `sprint: 1`, `priority: critical`, `status: ready`, `ai: ready`, `review: backend, review: security` | Sprint 1 — Identity & Access |
| Sprint 1 Task 1.9 — Build Web Login Screen | feature | web-admin | high | 3 | `type: feature`, `area: web-admin`, `sprint: 1`, `priority: high`, `status: ready`, `ai: ready`, `review: frontend` | Sprint 1 — Identity & Access |
| Sprint 1 Task 1.10 — Add Auth Test Coverage | test | identity | critical | 5 | `type: test`, `area: identity`, `sprint: 1`, `priority: critical`, `status: ready`, `ai: ready`, `review: qa, review: security` | Sprint 1 — Identity & Access |

---

## 25. GitHub Project Fields

Recommended fields:

| Field | Recommended Value |
|---|---|
| Status | Ready / In Progress / In Review / QA / Done / Blocked |
| Sprint | Sprint 1 |
| Priority | Critical / High / Medium / Low |
| Area | Based on module/service/app |
| Type | feature / bug / chore / docs / test / security / infra |
| Owner | Assigned developer/QA/DevOps/Product reviewer |
| Estimate | 1 / 2 / 3 / 5 / 8 |
| Risk | Low / Medium / High / Migration / Data Sensitive / Breaking Change |
| Platform | Backend / Web / Mobile / Infra / Docs / QA / Product |
| AI Agent | Ready / Needs Context / Generated / Needs Human Review / Do Not Use |
| Target Release | MVP |

---

## 26. Sprint Exit Criteria

Sprint 1 may be closed when:

```text
- All Critical and High sprint issues are Done or explicitly deferred by Product Owner.
- CI passes on develop.
- QA validates sprint core flows.
- Security/privacy review completed for sensitive changes.
- Documentation updated.
- Handoff notes to next sprint completed.
```

Additional exit criteria for this sprint:

```text
Sprint 2 dapat menggunakan actor context, role, permission, dan authenticated API baseline untuk menerapkan School Core scope.
```

---

## 27. Handoff Notes

Sprint 2 dapat menggunakan actor context, role, permission, dan authenticated API baseline untuk menerapkan School Core scope.

Next sprint should review:

```text
- completed APIs/contracts
- unresolved risks
- known bugs
- deferred issues
- data model decisions
- permission/scope decisions
- event/audit behavior
```
