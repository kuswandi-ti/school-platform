# Product Requirement Document — School Platform

Project: `school-platform`  
Document Type: Product Requirement Document  
Target Audience: Product Owner, Developer, QA, DevOps, AI Agent  
Status: Draft for MVP Implementation  
Language: Bahasa Indonesia  
Repository Target Path: `docs/25-product-requirement-document.md`

---

## 1. Executive Summary

`school-platform` adalah platform internal yayasan sekolah multi-unit yang dirancang untuk mendukung operasional dasar TK, SD, SMP, dan SMA dalam satu ekosistem sistem informasi yang terstruktur, aman, dan dapat dikembangkan secara bertahap.

Platform ini bertujuan untuk menyatukan data inti yayasan dan sekolah, mengelola proses PPDB, mendukung pengelolaan tagihan SPP secara manual, membantu proses akademik dasar seperti jadwal, absensi, nilai, dan rapor, serta menyediakan komunikasi dan dashboard pelaporan untuk kebutuhan yayasan, sekolah, guru, orang tua, dan siswa.

Pada tahap MVP, `school-platform` difokuskan sebagai sistem internal yayasan. Namun sejak awal, arsitektur produk dirancang SaaS-ready dengan pendekatan `foundation_id`, `school_id`, role-based access control, scope-based access control, audit log, service boundary yang jelas, private file storage, dan reporting berbasis read model.

MVP tidak bertujuan untuk langsung menjadi sistem sekolah lengkap dengan payroll, LMS penuh, library, asset management, payment gateway, WhatsApp integration, atau offline mobile write. Fokus utama MVP adalah menyediakan fondasi operasional yang stabil, aman, dan dapat dipilotkan pada unit sekolah dalam yayasan.

Keberhasilan MVP diukur dari kemampuan yayasan menjalankan alur utama berikut:

```text
login
→ setup data yayasan/sekolah
→ kelola siswa/guru/kelas
→ import data awal
→ jalankan PPDB
→ generate tagihan SPP
→ verifikasi pembayaran manual
→ input absensi dan nilai dasar
→ publish rapor dasar
→ kirim pengumuman/notifikasi
→ melihat dashboard ringkasan
→ audit aksi sensitif
→ backup dan restore teruji
```

---

## 2. Background and Problem Statement

Yayasan yang mengelola beberapa unit pendidikan seperti TK, SD, SMP, dan SMA umumnya menghadapi tantangan operasional yang berulang. Setiap unit sekolah memiliki kebutuhan administrasi yang mirip, tetapi sering kali data dan prosesnya berjalan terpisah. Kondisi ini menimbulkan duplikasi pekerjaan, kesulitan konsolidasi data, dan risiko kesalahan manual.

Masalah utama yang ingin diselesaikan oleh `school-platform` adalah:

### 2.1 Data Siswa, Guru, dan Kelas Tersebar

Data siswa, orang tua, guru, kelas, dan wali kelas sering tersebar di file spreadsheet, aplikasi terpisah, atau catatan manual. Akibatnya:

- data sulit diverifikasi;
- data antar unit sekolah tidak konsisten;
- perubahan data tidak selalu tercatat;
- pencarian data membutuhkan waktu;
- yayasan sulit mendapatkan gambaran jumlah siswa/guru secara akurat.

### 2.2 PPDB Belum Terpusat

Proses Penerimaan Peserta Didik Baru sering berjalan menggunakan formulir terpisah, file manual, atau komunikasi informal. Dampaknya:

- status pendaftar sulit dipantau;
- dokumen pendaftar sulit dilacak;
- keputusan diterima/ditolak tidak terdokumentasi rapi;
- konversi pendaftar menjadi siswa aktif berisiko duplikasi data;
- yayasan tidak memiliki ringkasan PPDB lintas unit secara cepat.

### 2.3 SPP Manual Sulit Direkonsiliasi

Pada MVP, pembayaran SPP tetap dilakukan secara manual menggunakan upload bukti pembayaran. Tantangan yang muncul:

- tagihan tidak selalu tersusun konsisten;
- diskon, beasiswa, bebas SPP, dan sibling discount sulit dilacak;
- bukti pembayaran tersebar;
- verifikasi pembayaran tidak selalu terdokumentasi;
- tunggakan sulit dipantau per siswa/per sekolah;
- perubahan atau void pembayaran membutuhkan audit.

### 2.4 Laporan Akademik Belum Terintegrasi

Proses akademik seperti jadwal, absensi, nilai, dan rapor sering berjalan terpisah. Dampaknya:

- guru sulit melihat jadwal dan assignment secara konsisten;
- absensi tidak selalu terkonsolidasi;
- nilai dan rapor membutuhkan rekap manual;
- revisi rapor setelah publish sulit dikontrol;
- orang tua tidak selalu mendapatkan informasi akademik tepat waktu.

### 2.5 Komunikasi Sekolah-Orang Tua Belum Konsisten

Pengumuman sekolah atau yayasan dapat tersebar di berbagai channel informal. Dampaknya:

- orang tua dapat melewatkan informasi penting;
- tidak ada log delivery/read status;
- pengumuman penting sulit ditargetkan ke sekolah, kelas, atau role tertentu;
- informasi sensitif berisiko tersebar tidak tepat sasaran.

### 2.6 Dashboard Yayasan Belum Real-Time

Yayasan membutuhkan ringkasan operasional lintas unit, misalnya jumlah siswa aktif, progres PPDB, pembayaran SPP, absensi, progres input nilai, dan pending approvals. Tanpa dashboard terpusat:

- keputusan yayasan lambat;
- data harus dikumpulkan manual;
- risiko salah interpretasi meningkat;
- unit sekolah sulit dibandingkan secara konsisten.

### 2.7 Audit dan Akses Data Belum Tertata

Sistem sekolah mengelola data anak, orang tua, keuangan, dokumen, dan rapor. Tanpa kontrol akses dan audit yang baik:

- perubahan data sensitif sulit dilacak;
- akses data lintas sekolah dapat terjadi tanpa sengaja;
- download/export dokumen tidak tercatat;
- tindakan penting seperti verifikasi pembayaran atau publish rapor tidak memiliki jejak audit;
- risiko privasi dan kepatuhan meningkat.

---

## 3. Product Vision

### 3.1 MVP Vision

MVP `school-platform` menjadi fondasi sistem informasi internal yayasan yang mampu menjalankan proses operasional inti secara terintegrasi, aman, dan dapat dipilotkan pada unit TK, SD, SMP, dan SMA.

MVP harus mampu mendukung:

- data master yayasan, sekolah, siswa, orang tua, guru, dan kelas;
- PPDB hingga konversi pendaftar menjadi siswa;
- pengelolaan tagihan SPP manual;
- upload dan verifikasi bukti pembayaran;
- jadwal, absensi, dan nilai dasar;
- rapor dasar dengan workflow approval/publish;
- pengumuman dan notifikasi dasar;
- dashboard ringkasan operasional;
- audit aksi sensitif;
- private file management;
- backup dan restore minimum.

### 3.2 Post-MVP Vision

Setelah MVP stabil, platform dapat dikembangkan untuk mendukung modul lanjutan seperti:

- payment gateway;
- WhatsApp notification;
- BK/UKS detail;
- library;
- asset/inventory;
- LMS;
- alumni/tracer;
- HR dan payroll;
- koperasi;
- global search;
- advanced reporting/BI;
- offline mobile capability;
- integrations with external education systems.

