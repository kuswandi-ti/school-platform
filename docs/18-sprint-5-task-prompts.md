# 18 — Sprint 5 Task Prompts

Sprint: Sprint 5 — Finance / SPP

Project: `school-platform`  
Status: AI Agent task prompt pack  
Usage: Copy one task prompt at a time into AI Agent. Keep tasks small and do not combine unrelated tasks.

---

## Global Rules for This Sprint

AI Agent must follow these rules for every task in this document:

```text
- Follow docs/01-technical-architecture.md.
- Follow docs/02-service-boundary.md.
- Follow docs/03-data-model-mvp.md.
- Follow docs/04-api-contract.md.
- Follow docs/05-event-contract.md.
- Follow docs/07-test-plan-acceptance-criteria.md.
- Follow docs/08-coding-standard.md.
- Follow docs/09-ai-agent-rules.md.
- Follow docs/10-sprint-backlog-mvp.md.
- Follow docs/11-github-repository-rules.md.
- Do not query another service database.
- Do not put business logic in API Gateway.
- Use Go, Chi, gRPC, pgx, sqlc, goose, slog, validator, testify.
- Use UUID primary keys.
- Use foundation_id and school_id correctly.
- Enforce authentication, permission, scope, and object-level authorization.
- Add audit log for sensitive actions.
- Publish events through the standard outbox/event mechanism when required.
- Use standard API response and error format.
- Add tests for success and negative cases.
- Do not log tokens, passwords, or Confidential data.
- Do not implement out-of-scope features.
```

---

# Task 5.1 — Create Finance Database Migrations

## Prompt

```text
You are working on `school-platform`.

Task:
Create Finance Database Migrations

Target:
services/finance-service

Goal:
Create finance_db schema for fees, policies, bills, payments, receipts, reconciliation, approvals.

Scope:
- Create fee_types, fee_schemes, fee_scheme_items, student_fee_policies, sibling_discount_rules, student_bills, student_bill_items, student_payments, payment_proofs, payment_receipts, payment_reconciliations, finance_approval_requests, finance_audit_logs.
- Add indexes/unique constraints.

Out of Scope:
- Payment gateway.
- Payroll.
- Accounting ledger.

Rules:
- Use NUMERIC(14,2).
- No float.
- No FK to school_core_db.
- Snapshot columns for bills.

Acceptance Criteria:
- Migrations run up/down.
- Invoice/payment/receipt numbers unique.
- Indexes exist.

Tests Required:
- Migration tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 5.2 — Implement Fee Type and Fee Scheme

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Fee Type and Fee Scheme

Target:
finance-service

Goal:
Implement fee type and fee scheme management.

Scope:
- CRUD fee_types.
- CRUD fee_schemes.
- CRUD fee_scheme_items.
- Billing frequency.
- Audit changes.

Out of Scope:
- Fee policy.
- Bill generation.

Rules:
- Bendahara/Admin scope.
- No cross-school management unless Admin Yayasan.
- Use decimal.

Acceptance Criteria:
- Fee types created.
- Fee schemes created.
- Duplicate code rejected.
- Audit recorded.

Tests Required:
- CRUD tests.
- Duplicate tests.
- Scope tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 5.3 — Implement Student Fee Policy and Approval

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Student Fee Policy and Approval

Target:
finance-service

Goal:
Implement per-student fee policy for free_spp, discounts, scholarship, custom_fee.

Scope:
- Create student_fee_policies.
- Submit/approve/reject policy.
- Policy period.
- Reason required.
- Audit.
- Publish fee_policy events.

Out of Scope:
- Generate bills.
- Complex rule engine.

Rules:
- Fee policy is not student status.
- Approval required.
- Use student_id reference only.
- Do not query school_core_db directly.

Acceptance Criteria:
- Policy created/submitted/approved/rejected.
- Reason required.
- Only approved policy applies later.
- Event/audit recorded.

Tests Required:
- Policy tests.
- Approval tests.
- Scope tests.
- Audit/event tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 5.4 — Implement Sibling Discount Rules

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Sibling Discount Rules

Target:
finance-service

Goal:
Create configurable sibling discount rules.

Scope:
- CRUD sibling_discount_rules.
- child_order.
- discount_type/value.
- effective_from/until.
- school_id nullable for foundation-level.
- Audit.

Out of Scope:
- Automatic verification without approval.
- Complex family graph.

Rules:
- Rules configurable.
- Application to student policy still needs approval.
- Use decimal.

Acceptance Criteria:
- Rules created.
- Rules effective period works.
- Duplicate active rule prevented if needed.

Tests Required:
- Rule tests.
- Effective period tests.
- Audit tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 5.5 — Implement Bill Generation with Snapshots

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Bill Generation with Snapshots

Target:
finance-service

Goal:
Generate SPP bills applying fee policies and snapshots.

Scope:
- Generate bills by school/period/fee type.
- Apply approved policy.
- Create bill and bill_items.
- Store student_snapshot_json and applied_policy_snapshot_json.
- Idempotency-Key.
- Publish bill.generated.

Out of Scope:
- Payment verification.
- Gateway payments.

Rules:
- Use decimal not float.
- Do not duplicate bills.
- Validate student references via gRPC/projection if available.
- Outbox event.

Acceptance Criteria:
- Bills generated.
- Discount/free_spp applied.
- Snapshots stored.
- Duplicate request prevented.
- Event published.

Tests Required:
- Normal bill tests.
- Discount tests.
- Free SPP tests.
- Idempotency tests.
- Event tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 5.6 — Implement Payment Proof Upload

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Payment Proof Upload

Target:
finance-service

Goal:
Allow parent to upload manual transfer proof.

Scope:
- Create payment pending_verification.
- Upload proof file.
- Link proof to payment.
- Parent scope check.
- Publish proof_uploaded event.

Out of Scope:
- Verify payment.
- Gateway callback.

Rules:
- Parent can only upload for own child.
- File Restricted/private.
- No raw file logs.

Acceptance Criteria:
- Proof uploaded.
- Payment pending.
- Out-of-scope parent rejected.
- Event published.

Tests Required:
- Upload tests.
- Parent scope tests.
- File privacy tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 5.7 — Implement Payment Verification and Rejection

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Payment Verification and Rejection

Target:
finance-service

Goal:
Allow Bendahara to verify/reject manual payments.

Scope:
- Verify payment.
- Reject payment with reason.
- Update bill paid/outstanding amount.
- Generate receipt on verified.
- Audit.
- Publish verified/rejected events.

Out of Scope:
- Void/refund.
- Gateway settlement.

Rules:
- Bendahara scope required.
- Use transaction.
- Use decimal.
- Receipt number generated.

Acceptance Criteria:
- Payment verified updates bill.
- Payment rejected stores reason.
- Receipt created.
- Audit/event recorded.

Tests Required:
- Verify tests.
- Reject tests.
- Bill status tests.
- Audit/event tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 5.8 — Implement Receipt Download

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Receipt Download

Target:
finance-service

Goal:
Generate and retrieve payment receipt metadata/PDF placeholder.

Scope:
- payment_receipts.
- Receipt number.
- Receipt detail endpoint.
- Optional PDF file metadata.
- Parent and Bendahara scoped access.

Out of Scope:
- Advanced PDF design.
- Email receipt.

Rules:
- Receipt immutable snapshot.
- Access scoped.
- Audit download if restricted.

Acceptance Criteria:
- Receipt created after verified payment.
- Authorized user can view.
- Unauthorized rejected.

Tests Required:
- Receipt tests.
- Access tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 5.9 — Implement Void Payment with Approval

## Prompt

```text
You are working on `school-platform`.

