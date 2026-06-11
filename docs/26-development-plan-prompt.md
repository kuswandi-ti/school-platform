# 26 — Development Plan Prompt

Project: `school-platform`  
Purpose: Prompt for AI Agent to generate Development Plan  
Target output: `docs/26-development-plan.md`

---

## Prompt

```text
Kamu adalah Senior Engineering Manager, Software Architect, Technical Project Manager, dan DevOps-minded Delivery Lead.

Tugasmu adalah membuat dokumen Development Plan untuk project `school-platform`.

Dokumen ini harus menjelaskan rencana implementasi MVP secara teknis, bertahap, praktis, dan bisa digunakan oleh Developer, QA, DevOps, Product Owner, dan AI Agent.

---

# 1. Context Project

`school-platform` adalah platform internal yayasan sekolah multi-unit untuk TK, SD, SMP, dan SMA.

Project ini menggunakan arsitektur:

- Monorepo
- Go microservices
- Custom Go API Gateway
- gRPC + protobuf untuk komunikasi internal
- PostgreSQL database per service
- RabbitMQ untuk domain events
- Redis
- MinIO / S3-compatible object storage
- Next.js untuk web admin
- Flutter untuk mobile app
- GitHub Actions untuk CI/CD
- Docker Compose untuk local development dan staging awal
- GitHub branch workflow: feature/* → develop → staging → main/production

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
7. docs/08-coding-standard.md
8. docs/09-ai-agent-rules.md
9. docs/10-sprint-backlog-mvp.md
10. dokumen sprint/task terkait

---

# 3. MVP Modules

MVP mencakup:

- API Gateway
- Identity & Access
- School Core
- File Management + Import Excel
- PPDB
- Finance/SPP manual
- Academic Basic
- Report Card/E-Rapor Basic
- Communication/Notification
- Reporting Dashboard
- Security, Observability, Backup, UAT Hardening

MVP tidak mencakup:

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

---

# 4. Team Structure

Tim MVP terdiri dari:

- Backend Developer
- Frontend Developer
- QA
- Infrastructure/DevOps
- AI Agent as assistant

Jelaskan juga bagaimana AI Agent digunakan untuk task kecil, bukan sebagai final authority.

---

# 5. Development Workflow

Gunakan workflow:

- feature/* → develop → staging → main/production
- Semua perubahan melalui Pull Request
- CI wajib pass
- Review wajib
- QA sign-off sebelum production
- Production deploy hanya dari main
- Production deploy wajib manual approval
- Hotfix dari main dan back-merge ke staging/develop

---

# 6. Sprint Plan

Gunakan urutan sprint berikut:

- Sprint 0 — Project Foundation
- Sprint 1 — Identity & Access
- Sprint 2 — School Core
- Sprint 3 — File Management + Import Excel
- Sprint 4 — PPDB
- Sprint 5 — Finance/SPP
- Sprint 6 — Academic Basic
- Sprint 7 — Report Card/E-Rapor Basic
- Sprint 8 — Communication/Notification
- Sprint 9 — Reporting Dashboard
- Sprint 10 — Security, Observability, Backup, UAT Hardening

---

# 7. Development Plan Goal

Buat dokumen Development Plan yang bisa dipakai sebagai panduan implementasi MVP.

Dokumen harus menjawab:

- bagaimana project dikembangkan
- siapa mengerjakan apa
- bagaimana sprint dijalankan
- apa dependency antar sprint
- bagaimana testing dilakukan
- bagaimana AI Agent digunakan
- bagaimana CI/CD, staging, production, dan release dijalankan
- bagaimana quality gate dijaga
- bagaimana security, audit, backup, dan observability dimasukkan sejak awal

Dokumen harus siap disimpan sebagai:

`docs/26-development-plan.md`

---

# 8. Output Format

Buat dokumen Markdown dengan struktur:

# Development Plan — School Platform MVP

## 1. Executive Summary
Jelaskan ringkasan strategi development MVP.

## 2. Development Principles
Tuliskan prinsip:
- local-first development
- small task delivery
- service boundary discipline
- test-driven quality gate
- security from start
- audit and permission from early sprint
- event-driven reporting/notification
- no cross-service database query
- documentation as source of truth
- AI Agent with human review

## 3. Team and Responsibilities
Buat tabel:
| Role | Main Responsibilities | Deliverables | Review Responsibility |
|---|---|---|---|

Roles:
- Backend Developer
- Frontend Developer
- QA
- Infrastructure/DevOps
- AI Agent
- Product Owner/Reviewer jika relevan

## 4. Repository and Branching Strategy
Jelaskan:
- branch develop, staging, main
- feature branch
- fix branch
- hotfix branch
- PR rule
- CI
- branch protection
- release tag
- hotfix workflow

## 5. GitHub Project Management
Jelaskan:
- GitHub Project: School Platform MVP
- labels
- milestones Sprint 0–Sprint 10
- issue lifecycle
- PR relation to issue
- AI Agent labels

Sertakan field:
- Status
- Sprint
- Priority
- Area
- Type
- Owner
- Estimate
- Risk
- Platform
- AI Agent
- Target Release

## 6. Environment Strategy
Jelaskan:
- local
- staging
- production
- Docker Compose
- GitHub Environments
- secrets separation
- local data policy
- production deployment approval

## 7. Development Phases
Buat fase:
- Phase 1 — Platform Foundation
- Phase 2 — Admission & Finance
- Phase 3 — Academic & Communication
- Phase 4 — Reporting & Production Readiness

Untuk setiap fase:
- objective
- related sprint
- key deliverables
- risk

## 8. Sprint Plan Overview
Buat tabel:
| Sprint | Objective | Main Modules | Key Deliverables | Dependencies | Owner | Exit Criteria |
|---|---|---|---|---|---|---|

## 9. Detailed Sprint Plan
Untuk Sprint 0 sampai Sprint 10, jelaskan:

### Sprint N — Nama Sprint
- Objective
- Scope
- Out of Scope
- Backend Tasks
- Frontend Tasks
- Mobile Tasks jika relevan
- QA Tasks
- DevOps Tasks
- AI Agent Suitable Tasks
- Human Review Required Tasks
- Dependencies
- Deliverables
- Acceptance Criteria
- Definition of Done
- Risks

Pastikan Sprint 0 sampai Sprint 10 lengkap.

## 10. Dependency Map
Jelaskan dependency antar sprint.

Contoh:
- Finance butuh School Core
- Academic butuh School Core
- Report Card butuh Academic Basic
- Reporting butuh event dari modul lain
- Communication butuh event dari domain lain
- PPDB conversion butuh School Core

Buat tabel dependency.

## 11. AI Agent Usage Plan
Jelaskan:
- jenis task yang boleh dikerjakan AI Agent
- jenis task yang wajib human review
- task yang tidak boleh dikerjakan AI Agent
- cara menggunakan AGENTS.md
- cara menggunakan SKILLS.md
- cara menggunakan sprint task prompt
- format output wajib AI Agent
- quality control terhadap output AI Agent

## 12. Quality Gates
Jelaskan quality gate:
- lint
- formatting
- unit test
- integration test
- API test
- permission/scope test
- event test
- audit test
- build
- review
- QA sign-off

Buat tabel quality gate per branch.

## 13. Testing Strategy
Jelaskan testing per sprint:
- unit test
- integration test
- API test
- permission/scope test
- event test
- audit test
- frontend test
- mobile test
- regression test
- UAT

Buat matrix test per module.

## 14. Security and Compliance Plan
Jelaskan:
- authentication
- authorization
- object-level auth
- field-level access
- audit
- file privacy
- signed URL
- secrets
- backup
- data anak
- data orang tua
- data keuangan
- Restricted/Confidential data

## 15. Observability Plan
Jelaskan:
- structured JSON logging
- request_id
- correlation_id
- healthz
- readyz
- metrics
- Prometheus
- Grafana
- Loki
- RabbitMQ monitoring
- DLQ monitoring

## 16. Backup and Restore Plan
Jelaskan:
- backup database per service
- object storage backup
- retention
- encryption
- restore test
- RPO
- RTO
- backup before risky migration

## 17. Release Plan
Jelaskan:
- develop to staging
- staging QA/UAT
- staging to main
- production approval
- release tag
- post-release verification
- rollback

## 18. Hotfix Plan
Jelaskan:
- kapan hotfix dipakai
- branch hotfix dari main
- PR ke main
- deploy production
- release tag
- back-merge ke staging dan develop

## 19. Risk Management
Buat tabel:
| Risk | Impact | Probability | Mitigation | Owner |
|---|---|---|---|---|

## 20. MVP Readiness Criteria
Tuliskan kriteria MVP siap pilot/production.

## 21. Appendix
Tambahkan:
- recommended GitHub labels
- GitHub Project fields
- sprint milestone list
- useful local commands
- useful CI commands
- AI Agent prompt usage example

---

# 9. Writing Rules

- Gunakan Bahasa Indonesia.
- Buat detail dan actionable.
- Jangan menambahkan modul di luar MVP.
- Gunakan keputusan arsitektur yang sudah ditentukan.
- Jangan membuat asumsi timeline tanggal spesifik kecuali diminta.
- Jika perlu estimasi, gunakan relatif: Sprint 0, Sprint 1, dst.
- Format harus siap disimpan sebagai `docs/26-development-plan.md`.
- Jangan menulis kode implementasi.
- Gunakan tabel untuk sprint overview, team responsibilities, quality gate, dependency, risk.
```

## GitHub Project Management Context

When generating or updating this document, also consider:

```text
docs/25-github-project-management.md
.github/CODEOWNERS
.github/pull_request_template.md
.github/workflows/ci.yml
.github/ISSUE_TEMPLATE/feature_task.yml
.github/ISSUE_TEMPLATE/bug_report.yml
.github/ISSUE_TEMPLATE/ai_agent_task.yml
.github/ISSUE_TEMPLATE/security_review.yml
.github/ISSUE_TEMPLATE/qa_uat.yml
```

Ensure GitHub labels, milestones, project fields, issue lifecycle, PR workflow, AI Agent task workflow, QA/UAT workflow, and release readiness are aligned with `docs/25-github-project-management.md`.

## Development Plan Final Document Status

The final generated document is now available at:

```text
docs/26-development-plan.md
```

Use this prompt file only when regenerating or updating the final document.

When making product, planning, workflow, sprint, GitHub issue, PR, QA/UAT, or implementation decisions, read the final document first.
