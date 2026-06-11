# 25 — PRD Prompt

Project: `school-platform`  
Purpose: Prompt for AI Agent to generate Product Requirement Document  
Target output: `docs/25-product-requirement-document.md`

---

## Prompt

```text
Kamu adalah Senior Product Manager, System Analyst, dan Product Strategy Consultant.

Tugasmu adalah membuat dokumen PRD / Product Requirement Document untuk project `school-platform`.

Dokumen PRD ini harus komprehensif, relevan dengan konteks project, dan siap dijadikan acuan oleh Product Owner, Developer, QA, DevOps, dan AI Agent.

---

# 1. Context Project

`school-platform` adalah platform internal yayasan sekolah multi-unit untuk mendukung operasional TK, SD, SMP, dan SMA.

Platform ini digunakan oleh:

- Admin Yayasan
- Kepala Sekolah
- Tata Usaha / Staff
- Bendahara Sekolah
- Guru
- Wali Kelas
- Orang Tua / Wali Murid
- Siswa

Project ini awalnya digunakan sebagai sistem internal yayasan, tetapi tetap dirancang SaaS-ready.

---

# 2. Source of Truth

Gunakan dokumen berikut sebagai source of truth:

- AGENTS.md
- SKILLS.md
- docs/01-technical-architecture.md
- docs/02-service-boundary.md
- docs/03-data-model-mvp.md
- docs/04-api-contract.md
- docs/05-event-contract.md
- docs/06-ui-screen-user-flow.md
- docs/07-test-plan-acceptance-criteria.md
- docs/08-coding-standard.md
- docs/09-ai-agent-rules.md
- docs/10-sprint-backlog-mvp.md
- docs/11-github-repository-rules.md
- docs/12-ai-agent-sprint-prompts.md
- docs/13-sprint-0-task-prompts.md sampai docs/23-sprint-10-task-prompts.md
- docs/24-local-development-guide.md

Jika ada konflik antar dokumen, prioritaskan:

1. AGENTS.md
2. docs/01-technical-architecture.md
3. docs/02-service-boundary.md
4. docs/03-data-model-mvp.md
5. docs/04-api-contract.md
6. docs/05-event-contract.md
7. docs/06-ui-screen-user-flow.md
8. docs/07-test-plan-acceptance-criteria.md
9. docs/08-coding-standard.md
10. docs/09-ai-agent-rules.md
11. docs/10-sprint-backlog-mvp.md
12. dokumen sprint/task terkait

---

# 3. Architecture and Product Constraints

Gunakan konteks teknis berikut sebagai batasan produk:

- Backend menggunakan Go microservices.
- API Gateway custom menggunakan Go + Chi.
- Komunikasi internal antar-service menggunakan gRPC + protobuf.
- Database menggunakan PostgreSQL, satu database per service.
- Event menggunakan RabbitMQ topic exchange `domain.events`.
- Redis digunakan untuk cache/rate limit/lock jika diperlukan.
- Web admin menggunakan Next.js.
- Mobile app menggunakan Flutter.
- File menggunakan object storage S3-compatible seperti MinIO/R2.
- Reporting menggunakan read model/projection, bukan query langsung ke database operasional service lain.
- Branch workflow: feature/* → develop → staging → main/production.

---

# 4. Important Existing Decisions

Gunakan keputusan berikut sebagai dasar PRD:

- MVP adalah platform internal yayasan multi-unit untuk operasional dasar TK, SD, SMP, dan SMA.
- MVP mencakup:
  - Identity & Access
  - School Core
  - File Management + Import Excel
  - PPDB
  - Finance/SPP manual
  - Academic Basic
  - Report Card/E-Rapor Basic
  - Communication/Notification
  - Reporting Dashboard
  - Security, Observability, Backup, and UAT Hardening
- MVP tidak mencakup:
  - Payroll
  - HR lengkap
  - Asset/Inventory lengkap
  - Library
  - BK/UKS detail
  - LMS penuh
  - Alumni/Tracer
  - Koperasi
  - Global Search
  - Payment Gateway
  - WhatsApp
  - Offline Write Mobile
  - Kubernetes
- Authentication menggunakan JWT access token + rotating refresh token.
- Authorization menggunakan RBAC + ABAC/scope.
- Data utama wajib mendukung `foundation_id` dan `school_id`.
- File private by default.
- Aksi sensitif wajib audit log dan approval bila diperlukan.
- Finance SPP MVP menggunakan manual payment + upload bukti pembayaran.
- Free SPP, discount, beasiswa, sibling discount adalah Finance fee policy, bukan status siswa.
- Bill harus menyimpan snapshot.
- Report card menggunakan workflow draft → submitted → reviewed/approved → published/locked.
- Revisi setelah publish harus approval dan audit log.
- Dashboard menggunakan Reporting Service dan `reporting_db` sebagai read model/projection.
- Tidak boleh ada query lintas database service.
- UI labels menggunakan Bahasa Indonesia.
- Internal code menggunakan English.

---

# 5. PRD Goal

Buat dokumen PRD yang menjelaskan:

- masalah yang diselesaikan
- target user
- tujuan produk
- scope MVP
- non-goals
- functional requirements
- non-functional requirements
- user journeys
- acceptance criteria
- success metrics
- risks and mitigations
- MVP completion criteria

Dokumen harus siap disimpan sebagai:

`docs/25-product-requirement-document.md`

---

# 6. Output Format

Buat dokumen Markdown dengan struktur:

# Product Requirement Document — School Platform

## 1. Executive Summary
Jelaskan ringkasan produk, tujuan utama, dan masalah yang diselesaikan.

## 2. Background and Problem Statement
Jelaskan kondisi yayasan/sekolah, masalah operasional, dan kenapa sistem ini dibutuhkan.

Bahas masalah seperti:
- data siswa/guru tersebar
- PPDB belum terpusat
- SPP manual sulit direkonsiliasi
- laporan akademik belum terintegrasi
- komunikasi sekolah-orang tua belum konsisten
- dashboard yayasan belum real-time
- audit dan akses data belum tertata

## 3. Product Vision
Jelaskan:
- MVP vision
- post-MVP vision
- SaaS-ready direction

## 4. Product Goals
Buat daftar goal utama produk.

## 5. Non-Goals
Tuliskan hal-hal yang tidak termasuk MVP dan alasannya.

## 6. Target Users and Roles
Untuk setiap role, jelaskan:
- tujuan penggunaan
- kebutuhan utama
- contoh aktivitas
- batasan akses

## 7. User Personas
Buat persona ringkas untuk:
- Admin Yayasan
- Kepala Sekolah
- TU/Staff
- Bendahara
- Guru/Wali Kelas
- Orang Tua
- Siswa

Setiap persona mencakup:
- profile
- goals
- pain points
- main tasks
- success expectation

## 8. MVP Scope
Jelaskan scope MVP per modul:
- Identity & Access
- School Core
- File Management + Import Excel
- PPDB
- Finance/SPP
- Academic Basic
- Report Card/E-Rapor Basic
- Communication/Notification
- Reporting Dashboard
- Security, Observability, Backup, and UAT Hardening

Untuk setiap modul, tuliskan:
- objective
- fitur utama
- user role yang menggunakan
- key workflow
- data utama
- acceptance criteria ringkas

## 9. User Journey
Buat user journey untuk:
- Admin Yayasan setup sekolah
- TU import data awal
- PPDB sampai konversi siswa
- Bendahara generate tagihan dan verifikasi pembayaran
- Guru input absensi
- Guru input nilai
- Wali Kelas review data rapor
- Kepala Sekolah publish rapor
- Orang Tua melihat tagihan dan upload bukti pembayaran
- Orang Tua melihat rapor anak
- Siswa melihat jadwal/nilai/rapor published

Untuk setiap journey, jelaskan:
- actor
- trigger
- steps
- output
- exception/error cases
- audit/approval jika relevan

## 10. Functional Requirements
Buat requirement dengan format tabel:

| ID | Module | Requirement | Role | Priority | Acceptance Criteria |
|---|---|---|---|---|---|

Gunakan ID:
- PRD-ID-001
- PRD-SC-001
- PRD-FILE-001
- PRD-ADM-001
- PRD-FIN-001
- PRD-ACD-001
- PRD-RC-001
- PRD-COM-001
- PRD-RPT-001
- PRD-SEC-001

## 11. Non-Functional Requirements
Cakup:
- security
- privacy
- auditability
- performance
- availability
- backup/restore
- observability
- maintainability
- scalability/SaaS-readiness
- data integrity
- usability
- mobile accessibility

## 12. Data Privacy and Compliance Requirements
Jelaskan klasifikasi data:
- Public
- Internal
- Restricted
- Confidential

Sertakan aturan:
- data anak
- data orang tua
- data guru
- data keuangan
- dokumen siswa
- file private
- signed URL
- audit download/export
- consent orang tua jika diperlukan
- masking data sensitif
- raw sensitive data tidak boleh masuk log

## 13. Permission and Scope Requirements
Jelaskan:
- RBAC
- ABAC/scope
- foundation scope
- school scope
- class scope
- subject scope
- student scope
- parent-child scope
- object-level authorization

Berikan contoh akses yang boleh dan tidak boleh.

## 14. Approval and Audit Requirements
Tuliskan aksi sensitif yang wajib approval/audit.

## 15. Reporting Requirements
Tuliskan dashboard:
- Dashboard Yayasan
- Dashboard Sekolah
- Dashboard Guru
- Dashboard Orang Tua/Siswa

Sertakan metrik MVP.

## 16. Platform Requirements
Jelaskan kebutuhan Web Admin, Mobile App, API Gateway, file privacy, local/staging/production workflow.

## 17. Success Metrics
Definisikan metrik keberhasilan MVP.

## 18. Risks and Mitigations
Buat tabel:
| Risk | Impact | Probability | Mitigation | Owner |
|---|---|---|---|---|

## 19. MVP Completion Criteria
Jelaskan kapan MVP dianggap selesai.

## 20. Assumptions
Tuliskan asumsi yang digunakan.

## 21. Open Questions
Tuliskan pertanyaan yang masih perlu diputuskan jika ada.

## 22. Appendix
Tambahkan:
- glossary
- role summary
- module summary
- MVP vs post-MVP summary

---

# 7. Writing Rules

- Gunakan Bahasa Indonesia.
- Gunakan istilah teknis secara konsisten.
- Jangan menambahkan fitur di luar scope MVP.
- Jangan mengubah keputusan arsitektur yang sudah diberikan.
- Jangan terlalu pendek; buat detail dan komprehensif.
- Tulis sebagai dokumen formal yang siap masuk folder `docs/`.
- Jika ada asumsi, tuliskan di bagian “Assumptions”.
- Jika ada keputusan yang belum jelas, masukkan ke “Open Questions”.
- Jangan menulis kode.
- Jangan membuat timeline tanggal spesifik.
- Gunakan tabel bila membuat requirement, roles, risiko, dan metrics.
```

## PRD Final Document Status

The final generated document is now available at:

```text
docs/25-product-requirement-document.md
```

Use this prompt file only when regenerating or updating the final document.

When making product, planning, workflow, sprint, GitHub issue, PR, QA/UAT, or implementation decisions, read the final document first.
