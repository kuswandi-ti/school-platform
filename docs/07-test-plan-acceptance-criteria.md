# 07 — Test Plan and Acceptance Criteria

Project: `school-platform`  
Status: Final decision for MVP  
Scope: Testing strategy, acceptance criteria, QA workflow, UAT checklist, and AI Agent Definition of Done.

---

## 1. Purpose

This document defines the MVP testing strategy and acceptance criteria standards.

Testing is mandatory because the system handles:

- student personal data
- parent/guardian data
- finance and payment data
- report cards
- sensitive documents
- multi-school access
- multi-role authorization
- approval workflow
- audit log
- event-driven projection and notification

---

## 2. Core Decision

Every MVP module and feature must have:

```text
test plan
acceptance criteria
permission/scope test
Definition of Done
```

before implementation.

QA sign-off is mandatory before production release.

---

## 3. Testing Levels

### 3.1 Unit Test

Used for isolated business logic.

Examples:

```text
- Calculate SPP discount
- Calculate final bill amount
- Validate fee policy period
- Validate status transition
- Generate document number
- Mask sensitive data
```

Required for:

```text
Finance
Academic grading
Approval
Numbering
Permission/scope
```

---

### 3.2 Integration Test

Used for service + database/dependency behavior.

Examples:

```text
- Create student in school_core_db
- Generate bill in finance_db
- Verify payment and update bill status
- Publish report card and lock it
- Import Excel and create import rows
```

Integration tests must use test database/environment, not development or production data.

---

### 3.3 API Test

Used for external REST API via API Gateway.

Must verify:

```text
HTTP status code
response format
error format
permission check
scope check
validation error
object-level authorization
```

Examples:

```text
POST /api/v1/auth/login
GET /api/v1/students
POST /api/v1/finance/bills/generate
POST /api/v1/finance/payments/{id}/verify
POST /api/v1/academic/report-cards/{id}/publish
```

---

### 3.4 Permission and Scope Test

This is a top priority.

Required cases:

```text
Admin Yayasan can access foundation scope.
Kepala Sekolah can only access own school.
TU/Staff can manage data only within school scope.
Bendahara can access finance only within school scope.
Guru can only access assigned class/subject.
Orang Tua can only access linked children.
Siswa can only access self data.
```

Negative examples:

```text
Guru cannot view student payment details.
Bendahara cannot publish report cards.
Orang Tua cannot access another child.
Kepala Sekolah SD cannot edit SMP data.
```

---

### 3.5 Event Test

Required for event-driven workflows.

Must verify:

```text
event_type
event_id
event_version
source_service
correlation_id
foundation_id
school_id if relevant
payload safety
no Confidential data detail
```

Important events:

```text
school.student.created
finance.bill.generated
finance.payment.verified
academic.attendance.marked
academic.report_card.published
communication.announcement.published
approval.request.created
```

---

### 3.6 UI Flow Test

Used for user-facing workflows.

Main flows:

```text
Login
Manage students/teachers/classes
Import Excel
PPDB
Generate bills
Upload payment proof
Verify payment
Input attendance
Input score
Publish report card
Announcement
Dashboard
```

MVP does not require full automated E2E coverage at the beginning, but manual checklist is required.

---

### 3.7 Regression Test

Required before staging-to-main release.

Regression checklist must cover:

```text
Auth
Role/scope
School Core
Import Excel
PPDB
Finance/SPP
Academic/report card
Notification
Reporting
File access
Approval
Audit log
```

---

### 3.8 Security Baseline Test

Required before production release.

Must include:

```text
No cross-school data leak
No parent-to-other-child access
No teacher-to-unassigned-class access
No public access to private files
No token/password in logs
Rate limit for sensitive endpoints
Signed URL authorization
Object-level authorization
```

---

## 4. Acceptance Criteria Format

Acceptance criteria should use Given-When-Then or checklist.

Example:

```text
Given Bendahara berada di sekolah SD dan ada pembayaran pending_verification
When Bendahara memverifikasi pembayaran
Then payment status menjadi verified
And bill status menjadi paid atau partially_paid
And receipt dibuat
And audit log finance.payment.verified tercatat
And event finance.payment.verified dipublish
And notifikasi dikirim ke orang tua
```

Negative example:

```text
Given Guru mencoba membuka detail pembayaran siswa
When request dikirim
Then sistem mengembalikan 403 Forbidden atau 404 Not Found sesuai policy
And tidak ada data pembayaran dikirim
```

---

## 5. Module Test Plan

## 5.1 Identity & Access

Acceptance criteria:

```text
- User can login with valid credential.
- Access token is issued.
- Refresh token is stored as hash.
- Refresh token is rotated.
- Logout revokes session.
- User receives correct roles, permissions, and scopes.
- Login failure is rate-limited.
```

Required tests:

