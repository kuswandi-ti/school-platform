# 02 — Service Boundary & Data Ownership

**Project:** `school-platform`  
**Document:** Service Boundary & Data Ownership  
**Status:** Final Decision Draft  
**Scope:** MVP Sistem Manajemen Sekolah/Yayasan  

---

## 1. Tujuan Dokumen

Dokumen ini mendefinisikan batas tanggung jawab setiap service pada arsitektur microservice `school-platform`.

Tujuan utama dokumen ini:

1. Menentukan **owner data** untuk setiap domain utama.
2. Mencegah duplikasi source of truth antar service.
3. Menetapkan pola komunikasi antar service menggunakan **gRPC** dan **RabbitMQ domain event**.
4. Menegaskan bahwa service **tidak boleh query langsung ke database service lain**.
5. Menjadi acuan untuk Backend Developer, Frontend Developer, QA, Infrastructure/DevOps, dan AI Agent.

---

## 2. Prinsip Utama Service Boundary

Service boundary MVP dikunci berdasarkan **data ownership**, bukan berdasarkan menu UI.

Prinsip utama:

```text
1 data utama = 1 owner service
service lain hanya menyimpan reference ID atau snapshot terbatas
tidak ada query langsung ke database service lain
komunikasi sinkron antar service menggunakan gRPC
komunikasi async menggunakan RabbitMQ domain event
reporting menggunakan projection/read model, bukan query lintas database operasional
```

Contoh:

```text
School Core Service adalah owner data siswa.
Finance Service boleh menyimpan student_id pada tagihan.
Finance Service tidak boleh menjadi source of truth biodata siswa.
```

Jika Finance membutuhkan nama siswa pada invoice/receipt, Finance boleh menyimpan snapshot terbatas:

```text
student_id
student_name_snapshot
class_name_snapshot
school_id
```

Snapshot tersebut dipakai untuk histori dokumen/transaksi, bukan untuk menggantikan master data siswa.

---

## 3. Service MVP

Service MVP yang dikunci:

```text
1. api-gateway
2. identity-service
3. school-core-service
4. admission-service
5. academic-service
6. finance-service
7. communication-service
8. reporting-service
```

Shared packages yang mendukung standardisasi:

```text
packages/proto
packages/openapi
packages/events
packages/shared-go
shared audit library
shared numbering library
shared error/response contract
shared file/storage helper
```

---

## 4. Ringkasan Data Ownership

| Domain/Data | Owner Service | Catatan |
|---|---|---|
| User account, credential, session | Identity Service | Source of truth auth |
| Role, permission, role assignment | Identity Service | Service domain tetap wajib cek permission/scope |
| Foundation, school/unit | School Core Service | Master data yayasan/sekolah |
| Academic year, semester | School Core Service | Global level yayasan untuk MVP |
| Student master | School Core Service | Finance/Academic hanya menyimpan student_id/snapshot |
| Guardian/parent master | School Core Service | Identity hanya menyimpan user account |
| Teacher master dasar | School Core Service | HR detail nanti phase lanjut |
| Class/rombel | School Core Service | Academic memakai class_id reference |
| Applicant PPDB | Admission Service | Sebelum diterima menjadi siswa |
| Student setelah PPDB conversion | School Core Service | Admission menyimpan converted_student_id |
| Curriculum, subject, schedule | Academic Service | Proses akademik |
| Attendance, grade, report card | Academic Service | Source of truth akademik |
| Fee type, fee scheme, fee policy | Finance Service | Termasuk free SPP/diskon/beasiswa/sibling discount |
| Bill, payment, receipt, reconciliation | Finance Service | Snapshot wajib untuk histori |
| Announcement, notification, template | Communication Service | Event-driven notification |
| Dashboard summary/projection | Reporting Service | Bukan source of truth operasional |
| File metadata | Service pemilik domain | File Service terpusat belum masuk MVP |
| Approval request | Service pemilik domain | Approval Service terpusat belum masuk MVP |
| Audit log | Service pemilik domain | Audit terpusat phase lanjut via event |
| Numbering sequence | Service pemilik domain | Numbering Service terpusat belum masuk MVP |

---

## 5. API Gateway Boundary