### 3.3 SaaS-Ready Direction

Walaupun tahap awal adalah sistem internal yayasan, arsitektur produk harus tetap SaaS-ready. Artinya:

- data utama memiliki `foundation_id`;
- data terkait sekolah memiliki `school_id`;
- akses user selalu berdasarkan role dan scope;
- service boundary jelas;
- tidak ada query langsung ke database service lain;
- file private by default;
- audit log tersedia untuk aksi sensitif;
- reporting menggunakan projection/read model;
- environment dan deployment dapat dipisahkan;
- konfigurasi produk dapat berkembang untuk multi-yayasan di masa depan.

---

## 4. Product Goals

Goal utama produk adalah:

1. Menyatukan data operasional dasar yayasan dan sekolah dalam satu platform.
2. Mengurangi ketergantungan terhadap spreadsheet manual untuk data inti.
3. Mendukung proses PPDB dari pendaftaran hingga konversi siswa.
4. Mendukung pengelolaan tagihan SPP manual dengan snapshot dan audit.
5. Mendukung proses verifikasi pembayaran manual berbasis upload bukti.
6. Mendukung proses akademik dasar: jadwal, absensi, nilai, dan rapor.
7. Mendukung workflow publish rapor yang terkendali dan dapat diaudit.
8. Meningkatkan komunikasi sekolah-orang tua melalui announcement dan notification.
9. Menyediakan dashboard ringkasan untuk yayasan, sekolah, guru, orang tua, dan siswa.
10. Menjaga keamanan data anak, orang tua, guru, keuangan, dan dokumen.
11. Menyediakan audit trail untuk aksi penting dan sensitif.
12. Menyediakan fondasi teknis yang scalable, maintainable, dan SaaS-ready.
13. Menyediakan struktur dokumentasi dan workflow yang dapat digunakan oleh tim manusia maupun AI Agent.

---

## 5. Non-Goals

Fitur berikut tidak termasuk MVP:

| Non-Goal | Alasan Tidak Masuk MVP |
|---|---|
| Payroll | Kompleksitas perhitungan gaji, pajak, tunjangan, dan HR membutuhkan scope tersendiri. |
| HR lengkap | MVP hanya membutuhkan data guru/staff dasar, bukan lifecycle HR penuh. |
| Asset/Inventory lengkap | Tidak terkait langsung dengan core flow MVP sekolah. |
| Library | Perlu domain terpisah seperti katalog, peminjaman, denda, dan inventory buku. |
| BK/UKS detail | Mengandung data sangat sensitif dan membutuhkan kontrol privasi lanjutan. |
| LMS penuh | Membutuhkan materi, assignment, kuis, diskusi, dan grading online kompleks. |
| Alumni/Tracer | Tidak dibutuhkan untuk pilot operasional aktif. |
| Koperasi | Domain transaksi dan inventory tersendiri. |
| Global Search | MVP cukup dengan local search/filter per modul. |
| Payment Gateway | MVP menggunakan manual payment + upload bukti pembayaran. |
| WhatsApp | MVP menggunakan in-app, FCM push, dan email terbatas. |
| Offline Write Mobile | MVP online-only untuk aksi utama. |
| Kubernetes | MVP menggunakan Docker Compose untuk local/staging awal. |
| Full BI/Data Warehouse | Reporting MVP menggunakan read model/projection sederhana. |
| Advanced Admission Scoring | PPDB MVP hanya sampai proses verifikasi dan keputusan dasar. |
| National E-Rapor Integration | MVP menyediakan e-rapor basic internal, bukan integrasi nasional. |

---

## 6. Target Users and Roles

| Role | Tujuan Penggunaan | Kebutuhan Utama | Contoh Aktivitas | Batasan Akses |
|---|---|---|---|---|
| Admin Yayasan | Mengelola konfigurasi lintas sekolah dan melihat ringkasan yayasan | Setup yayasan, sekolah, role, approval, dashboard lintas unit | Membuat data sekolah, melihat dashboard yayasan, approve aksi lintas sekolah | Scope yayasan; akses data sekolah sesuai izin |
| Kepala Sekolah | Mengawasi operasional unit sekolah | Dashboard sekolah, approval, publish rapor, validasi keputusan penting | Approve rapor, melihat data siswa, memantau PPDB/SPP/akademik | Hanya sekolah yang dipimpin |
| TU/Staff | Mengelola administrasi sekolah | Data siswa/guru/kelas, import data, PPDB operasional | Input data siswa, upload import Excel, verifikasi dokumen PPDB | Hanya sekolah terkait; tidak boleh akses finance sensitif tanpa izin |
| Bendahara Sekolah | Mengelola tagihan dan pembayaran | Fee scheme, generate bill, verifikasi pembayaran, laporan tunggakan | Generate SPP, verifikasi bukti pembayaran, melihat outstanding | Hanya data finance sekolah terkait |
| Guru | Mengelola aktivitas akademik yang ditugaskan | Jadwal, absensi, nilai per kelas/mapel | Melihat jadwal, input absensi, input nilai | Hanya kelas/mapel yang ditugaskan |
| Wali Kelas | Mengawasi kelas tertentu | Review absensi/nilai/rapor kelas | Review data rapor, mengajukan koreksi | Hanya kelas yang menjadi tanggung jawabnya |
| Orang Tua/Wali Murid | Melihat informasi anak dan melakukan pembayaran manual | Tagihan, upload bukti, rapor, pengumuman | Melihat tagihan anak, upload bukti, melihat rapor published | Hanya anak yang terhubung |
| Siswa | Melihat informasi akademik pribadi | Jadwal, nilai/rapor published, pengumuman | Melihat jadwal, melihat rapor yang sudah publish | Hanya data pribadi |

---

## 7. User Personas

### 7.1 Admin Yayasan

| Aspek | Detail |
|---|---|
| Profile | Pengelola yayasan yang mengawasi beberapa unit sekolah. |
| Goals | Mendapatkan kontrol dan visibilitas lintas unit. |
| Pain Points | Data tersebar, laporan lambat, sulit membandingkan kondisi antar sekolah. |
| Main Tasks | Setup sekolah, role, melihat dashboard, memantau pending approval. |
| Success Expectation | Bisa melihat ringkasan operasional yayasan secara cepat dan aman. |

### 7.2 Kepala Sekolah

| Aspek | Detail |
|---|---|
| Profile | Pemimpin unit sekolah. |
| Goals | Memastikan operasional sekolah berjalan tertib. |
| Pain Points | Sulit memantau SPP, absensi, nilai, PPDB, dan rapor dalam satu tempat. |
| Main Tasks | Review data, approve/publish rapor, memantau dashboard sekolah. |
| Success Expectation | Semua proses penting sekolah bisa dipantau dan disetujui dengan jejak audit. |

### 7.3 TU/Staff

| Aspek | Detail |
|---|---|
| Profile | Staf administrasi yang mengelola data operasional harian. |
| Goals | Menginput dan mengelola data dengan cepat dan minim duplikasi. |
| Pain Points | Data Excel tersebar, validasi manual, koreksi data sulit dilacak. |
| Main Tasks | Import data, kelola siswa/guru/kelas, proses PPDB. |
| Success Expectation | Data master rapi, valid, mudah dicari, dan sesuai scope sekolah. |

