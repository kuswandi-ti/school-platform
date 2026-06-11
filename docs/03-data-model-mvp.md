# 03 — Data Model MVP

**Project:** `school-platform`  
**Product Context:** Sistem Manajemen Sekolah/Yayasan untuk TK, SD, SMP, dan SMA  
**Architecture:** Microservice, database per service, event-driven projection  
**Status:** Final architecture decision draft  
**Last Updated:** 2026-06-08

---

## 1. Tujuan Dokumen

Dokumen ini mendefinisikan **data model MVP** untuk sistem manajemen sekolah/yayasan berbasis microservice.

Dokumen ini menjadi acuan untuk:

- desain database per service;
- migration awal setiap service;
- query SQLC;
- API/gRPC contract;
- event contract;
- audit, approval, file metadata, numbering, dan import data;
- task implementation oleh developer dan AI Agent.

Data model dalam dokumen ini **bukan ERD monolitik**, melainkan **ERD konseptual per service** sesuai keputusan service boundary.

---

## 2. Prinsip Data Model MVP

Keputusan utama:

```text
Data model MVP dibuat per service sesuai data ownership.
Setiap service memiliki database sendiri.
Tidak ada foreign key lintas database service.
Relasi lintas service memakai reference ID dan divalidasi melalui gRPC atau event.
```

Prinsip wajib:

1. Setiap tabel domain utama menggunakan `UUID` sebagai primary key teknis.
2. `foundation_id` menjadi tenant boundary utama.
3. `school_id` digunakan untuk data yang terkait unit sekolah.
4. Service hanya boleh mengakses database miliknya sendiri.
5. Tidak ada query langsung ke database service lain.
6. Relasi lintas service menggunakan reference ID.
7. Data historis penting wajib menyimpan snapshot.
8. Tabel `audit_logs`, `approval_requests`, `numbering_sequences`, `files`, dan `import_batches` dibuat di service pemilik domain sesuai kebutuhan.
9. Struktur tabel standar distandarkan melalui shared library/package.
10. Semua data sensitif mengikuti klasifikasi: `public`, `internal`, `restricted`, `confidential`.

---

## 3. Database MVP

Database MVP:

```text
identity_db
school_core_db
admission_db
academic_db
finance_db
communication_db
reporting_db
```

Mapping database ke service:

| Service | Database | Fungsi |
|---|---|---|
| `identity-service` | `identity_db` | User, credential, session, role, permission, role assignment |
| `school-core-service` | `school_core_db` | Yayasan, sekolah, tahun ajaran, semester, siswa, wali, guru, kelas, assignment |
| `admission-service` | `admission_db` | PPDB dan applicant sebelum dikonversi menjadi siswa |
| `academic-service` | `academic_db` | Kurikulum, mapel, jadwal, absensi, nilai, rapor |
| `finance-service` | `finance_db` | Fee type, fee scheme, fee policy, tagihan, pembayaran, receipt, rekonsiliasi |
| `communication-service` | `communication_db` | Pengumuman, notifikasi, template, delivery log, preference, surat dasar |
| `reporting-service` | `reporting_db` | Projection/read model dashboard dan summary |

---

## 4. Shared Field Standard

### 4.1 Field Umum Domain Table

Tabel domain utama sebaiknya memiliki field dasar berikut:

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL,
deleted_at TIMESTAMP NULL
```

Catatan:

- `school_id` boleh `NULL` untuk data level yayasan/global.
- Untuk tabel yang hanya relevan di satu sekolah, `school_id` sebaiknya `NOT NULL`.
- `deleted_at` digunakan jika data mendukung soft delete.

### 4.2 Field Audit Ringan pada Domain Table

Untuk tabel penting, tambahkan:

```sql
created_by UUID NULL,
updated_by UUID NULL,
deleted_by UUID NULL,
delete_reason TEXT NULL
```

### 4.3 Field Status

Tabel operasional penting harus memiliki:

```sql
status VARCHAR(50) NOT NULL
```

Status enum disimpan dalam format `lowercase_snake_case`.

Contoh:

```text
active
pending_verification
revision_requested
partially_paid
```

### 4.4 Optimistic Locking

Untuk data rawan konflik, dapat ditambahkan:

```sql
version INT NOT NULL DEFAULT 1
```

Cocok untuk:

- report card;
- grade book;
- payment verification;
- fee policy;
- approval request.

---

## 5. Identity Service Data Model

Database: `identity_db`

Identity Service menjadi owner untuk:

- user account;
- credential;
- password hash;
- refresh token/session;
- device;
- role;
- permission;
- user role assignment.

Identity Service **bukan owner** data siswa, guru, atau wali secara domain sekolah.

---

### 5.1 `users`

```sql
id UUID PRIMARY KEY,
email VARCHAR(255) UNIQUE,
phone VARCHAR(50) NULL,
password_hash TEXT NOT NULL,
display_name VARCHAR(150) NOT NULL,
avatar_file_id UUID NULL,
status VARCHAR(50) NOT NULL,
last_login_at TIMESTAMP NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Status:

```text
active
inactive
locked
invited
```

Catatan:

- `users` menyimpan akun login.
- Detail siswa/guru/wali tetap berada di `school_core_db`.
- `avatar_file_id` adalah reference ke metadata file sesuai service owner atau file domain terkait.

---

### 5.2 `roles`

```sql
id UUID PRIMARY KEY,
code VARCHAR(100) UNIQUE NOT NULL,
name VARCHAR(150) NOT NULL,
description TEXT NULL,
is_system BOOLEAN NOT NULL DEFAULT false,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Role MVP:

```text
admin_yayasan
kepala_sekolah
tu_staff
bendahara_sekolah
guru
orang_tua
siswa
```

Catatan:

- `wali_kelas` bukan role utama, tetapi assignment tambahan di School Core/Academic.

---

### 5.3 `permissions`

```sql
id UUID PRIMARY KEY,
code VARCHAR(150) UNIQUE NOT NULL,
name VARCHAR(150) NOT NULL,
description TEXT NULL,
module VARCHAR(100) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Format permission:

```text
domain.resource.action
```

Contoh:

```text
school.student.view
school.student.create
finance.payment.verify
academic.report_card.publish
identity.role.assign
```

---

### 5.4 `role_permissions`

```sql
id UUID PRIMARY KEY,
role_id UUID NOT NULL,
permission_id UUID NOT NULL,
created_at TIMESTAMP NOT NULL
```

Unique constraint:

```text
role_id + permission_id
```

---

### 5.5 `user_role_assignments`

