# 04 — API Contract

**Project:** `school-platform`  
**Document Type:** Technical Contract / API Standard  
**Status:** Final Decision Draft  
**Scope:** MVP  
**Last Updated:** 2026-06-08

---

## 1. Tujuan Dokumen

Dokumen ini mendefinisikan standar **API Contract** untuk sistem manajemen sekolah/yayasan berbasis microservice.

API contract digunakan sebagai acuan untuk:

- Backend Developer
- Frontend Developer
- Mobile Developer
- QA
- Infrastructure/DevOps
- AI Agent

Dokumen ini wajib dirujuk ketika membuat atau mengubah:

- REST API external melalui API Gateway
- gRPC/protobuf internal antar microservice
- format response
- format error
- pagination
- filtering
- sorting
- auth/context header
- idempotency
- async operation
- file upload/download API
- OpenAPI dan proto contract

---

## 2. Keputusan Utama

API Contract MVP dibagi menjadi dua lapisan:

```text
External API:
- REST/JSON melalui Custom Go API Gateway
- Dipakai oleh Next.js web admin dan Flutter mobile app
- Menggunakan prefix /api/v1
- Didokumentasikan dengan OpenAPI di packages/openapi

Internal API:
- gRPC/protobuf antar microservice
- Proto contract disimpan di packages/proto
- API Gateway melakukan REST-to-gRPC mapping
```

Keputusan penting:

```text
Frontend tidak boleh memanggil microservice langsung.
Microservice internal tidak expose REST publik langsung.
Semua akses frontend harus melalui API Gateway.
```

---

## 3. Prinsip API Contract

Prinsip dasar:

```text
1. Frontend hanya consume REST/JSON dari API Gateway.
2. API Gateway meneruskan request ke service internal melalui gRPC.
3. Microservice internal tidak expose REST publik langsung.
4. Semua request membawa request_id dan correlation_id.
5. Semua response memakai format standar.
6. Semua error memakai format standar.
7. Pagination, filtering, dan sorting harus konsisten.
8. Authorization tetap dicek di service internal, bukan hanya API Gateway.
9. Semua resource berbasis ID wajib object-level authorization.
10. Operasi rawan duplikasi wajib mendukung idempotency.
```

---

## 4. Arsitektur API

### 4.1 External API Flow

```text
Next.js / Flutter
↓
API Gateway REST/JSON
↓
gRPC internal
↓
Domain Service
↓
Database service terkait
```

Contoh:

```text
Flutter Parent App
↓
GET /api/v1/finance/bills
↓
API Gateway
↓
FinanceService.ListBills gRPC
↓
finance_db
```

### 4.2 Internal API Flow

```text
Service A
↓
gRPC
↓
Service B
```

Contoh:

```text
Admission Service
↓
SchoolCoreService.CreateStudentFromApplicant
↓
School Core Service
```

---

## 5. External REST API Standard

External API memakai format:

```text
/api/v1/{domain}/{resource}
```

Contoh:

```text
/api/v1/auth/login
/api/v1/students
/api/v1/admissions/applicants
/api/v1/finance/bills
/api/v1/academic/report-cards
/api/v1/announcements
/api/v1/dashboard/school
```

### 5.1 Versioning

Versioning API menggunakan URL prefix:

```text
/api/v1
```

Aturan:

```text
- MVP hanya memakai v1.
- Breaking change besar harus dibuat di versi baru.
- Tambah field response yang backward-compatible boleh tetap di v1.
```

---

## 6. Standard Response Format

Semua response external API wajib menggunakan format berikut.

### 6.1 Success Response — Single Object

```json
{
  "data": {},
  "meta": null,
  "error": null
}
```

Contoh:

```json
{
  "data": {
    "id": "5e8d9b44-3c8e-4ad1-a9c2-7fd00fb38a21",
    "full_name": "Andi Pratama",
    "status": "active"
  },
  "meta": null,
  "error": null
}
```

### 6.2 Success Response — List

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

### 6.3 Success Response — Action

```json
{
  "data": {
    "success": true
  },
  "meta": null,
  "error": null
}
```

### 6.4 Created Response

Untuk create resource:

```http
HTTP/1.1 201 Created
```

