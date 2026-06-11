# Sprint 9 Plan — Reporting Dashboard

Project: `school-platform`  
Sprint: Sprint 9 — Reporting Dashboard  
Target Output: `docs/38-sprint-9-plan.md`  
Document Type: Sprint Planning Document  
Audience: Product Owner, Developer, QA, DevOps, AI Agent  
Status: Draft for Sprint Execution  

---

## 1. Sprint Summary

Sprint 9 membangun Reporting Service sebagai read model/projection untuk dashboard Yayasan, Sekolah, Guru, dan Orang Tua/Siswa berbasis event dari service operasional.

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

Menyediakan dashboard MVP berbasis projection yang tidak melanggar service boundary.

---

## 3. Business Context

Yayasan dan sekolah membutuhkan ringkasan lintas modul: siswa, guru, PPDB, finance, absensi, progres akademik, rapor, dan notifikasi. Dashboard harus membantu pengambilan keputusan cepat.

---

## 4. Technical Context

Reporting Service only reads reporting_db. It consumes domain events from RabbitMQ and builds projections. It must not query operational service DBs directly.

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

- reporting_db migrations
- dashboard projection tables
- processed_events
- projection offsets
- RabbitMQ consumer infrastructure
- student/teacher summary projection
- admission summary projection
- finance summary projection
- attendance summary projection
- academic progress projection
- dashboard APIs
- scheduled rebuild/sync skeleton
- idempotent consumers

---

## 6. Out of Scope

- advanced BI
- global search
- data warehouse
- direct query to operational DBs

---

## 7. Target Users / Actors

- Admin Yayasan
- Kepala Sekolah
- Guru
- Orang Tua
- Siswa
- Reporting Service

---

## 8. User Stories

- As a Admin Yayasan, I want melihat dashboard lintas sekolah, so that dapat memantau kondisi yayasan.
- As a Kepala Sekolah, I want melihat dashboard sekolah, so that dapat memantau operasional unit.
- As a Guru, I want melihat ringkasan tugas akademik, so that dapat memprioritaskan pekerjaan.
- As a Orang Tua/Siswa, I want melihat ringkasan personal, so that informasi penting mudah diakses.

---

## 9. Functional Breakdown

- Foundation dashboard
- School dashboard
- Teacher dashboard
- Parent/student dashboard
- Event projections
- Processed events
- Scheduled rebuild skeleton

---

## 10. Technical Breakdown

### Backend

- reporting migrations
- projection consumers
- processed_events
- dashboard APIs
- scope filters

### API Gateway

- dashboard route forwarding

### Web Frontend

- foundation dashboard
- school dashboard
- teacher dashboard

### Mobile

- parent/student summary cards

### QA

- projection tests
- idempotency tests
- dashboard scope tests

### DevOps

- consumer worker setup
- RabbitMQ/DLQ monitoring

### Documentation

- dashboard metrics and projection rules

---

## 11. Service and Data Ownership

| Service/App | Ownership |
| --- | --- |
| Reporting Service | reporting_db projections and dashboard APIs |
| Operational Services | source of truth and domain event producers |

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
| REST | dashboard summary endpoints by role/scope |
| gRPC/proto | Not primary; reporting uses events/projections |
| Event | Consumes school/admission/finance/academic/report/communication events |
| OpenAPI | Dashboard API contracts |
| Event schema | processed_events idempotency based on event_id |

---

## 13. Data Model Impact

Potential entities/tables:

- processed_events
- projection_offsets
- foundation_dashboard_summaries
- school_dashboard_summaries
- teacher_dashboard_summaries
- parent_student_dashboard_summaries
- finance_summary_projections
- admission_summary_projections
- attendance_summary_projections
- academic_progress_projections

Migration rules:

```text
- Migration harus reviewable dan rollback-aware.
- Data sensitif harus diklasifikasi.
- Finance amount wajib decimal/NUMERIC jika ada.
- Tidak boleh membuat foreign key lintas database service.
```

---

## 14. Permission and Scope Requirements

- dashboard APIs enforce role/scope
- Admin Yayasan foundation scope
- Kepala Sekolah school scope
- Guru assignment scope
- parent-child scope
- student self scope

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

- dashboard access audit optional for high sensitivity
- export/download if introduced must be audited

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

Reporting Dashboard privacy requirements:

- dashboard must aggregate/minimize data
- no operational DB direct query
- no Confidential metrics exposed

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