```sql
id UUID PRIMARY KEY,
user_id UUID NOT NULL,
role_id UUID NOT NULL,
foundation_id UUID NOT NULL,
school_id UUID NULL,
class_id UUID NULL,
student_id UUID NULL,
employee_id UUID NULL,
subject_id UUID NULL,
scope_json JSONB NULL,
is_active BOOLEAN NOT NULL DEFAULT true,
starts_at TIMESTAMP NULL,
ends_at TIMESTAMP NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Fungsi:

- menyimpan role + scope user;
- mendukung multi-role;
- mendukung multi-school;
- mendukung orang tua dengan scope anak;
- mendukung guru dengan class/subject scope.

Contoh:

```text
Kepala Sekolah SD:
- user_id = U001
- role = kepala_sekolah
- foundation_id = F001
- school_id = SD001
```

```text
Orang Tua:
- user_id = U002
- role = orang_tua
- foundation_id = F001
- student_id = S001/S002 melalui scope_json atau multiple assignment rows
```

---

### 5.6 `user_sessions`

```sql
id UUID PRIMARY KEY,
user_id UUID NOT NULL,
refresh_token_hash TEXT NOT NULL,
device_id UUID NULL,
ip_address VARCHAR(100) NULL,
user_agent TEXT NULL,
expires_at TIMESTAMP NOT NULL,
revoked_at TIMESTAMP NULL,
last_used_at TIMESTAMP NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Aturan:

- refresh token disimpan dalam bentuk hash;
- refresh token dirotasi setiap digunakan;
- session dapat direvoke per device/session;
- reuse refresh token harus dianggap suspicious.

---

### 5.7 `user_devices`

```sql
id UUID PRIMARY KEY,
user_id UUID NOT NULL,
device_uid VARCHAR(255) NOT NULL,
platform VARCHAR(50) NOT NULL,
device_name VARCHAR(150) NULL,
fcm_token TEXT NULL,
is_active BOOLEAN NOT NULL DEFAULT true,
last_seen_at TIMESTAMP NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Platform:

```text
web
android
ios
```

Catatan:

- Device/session ownership berada di Identity.
- Communication Service dapat request device token via gRPC atau consume event device registration.

---

### 5.8 `password_reset_tokens`

```sql
id UUID PRIMARY KEY,
user_id UUID NOT NULL,
token_hash TEXT NOT NULL,
expires_at TIMESTAMP NOT NULL,
used_at TIMESTAMP NULL,
created_at TIMESTAMP NOT NULL
```

---

### 5.9 Identity Index Utama

```text
users.email
users.phone
users.status
user_role_assignments.user_id
user_role_assignments.foundation_id
user_role_assignments.school_id
user_sessions.user_id
user_sessions.expires_at
```

---

## 6. School Core Service Data Model

Database: `school_core_db`

School Core Service menjadi owner untuk:

- foundation;
- school/unit;
- academic year;
- semester;
- student master;
- guardian/parent master;
- teacher master dasar;
- grade level;
- class/rombel;
- student-class assignment;
- teacher assignment;
- homeroom assignment;
- room dasar.

---

### 6.1 `foundations`

```sql
id UUID PRIMARY KEY,
foundation_code VARCHAR(50) UNIQUE NOT NULL,
name VARCHAR(150) NOT NULL,
legal_name VARCHAR(200) NULL,
address TEXT NULL,
phone VARCHAR(50) NULL,
email VARCHAR(255) NULL,
logo_file_id UUID NULL,
timezone VARCHAR(100) NOT NULL DEFAULT 'Asia/Jakarta',
status VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Status:

```text
active
inactive
```

---

### 6.2 `schools`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_code VARCHAR(50) NOT NULL,
name VARCHAR(150) NOT NULL,
school_level VARCHAR(50) NOT NULL,
npsn VARCHAR(50) NULL,
address TEXT NULL,
phone VARCHAR(50) NULL,
email VARCHAR(255) NULL,
logo_file_id UUID NULL,
status VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Unique constraint:

```text
foundation_id + school_code
```

`school_level`:

```text
kindergarten
elementary
junior_high
senior_high
```

Contoh `school_code`:

```text
TK
SD
SMP
SMA
```

Jika ada lebih dari satu unit per jenjang:

```text
SD01
SD02
SMP01
```

---

### 6.3 `academic_years`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
name VARCHAR(50) NOT NULL,
start_date DATE NOT NULL,
end_date DATE NOT NULL,
status VARCHAR(50) NOT NULL,
is_active BOOLEAN NOT NULL DEFAULT false,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Status:

```text
draft
active
closed
archived
```

Catatan:

- Academic year berlaku global level yayasan pada MVP.

---

### 6.4 `semesters`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
academic_year_id UUID NOT NULL,
name VARCHAR(50) NOT NULL,
sequence INT NOT NULL,
start_date DATE NOT NULL,
end_date DATE NOT NULL,
status VARCHAR(50) NOT NULL,
is_active BOOLEAN NOT NULL DEFAULT false,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Status:

```text
draft
active
closed
archived
```

---