```json
{
  "data": {
    "id": "uuid",
    "status": "active"
  },
  "meta": null,
  "error": null
}
```

### 6.5 Accepted Response untuk Async Operation

Untuk proses async seperti import Excel atau generate tagihan massal:

```http
HTTP/1.1 202 Accepted
```

```json
{
  "data": {
    "job_id": "uuid",
    "status": "queued"
  },
  "meta": null,
  "error": null
}
```

Atau untuk import:

```json
{
  "data": {
    "import_batch_id": "uuid",
    "status": "pending"
  },
  "meta": null,
  "error": null
}
```

---

## 7. Standard Error Format

Semua error external API wajib menggunakan format:

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

### 7.1 Validation Error

```json
{
  "data": null,
  "meta": null,
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Data yang dikirim tidak valid.",
    "details": {
      "full_name": ["Nama siswa wajib diisi."],
      "birth_date": ["Tanggal lahir tidak valid."]
    }
  }
}
```

### 7.2 Business Rule Error

```json
{
  "data": null,
  "meta": null,
  "error": {
    "code": "BUSINESS_RULE_VIOLATION",
    "message": "Tagihan untuk periode ini sudah pernah dibuat.",
    "details": {
      "billing_period": "2026-07"
    }
  }
}
```

### 7.3 Resource Locked Error

```json
{
  "data": null,
  "meta": null,
  "error": {
    "code": "RESOURCE_LOCKED",
    "message": "Rapor sudah dipublish dan dikunci.",
    "details": {
      "status": "locked"
    }
  }
}
```

---

## 8. Standard Error Codes

Error code standar:

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

Penjelasan:

| Error Code | Makna |
|---|---|
| `UNAUTHORIZED` | Token tidak valid, expired, atau user belum login |
| `FORBIDDEN` | User login tetapi tidak memiliki permission/scope |
| `VALIDATION_ERROR` | Request tidak valid |
| `NOT_FOUND` | Resource tidak ditemukan atau di luar scope user |
| `CONFLICT` | Konflik data, misalnya nomor/record duplikat |
| `RATE_LIMITED` | Request terkena rate limit |
| `BUSINESS_RULE_VIOLATION` | Melanggar aturan bisnis |
| `APPROVAL_REQUIRED` | Aksi membutuhkan approval |
| `RESOURCE_LOCKED` | Resource sudah dikunci dan tidak boleh diubah |
| `INTERNAL_ERROR` | Error internal tidak terduga |
| `SERVICE_UNAVAILABLE` | Service dependency tidak tersedia |

Catatan keamanan:

```text
Untuk resource yang berada di luar scope user, sistem boleh mengembalikan NOT_FOUND agar tidak membocorkan keberadaan data.
```

---

## 9. HTTP Status Code Standard

| HTTP Status | Penggunaan |
|---|---|
| `200 OK` | Request berhasil |
| `201 Created` | Resource berhasil dibuat |
| `202 Accepted` | Proses async diterima |
| `204 No Content` | Delete/aksi berhasil tanpa body response |
| `400 Bad Request` | Request malformed atau business precondition gagal |
| `401 Unauthorized` | Token tidak valid/belum login |
| `403 Forbidden` | Tidak memiliki permission/scope |
| `404 Not Found` | Resource tidak ditemukan/di luar scope |
| `409 Conflict` | Konflik data/duplikasi |
| `422 Unprocessable Entity` | Validasi gagal |
| `423 Locked` | Resource terkunci |
| `429 Too Many Requests` | Rate limit |
| `500 Internal Server Error` | Error internal |
| `503 Service Unavailable` | Service dependency gagal |

---

## 10. Authentication Header

Semua endpoint protected wajib menggunakan:

```http
Authorization: Bearer <access_token>
```

Web:

```text
- Access token digunakan untuk API call.
- Refresh token disimpan dalam httpOnly secure cookie.
```

Mobile:

```text
- Access token dikirim sebagai Bearer token.
- Refresh token disimpan di secure storage.
```

---

## 11. Request ID dan Correlation ID

Header standar:

```http
X-Request-ID: req_xxx
X-Correlation-ID: corr_xxx
```

Aturan:

```text
- Jika frontend tidak mengirim, API Gateway wajib membuat otomatis.
- request_id unik untuk satu request external.
- correlation_id mengikuti alur lintas service dan event.
- correlation_id wajib diteruskan ke gRPC metadata dan RabbitMQ event.
```

---

## 12. Context Header

Frontend boleh mengirim selected context melalui header.

Header yang disarankan:

```http
X-School-ID: <school_id>
X-Academic-Year-ID: <academic_year_id>
X-Semester-ID: <semester_id>
```

Tambahan context jika diperlukan:

```http
X-Class-ID: <class_id>
X-Subject-ID: <subject_id>
X-Student-ID: <student_id>
X-Billing-Month: 2026-07
```

Aturan:

```text
- Context dari frontend tidak boleh dipercaya mentah.
- Backend wajib validasi context terhadap token/scope/assignment.
- foundation_id mayoritas berasal dari token/session context, bukan query bebas frontend.
```

Role-specific behavior:

| Role | Context Rule |
|---|---|
| Admin Yayasan | Boleh memilih school_id sesuai foundation scope |
| Kepala Sekolah | school_id harus sesuai scope token |
| TU/Staff | school_id harus sesuai scope token |
| Bendahara | school_id + billing month divalidasi |
| Guru | class_id/subject_id divalidasi dari assignment |
| Orang Tua | student_id harus anaknya sendiri |
| Siswa | student_id harus dirinya sendiri |

---

## 13. Pagination Standard

Query parameter:

```text
page
per_page
sort_by
sort_direction
```

Default:

```text
page = 1
per_page = 20
max per_page = 100
sort_direction = asc | desc
```

Contoh:

```http
GET /api/v1/students?page=1&per_page=20&sort_by=full_name&sort_direction=asc
```

Response meta:

```json
{
  "meta": {
    "page": 1,
    "per_page": 20,
    "total": 150,
    "total_pages": 8
  }
}
```

---

## 14. Filtering Standard

Query parameter umum:

```text
search
school_id
academic_year_id
semester_id
class_id
student_id
status
date_from
date_to
created_from
created_to
```

Contoh student search:

```http
GET /api/v1/students?search=andi&class_id=uuid&status=active
```

Contoh finance:

```http
GET /api/v1/finance/bills?billing_period=2026-07&status=unpaid
```

Contoh PPDB:

```http
GET /api/v1/admissions/applicants?admission_period_id=uuid&status=verified
```

Aturan:

```text
- Filtering harus tetap mengikuti scope user.
- User tidak boleh bebas memilih foundation_id/school_id di luar scope.
- Jika filter tidak valid, kembalikan VALIDATION_ERROR.
```

---

## 15. Sorting Standard

Parameter:

```text
sort_by
sort_direction
```

Contoh:

```http
GET /api/v1/finance/payments?sort_by=created_at&sort_direction=desc
```

Aturan:

```text
- sort_by hanya boleh dari whitelist field per endpoint.
- Jangan langsung memasukkan sort_by user ke SQL mentah.
- Query tetap harus parameterized/aman.
```

---

## 16. Date, Time, Currency, dan Amount

Format standar:

```text
Date: YYYY-MM-DD
Datetime: RFC3339 / ISO-8601
Time: HH:mm:ss
Currency: IDR
Amount API: number rupiah, contoh 500000
Database amount: NUMERIC(14,2)
```

Contoh:

```json
{
  "amount": 500000,
  "currency": "IDR",
  "due_date": "2026-07-10",
  "paid_at": "2026-07-08T10:30:00Z"
}
```

Aturan finance:

```text
Finance calculation tidak boleh memakai float.
Gunakan decimal/money library di backend.
```

---

## 17. Idempotency Standard

Operasi rawan duplikasi wajib mendukung idempotency.

Header:

```http
Idempotency-Key: unique-client-generated-key
```

Endpoint yang wajib disiapkan:

```text
- generate tagihan
- create payment
- verify payment
- import confirm
- convert applicant to student
- publish report card
- payment gateway callback nanti
```

Aturan:

```text
- Jika request dengan Idempotency-Key yang sama dikirim ulang, sistem tidak boleh membuat data duplikat.
- Response boleh mengembalikan hasil request pertama.
- Idempotency key harus scoped by actor/foundation/school/action.
```

Contoh:

```http
POST /api/v1/finance/bills/generate
Idempotency-Key: generate-spp-sd-2026-07-001
```

---

## 18. Async Operation API

Untuk operasi async:

```text
- import Excel
- generate tagihan massal
- generate dokumen/PDF
- rebuild reporting projection
```

Gunakan:

```http
HTTP/1.1 202 Accepted
```

Response:

```json
{
  "data": {
    "job_id": "uuid",
    "status": "queued"
  },
  "meta": null,
  "error": null
}
```

Cek status:

```http
GET /api/v1/jobs/{job_id}
```

Atau domain-specific:

```http
GET /api/v1/imports/{import_batch_id}
GET /api/v1/finance/bill-generation-jobs/{job_id}
```

Rekomendasi MVP:

```text
Gunakan endpoint domain-specific untuk status import dan generate tagihan agar authorization lebih jelas.
```

---

## 19. File Upload API

File upload dilakukan melalui endpoint domain-specific.

Contoh:

```http
POST /api/v1/admissions/applicants/{applicant_id}/documents
POST /api/v1/finance/payments/{payment_id}/proofs
POST /api/v1/students/{student_id}/documents
```

Request:

```http
Content-Type: multipart/form-data
```

Response:

```json
{
  "data": {
    "file_id": "uuid",
    "original_filename": "akta.pdf",
    "mime_type": "application/pdf",
    "classification": "restricted",
    "status": "active"
  },
  "meta": null,
  "error": null
}
```

Aturan:

```text
- File private by default.
- Validasi MIME, extension, size.
- Metadata file disimpan di service pemilik domain.
- File sensitif Restricted/Confidential wajib audit untuk download/export.
```

---

## 20. Signed URL API

Signed URL harus domain-specific agar authorization jelas.

Contoh:

```http
POST /api/v1/students/{student_id}/documents/{file_id}/signed-url
POST /api/v1/finance/receipts/{receipt_id}/signed-url
POST /api/v1/academic/report-cards/{report_card_id}/signed-url
```

Response:

```json
{
  "data": {
    "url": "https://signed-url.example",
    "expires_at": "2026-07-08T10:40:00Z"
  },
  "meta": null,
  "error": null
}
```

Aturan expiry:

```text
Internal: sekitar 30 menit
Restricted: sekitar 10 menit
Confidential: sekitar 3 menit
```

---

## 21. OpenAPI Contract

External REST API wajib didokumentasikan dengan OpenAPI.

Lokasi:

```text
packages/openapi/
```

File utama:

```text
packages/openapi/api-gateway.v1.yaml
```

Aturan:

```text
- Setiap endpoint external wajib masuk OpenAPI.
- Request/response schema harus jelas.
- Error schema harus memakai standar dokumen ini.
- Frontend dan QA memakai OpenAPI sebagai referensi.
- OpenAPI wajib diperbarui jika endpoint berubah.
```

Struktur yang disarankan:

```text
packages/openapi/
├── api-gateway.v1.yaml
├── schemas/
│   ├── common.yaml
│   ├── auth.yaml
│   ├── school-core.yaml
│   ├── admission.yaml
│   ├── finance.yaml
│   ├── academic.yaml
│   ├── communication.yaml
│   └── reporting.yaml
└── examples/
```

---

## 22. Internal gRPC Contract

Internal service menggunakan gRPC/protobuf.

Lokasi:

```text
packages/proto/
```

Struktur:

```text
packages/proto/identity/v1/identity.proto
packages/proto/schoolcore/v1/school_core.proto
packages/proto/admission/v1/admission.proto
packages/proto/academic/v1/academic.proto
packages/proto/finance/v1/finance.proto
packages/proto/communication/v1/communication.proto
packages/proto/reporting/v1/reporting.proto
```

Package naming:

```proto
package schoolplatform.identity.v1;
package schoolplatform.finance.v1;
```

Service naming:

```proto
service IdentityService {}
service SchoolCoreService {}
service AdmissionService {}
service AcademicService {}
service FinanceService {}
service CommunicationService {}
service ReportingService {}
```

---

## 23. gRPC Context Standard

Setiap call internal harus membawa request context melalui metadata atau message context.

Metadata wajib:

```text
x-request-id
x-correlation-id
x-actor-user-id
x-foundation-id
x-school-id
x-role
x-scope
```

