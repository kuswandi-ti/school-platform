# Sprint 8 Plan — Communication / Notification

Project: `school-platform`  
Sprint: Sprint 8 — Communication / Notification  
Target Output: `docs/37-sprint-8-plan.md`  
Document Type: Sprint Planning Document  
Audience: Product Owner, Developer, QA, DevOps, AI Agent  
Status: Draft for Sprint Execution  

---

## 1. Sprint Summary

Sprint 8 membangun announcement dan notification system berbasis event: announcement CRUD/publish, target audience, templates, in-app notifications, delivery logs, provider abstraction FCM/email, preferences, dan event consumers.

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

Menyediakan komunikasi internal dan notifikasi dasar yang aman, tertarget, dan event-driven.

---

## 3. Business Context

Sekolah dan yayasan membutuhkan sarana pengumuman dan notifikasi yang lebih konsisten daripada channel informal. Orang tua/siswa/guru perlu menerima informasi yang relevan sesuai scope.

---

## 4. Technical Context

Communication Service owns announcements, notifications, templates, delivery logs, preferences, and event consumers. Business services must not call FCM/email directly.

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

- communication_db migrations
- announcements
- announcement_targets
- notifications
- notification_templates
- notification_deliveries
- notification_preferences
- announcement publish
- target by foundation/school/class/role/user
- event consumers
- in-app notification
- read/unread
- FCM mock/provider abstraction
- email mock/provider abstraction
- delivery log
- notification preferences

---

## 6. Out of Scope

- WhatsApp
- SMS
- marketing automation
- advanced campaign builder

---

## 7. Target Users / Actors

- Admin Yayasan
- Kepala Sekolah
- TU/Staff
- Guru
- Orang Tua
- Siswa
- Communication Service

---

## 8. User Stories

- As a Admin Yayasan, I want mengirim pengumuman lintas sekolah, so that informasi yayasan tersampaikan.
- As a Kepala Sekolah, I want mengirim pengumuman sekolah/kelas, so that komunikasi sekolah tertata.
- As a Orang Tua, I want menerima notifikasi tagihan/rapor, so that tidak melewatkan informasi penting.
- As a Siswa, I want melihat pengumuman dan notifikasi, so that mengetahui informasi akademik.

---

## 9. Functional Breakdown

- Announcement CRUD/publish
- Target audience
- In-app notification
- Read/unread
- Notification templates
- Event consumers
- Delivery logs
- Preferences

---

## 10. Technical Breakdown

### Backend

- communication migrations
- announcement service
- notification service
- template service
- event consumers
- delivery logs

### API Gateway

- communication route forwarding

### Web Frontend

- announcement UI
- notification center

### Mobile

- notification list
- announcement view
- push abstraction integration

### QA

- target tests
- privacy tests
- consumer idempotency

### DevOps

- queue consumer config
- provider env placeholders

### Documentation

- notification event and privacy rules

---

## 11. Service and Data Ownership

| Service/App | Ownership |
| --- | --- |
| Communication Service | announcement, notification, templates, delivery, preferences |
| Operational services | publish domain events only |

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
| REST | announcement CRUD/publish, notification list/read, preferences |
| gRPC/proto | Optional user/target lookup if needed |
| Event | Consumes finance/admission/report/academic events; may publish notification delivery events |
| OpenAPI | Communication endpoint contracts |
| Event schema | No Confidential data in notification body/payload |

---

## 13. Data Model Impact

Potential entities/tables:

- announcements
- announcement_targets
- notifications
- notification_templates
- notification_deliveries
- notification_preferences

Migration rules:

```text
- Migration harus reviewable dan rollback-aware.
- Data sensitif harus diklasifikasi.
- Finance amount wajib decimal/NUMERIC jika ada.
- Tidak boleh membuat foreign key lintas database service.
```

---

## 14. Permission and Scope Requirements

- announcement authoring by role/scope
- target by foundation/school/class/role/user
- user sees only own notifications
- critical notification cannot be fully disabled

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

- announcement publish
- cross-school/foundation announcement
- template changes if sensitive

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

Communication / Notification privacy requirements:

- notification body must not contain Confidential data
- Restricted info minimized
- preferences respected except critical notifications

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

Sprint 8 acceptance criteria:

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
- milestone is Sprint 8
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