### 5.1 Tanggung Jawab

`api-gateway` bukan owner data domain. API Gateway bertanggung jawab untuk:

```text
- Menyediakan external REST/JSON API untuk Next.js dan Flutter
- Validasi access token
- Extract identity, foundation_id, school_id, role, permission, dan scope
- Basic guard
- Request routing
- REST/JSON to gRPC mapping
- Response standardization
- Error mapping dari gRPC ke HTTP
- Rate limiting dasar
- Request ID dan correlation ID
- CORS dan security header untuk API publik
```

### 5.2 Yang Tidak Boleh Dilakukan API Gateway

API Gateway tidak boleh:

```text
- Menyimpan business data domain
- Menghitung tagihan/SPP
- Memproses nilai/rapor
- Menentukan keputusan PPDB
- Menjadi tempat business logic utama
- Query langsung ke database service domain
```

Business logic harus berada di service pemilik domain.

### 5.3 Storage API Gateway

API Gateway boleh memiliki storage/cache kecil untuk:

```text
- route config
- gateway config
- rate limit cache via Redis
```

Namun tidak boleh menjadi source of truth data bisnis.

---

## 6. Identity Service Boundary

### 6.1 Owner Data

`identity-service` menjadi owner untuk:

```text
- user account
- credential/password hash
- user sessions
- refresh token hash
- devices/session records
- roles
- permissions
- role permissions
- user role assignments
- login/security events
```

Tabel utama:

```text
users
user_sessions
user_devices
roles
permissions
role_permissions
user_role_assignments
identity_audit_logs
```

### 6.2 Bukan Owner

Identity Service bukan owner untuk:

```text
- detail siswa
- detail guru/pegawai
- detail orang tua/wali sebagai entitas sekolah
- data akademik
- data finance
- data PPDB
```

Catatan penting:

```text
User account ≠ Student record
User account ≠ Teacher record
User account ≠ Guardian record
```

Relasi dilakukan melalui reference ID:

```text
users.id
student_id
teacher_id / employee_id
guardian_id
```

Contoh:

```text
Orang tua memiliki user account di Identity Service.
Data detail guardian/parent berada di School Core Service.
Identity hanya menyimpan user_id, role, permission, dan scope assignment.
```

---

## 7. School Core Service Boundary

### 7.1 Owner Data

`school-core-service` menjadi owner data inti yayasan/sekolah:

```text
- foundation
- school/unit TK/SD/SMP/SMA
- school profile
- academic year
- semester
- calendar academic global/sekolah jika MVP membutuhkan
- student master
- guardian/parent master
- teacher master dasar
- grade level
- class/rombel
- room dasar jika hanya untuk kebutuhan kelas
- student-class assignment
- teacher assignment dasar
- homeroom assignment
```

Tabel utama:

```text
foundations
schools
academic_years
semesters
students
guardians
student_guardians
teachers
grade_levels
classes
rooms
student_class_assignments
teacher_assignments
homeroom_assignments
school_core_audit_logs
```

### 7.2 Guru dan HR

Untuk MVP, data guru master dasar berada di School Core Service.

Jika HR lengkap dibuat pada phase lanjut:

```text
School Core Service:
- teacher master untuk kebutuhan sekolah/akademik

HR/Payroll Service nanti:
- kontrak kerja
- cuti
- appraisal
- payroll profile
- dokumen HR lengkap
```

HR Service tidak masuk MVP.

### 7.3 Bukan Owner

School Core Service bukan owner untuk:

```text
- user login dan credential
- role/permission/session
- tagihan/pembayaran
- nilai/rapor
- PPDB applicant sebelum diterima
- notifikasi
- dashboard summary
```

---

## 8. Admission Service Boundary

### 8.1 Owner Data

`admission-service` menjadi owner untuk proses PPDB:

```text
- admission period / gelombang PPDB
- applicant / calon siswa
- applicant guardian data saat pendaftaran
- applicant document metadata reference
- document verification status
- acceptance/rejection decision
- admission workflow
- conversion request to student
```

Tabel utama:

```text
admission_periods
applicants
applicant_guardians
applicant_documents
applicant_verifications
admission_decisions
admission_audit_logs
```