Proto context yang direkomendasikan:

```proto
message RequestContext {
  string request_id = 1;
  string correlation_id = 2;
  string actor_user_id = 3;
  string foundation_id = 4;
  string school_id = 5;
  repeated string roles = 6;
  map<string, string> scope = 7;
}
```

Contoh:

```proto
message GetStudentRequest {
  RequestContext context = 1;
  string student_id = 2;
}
```

Aturan:

```text
- API Gateway meneruskan context dari token ke gRPC metadata/context.
- Service internal tetap wajib validasi authorization sesuai domain.
- Jangan mengandalkan API Gateway sebagai satu-satunya authorization layer.
```

---

## 24. gRPC Error Mapping

API Gateway wajib memetakan gRPC error ke HTTP error standar.

| gRPC Code | HTTP Status | Error Code |
|---|---:|---|
| `Unauthenticated` | 401 | `UNAUTHORIZED` |
| `PermissionDenied` | 403 | `FORBIDDEN` |
| `NotFound` | 404 | `NOT_FOUND` |
| `InvalidArgument` | 422 | `VALIDATION_ERROR` |
| `AlreadyExists` | 409 | `CONFLICT` |
| `FailedPrecondition` | 400 / 423 | `BUSINESS_RULE_VIOLATION` / `RESOURCE_LOCKED` |
| `ResourceExhausted` | 429 | `RATE_LIMITED` |
| `Internal` | 500 | `INTERNAL_ERROR` |
| `Unavailable` | 503 | `SERVICE_UNAVAILABLE` |

---

## 25. Proto Compatibility Rules

Aturan perubahan protobuf:

```text
- Jangan reuse field number.
- Field yang dihapus harus ditandai reserved.
- Tambah field baru boleh jika backward-compatible.
- Breaking change harus menggunakan version baru.
- Proto wajib diperbarui bersama service consumer/provider.
```

Contoh:

```proto
message Student {
  string id = 1;
  string full_name = 2;

  reserved 3;
  reserved "old_field_name";
}
```

---

## 26. Endpoint MVP — Auth

```http
POST /api/v1/auth/login
POST /api/v1/auth/refresh
POST /api/v1/auth/logout
POST /api/v1/auth/forgot-password
POST /api/v1/auth/reset-password
GET  /api/v1/me
GET  /api/v1/me/permissions
GET  /api/v1/me/context
```

Login request:

```json
{
  "email": "user@example.com",
  "password": "user-provided-password"
}
```

Login success data:

```json
{
  "user_id": "uuid",
  "display_name": "User Name",
  "access_token": "jwt",
  "refresh_token": "opaque-token",
  "token_type": "Bearer",
  "expires_in": 900
}
```

Login errors use `VALIDATION_ERROR`, `UNAUTHORIZED`, `FORBIDDEN`, or `SERVICE_UNAVAILABLE` in the standard response envelope. Password hashes are never returned.

Refresh request:

```json
{
  "refresh_token": "opaque-token"
}
```

Refresh success data returns a new `access_token` and `refresh_token`. The submitted token is revoked atomically, its `last_used_at` is updated, and any later reuse returns `UNAUTHORIZED`.

Logout request requires `Authorization: Bearer <access_token>` and the current refresh token:

```json
{
  "refresh_token": "opaque-token"
}
```

Logout verifies that the session belongs to the authenticated actor, revokes it, and returns `{"logged_out": true}` inside the standard response envelope. A refresh attempt using that token returns `UNAUTHORIZED`.

Catatan:

```text
- Login/refresh/forgot-password wajib rate limit.
- Refresh token harus rotating dan disimpan hash.
- /me/context mengembalikan role/scope/context yang bisa dipilih user.
```

---

## 27. Endpoint MVP — Foundation, School, Academic Year

```http
GET /api/v1/foundations/current
GET /api/v1/schools
GET /api/v1/schools/{school_id}
GET /api/v1/academic-years
GET /api/v1/semesters
GET /api/v1/calendar-events
```

Admin Yayasan bisa melihat lintas unit sesuai foundation scope. Role sekolah hanya melihat unit sesuai scope.

---

## 28. Endpoint MVP — Students, Teachers, Classes