Sprint 8 is done when:

```text
- all blocking Critical/High Sprint 8 issues are Done
- core flow acceptance criteria pass
- no Critical/High bug remains for sprint scope
- handoff notes for next sprint are documented
```

---

## 21. Dependencies

- Sprint 1 Identity
- Sprint 2 School Core
- Sprint 5 Finance events
- Sprint 7 Report Card events

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
| Sensitive data in notification | Privacy breach | Medium | Template review and tests | Security Reviewer |
| Duplicate notifications | User confusion | Medium | Consumer idempotency | Backend Developer |
| Targeting wrong audience | Information leak | Medium | Target scope tests | QA |

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
docs/21-sprint-8-task-prompts.md
docs/37-sprint-8-plan.md
```

---

## 24. GitHub Issue Plan

| Issue Title | Type | Area | Priority | Estimate | Labels | Milestone |
| --- | --- | --- | --- | --- | --- | --- |
| Sprint 8 Task 8.1 — Create Communication Database Migrations | feature | communication | critical | 5 | `type: feature`, `area: communication`, `sprint: 8`, `priority: critical`, `status: ready`, `ai: ready`, `risk: migration` | Sprint 8 — Communication / Notification |
| Sprint 8 Task 8.2 — Implement Announcement CRUD and Publish | feature | communication | critical | 5 | `type: feature`, `area: communication`, `sprint: 8`, `priority: critical`, `status: ready`, `ai: ready`, `review: backend` | Sprint 8 — Communication / Notification |
| Sprint 8 Task 8.3 — Implement Announcement Targeting | feature | communication | critical | 5 | `type: feature`, `area: communication`, `sprint: 8`, `priority: critical`, `status: ready`, `ai: ready`, `review: security` | Sprint 8 — Communication / Notification |
| Sprint 8 Task 8.4 — Implement In-App Notifications | feature | communication | critical | 5 | `type: feature`, `area: communication`, `sprint: 8`, `priority: critical`, `status: ready`, `ai: ready`, `review: backend` | Sprint 8 — Communication / Notification |
| Sprint 8 Task 8.5 — Implement Read/Unread and Preferences | feature | communication | high | 3 | `type: feature`, `area: communication`, `sprint: 8`, `priority: high`, `status: ready`, `ai: ready`, `review: backend` | Sprint 8 — Communication / Notification |
| Sprint 8 Task 8.6 — Implement Notification Templates | feature | communication | high | 3 | `type: feature`, `area: communication`, `sprint: 8`, `priority: high`, `status: ready`, `ai: ready`, `review: product` | Sprint 8 — Communication / Notification |
| Sprint 8 Task 8.7 — Implement Event Consumers | feature | communication | critical | 5 | `type: feature`, `area: communication`, `sprint: 8`, `priority: critical`, `status: ready`, `ai: ready`, `review: backend` | Sprint 8 — Communication / Notification |
| Sprint 8 Task 8.8 — Add FCM/Email Provider Abstraction | feature | communication | medium | 3 | `type: feature`, `area: communication`, `sprint: 8`, `priority: medium`, `status: ready`, `ai: ready`, `review: infra` | Sprint 8 — Communication / Notification |
| Sprint 8 Task 8.9 — Build Notification Center UI | feature | web-admin | high | 3 | `type: feature`, `area: web-admin`, `sprint: 8`, `priority: high`, `status: ready`, `ai: ready`, `review: frontend` | Sprint 8 — Communication / Notification |
| Sprint 8 Task 8.10 — Add Notification Privacy and Idempotency Tests | test | communication | critical | 5 | `type: test`, `area: communication`, `sprint: 8`, `priority: critical`, `status: ready`, `ai: ready`, `review: qa` | Sprint 8 — Communication / Notification |

---

## 25. GitHub Project Fields

Recommended fields:

| Field | Recommended Value |
|---|---|
| Status | Ready / In Progress / In Review / QA / Done / Blocked |
| Sprint | Sprint 8 |
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

Sprint 8 may be closed when:

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
Sprint 9 uses notification/announcement events for dashboard metrics and operational summaries.
```

---

## 27. Handoff Notes

Sprint 9 uses notification/announcement events for dashboard metrics and operational summaries.

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
