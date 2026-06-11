# Sprint 3 Plan — File Management + Import Excel

Project: `school-platform`  
Sprint: Sprint 3 — File Management + Import Excel  
Target Output: `docs/32-sprint-3-plan.md`  
Document Type: Sprint Planning Document  
Audience: Product Owner, Developer, QA, DevOps, AI Agent  
Status: Draft for Sprint Execution  

---

## 1. Sprint Summary

Sprint 3 membangun fondasi private file management dan import Excel data awal untuk students, guardians, teachers, classes, dan assignments.

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

Menyediakan mekanisme file private dan import data awal yang aman, tervalidasi, dan dapat diaudit.

---

## 3. Business Context

Yayasan/sekolah kemungkinan memiliki data awal dalam Excel. Sprint ini mengurangi input manual berulang dan menyiapkan upload dokumen yang akan dipakai PPDB dan Finance.

---

## 4. Technical Context

File management menggunakan object storage S3-compatible seperti MinIO. File private by default dan akses menggunakan backend authorization + signed URL. Import harus validation-preview-confirm-report.

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

- file metadata structure
- MinIO/S3 storage abstraction
- private file upload
- signed URL authorization
- MIME/extension/size validation
- classification public/internal/restricted/confidential
- import_batches
- import_batch_rows
- template download
- Excel upload
- validation preview
- confirm import
- import report
- error report

---

## 6. Out of Scope

- import grades
- historical payments import
- payroll import
- asset/library/BK/UKS/alumni/koperasi import
- virus scanning integration
- central File Service

---

## 7. Target Users / Actors

- Admin Yayasan
- TU/Staff
- Bendahara
- QA
- DevOps

---

## 8. User Stories

- As a TU/Staff, I want mengupload Excel data awal, so that data siswa/guru/kelas tidak perlu diinput satu per satu.
- As a TU/Staff, I want melihat preview validasi, so that kesalahan bisa diperbaiki sebelum import.
- As a Admin Yayasan, I want memastikan file private, so that data siswa tidak bocor.
- As a QA, I want melihat import report, so that hasil import dapat diverifikasi.

---

## 9. Functional Breakdown

- File upload
- File metadata
- Signed URL
- Template download
- Excel upload
- Validation preview
- Confirm import
- Import report
- Error report

---

## 10. Technical Breakdown

### Backend

- storage abstraction
- file metadata repository
- signed URL service
- import batch service
- row validation
- confirm import integration with School Core

### API Gateway

- file upload routing
- authenticated file endpoints

### Web Frontend

- upload UI
- template download UI
- preview table
- error report UI

### Mobile

- no major scope

### QA

- file validation tests
- import valid/invalid scenarios
- privacy tests

### DevOps

- MinIO bucket setup
- local env
- object storage backup consideration

### Documentation

- template format
- file privacy rules

---

## 11. Service and Data Ownership

| Service/App | Ownership |
| --- | --- |
| File Management logic | May be implemented in relevant service/shared pattern for MVP; metadata ownership must be explicit |
| School Core Service | final imported students/guardians/teachers/classes |
| Object Storage | binary files only, not source of truth metadata |

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
| REST | upload file, download template, create import batch, validate import, confirm import, get report |
| gRPC/proto | School Core validation/import operation if separated |
| Event | Optional `file.uploaded`, `import.completed` |
| OpenAPI | File/import endpoint contracts |
| Event schema | Do not include raw row data in event payload |

---

## 13. Data Model Impact

Potential entities/tables:

- file_metadata
- import_batches
- import_batch_rows
- import_errors

Migration rules:

```text
- Migration harus reviewable dan rollback-aware.
- Data sensitif harus diklasifikasi.
- Finance amount wajib decimal/NUMERIC jika ada.
- Tidak boleh membuat foreign key lintas database service.
```

---

## 14. Permission and Scope Requirements

- TU/Staff school scope
- Admin Yayasan foundation scope
- file access authorized by owner/scope
- object-level authorization for file_id/import_batch_id

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

- file upload restricted/confidential
- signed URL generation for restricted/confidential
- confirm import
- download import report if sensitive

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

File Management + Import Excel privacy requirements:

- import file classified Restricted
- private bucket
- signed URL expiry
- no raw import data in logs
- MIME/size validation

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

Sprint 3 acceptance criteria:

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
- milestone is Sprint 3
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

Sprint 3 is done when:

