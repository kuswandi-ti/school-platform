# Sprint 6 Plan — Academic Basic

Project: `school-platform`  
Sprint: Sprint 6 — Academic Basic  
Target Output: `docs/35-sprint-6-plan.md`  
Document Type: Sprint Planning Document  
Audience: Product Owner, Developer, QA, DevOps, AI Agent  
Status: Draft for Sprint Execution  

---

## 1. Sprint Summary

Sprint 6 membangun fondasi akademik dasar: curriculum, subject, subject group, class subject, schedule, teacher schedule view, attendance input, dan attendance correction.

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

Menyediakan fondasi akademik dasar agar guru dapat melihat jadwal dan menginput absensi sesuai assignment.

---

## 3. Business Context

Sekolah membutuhkan data jadwal dan absensi yang tertib agar operasional kelas dapat dipantau oleh guru, wali kelas, dan kepala sekolah.

---

## 4. Technical Context

Academic Service owns curriculum, subject, schedule, and attendance. School Core owns student/teacher/class. Academic stores reference IDs and must not query school_core_db directly.

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

- academic_db migrations
- curriculums
- learning_phases
- subjects
- subject_groups
- class_subjects
- schedules
- student_attendances
- curriculum baseline
- subject CRUD
- class subject assignment
- schedule CRUD
- teacher schedule view
- attendance input
- attendance correction
- academic basic events
- audit correction

---

## 6. Out of Scope

- report card
- grade book
- full LMS
- advanced timetable optimization
- BK/UKS detail

---

## 7. Target Users / Actors

- Admin Yayasan
- Kepala Sekolah
- TU/Staff
- Guru
- Wali Kelas
- Siswa

---

## 8. User Stories

- As a TU/Staff, I want membuat jadwal pelajaran, so that guru dan siswa memiliki jadwal yang jelas.
- As a Guru, I want melihat jadwal mengajar, so that tahu kelas dan mapel yang harus diajar.
- As a Guru, I want menginput absensi, so that kehadiran siswa tercatat.
- As a Wali Kelas, I want melihat absensi kelas, so that dapat memantau siswa.

---

## 9. Functional Breakdown

- Curriculum setup
- Subject management
- Class subject assignment
- Schedule management
- Teacher schedule view
- Attendance input
- Attendance correction

---

## 10. Technical Breakdown

### Backend

- academic migrations
- subject services
- schedule services
- attendance services
- assignment scope validation
- events/audit

### API Gateway

- academic route forwarding

### Web Frontend

- curriculum/subject UI
- schedule UI
- attendance UI

### Mobile

- teacher schedule/attendance optional
- student schedule view optional

### QA

- assignment scope tests
- attendance tests
- correction audit tests

### DevOps

- academic_db env/migrations

### Documentation

- academic API and workflow updates

---

## 11. Service and Data Ownership

| Service/App | Ownership |
| --- | --- |
| Academic Service | curriculum, subject, schedule, attendance |
| School Core Service | student, teacher, class references |

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
| REST | curriculum/subject/schedule/attendance endpoints |
| gRPC/proto | School Core lookup for assignment validation if needed |
| Event | `academic.schedule_created`, `academic.attendance_recorded`, `academic.attendance_corrected` |
| OpenAPI | Academic endpoint contracts |
| Event schema | Attendance summary only; avoid raw sensitive notes if any |

---

## 13. Data Model Impact

Potential entities/tables:

- curriculums
- learning_phases
- subjects
- subject_groups
- class_subjects
- schedules
- student_attendances

Migration rules:

```text
- Migration harus reviewable dan rollback-aware.
- Data sensitif harus diklasifikasi.
- Finance amount wajib decimal/NUMERIC jika ada.
- Tidak boleh membuat foreign key lintas database service.
```

---

## 14. Permission and Scope Requirements

- TU/Staff school scope for setup
- Guru subject/class assignment scope
- Wali Kelas class scope
- object-level auth for schedule_id/attendance_id

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

- attendance correction
- schedule changes if sensitive
- subject/class assignment changes if implemented here

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

Academic Basic privacy requirements:

- attendance data Restricted
- no sensitive notes in logs
- teacher/student data only by reference

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

