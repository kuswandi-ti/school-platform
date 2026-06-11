# Dokumen Arsitektur Teknis Final

**Project:** `school-platform`  
**Jenis Produk:** Platform internal yayasan multi-unit sekolah  
**Target Unit:** TK, SD, SMP, SMA  
**Status Dokumen:** Final architecture baseline untuk MVP  
**Versi:** 1.0  

---

## 1. Ringkasan Eksekutif

`school-platform` adalah sistem manajemen sekolah/yayasan berbasis microservice untuk yayasan yang mengelola beberapa unit sekolah: TK, SD, SMP, dan SMA.

Sistem dirancang sebagai **platform internal yayasan** terlebih dahulu, tetapi tetap dibuat **SaaS-ready** agar memungkinkan pengembangan menjadi platform SaaS di masa depan dengan refactor yang terkendali.

MVP berfokus pada operasional dasar:

- Identity & Access
- School Core
- PPDB
- Academic dasar
- Finance/SPP manual
- Communication/Notification
- Reporting dashboard
- File Management dasar
- Import Excel data awal

Fitur besar seperti HR lengkap, payroll, aset/inventaris lengkap, perpustakaan, BK/UKS detail, LMS penuh, koperasi, global search, payment gateway, WhatsApp, offline write mobile, dan Kubernetes **tidak masuk MVP**.

---

## 2. Tujuan Sistem

Tujuan utama sistem adalah menyediakan platform terpadu untuk yayasan dan unit sekolah dalam mengelola:

1. Data yayasan dan unit sekolah.
2. Data siswa, orang tua/wali, guru, kelas, dan rombel.
3. PPDB dari pendaftaran sampai konversi menjadi siswa.
4. Tagihan SPP, pembayaran manual, upload bukti pembayaran, verifikasi, dan kwitansi.
5. Akademik dasar seperti mata pelajaran, jadwal, absensi, nilai, dan rapor dasar.
6. Pengumuman dan notifikasi ke pengguna.
7. Dashboard yayasan, dashboard sekolah, dashboard guru, dan dashboard orang tua/siswa sederhana.
8. File/dokumen secara private dan aman.
9. Audit, approval, backup, observability, dan security baseline sejak MVP.

---

## 3. Product Model

### 3.1 Keputusan

Produk adalah **Internal Yayasan** untuk tahap awal.

Namun arsitektur tetap disiapkan agar dapat berkembang menjadi SaaS melalui disiplin berikut:

- Setiap data utama menyimpan `foundation_id`.
- Data sekolah menyimpan `school_id`.
- Authorization berbasis RBAC + ABAC/scope.
- Storage path mencantumkan foundation/school context.
- Event payload membawa tenant context.
- Reporting menggunakan projection/read model.
- Tidak ada query lintas database service.

### 3.2 Estimasi Refactor ke SaaS

Jika prinsip di atas dijaga, refactor menuju SaaS diperkirakan terkendali di kisaran **3/10 sampai 5/10**.

Refactor akan lebih besar jika:

- Query tidak konsisten memakai `foundation_id`/`school_id`.
- Authorization hanya di frontend.
- File storage tidak scoped.
- Reporting query langsung ke database operasional.
- Event payload tidak membawa tenant context.

---

## 4. Nama Produk, Domain, dan Identitas Sistem

### 4.1 Nama Produk

Product Display Name belum dikunci permanen dan boleh mengikuti branding yayasan.

Namun nama teknis internal dikunci sebagai:

```text
school-platform
```

### 4.2 Repository

```text
Repository name: school-platform
```

### 4.3 Service Name

Service MVP menggunakan pola:

```text
<domain>-service
```

Service MVP:

```text
api-gateway
identity-service
school-core-service
admission-service
academic-service
finance-service
communication-service
reporting-service
```

### 4.4 Database Name

```text
identity_db
school_core_db
admission_db
academic_db
finance_db
communication_db
reporting_db
```

### 4.5 Domain Environment

```text
local           : localhost
staging web     : staging.<domain>
staging API     : api-staging.<domain>
production web  : app.<domain>
production API  : api.<domain>
```

### 4.6 Identifier Utama

```text
foundation_id
foundation_code
school_id
school_code
school_level
```

### 4.7 Naming Language

- Kode internal memakai Bahasa Inggris.
- UI label memakai Bahasa Indonesia.
- Branding seperti app name, logo, warna utama, nama yayasan, alamat, dan kontak dibuat configurable.

---

## 5. Tech Stack Final

### 5.1 Backend

```text
Language       : Go
HTTP Router    : Chi
Internal RPC   : gRPC + protobuf
DB Driver      : pgx
SQL Generator  : sqlc
Migration Tool : goose
Logger         : slog
Config         : environment variable + envconfig/simple loader
Validation     : go-playground/validator
Testing        : Go testing package + testify
Redis Client   : go-redis
RabbitMQ       : amqp091-go
UUID           : google/uuid
Decimal/Money  : shopspring/decimal
Task Runner    : Makefile
Quality Gate   : golangci-lint, gofmt, go vet, go test
```

Gin tidak digunakan pada MVP agar standar backend konsisten.