### 7.4 Bendahara

| Aspek | Detail |
|---|---|
| Profile | Pengelola keuangan sekolah. |
| Goals | Mengelola tagihan SPP dan pembayaran manual secara akurat. |
| Pain Points | Tagihan dan bukti pembayaran sulit direkonsiliasi. |
| Main Tasks | Setup fee, generate bill, verifikasi pembayaran, pantau tunggakan. |
| Success Expectation | Status pembayaran siswa jelas, bukti tersimpan, dan transaksi sensitif diaudit. |

### 7.5 Guru/Wali Kelas

| Aspek | Detail |
|---|---|
| Profile | Pengajar dan/atau wali kelas. |
| Goals | Mengelola absensi, nilai, dan rapor sesuai assignment. |
| Pain Points | Jadwal dan data siswa tidak selalu sinkron; rekap nilai manual. |
| Main Tasks | Input absensi, input nilai, review rapor kelas. |
| Success Expectation | Pekerjaan akademik lebih terstruktur dan sesuai akses. |

### 7.6 Orang Tua/Wali Murid

| Aspek | Detail |
|---|---|
| Profile | Wali murid yang memantau informasi anak. |
| Goals | Mendapat akses cepat ke tagihan, pembayaran, pengumuman, dan rapor. |
| Pain Points | Informasi tersebar di banyak channel dan tidak terdokumentasi. |
| Main Tasks | Lihat tagihan, upload bukti, baca pengumuman, lihat rapor. |
| Success Expectation | Dapat memantau informasi anak dengan mudah dan aman. |

### 7.7 Siswa

| Aspek | Detail |
|---|---|
| Profile | Peserta didik yang membutuhkan akses informasi akademik pribadi. |
| Goals | Melihat jadwal, nilai, dan rapor yang sudah dipublish. |
| Pain Points | Informasi akademik sulit diakses secara mandiri. |
| Main Tasks | Lihat jadwal, lihat nilai/rapor published, baca pengumuman. |
| Success Expectation | Informasi pribadi tersedia dengan jelas dan sesuai izin. |

---

## 8. MVP Scope

### 8.1 Identity & Access

| Item | Detail |
|---|---|
| Objective | Menyediakan autentikasi, otorisasi, role, permission, dan user context. |
| Fitur Utama | Login, JWT access token, rotating refresh token, logout, role/permission, user scope. |
| User Role | Semua user. |
| Key Workflow | Login → token issued → context loaded → role/scope enforced. |
| Data Utama | users, roles, permissions, role assignments, sessions, refresh tokens. |
| Acceptance Criteria Ringkas | User dapat login sesuai role; refresh token rotate; akses dibatasi role/scope; token tidak tersimpan/log mentah. |

### 8.2 School Core

| Item | Detail |
|---|---|
| Objective | Mengelola data master yayasan, sekolah, tahun ajaran, semester, siswa, orang tua, guru, kelas, assignment. |
| Fitur Utama | CRUD data master, student/guardian/teacher/class, assignment siswa-kelas, wali kelas, guru mapel. |
| User Role | Admin Yayasan, Kepala Sekolah, TU/Staff. |
| Key Workflow | Setup yayasan/sekolah → setup tahun ajaran/semester → import/input siswa/guru/kelas. |
| Data Utama | foundations, schools, academic years, semesters, students, guardians, teachers, classes, assignments. |
| Acceptance Criteria Ringkas | Data tersimpan sesuai foundation/school scope; pencarian/filter berjalan; perubahan sensitif diaudit. |

### 8.3 File Management + Import Excel

| Item | Detail |
|---|---|
| Objective | Menyediakan private file management dan import Excel data awal. |
| Fitur Utama | Upload file, metadata, signed URL, import template, validation preview, confirm import, import report. |
| User Role | Admin Yayasan, TU/Staff, Bendahara, Guru sesuai modul. |
| Key Workflow | Download template → upload Excel → validate → preview → confirm → import report. |
| Data Utama | file metadata, import_batches, import_batch_rows. |
| Acceptance Criteria Ringkas | File private; validasi sebelum import; error report tersedia; raw data tidak masuk log. |

### 8.4 PPDB

| Item | Detail |
|---|---|
| Objective | Mengelola proses penerimaan peserta didik baru. |
| Fitur Utama | Admission period, applicant, guardian, document upload, verification, accept/reject, conversion to student. |
| User Role | Admin Yayasan, Kepala Sekolah, TU/Staff, Orang Tua calon siswa. |
| Key Workflow | Create period → applicant submit → upload docs → verification → decision → convert to student. |
| Data Utama | admission_periods, applicants, applicant_guardians, applicant_documents, admission_decisions. |
| Acceptance Criteria Ringkas | Applicant dapat diproses; dokumen private; keputusan diaudit; konversi idempotent melalui School Core. |

### 8.5 Finance/SPP

| Item | Detail |
|---|---|
| Objective | Mengelola tagihan SPP manual, bukti pembayaran, dan verifikasi pembayaran. |
| Fitur Utama | Fee type, scheme, policy, sibling discount, bill generation, payment proof, verify/reject, receipt, outstanding. |
| User Role | Bendahara, Kepala Sekolah, Admin Yayasan, Orang Tua. |
| Key Workflow | Setup fee → assign policy → generate bill → parent upload proof → treasurer verify → receipt. |
| Data Utama | fee_types, fee_schemes, student_fee_policies, bills, bill_items, payments, payment_proofs. |
| Acceptance Criteria Ringkas | Amount menggunakan decimal; bill snapshot; generate idempotent; verification diaudit; parent hanya akses bill anak. |

### 8.6 Academic Basic

| Item | Detail |
|---|---|
| Objective | Mendukung proses akademik dasar. |
| Fitur Utama | Curriculum, subject, class subject, schedule, teacher schedule, attendance, attendance correction. |
| User Role | Admin Yayasan, Kepala Sekolah, TU/Staff, Guru, Wali Kelas. |
| Key Workflow | Setup subject/schedule → guru lihat jadwal → input absensi → correction jika perlu. |
| Data Utama | curriculums, subjects, class_subjects, schedules, student_attendances. |
| Acceptance Criteria Ringkas | Guru hanya akses assignment; absensi tersimpan; correction butuh reason dan audit. |

### 8.7 Report Card / E-Rapor Basic

| Item | Detail |
|---|---|
| Objective | Mendukung input nilai dan publish rapor dasar. |
| Fitur Utama | Assessment scheme, grade book, score input, review, report card generation, publish/lock, revision approval. |
| User Role | Guru, Wali Kelas, Kepala Sekolah, Orang Tua, Siswa. |
| Key Workflow | Guru input nilai → submit → wali kelas review → kepala sekolah publish → parent/student view. |
| Data Utama | assessment_components, grade_books, student_scores, report_cards, report_card_items. |
| Acceptance Criteria Ringkas | Rapor published terkunci; revisi setelah publish butuh approval; parent/student hanya lihat published report. |

### 8.8 Communication / Notification

