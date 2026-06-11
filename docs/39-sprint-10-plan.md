# Sprint 10 Plan — Security, Observability, Backup, and UAT Hardening

Project: `school-platform`  
Sprint: Sprint 10 — Security, Observability, Backup, and UAT Hardening  
Target Output: `docs/39-sprint-10-plan.md`  
Document Type: Sprint Planning Document  
Audience: Product Owner, Developer, QA, DevOps, AI Agent  
Status: Draft for Sprint Execution  

---

## 1. Sprint Summary

Sprint 10 melakukan hardening MVP sebelum pilot/production: security review, permission/scope regression, audit log review, metrics/logging, backup/restore, UAT checklist, release readiness, dan rollback documentation.

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

Memastikan MVP siap pilot/production dengan kualitas, keamanan, observability, backup/restore, QA/UAT, dan release governance yang memadai.

---

## 3. Business Context

Sebelum digunakan secara operasional, yayasan membutuhkan keyakinan bahwa core flow aman, stabil, bisa dipantau, bisa dipulihkan, dan tidak memiliki Critical/High bug.

---

## 4. Technical Context

Sprint ini fokus pada hardening lintas modul, bukan fitur baru. Semua critical paths diuji ulang, backup/restore dibuktikan, observability baseline tersedia, dan release workflow dipastikan aman.

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

- security baseline review
- permission/scope audit
- object-level authorization tests
- audit log consistency
- structured JSON logging review
- Prometheus metrics
- Loki/Grafana setup
- RabbitMQ queue/DLQ monitoring
- backup script
- restore procedure
- restore test documentation
- UAT checklist
- regression checklist
- production deployment readiness
- rollback documentation
- bug fixing

---

## 6. Out of Scope

- Kubernetes
- advanced SIEM
- full penetration test unless separately planned
- WAF unless needed
- new feature development

---

## 7. Target Users / Actors

- Product Owner
- QA
- DevOps
- Backend Developer
- Frontend Developer
- Mobile Developer
- Security Reviewer
- User Representative

---

## 8. User Stories

- As a QA, I want menjalankan regression dan UAT, so that release readiness dapat diputuskan.
- As a DevOps, I want menguji backup/restore, so that risiko kehilangan data berkurang.
- As a Security Reviewer, I want memvalidasi permission/scope, so that data sensitif terlindungi.
- As a Product Owner, I want melihat readiness checklist, so that dapat memberi sign-off.

---

## 9. Functional Breakdown

- Security review
- Permission regression
- Audit review
- Observability baseline
- Backup/restore
- UAT checklist
- Release readiness
- Rollback plan
- Bug fixing

---

## 10. Technical Breakdown

### Backend

- fix missing auth checks
- add missing audit
- metrics endpoints
- hardening bug fixes

### API Gateway

- rate limit review
- auth middleware review
- request/correlation propagation review

### Web Frontend

- UAT fixes
- validation polish
- error states

### Mobile

- UAT fixes
- auth/session checks
- parent/student core flow

### QA

- regression
- UAT
- bug verification
- release sign-off

### DevOps

- backup scripts
- restore test
- Grafana/Loki/Prometheus
- RabbitMQ/DLQ monitoring
- release/rollback docs

### Documentation

- release checklist
- backup/restore docs
- UAT results
- known issues

---

## 11. Service and Data Ownership

| Service/App | Ownership |
| --- | --- |
| All Services | security, audit, metrics, tests |
| DevOps | backup/restore, observability, deployment readiness |
| QA | UAT/regression readiness |

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
| REST | Review all core endpoints for permission/scope and error consistency |
| gRPC/proto | Review cross-service calls and no direct DB violation |
| Event | Review event idempotency, DLQ, retry, privacy |
| OpenAPI | Ensure core APIs documented enough for QA/dev handoff |
| Event schema | Ensure no Confidential raw data |

---

## 13. Data Model Impact

Potential entities/tables:

- audit_logs if implemented
- metrics
- backup artifacts
- restore test records
- processed_events review
- all service DB backups

Migration rules:

```text
- Migration harus reviewable dan rollback-aware.
- Data sensitif harus diklasifikasi.
- Finance amount wajib decimal/NUMERIC jika ada.
- Tidak boleh membuat foreign key lintas database service.
```

---

## 14. Permission and Scope Requirements

- full permission/scope regression
- object-level authorization tests
- parent-child scope
- teacher assignment scope
- school/foundation scope
- finance access scope

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

- review all sensitive actions from PRD
- audit log consistency
- download/export audit if implemented
- payment/report/role actions

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

Security, Observability, Backup, and UAT Hardening privacy requirements:

- Restricted/Confidential review
- no sensitive logs
- backup Confidential
- signed URL review
- notification body privacy

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

