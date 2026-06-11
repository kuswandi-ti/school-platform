# Sprint 7 Plan — Report Card / E-Rapor Basic

Project: `school-platform`  
Sprint: Sprint 7 — Report Card / E-Rapor Basic  
Target Output: `docs/36-sprint-7-plan.md`  
Document Type: Sprint Planning Document  
Audience: Product Owner, Developer, QA, DevOps, AI Agent  
Status: Draft for Sprint Execution  

---

## 1. Sprint Summary

Sprint 7 membangun proses nilai dan rapor dasar: assessment components, schemes, grade book, score input, grade book submit/review, report card generation, publish/lock, dan revision approval.

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

Menyediakan workflow e-rapor basic yang aman, dapat direview, dipublish, dikunci, dan hanya dapat dilihat orang tua/siswa setelah publish.

---

## 3. Business Context

Sekolah membutuhkan proses nilai dan rapor yang lebih tertib daripada rekap manual. Orang tua dan siswa membutuhkan akses rapor yang sudah final/published.

---

## 4. Technical Context

Academic Service mengelola score/grade/report card basic. Published report card locked. Revision after publish requires approval and audit.

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

- assessment_components
- assessment_schemes
- grade_books
- student_scores
- report_templates
- report_cards
- report_card_items
- academic approval requests
- score input
- grade book submit
- Wali Kelas review
- Kepala Sekolah publish
- report card lock
- parent/student published view
- revision after publish with approval
- report card events
- audit sensitive actions

---

## 6. Out of Scope

- full LMS
- advanced analytics
- national e-rapor integration
- offline score input
- final PDF design beyond MVP placeholder

---

## 7. Target Users / Actors

- Guru
- Wali Kelas
- Kepala Sekolah
- Orang Tua/Wali Murid
- Siswa
- TU/Staff

---

## 8. User Stories

- As a Guru, I want menginput nilai siswa, so that nilai mapel terdokumentasi.
- As a Wali Kelas, I want mereview rapor kelas, so that data rapor siap dipublish.
- As a Kepala Sekolah, I want mempublish rapor, so that rapor resmi terkunci dan tersedia.
- As a Orang Tua, I want melihat rapor anak yang published, so that dapat memantau perkembangan anak.
- As a Siswa, I want melihat rapor published, so that mendapat informasi akademik pribadi.

---

## 9. Functional Breakdown

- Assessment scheme
- Grade book
- Score input
- Grade book submit
- Homeroom review
- Report card generation
- Publish/lock
- Revision approval
- Parent/student view

---

## 10. Technical Breakdown

### Backend

- grade migrations
- score service
- report card service
- publish/lock workflow
- revision approval
- events/audit

### API Gateway

- report card route forwarding

### Web Frontend

- score input UI
- grade book review UI
- publish UI
- report view

### Mobile

- parent/student published report view

### QA

- score workflow
- publish/lock tests
- revision approval tests
- parent-child scope tests

### DevOps

- academic service worker/config if needed

### Documentation

- report card workflow and API updates

---

## 11. Service and Data Ownership

| Service/App | Ownership |
| --- | --- |
| Academic Service | assessment, score, grade book, report card |
| School Core Service | student/teacher/class references |
| Communication Service | later notification after publish |

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
| REST | assessment, grade book, score, report card review/publish/view endpoints |
| gRPC/proto | School Core lookup for student/class snapshot if needed |
| Event | `report_card.submitted`, `report_card.reviewed`, `report_card.published`, `report_card.revision_requested` |
| OpenAPI | Report card endpoint contracts |
| Event schema | Do not include full scores in notification event payload unless safe and necessary |

---

## 13. Data Model Impact

Potential entities/tables:

- assessment_components
- assessment_schemes
- grade_books
- student_scores
- report_templates
- report_cards
- report_card_items
- academic_approval_requests

Migration rules:

```text
- Migration harus reviewable dan rollback-aware.
- Data sensitif harus diklasifikasi.
- Finance amount wajib decimal/NUMERIC jika ada.
- Tidak boleh membuat foreign key lintas database service.
```

---

## 14. Permission and Scope Requirements

- Guru subject/class assignment
- Wali Kelas class scope
- Kepala Sekolah school publish scope
- parent-child scope for published view
- student self scope for published view

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

- grade book submit
- report card review
- report card publish
- revision after publish
- score revision after lock

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

Report Card / E-Rapor Basic privacy requirements:

- scores/report cards Restricted
- parent/student only see published
- no draft/submitted report to parent/student
- student snapshot stored safely

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

