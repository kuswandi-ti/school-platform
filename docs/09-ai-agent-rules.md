# 09 — AI Agent Rules

Project: `school-platform`  
Status: Mandatory implementation rules  
Audience: AI Agent, developers, reviewers, QA

---

## 1. Purpose

This document defines mandatory rules for AI Agent implementation.

The goal is to keep AI-assisted development:

- focused
- consistent
- secure
- testable
- aligned with architecture decisions
- safe for microservice boundaries
- resistant to feature creep

AI Agent must read and follow this document before implementing any task.

---

## 2. Non-Negotiable Architecture Rules

AI Agent must never violate these rules.

```text
1. Do not query another service database.
2. Do not put business logic in API Gateway.
3. Do not bypass service boundary.
4. Do not remove foundation_id or school_id from domain data.
5. Do not trust frontend-provided scope without backend validation.
6. Do not create public access to private/sensitive files.
7. Do not log token, password, or Confidential data.
8. Do not hardcode production secrets.
9. Do not change API/proto/event contract without updating docs.
10. Do not implement features outside the assigned scope.
```

---

## 3. Service Boundary Rules

Service data ownership:

```text
Identity Service:
- user account
- credentials
- sessions
- refresh tokens
- roles
- permissions
- role assignments

School Core Service:
- foundation
- school
- academic year
- semester
- student master
- guardian/parent master
- teacher master dasar
- class/rombel
- student-class assignment
- teacher assignment
- homeroom assignment

Admission Service:
- PPDB process
- applicant before accepted/converted

Academic Service:
- curriculum
- subject
- schedule
- attendance
- grade
- report card
- report template

Finance Service:
- fee type
- fee scheme
- fee policy
- sibling discount
- bill
- payment
- receipt
- reconciliation
- finance approval

Communication Service:
- announcement
- notification
- template
- delivery log
- preference
- basic letters if included

Reporting Service:
- dashboard projection/read model only
```

Rule:

```text
A service may store reference IDs to data owned by another service, but must not become the source of truth for that data.
```

---

## 4. Database Rules

AI Agent must:

```text
Use pgx + sqlc.
Use parameterized queries.
Use service-owned database only.
Include foundation_id in domain queries.
Include school_id in school-scoped queries.
Use UUID for primary keys.
Use snake_case for database naming.
Use plural table names.
Use lowercase snake_case enum/status values.
```

AI Agent must not:

```text
write raw SQL string concatenation from user input
join across service databases
create cross-database foreign keys
access reporting_db from operational service except via defined reporting workflow
```

---

## 5. Authorization Rules

Every protected usecase must receive actor context:

```text
user_id
foundation_id
school_id
role
permissions
scope
request_id
correlation_id
```

AI Agent must implement:

```text
authentication check
permission check
scope check
object-level authorization
```

Common scope rules:

```text
Admin Yayasan → foundation scope
Kepala Sekolah → school scope
TU/Staff → school scope
Bendahara → school + finance scope
Guru → assigned class/subject scope
Wali Kelas → assigned class scope
Orang Tua → linked student scope
Siswa → self scope
```

Frontend hiding menu is not authorization.

---

## 6. API Rules

External API:

```text
REST/JSON via API Gateway
prefix /api/v1
standard response format
standard error format
OpenAPI must be updated if endpoint changes
```

Internal API:

```text
gRPC/protobuf
proto contract in packages/proto
all calls carry request/correlation/actor/tenant context
```

Response format:

```json
{
  "data": {},
  "meta": null,
  "error": null
}
```

Error format:

```json
{
  "data": null,
  "meta": null,
  "error": {
    "code": "FORBIDDEN",
    "message": "Anda tidak memiliki akses ke data ini.",
    "details": {}
  }
}
```

AI Agent must not create ad-hoc response formats.

---

## 7. Error Handling Rules

Use standard error codes:

```text
UNAUTHORIZED
FORBIDDEN
VALIDATION_ERROR
NOT_FOUND
CONFLICT
BUSINESS_RULE_VIOLATION
APPROVAL_REQUIRED
RESOURCE_LOCKED
RATE_LIMITED
INTERNAL_ERROR
SERVICE_UNAVAILABLE
```

Rules:

```text
Do not expose stack trace to frontend.
Log internal errors with correlation_id.
For out-of-scope resources, NOT_FOUND is allowed to avoid leaking existence.
```

---

## 8. Audit Log Rules

Sensitive actions must create audit logs.

Examples:

```text
identity.role.assigned
school.student.updated
admission.applicant.accepted
finance.payment.verified
finance.payment.voided
finance.fee_policy.approved
academic.report_card.published
academic.report_card.revised
file.downloaded
```

Audit log must include:

```text
actor_user_id
foundation_id
school_id if relevant
action
entity_type
entity_id
old_values_json if relevant
new_values_json if relevant
reason if required
request_id
correlation_id
occurred_at
```

AI Agent must not implement sensitive action without audit.

---

## 9. Approval Rules

Actions that require approval:

```text
void payment
refund
free_spp
large discount/beasiswa
sibling discount application
publish report card
revise report card after publish
role sensitive assignment
export Restricted/Confidential data if required
```