```http
GET    /api/v1/students
POST   /api/v1/students
GET    /api/v1/students/{student_id}
PATCH  /api/v1/students/{student_id}
GET    /api/v1/students/{student_id}/guardians
POST   /api/v1/students/{student_id}/guardians

GET    /api/v1/teachers
POST   /api/v1/teachers
GET    /api/v1/teachers/{teacher_id}
PATCH  /api/v1/teachers/{teacher_id}

GET    /api/v1/classes
POST   /api/v1/classes
GET    /api/v1/classes/{class_id}
PATCH  /api/v1/classes/{class_id}
POST   /api/v1/classes/{class_id}/students
POST   /api/v1/classes/{class_id}/homeroom-teacher
```

Aturan:

```text
- Semua endpoint wajib scope by foundation_id dan school_id.
- Guru/orang tua/siswa tidak boleh mengakses data master bebas.
```

---

## 29. Endpoint MVP — Import

```http
GET  /api/v1/imports/templates/{type}
POST /api/v1/imports/{type}/upload
GET  /api/v1/imports/{batch_id}
POST /api/v1/imports/{batch_id}/confirm
GET  /api/v1/imports/{batch_id}/errors
```

`type` MVP:

```text
students
teachers
classes
```

Aturan:

```text
- Upload hanya validasi dan preview.
- Data baru masuk setelah confirm.
- Import file adalah Restricted data.
- Import tidak boleh langsung insert tanpa preview.
```

---

## 30. Endpoint MVP — Admission / PPDB

```http
GET    /api/v1/admissions/periods
POST   /api/v1/admissions/periods
GET    /api/v1/admissions/periods/{period_id}
PATCH  /api/v1/admissions/periods/{period_id}

GET    /api/v1/admissions/applicants
POST   /api/v1/admissions/applicants
GET    /api/v1/admissions/applicants/{applicant_id}
PATCH  /api/v1/admissions/applicants/{applicant_id}
POST   /api/v1/admissions/applicants/{applicant_id}/documents
POST   /api/v1/admissions/applicants/{applicant_id}/verify
POST   /api/v1/admissions/applicants/{applicant_id}/accept
POST   /api/v1/admissions/applicants/{applicant_id}/reject
POST   /api/v1/admissions/applicants/{applicant_id}/convert-to-student
```

Aturan:

```text
- Convert applicant to student wajib idempotency.
- Admission tidak boleh insert langsung ke school_core_db.
- Konversi ke student dilakukan via gRPC ke School Core.
```

---

## 31. Endpoint MVP — Finance / SPP

```http
GET   /api/v1/finance/fee-types
POST  /api/v1/finance/fee-types
PATCH /api/v1/finance/fee-types/{fee_type_id}

GET   /api/v1/finance/fee-schemes
POST  /api/v1/finance/fee-schemes
GET   /api/v1/finance/fee-schemes/{scheme_id}
PATCH /api/v1/finance/fee-schemes/{scheme_id}

GET   /api/v1/finance/fee-policies
POST  /api/v1/finance/fee-policies
GET   /api/v1/finance/fee-policies/{policy_id}
POST  /api/v1/finance/fee-policies/{policy_id}/submit
POST  /api/v1/finance/fee-policies/{policy_id}/disable-request

GET   /api/v1/finance/sibling-discount-rules
POST  /api/v1/finance/sibling-discount-rules

POST  /api/v1/finance/bills/generate
GET   /api/v1/finance/bills
GET   /api/v1/finance/bills/{bill_id}

GET   /api/v1/finance/payments
POST  /api/v1/finance/payments
GET   /api/v1/finance/payments/{payment_id}
POST  /api/v1/finance/payments/{payment_id}/proofs
POST  /api/v1/finance/payments/{payment_id}/verify
POST  /api/v1/finance/payments/{payment_id}/reject
POST  /api/v1/finance/payments/{payment_id}/void-request

GET   /api/v1/finance/receipts/{receipt_id}
POST  /api/v1/finance/receipts/{receipt_id}/signed-url
GET   /api/v1/finance/reports/payments
GET   /api/v1/finance/reports/outstanding
```

Aturan:

```text
- Generate tagihan wajib idempotent.
- Verify payment wajib audit log dan event finance.payment.verified.
- Void payment wajib approval.
- Bill item wajib menyimpan snapshot fee policy.
```

