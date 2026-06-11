# 05 — Event Contract

Project: `school-platform`  
Status: Final decision for MVP  
Scope: RabbitMQ domain event contract for async processing, reporting projection, notification, and audit consolidation.

---

## 1. Purpose

Event Contract defines the standard format, naming, routing, reliability rules, and security rules for domain events exchanged between services.

Events are used for:

- reporting projection
- notification delivery
- audit consolidation
- asynchronous side effects
- integration between bounded contexts where immediate response is not required

Events are **not** a replacement for synchronous validation or commands that require immediate consistency.

---

## 2. Core Decision

MVP uses RabbitMQ with a topic exchange:

```text
Exchange name : domain.events
Exchange type : topic
Routing key   : event_type
```

Routing key uses this convention:

```text
domain.entity.action_past_tense
```

Examples:

```text
school.student.created
admission.applicant.accepted
finance.payment.verified
academic.report_card.published
communication.announcement.published
approval.request.created
audit.log.created
```

---

## 3. Event Naming Convention

Events must use past-tense naming because an event represents something that already happened.

Format:

```text
<domain>.<entity>.<action_past_tense>
```

Rules:

- lowercase only
- dot-separated
- English internal naming
- past tense
- no command-style event name
- no UI-specific naming

Correct:

```text
finance.payment.verified
academic.grade_book.submitted
school.student.status_changed
```

Incorrect:

```text
verify_payment
payment.verify
doPaymentVerification
PembayaranDiverifikasi
```

---

## 4. Standard Event Envelope

All domain events must use the same envelope shape.

```json
{
  "event_id": "uuid",
  "event_type": "finance.payment.verified",
  "event_version": 1,
  "source_service": "finance-service",
  "occurred_at": "2026-06-08T10:30:00Z",
  "published_at": "2026-06-08T10:30:01Z",
  "request_id": "req_123",
  "correlation_id": "corr_456",
  "actor": {
    "user_id": "uuid",
    "role": "bendahara_sekolah"
  },
  "tenant": {
    "foundation_id": "uuid",
    "school_id": "uuid"
  },
  "entity": {
    "entity_type": "payment",
    "entity_id": "uuid"
  },
  "payload": {},
  "metadata": {}
}
```

Required fields:

```text
event_id
event_type
event_version
source_service
occurred_at
correlation_id
tenant.foundation_id
entity.entity_type
entity.entity_id
payload
```

Conditionally required fields:

```text
tenant.school_id
```

`school_id` is required when the event is related to a specific school/unit.

Optional fields:

```text
request_id
published_at
actor
metadata
```

---

## 5. Event Context Rules

Every event must carry tenant and tracing context.

Required context:

```text
foundation_id
school_id if relevant
request_id if available
correlation_id
actor_user_id if action is user-triggered
source_service
```

Correlation ID must be propagated from:

```text
HTTP request
→ API Gateway
→ gRPC metadata
→ Service usecase
→ Database transaction
→ Outbox event
→ RabbitMQ headers/payload
→ Consumer logs
```

---

## 6. RabbitMQ Routing

Exchange:

```text
domain.events
```

Queue examples:

```text
reporting-service.events
communication-service.events
audit-service.events
```

Retry/DLQ examples:

```text
reporting-service.events.retry
reporting-service.events.dlq

communication-service.events.retry
communication-service.events.dlq
```

Binding examples:

```text
reporting-service:
- school.*
- admission.*
- finance.*
- academic.*
- approval.*

communication-service:
- finance.bill.generated
- finance.payment.verified
- finance.payment.rejected
- finance.bill.overdue
- academic.attendance.absent_recorded
- academic.report_card.published
- admission.applicant.accepted
- admission.applicant.rejected
- approval.request.created
- communication.announcement.published
```

---

## 7. Event Versioning

Every event must have:

```text
event_version
```

Initial version:

```text
1
```

Backward-compatible changes:

- adding optional fields
- adding metadata
- adding payload field that consumers may ignore

Breaking changes:

- renaming a field
- removing a field
- changing field type
- changing semantics

Breaking changes require a new version and schema update.

Consumers must ignore unknown fields.

---

## 8. Idempotency

Consumers must be idempotent.

Every consumer must track processed event IDs:

```text
processed_events
- id
- event_id
- event_type
- source_service
- processed_at
```

Rule:

```text
If event_id was already processed, skip processing.
```

This prevents:

- duplicate notifications
- duplicate reporting aggregation
- duplicate audit consolidation
- duplicate side effects

---

## 9. Retry and Dead Letter Queue

Retry policy for MVP:

```text
Retry count : 3–5 times
Delay       : 30 seconds → 2 minutes → 10 minutes
Final fail  : send to DLQ
```

DLQ must be monitored by observability stack.

Alerts must exist for:

```text
DLQ message count > 0
Repeated consumer failure
Queue depth too high
Consumer down
```

---

## 10. Transactional Outbox Pattern

Important services must use outbox pattern.

Required for MVP:

```text
identity-service
school-core-service
admission-service
academic-service
finance-service
communication-service
```

Flow:

```text
Start DB transaction
→ Change domain data
→ Insert outbox event in same transaction
→ Commit transaction
→ Outbox worker publishes event to RabbitMQ
→ Mark outbox event as published
```

Suggested table:

```text
outbox_events
- id
- event_id
- event_type
- event_version
- aggregate_type
- aggregate_id
- payload_json
- status
- retry_count
- next_retry_at
- published_at
- created_at
```

Status:

```text
pending
published
failed
```

Reason:

```text
Domain data change and event publication must not be separated unsafely.
```

---

## 11. Security Rules

Events must not contain:

```text
password
access token
refresh token
private key
raw document content
Confidential data detail
BK/UKS detail
payroll detail
full NIK unless explicitly required
raw payment proof file
```

For sensitive domain actions:

```text
event payload should include entity_id and safe metadata only.
```

Optional metadata:

```json
{
  "metadata": {
    "classification": "restricted",
    "contains_sensitive_data": false
  }
}
```

Confidential data must be fetched only from the owner service using proper authorization, not embedded inside events.

---

## 12. Event Schema Location

Event schemas must be stored in:

```text
packages/events/
```

Suggested structure:

```text
packages/events/
├── envelope.schema.json
├── identity/
│   └── user.created.v1.schema.json
├── school/
│   └── student.created.v1.schema.json
├── admission/
│   └── applicant.accepted.v1.schema.json
├── finance/
│   ├── bill.generated.v1.schema.json
│   └── payment.verified.v1.schema.json
├── academic/
│   └── report_card.published.v1.schema.json
└── communication/
    └── announcement.published.v1.schema.json
```

Rule:

```text
If an event is added or changed, packages/events must be updated.
```

---

## 13. MVP Event List

### Identity Service

```text
identity.user.created
identity.user.invited
identity.user.activated
identity.user.disabled
identity.role.assigned
identity.role.revoked
identity.session.revoked
identity.password_reset.requested
```

### School Core Service

```text
school.foundation.created
school.school.created
school.academic_year.activated
school.semester.activated
school.semester.closed
school.student.created
school.student.updated
school.student.status_changed
school.student.transferred
school.student.graduated
school.guardian.created
school.teacher.created
school.teacher.updated
school.class.created
school.student_class.assigned
school.homeroom.assigned
```

### Admission Service

```text
admission.period.opened
admission.period.closed
admission.applicant.submitted
admission.applicant.document_uploaded
admission.applicant.verified
admission.applicant.accepted
admission.applicant.rejected
admission.applicant.converted_to_student
```

### Finance Service

```text
finance.fee_type.created
finance.fee_policy.submitted
finance.fee_policy.approved
finance.fee_policy.rejected
finance.bill.generated
finance.bill.overdue
finance.payment.created
finance.payment.proof_uploaded
finance.payment.verified
finance.payment.rejected
finance.payment.void_requested
finance.payment.voided
finance.receipt.generated
finance.reconciliation.closed
```

### Academic Service

```text
academic.subject.created
academic.schedule.created
academic.attendance.marked
academic.attendance.corrected
academic.grade_book.created
academic.grade_book.submitted
academic.grade_book.approved
academic.report_card.generated
academic.report_card.published
academic.report_card.revision_requested
academic.report_card.revised
```

### Communication Service

```text
communication.announcement.created
communication.announcement.published
communication.notification.created
communication.notification.delivered
communication.notification.failed
communication.letter.approved
communication.letter.generated
```

### Approval

```text
approval.request.created
approval.request.approved
approval.request.rejected
approval.request.revision_requested
```

### Audit

```text
audit.log.created
```

---

## 14. Example Events

### school.student.created