Task:
Implement Void Payment with Approval

Target:
finance-service

Goal:
Request and approve void payment.

Scope:
- Void request.
- Approval request.
- Approve/reject void.
- Recalculate bill outstanding.
- Audit.
- Publish void events.

Out of Scope:
- Refund processing.
- Gateway reversal.

Rules:
- Void requires approval.
- Requester cannot bypass approval.
- Reason required.
- Use transaction.

Acceptance Criteria:
- Void request created.
- Approval voids payment.
- Bill recalculated.
- Rejection keeps payment verified.
- Audit/event recorded.

Tests Required:
- Void request tests.
- Approval tests.
- Recalculate tests.
- Permission tests.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

# Task 5.10 — Sprint 5 Final Verification

## Prompt

```text
You are working on `school-platform`.

Task:
Sprint 5 Final Verification

Target:
finance-service

Goal:
Verify Finance/SPP end-to-end flow.

Scope:
- Run fee → policy → bill → proof → verify → receipt → void flow.
- Verify scope and calculations.
- Produce report.

Out of Scope:
- Payment gateway.

Rules:
- No float usage.
- No cross-service DB query.
- No unauthorized bill access.

Acceptance Criteria:
- Finance core flow works.
- No Critical/High bug.

Tests Required:
- Full finance tests.
- Manual smoke test.

Expected AI Agent Output:
- Summary of implementation.
- Files created/modified.
- Tests added.
- Commands to run tests.
- Notes about migrations/contracts/events if changed.
```

---

## Final Planning Context Before Implementation

Before using this task prompt for coding, read:

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/25-product-requirement-document.md
docs/26-development-plan.md
docs/27-workflow.md
docs/25-github-project-management.md
docs/09-ai-agent-rules.md
docs/08-coding-standard.md
docs/34-sprint-5-plan.md if it exists
docs/18-sprint-5-task-prompts.md
```

Rules:

- Use final PRD as product scope reference.
- Use Development Plan as sprint and delivery reference.
- Use Workflow as daily SOP reference.
- Use GitHub Project Management guide for issue/PR/QA/release tracking.
- Implement only one selected issue/task at a time.

## Sprint Plan and GitHub Setup Context

Before using this task prompt for implementation, read:

```text
docs/34-sprint-5-plan.md
docs/40-github-repository-setup-labels-project.md
docs/25-github-project-management.md
docs/27-workflow.md
```

Use `docs/34-sprint-5-plan.md` for sprint-level issue plan, dependencies, risks, and exit criteria.

Use `docs/40-github-repository-setup-labels-project.md` for labels, milestones, project fields, and repository workflow alignment.