```text
Login success
Login wrong password
Login inactive user
Refresh token rotation
Refresh token reuse detection
Logout revoke token
Get current user context
Role assignment scope
Permission check
```

Negative/security tests:

```text
Unauthenticated user cannot access protected endpoint.
Expired token is rejected.
Revoked refresh token cannot be used.
School A user cannot access School B data.
```

---

## 5.2 School Core

Acceptance criteria:

```text
- Admin/TU can manage student within scope.
- Student stores foundation_id and school_id.
- Guardian is linked to student.
- Teacher and class can be managed.
- Student-class assignment works.
- Homeroom assignment works.
```

Required tests:

```text
Create student
Update student
Search/filter student
Create guardian
Link guardian to student
Create teacher
Create class
Assign student to class
Assign homeroom teacher
Student status change
```

Negative/scope tests:

```text
Kepala Sekolah SD cannot edit SMP student.
Guru cannot edit student master data.
Parent cannot view another student's data.
```

---

## 5.3 Import Excel

Acceptance criteria:

```text
- User can download template.
- User can upload Excel file.
- System validates columns and rows.
- System displays preview.
- Import runs only after confirmation.
- Import produces import report.
- Import file is private Restricted data.
```

Required tests:

```text
Upload valid file
Upload invalid extension
Missing required columns
Duplicate student number
Invalid birth date
Unknown class_code
Preview valid rows
Confirm import
Download error report
```

Negative tests:

```text
Guru cannot import student data.
Import file is not public.
Raw import data is not written to application log.
```

---

## 5.4 Admission / PPDB

Acceptance criteria:

```text
- Admission period can be created.
- Applicant can be submitted.
- PPDB document can be uploaded.
- TU can verify document.
- Kepala Sekolah can accept/reject applicant.
- Accepted applicant can be converted to student.
- School Core becomes owner of student after conversion.
```

Required tests:

```text
Create admission period
Submit applicant
Upload document
Verify document
Accept applicant
Reject applicant
Convert applicant to student
Prevent double conversion
```

Negative/scope tests:

```text
TU cannot final-approve if policy requires Kepala Sekolah.
Kepala Sekolah SD cannot approve SMP applicant.
Convert applicant without acceptance is rejected.
```

---

## 5.5 Finance / SPP

Acceptance criteria:

```text
- Bendahara can create fee type/scheme.
- Fee policy can be submitted and approved.
- Bill generation applies approved fee policy.
- Bill stores base_amount, discount_amount, final_amount, applied_policy snapshot.
- Parent can view own child's bills.
- Parent can upload payment proof.
- Bendahara can verify/reject payment.
- Receipt is created after payment verification.
- Void/refund requires approval.
```

Required tests:

```text
Create fee type
Create fee scheme
Create free_spp policy
Create sibling discount policy
Approve fee policy
Generate bills
Generate bills idempotency
Upload payment proof
Verify payment
Reject payment
Receipt generated
Void payment request
Void payment approval
Outstanding amount recalculated
```

Negative/security tests:

```text
Guru cannot view payments.
Parent cannot view another child's bill.
Bendahara cannot void without approval.
Duplicate generate bill request with same Idempotency-Key does not duplicate bills.
Finance calculation does not use float.
```

---

## 5.6 Academic / Attendance / Grade / Report Card

Acceptance criteria:

```text
- Guru can see assigned class/subject.
- Guru can input attendance for assigned class.
- Guru can input score.
- Guru can submit grade.
- Wali Kelas can review report summary.
- Kepala Sekolah can publish report card.
- Published report card is locked.
- Revision after publish requires approval and audit log.
```

Required tests:

```text
Create subject
Create schedule
Mark attendance
Correct attendance if allowed
Create grade book
Input score
Submit grade book
Review report card
Publish report card
Generate report PDF
Revision request after publish
```

Negative/scope tests:

```text
Guru cannot input score for unassigned class.
Guru cannot edit score after report card is published.
Parent can only view published report card.
Student cannot view another student's report card.
```

---

## 5.7 Communication / Notification

Acceptance criteria:

```text
- Announcement can be created within scope.
- Announcement can be published by authorized role.
- Notification is created from event.
- In-app notification is stored.
- FCM is sent if user has active device token.
- Email is sent only for selected events.
- Confidential data is not sent in notification body.
```

Required tests:

```text
Create announcement
Publish announcement
Target announcement by role/school/class
Notification created
Notification read/unread
FCM delivery mocked
Email delivery mocked
Notification preference respected
Critical notification cannot be fully disabled
```

Negative tests:

```text
Guru cannot publish foundation announcement.
BK/UKS notification must not include Confidential detail.
```

---

## 5.8 Reporting Dashboard

Acceptance criteria:

```text
- Dashboard reads from reporting_db.
- Reporting updates from event projection.
- Dashboard is near real-time.
- Scheduled rebuild can repair summary.
```

