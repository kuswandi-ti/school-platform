# Sprint 5 Plan — Finance / SPP

Project: `school-platform`  
Sprint: Sprint 5 — Finance / SPP  
Target Output: `docs/34-sprint-5-plan.md`  
Document Type: Sprint Planning Document  
Audience: Product Owner, Developer, QA, DevOps, AI Agent  
Status: Draft for Sprint Execution  

---

## 1. Sprint Summary

Sprint 5 membangun proses Finance/SPP MVP berbasis manual payment: fee type, fee scheme, fee policy, sibling discount, bill generation, payment proof upload, verification/rejection, receipt, dan void approval.

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

Menyediakan alur tagihan SPP manual yang akurat, idempotent, aman, dan dapat diaudit.

---

## 3. Business Context

Bendahara membutuhkan sistem untuk generate tagihan, memantau tunggakan, memverifikasi bukti pembayaran, dan memberikan status pembayaran yang jelas kepada orang tua.

---

## 4. Technical Context

Finance Service owns all fee policy, bill, payment, proof, receipt, and reconciliation records. Amount must use decimal/NUMERIC, never float. Bill must store snapshot.

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

- finance_db migrations
- fee_types
- fee_schemes
- fee_scheme_items
- student_fee_policies
- sibling_discount_rules
- student_bills
- student_bill_items
- student_payments
- payment_proofs
- payment_receipts
- payment_reconciliations
- finance approval requests
- fee policy approval
- bill generation with snapshots
- manual payment proof
- verify/reject payment
- receipt generation
- outstanding/tunggakan
- void payment with approval
- finance events
- audit sensitive actions

---

## 6. Out of Scope

- payment gateway
- automatic bank reconciliation
- full accounting ledger
- payroll
- tax

---

## 7. Target Users / Actors

- Bendahara
- Kepala Sekolah
- Admin Yayasan
- Orang Tua/Wali Murid
- Finance Service
- School Core Service

---

## 8. User Stories

- As a Bendahara, I want membuat fee scheme dan generate tagihan, so that SPP dapat ditagihkan secara konsisten.
- As a Orang Tua, I want melihat tagihan dan upload bukti, so that pembayaran manual dapat diverifikasi.
- As a Bendahara, I want memverifikasi atau menolak pembayaran, so that status pembayaran akurat.
- As a Kepala Sekolah, I want menyetujui void/kebijakan tertentu, so that aksi finance sensitif terkendali.

---

## 9. Functional Breakdown

- Fee type/scheme
- Student fee policy
- Sibling discount
- Bill generation
- Bill snapshot
- Payment proof upload
- Payment verification/rejection
- Receipt
- Outstanding
- Void approval

---

## 10. Technical Breakdown

### Backend

- finance migrations
- decimal money handling
- fee services
- bill generation
- payment verification
- receipt generation
- void approval
- events/audit

### API Gateway

- finance route forwarding
- auth context propagation

### Web Frontend

- fee setup UI
- bill generation UI
- payment verification UI
- outstanding UI

### Mobile

- parent bill list
- upload proof
- payment status

### QA

- idempotency tests
- decimal tests
- scope tests
- payment verification tests

### DevOps

- finance_db env
- file proof storage config

### Documentation

- finance rules and API updates

---

## 11. Service and Data Ownership

| Service/App | Ownership |
| --- | --- |
| Finance Service | fee, policy, bill, bill item, payment, proof relation, receipt, reconciliation |
| School Core Service | student and guardian reference data |
| File Management | payment proof file object/metadata |

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
| REST | fee, scheme, policy, bill generation, payment proof upload, verify/reject, outstanding |
| gRPC/proto | School Core lookup/validation for student/guardian/sibling if needed |
| Event | `finance.bill_generated`, `finance.payment_submitted`, `finance.payment_verified`, `finance.payment_rejected`, `finance.payment_voided` |
| OpenAPI | Finance endpoint contracts |
| Event schema | No raw proof file or confidential finance detail in notification payload |

---

## 13. Data Model Impact

Potential entities/tables:

- fee_types
- fee_schemes
- fee_scheme_items
- student_fee_policies
- sibling_discount_rules
- student_bills
- student_bill_items
- student_payments
- payment_proofs
- payment_receipts
- payment_reconciliations
- finance_approval_requests

Migration rules:

```text
- Migration harus reviewable dan rollback-aware.
- Data sensitif harus diklasifikasi.
- Finance amount wajib decimal/NUMERIC jika ada.
- Tidak boleh membuat foreign key lintas database service.
```

---

## 14. Permission and Scope Requirements