### 8.2 Relasi dengan School Core

Sebelum diterima:

```text
Applicant adalah milik Admission Service.
```

Setelah diterima dan dikonversi:

```text
Admission Service memanggil School Core Service via gRPC.
School Core Service membuat student + guardian.
School Core Service menjadi owner student record.
Admission Service menyimpan converted_student_id sebagai reference.
```

Flow:

```text
Applicant accepted
↓
Admission Service calls School Core gRPC: ConvertApplicantToStudent / CreateStudent
↓
School Core creates student + guardian
↓
School Core publishes school.student.created
↓
Admission stores converted_student_id
↓
Admission publishes admission.applicant.converted_to_student
```

Admission Service tidak boleh langsung insert ke `school_core_db`.

### 8.3 Relasi dengan Finance

Jika PPDB membutuhkan tagihan pendaftaran/daftar ulang:

```text
Admission Service meminta Finance Service membuat registration bill.
Finance Service menjadi owner tagihan.
Admission Service hanya menyimpan bill_id/reference jika diperlukan.
```

---

## 9. Academic Service Boundary

### 9.1 Owner Data

`academic-service` menjadi owner proses akademik:

```text
- curriculum
- learning phase
- subject
- subject group
- class subject
- schedule
- attendance
- assessment component
- assessment scheme
- grade book
- student score
- report template
- report card
- report card publication
- report card revision workflow
```

Tabel utama:

```text
curriculums
learning_phases
subjects
subject_groups
class_subjects
schedules
student_attendances
assessment_components
assessment_schemes
grade_books
student_scores
report_templates
report_cards
report_card_items
academic_approval_requests
academic_audit_logs
```

### 9.2 Data Reference dari School Core

Academic Service menggunakan reference ID:

```text
foundation_id
school_id
academic_year_id
semester_id
student_id
teacher_id
class_id
room_id
```

Academic Service tidak menjadi owner biodata siswa/guru/kelas.

Untuk dokumen historis seperti rapor, Academic boleh menyimpan snapshot:

```text
student_snapshot_json
class_name_snapshot
teacher_name_snapshot
```

Snapshot dipakai agar rapor historis tidak berubah jika data master berubah.

### 9.3 Bukan Owner

Academic Service bukan owner untuk:

```text
- master data siswa/guru/kelas
- user login
- tagihan/pembayaran
- notifikasi delivery
- dashboard summary
```

---

## 10. Finance Service Boundary

### 10.1 Owner Data

`finance-service` menjadi owner untuk:

```text
- fee type
- fee scheme
- fee scheme item
- student fee policy
- free SPP
- discount
- scholarship
- sibling discount rule
- custom fee
- bill/invoice
- bill item
- payment
- payment proof reference
- receipt
- void/refund request
- reconciliation
- finance approval
```

Tabel utama:

```text
fee_types
fee_schemes
fee_scheme_items
student_fee_policies
sibling_discount_rules
student_bills
student_bill_items
student_payments
payment_proofs
payment_receipts
payment_reconciliations
finance_approval_requests
finance_audit_logs
finance_numbering_sequences
```

### 10.2 Data Reference

Finance Service menggunakan reference ID:

```text
student_id
guardian_id
school_id
academic_year_id
semester_id
class_id optional
```

Finance Service tidak menjadi owner biodata siswa.

Namun Finance wajib menyimpan snapshot pada dokumen/transaksi historis:

```text
student_name_snapshot
class_name_snapshot
guardian_name_snapshot
fee_policy_snapshot_json
```

Alasan:

```text
Invoice/receipt/tagihan lama tidak boleh berubah jika nama siswa, kelas, atau fee policy berubah di masa depan.
```

### 10.3 Fee Policy Bukan Status Siswa

Free SPP, diskon, beasiswa, dan sibling discount dikelola sebagai `student_fee_policies` di Finance Service.

Status siswa tetap berada di School Core:

```text
active
inactive
transferred
graduated
dropped_out
```

Fee policy MVP:

```text
normal
free_spp
percentage_discount
fixed_amount_discount
sibling_discount
scholarship
custom_fee
```

### 10.4 Relasi dengan File