### 6.5 `students`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
student_number VARCHAR(100) NULL,
nisn VARCHAR(100) NULL,
full_name VARCHAR(150) NOT NULL,
gender VARCHAR(20) NOT NULL,
birth_place VARCHAR(100) NULL,
birth_date DATE NULL,
religion VARCHAR(50) NULL,
address TEXT NULL,
phone VARCHAR(50) NULL,
email VARCHAR(255) NULL,
photo_file_id UUID NULL,
status VARCHAR(50) NOT NULL,
entry_year INT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL,
deleted_at TIMESTAMP NULL
```

Status:

```text
active
inactive
transferred
graduated
dropped_out
```

Unique constraint:

```text
foundation_id + school_id + student_number
foundation_id + nisn, nullable jika NISN tersedia
```

Catatan penting:

```text
free_spp, diskon, beasiswa, dan sibling_discount bukan status siswa.
Kebijakan pembayaran dikelola di Finance Service melalui student_fee_policies.
```

---

### 6.6 `guardians`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
full_name VARCHAR(150) NOT NULL,
relationship_type VARCHAR(50) NOT NULL,
phone VARCHAR(50) NULL,
email VARCHAR(255) NULL,
address TEXT NULL,
occupation VARCHAR(100) NULL,
user_id UUID NULL,
status VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

`relationship_type`:

```text
father
mother
guardian
other
```

Catatan:

- `user_id` adalah reference ke Identity Service.
- Guardian domain data tetap milik School Core.

---

### 6.7 `student_guardians`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
student_id UUID NOT NULL,
guardian_id UUID NOT NULL,
relationship_type VARCHAR(50) NOT NULL,
is_primary BOOLEAN NOT NULL DEFAULT false,
can_login BOOLEAN NOT NULL DEFAULT true,
can_receive_notification BOOLEAN NOT NULL DEFAULT true,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Fungsi:

- relasi many-to-many siswa dan guardian;
- mendukung satu guardian memiliki lebih dari satu anak;
- dasar untuk sibling detection.

---

### 6.8 `teachers`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
employee_number VARCHAR(100) NULL,
user_id UUID NULL,
full_name VARCHAR(150) NOT NULL,
gender VARCHAR(20) NULL,
birth_place VARCHAR(100) NULL,
birth_date DATE NULL,
email VARCHAR(255) NULL,
phone VARCHAR(50) NULL,
address TEXT NULL,
photo_file_id UUID NULL,
employment_status VARCHAR(50) NULL,
status VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Status:

```text
active
inactive
resigned
```

Catatan:

- Data guru master dasar berada di School Core untuk MVP.
- HR detail dan payroll masuk phase lanjutan.

---

### 6.9 `grade_levels`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
level_code VARCHAR(50) NOT NULL,
name VARCHAR(100) NOT NULL,
sequence INT NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Contoh:

```text
TK-A
TK-B
1
2
3
...
12
```

---

### 6.10 `classes`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
academic_year_id UUID NOT NULL,
grade_level_id UUID NOT NULL,
class_code VARCHAR(50) NOT NULL,
name VARCHAR(100) NOT NULL,
capacity INT NULL,
room_id UUID NULL,
status VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Unique constraint:

```text
foundation_id + school_id + academic_year_id + class_code
```

Status:

```text
active
inactive
archived
```

---

### 6.11 `student_class_assignments`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
academic_year_id UUID NOT NULL,
semester_id UUID NULL,
student_id UUID NOT NULL,
class_id UUID NOT NULL,
status VARCHAR(50) NOT NULL,
assigned_at TIMESTAMP NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Status:

```text
active
moved
completed
```

---

### 6.12 `teacher_assignments`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
academic_year_id UUID NOT NULL,
semester_id UUID NOT NULL,
teacher_id UUID NOT NULL,
class_id UUID NULL,
subject_id UUID NULL,
assignment_type VARCHAR(50) NOT NULL,
status VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

`subject_id` adalah reference ke Academic Service.

`assignment_type`:

```text
subject_teacher
class_teacher
extracurricular_coach
```

---

### 6.13 `homeroom_assignments`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
academic_year_id UUID NOT NULL,
semester_id UUID NULL,
teacher_id UUID NOT NULL,
class_id UUID NOT NULL,
status VARCHAR(50) NOT NULL,
approved_by UUID NULL,
approved_at TIMESTAMP NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Status:

```text
active
inactive
```

---

### 6.14 `rooms`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
room_code VARCHAR(50) NOT NULL,
name VARCHAR(100) NOT NULL,
capacity INT NULL,
status VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

---

### 6.15 School Core Index Utama

```text
schools.foundation_id
schools.school_code
students.foundation_id + school_id
students.full_name
students.student_number
students.nisn
students.status
guardians.phone
guardians.email
teachers.foundation_id + school_id
teachers.full_name
classes.foundation_id + school_id + academic_year_id
student_class_assignments.student_id
student_class_assignments.class_id
```

---

## 7. Admission Service Data Model

Database: `admission_db`

Admission Service menjadi owner proses PPDB dan applicant sebelum diterima.

Setelah applicant diterima dan dikonversi menjadi siswa, School Core menjadi owner student record.

---

### 7.1 `admission_periods`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
academic_year_id UUID NOT NULL,
name VARCHAR(150) NOT NULL,
start_date DATE NOT NULL,
end_date DATE NOT NULL,
capacity INT NULL,
status VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Status:

```text
draft
open
closed
archived
```

---

### 7.2 `applicants`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
admission_period_id UUID NOT NULL,
registration_number VARCHAR(100) NOT NULL,
target_grade_level_id UUID NULL,
full_name VARCHAR(150) NOT NULL,
gender VARCHAR(20) NOT NULL,
birth_place VARCHAR(100) NULL,
birth_date DATE NULL,
previous_school VARCHAR(150) NULL,
status VARCHAR(50) NOT NULL,
converted_student_id UUID NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Status:

```text
draft
submitted
verified
accepted
rejected
converted
cancelled
```

Unique constraint:

```text
foundation_id + school_id + registration_number
```

Catatan:

- `converted_student_id` adalah reference ke School Core `students.id`.
- Tidak ada foreign key lintas database.

---

### 7.3 `applicant_guardians`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
applicant_id UUID NOT NULL,
full_name VARCHAR(150) NOT NULL,
relationship_type VARCHAR(50) NOT NULL,
phone VARCHAR(50) NULL,
email VARCHAR(255) NULL,
address TEXT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

---

### 7.4 `applicant_documents`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
applicant_id UUID NOT NULL,
file_id UUID NOT NULL,
document_type VARCHAR(100) NOT NULL,
status VARCHAR(50) NOT NULL,
uploaded_at TIMESTAMP NOT NULL,
verified_by UUID NULL,
verified_at TIMESTAMP NULL,
notes TEXT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Status:

```text
uploaded
verified
rejected
```

---

### 7.5 `applicant_verifications`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
applicant_id UUID NOT NULL,
verification_type VARCHAR(100) NOT NULL,
status VARCHAR(50) NOT NULL,
notes TEXT NULL,
verified_by UUID NOT NULL,
verified_at TIMESTAMP NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

---

### 7.6 `admission_decisions`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
applicant_id UUID NOT NULL,
decision VARCHAR(50) NOT NULL,
reason TEXT NULL,
decided_by UUID NOT NULL,
decided_at TIMESTAMP NOT NULL,
approval_request_id UUID NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Decision:

```text
accepted
rejected
revision_requested
```

---

### 7.7 Admission Index Utama

```text
admission_periods.foundation_id + school_id + academic_year_id
applicants.foundation_id + school_id + admission_period_id
applicants.registration_number
applicants.full_name
applicants.status
applicant_documents.applicant_id
```

---

## 8. Finance Service Data Model

Database: `finance_db`

Finance Service menjadi owner untuk:

- fee type;
- fee scheme;
- student fee policy;
- sibling discount rule;
- bill/invoice;
- bill item;
- payment;
- payment proof reference;
- receipt;
- reconciliation;
- finance approval.

Finance Service menggunakan reference ID dari School Core seperti `student_id`, `guardian_id`, `class_id`, `academic_year_id`, dan `semester_id`.

---

### 8.1 Money Handling

Keputusan:

```text
Finance calculation tidak boleh memakai float.
Database menggunakan NUMERIC(14,2).
Go menggunakan shopspring/decimal.
API menampilkan amount sebagai number rupiah, misalnya 500000.
```

---

### 8.2 `fee_types`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NULL,
code VARCHAR(50) NOT NULL,
name VARCHAR(150) NOT NULL,
category VARCHAR(100) NOT NULL,
is_recurring BOOLEAN NOT NULL DEFAULT false,
status VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Contoh:

```text
SPP
Daftar Ulang
Uang Pangkal
Seragam
Buku
Kegiatan
```

Unique constraint:

```text
foundation_id + school_id + code
```

---

### 8.3 `fee_schemes`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
academic_year_id UUID NOT NULL,
name VARCHAR(150) NOT NULL,
grade_level_id UUID NULL,
status VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

---

### 8.4 `fee_scheme_items`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
fee_scheme_id UUID NOT NULL,
fee_type_id UUID NOT NULL,
amount NUMERIC(14,2) NOT NULL,
billing_frequency VARCHAR(50) NOT NULL,
due_day INT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

`billing_frequency`:

```text
monthly
semester
yearly
one_time
```

---

### 8.5 `student_fee_policies`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
student_id UUID NOT NULL,
fee_type_id UUID NOT NULL,
policy_type VARCHAR(50) NOT NULL,
discount_type VARCHAR(50) NULL,
discount_value NUMERIC(14,2) NULL,
custom_amount NUMERIC(14,2) NULL,
start_period VARCHAR(20) NOT NULL,
end_period VARCHAR(20) NULL,
status VARCHAR(50) NOT NULL,
reason TEXT NOT NULL,
approval_request_id UUID NULL,
created_by UUID NOT NULL,
approved_by UUID NULL,
approved_at TIMESTAMP NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

`policy_type`:

```text
normal
free_spp
percentage_discount
fixed_amount_discount
sibling_discount
scholarship
custom_fee
```

Status:

```text
draft
submitted
approved
rejected
inactive
```

Aturan:

- free SPP = 100% discount untuk SPP pada periode tertentu;
- semua fee policy wajib punya reason, approval status, dan audit log;
- fokus MVP pada SPP bulanan.

---

### 8.6 `sibling_discount_rules`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NULL,
fee_type_id UUID NOT NULL,
child_order INT NOT NULL,
discount_type VARCHAR(50) NOT NULL,
discount_value NUMERIC(14,2) NOT NULL,
is_active BOOLEAN NOT NULL DEFAULT true,
effective_from DATE NOT NULL,
effective_until DATE NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Contoh:

```text
child_order = 1, discount_value = 0
child_order = 2, discount_value = 25
child_order = 3, discount_value = 50
```

---

### 8.7 `student_bills`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
student_id UUID NOT NULL,
academic_year_id UUID NOT NULL,
semester_id UUID NULL,
invoice_number VARCHAR(100) NOT NULL,
billing_period VARCHAR(20) NOT NULL,
total_amount NUMERIC(14,2) NOT NULL,
paid_amount NUMERIC(14,2) NOT NULL DEFAULT 0,
outstanding_amount NUMERIC(14,2) NOT NULL,
status VARCHAR(50) NOT NULL,
due_date DATE NULL,
student_snapshot_json JSONB NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Status:

```text
unpaid
partially_paid
paid
overdue
cancelled
void
```

Unique constraint:

```text
foundation_id + school_id + invoice_number
```

Catatan:

- `student_snapshot_json` menyimpan snapshot nama siswa, kelas, dan data ringkas invoice.
- Tagihan historis tidak boleh berubah jika data master berubah.

---

### 8.8 `student_bill_items`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
bill_id UUID NOT NULL,
fee_type_id UUID NOT NULL,
description VARCHAR(255) NOT NULL,
base_amount NUMERIC(14,2) NOT NULL,
discount_amount NUMERIC(14,2) NOT NULL DEFAULT 0,
final_amount NUMERIC(14,2) NOT NULL,
applied_policy_id UUID NULL,
applied_policy_snapshot_json JSONB NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Aturan:

- wajib menyimpan `base_amount`, `discount_amount`, `final_amount`, dan snapshot policy.
- perubahan fee policy di masa depan tidak mengubah tagihan lama.

---

### 8.9 `student_payments`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
student_id UUID NOT NULL,
bill_id UUID NOT NULL,
payment_number VARCHAR(100) NOT NULL,
payment_method VARCHAR(50) NOT NULL,
amount NUMERIC(14,2) NOT NULL,
status VARCHAR(50) NOT NULL,
paid_at TIMESTAMP NULL,
verified_by UUID NULL,
verified_at TIMESTAMP NULL,
rejected_by UUID NULL,
rejected_at TIMESTAMP NULL,
rejection_reason TEXT NULL,
external_reference VARCHAR(255) NULL,
gateway_transaction_id VARCHAR(255) NULL,
idempotency_key VARCHAR(255) NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

`payment_method`:

```text
cash
bank_transfer_manual
payment_gateway
qris
virtual_account
```

Status:

```text
draft
pending_verification
verified
rejected
void_requested
voided
refunded
```

Unique constraint:

```text
foundation_id + school_id + payment_number
idempotency_key nullable unique sesuai operation scope
```

---

### 8.10 `payment_proofs`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
payment_id UUID NOT NULL,
file_id UUID NOT NULL,
uploaded_by UUID NOT NULL,
uploaded_at TIMESTAMP NOT NULL,
created_at TIMESTAMP NOT NULL
```

