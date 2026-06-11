# 06 — UI Screen List and User Flow

Project: `school-platform`  
Status: Final decision for MVP  
Scope: Web and mobile MVP screens, role-based navigation, global context, and main end-to-end flows.

---

## 1. Purpose

This document defines the MVP UI screen list and user flows for the school foundation management system.

Screens are organized by:

- role
- platform
- workflow
- permission/scope
- active context
- API dependency
- sensitive action requirements

This document is intended to guide frontend implementation and AI Agent task breakdown.

---

## 2. Core Decision

MVP screens must be documented before major frontend implementation.

Platform split:

```text
Web:
- Admin Yayasan
- Kepala Sekolah
- Tata Usaha / Staff
- Bendahara Sekolah
- Guru / Wali Kelas

Mobile:
- Orang Tua / Wali Murid
- Siswa
- Guru for quick actions
```

Web is the primary channel for administrative and operational work.

Mobile is the primary channel for fast access, notification, billing visibility, attendance visibility, announcement, published grade/report card, and quick teacher workflows.

---

## 3. Global Context

The system must support global context.

Common context:

```text
selected_foundation
selected_school
selected_academic_year
selected_semester
```

Role-specific context:

```text
selected_class          → Guru / Wali Kelas
selected_subject        → Guru
selected_child          → Orang Tua
selected_billing_month  → Bendahara / Finance
```

Rules:

- show context visibly on relevant pages
- do not trust frontend context blindly
- backend must validate context against user scope
- switcher only appears if user has more than one valid option

Example page context label:

```text
SD ABC · Tahun Ajaran 2026/2027 · Semester Ganjil
```

---

## 4. Common Web Screens

Common screens for authenticated web users:

```text
1. Login
2. Forgot Password
3. Reset Password
4. Select Role / Switch Role
5. Select School / Switch School
6. Profile
7. Change Password
8. Notification Center
9. Forbidden / 403
10. Not Found / 404
```

Role/school switcher appears only when applicable.

---

## 5. Admin Yayasan — Web Screens

Admin Yayasan has foundation-level visibility.

MVP screens:

```text
1. Dashboard Yayasan
2. Daftar Unit Sekolah
3. Detail Unit Sekolah
4. Manajemen User
5. Manajemen Role & Permission
6. Tahun Ajaran
7. Semester
8. Kalender Yayasan
9. Data Siswa Lintas Unit
10. Data Guru/Staff Lintas Unit
11. PPDB Summary Lintas Unit
12. Finance Summary Lintas Unit
13. Fee Policy / Diskon / Beasiswa Lintas Unit
14. Approval Center Yayasan
15. Pengumuman Yayasan
16. Laporan Yayasan
17. Pengaturan Branding
18. Audit Log Ringkas / Aktivitas Penting
```

Main actions:

```text
- View cross-unit dashboard
- Manage school units
- Manage users and sensitive roles
- Activate academic year/semester
- View aggregate reports
- Approve foundation-level policy
- Publish foundation announcement
- Manage configurable branding
```

---

## 6. Kepala Sekolah — Web Screens

Kepala Sekolah operates within school scope.

MVP screens:

```text
1. Dashboard Sekolah
2. Data Siswa
3. Detail Siswa
4. Data Guru
5. Detail Guru
6. Kelas/Rombel
7. Detail Kelas
8. Kalender Sekolah
9. PPDB Sekolah
10. Detail Pendaftar PPDB
11. Approval PPDB
12. Keuangan Ringkas Sekolah
13. Approval Keuangan
14. Akademik Sekolah
15. Jadwal Pelajaran
16. Absensi Siswa
17. Progress Input Nilai
18. Review dan Publish Rapor
19. Pengumuman Sekolah
20. Surat/Dokumen Sekolah jika masuk MVP
21. Approval Center Sekolah
22. Laporan Sekolah
```

Main actions:

```text
- approve/reject PPDB
- approve financial sensitive actions
- publish report cards
- publish school announcements
- monitor attendance
- monitor academic progress
- review school reports
```

---

## 7. Tata Usaha / Staff — Web Screens

TU/Staff manages school administration data.

MVP screens:

```text
1. Dashboard Staff
2. Data Siswa
3. Tambah/Edit Siswa
4. Detail Siswa
5. Data Orang Tua/Wali
6. Data Guru
7. Tambah/Edit Guru
8. Kelas/Rombel
9. Assignment Siswa ke Kelas
10. Assignment Guru/Wali Kelas
11. Import Data
12. Import Preview
13. Import Result/Error Report
14. PPDB List
15. Verifikasi Berkas PPDB
16. Detail Pendaftar PPDB
17. Dokumen Siswa
18. Kalender Sekolah
19. Surat/Dokumen Administrasi jika masuk MVP
```

Main actions:

```text
- create/update student
- create/update teacher
- manage guardian data
- manage classes
- assign students to class
- import Excel data
- verify PPDB documents
- manage student documents
```

---

## 8. Bendahara Sekolah — Web Screens

Bendahara manages finance and payment operations.

MVP screens:

```text
1. Dashboard Keuangan
2. Jenis Tagihan
3. Skema Tagihan
4. Fee Policy / Diskon / Beasiswa
5. Sibling Discount Rules
6. Generate Tagihan
7. Preview Generate Tagihan
8. Daftar Tagihan
9. Detail Tagihan
10. Daftar Pembayaran
11. Detail Pembayaran
12. Verifikasi Bukti Pembayaran
13. Input Pembayaran Manual/Tunai
14. Daftar Tunggakan
15. Kwitansi/Receipt
16. Request Void/Refund
17. Laporan Keuangan
18. Export Laporan
19. Rekonsiliasi Manual/Semi-manual
```

Main actions:

```text
- create fee type/scheme
- manage fee policy
- generate bills
- verify/reject payment proof
- input manual cash payment
- generate receipt
- request void/refund
- view outstanding bills
- export finance reports
```

---

## 9. Guru / Wali Kelas — Web Screens

Guru uses web for academic work that requires tables/forms.

MVP screens:

```text
1. Dashboard Guru
2. Jadwal Mengajar
3. Kelas Saya
4. Detail Kelas
5. Input Absensi
6. Riwayat Absensi
7. Grade Book / Buku Nilai
8. Input Nilai
9. Submit Nilai
10. Review Rapor Kelas untuk Wali Kelas
11. Catatan Wali Kelas
12. Pengumuman Kelas/Sekolah
13. Profil Siswa Terbatas sesuai scope
```

Main actions:

```text
- view teaching schedule
- input attendance
- input scores
- submit grade book
- review class report summary as homeroom teacher
- add homeroom note
- view scoped student profile
```

Restrictions:

```text
Guru must not see student finance details by default.
Guru must not see Confidential BK/UKS data without special permission.
```

---

## 10. Common Mobile Screens

Common mobile screens:

```text
1. Login
2. Forgot Password
3. Reset Password
4. Home
5. Notification Center
6. Profile
7. Change Password
8. Switch Child
9. Offline/No Internet State
```

Offline write is not part of MVP.

---

## 11. Orang Tua / Wali Murid — Mobile Screens

MVP screens:

```text
1. Home Orang Tua
2. Switch Anak
3. Profil Anak
4. Tagihan Anak
5. Detail Tagihan
6. Upload Bukti Pembayaran
7. Riwayat Pembayaran
8. Kwitansi/Receipt
9. Absensi Anak
10. Nilai/Rapor Published
11. Detail Rapor
12. Pengumuman
13. Detail Pengumuman
14. Notifikasi
```

Main actions:

```text
- view child summary
- view bills
- upload payment proof
- view payment status
- view attendance
- view published score/report card
- read announcements
```

Scope rule:

```text
Parent can only access linked children.
```

---

## 12. Siswa — Mobile Screens