- Bendahara school finance scope
- Kepala Sekolah approval scope
- Admin Yayasan foundation oversight
- parent-child scope for bill view and proof upload
- object-level auth for bill_id/payment_id

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

- fee policy create/update/approval
- bill generation
- payment proof upload
- payment verify/reject
- payment void
- receipt generation

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

Finance / SPP privacy requirements:

- finance data Restricted
- payment proof Restricted file
- signed URL authorization
- no raw finance data in logs
- audit download/export

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

Sprint 5 acceptance criteria:

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
- milestone is Sprint 5
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

Sprint 5 is done when:

```text
- all blocking Critical/High Sprint 5 issues are Done
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
| Wrong billing amount | Finance trust issue | Medium | Decimal, snapshot, table-driven tests | Backend/QA |
| Duplicate bills | Overbilling | Medium | Unique constraints and idempotency | Backend Developer |
| Parent accesses wrong bill | Privacy breach | Medium | Parent-child object-level auth tests | QA |

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
docs/18-sprint-5-task-prompts.md
docs/34-sprint-5-plan.md
```

---

## 24. GitHub Issue Plan

| Issue Title | Type | Area | Priority | Estimate | Labels | Milestone |
| --- | --- | --- | --- | --- | --- | --- |
| Sprint 5 Task 5.1 — Create Finance Database Migrations | feature | finance | critical | 5 | `type: feature`, `area: finance`, `sprint: 5`, `priority: critical`, `status: ready`, `ai: ready`, `risk: migration, risk: data-sensitive` | Sprint 5 — Finance / SPP |
| Sprint 5 Task 5.2 — Implement Fee Type and Fee Scheme | feature | finance | critical | 5 | `type: feature`, `area: finance`, `sprint: 5`, `priority: critical`, `status: ready`, `ai: ready`, `review: backend` | Sprint 5 — Finance / SPP |
| Sprint 5 Task 5.3 — Implement Student Fee Policy and Approval | feature | finance | critical | 5 | `type: feature`, `area: finance`, `sprint: 5`, `priority: critical`, `status: ready`, `ai: ready`, `review: security` | Sprint 5 — Finance / SPP |
| Sprint 5 Task 5.4 — Implement Sibling Discount Rules | feature | finance | high | 3 | `type: feature`, `area: finance`, `sprint: 5`, `priority: high`, `status: ready`, `ai: ready`, `review: backend` | Sprint 5 — Finance / SPP |
| Sprint 5 Task 5.5 — Bill Generation with Snapshots | feature | finance | critical | 5 | `type: feature`, `area: finance`, `sprint: 5`, `priority: critical`, `status: ready`, `ai: ready`, `risk: data-sensitive` | Sprint 5 — Finance / SPP |
| Sprint 5 Task 5.6 — Parent Payment Proof Upload | feature | finance | critical | 5 | `type: feature`, `area: finance`, `sprint: 5`, `priority: critical`, `status: ready`, `ai: ready`, `review: security` | Sprint 5 — Finance / SPP |
| Sprint 5 Task 5.7 — Payment Verification and Rejection | feature | finance | critical | 5 | `type: feature`, `area: finance`, `sprint: 5`, `priority: critical`, `status: ready`, `ai: ready`, `review: backend` | Sprint 5 — Finance / SPP |
| Sprint 5 Task 5.8 — Receipt and Outstanding View | feature | finance | high | 3 | `type: feature`, `area: finance`, `sprint: 5`, `priority: high`, `status: ready`, `ai: ready`, `review: frontend` | Sprint 5 — Finance / SPP |
| Sprint 5 Task 5.9 — Void Payment Approval | feature | finance | high | 5 | `type: feature`, `area: finance`, `sprint: 5`, `priority: high`, `status: ready`, `ai: ready`, `review: security` | Sprint 5 — Finance / SPP |
| Sprint 5 Task 5.10 — Finance Test Suite | test | finance | critical | 5 | `type: test`, `area: finance`, `sprint: 5`, `priority: critical`, `status: ready`, `ai: ready`, `review: qa` | Sprint 5 — Finance / SPP |

---

## 25. GitHub Project Fields

Recommended fields:

| Field | Recommended Value |
|---|---|
| Status | Ready / In Progress / In Review / QA / Done / Blocked |
| Sprint | Sprint 5 |
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

Sprint 5 may be closed when:

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
Sprint 8 consumes finance events for notification; Sprint 9 consumes finance events for dashboard.
```

---

## 27. Handoff Notes

Sprint 8 consumes finance events for notification; Sprint 9 consumes finance events for dashboard.

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