---

### 8.11 `payment_receipts`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
payment_id UUID NOT NULL,
receipt_number VARCHAR(100) NOT NULL,
issued_at TIMESTAMP NOT NULL,
issued_by UUID NOT NULL,
pdf_file_id UUID NULL,
created_at TIMESTAMP NOT NULL
```

Unique constraint:

```text
foundation_id + school_id + receipt_number
```

---

### 8.12 `payment_reconciliations`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
period_key VARCHAR(20) NOT NULL,
status VARCHAR(50) NOT NULL,
total_payments NUMERIC(14,2) NOT NULL DEFAULT 0,
total_verified NUMERIC(14,2) NOT NULL DEFAULT 0,
total_voided NUMERIC(14,2) NOT NULL DEFAULT 0,
closed_by UUID NULL,
closed_at TIMESTAMP NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

---

### 8.13 Finance Index Utama

```text
fee_types.foundation_id + school_id + code
student_fee_policies.student_id
student_fee_policies.status
student_bills.student_id
student_bills.invoice_number
student_bills.billing_period
student_bills.status
student_payments.bill_id
student_payments.student_id
student_payments.payment_number
student_payments.status
```

---

## 9. Academic Service Data Model

Database: `academic_db`

Academic Service menjadi owner untuk:

- curriculum;
- subject;
- schedule;
- attendance;
- grade book;
- score/input nilai;
- report template;
- report card.

Academic memakai reference ID dari School Core seperti `student_id`, `teacher_id`, `class_id`, `academic_year_id`, dan `semester_id`.

---

### 9.1 `curriculums`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
name VARCHAR(150) NOT NULL,
code VARCHAR(50) NOT NULL,
is_default BOOLEAN NOT NULL DEFAULT false,
status VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Contoh:

```text
Kurikulum Merdeka
```

---

### 9.2 `learning_phases`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
curriculum_id UUID NOT NULL,
code VARCHAR(50) NOT NULL,
name VARCHAR(100) NOT NULL,
school_level VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Contoh fase:

```text
Pondasi
A
B
C
D
E
F
```

---

### 9.3 `subjects`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NULL,
curriculum_id UUID NOT NULL,
code VARCHAR(50) NOT NULL,
name VARCHAR(150) NOT NULL,
school_level VARCHAR(50) NOT NULL,
status VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

---

### 9.4 `subject_groups`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NULL,
name VARCHAR(150) NOT NULL,
school_level VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

---

### 9.5 `class_subjects`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
academic_year_id UUID NOT NULL,
semester_id UUID NOT NULL,
class_id UUID NOT NULL,
subject_id UUID NOT NULL,
teacher_id UUID NULL,
status VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

`class_id` dan `teacher_id` adalah reference ke School Core.

---

### 9.6 `schedules`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
academic_year_id UUID NOT NULL,
semester_id UUID NOT NULL,
class_id UUID NOT NULL,
subject_id UUID NOT NULL,
teacher_id UUID NOT NULL,
day_of_week INT NOT NULL,
start_time TIME NOT NULL,
end_time TIME NOT NULL,
room_id UUID NULL,
status VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

`day_of_week`:

```text
1 = Monday
2 = Tuesday
...
7 = Sunday
```

---

### 9.7 `student_attendances`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
academic_year_id UUID NOT NULL,
semester_id UUID NOT NULL,
class_id UUID NOT NULL,
student_id UUID NOT NULL,
subject_id UUID NULL,
teacher_id UUID NULL,
attendance_date DATE NOT NULL,
status VARCHAR(50) NOT NULL,
notes TEXT NULL,
created_by UUID NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Status:

```text
present
sick
excused
absent
late
```

Unique recommendation:

```text
foundation_id + school_id + class_id + student_id + attendance_date + subject_id
```

Catatan:

- Jika `subject_id` nullable, perlu strategi unique index khusus PostgreSQL untuk null handling.

---

### 9.8 `assessment_components`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NULL,
name VARCHAR(150) NOT NULL,
code VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Contoh:

```text
formatif
sumatif
proyek
praktik
catatan
```

---

### 9.9 `assessment_schemes`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
academic_year_id UUID NOT NULL,
semester_id UUID NOT NULL,
subject_id UUID NOT NULL,
class_id UUID NULL,
component_id UUID NOT NULL,
weight_percentage NUMERIC(5,2) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

---

### 9.10 `grade_books`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
academic_year_id UUID NOT NULL,
semester_id UUID NOT NULL,
class_id UUID NOT NULL,
subject_id UUID NOT NULL,
teacher_id UUID NOT NULL,
status VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Status:

```text
draft
submitted
approved
published
revision_requested
locked
```

---

### 9.11 `student_scores`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
grade_book_id UUID NOT NULL,
student_id UUID NOT NULL,
component_id UUID NOT NULL,
score NUMERIC(5,2) NULL,
description TEXT NULL,
status VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

---

### 9.12 `report_templates`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NULL,
school_level VARCHAR(50) NOT NULL,
name VARCHAR(150) NOT NULL,
template_json JSONB NOT NULL,
status VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

---

### 9.13 `report_cards`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
academic_year_id UUID NOT NULL,
semester_id UUID NOT NULL,
student_id UUID NOT NULL,
class_id UUID NOT NULL,
template_id UUID NOT NULL,
status VARCHAR(50) NOT NULL,
published_at TIMESTAMP NULL,
published_by UUID NULL,
locked_at TIMESTAMP NULL,
pdf_file_id UUID NULL,
student_snapshot_json JSONB NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Status:

```text
draft
reviewed
approved
published
locked
revision_requested
revised
```

Catatan:

- `student_snapshot_json` wajib agar rapor historis tidak berubah jika data siswa/kelas berubah.
- Revisi setelah publish wajib approval dan audit log.

---

### 9.14 `report_card_items`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
report_card_id UUID NOT NULL,
subject_id UUID NULL,
aspect_code VARCHAR(100) NULL,
title VARCHAR(150) NOT NULL,
score NUMERIC(5,2) NULL,
description TEXT NULL,
sort_order INT NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Catatan:

- Untuk TK/PAUD, `subject_id` dapat null dan memakai `aspect_code`.

---

### 9.15 Academic Index Utama

```text
subjects.foundation_id + school_id + code
schedules.foundation_id + school_id + class_id
schedules.teacher_id
student_attendances.student_id
student_attendances.class_id + attendance_date
grade_books.class_id + subject_id + teacher_id
student_scores.grade_book_id + student_id
report_cards.student_id
report_cards.class_id
report_cards.status
```

---

## 10. Communication Service Data Model

Database: `communication_db`

Communication Service menjadi owner untuk:

- announcement;
- announcement target;
- notification;
- notification template;
- delivery log;
- notification preference;
- surat dasar jika masuk MVP.

---

### 10.1 `announcements`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NULL,
title VARCHAR(200) NOT NULL,
body TEXT NOT NULL,
priority VARCHAR(50) NOT NULL,
status VARCHAR(50) NOT NULL,
target_scope VARCHAR(50) NOT NULL,
published_at TIMESTAMP NULL,
published_by UUID NULL,
expires_at TIMESTAMP NULL,
created_by UUID NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Priority:

```text
low
normal
high
urgent
```

Status:

```text
draft
submitted
published
rejected
archived
```

---

### 10.2 `announcement_targets`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NULL,
announcement_id UUID NOT NULL,
target_type VARCHAR(50) NOT NULL,
target_id UUID NULL,
created_at TIMESTAMP NOT NULL
```