Sprint 10 acceptance criteria:

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
- milestone is Sprint 10
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

Sprint 10 is done when:

```text
- all blocking Critical/High Sprint 10 issues are Done
- core flow acceptance criteria pass
- no Critical/High bug remains for sprint scope
- handoff notes for next sprint are documented
```

---

## 21. Dependencies

- Sprint 0–9 completion
- staging environment
- QA test data
- DevOps backup target

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
| Critical bug found late | Release delayed | Medium | Prioritize core flow and fix blocker only | QA/Product Owner |
| Restore test fails | Production readiness blocked | Low-Medium | Run restore rehearsal and document fix | DevOps |
| Sensitive data in logs/events | Privacy incident | Medium | Log/event payload review | Security Reviewer |
| Insufficient observability | Difficult incident response | Medium | Prometheus/Grafana/Loki baseline | DevOps |

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
docs/23-sprint-10-task-prompts.md
docs/39-sprint-10-plan.md
```

---

## 24. GitHub Issue Plan

| Issue Title | Type | Area | Priority | Estimate | Labels | Milestone |
| --- | --- | --- | --- | --- | --- | --- |
| Sprint 10 Task 10.1 — Security Baseline Review | security | security | critical | 5 | `type: security`, `area: security`, `sprint: 10`, `priority: critical`, `status: ready`, `ai: ready`, `review: security` | Sprint 10 — Security, Observability, Backup, and UAT Hardening |
| Sprint 10 Task 10.2 — Permission and Scope Regression | test | security | critical | 5 | `type: test`, `area: security`, `sprint: 10`, `priority: critical`, `status: ready`, `ai: ready`, `review: qa` | Sprint 10 — Security, Observability, Backup, and UAT Hardening |
| Sprint 10 Task 10.3 — Object-Level Authorization Test Sweep | test | security | critical | 5 | `type: test`, `area: security`, `sprint: 10`, `priority: critical`, `status: ready`, `ai: ready`, `review: security` | Sprint 10 — Security, Observability, Backup, and UAT Hardening |
| Sprint 10 Task 10.4 — Audit Log Consistency Review | security | observability | high | 3 | `type: security`, `area: observability`, `sprint: 10`, `priority: high`, `status: ready`, `ai: ready`, `review: backend` | Sprint 10 — Security, Observability, Backup, and UAT Hardening |
| Sprint 10 Task 10.5 — Structured Logging and Sensitive Log Review | security | observability | critical | 3 | `type: security`, `area: observability`, `sprint: 10`, `priority: critical`, `status: ready`, `ai: ready`, `review: security` | Sprint 10 — Security, Observability, Backup, and UAT Hardening |
| Sprint 10 Task 10.6 — Add Metrics and Monitoring Baseline | infra | observability | high | 5 | `type: infra`, `area: observability`, `sprint: 10`, `priority: high`, `status: ready`, `ai: ready`, `review: infra` | Sprint 10 — Security, Observability, Backup, and UAT Hardening |
| Sprint 10 Task 10.7 — Add RabbitMQ Queue and DLQ Monitoring | infra | observability | high | 3 | `type: infra`, `area: observability`, `sprint: 10`, `priority: high`, `status: ready`, `ai: ready`, `review: infra` | Sprint 10 — Security, Observability, Backup, and UAT Hardening |
| Sprint 10 Task 10.8 — Implement Backup and Restore Procedure | infra | infra | critical | 5 | `type: infra`, `area: infra`, `sprint: 10`, `priority: critical`, `status: ready`, `ai: ready`, `review: infra` | Sprint 10 — Security, Observability, Backup, and UAT Hardening |
| Sprint 10 Task 10.9 — Execute UAT and Regression Checklist | test | qa | critical | 5 | `type: test`, `area: qa`, `sprint: 10`, `priority: critical`, `status: ready`, `ai: ready`, `review: qa` | Sprint 10 — Security, Observability, Backup, and UAT Hardening |
| Sprint 10 Task 10.10 — Prepare Release Readiness and Rollback Package | docs | ci-cd | critical | 3 | `type: docs`, `area: ci-cd`, `sprint: 10`, `priority: critical`, `status: ready`, `ai: ready`, `review: product` | Sprint 10 — Security, Observability, Backup, and UAT Hardening |

---

## 25. GitHub Project Fields

Recommended fields:

| Field | Recommended Value |
|---|---|
| Status | Ready / In Progress / In Review / QA / Done / Blocked |
| Sprint | Sprint 10 |
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

Sprint 10 may be closed when:

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
After Sprint 10, MVP can enter pilot/production release process if readiness criteria are met.
```

---

## 27. Handoff Notes

After Sprint 10, MVP can enter pilot/production release process if readiness criteria are met.

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