MVP screens:

```text
1. Home Siswa
2. Jadwal Hari Ini
3. Absensi Saya
4. Nilai/Rapor Published
5. Detail Rapor
6. Pengumuman
7. Notifikasi
8. Profil Saya
```

Main actions:

```text
- view schedule
- view own attendance
- view published report card
- read announcements
```

Scope rule:

```text
Student can only access self data.
```

---

## 13. Guru — Mobile Screens

Mobile guru is limited to quick workflows.

MVP screens:

```text
1. Home Guru
2. Jadwal Hari Ini
3. Pilih Kelas
4. Input Absensi Cepat
5. Daftar Kelas Saya
6. Pengumuman
7. Notifikasi
```

Main actions:

```text
- view today schedule
- input quick attendance
- view assigned classes
- read announcements
- receive notifications
```

Not included in MVP mobile:

```text
complex grade input
report card editing
advanced academic analytics
```

---

## 14. Main User Flows

### 14.1 Login and Context Selection

```text
User opens app
→ Login
→ API Gateway validates credential through Identity Service
→ System returns role, permission, and scope
→ If multi-role/multi-school, show switcher
→ User selects active context
→ User enters role-specific dashboard
```

Context:

```text
selected_school
selected_academic_year
selected_semester
selected_role
selected_child if parent
```

---

### 14.2 Manage Students, Teachers, and Classes

```text
TU/Staff opens Data Siswa/Guru/Kelas
→ Filter by school, academic year, class, status
→ Create/update data
→ Backend validates permission and scope
→ School Core stores data
→ Audit log is created
→ Domain event is published
→ Reporting projection is updated
```

---

### 14.3 Import Excel Initial Data

```text
TU/Staff opens Import Data
→ Download template
→ Upload Excel file
→ System validates structure and rows
→ System displays preview: valid/warning/error
→ User confirms import
→ School Core processes import
→ Import report is created
→ Audit log is created
```

Rule:

```text
Import must not insert data directly without validation and preview.
```

---

### 14.4 PPDB to Student Conversion

```text
Admin/TU creates admission period
→ Applicant is registered
→ PPDB documents are uploaded
→ TU/Staff verifies documents
→ Kepala Sekolah accepts/rejects applicant
→ Accepted applicant is converted to student
→ Admission Service calls School Core via gRPC
→ School Core creates student and guardian
→ Event is published
→ Reporting and Notification are updated
```

---

### 14.5 Generate SPP Bills

```text
Bendahara opens Generate Tagihan
→ Select school, academic year, billing month, fee type
→ System displays preview
→ Preview shows normal, discount, free_spp, total bill
→ Bendahara confirms generate
→ Finance Service generates bills
→ Bill stores fee policy snapshot
→ finance.bill.generated event is published
→ Parent receives notification
```

Rule:

```text
Mass bill generation must use Idempotency-Key.
```

---

### 14.6 Manual Payment and Proof Upload

```text
Parent opens child bill
→ Parent transfers manually
→ Parent uploads payment proof
→ Payment status becomes pending_verification
→ Bendahara sees pending payment
→ Bendahara verifies or rejects
→ If verified, bill status is updated and receipt is created
→ Audit log is created
→ Notification is sent
```

---

### 14.7 Payment Void with Approval

```text
Bendahara opens payment detail
→ Click Request Void
→ Enter reason
→ Approval request is created
→ Kepala Sekolah/Admin Yayasan approves or rejects
→ If approved, payment becomes voided
→ Bill outstanding amount is recalculated
→ Audit log is created
→ Reporting projection is updated
```

Rule:

```text
Void payment always requires approval and audit log.
```

---

### 14.8 Teacher Attendance Input

```text
Guru opens schedule/class
→ Select class/subject
→ Input student attendance
→ Submit
→ Academic Service validates teacher assignment
→ Attendance is saved
→ academic.attendance.marked event is published
→ Reporting updates attendance summary
→ If absent, notification may be sent to parent
```