Required tests:

```text
StudentCreated updates student summary
PaymentVerified updates finance summary
AttendanceMarked updates attendance summary
ReportCardPublished updates academic progress
Foundation dashboard returns aggregate metrics
School dashboard returns scoped metrics
```

Negative/scope tests:

```text
Kepala Sekolah only sees own school dashboard.
Parent only sees linked child dashboard.
Reporting Service must not query operational service databases directly.
```

---

## 5.9 File Management

Acceptance criteria:

```text
- File is private by default.
- Upload validates MIME, extension, and size.
- File metadata is stored.
- Signed URL is generated after authorization.
- Restricted/Confidential download is audited.
- Official file is not overwritten.
```

Required tests:

```text
Upload valid PDF
Reject invalid extension
Reject oversized file
Generate signed URL
Signed URL expiry follows classification
Audit restricted file download
Version official document
Soft delete/archive file
```

Negative/security tests:

```text
Out-of-scope user cannot generate signed URL.
Confidential file cannot be accessed without special permission.
File does not have permanent public URL.
```

---

## 5.10 Approval

Acceptance criteria:

```text
- Approval request can be created for sensitive action.
- Authorized approver can approve/reject/request revision.
- Status changes according to workflow.
- Reason is required for selected actions.
- Audit log is recorded.
```

Required tests:

```text
Create approval request
Approve request
Reject request
Request revision
Re-submit revised request
Multi-level structure prepared but limited
```

Negative tests:

```text
Requester cannot approve own request if policy forbids it.
Approver outside school scope cannot approve school-level request.
Sensitive action cannot be finalized without approval.
```

---

## 5.11 Audit Log

Acceptance criteria:

```text
- Sensitive action creates audit log.
- Audit log contains actor, action, entity, request_id, correlation_id.
- Old/new values are stored if relevant.
- Sensitive data is masked.
- Application log does not contain raw sensitive data.
```

Required tests:

```text
Role change creates audit log
Payment verified creates audit log
Report card published creates audit log
Restricted file download creates audit log
Audit includes correlation_id
Sensitive values are masked
```

---

## 6. Release Acceptance Criteria

Before production release:

```text
CI pass
Migration reviewed
No open Critical/High bug
QA sign-off
Security/scope tests pass for changed features
Backup/snapshot performed if major migration exists
Release notes created
Rollback plan available
Manual production approval completed
```

MVP pilot acceptance:

```text
- At least one pilot school unit can run core workflow.
- Students/teachers/classes can be managed.
- PPDB can run.
- Manual SPP billing/payment can run.
- Attendance and basic report card can run.
- Dashboard appears.
- Backup and restore test has been performed at least once.
```

---

## 7. Bug Severity

### Critical

```text
Data leak across school/user
Major payment/billing calculation error
Mass login failure
Production down
Student/report/payment data loss
```

### High

```text
Main flow broken
Approval not working
Verified payment does not update bill
Report card cannot be published
Important dashboard metric wrong
```

### Medium

```text
Validation issue
UI issue with workaround
Export format issue
Partial notification failure
```

### Low

```text
Typo
Spacing
Minor UI issue
Unclear label
```

Production release rule:

```text
Production must not be released if any Critical/High bug remains open in MVP core flows.
```

---

## 8. Definition of Done for AI Agent Tasks

Every AI Agent task is done only when:

```text
Code is implemented within scope.
Architecture is not changed without instruction.
Tests are added.
Relevant test/lint passes.
Permission and scope checks exist.
Audit log is added for sensitive actions.
Event is published if required.
Response/error format is consistent.
No cross-service database query exists.
foundation_id/school_id are not hardcoded.
Sensitive data is not written to logs.
OpenAPI/proto/event docs are updated if changed.
```

---

## 9. UAT Checklist MVP

Minimum UAT checklist:

```text
1. Admin Yayasan login and views foundation dashboard.
2. Admin/Kepala Sekolah selects active academic year and semester.
3. TU imports student, teacher, and class data.
4. TU verifies import results.
5. PPDB creates new applicant.
6. Kepala Sekolah accepts applicant and converts to student.
7. Bendahara generates SPP bill.
8. Parent views bill and uploads payment proof.
9. Bendahara verifies payment.
10. System creates receipt.
11. Guru inputs attendance.
12. Guru inputs grade.
13. Kepala Sekolah publishes report card.
14. Parent/student views published report card.
15. Kepala Sekolah views school dashboard.
16. Admin Yayasan views cross-unit summary.
17. Sensitive actions are recorded in audit log.
```

---

## 10. Final Summary

Testing is part of design, not an afterthought.

Priority testing areas:

```text
role/scope authorization
finance
report card
approval
audit log
file access
event processing
data privacy
```

QA sign-off is mandatory before production release.