```json
{
  "event_id": "uuid",
  "event_type": "school.student.created",
  "event_version": 1,
  "source_service": "school-core-service",
  "occurred_at": "2026-06-08T10:30:00Z",
  "published_at": "2026-06-08T10:30:01Z",
  "request_id": "req_001",
  "correlation_id": "corr_001",
  "actor": {
    "user_id": "uuid",
    "role": "tu_staff"
  },
  "tenant": {
    "foundation_id": "uuid",
    "school_id": "uuid"
  },
  "entity": {
    "entity_type": "student",
    "entity_id": "uuid"
  },
  "payload": {
    "student_id": "uuid",
    "student_number": "SD20260001",
    "full_name": "Andi Pratama",
    "status": "active",
    "academic_year_id": "uuid",
    "class_id": "uuid"
  },
  "metadata": {
    "classification": "restricted"
  }
}
```

### finance.bill.generated

```json
{
  "event_id": "uuid",
  "event_type": "finance.bill.generated",
  "event_version": 1,
  "source_service": "finance-service",
  "occurred_at": "2026-06-08T10:30:00Z",
  "published_at": "2026-06-08T10:30:01Z",
  "request_id": "req_002",
  "correlation_id": "corr_002",
  "actor": {
    "user_id": "uuid",
    "role": "bendahara_sekolah"
  },
  "tenant": {
    "foundation_id": "uuid",
    "school_id": "uuid"
  },
  "entity": {
    "entity_type": "bill",
    "entity_id": "uuid"
  },
  "payload": {
    "bill_id": "uuid",
    "student_id": "uuid",
    "invoice_number": "INV/SD/2026/07/000001",
    "billing_period": "2026-07",
    "total_amount": 500000,
    "due_date": "2026-07-10",
    "status": "unpaid"
  },
  "metadata": {
    "classification": "restricted"
  }
}
```

### finance.payment.verified

```json
{
  "event_id": "uuid",
  "event_type": "finance.payment.verified",
  "event_version": 1,
  "source_service": "finance-service",
  "occurred_at": "2026-06-08T10:30:00Z",
  "published_at": "2026-06-08T10:30:01Z",
  "request_id": "req_003",
  "correlation_id": "corr_003",
  "actor": {
    "user_id": "uuid",
    "role": "bendahara_sekolah"
  },
  "tenant": {
    "foundation_id": "uuid",
    "school_id": "uuid"
  },
  "entity": {
    "entity_type": "payment",
    "entity_id": "uuid"
  },
  "payload": {
    "payment_id": "uuid",
    "bill_id": "uuid",
    "student_id": "uuid",
    "payment_number": "PAY/SD/2026/07/000001",
    "amount": 500000,
    "payment_method": "bank_transfer_manual",
    "verified_at": "2026-06-08T10:30:00Z"
  },
  "metadata": {
    "classification": "restricted"
  }
}
```

### academic.report_card.published

```json
{
  "event_id": "uuid",
  "event_type": "academic.report_card.published",
  "event_version": 1,
  "source_service": "academic-service",
  "occurred_at": "2026-06-08T10:30:00Z",
  "published_at": "2026-06-08T10:30:01Z",
  "request_id": "req_004",
  "correlation_id": "corr_004",
  "actor": {
    "user_id": "uuid",
    "role": "kepala_sekolah"
  },
  "tenant": {
    "foundation_id": "uuid",
    "school_id": "uuid"
  },
  "entity": {
    "entity_type": "report_card",
    "entity_id": "uuid"
  },
  "payload": {
    "report_card_id": "uuid",
    "student_id": "uuid",
    "class_id": "uuid",
    "academic_year_id": "uuid",
    "semester_id": "uuid",
    "published_at": "2026-06-08T10:30:00Z"
  },
  "metadata": {
    "classification": "restricted"
  }
}
```

---

## 15. Consumer Responsibilities

Every consumer must:

- validate envelope
- verify `event_version`
- check idempotency using `event_id`
- process event within its own transaction where possible
- log using `correlation_id`
- avoid leaking sensitive data
- retry safely
- send permanently failed messages to DLQ

---

## 16. AI Agent Rules for Event Work

AI Agent must:

- use standard envelope
- define typed event payload structs
- use event type constants
- preserve correlation ID
- preserve tenant context
- add event schema when new event is created
- add tests for event publication
- add idempotency logic for event consumers
- avoid Confidential data in payload
- not publish events directly from HTTP handlers

AI Agent must not:

- invent event naming conventions
- put raw document/token/password data in events
- skip outbox for important domain changes
- make consumers non-idempotent
- bypass RabbitMQ with ad-hoc async calls

---

## 17. Final Summary

MVP event contract uses:

```text
RabbitMQ topic exchange: domain.events
Routing key: event_type
Naming: domain.entity.action_past_tense
Envelope: standardized JSON
Reliability: outbox + idempotent consumers + retry + DLQ
Security: no Confidential detail in payload
Schemas: packages/events
```