Approval must include:

```text
requester
approver
reason
status
before/after snapshot if relevant
audit log
```

MVP mostly uses one-level approval, with structure prepared for multi-level.

---

## 10. Event Rules

AI Agent must:

```text
use standard event envelope
use event type convention: domain.entity.action_past_tense
include event_id
include event_version
include source_service
include occurred_at
include correlation_id
include tenant context
use outbox pattern for important domain changes
make consumers idempotent
update packages/events when adding/changing event
```

AI Agent must not:

```text
publish events directly from HTTP handlers
include token/password/raw document/Confidential detail in event payload
create event names outside convention
skip event tests for event-driven workflows
```

---

## 11. File Rules

All personal and operational files are private by default.

AI Agent must:

```text
validate MIME type
validate extension
validate size
store metadata
store file in private object storage
generate signed URL only after authorization
audit Restricted/Confidential download/export
use short expiry for signed URL
```

AI Agent must not:

```text
save private files in public path
return permanent public URL for private files
log raw file content
overwrite official documents
```

---

## 12. Finance Rules

AI Agent must:

```text
use decimal/money library for finance calculation
never use float for money
store bill snapshot
store fee policy snapshot when generating bill
use Idempotency-Key for bill generation/payment operations
require approval for sensitive finance actions
audit payment verification, rejection, void, refund, fee policy change
```

Free SPP, discount, scholarship, and sibling discount are:

```text
student_fee_policy in Finance Service
not student status
```

---

## 13. Academic Rules

AI Agent must enforce:

```text
teacher assignment check
class/subject scope
report card lock after publish
approval for revision after publish
audit report card publish/revision
parent/student can only view published report card
```

AI Agent must not:

```text
allow teacher to edit published grades
allow teacher to access unassigned class
allow parent/student to view unpublished report card
```

---

## 14. Reporting Rules

Reporting Service:

```text
reads reporting_db projection
consumes domain events
may run scheduled rebuild
is not operational source of truth
```

AI Agent must not:

```text
make Reporting Service query operational service databases directly
use dashboard endpoint to access domain databases directly
```

---

## 15. Notification Rules

AI Agent must:

```text
send notification through Communication/Notification Service
use event-driven notification
respect notification templates
respect delivery log
respect user preference except critical notification
avoid Confidential detail in notification body
```

Do not call notification provider directly from business service.

---

## 16. Security Rules

AI Agent must enforce:

```text
HTTPS-ready production config
rate limit sensitive endpoints
input validation
strict CORS configuration
security headers for web
private file storage
least privilege database user
no secrets in repository
no sensitive raw logs
```

Sensitive data must not be logged:

```text
password
access token
refresh token
NIK full
BK/UKS detail
payroll
document content
payment proof raw content
```

---

## 17. Testing Rules

Every task must include relevant tests.

Required test types based on task:

```text
unit test for business logic
integration test for repository/database behavior
API test for endpoint behavior
permission/scope test for protected access
event test for event publication/consumption
audit test for sensitive action
```

AI Agent must include negative tests for sensitive features.

---

## 18. Documentation Update Rules

AI Agent must update documentation if changing:

```text
API endpoint → packages/openapi and docs/04-api-contract.md if needed
gRPC proto → packages/proto
event → packages/events and docs/05-event-contract.md if needed
database schema → migrations and docs/03-data-model-mvp.md if needed
new screen → docs/06-ui-screen-user-flow.md if needed
security behavior → relevant docs
```

---

## 19. Task Scope Rules

AI Agent must keep task small.

Good task:

```text
Implement create fee type endpoint in finance-service with migration, sqlc query, usecase, handler mapping, permission check, audit log, and tests.
```

Bad task:

```text
Build the whole Finance Service.
```

Each task must define:

```text
Context
Goal
Scope
Out of scope
Files to create/modify
Rules
Acceptance criteria
Tests
Expected output
```

---

## 20. Definition of Done

A task is done only if:

```text
Code matches requested scope.
No unrelated feature is added.
Tests are added.
Relevant tests pass.
Permission/scope checks exist.
Audit log exists for sensitive action.
Event is published if required.
No cross-service database access exists.
No sensitive data is logged.
Response/error format follows convention.
Migration belongs to correct service.
OpenAPI/proto/event docs are updated if needed.
Summary and test instructions are provided.
```

---

## 21. Review Checklist

Before merging AI-generated code, reviewer must check:

```text
Does it follow service boundary?
Does it query only owned database?
Does it use actor context?
Does it enforce permission/scope?
Does it handle object-level authorization?
Does it create audit log if sensitive?
Does it publish event if required?
Does it avoid sensitive logs?
Does it include tests?
Does it update contract docs?
Does it preserve existing conventions?
```

---

## 22. Final Summary

AI Agent must prioritize:

```text
architecture discipline
small scoped tasks
tests
permission/scope correctness
auditability
event consistency
data privacy
no cross-service database access
```

If unsure, AI Agent should avoid expanding scope and follow the existing architecture decisions.