| Item | Detail |
|---|---|
| Objective | Mengirim pengumuman dan notifikasi internal. |
| Fitur Utama | Announcement, target audience, in-app notification, FCM/email provider abstraction, delivery log, preferences. |
| User Role | Admin Yayasan, Kepala Sekolah, TU/Staff, Guru, Orang Tua, Siswa. |
| Key Workflow | Create announcement → target audience → publish → notification delivered → read/unread. |
| Data Utama | announcements, announcement_targets, notifications, templates, deliveries, preferences. |
| Acceptance Criteria Ringkas | Notifikasi event-driven; confidential detail tidak masuk body; critical notification tidak bisa dimatikan penuh. |

### 8.9 Reporting Dashboard

| Item | Detail |
|---|---|
| Objective | Menyediakan dashboard berbasis read model/projection. |
| Fitur Utama | Dashboard yayasan, sekolah, guru, parent/student; projection consumer; summary metrics. |
| User Role | Admin Yayasan, Kepala Sekolah, Guru, Orang Tua, Siswa. |
| Key Workflow | Operational services publish events → Reporting consumes → dashboard updates. |
| Data Utama | reporting projections, processed_events, dashboard summaries. |
| Acceptance Criteria Ringkas | Reporting hanya baca reporting_db; event idempotent; dashboard sesuai role/scope. |

### 8.10 Security, Observability, Backup, and UAT Hardening

| Item | Detail |
|---|---|
| Objective | Memastikan MVP siap pilot/production. |
| Fitur Utama | Security review, permission regression, logging, metrics, backup, restore test, UAT checklist. |
| User Role | DevOps, QA, Product Owner, Technical Lead. |
| Key Workflow | Hardening → regression → backup/restore test → UAT → release readiness. |
| Data Utama | audit logs, metrics, logs, backup artifacts. |
| Acceptance Criteria Ringkas | No Critical/High core bug; restore test dilakukan; logs aman; release checklist pass. |

---

## 9. User Journey

### 9.1 Admin Yayasan Setup Sekolah

| Item | Detail |
|---|---|
| Actor | Admin Yayasan |
| Trigger | Yayasan mulai menggunakan platform atau menambah unit sekolah. |
| Steps | Login → create/update foundation → create school TK/SD/SMP/SMA → setup academic year/semester → assign roles. |
| Output | Struktur yayasan dan sekolah siap digunakan. |
| Exception/Error Cases | School duplicate, missing required data, insufficient permission. |
| Audit/Approval | Role assignment dan perubahan data penting diaudit. |

### 9.2 TU Import Data Awal

| Item | Detail |
|---|---|
| Actor | TU/Staff |
| Trigger | Data awal siswa/guru/kelas perlu dimigrasikan dari Excel. |
| Steps | Download template → isi data → upload → validate → preview → confirm → import report. |
| Output | Data siswa/guru/kelas masuk ke School Core. |
| Exception/Error Cases | Format salah, duplicate student, invalid class_code, missing guardian. |
| Audit/Approval | Import batch dicatat; file Restricted; action diaudit. |

### 9.3 PPDB sampai Konversi Siswa

| Item | Detail |
|---|---|
| Actor | Orang Tua calon siswa, TU/Staff, Kepala Sekolah |
| Trigger | Admission period dibuka. |
| Steps | Applicant submit → upload documents → staff verify → decision accept/reject → accepted applicant converted to student. |
| Output | Applicant diterima menjadi student di School Core. |
| Exception/Error Cases | Dokumen tidak lengkap, duplicate applicant, conversion failed, already converted. |
| Audit/Approval | Decision dan conversion diaudit. |

### 9.4 Bendahara Generate Tagihan dan Verifikasi Pembayaran

| Item | Detail |
|---|---|
| Actor | Bendahara, Orang Tua |
| Trigger | Awal periode tagihan SPP. |
| Steps | Setup fee scheme → apply policy → generate bill → parent upload proof → treasurer verify/reject → receipt generated. |
| Output | Tagihan dan status pembayaran tercatat. |
| Exception/Error Cases | Duplicate bill, invalid amount, proof unreadable, parent not linked to student. |
| Audit/Approval | Verification/rejection/void diaudit; void memerlukan approval jika relevan. |

### 9.5 Guru Input Absensi

| Item | Detail |
|---|---|
| Actor | Guru |
| Trigger | Sesi kelas berlangsung. |
| Steps | Login → pilih jadwal/kelas → input attendance → submit. |
| Output | Absensi siswa tersimpan. |
| Exception/Error Cases | Guru tidak assigned, kelas tidak aktif, duplicate attendance. |
| Audit/Approval | Correction setelah submit membutuhkan reason dan audit. |

### 9.6 Guru Input Nilai

| Item | Detail |
|---|---|
| Actor | Guru |
| Trigger | Penilaian selesai dilakukan. |
| Steps | Login → pilih subject/class → input score → save draft → submit grade book. |
| Output | Nilai siswa tersimpan dan siap review. |
| Exception/Error Cases | Guru tidak assigned, invalid score, period locked. |
| Audit/Approval | Submit/revision dicatat; perubahan setelah lock perlu aturan khusus. |

### 9.7 Wali Kelas Review Data Rapor

| Item | Detail |
|---|---|
| Actor | Wali Kelas |
| Trigger | Nilai mapel sudah disubmit. |
| Steps | Buka kelas → review nilai/absensi/deskripsi → request revision jika perlu → submit for approval. |
| Output | Rapor kelas siap disetujui Kepala Sekolah. |
| Exception/Error Cases | Nilai belum lengkap, data siswa tidak valid, deskripsi kosong. |
| Audit/Approval | Review dan revision request diaudit. |

### 9.8 Kepala Sekolah Publish Rapor

| Item | Detail |
|---|---|
| Actor | Kepala Sekolah |
| Trigger | Rapor sudah siap publish. |
| Steps | Review final → approve/publish → report card locked → parent/student notified. |
| Output | Rapor published dan terkunci. |
| Exception/Error Cases | Nilai belum lengkap, approval missing, publish duplicate. |
| Audit/Approval | Publish dan revision after publish wajib audit. |

### 9.9 Orang Tua Melihat Tagihan dan Upload Bukti Pembayaran

| Item | Detail |
|---|---|
| Actor | Orang Tua/Wali Murid |
| Trigger | Ada tagihan aktif. |
| Steps | Login mobile/web → pilih anak → lihat tagihan → upload bukti → tunggu verifikasi. |
| Output | Bukti pembayaran tersimpan dan masuk antrean verifikasi. |
| Exception/Error Cases | File invalid, parent-child scope invalid, bill sudah paid/void. |
| Audit/Approval | Upload proof dicatat; file Restricted. |

### 9.10 Orang Tua Melihat Rapor Anak

| Item | Detail |
|---|---|
| Actor | Orang Tua/Wali Murid |
| Trigger | Rapor sudah published. |
| Steps | Login → pilih anak → buka rapor published → lihat detail. |
| Output | Orang tua dapat melihat rapor anak. |
| Exception/Error Cases | Rapor belum published, parent tidak linked, access denied. |
| Audit/Approval | Download/export jika ada harus diaudit. |

### 9.11 Siswa Melihat Jadwal/Nilai/Rapor Published

| Item | Detail |
|---|---|
| Actor | Siswa |
| Trigger | Siswa ingin melihat informasi akademik pribadi. |
| Steps | Login → buka jadwal/nilai/rapor → lihat data yang sudah diizinkan. |
| Output | Siswa melihat data akademik pribadi. |
| Exception/Error Cases | Data belum published, scope invalid, akun belum aktif. |
| Audit/Approval | Akses data pribadi mengikuti scope student. |