`target_type`:

```text
foundation
school
class
role
user
student
parent
teacher
```

---

### 10.3 `notifications`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NULL,
recipient_user_id UUID NOT NULL,
event_type VARCHAR(150) NOT NULL,
category VARCHAR(100) NOT NULL,
priority VARCHAR(50) NOT NULL,
title VARCHAR(200) NOT NULL,
body TEXT NOT NULL,
data_json JSONB NULL,
read_at TIMESTAMP NULL,
created_at TIMESTAMP NOT NULL
```

---

### 10.4 `notification_templates`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NULL,
event_type VARCHAR(150) NOT NULL,
channel VARCHAR(50) NOT NULL,
language VARCHAR(20) NOT NULL,
title_template TEXT NOT NULL,
body_template TEXT NOT NULL,
is_active BOOLEAN NOT NULL DEFAULT true,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Channel:

```text
in_app
fcm
email
whatsapp
sms
```

MVP channel:

```text
in_app
fcm
email
```

---

### 10.5 `notification_deliveries`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
notification_id UUID NOT NULL,
channel VARCHAR(50) NOT NULL,
status VARCHAR(50) NOT NULL,
provider VARCHAR(100) NULL,
provider_message_id VARCHAR(255) NULL,
sent_at TIMESTAMP NULL,
failed_at TIMESTAMP NULL,
failure_reason TEXT NULL,
retry_count INT NOT NULL DEFAULT 0,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Status:

```text
pending
sent
failed
retrying
cancelled
```

---

### 10.6 `notification_preferences`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
user_id UUID NOT NULL,
channel VARCHAR(50) NOT NULL,
category VARCHAR(100) NOT NULL,
is_enabled BOOLEAN NOT NULL DEFAULT true,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Catatan:

- Notifikasi penting seperti tagihan, pembayaran ditolak, absensi alfa, security, dan emergency tidak boleh dimatikan sepenuhnya.

---

### 10.7 `letters`

Jika surat sederhana masuk MVP:

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NULL,
letter_number VARCHAR(100) NULL,
letter_type VARCHAR(100) NOT NULL,
subject VARCHAR(200) NOT NULL,
body TEXT NOT NULL,
status VARCHAR(50) NOT NULL,
requested_by UUID NOT NULL,
approved_by UUID NULL,
approved_at TIMESTAMP NULL,
pdf_file_id UUID NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Status:

```text
draft
submitted
revision_requested
approved
rejected
published
```

---

### 10.8 Communication Index Utama

```text
announcements.foundation_id + school_id
announcements.status
announcements.published_at
announcement_targets.announcement_id
notifications.recipient_user_id
notifications.read_at
notification_deliveries.status
notification_deliveries.channel
```

---

## 11. Reporting Service Data Model

Database: `reporting_db`

Reporting Service menyimpan projection/read model untuk dashboard dan summary.

Reporting Service **bukan source of truth** data operasional.

---

### 11.1 `foundation_dashboard_metrics`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
academic_year_id UUID NULL,
semester_id UUID NULL,
total_active_students INT NOT NULL DEFAULT 0,
total_teachers INT NOT NULL DEFAULT 0,
total_staff INT NOT NULL DEFAULT 0,
total_applicants INT NOT NULL DEFAULT 0,
total_billed NUMERIC(14,2) NOT NULL DEFAULT 0,
total_paid NUMERIC(14,2) NOT NULL DEFAULT 0,
total_outstanding NUMERIC(14,2) NOT NULL DEFAULT 0,
collection_rate NUMERIC(5,2) NOT NULL DEFAULT 0,
updated_at TIMESTAMP NOT NULL
```

