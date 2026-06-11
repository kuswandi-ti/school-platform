# Sprint 4 Plan — PPDB

Project: `school-platform`  
Sprint: Sprint 4 — PPDB  
Target Output: `docs/33-sprint-4-plan.md`  
Document Type: Sprint Planning Document  
Audience: Product Owner, Developer, QA, DevOps, AI Agent  
Status: Draft for Sprint Execution  

---

## 1. Sprint Summary

Sprint 4 membangun proses PPDB dari periode pendaftaran, applicant submission, upload dokumen, verifikasi dokumen, keputusan accept/reject, sampai konversi applicant menjadi student di School Core.

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

Menyediakan alur PPDB MVP yang tertib, terdokumentasi, aman, dan dapat mengonversi pendaftar diterima menjadi siswa aktif.

---

## 3. Business Context

PPDB adalah alur penerimaan siswa baru. Yayasan/sekolah membutuhkan status pendaftar yang jelas, dokumen terkontrol, keputusan tercatat, dan konversi siswa tanpa duplikasi.

---

## 4. Technical Context

Admission Service owns applicant data before conversion. School Core owns student data after conversion. Conversion must use gRPC and must be idempotent.

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

- admission_db migrations
- admission_periods
- applicants
- applicant_guardians
- applicant_documents
- applicant_verifications
- admission_decisions
- admission period CRUD
- applicant submission
- document upload
- document verification
- accept/reject decision
- conversion to student through School Core gRPC
- converted_student_id tracking
- idempotency for conversion
- PPDB events
- audit for decisions

---

## 6. Out of Scope

- complex scoring/selection
- public marketing website
- payment gateway
- advanced admission analytics
- direct write to school_core_db

---

## 7. Target Users / Actors

- Admin Yayasan
- Kepala Sekolah
- TU/Staff
- Orang Tua calon siswa
- Admission Service
- School Core Service

---

## 8. User Stories

- As a TU/Staff, I want membuka periode PPDB, so that pendaftaran dapat dikelola per sekolah.
- As a Orang Tua, I want mengisi data calon siswa dan upload dokumen, so that pendaftaran dapat dilakukan terstruktur.
- As a Kepala Sekolah, I want memberi keputusan accept/reject, so that keputusan PPDB terdokumentasi.
- As a TU/Staff, I want mengonversi applicant diterima menjadi siswa, so that data siswa aktif terbentuk tanpa input ulang.

---

## 9. Functional Breakdown

- Admission period management
- Applicant data capture
- Guardian data capture
- Document upload
- Document verification
- Accept/reject decision
- Conversion to student
- Applicant status tracking

---

## 10. Technical Breakdown

### Backend

- admission migrations
- workflow/status service
- document integration
- decision service
- gRPC conversion to School Core
- events
- audit

### API Gateway

- admission route forwarding
- file upload route integration

### Web Frontend

- period UI
- applicant list/detail
- verification UI
- decision UI
- conversion action

### Mobile

- optional applicant/parent view if included; otherwise no major scope

### QA

- workflow tests
- document tests
- conversion idempotency
- scope tests

### DevOps

- admission_db env/migrations

### Documentation

- PPDB workflow and API updates

---

## 11. Service and Data Ownership

| Service/App | Ownership |
| --- | --- |
| Admission Service | admission period, applicant, applicant guardian, verification, decision |
| File Management | applicant document file storage/metadata |
| School Core Service | student/guardian after conversion |

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
| REST | admission period CRUD, applicant submit/list/detail, document verify, decision, conversion |
| gRPC/proto | School Core convert/create student from applicant snapshot |
| Event | `admission.applicant_submitted`, `admission.document_verified`, `admission.applicant_accepted`, `admission.applicant_rejected`, `admission.applicant_converted` |
| OpenAPI | Admission endpoints |
| Event schema | Use IDs and safe metadata; avoid raw document contents |

---

## 13. Data Model Impact

Potential entities/tables:

- admission_periods
- applicants
- applicant_guardians
- applicant_documents
- applicant_verifications
- admission_decisions

Migration rules:

```text
- Migration harus reviewable dan rollback-aware.
- Data sensitif harus diklasifikasi.
- Finance amount wajib decimal/NUMERIC jika ada.
- Tidak boleh membuat foreign key lintas database service.
```

---

## 14. Permission and Scope Requirements

- Admin Yayasan foundation scope
- Kepala Sekolah school decision scope
- TU/Staff school operational scope
- parent/applicant access to own submission if external account exists
- object-level auth for applicant_id and document_id

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