Bukti pembayaran dan receipt PDF mengikuti file management policy:

```text
payment_proofs.file_id
payment_receipts.pdf_file_id
```

Metadata file disimpan di Finance Service atau tabel file lokal service pemilik domain.
File fisik disimpan di MinIO/R2 object storage.

---

## 11. Communication Service Boundary

### 11.1 Owner Data

`communication-service` menjadi owner untuk:

```text
- announcement
- announcement target
- notification
- notification template
- notification delivery log
- notification preference
- email/push delivery status
- surat menyurat dasar jika masuk MVP
```

Tabel utama:

```text
announcements
announcement_targets
notifications
notification_templates
notification_deliveries
notification_preferences
letters
communication_audit_logs
communication_numbering_sequences
```

### 11.2 Device Token

Untuk MVP, `user_devices` direkomendasikan berada di Identity Service karena terkait session/device ownership.

Communication Service dapat memperoleh target device dengan salah satu cara:

```text
1. gRPC call ke Identity Service untuk mendapatkan device token aktif.
2. Consume event identity.user_device.registered sebagai local read model jika nanti diperlukan.
```

Keputusan MVP:

```text
user_devices di Identity Service.
Communication Service request device tokens via gRPC atau read model event jika dibutuhkan.
```

### 11.3 Bukan Owner

Communication Service bukan owner untuk:

```text
- data siswa
- data guru
- tagihan
- pembayaran
- nilai/rapor
```

Communication Service menerima domain event seperti:

```text
finance.bill.generated
finance.payment.verified
academic.report_card.published
admission.applicant.accepted
communication.announcement.published
approval.request.created
```

Kemudian Communication Service membuat notifikasi berdasarkan template, priority, recipient, dan user preference.

---

## 12. Reporting Service Boundary

### 12.1 Owner Data

`reporting-service` menjadi owner untuk:

```text
- dashboard projection
- reporting read model
- aggregate metrics
- reporting cache/materialized summary
```

Tabel utama:

```text
foundation_dashboard_metrics
school_dashboard_metrics
student_summary_metrics
teacher_summary_metrics
admission_summary_metrics
finance_summary_metrics
attendance_summary_metrics
academic_progress_metrics
approval_pending_metrics
notification_summary_metrics
reporting_projection_offsets
processed_events
```

### 12.2 Sumber Data Reporting

Reporting Service mendapatkan data dari:

```text
- RabbitMQ domain events
- scheduled rebuild/sync resmi
- gRPC read endpoint jika benar-benar diperlukan dan terkontrol
```

Reporting Service tidak boleh query langsung:

```text
school_core_db
finance_db
academic_db
admission_db
identity_db
communication_db
```

### 12.3 Bukan Source of Truth

Reporting Service bukan source of truth data operasional.

Jika terjadi selisih:

```text
Service domain owner = source of truth
reporting_db = projection/read model yang bisa direbuild
```

---

## 13. File Management Boundary

MVP belum membuat File Service terpusat.

Keputusan MVP:

```text
File metadata disimpan di service pemilik domain.
Object storage shared menggunakan S3-compatible storage: MinIO/R2.
Shared file library/package digunakan untuk upload, signed URL, checksum, validation, dan classification.
```

Contoh ownership file metadata:

| File | Metadata Owner |
|---|---|
| Dokumen PPDB | Admission Service |
| Dokumen siswa/guru | School Core Service |
| Bukti pembayaran | Finance Service |
| Receipt PDF | Finance Service |
| Rapor PDF | Academic Service |
| Surat PDF | Communication Service |

File Service terpusat dapat dipertimbangkan pada phase lanjut jika kebutuhan file management menjadi kompleks.

---

## 14. Approval Boundary

MVP belum membuat Approval Service terpusat.

Keputusan MVP:

```text
approval_requests berada di masing-masing service pemilik domain.
```

Contoh:

```text
Finance Service:
- approval void/refund
- approval fee policy/diskon/beasiswa

Academic Service:
- approval publish/revisi rapor
- approval revisi nilai setelah publish

Admission Service:
- approval penerimaan/penolakan PPDB jika dibutuhkan

School Core Service:
- approval mutasi/lulus/perubahan data sensitif

Identity Service:
- approval role sensitif jika dibutuhkan
```