---

### 11.2 `school_dashboard_metrics`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
academic_year_id UUID NULL,
semester_id UUID NULL,
total_active_students INT NOT NULL DEFAULT 0,
total_teachers INT NOT NULL DEFAULT 0,
total_classes INT NOT NULL DEFAULT 0,
total_applicants INT NOT NULL DEFAULT 0,
total_billed NUMERIC(14,2) NOT NULL DEFAULT 0,
total_paid NUMERIC(14,2) NOT NULL DEFAULT 0,
total_outstanding NUMERIC(14,2) NOT NULL DEFAULT 0,
collection_rate NUMERIC(5,2) NOT NULL DEFAULT 0,
updated_at TIMESTAMP NOT NULL
```

---

### 11.3 `student_summary_metrics`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
academic_year_id UUID NULL,
grade_level_id UUID NULL,
class_id UUID NULL,
active_count INT NOT NULL DEFAULT 0,
inactive_count INT NOT NULL DEFAULT 0,
transferred_count INT NOT NULL DEFAULT 0,
graduated_count INT NOT NULL DEFAULT 0,
dropped_out_count INT NOT NULL DEFAULT 0,
updated_at TIMESTAMP NOT NULL
```

---

### 11.4 `finance_summary_metrics`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NULL,
academic_year_id UUID NULL,
period_key VARCHAR(20) NOT NULL,
total_billed NUMERIC(14,2) NOT NULL DEFAULT 0,
total_paid NUMERIC(14,2) NOT NULL DEFAULT 0,
total_outstanding NUMERIC(14,2) NOT NULL DEFAULT 0,
collection_rate NUMERIC(5,2) NOT NULL DEFAULT 0,
updated_at TIMESTAMP NOT NULL
```

---

### 11.5 `attendance_summary_metrics`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
class_id UUID NULL,
attendance_date DATE NOT NULL,
total_students INT NOT NULL DEFAULT 0,
present_count INT NOT NULL DEFAULT 0,
sick_count INT NOT NULL DEFAULT 0,
excused_count INT NOT NULL DEFAULT 0,
absent_count INT NOT NULL DEFAULT 0,
late_count INT NOT NULL DEFAULT 0,
not_marked_count INT NOT NULL DEFAULT 0,
updated_at TIMESTAMP NOT NULL
```

---

### 11.6 `academic_progress_metrics`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NOT NULL,
academic_year_id UUID NOT NULL,
semester_id UUID NOT NULL,
class_id UUID NULL,
total_grade_books INT NOT NULL DEFAULT 0,
submitted_grade_books INT NOT NULL DEFAULT 0,
approved_grade_books INT NOT NULL DEFAULT 0,
published_report_cards INT NOT NULL DEFAULT 0,
total_report_cards INT NOT NULL DEFAULT 0,
updated_at TIMESTAMP NOT NULL
```

---

### 11.7 `approval_pending_metrics`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NULL,
module VARCHAR(100) NOT NULL,
action VARCHAR(150) NOT NULL,
pending_count INT NOT NULL DEFAULT 0,
updated_at TIMESTAMP NOT NULL
```

---

### 11.8 `notification_summary_metrics`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NULL,
period_key VARCHAR(20) NOT NULL,
total_notifications INT NOT NULL DEFAULT 0,
sent_count INT NOT NULL DEFAULT 0,
failed_count INT NOT NULL DEFAULT 0,
updated_at TIMESTAMP NOT NULL
```

---

### 11.9 `reporting_projection_offsets`

```sql
id UUID PRIMARY KEY,
source_service VARCHAR(100) NOT NULL,
event_type VARCHAR(150) NOT NULL,
last_event_id UUID NULL,
last_processed_at TIMESTAMP NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Fungsi:

- tracking event projection;
- membantu idempotency;
- membantu scheduled rebuild dan recovery.

---

## 12. Standard Tables per Service

Beberapa tabel standar dapat dibuat di service pemilik domain sesuai kebutuhan.

---