- document verification
- accept/reject decision
- conversion to student
- sensitive applicant update

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

PPDB privacy requirements:

- applicant and documents classified Restricted
- file private
- signed URL authorization
- no raw document data in logs

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

Sprint 4 acceptance criteria:

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
- milestone is Sprint 4
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

Sprint 4 is done when:

```text
- all blocking Critical/High Sprint 4 issues are Done
- core flow acceptance criteria pass
- no Critical/High bug remains for sprint scope
- handoff notes for next sprint are documented
```

---

## 21. Dependencies

- Sprint 2 School Core
- Sprint 3 File Management
- Sprint 1 Identity

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
| Double conversion creates duplicate student | Data duplication | Medium | Idempotency and converted_student_id | Backend Developer |
| Applicant document leak | Privacy incident | Medium | Private file + signed URL auth | Security Reviewer |
| PPDB workflow status inconsistent | Operational confusion | Medium | Explicit status transition tests | QA |

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
docs/17-sprint-4-task-prompts.md
docs/33-sprint-4-plan.md
```

---

## 24. GitHub Issue Plan

| Issue Title | Type | Area | Priority | Estimate | Labels | Milestone |
| --- | --- | --- | --- | --- | --- | --- |
| Sprint 4 Task 4.1 — Create Admission Database Migrations | feature | admission | critical | 5 | `type: feature`, `area: admission`, `sprint: 4`, `priority: critical`, `status: ready`, `ai: ready`, `risk: migration, review: backend` | Sprint 4 — PPDB |
| Sprint 4 Task 4.2 — Implement Admission Period Management | feature | admission | high | 3 | `type: feature`, `area: admission`, `sprint: 4`, `priority: high`, `status: ready`, `ai: ready`, `review: backend` | Sprint 4 — PPDB |
| Sprint 4 Task 4.3 — Implement Applicant Submission | feature | admission | critical | 5 | `type: feature`, `area: admission`, `sprint: 4`, `priority: critical`, `status: ready`, `ai: ready`, `risk: data-sensitive, review: backend` | Sprint 4 — PPDB |
| Sprint 4 Task 4.4 — Implement Applicant Document Upload Integration | feature | admission | critical | 5 | `type: feature`, `area: admission`, `sprint: 4`, `priority: critical`, `status: ready`, `ai: ready`, `review: security` | Sprint 4 — PPDB |
| Sprint 4 Task 4.5 — Implement Document Verification Workflow | feature | admission | high | 3 | `type: feature`, `area: admission`, `sprint: 4`, `priority: high`, `status: ready`, `ai: ready`, `review: backend` | Sprint 4 — PPDB |
| Sprint 4 Task 4.6 — Implement Accept/Reject Decision | feature | admission | critical | 5 | `type: feature`, `area: admission`, `sprint: 4`, `priority: critical`, `status: ready`, `ai: ready`, `review: product, review: security` | Sprint 4 — PPDB |
| Sprint 4 Task 4.7 — Implement Applicant Conversion to Student via gRPC | feature | admission | critical | 5 | `type: feature`, `area: admission`, `sprint: 4`, `priority: critical`, `status: ready`, `ai: ready`, `review: backend` | Sprint 4 — PPDB |
| Sprint 4 Task 4.8 — Build PPDB Admin Screens | feature | web-admin | high | 5 | `type: feature`, `area: web-admin`, `sprint: 4`, `priority: high`, `status: ready`, `ai: ready`, `review: frontend` | Sprint 4 — PPDB |
| Sprint 4 Task 4.9 — Add PPDB Events and Audit | feature | admission | high | 3 | `type: feature`, `area: admission`, `sprint: 4`, `priority: high`, `status: ready`, `ai: ready`, `review: backend` | Sprint 4 — PPDB |
| Sprint 4 Task 4.10 — Add PPDB Workflow Tests | test | admission | critical | 5 | `type: test`, `area: admission`, `sprint: 4`, `priority: critical`, `status: ready`, `ai: ready`, `review: qa` | Sprint 4 — PPDB |

---

## 25. GitHub Project Fields

Recommended fields:

| Field | Recommended Value |
|---|---|
| Status | Ready / In Progress / In Review / QA / Done / Blocked |
| Sprint | Sprint 4 |
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

Sprint 4 may be closed when:

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
Sprint 9 consumes PPDB events for admission summary dashboards.
```

---

## 27. Handoff Notes

Sprint 9 consumes PPDB events for admission summary dashboards.

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