Sprint 9 acceptance criteria:

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
- milestone is Sprint 9
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

Sprint 9 is done when:

```text
- all blocking Critical/High Sprint 9 issues are Done
- core flow acceptance criteria pass
- no Critical/High bug remains for sprint scope
- handoff notes for next sprint are documented
```

---

## 21. Dependencies

- Sprint 4 PPDB events
- Sprint 5 Finance events
- Sprint 6 Academic events
- Sprint 7 Report events
- Sprint 8 Communication events

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
| Reporting queries operational DB | Architecture violation | Medium | Code review and tests around repository dependencies | Backend Lead |
| Duplicate events corrupt metrics | Wrong dashboard | Medium | processed_events idempotency | Backend Developer |
| Dashboard exposes cross-school data | Privacy breach | Medium | Scope tests | QA |

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
docs/22-sprint-9-task-prompts.md
docs/38-sprint-9-plan.md
```

---

## 24. GitHub Issue Plan

| Issue Title | Type | Area | Priority | Estimate | Labels | Milestone |
| --- | --- | --- | --- | --- | --- | --- |
| Sprint 9 Task 9.1 — Create Reporting Database Migrations | feature | reporting | critical | 5 | `type: feature`, `area: reporting`, `sprint: 9`, `priority: critical`, `status: ready`, `ai: ready`, `risk: migration` | Sprint 9 — Reporting Dashboard |
| Sprint 9 Task 9.2 — Implement Processed Events and Idempotency | feature | reporting | critical | 5 | `type: feature`, `area: reporting`, `sprint: 9`, `priority: critical`, `status: ready`, `ai: ready`, `review: backend` | Sprint 9 — Reporting Dashboard |
| Sprint 9 Task 9.3 — Implement Student and Teacher Summary Projection | feature | reporting | high | 3 | `type: feature`, `area: reporting`, `sprint: 9`, `priority: high`, `status: ready`, `ai: ready`, `review: backend` | Sprint 9 — Reporting Dashboard |
| Sprint 9 Task 9.4 — Implement Admission Summary Projection | feature | reporting | high | 3 | `type: feature`, `area: reporting`, `sprint: 9`, `priority: high`, `status: ready`, `ai: ready`, `review: backend` | Sprint 9 — Reporting Dashboard |
| Sprint 9 Task 9.5 — Implement Finance Summary Projection | feature | reporting | critical | 5 | `type: feature`, `area: reporting`, `sprint: 9`, `priority: critical`, `status: ready`, `ai: ready`, `review: backend` | Sprint 9 — Reporting Dashboard |
| Sprint 9 Task 9.6 — Implement Attendance and Academic Progress Projection | feature | reporting | high | 5 | `type: feature`, `area: reporting`, `sprint: 9`, `priority: high`, `status: ready`, `ai: ready`, `review: backend` | Sprint 9 — Reporting Dashboard |
| Sprint 9 Task 9.7 — Implement Dashboard APIs | feature | reporting | critical | 5 | `type: feature`, `area: reporting`, `sprint: 9`, `priority: critical`, `status: ready`, `ai: ready`, `review: security` | Sprint 9 — Reporting Dashboard |
| Sprint 9 Task 9.8 — Build Web Dashboard Screens | feature | web-admin | high | 5 | `type: feature`, `area: web-admin`, `sprint: 9`, `priority: high`, `status: ready`, `ai: ready`, `review: frontend` | Sprint 9 — Reporting Dashboard |
| Sprint 9 Task 9.9 — Build Parent/Student Summary View | feature | mobile | medium | 3 | `type: feature`, `area: mobile`, `sprint: 9`, `priority: medium`, `status: ready`, `ai: ready`, `review: mobile` | Sprint 9 — Reporting Dashboard |
| Sprint 9 Task 9.10 — Add Projection and Dashboard Scope Tests | test | reporting | critical | 5 | `type: test`, `area: reporting`, `sprint: 9`, `priority: critical`, `status: ready`, `ai: ready`, `review: qa` | Sprint 9 — Reporting Dashboard |

---

## 25. GitHub Project Fields

Recommended fields:

| Field | Recommended Value |
|---|---|
| Status | Ready / In Progress / In Review / QA / Done / Blocked |
| Sprint | Sprint 9 |
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

Sprint 9 may be closed when:

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
Sprint 10 uses dashboard and observability readiness to validate MVP release.
```

---

## 27. Handoff Notes

Sprint 10 uses dashboard and observability readiness to validate MVP release.

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