---

## 10. Functional Requirements

| ID | Module | Requirement | Role | Priority | Acceptance Criteria |
|---|---|---|---|---|---|
| PRD-ID-001 | Identity & Access | Sistem harus menyediakan login berbasis email/password. | Semua user | Critical | User valid dapat login; user invalid ditolak; error tidak membocorkan credential detail. |
| PRD-ID-002 | Identity & Access | Sistem harus menggunakan JWT access token dan rotating refresh token. | Semua user | Critical | Refresh token di-rotate; token lama tidak dapat digunakan ulang; token tidak disimpan mentah. |
| PRD-ID-003 | Identity & Access | Sistem harus menyediakan role dan permission. | Admin Yayasan | Critical | Role/permission dapat digunakan untuk membatasi akses endpoint dan UI. |
| PRD-ID-004 | Identity & Access | Sistem harus membangun actor context. | Semua user | Critical | Context memuat user_id, foundation_id, school_id, roles, permissions, scope. |
| PRD-SC-001 | School Core | Sistem harus mengelola data foundation dan school. | Admin Yayasan | Critical | Foundation/school dapat dibuat, diubah, dan difilter sesuai scope. |
| PRD-SC-002 | School Core | Sistem harus mengelola tahun ajaran dan semester. | Admin Yayasan, Kepala Sekolah | High | Academic year/semester aktif dapat dipilih sebagai global context. |
| PRD-SC-003 | School Core | Sistem harus mengelola data siswa. | TU/Staff | Critical | Student CRUD berjalan sesuai school scope; perubahan sensitif diaudit. |
| PRD-SC-004 | School Core | Sistem harus mengelola guardian dan relasi student_guardian. | TU/Staff | High | Orang tua hanya dapat dikaitkan dengan siswa yang valid dalam scope sekolah. |
| PRD-SC-005 | School Core | Sistem harus mengelola data guru dan assignment. | TU/Staff, Kepala Sekolah | High | Guru dapat diassign ke class/subject/homeroom sesuai kebutuhan. |
| PRD-FILE-001 | File Management | Sistem harus menyimpan file private by default. | Semua role terkait | Critical | File tidak dapat diakses tanpa authorization backend. |
| PRD-FILE-002 | File Management | Sistem harus menyediakan signed URL untuk akses file private. | Semua role terkait | High | URL memiliki expiry sesuai klasifikasi data. |
| PRD-FILE-003 | Import Excel | Sistem harus mendukung import data siswa, guardian, guru, kelas, dan assignment. | TU/Staff | High | Import melalui validate-preview-confirm-report. |
| PRD-FILE-004 | Import Excel | Sistem harus menyediakan error report. | TU/Staff | Medium | Baris gagal import dapat ditinjau dan dikoreksi. |
| PRD-ADM-001 | PPDB | Sistem harus menyediakan admission period. | Admin Yayasan, TU/Staff | High | Period dapat dibuat dengan status dan tanggal valid. |
| PRD-ADM-002 | PPDB | Sistem harus menyimpan applicant dan guardian calon siswa. | Orang Tua, TU/Staff | High | Applicant tersimpan dengan school target dan status. |
| PRD-ADM-003 | PPDB | Sistem harus mendukung upload dokumen pendaftar. | Orang Tua, TU/Staff | High | Dokumen private dan dapat diverifikasi. |
| PRD-ADM-004 | PPDB | Sistem harus mendukung keputusan accept/reject. | Kepala Sekolah, Admin Yayasan | High | Decision tercatat dan diaudit. |
| PRD-ADM-005 | PPDB | Sistem harus mengonversi accepted applicant menjadi student. | TU/Staff | Critical | Conversion idempotent dan dilakukan melalui School Core. |
| PRD-FIN-001 | Finance | Sistem harus mengelola fee type dan fee scheme. | Bendahara | Critical | Fee dapat dibuat dan digunakan untuk bill generation. |
| PRD-FIN-002 | Finance | Sistem harus mengelola fee policy seperti free_spp, discount, scholarship, sibling_discount. | Bendahara, Kepala Sekolah | Critical | Policy tersimpan dengan period, reason, status approval. |
| PRD-FIN-003 | Finance | Sistem harus generate bill secara idempotent. | Bendahara | Critical | Tidak ada duplicate bill untuk student/period/fee yang sama. |
| PRD-FIN-004 | Finance | Bill harus menyimpan snapshot amount dan policy. | Bendahara | Critical | Perubahan policy berikutnya tidak mengubah bill lama. |
| PRD-FIN-005 | Finance | Orang tua harus dapat upload bukti pembayaran. | Orang Tua | High | Proof tersimpan sebagai Restricted file dan masuk antrean verifikasi. |
| PRD-FIN-006 | Finance | Bendahara harus dapat verify/reject payment. | Bendahara | Critical | Status payment berubah; receipt dibuat jika verified; action diaudit. |
| PRD-FIN-007 | Finance | Sistem harus mendukung void payment dengan approval jika diperlukan. | Bendahara, Kepala Sekolah | High | Void tidak dapat dilakukan tanpa reason/approval/audit. |
| PRD-ACD-001 | Academic Basic | Sistem harus mengelola curriculum dan subject. | Admin Yayasan, Kepala Sekolah, TU/Staff | High | Subject dapat diassign ke class/teacher. |
| PRD-ACD-002 | Academic Basic | Sistem harus mengelola schedule. | TU/Staff, Guru | High | Jadwal dapat dilihat oleh guru sesuai assignment. |
| PRD-ACD-003 | Academic Basic | Guru harus dapat input absensi. | Guru | Critical | Guru hanya dapat input untuk class/subject assignment. |
| PRD-ACD-004 | Academic Basic | Correction absensi harus membutuhkan reason dan audit. | Guru, Wali Kelas | High | Correction tercatat dengan actor dan alasan. |
| PRD-RC-001 | Report Card | Sistem harus mendukung assessment scheme dan grade book. | Guru, Kepala Sekolah | Critical | Guru dapat input nilai sesuai assignment. |
| PRD-RC-002 | Report Card | Wali Kelas harus dapat review rapor kelas. | Wali Kelas | High | Review hanya untuk kelas yang diampu. |
| PRD-RC-003 | Report Card | Kepala Sekolah harus dapat publish rapor. | Kepala Sekolah | Critical | Published report card terkunci dan diaudit. |
| PRD-RC-004 | Report Card | Revisi setelah publish harus approval dan audit. | Guru, Wali Kelas, Kepala Sekolah | Critical | Rapor published tidak dapat diubah langsung. |
| PRD-RC-005 | Report Card | Orang tua dan siswa hanya dapat melihat rapor published. | Orang Tua, Siswa | High | Draft/submitted report tidak dapat diakses parent/student. |
| PRD-COM-001 | Communication | Sistem harus mendukung announcement. | Admin Yayasan, Kepala Sekolah, TU/Staff | High | Announcement dapat ditargetkan ke role/school/class. |
| PRD-COM-002 | Notification | Sistem harus mendukung in-app notification. | Semua user | High | Notification dapat dibaca dan ditandai read/unread. |
| PRD-COM-003 | Notification | Sistem harus menggunakan event-driven notification. | System | High | Business service tidak memanggil provider langsung. |
| PRD-COM-004 | Notification | Confidential detail tidak boleh masuk notification body. | System | Critical | Payload notification aman dari data sensitif mentah. |
| PRD-RPT-001 | Reporting | Reporting harus menggunakan reporting_db. | Admin Yayasan, Kepala Sekolah | Critical | Dashboard tidak query DB operasional langsung. |
| PRD-RPT-002 | Reporting | Reporting consumer harus idempotent. | System | High | Duplicate event tidak menggandakan metrics. |
| PRD-RPT-003 | Reporting | Dashboard harus dibatasi role/scope. | Semua role terkait | Critical | User hanya melihat data sesuai foundation/school/class/child scope. |
| PRD-SEC-001 | Security | Semua resource by ID harus memiliki object-level authorization. | System | Critical | Akses lintas scope ditolak. |
| PRD-SEC-002 | Audit | Aksi sensitif harus memiliki audit log. | System | Critical | Audit menyimpan actor, action, resource, request_id/correlation_id. |
| PRD-SEC-003 | Backup | Sistem harus memiliki backup dan restore test. | DevOps | Critical | Restore test terdokumentasi sebelum MVP release. |
| PRD-SEC-004 | Observability | Sistem harus memiliki structured logs dan metrics dasar. | DevOps | High | Logs memiliki request_id/correlation_id dan tidak memuat data sensitif. |