Semua query database wajib melalui `sqlc` dan parameterized. Business logic berada di `domain/usecase`, bukan di handler/API Gateway.

Finance calculation tidak boleh menggunakan `float`.

### 5.2 Frontend Web

```text
Framework       : Next.js 14 + TypeScript
UI              : Tailwind CSS + shadcn/ui
Server State    : React Query
Client State    : Zustand
Form            : React Hook Form
Validation      : Zod
```

### 5.3 Frontend Mobile

```text
Framework       : Flutter 3 + Dart
State           : Riverpod
HTTP Client     : Dio + Retrofit
Secure Storage  : Flutter Secure Storage
Push            : FCM
```

### 5.4 Database, Cache, Queue, Storage

```text
Database        : PostgreSQL, 1 instance, database per service
Cache/Session   : Redis
Async Events    : RabbitMQ
Object Storage  : MinIO for local/on-premise, Cloudflare R2 or MinIO for production
```

---

## 6. Arsitektur Umum

### 6.1 Model Arsitektur

Sistem menggunakan arsitektur microservice.

Frontend tidak memanggil microservice langsung. Semua request frontend masuk melalui `api-gateway`.

```text
Next.js / Flutter
        ↓
Custom Go API Gateway
        ↓ REST-to-gRPC
Internal Microservices
        ↓
Database masing-masing service
```

Komunikasi internal:

```text
Synchronous : gRPC/protobuf
Asynchronous: RabbitMQ domain event
```

### 6.2 Service MVP

```text
api-gateway
identity-service
school-core-service
admission-service
academic-service
finance-service
communication-service
reporting-service
```

### 6.3 Prinsip Boundary

```text
1 data utama = 1 owner service
service lain hanya menyimpan reference ID/snapshot terbatas
tidak ada query langsung ke database service lain
reporting menggunakan projection/read model
```

---

## 7. Service Boundary dan Data Ownership

### 7.1 API Gateway

API Gateway bukan owner data domain.

Tanggung jawab:

- External REST/JSON API.
- Validasi token.
- Scope extraction.
- Routing.
- REST-to-gRPC mapping.
- Response standardization.
- Rate limiting dasar.
- Correlation ID.

Tidak boleh:

- Menyimpan business data.
- Menaruh business logic tagihan/nilai/PPDB.
- Query database domain service.

### 7.2 Identity Service

Owner:

- User account.
- Credential.
- Session.
- Refresh token.
- Role.
- Permission.
- Role assignment.

Bukan owner:

- Detail siswa.
- Detail guru.
- Detail orang tua/wali.
- Data akademik.
- Data finance.

### 7.3 School Core Service

Owner:

- Foundation.
- School.
- Academic year.
- Semester.
- Student master.
- Guardian/parent master.
- Teacher master dasar.
- Class/rombel.
- Student-class assignment.
- Teacher assignment.
- Homeroom assignment.

### 7.4 Admission Service

Owner:

- Proses PPDB.
- Applicant sebelum diterima.
- Applicant document.
- Verification.
- Admission decision.

Setelah applicant dikonversi menjadi siswa, School Core menjadi owner student record.

### 7.5 Academic Service

Owner:

- Curriculum.
- Subject.
- Schedule.
- Attendance.
- Grade.
- Report card.
- Report template.

### 7.6 Finance Service

Owner:

- Fee type.
- Fee scheme.
- Fee policy.
- Sibling discount.
- Bill.
- Payment.
- Receipt.
- Reconciliation.
- Finance approval.

Free SPP, diskon, beasiswa, dan sibling discount adalah fee policy di Finance Service, bukan status siswa.

### 7.7 Communication Service

Owner:

- Announcement.
- Notification.
- Notification template.
- Delivery log.
- Notification preference.
- Surat menyurat dasar jika masuk MVP.

### 7.8 Reporting Service

Owner:

- Dashboard summary.
- Read model/projection di `reporting_db`.

Reporting bukan source of truth data operasional.

### 7.9 File, Approval, Audit, Numbering

Pada MVP:

- File metadata dikelola di service pemilik domain.
- Approval dikelola di service pemilik domain.
- Audit log dikelola lokal per service.
- Numbering sequence dikelola di service pemilik domain.

Gunakan shared library/package untuk standardisasi.

---

## 8. Data Isolation dan Database Strategy

### 8.1 Keputusan

Menggunakan PostgreSQL shared infrastructure:

```text
1 PostgreSQL instance
N database per service
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

### 8.2 Prinsip

- Setiap service hanya boleh mengakses database miliknya.
- Setiap service memakai database user/permission khusus.
- Tidak ada foreign key lintas database service.
- Relasi lintas service menggunakan reference ID.
- Validasi lintas service melalui gRPC atau domain event.
- Semua data domain utama menyimpan `foundation_id`.
- Data terkait unit sekolah menyimpan `school_id`.

---

## 9. Data Model MVP

### 9.1 Prinsip Data Model

- ERD dibuat per service, bukan satu ERD monolitik.
- Primary key teknis menggunakan UUID.
- `foundation_id` menjadi tenant boundary utama.
- `school_id` digunakan untuk data terkait unit sekolah.
- Status enum disimpan lowercase snake_case.
- Data historis penting menyimpan snapshot.

### 9.2 Ownership Ringkas

```text
Identity:
- user, credential, session, role, permission, role assignment