---

## 32. Endpoint MVP — Academic

```http
GET   /api/v1/academic/subjects
POST  /api/v1/academic/subjects
PATCH /api/v1/academic/subjects/{subject_id}

GET   /api/v1/academic/schedules
POST  /api/v1/academic/schedules
PATCH /api/v1/academic/schedules/{schedule_id}

GET   /api/v1/academic/attendances
POST  /api/v1/academic/attendances
PATCH /api/v1/academic/attendances/{attendance_id}

GET   /api/v1/academic/grade-books
GET   /api/v1/academic/grade-books/{grade_book_id}
POST  /api/v1/academic/grade-books/{grade_book_id}/scores
POST  /api/v1/academic/grade-books/{grade_book_id}/submit
POST  /api/v1/academic/grade-books/{grade_book_id}/approve

GET   /api/v1/academic/report-cards
GET   /api/v1/academic/report-cards/{report_card_id}
POST  /api/v1/academic/report-cards/{report_card_id}/publish
POST  /api/v1/academic/report-cards/{report_card_id}/revision-request
POST  /api/v1/academic/report-cards/{report_card_id}/signed-url
```

Aturan:

```text
- Guru hanya bisa akses kelas/mapel assignment-nya.
- Rapor published menjadi locked.
- Revisi setelah publish wajib approval dan audit log.
```

---

## 33. Endpoint MVP — Communication / Notification

```http
GET   /api/v1/announcements
POST  /api/v1/announcements
GET   /api/v1/announcements/{announcement_id}
PATCH /api/v1/announcements/{announcement_id}
POST  /api/v1/announcements/{announcement_id}/publish
POST  /api/v1/announcements/{announcement_id}/archive

GET   /api/v1/notifications
POST  /api/v1/notifications/{notification_id}/read
POST  /api/v1/notifications/read-all
GET   /api/v1/notification-preferences
PATCH /api/v1/notification-preferences
```

Aturan:

```text
- Notification dibuat event-driven.
- Confidential data tidak boleh dikirim detail.
- Critical notification tidak boleh dimatikan sepenuhnya.
```

---

## 34. Endpoint MVP — Reporting Dashboard

```http
GET /api/v1/dashboard/foundation
GET /api/v1/dashboard/school
GET /api/v1/dashboard/teacher
GET /api/v1/dashboard/parent
GET /api/v1/dashboard/student
```

Aturan:

```text
- Dashboard membaca reporting_db melalui Reporting Service.
- Dashboard tidak query database operasional service lain.
- Dashboard scoped sesuai role.
```

---

## 35. Endpoint MVP — Approval

Approval dapat dibuat sebagai unified API di Gateway yang diarahkan ke service owner.

```http
GET  /api/v1/approvals/pending
GET  /api/v1/approvals/requested-by-me
GET  /api/v1/approvals/{approval_id}
POST /api/v1/approvals/{approval_id}/approve
POST /api/v1/approvals/{approval_id}/reject
POST /api/v1/approvals/{approval_id}/request-revision
```

Aturan:

```text
- API Gateway menentukan service owner berdasarkan module/entity_type.
- Service owner tetap melakukan permission/scope check.
- Reason wajib untuk reject/request revision dan aksi sensitif tertentu.
```

---

## 36. Public API Terbatas

Endpoint public harus sangat terbatas.

Contoh yang boleh:

```http
GET  /api/v1/public/schools
GET  /api/v1/public/admissions/periods
POST /api/v1/public/admissions/register
```

Aturan:

```text
- Public API wajib rate limited.
- Public API tidak boleh membocorkan data siswa/guru/orang tua.
- PPDB public register hanya menerima data sesuai form yang diizinkan.
```

Untuk MVP internal awal, public API bisa diminimalkan.

---

## 37. API Security Rules

Semua endpoint protected wajib:

```text
- Auth required
- Permission check
- Scope check
- Object-level authorization
- Input validation
- Rate limit untuk endpoint sensitif
- Audit log untuk aksi sensitif
- No sensitive data in logs
```

Endpoint sensitif yang perlu rate limit:

```text
/auth/login
/auth/refresh
/auth/forgot-password
/files/signed-url
/payment-proof/upload
/public/admissions/register
```