---

## 11. Non-Functional Requirements

| Category | Requirement | Priority | Acceptance Criteria |
|---|---|---|---|
| Security | Backend wajib enforce authentication dan authorization. | Critical | Endpoint protected menolak request tanpa token valid. |
| Security | Object-level authorization wajib untuk resource by ID. | Critical | User tidak dapat akses data luar scope. |
| Privacy | Restricted/Confidential data tidak boleh masuk log mentah. | Critical | Log review tidak menemukan token, password, data confidential mentah. |
| Auditability | Aksi sensitif wajib tercatat. | Critical | Audit log menyimpan actor, action, resource, timestamp, request_id. |
| Performance | Dashboard MVP dapat menampilkan data ringkasan secara responsif. | Medium | Dashboard memakai reporting projection, bukan heavy cross-service query. |
| Availability | MVP dapat berjalan stabil di staging/production awal. | High | Health/readiness endpoint tersedia. |
| Backup/Restore | Backup harian dan restore test tersedia. | Critical | Restore test terdokumentasi dan berhasil. |
| Observability | Structured JSON log, request_id, correlation_id tersedia. | High | Trace dasar antar request/event dapat dicari. |
| Maintainability | Service boundary jelas dan tidak ada cross-service DB query. | Critical | Code review menolak query langsung ke DB service lain. |
| Scalability | Data model mendukung foundation_id dan school_id. | Critical | Semua data utama ter-scope dengan benar. |
| Data Integrity | Finance menggunakan decimal/NUMERIC, bukan float. | Critical | Tidak ada float untuk amount. |
| Usability | UI menggunakan label Bahasa Indonesia. | High | User operasional memahami menu dan alur utama. |
| Mobile Accessibility | Parent/student/guru quick features dapat diakses mobile. | Medium | Mobile app/web responsive menyediakan flow utama. |

---

## 12. Data Privacy and Compliance Requirements

### 12.1 Data Classification

| Classification | Description | Examples | Access Rule |
|---|---|---|---|
| Public | Data yang boleh diketahui publik. | Nama sekolah, informasi umum yayasan. | Dapat diakses publik jika memang disetujui. |
| Internal | Data operasional non-sensitif. | Kalender umum, pengumuman internal umum. | Hanya user internal terkait. |
| Restricted | Data pribadi dan operasional penting. | Data siswa, orang tua, guru, absensi, nilai, tagihan, pembayaran, dokumen PPDB. | Role/scope terbatas, audit untuk download/export. |
| Confidential | Data sangat sensitif. | Credential, token, backup, dokumen legal, data kesehatan/BK detail jika ada, severe cases. | Akses sangat terbatas, tidak boleh muncul di log/event/notification mentah. |

### 12.2 Privacy Rules

- Data anak harus diakses hanya oleh pihak yang memiliki relasi dan izin.
- Orang tua hanya dapat melihat data anak yang terhubung.
- Siswa hanya dapat melihat data pribadi yang sudah dipublish/diizinkan.
- Guru hanya dapat melihat data kelas/mapel yang ditugaskan.
- Kepala Sekolah hanya dapat melihat data sekolah yang dipimpin.
- Admin Yayasan dapat melihat data lintas unit sesuai role dan permission.
- File private by default.
- Download/export Restricted/Confidential wajib permission dan audit.
- Signed URL memiliki masa berlaku pendek sesuai klasifikasi.
- Consent orang tua diperlukan untuk penggunaan data anak dalam konteks pendaftaran, komunikasi, dan publikasi jika digunakan.
- Raw sensitive data tidak boleh masuk application log, event payload, notification body, atau error response.

---

## 13. Permission and Scope Requirements

### 13.1 Access Model

`school-platform` menggunakan kombinasi:

```text
RBAC + ABAC/scope + object-level authorization
```

RBAC menentukan role dan permission.  
ABAC/scope menentukan batas akses berdasarkan foundation, school, class, subject, student, atau child relation.  
Object-level authorization memastikan setiap resource by ID benar-benar berada dalam scope user.

### 13.2 Scope Types

| Scope | Description | Example |
|---|---|---|
| Foundation Scope | Akses lintas unit dalam satu yayasan. | Admin Yayasan melihat dashboard yayasan. |
| School Scope | Akses hanya dalam satu sekolah. | Kepala Sekolah SD hanya melihat data SD. |
| Class Scope | Akses hanya kelas tertentu. | Wali Kelas 5A melihat siswa kelas 5A. |
| Subject Scope | Akses mapel tertentu. | Guru Matematika melihat kelas/mapel assignment. |
| Student Scope | Akses data siswa tertentu. | Siswa melihat rapor pribadi. |
| Parent-Child Scope | Orang tua mengakses data anak yang terhubung. | Orang tua melihat tagihan anaknya. |

### 13.3 Examples

| Scenario | Allowed? | Reason |
|---|---|---|
| Bendahara SD melihat tagihan siswa SD dalam school scope | Yes | Role dan school scope sesuai. |
| Bendahara SD melihat tagihan siswa SMP tanpa assignment | No | Cross-school scope tidak valid. |
| Guru melihat nilai kelas yang diajar | Yes | Subject/class assignment valid. |
| Guru melihat nilai kelas lain | No | Object-level authorization gagal. |
| Orang tua melihat rapor anak sendiri yang published | Yes | Parent-child relation valid dan report published. |
| Orang tua melihat rapor anak lain | No | Parent-child scope tidak valid. |
| Siswa melihat draft rapor | No | Hanya published report yang boleh dilihat. |
| API Gateway menentukan final business authorization | No | Business authorization tetap di service. |

---

## 14. Approval and Audit Requirements

### 14.1 Actions Requiring Audit

Aksi berikut wajib diaudit:

- login/logout/security events tertentu;
- role assignment;
- permission update;
- student sensitive data update;
- guardian relation update;
- teacher assignment update;
- PPDB accept/reject decision;
- applicant conversion to student;
- file download/export Restricted/Confidential;
- import Excel confirm;
- fee policy create/update/approve;
- bill generation;
- payment proof upload;
- payment verification/rejection;
- payment void;
- attendance correction;
- grade book submit/revision;
- report card publish;
- report card revision after publish;
- announcement publish;
- production release approval if recorded in workflow.

### 14.2 Actions Requiring Approval

Aksi berikut memerlukan approval dalam MVP atau disiapkan untuk approval:

| Action | Approver | Notes |
|---|---|---|
| Fee policy sensitive change | Kepala Sekolah/Admin Yayasan | Terutama free_spp, scholarship, custom_fee. |
| Large/exceptional finance action | Kepala Sekolah/Admin Yayasan | Sesuai threshold yang ditentukan. |
| Payment void | Kepala Sekolah | Harus ada reason dan audit. |
| Report card publish | Kepala Sekolah | Setelah review Wali Kelas. |
| Report card revision after publish | Kepala Sekolah | Published report locked. |
| Cross-school/foundation announcement | Admin Yayasan | Jika target lintas unit. |
| Sensitive role assignment | Admin Yayasan | Terutama role admin/finance/approval. |

---

## 15. Reporting Requirements

### 15.1 Dashboard Yayasan

Metrics MVP:

- jumlah siswa aktif per unit sekolah;
- jumlah guru/staff per unit;
- ringkasan PPDB per sekolah;
- total tagihan;
- total pembayaran;
- outstanding/tunggakan;
- collection rate;
- pending approvals;
- pengumuman penting;
- status backup/health jika relevan.

### 15.2 Dashboard Sekolah

Metrics MVP:

- jumlah siswa aktif;
- jumlah kelas;
- jumlah guru;
- PPDB sekolah;
- tagihan dan pembayaran sekolah;
- outstanding per periode;
- absensi hari ini;
- progres input nilai;
- progres publish rapor;
- pending approval sekolah.

### 15.3 Dashboard Guru

Metrics MVP:

- jadwal mengajar;
- kelas/mapel assignment;
- absensi yang perlu diinput;
- progres nilai;
- pengumuman terkait guru;
- notifikasi tugas/revision.

### 15.4 Dashboard Orang Tua/Siswa

Metrics MVP:

- tagihan aktif;
- status pembayaran;
- pengumuman terbaru;
- jadwal;
- rapor published;
- notifikasi penting.

### 15.5 Reporting Constraints

- Reporting Service membaca `reporting_db` saja.
- Reporting tidak melakukan query langsung ke database operasional service lain.
- Data dashboard dibangun dari domain events.
- Consumer harus idempotent.
- Dashboard harus enforce role/scope.
- Reporting bukan source of truth.

---

## 16. Platform Requirements

### 16.1 Web Admin

Web Admin digunakan oleh:

- Admin Yayasan;
- Kepala Sekolah;
- TU/Staff;
- Bendahara;
- Guru.

Kebutuhan utama:

- dashboard;
- data management;
- approval;
- import;
- finance operation;
- academic operation;
- reporting;
- QA/admin workflow.

### 16.2 Mobile App

Mobile App digunakan oleh:

- Orang Tua;
- Siswa;
- Guru untuk quick features jika relevan.

Kebutuhan utama:

- login;
- lihat tagihan;
- upload bukti;
- lihat pengumuman;
- lihat jadwal;
- lihat rapor published;
- notifikasi.

### 16.3 API Gateway

API Gateway adalah satu-satunya external API entrypoint.

Kebutuhan:

- REST/JSON external API;
- token validation;
- request_id/correlation_id;
- routing ke service;
- rate limit dasar jika diperlukan;
- tidak berisi business logic.

### 16.4 File Privacy

- File disimpan di object storage S3-compatible.
- File private by default.
- Backend melakukan RBAC/ABAC/object-level authorization sebelum memberikan signed URL.
- File classification wajib ditentukan.
- Official files versioned, bukan overwritten.

### 16.5 Environment Workflow

| Environment | Purpose |
|---|---|
| Local | Development individual menggunakan Docker Compose. |
| Staging | QA/UAT dan release candidate. |
| Production | Pilot/operasional resmi. |

Branch flow:

```text
feature/* → develop → staging → main/production
```

---

## 17. Success Metrics

| Metric | Target MVP |
|---|---|
| Data master readiness | Siswa, guru, kelas, guardian dapat dibuat/import. |
| PPDB readiness | Applicant dapat diproses hingga conversion. |
| Finance readiness | Bill dapat digenerate dan payment manual dapat diverifikasi. |
| Academic readiness | Jadwal, absensi, nilai dasar dapat berjalan. |
| Report card readiness | Rapor dapat dipublish dan dilihat parent/student. |
| Communication readiness | Announcement dan notification dasar berjalan. |
| Reporting readiness | Dashboard dasar tersedia sesuai role/scope. |
| Security readiness | Permission/scope dan object-level auth diuji. |
| Audit readiness | Aksi sensitif memiliki audit log. |
| Backup readiness | Backup dan restore test terdokumentasi. |
| QA readiness | Tidak ada Critical/High bug pada core flow. |
| AI Agent readiness | Dokumen dan prompt mendukung task kecil dan review. |

---

## 18. Risks and Mitigations

| Risk | Impact | Probability | Mitigation | Owner |
|---|---|---|---|---|
| Scope creep | MVP melebar dan terlambat | High | Gunakan non-goals dan sprint scope sebagai batas. | Product Owner |
| Data migration error | Data awal tidak valid | Medium | Validation-preview-confirm-report untuk import. | Backend, QA |
| Permission/scope bug | Data lintas sekolah bocor | High | Permission test, object-level auth, security review. | Backend, Security Reviewer |
| Finance calculation bug | Tagihan salah | Medium | Decimal/NUMERIC, snapshot, test idempotency. | Backend, QA |
| Privacy leak | Data anak/keuangan terekspos | Medium | Data classification, masking, audit, signed URL. | Security Reviewer |
| Reporting delay | Dashboard tidak update cepat | Medium | Event projection, scheduled rebuild/sync. | Backend |
| AI Agent output inconsistency | Kode/dokumen tidak sesuai standar | Medium | AGENTS.md, SKILLS.md, prompt task kecil, human review. | Tech Lead |
| Service boundary violation | Arsitektur sulit maintain | Medium | Code review, lint/checklist, docs boundary. | Backend Lead |
| Adoption issue | User operasional sulit menggunakan sistem | Medium | UI Bahasa Indonesia, UAT, training, feedback loop. | Product Owner |
| Backup restore failure | Risiko kehilangan data | Low-Medium | Restore test bulanan dan sebelum production. | DevOps |
| Security hardening terlambat | Release tertunda | Medium | Security checks sejak awal sprint, bukan hanya akhir. | DevOps, QA |

---

## 19. MVP Completion Criteria

MVP dianggap selesai jika seluruh kriteria berikut terpenuhi:

1. User dapat login sesuai role.
2. JWT access token dan rotating refresh token berjalan.
3. Role, permission, dan scope diterapkan di backend.
4. Foundation, school, academic year, semester dapat dikelola.
5. Student, guardian, teacher, class, dan assignment dapat dikelola.
6. Import Excel data awal berjalan dengan validation-preview-confirm-report.
7. File private by default dan signed URL berjalan.
8. PPDB dapat berjalan dari applicant hingga conversion to student.
9. Fee type, fee scheme, fee policy, dan bill generation berjalan.
10. Bill menyimpan snapshot amount dan policy.
11. Orang tua dapat upload bukti pembayaran.
12. Bendahara dapat verify/reject payment.
13. Payment verification dan void diaudit.
14. Academic basic seperti subject, schedule, dan attendance berjalan.
15. Guru dapat input nilai sesuai assignment.
16. Rapor dapat direview, dipublish, locked, dan dilihat parent/student.
17. Revision after publish membutuhkan approval dan audit.
18. Announcement dan notification dasar berjalan.
19. Dashboard dasar tersedia sesuai role/scope.
20. Reporting menggunakan reporting_db dan event projection.
21. Audit log tersedia untuk aksi sensitif.
22. Backup dan restore test berhasil serta terdokumentasi.
23. Structured logging dan metrics dasar tersedia.
24. QA/UAT core flow pass.
25. Tidak ada Critical/High bug terbuka pada core flow MVP.
26. Dokumentasi teknis, workflow, dan prompt AI Agent tersedia.

---

## 20. Assumptions

Asumsi yang digunakan dalam PRD ini:

1. Project digunakan pertama kali untuk satu yayasan yang memiliki beberapa unit sekolah.
2. Unit sekolah minimum mencakup TK, SD, SMP, dan SMA.
3. Sistem dirancang internal-first, tetapi tetap SaaS-ready.
4. Payment gateway belum digunakan pada MVP.
5. Pembayaran SPP MVP menggunakan manual payment dan upload bukti.
6. WhatsApp belum digunakan pada MVP.
7. Mobile app MVP online-only.
8. Web admin menjadi platform utama untuk operasional staff, bendahara, kepala sekolah, admin yayasan, dan guru.
9. Orang tua dan siswa menggunakan mobile app atau mobile-friendly interface untuk akses dasar.
10. Data production tidak digunakan di local development kecuali sudah dianonimkan dan disetujui.
11. Semua service menggunakan PostgreSQL database masing-masing.
12. Reporting tidak membaca database operasional secara langsung.
13. Semua file Restricted/Confidential harus private.
14. AI Agent membantu penyusunan dokumen dan implementasi task kecil, tetapi output tetap direview manusia.

---

## 21. Open Questions

Pertanyaan yang masih perlu diputuskan:

1. Apakah satu user dapat memiliki role berbeda di beberapa sekolah dalam yayasan yang sama?
2. Apakah siswa memiliki akun login sendiri sejak MVP, atau opsional per jenjang?
3. Apakah semua orang tua wajib memiliki akun, atau bisa dibuat saat pembayaran/PPDB?
4. Bagaimana format final nomor dokumen untuk rapor, receipt, dan dokumen PPDB?
5. Berapa threshold finance action yang membutuhkan approval Admin Yayasan?
6. Apakah fee policy tertentu dapat disetujui Kepala Sekolah saja atau harus Admin Yayasan?
7. Apakah report card MVP perlu export PDF pada tahap pertama, atau cukup view digital?
8. Apakah mobile app MVP mencakup guru quick features, atau hanya parent/student?
9. Apakah email provider sudah ditentukan untuk production?
10. Apakah FCM wajib aktif pada pilot pertama atau cukup provider abstraction?
11. Apakah struktur kelas TK/PAUD membutuhkan model berbeda dari SD/SMP/SMA?
12. Apakah format rapor TK/PAUD sudah ditentukan oleh yayasan?
13. Apakah data alumni langsung dibuat saat siswa graduated, atau post-MVP?
14. Siapa final approver untuk production release pada tahap pilot?
15. Apakah staging server sudah tersedia sebelum Sprint 0 selesai?

---

## 22. Appendix

### 22.1 Glossary

| Term | Meaning |
|---|---|
| Foundation | Yayasan yang mengelola beberapa unit sekolah. |
| School | Unit sekolah seperti TK, SD, SMP, SMA. |
| RBAC | Role-Based Access Control. |
| ABAC | Attribute-Based Access Control. |
| Scope | Batas akses berdasarkan foundation, school, class, subject, student, atau child relation. |
| Object-Level Authorization | Validasi akses terhadap resource spesifik by ID. |
| PPDB | Penerimaan Peserta Didik Baru. |
| Fee Policy | Kebijakan biaya siswa seperti diskon, bebas SPP, beasiswa. |
| Bill Snapshot | Salinan amount/policy pada saat tagihan dibuat. |
| Report Card | Rapor siswa. |
| Published | Status data yang sudah dirilis ke parent/student. |
| Locked | Status data yang tidak boleh diubah tanpa approval/revision. |
| Read Model | Data projection untuk kebutuhan baca/dashboard. |
| Audit Log | Catatan aksi penting/sensitif. |
| Signed URL | URL sementara untuk akses file private. |
| Restricted Data | Data pribadi/operasional penting. |
| Confidential Data | Data sangat sensitif seperti credential, token, backup, legal/health/severe cases. |

### 22.2 Role Summary

| Role | Summary |
|---|---|
| Admin Yayasan | Mengelola yayasan, konfigurasi lintas sekolah, dan dashboard yayasan. |
| Kepala Sekolah | Mengawasi sekolah, approval, dan publish rapor. |
| TU/Staff | Mengelola data administrasi, import, dan PPDB. |
| Bendahara | Mengelola tagihan, pembayaran, dan verifikasi. |
| Guru | Mengelola jadwal, absensi, dan nilai sesuai assignment. |
| Wali Kelas | Mereview data kelas dan rapor. |
| Orang Tua/Wali Murid | Melihat data anak, tagihan, pembayaran, pengumuman, dan rapor. |
| Siswa | Melihat informasi akademik pribadi yang dipublish. |

### 22.3 Module Summary

| Module | MVP Purpose |
|---|---|
| Identity & Access | Login, token, role, permission, user context. |
| School Core | Master data yayasan/sekolah/siswa/guru/kelas. |
| File Management + Import Excel | Private file dan migrasi data awal. |
| PPDB | Pendaftaran dan konversi siswa baru. |
| Finance/SPP | Tagihan dan pembayaran manual. |
| Academic Basic | Jadwal, absensi, dan nilai dasar. |
| Report Card/E-Rapor Basic | Review, publish, lock rapor. |
| Communication/Notification | Announcement dan notification. |
| Reporting Dashboard | Dashboard berbasis projection. |
| Security/Observability/Backup/UAT | Production readiness. |

### 22.4 MVP vs Post-MVP Summary

| Area | MVP | Post-MVP |
|---|---|---|
| Payment | Manual payment + upload proof | Payment gateway, bank reconciliation |
| Notification | In-app, FCM/email abstraction | WhatsApp, SMS, advanced campaign |
| Academic | Schedule, attendance, score, report basic | LMS, advanced analytics, integrations |
| Reporting | Dashboard basic projection | BI/data warehouse |
| Search | Local search/filter | Global Search Service |
| Mobile | Online-only main views | Offline write/sync |
| Infrastructure | Docker Compose local/staging | Kubernetes if scale requires |
| HR/Payroll | Not included | Full HR/payroll module |
| Library/Asset | Not included | Dedicated modules |
| BK/UKS | Not included in detail | Sensitive module with stronger privacy |