### 12.1 `audit_logs`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NULL,
actor_user_id UUID NULL,
actor_role VARCHAR(100) NULL,
action VARCHAR(150) NOT NULL,
module VARCHAR(100) NOT NULL,
entity_type VARCHAR(100) NOT NULL,
entity_id UUID NOT NULL,
old_values_json JSONB NULL,
new_values_json JSONB NULL,
metadata_json JSONB NULL,
ip_address VARCHAR(100) NULL,
user_agent TEXT NULL,
request_id VARCHAR(100) NOT NULL,
correlation_id VARCHAR(100) NOT NULL,
occurred_at TIMESTAMP NOT NULL,
created_at TIMESTAMP NOT NULL
```

Aturan:

- Application log berbeda dari audit log.
- Aksi sensitif wajib audit log.
- Data sensitif harus dimasking.

---

### 12.2 `approval_requests`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NULL,
module VARCHAR(100) NOT NULL,
entity_type VARCHAR(100) NOT NULL,
entity_id UUID NOT NULL,
action VARCHAR(150) NOT NULL,
status VARCHAR(50) NOT NULL,
requested_by UUID NOT NULL,
requested_at TIMESTAMP NOT NULL,
current_approver_role VARCHAR(100) NULL,
current_approver_id UUID NULL,
approved_by UUID NULL,
approved_at TIMESTAMP NULL,
rejected_by UUID NULL,
rejected_at TIMESTAMP NULL,
revision_requested_by UUID NULL,
revision_requested_at TIMESTAMP NULL,
reason TEXT NULL,
revision_note TEXT NULL,
before_json JSONB NULL,
after_json JSONB NULL,
metadata_json JSONB NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Status:

```text
draft
submitted
revision_requested
revised
approved
rejected
cancelled
```

---

### 12.3 `approval_steps`

```sql
id UUID PRIMARY KEY,
approval_request_id UUID NOT NULL,
step_order INT NOT NULL,
approver_role VARCHAR(100) NOT NULL,
approver_id UUID NULL,
status VARCHAR(50) NOT NULL,
acted_by UUID NULL,
acted_at TIMESTAMP NULL,
note TEXT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Catatan:

- MVP mayoritas approval 1 level.
- `approval_steps` disiapkan untuk multi-level terbatas.

---

### 12.4 `numbering_sequences`

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NULL,
system_key VARCHAR(100) NOT NULL,
period_key VARCHAR(50) NOT NULL,
prefix VARCHAR(50) NOT NULL,
current_number BIGINT NOT NULL,
padding_length INT NOT NULL,
reset_policy VARCHAR(50) NOT NULL,
scope_level VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Unique constraint:

```text
foundation_id + school_id + system_key + period_key
```

Aturan:

- Nomor yang sudah dipakai tidak boleh digunakan ulang.
- Gunakan transaction/row lock.
- Dikelola per service pada MVP.

---

### 12.5 `files`

Jika metadata file dikelola lokal di service:

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NULL,
owner_service VARCHAR(100) NOT NULL,
module VARCHAR(100) NOT NULL,
entity_type VARCHAR(100) NOT NULL,
entity_id UUID NOT NULL,
file_category VARCHAR(100) NOT NULL,
original_filename VARCHAR(255) NOT NULL,
stored_filename VARCHAR(255) NOT NULL,
storage_disk VARCHAR(100) NOT NULL,
storage_bucket VARCHAR(255) NOT NULL,
storage_path TEXT NOT NULL,
mime_type VARCHAR(150) NOT NULL,
file_extension VARCHAR(50) NOT NULL,
size_bytes BIGINT NOT NULL,
checksum_sha256 VARCHAR(128) NOT NULL,
classification VARCHAR(50) NOT NULL,
visibility VARCHAR(50) NOT NULL,
uploaded_by UUID NOT NULL,
uploaded_at TIMESTAMP NOT NULL,
status VARCHAR(50) NOT NULL,
deleted_at TIMESTAMP NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Classification:

```text
public
internal
restricted
confidential
```

Visibility:

```text
public
private
```

Status:

```text
pending_scan
active
archived
deleted
quarantined
```

---

### 12.6 `import_batches`

Untuk import Excel MVP di School Core:

```sql
id UUID PRIMARY KEY,
foundation_id UUID NOT NULL,
school_id UUID NULL,
import_type VARCHAR(100) NOT NULL,
file_id UUID NOT NULL,
status VARCHAR(50) NOT NULL,
mode VARCHAR(50) NOT NULL,
total_rows INT NOT NULL DEFAULT 0,
success_rows INT NOT NULL DEFAULT 0,
failed_rows INT NOT NULL DEFAULT 0,
warning_rows INT NOT NULL DEFAULT 0,
uploaded_by UUID NOT NULL,
started_at TIMESTAMP NULL,
finished_at TIMESTAMP NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Mode:

```text
create_only
upsert_limited
```

Status:

```text
uploaded
validated
confirmed
processing
completed
failed
cancelled
```

---

### 12.7 `import_batch_rows`

```sql
id UUID PRIMARY KEY,
import_batch_id UUID NOT NULL,
row_number INT NOT NULL,
status VARCHAR(50) NOT NULL,
entity_type VARCHAR(100) NOT NULL,
entity_id UUID NULL,
error_message TEXT NULL,
warning_message TEXT NULL,
raw_data_json JSONB NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

---

### 12.8 `outbox_events`

Untuk service penting yang publish event:

```sql
id UUID PRIMARY KEY,
event_id UUID NOT NULL,
event_type VARCHAR(150) NOT NULL,
event_version INT NOT NULL,
aggregate_type VARCHAR(100) NOT NULL,
aggregate_id UUID NOT NULL,
payload_json JSONB NOT NULL,
status VARCHAR(50) NOT NULL,
retry_count INT NOT NULL DEFAULT 0,
next_retry_at TIMESTAMP NULL,
published_at TIMESTAMP NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
```

Status:

```text
pending
published
failed
```

---

### 12.9 `processed_events`

Untuk consumer idempotency:

```sql
id UUID PRIMARY KEY,
event_id UUID NOT NULL,
event_type VARCHAR(150) NOT NULL,
source_service VARCHAR(100) NOT NULL,
processed_at TIMESTAMP NOT NULL,
created_at TIMESTAMP NOT NULL
```

Unique constraint:

```text
event_id
```

---

## 13. Relasi Lintas Service Utama

Karena database per service, relasi lintas service berupa reference ID.

---

### 13.1 Identity ↔ School Core

```text
identity_db.users.id ↔ school_core_db.guardians.user_id
identity_db.users.id ↔ school_core_db.teachers.user_id
identity_db.users.id ↔ school_core_db.students.user_id jika siswa punya akun nanti
```

Owner:

```text
users = Identity Service
guardians/teachers/students = School Core Service
```

---

### 13.2 Admission → School Core

```text
admission_db.applicants.converted_student_id → school_core_db.students.id
```

Owner:

```text
applicant = Admission Service
student = School Core Service
```

---

### 13.3 Finance → School Core

```text
finance_db.student_bills.student_id → school_core_db.students.id
finance_db.student_payments.student_id → school_core_db.students.id
```

Finance wajib menyimpan snapshot untuk invoice dan receipt.

---

### 13.4 Academic → School Core