School Core:
- foundation, school, academic year, semester, student, guardian, teacher, class, assignment

Admission:
- PPDB dan applicant sebelum converted

Academic:
- curriculum, subject, schedule, attendance, grade, report card

Finance:
- fee type, fee scheme, fee policy, bill, payment, receipt, reconciliation

Communication:
- announcement, notification, template, delivery, preference

Reporting:
- projection/read model untuk dashboard dan summary
```

### 9.3 Standard Tables per Service

Beberapa tabel standar dapat ada pada service pemilik domain:

```text
audit_logs
approval_requests
approval_steps
numbering_sequences
files
import_batches
import_batch_rows
outbox_events
processed_events
```

---

## 10. API Contract

### 10.1 External API

External API menggunakan REST/JSON melalui Custom Go API Gateway.

```text
/api/v1
```

Digunakan oleh:

- Next.js.
- Flutter.

Didokumentasikan di:

```text
packages/openapi
```

### 10.2 Internal API

Internal API menggunakan gRPC/protobuf antar microservice.

Proto contract disimpan di:

```text
packages/proto
```

### 10.3 Standard Response

Success response:

```json
{
  "data": {},
  "meta": null,
  "error": null
}
```

List response:

```json
{
  "data": [],
  "meta": {
    "page": 1,
    "per_page": 20,
    "total": 100,
    "total_pages": 5
  },
  "error": null
}
```

Error response:

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

### 10.4 Standard Error Code

```text
UNAUTHORIZED
FORBIDDEN
VALIDATION_ERROR
NOT_FOUND
CONFLICT
RATE_LIMITED
BUSINESS_RULE_VIOLATION
APPROVAL_REQUIRED
RESOURCE_LOCKED
INTERNAL_ERROR
SERVICE_UNAVAILABLE
```

### 10.5 Context Header

```text
Authorization: Bearer <access_token>
X-Request-ID
X-Correlation-ID
X-School-ID
X-Academic-Year-ID
X-Semester-ID
Idempotency-Key
```

Backend tetap wajib validasi context dari token/scope.

### 10.6 Async Operation

Operasi async seperti import Excel, generate tagihan massal, dan generate dokumen mengembalikan:

```text
202 Accepted
```

dengan `job_id` atau `import_batch_id`.

### 10.7 Idempotency

`Idempotency-Key` disiapkan untuk:

- Generate tagihan.
- Create payment.
- Verify payment.
- Convert applicant to student.
- Publish report card.

---

## 11. Event Contract

### 11.1 RabbitMQ Exchange

```text
Exchange: domain.events
Type    : topic
```

Routing key menggunakan `event_type`.

Format:

```text
domain.entity.action_past_tense
```

Contoh:

```text
finance.payment.verified
academic.report_card.published
school.student.created
```

### 11.2 Event Envelope

Semua event memakai envelope standar:

```json
{
  "event_id": "uuid",
  "event_type": "finance.payment.verified",
  "event_version": 1,
  "source_service": "finance-service",
  "occurred_at": "2026-06-08T10:30:00Z",
  "published_at": "2026-06-08T10:30:01Z",
  "request_id": "req_xxx",
  "correlation_id": "corr_xxx",
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

### 11.3 Event Rules

- Event wajib membawa `foundation_id` dan `school_id` jika relevan.
- Event payload tidak boleh membawa data Confidential detail.
- Event tidak boleh membawa token, password, isi dokumen, atau data sensitif mentah.
- Service penting menggunakan outbox pattern.
- Consumer wajib idempotent dengan menyimpan `processed_event_id`.
- Retry dan DLQ wajib disiapkan untuk consumer penting.
- Event schema disimpan di `packages/events`.

---

## 12. Authentication & Authorization

### 12.1 Authentication

Menggunakan:

```text
JWT access token + rotating refresh token
```

Aturan:

- Access token short-lived.
- Refresh token disimpan aman dan di-hash di database.
- Refresh token dirotasi setiap digunakan.
- Web menggunakan httpOnly secure cookie untuk refresh token.
- Mobile menggunakan secure storage.

### 12.2 Authorization

Menggunakan:

```text
RBAC + ABAC/scope
```

RBAC untuk role dan permission.  
ABAC/scope untuk membatasi akses berdasarkan:

```text
foundation_id
school_id
class_id
subject_id
student_id
employee_id
ownership data
```

API Gateway hanya melakukan validasi token, ekstraksi identity/scope, dan basic guard. Authorization detail tetap dilakukan di masing-masing service.

---

## 13. Role Matrix MVP

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

`wali_kelas` adalah assignment tambahan untuk Guru, bukan role utama terpisah.

Scope default:

```text
Admin Yayasan       : foundation scope
Kepala Sekolah      : school scope
TU/Staff            : school scope
Bendahara Sekolah   : school + finance scope
Guru                : school + class + subject scope
Wali Kelas          : class scope
Orang Tua           : student scope untuk anaknya
Siswa               : self scope
```

Permission format:

```text
domain.resource.action
```

Contoh:

```text
school.student.view
finance.payment.verify
academic.report_card.publish
identity.role.assign
```

---

## 14. Approval Workflow

Approval hanya dipakai untuk aksi penting dan sensitif.

MVP mayoritas approval menggunakan 1 level.

Approver utama:

```text
Kepala Sekolah : aksi operasional level sekolah
Admin Yayasan  : aksi lintas sekolah, role sensitif, pengumuman yayasan, finance bernilai besar
```

Aksi sensitif:

- Void pembayaran.
- Refund.
- Diskon/beasiswa/free SPP.
- Publish rapor.
- Revisi nilai setelah publish.
- Ubah role.
- Export/download Restricted/Confidential.

Approval bisa dikembalikan untuk revisi melalui status:

```text
revision_requested
revised
```

Approval disimpan di service pemilik domain dengan struktur tabel standar `approval_requests` dan `approval_steps`.

---

## 15. Audit Log dan Data History

MVP menggunakan audit log lokal per service.

Setiap service mencatat aksi penting di domainnya sendiri.

Gunakan shared audit library/package untuk standardisasi:

- action naming.
- request_id.
- correlation_id.
- actor.
- entity.
- masking data sensitif.
- payload audit.

Audit action format:

```text
domain.resource.action
```

Jangka panjang, service publish `AuditLogCreated` event ke RabbitMQ untuk konsolidasi audit/reporting.

Audit lokal tetap menjadi source of truth domain.

---

## 16. Data Privacy dan Klasifikasi Data

Klasifikasi data:

```text
public
internal
restricted
confidential
```

### Public

Data tanpa login dan tidak mengandung data pribadi.

Contoh:

- Profil umum.
- Informasi PPDB publik.
- Logo/brosur publik.

### Internal

Data operasional sekolah/yayasan untuk user internal sesuai scope.

### Restricted

Data pribadi dan penting:

- Siswa.
- Guru.
- Orang tua.
- PPDB.
- Nilai.
- Absensi.
- Tagihan.
- Pembayaran.
- Dokumen siswa.

Wajib RBAC + ABAC/scope, private storage, masking, dan audit untuk perubahan/download/export.

### Confidential

Data sangat sensitif:

- BK.
- UKS/kesehatan.
- Payroll.
- Credential/token.
- Dokumen hukum.
- Backup.

Wajib permission khusus, signed URL pendek, reason/approval jika perlu, audit view/download/export, dan masking default.

---

## 17. Academic Year, Semester, dan Calendar

Academic Year dan Semester dibuat global di level Yayasan.

Semua unit TK, SD, SMP, dan SMA memakai referensi `academic_year_id` dan `semester_id` yang sama untuk MVP.

Status:

```text
draft
active
closed
archived
```

Calendar menggunakan model hybrid:

- Event global yayasan.
- Event khusus sekolah.
- Event khusus kelas/jenjang jika diperlukan.

`school_id` nullable:

```text
null  = event yayasan/global
filled = event khusus sekolah
```

Saat semester/tahun ajaran ditutup, data akademik dikunci sebagian dan perubahan setelah penutupan harus melalui approval/audit log.

---

## 18. Payment Flow dan Finance Policy

### 18.1 Payment Flow MVP

MVP menggunakan:

```text
manual payment + upload bukti pembayaran
```

Metode:

- Transfer bank manual.
- Pembayaran tunai/manual via bendahara.

Flow:

```text
Orang tua melihat tagihan
↓
Upload bukti pembayaran
↓
Bendahara verifikasi
↓
Payment verified
↓
Bill status updated
↓
Receipt generated
↓
Audit log + notification
```

Payment gateway belum masuk MVP.

### 18.2 Fee Policy

Free SPP, diskon, beasiswa, dan sibling discount dikelola sebagai `student_fee_policy` di Finance Service.

Status siswa tetap akademik/administratif:

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

Policy fokus awal pada SPP bulanan.

Semua fee policy wajib memiliki:

- Periode berlaku.
- Reason.
- Approval status.
- Audit log.

Generate tagihan wajib menyimpan snapshot:

- base_amount.
- discount_amount.
- final_amount.
- applied_policy.

---

## 19. Numbering dan Dokumen Administratif

Numbering configurable berdasarkan document type dan dikelola per service pada MVP.

Setiap jenis nomor memiliki:

```text
system_key
display_name
prefix
format_pattern
reset_policy
scope_level
padding_length
```

Admin Yayasan boleh mengubah:

- display_name.
- prefix.
- format_pattern.
- padding_length.
- reset_policy.
- scope sesuai aturan.

Tidak boleh mengubah bebas:

- `system_key` untuk dokumen inti.

Nomor yang sudah dipakai tidak boleh digunakan ulang meskipun dokumen dibatalkan/void.

Pembatalan memakai status `cancelled`/`void` dan audit log.

---

## 20. Reporting Strategy

MVP menggunakan Reporting Service dengan `reporting_db` sebagai read model/projection.

Dashboard tidak query langsung ke database operasional seperti:

```text
school_core_db
finance_db
academic_db
admission_db
```

Service operasional publish domain event ke RabbitMQ. Reporting Service consume event dan memperbarui projection.

Dashboard cukup near real-time, delay beberapa detik sampai beberapa menit dapat diterima.

Scheduled rebuild/sync harian atau periodik wajib disiapkan untuk menjaga akurasi.

Metrik MVP:

- Total siswa aktif per unit.
- Total guru/staff per unit.
- PPDB summary per unit.
- Tagihan, pembayaran, tunggakan, collection rate.
- Absensi siswa hari ini.
- Progress input nilai dan publish rapor.
- Pending approval penting.
- Pengumuman/notifikasi penting.

Dashboard minimal:

- Dashboard Yayasan.
- Dashboard Sekolah.
- Dashboard Guru.
- Dashboard Orang Tua/Siswa sederhana.

---

## 21. Search dan Filtering

MVP belum membuat global search lintas modul.

MVP menggunakan search dan filtering lokal per modul/service:

- Siswa.
- Guru.
- PPDB.
- Tagihan.
- Pembayaran.
- Jadwal.
- Nilai.
- Rapor.
- Pengumuman.
- Surat.

Teknologi:

- PostgreSQL index.
- Full-text search jika diperlukan.
- `pg_trgm` jika diperlukan.

Filtering distandarkan:

```text
foundation_id
school_id
academic_year_id
semester_id
status
date range
pagination
sorting
```

Untuk role tertentu, `foundation_id` dan `school_id` diambil dari token/scope, bukan bebas dari frontend.

Global Search masuk phase lanjutan melalui Search Service.

---

## 22. File Management dan Dokumen

Object storage menggunakan S3-compatible API:

```text
Local/on-premise : MinIO
Production       : Cloudflare R2 atau MinIO
```

Semua file private by default kecuali file public non-PII.

Akses file private:

```text
backend authorization + RBAC + ABAC/scope + signed URL
```

Signed URL expiry:

```text
Internal     : sekitar 30 menit
Restricted   : sekitar 10 menit
Confidential : sekitar 3 menit
```

Metadata file disimpan di database service terkait.

Upload validasi:

- MIME type.
- Extension.
- Size.
- Permission.
- Scope.
- Classification policy.

File MVP:

```text
PDF
JPG/JPEG
PNG
DOCX
XLSX
CSV jika diperlukan
```

File resmi seperti rapor, surat, kwitansi, invoice, dan slip gaji tidak boleh dioverwrite; revisi membuat versi baru.

Delete file memakai soft delete/archive. Physical purge berdasarkan retention policy dan approval/admin job.

---

## 23. Notification Channel

MVP Notification Channel:

```text
in-app notification
FCM push notification
email terbatas
```

Email terbatas untuk:

- Auth.
- Reset password.
- Undangan user.
- Status penting PPDB.
- Dokumen/receipt resmi tertentu.

WhatsApp belum masuk MVP dan disiapkan sebagai phase lanjutan.

SMS bukan prioritas, hanya fallback emergency jika diperlukan.

Semua notifikasi dikirim melalui Communication/Notification Service berbasis event dari RabbitMQ.

Notifikasi penting tidak boleh dimatikan sepenuhnya:

- Tagihan.
- Pembayaran ditolak.
- Absensi alfa.
- Approval penting.
- Security.
- Emergency.

Confidential data tidak boleh dikirim detail melalui notifikasi.

---

## 24. Offline Mode Mobile

MVP menggunakan model online-only untuk seluruh fitur utama.

Aplikasi mobile boleh menyimpan cache ringan/read-only:

- Profil terakhir.
- Jadwal terakhir.
- Pengumuman terakhir.
- Tagihan terakhir.

Tidak dibuat di MVP:

- Offline write.
- Sync queue.
- Conflict resolution.
- Local database kompleks.

Offline mode phase lanjutan terbatas untuk:

- Input absensi guru.
- Cache jadwal.
- Cache pengumuman.
- Draft nilai guru jika benar-benar dibutuhkan.

Fitur sensitif tetap online-only:

- Pembayaran.
- Approval.
- Publish rapor.
- Revisi nilai.
- BK/UKS.
- Payroll.
- Akses confidential document.

---

## 25. CI/CD dan Deployment

MVP menggunakan:

- Monorepo.
- Docker Compose untuk development dan staging awal.
- GitHub Actions.
- Container registry.
- Manual approval untuk production.
- Kubernetes belum masuk MVP.

Branch strategy:

```text
feature/* → develop → staging → main/production
```

CI/CD mencakup:

- lint.
- test.
- build.
- Docker image build.
- push image ke registry.
- deploy ke staging.
- production deploy dengan manual approval.

Setiap microservice memiliki migration sendiri dan hanya migrate database miliknya sendiri.

Rollback menggunakan image tag berbasis commit SHA dan migration backward-compatible.

---

## 26. Observability

MVP wajib mencakup:

- Structured JSON logging.
- request_id.
- correlation_id.
- health check.
- readiness check.
- basic metrics.
- centralized log.
- alert dasar.

Stack awal:

```text
Logging : slog + Loki + Grafana
Metrics : Prometheus + Grafana
Tracing : OpenTelemetry prepared
Alert   : Grafana Alerting
```

Correlation ID wajib diteruskan melalui:

- HTTP headers.
- gRPC metadata.
- RabbitMQ event headers.

Application log tidak boleh berisi data sensitif mentah.

Audit log tetap terpisah dari application log.

---

## 27. Backup, Restore, dan Disaster Recovery

MVP target:

```text
RPO maksimal 24 jam
RTO 4–8 jam
```

Backup mencakup:

- Seluruh PostgreSQL database per service.
- Object storage MinIO/R2.
- Konfigurasi deployment.
- Secrets terenkripsi.
- RabbitMQ definitions jika diperlukan.

Strategi:

```text
PostgreSQL backup harian menggunakan pg_dump per database
Object storage backup/sync harian
Backup disimpan di lokasi terpisah dari server production
Retensi daily 30 hari, weekly 12 minggu, monthly 12 bulan
```

Restore test wajib dilakukan minimal bulanan.

Backup diperlakukan sebagai Confidential data.

---

## 28. Security Baseline

Security wajib masuk MVP.

Baseline:

- HTTPS wajib di production.
- JWT short-lived + rotating refresh token.
- Password hashing Argon2id atau bcrypt.
- Refresh token disimpan hash.
- RBAC + ABAC/scope di service internal.
- Object-level authorization untuk resource by ID.
- Rate limiting endpoint sensitif.
- Input validation semua request.
- Query parameterized melalui pgx/sqlc.
- CORS ketat.
- Security headers untuk web.
- File private by default dan akses via signed URL.
- Secrets tidak boleh masuk repository.
- Database user dipisah per service dengan least privilege.
- Service internal/database/message broker tidak diekspos publik.
- Application log tidak boleh berisi data sensitif mentah.
- Audit log untuk aksi sensitif.
- Export Restricted/Confidential wajib permission, scope check, dan audit.
- Backup terenkripsi dan diperlakukan sebagai Confidential.

Phase lanjutan:

- 2FA admin.
- Secret manager.
- Virus scanning file.
- Container image scanning/SBOM.
- WAF/API gateway product seperti Kong.
- Penetration test berkala.
- Advanced anomaly detection.

---

## 29. Data Migration dan Import Awal

Import Excel masuk MVP untuk migrasi data awal.

Scope MVP:

- Siswa.
- Orang tua/wali siswa.
- Guru.
- Kelas/rombel.
- Assignment siswa ke kelas melalui `class_code`.
- Assignment wali kelas/guru mapel opsional jika data siap.

Flow:

```text
download template
↓
upload file
↓
validasi
↓
preview hasil
↓
confirm import
↓
process import
↓
import report
```

Import tidak boleh langsung insert tanpa validasi dan preview.

Mode MVP:

- create_only.
- upsert terbatas jika diperlukan.

File import diperlakukan sebagai Restricted data.

Import nilai, pembayaran historis, payroll, aset, perpustakaan, BK/UKS, alumni, dan koperasi ditunda.

---

## 30. UI/UX dan Multi-Role Navigation

Web digunakan untuk:

- Admin Yayasan.
- Kepala Sekolah.
- TU/Staff.
- Bendahara Sekolah.
- Guru.

Mobile digunakan untuk:

- Orang Tua/Wali Murid.
- Siswa.
- Guru fitur cepat.

Global context wajib:

```text
selected_foundation
selected_school
selected_academic_year
selected_semester
```

Tambahan context:

```text
selected_class          : Guru/Wali Kelas
selected_subject        : Guru
selected_child          : Orang Tua
selected_billing_month  : Bendahara/Finance
```

Menu dan navigasi berbasis role, permission, dan scope.

Frontend boleh menyembunyikan menu, tetapi backend tetap wajib permission/scope check.

Aksi sensitif wajib confirmation dialog, reason jika diperlukan, approval flow, dan audit log.

---

## 31. Legal dan Kepatuhan Data

MVP harus memiliki baseline legal/compliance untuk perlindungan data pribadi.

Sistem mengacu pada prinsip UU PDP Indonesia:

- Pengumpulan data harus punya tujuan jelas.
- Akses dibatasi sesuai role/scope.
- Data sensitif diproteksi lebih ketat.
- Pemrosesan data transparan.

Dokumen minimal:

- Privacy Policy.
- Terms of Use/Ketentuan Penggunaan.
- Data Retention Policy.
- Access Control Policy.
- SOP permintaan/perubahan/penghapusan/arsip data.

Consent orang tua/wali perlu disiapkan untuk:

- Pemrosesan data anak.
- PPDB.
- Komunikasi/notifikasi.
- Publikasi foto/dokumentasi jika digunakan.

---

## 32. Definisi MVP

MVP adalah platform internal yayasan multi-unit yang dapat digunakan untuk operasional dasar TK, SD, SMP, dan SMA.

MVP mencakup:

- API Gateway.
- Identity & Access.
- School Core.
- PPDB.
- Academic dasar.
- Finance/SPP manual.
- Communication/Notification.
- Reporting dashboard.
- File Management dasar.
- Import Excel data awal.

Pengguna MVP:

- Admin Yayasan.
- Kepala Sekolah.
- TU/Staff.
- Bendahara Sekolah.
- Guru/Wali Kelas.
- Orang Tua/Wali Murid.
- Siswa.

MVP dianggap selesai jika yayasan dan minimal satu unit sekolah pilot bisa menjalankan alur:

```text
login → kelola siswa/guru/kelas → PPDB → generate tagihan → pembayaran manual → absensi/nilai/rapor dasar → pengumuman/notifikasi → dashboard/laporan ringkas
```

---

## 33. Team Structure dan Workflow

Tim MVP terdiri dari 4 peran:

```text
Backend Developer
Frontend Developer
QA
Infrastructure/DevOps
```

Backend:

- Go microservices.
- API Gateway.
- gRPC/protobuf.
- Database migration.
- Business logic.
- Domain event.
- Authorization service-side.
- Test backend.

Frontend:

- Next.js web admin.
- Flutter mobile.
- UI/UX implementation.
- Role-based navigation.
- Form validation.
- State management.
- API integration.

QA:

- Test scenario.
- Test case.
- UAT checklist.
- Regression test.
- API testing.
- Bug verification.
- Release validation.

Infrastructure/DevOps:

- Monorepo structure.
- Docker Compose.
- GitHub Actions.
- Container registry.
- Staging/production deployment.
- Secrets.
- Observability.
- Backup/restore.
- Server security baseline.

Workflow:

```text
feature/* → develop → staging → main/production
```

Production deploy hanya dari main dengan manual approval.

QA sign-off wajib sebelum production release.

---

## 34. GitHub Repository Rules

Branch utama:

```text
develop
staging
main
```

Aturan:

- Semua perubahan melalui Pull Request.
- Semua PR wajib melewati CI.
- Semua PR wajib review.
- Force push/delete branch dilarang untuk branch utama.

Branch `develop`:

- Menerima PR dari `feature/*`.
- Wajib CI pass.
- Minimal 1 approval.

Branch `staging`:

- Menerima PR dari `develop`.
- Deploy staging/QA/UAT.
- Wajib CI pass.
- Wajib QA sign-off sebelum lanjut ke main.

Branch `main`:

- Hanya untuk production.
- Menerima PR dari `staging`.
- Wajib CI pass.
- Wajib QA sign-off.
- Wajib Infrastructure/DevOps approval.
- Wajib manual approval melalui GitHub Environment production.

Gunakan:

- GitHub Environments.
- CODEOWNERS.
- Branch naming convention.
- PR template.
- Release tag untuk production deployment.

---

## 35. Local Development Standard

Development dilakukan di lokal terlebih dahulu menggunakan monorepo dan Docker Compose.

Docker Compose menyediakan:

- PostgreSQL.
- Redis.
- RabbitMQ.
- MinIO.
- Optional Mailpit/Grafana/Loki.

Developer dapat menjalankan full stack atau subset service.

Database lokal:

```text
identity_db
school_core_db
admission_db
academic_db
finance_db
communication_db
reporting_db
```

Seed data lokal wajib untuk:

- 1 yayasan.
- Unit TK/SD/SMP/SMA.
- Role MVP.
- User dummy.
- Tahun ajaran aktif.
- Semester aktif.
- Kelas.
- Siswa.
- Guru.
- PPDB.
- Tagihan dummy.

External provider seperti email, FCM, payment gateway, dan WhatsApp menggunakan mock/log-only di lokal.

Staging digunakan untuk integrasi, QA, dan UAT, bukan tempat development utama.

---

## 36. UI Screen List dan User Flow

Screen list dan user flow MVP wajib didokumentasikan sebelum implementasi frontend besar.

Web MVP mencakup:

- Dashboard.
- Data siswa/guru/kelas.
- Import data.
- PPDB.
- Finance/SPP.
- Akademik dasar.
- Absensi.
- Nilai/rapor.
- Pengumuman.
- Approval.
- Laporan.
- Pengaturan dasar.

Mobile MVP mencakup:

- Login.
- Home.
- Notifikasi.
- Tagihan.
- Upload bukti pembayaran.
- Absensi.
- Nilai/rapor published.
- Pengumuman.
- Jadwal.
- Profil.

User flow utama:

- Login dan context selection.
- Kelola siswa/guru/kelas.
- Import Excel.
- PPDB sampai konversi siswa.
- Generate tagihan.
- Pembayaran manual dan upload bukti.
- Verifikasi pembayaran.
- Void pembayaran dengan approval.
- Input absensi.
- Input nilai dan publish rapor.
- Pengumuman/notifikasi.
- Dashboard reporting.

Setiap screen wajib punya:

- Role.
- Platform.
- Route.
- Purpose.
- Permission.
- Global context.
- API.
- Action.
- Validation.
- Empty/error state.
- Audit requirement jika sensitif.

---

## 37. Test Plan dan Acceptance Criteria

Setiap modul dan fitur MVP wajib punya test plan dan acceptance criteria sebelum implementasi.

Testing MVP mencakup:

- Unit test.
- Integration test.
- API test.
- Permission/scope test.
- Event test.
- UI flow test.
- Regression test.
- UAT checklist.
- Security baseline test.

Modul wajib test plan:

- Identity & Access.
- School Core.
- Import Excel.
- PPDB.
- Finance/SPP.
- Academic/Absensi/Nilai/Rapor.
- Communication/Notification.
- Reporting Dashboard.
- File Management.
- Approval.
- Audit Log.

Test permission/scope menjadi prioritas utama.

Production tidak boleh release jika masih ada Critical/High bug terbuka pada flow utama MVP.

---

## 38. Sprint Roadmap

Urutan sprint MVP:

```text
Sprint 0  : Project Foundation
Sprint 1  : Identity & Access
Sprint 2  : School Core
Sprint 3  : File Management + Import Excel
Sprint 4  : PPDB
Sprint 5  : Finance / SPP
Sprint 6  : Academic Dasar
Sprint 7  : Report Card / E-Rapor Dasar
Sprint 8  : Communication / Notification
Sprint 9  : Reporting Dashboard
Sprint 10 : Security, Observability, Backup, dan UAT Hardening
```

Milestone:

```text
Milestone 1: Platform Foundation
Milestone 2: Admission & Finance
Milestone 3: Academic & Communication
Milestone 4: Reporting & Production Readiness
```

Setiap sprint wajib memiliki:

- Scope jelas.
- Out of scope.
- Acceptance criteria.
- Test plan.
- Definition of Done.
- Task kecil yang bisa dikerjakan AI Agent.

Security, audit, permission/scope, logging, correlation_id, dan test tidak ditunda sampai akhir.

---

## 39. Coding Standard dan Project Convention

Project menggunakan coding standard dan convention terpusat.

Kode internal memakai Bahasa Inggris. UI label memakai Bahasa Indonesia.

Struktur monorepo:

```text
apps/
services/
packages/
infra/
deploy/
docs/
scripts/
```

Struktur Go service:

```text
cmd
internal/app
internal/config
internal/domain
internal/usecase
internal/repository
internal/transport
internal/event
internal/authz
internal/audit
internal/db
```

Business logic berada di domain/usecase.

API Gateway bukan tempat business logic.

Database naming:

- snake_case.
- table plural.
- UUID primary key.
- `foundation_id`.
- `school_id`.
- `created_at`.
- `updated_at`.
- status enum lowercase snake_case.

AI Agent wajib mengikuti `docs/AI_AGENT_RULES.md`.

---

## 40. AI Agent Development Rules

AI Agent tidak boleh:

- Query lintas database service.
- Menaruh business logic di API Gateway.
- Membuat endpoint tanpa permission/scope check.
- Menghapus `foundation_id` atau `school_id`.
- Membuat file public untuk data pribadi.
- Menulis token/password/data confidential ke log.
- Mengubah API/proto/event contract tanpa update dokumen.
- Membuat fitur di luar scope task.
- Mengabaikan test.

AI Agent wajib:

- Mengikuti service boundary.
- Membuat test.
- Menambahkan audit untuk aksi sensitif.
- Menggunakan response/error format standar.
- Menggunakan shared package untuk logging, audit, numbering, event, dan file jika tersedia.
- Menjelaskan file yang diubah dan cara test.

---

## 41. Non-Goal MVP

Fitur berikut tidak masuk MVP:

- Payroll.
- HR lengkap.
- Asset/inventory lengkap.
- Library/perpustakaan lengkap.
- BK/UKS detail.
- LMS penuh.
- Alumni/tracer.
- Koperasi.
- Global Search.
- Payment Gateway.
- WhatsApp.
- Offline Write Mobile.
- Kubernetes.
- Advanced analytics.

---

## 42. Kesimpulan Arsitektur

MVP `school-platform` menggunakan arsitektur microservice dengan Go backend, Next.js web admin, Flutter mobile, PostgreSQL database per service, RabbitMQ event-driven communication, Redis, dan S3-compatible object storage.

Keputusan arsitektur utama:

- Platform internal yayasan, SaaS-ready.
- Custom Go API Gateway dulu; Kong nanti jika kebutuhan API management nyata.
- External API REST/JSON via Gateway.
- Internal API gRPC/protobuf.
- Domain event via RabbitMQ.
- Database per service.
- Reporting via projection/read model.
- Auth JWT + rotating refresh token.
- Authorization RBAC + ABAC/scope.
- File private by default.
- Security, audit, observability, backup, dan test masuk MVP.
- AI Agent bekerja dengan guardrails eksplisit.

Dokumen ini menjadi baseline utama untuk dokumen turunan berikut:

```text
02-service-boundary.md
03-data-model-mvp.md
04-api-contract.md
05-event-contract.md
06-ui-screen-user-flow.md
07-test-plan-acceptance-criteria.md
08-coding-standard.md
09-ai-agent-rules.md
10-sprint-backlog-mvp.md
```

