# Sprint 2 Plan — School Core

Project: `school-platform`  
Sprint: Sprint 2 — School Core  
Target Output: `docs/31-sprint-2-plan.md`  
Document Type: Sprint Planning Document  
Audience: Product Owner, Developer, QA, DevOps, AI Agent  
Status: Draft for Sprint Execution  

---

## 1. Sprint Summary

Sprint 2 membangun master data inti yayasan/sekolah: foundation, school, academic year, semester, students, guardians, teachers, grade levels, classes, student-class assignment, teacher assignment, dan homeroom assignment.

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

Menyediakan data master operasional yang menjadi dependency utama PPDB, Finance, Academic, Report Card, Communication, dan Reporting.

---

## 3. Business Context

Yayasan membutuhkan satu sumber data siswa, orang tua, guru, kelas, dan sekolah. Data ini harus rapi, scoped per yayasan/sekolah, dan siap digunakan modul berikutnya.

---

## 4. Technical Context

School Core Service menjadi owner data foundation, school, student, guardian, teacher, class, dan assignment. Service lain hanya boleh menggunakan reference ID atau gRPC/event, bukan query langsung ke school_core_db.

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

- school_core_db migrations
- foundations
- schools
- academic_years
- semesters
- students
- guardians
- student_guardians
- teachers
- grade_levels
- classes
- student_class_assignments
- teacher_assignments
- homeroom_assignments
- CRUD core data
- search/filter
- school/foundation scope checks
- school core events
- audit for sensitive changes

---

## 6. Out of Scope

- PPDB process
- Excel import
- finance
- academic grade/report card
- payroll/HR lengkap

---

## 7. Target Users / Actors

- Admin Yayasan
- Kepala Sekolah
- TU/Staff
- Backend Service Consumers
- QA

---

## 8. User Stories

- As a Admin Yayasan, I want membuat unit sekolah, so that yayasan dapat mengelola TK/SD/SMP/SMA.
- As a TU/Staff, I want mengelola data siswa dan orang tua, so that administrasi sekolah tersusun rapi.
- As a Kepala Sekolah, I want melihat data guru/kelas, so that dapat memantau struktur sekolah.
- As a Service lain, I want mereferensikan student_id/school_id, so that tidak perlu memiliki data master sendiri.

---

## 9. Functional Breakdown

- Foundation and school management
- Academic year and semester management
- Student/guardian management
- Teacher management
- Grade level and class management
- Student-class assignment
- Teacher and homeroom assignment
- Search/filter core data

---

## 10. Technical Breakdown

### Backend

- migrations
- repositories/sqlc
- CRUD services
- scope checks
- events
- audit

### API Gateway

- route forwarding to school-core-service
- protected route integration

### Web Frontend

- school setup screens
- student list/form
- teacher list/form
- class list/form

### Mobile

- no major scope; optional profile context usage

### QA

- CRUD tests
- scope tests
- duplicate validation
- audit verification

### DevOps

- school_core_db env/migration command
- local seed/test data

### Documentation

- data model/API updates if needed

---

## 11. Service and Data Ownership

| Service/App | Ownership |
| --- | --- |
| School Core Service | foundation, school, academic year, semester, student, guardian, teacher, class, assignment |
| Identity Service | user credentials and role assignments only |
| API Gateway | routing only |

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
| REST | CRUD endpoints for foundation/school/student/guardian/teacher/class/assignment |
| gRPC/proto | School Core lookup/validation endpoints for later PPDB/Finance/Academic |
| Event | `school.student_created`, `school.student_updated`, `school.teacher_created`, `school.class_created` if defined |
| OpenAPI | Core CRUD contracts |
| Event schema | No unnecessary personal data; use IDs and safe summary |

---

## 13. Data Model Impact

Potential entities/tables:

- foundations
- schools
- academic_years
- semesters
- students
- guardians
- student_guardians
- teachers
- grade_levels
- classes
- student_class_assignments
- teacher_assignments
- homeroom_assignments

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
- Kepala Sekolah school scope
- TU/Staff school scope
- object-level check for student/class/teacher by school_id

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

- student sensitive data update
- guardian relation update
- teacher assignment update
- homeroom assignment update

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

School Core privacy requirements:

- student/guardian/teacher data classified Restricted
- no raw sensitive data in logs
- exports/downloads deferred or audited if introduced

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

Sprint 2 acceptance criteria:

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
- milestone is Sprint 2
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