Sprint 6 acceptance criteria:

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
- milestone is Sprint 6
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

Sprint 6 is done when:

```text
- all blocking Critical/High Sprint 6 issues are Done
- core flow acceptance criteria pass
- no Critical/High bug remains for sprint scope
- handoff notes for next sprint are documented
```

---

## 21. Dependencies

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
| Teacher accesses wrong class | Privacy and academic integrity issue | Medium | Assignment scope tests | Backend/QA |
| Duplicate attendance | Data inconsistency | Medium | Unique constraints by class/date/student/session | Backend |
| Schedule model too rigid | Rework | Medium | Keep MVP simple but extensible | Tech Lead |

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
docs/19-sprint-6-task-prompts.md
docs/35-sprint-6-plan.md
```

---

## 24. GitHub Issue Plan

| Issue Title | Type | Area | Priority | Estimate | Labels | Milestone |
| --- | --- | --- | --- | --- | --- | --- |
| Sprint 6 Task 6.1 — Create Academic Basic Migrations | feature | academic | critical | 5 | `type: feature`, `area: academic`, `sprint: 6`, `priority: critical`, `status: ready`, `ai: ready`, `risk: migration` | Sprint 6 — Academic Basic |
| Sprint 6 Task 6.2 — Implement Curriculum and Subject Management | feature | academic | high | 5 | `type: feature`, `area: academic`, `sprint: 6`, `priority: high`, `status: ready`, `ai: ready`, `review: backend` | Sprint 6 — Academic Basic |
| Sprint 6 Task 6.3 — Implement Class Subject Assignment | feature | academic | critical | 5 | `type: feature`, `area: academic`, `sprint: 6`, `priority: critical`, `status: ready`, `ai: ready`, `review: backend` | Sprint 6 — Academic Basic |
| Sprint 6 Task 6.4 — Implement Schedule CRUD | feature | academic | critical | 5 | `type: feature`, `area: academic`, `sprint: 6`, `priority: critical`, `status: ready`, `ai: ready`, `review: backend` | Sprint 6 — Academic Basic |
| Sprint 6 Task 6.5 — Implement Teacher Schedule View | feature | academic | high | 3 | `type: feature`, `area: academic`, `sprint: 6`, `priority: high`, `status: ready`, `ai: ready`, `review: frontend` | Sprint 6 — Academic Basic |
| Sprint 6 Task 6.6 — Implement Attendance Input | feature | academic | critical | 5 | `type: feature`, `area: academic`, `sprint: 6`, `priority: critical`, `status: ready`, `ai: ready`, `review: backend` | Sprint 6 — Academic Basic |
| Sprint 6 Task 6.7 — Implement Attendance Correction | feature | academic | high | 3 | `type: feature`, `area: academic`, `sprint: 6`, `priority: high`, `status: ready`, `ai: ready`, `review: security` | Sprint 6 — Academic Basic |
| Sprint 6 Task 6.8 — Build Academic Web Screens | feature | web-admin | high | 5 | `type: feature`, `area: web-admin`, `sprint: 6`, `priority: high`, `status: ready`, `ai: ready`, `review: frontend` | Sprint 6 — Academic Basic |
| Sprint 6 Task 6.9 — Add Teacher Assignment Scope Tests | test | academic | critical | 5 | `type: test`, `area: academic`, `sprint: 6`, `priority: critical`, `status: ready`, `ai: ready`, `review: qa` | Sprint 6 — Academic Basic |
| Sprint 6 Task 6.10 — Add Attendance Audit Tests | test | academic | high | 3 | `type: test`, `area: academic`, `sprint: 6`, `priority: high`, `status: ready`, `ai: ready`, `review: qa` | Sprint 6 — Academic Basic |

---

## 25. GitHub Project Fields

Recommended fields:

| Field | Recommended Value |
|---|---|
| Status | Ready / In Progress / In Review / QA / Done / Blocked |
| Sprint | Sprint 6 |
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

Sprint 6 may be closed when:

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
Sprint 7 uses Academic subject/class and attendance context for report card and grade book.
```

---

## 27. Handoff Notes

Sprint 7 uses Academic subject/class and attendance context for report card and grade book.

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