---

### 14.9 Grade Input and Report Card Publish

```text
Guru opens Grade Book
→ Input score and description
→ Submit grade book
→ Wali Kelas reviews summary
→ Kepala Sekolah publishes report card
→ Report card status becomes published/locked
→ PDF may be generated
→ Parent/student can view published report card
→ academic.report_card.published event is published
```

Rule:

```text
Revision after publish requires approval and audit log.
```

---

### 14.10 Announcement and Notification

```text
Authorized user creates announcement
→ User selects target scope
→ If required, approval/publish flow runs
→ Communication Service publishes announcement
→ Notification is created for target recipients
→ In-app notification is stored
→ FCM is sent if device token exists
→ Email is sent only for selected important events
```

Rule:

```text
Confidential data must not be sent in notification body.
```

---

### 14.11 Reporting Dashboard

```text
User opens dashboard
→ Frontend calls API Gateway
→ API Gateway calls Reporting Service
→ Reporting Service reads reporting_db projection
→ Dashboard displays near real-time metrics
```

Rule:

```text
Dashboard must not query operational service databases directly.
```

---

## 15. UI Component Standards

Web components:

```text
Page Header
Breadcrumb
Global Context Bar
Filter Bar
Data Table
Pagination
Status Badge
Action Button
Confirmation Dialog
Reason Modal
Empty State
Loading State
Error State
Success Toast
Permission Guard
```

Mobile components:

```text
App Bar
Bottom Navigation
Card Summary
List Item
Filter/Chip
Action Sheet
Upload Component
Notification Badge
Empty State
Loading State
Error State
```

---

## 16. Page State Standards

Every screen must handle:

```text
loading
empty
error
success
forbidden
unauthorized/session expired
```

Examples:

```text
No active bills → "Belum ada tagihan aktif."
No permission → hide action or show 403.
Session expired → redirect to login.
```

---

## 17. Sensitive Action UI Pattern

Sensitive actions must use confirmation dialog and reason when required.

Sensitive actions include:

```text
Change role
Deactivate user
Approve/reject PPDB
Void payment
Refund
Create/update fee policy
Publish report card
Revise report card
Export Restricted/Confidential data
Download Confidential file
```

Pattern:

```text
Click action
→ Confirmation modal
→ Display impact summary
→ Require reason if needed
→ Submit
→ Approval/audit flow
```

---

## 18. Screen Specification Template

Every screen spec should use this format:

```text
Screen Name:
Role:
Platform:
Route:
Purpose:
Global Context:
Permissions:
Data/API:
Main Components:
Actions:
Validation:
Empty State:
Error State:
Audit Requirement:
Notes:
```

Example:

```text
Screen Name: Verifikasi Pembayaran
Role: Bendahara Sekolah
Platform: Web
Route: /finance/payments/:payment_id/verify
Purpose: Memverifikasi bukti pembayaran orang tua
Global Context: selected_school, selected_academic_year, selected_billing_month
Permissions: finance.payment.verify
Data/API:
- GET /api/v1/finance/payments/{payment_id}
- POST /api/v1/finance/payments/{payment_id}/verify
Actions:
- Verify
- Reject with reason
Audit Requirement:
- finance.payment.verified
- finance.payment.rejected
```

---

## 19. Out of MVP Screens

Do not implement these in MVP:

```text
Payroll detail
HR lengkap
Asset inventory lengkap
Library/perpustakaan lengkap
BK detail
UKS detail
LMS penuh
Alumni/tracer
Koperasi
Payment gateway
WhatsApp settings
Global search
Offline sync management
Advanced analytics
```

---

## 20. Final Summary

MVP UI uses:

```text
Web for administrative and operational work.
Mobile for quick access and non-admin users.
Screen list is role-based and workflow-based.
Every screen must define permission, context, API, action, and state.
Sensitive actions require confirmation, reason if needed, approval flow, and audit log.
```