```text
- all blocking Critical/High Sprint 3 issues are Done
- core flow acceptance criteria pass
- no Critical/High bug remains for sprint scope
- handoff notes for next sprint are documented
```

---

## 21. Dependencies

- Sprint 2 School Core
- MinIO local setup from Sprint 0
- Identity actor context

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
| Invalid import corrupts data | Data quality incident | Medium | Validation-preview-confirm and transaction/idempotency | Backend/QA |
| Signed URL leak | Privacy incident | Medium | Short expiry and backend auth before URL generation | Security Reviewer |
| Large Excel performance issue | Import slow/fail | Medium | Batch processing and limits | Backend Developer |

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
docs/16-sprint-3-task-prompts.md
docs/32-sprint-3-plan.md
```

---

## 24. GitHub Issue Plan

| Issue Title | Type | Area | Priority | Estimate | Labels | Milestone |
| --- | --- | --- | --- | --- | --- | --- |
| Sprint 3 Task 3.1 — Create File Metadata and Import Batch Migrations | feature | file-management | critical | 5 | `type: feature`, `area: file-management`, `sprint: 3`, `priority: critical`, `status: ready`, `ai: ready`, `risk: migration, review: backend` | Sprint 3 — File Management + Import Excel |
| Sprint 3 Task 3.2 — Implement S3/MinIO Storage Abstraction | feature | file-management | critical | 5 | `type: feature`, `area: file-management`, `sprint: 3`, `priority: critical`, `status: ready`, `ai: ready`, `risk: data-sensitive, review: security` | Sprint 3 — File Management + Import Excel |
| Sprint 3 Task 3.3 — Implement Private File Upload | feature | file-management | critical | 5 | `type: feature`, `area: file-management`, `sprint: 3`, `priority: critical`, `status: ready`, `ai: ready`, `review: backend` | Sprint 3 — File Management + Import Excel |
| Sprint 3 Task 3.4 — Implement Signed URL Authorization | feature | file-management | critical | 5 | `type: feature`, `area: file-management`, `sprint: 3`, `priority: critical`, `status: ready`, `ai: ready`, `risk: data-sensitive, review: security` | Sprint 3 — File Management + Import Excel |
| Sprint 3 Task 3.5 — Create Import Template Download | feature | file-management | high | 3 | `type: feature`, `area: file-management`, `sprint: 3`, `priority: high`, `status: ready`, `ai: ready`, `review: backend` | Sprint 3 — File Management + Import Excel |
| Sprint 3 Task 3.6 — Implement Excel Upload and Validation Preview | feature | file-management | critical | 5 | `type: feature`, `area: file-management`, `sprint: 3`, `priority: critical`, `status: ready`, `ai: ready`, `review: backend` | Sprint 3 — File Management + Import Excel |
| Sprint 3 Task 3.7 — Implement Confirm Import to School Core | feature | school-core | critical | 5 | `type: feature`, `area: school-core`, `sprint: 3`, `priority: critical`, `status: ready`, `ai: ready`, `review: backend` | Sprint 3 — File Management + Import Excel |
| Sprint 3 Task 3.8 — Build Import UI | feature | web-admin | high | 5 | `type: feature`, `area: web-admin`, `sprint: 3`, `priority: high`, `status: ready`, `ai: ready`, `review: frontend` | Sprint 3 — File Management + Import Excel |
| Sprint 3 Task 3.9 — Add Import Error Report UI | feature | web-admin | medium | 3 | `type: feature`, `area: web-admin`, `sprint: 3`, `priority: medium`, `status: ready`, `ai: ready`, `review: frontend` | Sprint 3 — File Management + Import Excel |
| Sprint 3 Task 3.10 — Add File Privacy and Import Tests | test | file-management | critical | 5 | `type: test`, `area: file-management`, `sprint: 3`, `priority: critical`, `status: ready`, `ai: ready`, `review: qa, review: security` | Sprint 3 — File Management + Import Excel |

---

## 25. GitHub Project Fields

Recommended fields:

| Field | Recommended Value |
|---|---|
| Status | Ready / In Progress / In Review / QA / Done / Blocked |
| Sprint | Sprint 3 |
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

Sprint 3 may be closed when:

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
Sprint 4 uses file upload for applicant documents; Sprint 5 uses file upload for payment proof.
```

---

## 27. Handoff Notes

Sprint 4 uses file upload for applicant documents; Sprint 5 uses file upload for payment proof.

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