Struktur tabel approval harus distandarkan melalui shared library/package.

---

## 15. Audit Boundary

MVP menggunakan audit log lokal per service.

Keputusan MVP:

```text
audit_logs berada di masing-masing service pemilik domain.
```

Setiap service wajib mencatat aksi sensitif di domainnya sendiri.

Jangka panjang:

```text
AuditLogCreated event dipublish ke RabbitMQ.
Audit/Reporting Service dapat mengonsumsi event tersebut untuk pencarian/konsolidasi audit lintas service.
```

Audit lokal tetap menjadi source of truth domain.

---

## 16. Numbering Boundary

MVP belum membuat Numbering Service terpusat.

Keputusan MVP:

```text
numbering_sequences dikelola di masing-masing service pemilik domain.
Shared numbering library digunakan untuk validasi format, reset policy, concurrency lock, dan anti-duplikasi.
```

Contoh ownership numbering:

| Nomor | Owner Service |
|---|---|
| Nomor PPDB | Admission Service |
| Invoice/payment/receipt | Finance Service |
| NIS/nomor pegawai internal | School Core Service |
| Nomor rapor jika dipakai | Academic Service |
| Nomor surat | Communication Service |

Nomor yang sudah dipakai tidak boleh digunakan ulang meskipun dokumen dibatalkan/void.

---

## 17. Komunikasi Antar Service

### 17.1 gRPC untuk Kebutuhan Sinkron

gRPC digunakan jika service membutuhkan response langsung.

Contoh:

```text
API Gateway → Identity Service: login, validate token, get user context
API Gateway → School Core Service: CRUD siswa/guru/kelas
API Gateway → Finance Service: tagihan/pembayaran
Admission Service → School Core Service: convert applicant to student
Finance Service → School Core Service: validate student exists / get student snapshot
Academic Service → School Core Service: validate class/student/teacher assignment
Communication Service → Identity Service: get user/device target
```

### 17.2 RabbitMQ Event untuk Async

RabbitMQ domain event digunakan untuk:

```text
- notification
- reporting projection
- audit consolidation
- async workflow
- rebuild/sync trigger
```

Contoh event:

```text
school.student.created
admission.applicant.accepted
admission.applicant.converted_to_student
finance.bill.generated
finance.payment.verified
academic.attendance.marked
academic.report_card.published
communication.announcement.published
approval.request.created
audit.log.created
```

---

## 18. Relasi Lintas Service Utama

### 18.1 Identity ↔ School Core

```text
users.id ↔ guardians.user_id
users.id ↔ teachers.user_id
users.id ↔ students.user_id jika siswa punya akun
```

Owner:

```text
users = Identity Service
guardians/teachers/students = School Core Service
```

### 18.2 Admission → School Core

```text
applicants.converted_student_id → students.id
```

Owner:

```text
applicants = Admission Service
students = School Core Service
```

### 18.3 Finance → School Core

```text
student_bills.student_id → students.id
student_payments.student_id → students.id
```

Finance menyimpan snapshot untuk invoice/receipt.

### 18.4 Academic → School Core

```text
student_attendances.student_id → students.id
student_attendances.class_id → classes.id
student_attendances.teacher_id → teachers.id
report_cards.student_id → students.id
```

Academic menyimpan snapshot pada report card.

### 18.5 Communication → Identity/School Core

```text
notifications.recipient_user_id → users.id
announcement_targets.target_id → class_id/user_id/student_id tergantung target_type
```

### 18.6 Reporting ← Semua Service

Reporting menerima event dari semua service dan membangun projection.

Reporting tidak menjadi owner data operasional.

---

## 19. Anti-Pattern yang Dilarang

Aturan ini wajib diikuti developer dan AI Agent.

```text
Finance Service tidak boleh query school_core_db.
Academic Service tidak boleh query identity_db.
Reporting Service tidak boleh query semua database operasional langsung.
API Gateway tidak boleh punya business logic tagihan/nilai/PPDB.
Communication Service tidak boleh menyimpan ulang seluruh data siswa sebagai source of truth.
Service tidak boleh membuat foundation_id/school_id sendiri tanpa validasi.
Frontend tidak boleh menjadi satu-satunya filter authorization.
Service tidak boleh expose REST publik langsung selain API Gateway.
Service tidak boleh menyimpan token/password/data confidential mentah di log/event.
```