Sprint 7 acceptance criteria:

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
- milestone is Sprint 7
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

Sprint 7 is done when:

```text
- all blocking Critical/High Sprint 7 issues are Done
- core flow acceptance criteria pass
- no Critical/High bug remains for sprint scope
- handoff notes for next sprint are documented
```

---

## 21. Dependencies

- Sprint 6 Academic Basic
- Sprint 2 School Core
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
| Published report changed without approval | Integrity issue | Medium | Lock status and revision approval tests | Backend/QA |
| Parent sees wrong report | Privacy breach | Medium | Parent-child object-level tests | QA |
| Score calculation ambiguity | Business disagreement | Medium | Keep MVP calculation explicit and documented | Product Owner |

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
docs/20-sprint-7-task-prompts.md
docs/36-sprint-7-plan.md
```

---

## 24. GitHub Issue Plan

| Issue Title | Type | Area | Priority | Estimate | Labels | Milestone |
| --- | --- | --- | --- | --- | --- | --- |
| Sprint 7 Task 7.1 — Create Report Card Migrations | feature | academic | critical | 5 | `type: feature`, `area: academic`, `sprint: 7`, `priority: critical`, `status: ready`, `ai: ready`, `risk: migration, risk: data-sensitive` | Sprint 7 — Report Card / E-Rapor Basic |
| Sprint 7 Task 7.2 — Implement Assessment Components and Schemes | feature | academic | high | 5 | `type: feature`, `area: academic`, `sprint: 7`, `priority: high`, `status: ready`, `ai: ready`, `review: backend` | Sprint 7 — Report Card / E-Rapor Basic |
| Sprint 7 Task 7.3 — Implement Grade Book and Score Input | feature | academic | critical | 5 | `type: feature`, `area: academic`, `sprint: 7`, `priority: critical`, `status: ready`, `ai: ready`, `review: backend` | Sprint 7 — Report Card / E-Rapor Basic |
| Sprint 7 Task 7.4 — Implement Grade Book Submit Workflow | feature | academic | high | 3 | `type: feature`, `area: academic`, `sprint: 7`, `priority: high`, `status: ready`, `ai: ready`, `review: backend` | Sprint 7 — Report Card / E-Rapor Basic |
| Sprint 7 Task 7.5 — Implement Homeroom Report Review | feature | academic | high | 5 | `type: feature`, `area: academic`, `sprint: 7`, `priority: high`, `status: ready`, `ai: ready`, `review: product` | Sprint 7 — Report Card / E-Rapor Basic |
| Sprint 7 Task 7.6 — Implement Report Card Publish and Lock | feature | academic | critical | 5 | `type: feature`, `area: academic`, `sprint: 7`, `priority: critical`, `status: ready`, `ai: ready`, `review: security` | Sprint 7 — Report Card / E-Rapor Basic |
| Sprint 7 Task 7.7 — Implement Revision After Publish Approval | feature | academic | critical | 5 | `type: feature`, `area: academic`, `sprint: 7`, `priority: critical`, `status: ready`, `ai: ready`, `review: security` | Sprint 7 — Report Card / E-Rapor Basic |
| Sprint 7 Task 7.8 — Build Score and Report Card Web Screens | feature | web-admin | high | 5 | `type: feature`, `area: web-admin`, `sprint: 7`, `priority: high`, `status: ready`, `ai: ready`, `review: frontend` | Sprint 7 — Report Card / E-Rapor Basic |
| Sprint 7 Task 7.9 — Build Parent/Student Published Report View | feature | mobile | high | 3 | `type: feature`, `area: mobile`, `sprint: 7`, `priority: high`, `status: ready`, `ai: ready`, `review: mobile` | Sprint 7 — Report Card / E-Rapor Basic |
| Sprint 7 Task 7.10 — Add Report Card Scope and Audit Tests | test | academic | critical | 5 | `type: test`, `area: academic`, `sprint: 7`, `priority: critical`, `status: ready`, `ai: ready`, `review: qa` | Sprint 7 — Report Card / E-Rapor Basic |

---

## 25. GitHub Project Fields

Recommended fields:

| Field | Recommended Value |
|---|---|
| Status | Ready / In Progress / In Review / QA / Done / Blocked |
| Sprint | Sprint 7 |
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

Sprint 7 may be closed when:

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
Sprint 8 sends notifications for report published; Sprint 9 projects report progress metrics.
```

---

## 27. Handoff Notes

Sprint 8 sends notifications for report published; Sprint 9 projects report progress metrics.

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