```text
academic_db.student_attendances.student_id → school_core_db.students.id
academic_db.student_attendances.class_id → school_core_db.classes.id
academic_db.student_attendances.teacher_id → school_core_db.teachers.id
academic_db.report_cards.student_id → school_core_db.students.id
```

Academic wajib menyimpan snapshot pada `report_cards`.

---

### 13.5 Communication → Identity / School Core

```text
communication_db.notifications.recipient_user_id → identity_db.users.id
communication_db.announcement_targets.target_id → user_id/class_id/student_id/role_id sesuai target_type
```

---

### 13.6 Reporting ← Semua Service

Reporting menerima event dari semua service.

Reporting tidak menjadi owner data operasional.

---

## 14. Index dan Constraint Standard

Index wajib hampir di semua tabel domain:

```text
foundation_id
school_id
status
created_at
academic_year_id
semester_id
student_id
class_id
```

Search/filter index:

```text
students.full_name
students.student_number
students.nisn
teachers.full_name
applicants.registration_number
applicants.full_name
student_bills.invoice_number
student_payments.payment_number
payment_receipts.receipt_number
```

Untuk PostgreSQL:

```text
B-tree index untuk filter biasa
GIN/pg_trgm untuk pencarian nama jika diperlukan
```

Constraint penting:

```text
students: foundation_id + school_id + student_number
classes: foundation_id + school_id + academic_year_id + class_code
student_bills: foundation_id + school_id + invoice_number
student_payments: foundation_id + school_id + payment_number
payment_receipts: foundation_id + school_id + receipt_number
applicants: foundation_id + school_id + registration_number
numbering_sequences: foundation_id + school_id + system_key + period_key
```

---

## 15. Enum/Status Standard

### 15.1 Student Status

```text
active
inactive
transferred
graduated
dropped_out
```

### 15.2 Approval Status

```text
draft
submitted
revision_requested
revised
approved
rejected
cancelled
```

### 15.3 Bill Status

```text
unpaid
partially_paid
paid
overdue
cancelled
void
```

### 15.4 Payment Status

```text
draft
pending_verification
verified
rejected
void_requested
voided
refunded
```

### 15.5 Report Card Status

```text
draft
reviewed
approved
published
locked
revision_requested
revised
```

### 15.6 Admission Status

```text
draft
submitted
verified
accepted
rejected
converted
cancelled
```

### 15.7 File Classification

```text
public
internal
restricted
confidential
```

### 15.8 File Status

```text
pending_scan
active
archived
deleted
quarantined
```

### 15.9 Notification Delivery Status

```text
pending
sent
failed
retrying
cancelled
```

---

## 16. Snapshot Rules

Data historis penting wajib menyimpan snapshot.

Wajib snapshot:

| Domain | Snapshot |
|---|---|
| Bill/Invoice | student name, class, fee policy, billing period |
| Bill Item | base amount, discount amount, final amount, applied policy |
| Receipt | payment, payer/student identity summary, amount |
| Report Card | student, class, academic year, semester, template |
| Admission Decision | applicant data saat keputusan dibuat |
| Official Letter | subject, recipient, signer, document number |

Alasan:

```text
Data historis tidak boleh berubah jika data master berubah di masa depan.
```

Contoh:

```text
Jika siswa pindah kelas setelah invoice dibuat, invoice lama tetap menampilkan kelas saat invoice diterbitkan.
Jika fee policy berubah, tagihan lama tetap memakai applied policy snapshot saat generate.
```

---

## 17. Data Privacy Rules pada Data Model

Klasifikasi data:

| Classification | Contoh |
|---|---|
| `public` | logo, banner, brosur PPDB publik |
| `internal` | kalender, data operasional umum sekolah |
| `restricted` | data siswa, orang tua, nilai, tagihan, pembayaran, dokumen siswa |
| `confidential` | BK, UKS/kesehatan, payroll, credential/token, dokumen hukum, backup |

Aturan:

- token/password tidak boleh disimpan plain;
- refresh token hanya hash;
- file Restricted/Confidential private by default;
- Confidential data tidak boleh masuk application log;
- download/export Restricted/Confidential wajib audit;
- field-level access control diterapkan di service/usecase.

---

## 18. Data Model Rules untuk AI Agent

AI Agent wajib mengikuti aturan berikut:

```text
1. Jangan membuat satu database monolitik.
2. Jangan membuat foreign key lintas database service.
3. Jangan membuat query langsung ke database service lain.
4. Jangan menghapus foundation_id dari tabel domain.
5. Jangan menghapus school_id dari tabel yang terkait unit sekolah.
6. Jangan menjadikan free_spp/diskon sebagai status siswa.
7. Jangan menyimpan token/password secara plain.
8. Jangan memakai float untuk finance amount.
9. Jangan overwrite dokumen resmi; buat versi/snapshot.
10. Jangan membuat status enum baru tanpa memperbarui dokumen.
11. Jangan membuat migration tanpa index/scope yang sesuai.
12. Jangan membuat data historis tanpa snapshot jika domain membutuhkan histori.
```

---

## 19. Ringkasan Keputusan Final

Keputusan final data model MVP:

```text
Data model MVP dibuat per service sesuai data ownership, bukan satu ERD monolitik.
```

Database MVP:

```text
identity_db
school_core_db
admission_db
academic_db
finance_db
communication_db
reporting_db
```

Prinsip utama:

```text
- UUID sebagai primary key teknis.
- foundation_id sebagai tenant boundary utama.
- school_id untuk data terkait unit sekolah.
- Tidak ada foreign key lintas database service.
- Relasi lintas service menggunakan reference ID.
- Validasi lintas service melalui gRPC atau domain event.
- Reporting hanya projection/read model, bukan source of truth.
- Tabel audit_logs, approval_requests, numbering_sequences, files, import_batches, outbox_events, dan processed_events dibuat di service pemilik domain sesuai kebutuhan.
- Data historis penting seperti invoice, receipt, report card, dan bill item wajib menyimpan snapshot.
```

Ownership ringkas:

```text
Identity Service:
- user, credential, session, role, permission, role assignment

School Core Service:
- foundation, school, academic year, semester, student, guardian, teacher, class, assignment

Admission Service:
- PPDB dan applicant sebelum dikonversi menjadi siswa

Academic Service:
- curriculum, subject, schedule, attendance, grade, report card, report template

Finance Service:
- fee type, fee scheme, fee policy, sibling discount, bill, payment, receipt, reconciliation

Communication Service:
- announcement, notification, template, delivery log, preference, surat dasar jika masuk MVP

Reporting Service:
- projection/read model dashboard dan summary
```

---

## 20. Dokumen Terkait

Dokumen ini harus dibaca bersama:

```text
01-technical-architecture.md
02-service-boundary.md
04-api-contract.md
05-event-contract.md
07-test-plan-acceptance-criteria.md
08-coding-standard.md
09-ai-agent-rules.md
```