---

## 38. Object-Level Authorization

Setiap endpoint berbasis ID wajib mengecek ownership/scope.

Contoh:

```http
GET /api/v1/students/{student_id}
```

Wajib validasi:

```text
- student_id berada dalam foundation_id user.
- Jika user school-scoped, student_id harus berada di school_id user.
- Jika user orang tua, student_id harus anaknya sendiri.
- Jika user siswa, student_id harus dirinya sendiri.
```

Aturan:

```text
Jangan mengambil data by ID tanpa scope filter.
```

Contoh query aman:

```sql
SELECT *
FROM students
WHERE id = $1
  AND foundation_id = $2
  AND school_id = $3;
```

---

## 39. Backward Compatibility Rules

External API:

```text
- Jangan hapus field response tanpa versioning.
- Tambah field response boleh jika backward-compatible.
- Rename field harus lewat versi baru atau deprecation.
- Breaking change wajib update OpenAPI dan frontend.
```

Internal protobuf:

```text
- Jangan reuse field number.
- Field dihapus harus reserved.
- Tambah field baru harus backward-compatible.
- Breaking change gunakan package/version baru.
```

Event:

```text
- Event payload backward-compatible jika menambah field.
- Breaking change wajib menaikkan event_version.
```

---

## 40. AI Agent Rules untuk API

AI Agent wajib mengikuti aturan berikut saat membuat/mengubah API:

```text
1. Tidak boleh membuat endpoint tanpa permission/scope check.
2. Tidak boleh membuat response format sendiri di luar standar.
3. Tidak boleh membuat error format sendiri di luar standar.
4. Tidak boleh bypass API Gateway dari frontend.
5. Tidak boleh expose REST public langsung dari microservice internal.
6. Tidak boleh membuat query by ID tanpa foundation_id/school_id scope.
7. Tidak boleh menaruh business logic di API Gateway.
8. Tidak boleh menambahkan endpoint tanpa memperbarui OpenAPI.
9. Tidak boleh mengubah gRPC contract tanpa memperbarui proto.
10. Operasi rawan duplikasi wajib mendukung Idempotency-Key.
11. Aksi sensitif wajib audit log.
12. Endpoint async wajib mengembalikan 202 Accepted dengan job_id/import_batch_id.
```

---

## 41. Definition of Done untuk API Task

Sebuah task API dianggap selesai jika:

```text
- Endpoint/gRPC method dibuat sesuai contract.
- Request validation tersedia.
- Permission dan scope check tersedia.
- Object-level authorization tersedia untuk resource by ID.
- Response dan error format sesuai standar.
- Audit log dibuat untuk aksi sensitif.
- Event dipublish jika diwajibkan.
- Idempotency diterapkan jika operasi rawan duplikasi.
- Unit/integration/API test dibuat.
- OpenAPI/proto/event schema diperbarui jika berubah.
- Tidak ada query lintas database service.
- Tidak ada data sensitif masuk log.
```

---

## 42. Ringkasan Keputusan Final

```text
External API menggunakan REST/JSON melalui Custom Go API Gateway.
Internal API menggunakan gRPC/protobuf antar microservice.
Frontend tidak boleh memanggil microservice langsung.
Microservice internal tidak expose REST publik langsung.
```

```text
External API memakai prefix /api/v1 dan didokumentasikan dengan OpenAPI di packages/openapi.
Internal proto contract disimpan di packages/proto.
```

```text
Semua response, error, pagination, filtering, sorting, auth header, context header, dan correlation ID distandarkan.
```

```text
Semua gRPC call membawa request_id, correlation_id, actor_user_id, foundation_id, school_id, role, dan scope melalui metadata/context.
```

```text
Semua endpoint wajib menerapkan authentication, permission, scope check, object-level authorization, input validation, dan audit log untuk aksi sensitif.
```

```text
Operasi async seperti import Excel, generate tagihan massal, dan generate dokumen mengembalikan 202 Accepted dengan job_id/import_batch_id.
```

```text
Idempotency-Key disiapkan untuk generate tagihan, create payment, verify payment, convert applicant to student, dan publish report card.
```

```text
OpenAPI dan protobuf wajib diperbarui jika API berubah.
```