---

## 20. Boundary Decision Matrix

| Kebutuhan | Keputusan MVP |
|---|---|
| External API | API Gateway REST/JSON |
| Internal sync communication | gRPC/protobuf |
| Async communication | RabbitMQ domain events |
| Dashboard | Reporting Service + reporting_db projection |
| File metadata | Service pemilik domain |
| Object storage | MinIO/R2 S3-compatible |
| Approval | Lokal per service |
| Audit log | Lokal per service |
| Numbering | Lokal per service + shared library |
| Global search | Tidak masuk MVP |
| File Service terpusat | Tidak masuk MVP |
| Audit Service terpusat | Phase lanjut |
| Approval Service terpusat | Phase lanjut jika dibutuhkan |
| Numbering Service terpusat | Phase lanjut jika dibutuhkan |

---

## 21. Implikasi untuk AI Agent

Setiap task AI Agent wajib mematuhi service boundary ini.

AI Agent wajib:

```text
- Menentukan service owner sebelum membuat tabel/API/logic.
- Tidak membuat query lintas database service.
- Tidak menaruh business logic di API Gateway.
- Menggunakan gRPC jika butuh validasi sinkron antar service.
- Menggunakan event jika kebutuhan async/projection/notification.
- Menyimpan snapshot hanya untuk histori, bukan source of truth baru.
- Menambahkan audit log untuk aksi sensitif.
- Menambahkan event jika flow membutuhkan reporting/notification.
- Menjaga foundation_id dan school_id pada semua data domain.
```

Jika AI Agent ragu sebuah data milik service mana, task harus dihentikan dan boundary harus dikonfirmasi terlebih dahulu.

---

## 22. Kesimpulan Final

Keputusan final service boundary MVP:

```text
Service boundary MVP dikunci berdasarkan data ownership.

API Gateway bukan owner data domain; hanya menangani external REST/JSON API, validasi token, scope extraction, routing, REST-to-gRPC mapping, response standardization, rate limiting dasar, dan correlation ID.

Identity Service menjadi owner user account, credential, session, refresh token, role, permission, dan role assignment.

School Core Service menjadi owner foundation, school, academic year, semester, student master, guardian/parent master, teacher master dasar, class/rombel, student-class assignment, teacher assignment, dan homeroom assignment.

Admission Service menjadi owner proses PPDB dan applicant sebelum diterima. Setelah applicant dikonversi menjadi siswa, School Core menjadi owner student record.

Academic Service menjadi owner curriculum, subject, schedule, attendance, grade, report card, dan report template.

Finance Service menjadi owner fee type, fee scheme, fee policy, sibling discount, bill, payment, receipt, reconciliation, dan finance approval.

Communication Service menjadi owner announcement, notification, notification template, delivery log, notification preference, dan surat menyurat dasar jika masuk MVP.

Reporting Service menjadi owner dashboard summary/read model/projection di reporting_db, bukan source of truth data operasional.

File metadata, approval, audit log, dan numbering dikelola di masing-masing service pemilik domain pada MVP, dengan shared library/package untuk standardisasi.

Komunikasi antar service menggunakan gRPC untuk kebutuhan sinkron dan RabbitMQ domain event untuk kebutuhan async/projection/notification/reporting.

Service dilarang query langsung ke database service lain.
```

---

## 23. Referensi Keputusan Terkait

Dokumen ini terkait langsung dengan keputusan:

```text
01 — Technical Architecture
02 — Model Isolasi Data
04 — Authentication & Authorization
06 — Approval Workflow
07 — Audit Log dan Data History
12 — Nomor Dokumen dan Format Administrasi
13 — Reporting Strategy
15 — File Management dan Dokumen
16 — Notification Channel
21 — Security Baseline
C  — Entity Relationship dan Data Model MVP
D  — API Contract
E  — Event Contract
I  — Coding Standard dan Project Convention
J  — Migration Tool dan Go Framework Final
```