Sprint 2 is done when:

```text
- all blocking Critical/High Sprint 2 issues are Done
- core flow acceptance criteria pass
- no Critical/High bug remains for sprint scope
- handoff notes for next sprint are documented
```

---

## 21. Dependencies

- Sprint 1 Identity & Access
- actor context and permission baseline

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
| Data model tidak fleksibel untuk TK/SD/SMP/SMA | Rework besar | Medium | Gunakan grade_level/class yang fleksibel | Backend Lead |
| Cross-school data access | Privacy leak | High | Object-level authorization tests | Backend/QA |
| Duplicate student/guardian data | Data quality buruk | Medium | Validation and unique constraints where appropriate | Backend Developer |

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
docs/15-sprint-2-task-prompts.md
docs/31-sprint-2-plan.md
```

---

## 24. GitHub Issue Plan

| Issue Title | Type | Area | Priority | Estimate | Labels | Milestone |
| --- | --- | --- | --- | --- | --- | --- |
| Sprint 2 Task 2.1 — Create School Core Database Migrations | feature | school-core | critical | 5 | `type: feature`, `area: school-core`, `sprint: 2`, `priority: critical`, `status: ready`, `ai: ready`, `risk: migration, review: backend` | Sprint 2 — School Core |
| Sprint 2 Task 2.2 — Implement Foundation and School CRUD | feature | school-core | critical | 5 | `type: feature`, `area: school-core`, `sprint: 2`, `priority: critical`, `status: ready`, `ai: ready`, `review: backend` | Sprint 2 — School Core |
| Sprint 2 Task 2.3 — Implement Academic Year and Semester | feature | school-core | high | 3 | `type: feature`, `area: school-core`, `sprint: 2`, `priority: high`, `status: ready`, `ai: ready`, `review: backend` | Sprint 2 — School Core |
| Sprint 2 Task 2.4 — Implement Student and Guardian Management | feature | school-core | critical | 5 | `type: feature`, `area: school-core`, `sprint: 2`, `priority: critical`, `status: ready`, `ai: ready`, `risk: data-sensitive, review: security` | Sprint 2 — School Core |
| Sprint 2 Task 2.5 — Implement Teacher Management | feature | school-core | high | 3 | `type: feature`, `area: school-core`, `sprint: 2`, `priority: high`, `status: ready`, `ai: ready`, `review: backend` | Sprint 2 — School Core |
| Sprint 2 Task 2.6 — Implement Grade Level and Class Management | feature | school-core | high | 3 | `type: feature`, `area: school-core`, `sprint: 2`, `priority: high`, `status: ready`, `ai: ready`, `review: backend` | Sprint 2 — School Core |
| Sprint 2 Task 2.7 — Implement Student-Class Assignment | feature | school-core | critical | 5 | `type: feature`, `area: school-core`, `sprint: 2`, `priority: critical`, `status: ready`, `ai: ready`, `review: backend` | Sprint 2 — School Core |
| Sprint 2 Task 2.8 — Implement Teacher and Homeroom Assignment | feature | school-core | high | 5 | `type: feature`, `area: school-core`, `sprint: 2`, `priority: high`, `status: ready`, `ai: ready`, `review: backend` | Sprint 2 — School Core |
| Sprint 2 Task 2.9 — Build Web Admin Master Data Screens | feature | web-admin | high | 5 | `type: feature`, `area: web-admin`, `sprint: 2`, `priority: high`, `status: ready`, `ai: ready`, `review: frontend` | Sprint 2 — School Core |
| Sprint 2 Task 2.10 — Add School Core Scope and Audit Tests | test | school-core | critical | 5 | `type: test`, `area: school-core`, `sprint: 2`, `priority: critical`, `status: ready`, `ai: ready`, `review: qa, review: security` | Sprint 2 — School Core |

---

## 25. GitHub Project Fields

Recommended fields:

| Field | Recommended Value |
|---|---|
| Status | Ready / In Progress / In Review / QA / Done / Blocked |
| Sprint | Sprint 2 |
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

Sprint 2 may be closed when:

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
Sprint 3 menggunakan School Core data untuk import target; Sprint 4 menggunakan School Core untuk applicant conversion; Sprint 5 menggunakan student data untuk bills.
```

---

## 27. Handoff Notes

Sprint 3 menggunakan School Core data untuk import target; Sprint 4 menggunakan School Core untuk applicant conversion; Sprint 5 menggunakan student data untuk bills.

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
