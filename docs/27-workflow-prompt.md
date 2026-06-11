# 27 — Workflow Prompt

Project: `school-platform`  
Purpose: Prompt for AI Agent to generate Workflow/SOP document  
Target output: `docs/27-workflow.md`

---

## Prompt

```text
Kamu adalah Senior Technical Project Manager, DevOps Lead, Engineering Process Consultant, dan QA Process Designer.

Tugasmu adalah membuat dokumen Workflow / SOP kerja harian untuk project `school-platform`.

Dokumen ini harus praktis, actionable, dan bisa dipakai oleh Developer, QA, DevOps, Product Owner, dan AI Agent.

---

# 1. Context Project

`school-platform` adalah sistem manajemen yayasan sekolah multi-unit untuk TK, SD, SMP, dan SMA.

Project ini menggunakan:

- Monorepo
- Go microservices
- Custom Go API Gateway
- PostgreSQL database per service
- RabbitMQ domain events
- Redis
- MinIO / S3-compatible object storage
- Next.js web admin
- Flutter mobile app
- Docker Compose untuk local development
- GitHub Actions untuk CI/CD
- GitHub Projects untuk tracking
- GitHub Issues untuk task
- Pull Request untuk semua perubahan
- Branch develop, staging, main
- AI Agent untuk membantu implementasi task kecil

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
2. SKILLS.md
3. docs/11-github-repository-rules.md
4. docs/24-local-development-guide.md
5. docs/09-ai-agent-rules.md
6. dokumen sprint/task terkait

---

# 3. Existing Workflow Decisions

Gunakan keputusan berikut:

- Branch workflow: feature/* → develop → staging → main/production
- Tidak boleh direct push ke protected branches
- Semua perubahan via Pull Request
- CI wajib pass
- Review wajib
- QA sign-off sebelum production
- Production deploy hanya dari main
- Production deploy harus manual approval via GitHub Environment
- Hotfix dari main menggunakan hotfix/*
- Hotfix wajib back-merge ke staging dan develop
- AI Agent boleh mengerjakan task kecil yang jelas scope-nya
- AI Agent wajib mengikuti AGENTS.md dan SKILLS.md
- AI Agent tidak boleh mengubah arsitektur tanpa instruksi
- AI Agent tidak boleh menangani production secrets atau final production approval
- Development dilakukan local-first
- Staging digunakan untuk QA/UAT
- Production hanya dari main

---

# 4. Goal

Buat dokumen Workflow yang menjadi SOP kerja harian untuk:

- Product Owner / Reviewer
- Backend Developer
- Frontend Developer
- Mobile Developer
- QA
- Infrastructure/DevOps
- AI Agent

Dokumen harus siap disimpan sebagai:

`docs/27-workflow.md`

---

# 5. Output Format

Buat dokumen Markdown dengan struktur:

# Workflow — School Platform

## 1. Purpose
Jelaskan tujuan dokumen workflow.

## 2. Workflow Principles
Tuliskan prinsip:
- local-first development
- PR-based workflow
- protected branch
- small task
- service boundary discipline
- test before PR
- AI Agent as assistant, not final authority
- human review required
- QA sign-off
- production approval
- documentation as source of truth

## 3. Roles in Workflow
Buat tabel:
| Role | Responsibility | Main Output |
|---|---|---|

Roles:
- Product Owner / Reviewer
- Backend Developer
- Frontend Developer
- Mobile Developer
- QA
- Infrastructure/DevOps
- AI Agent

## 4. GitHub Repository Workflow
Jelaskan branch:
- develop
- staging
- main
- feature/*
- fix/*
- hotfix/*
- docs/*
- chore/*
- refactor/*
- test/*

Sertakan aturan merge:
- feature/* → develop
- fix/* → develop
- docs/* → develop
- develop → staging
- staging → main
- hotfix/* → main → back-merge staging/develop

## 5. Branch Protection Workflow
Jelaskan protection untuk:
- develop
- staging
- main

Cakup:
- required PR
- approval
- status checks
- conversation resolution
- block force push
- block deletion
- production environment approval

## 6. GitHub Labels Workflow
Buat daftar label dan jelaskan kapan dipakai:

Type Labels:
- type: feature
- type: bug
- type: chore
- type: docs
- type: refactor
- type: test
- type: security
- type: infra
- type: spike
- type: hotfix

Area Labels:
- area: api-gateway
- area: identity
- area: school-core
- area: admission
- area: academic
- area: finance
- area: communication
- area: reporting
- area: web-admin
- area: mobile
- area: infra
- area: docs
- area: security
- area: observability
- area: ci-cd
- area: file-management

Sprint Labels:
- sprint: 0 sampai sprint: 10

Priority Labels:
- priority: critical
- priority: high
- priority: medium
- priority: low

Status Labels:
- status: ready
- status: in-progress
- status: blocked
- status: needs-review
- status: needs-qa
- status: qa-passed
- status: done

AI Agent Labels:
- ai: ready
- ai: needs-context
- ai: generated
- ai: needs-human-review
- ai: do-not-use-agent

Risk Labels:
- risk: low
- risk: medium
- risk: high
- risk: breaking-change
- risk: migration
- risk: data-sensitive

Review Labels:
- review: backend
- review: frontend
- review: mobile
- review: infra
- review: qa
- review: security
- review: product

## 7. GitHub Project Workflow
Jelaskan GitHub Project:

Project name:
`School Platform MVP`

Fields:
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

Views:
- MVP Board
- Sprint Board
- Backlog Table
- By Area
- By Priority
- QA/UAT
- AI Agent Tasks
- Release Readiness

Jelaskan fungsi tiap view.

## 8. Milestone Workflow
Jelaskan penggunaan milestone:
- Sprint 0 — Project Foundation
- Sprint 1 — Identity & Access
- Sprint 2 — School Core
- Sprint 3 — File Management + Import Excel
- Sprint 4 — PPDB
- Sprint 5 — Finance / SPP
- Sprint 6 — Academic Dasar
- Sprint 7 — Report Card / E-Rapor Dasar
- Sprint 8 — Communication / Notification
- Sprint 9 — Reporting Dashboard
- Sprint 10 — Security, Observability, Backup, UAT Hardening
- MVP Release

## 9. Issue Workflow
Jelaskan lifecycle issue:
- Backlog
- Ready
- In Progress
- In Review
- QA
- Done
- Blocked

Untuk setiap status, jelaskan:
- arti
- siapa yang mengubah status
- syarat pindah status

## 10. Task Creation Workflow
Jelaskan cara membuat task dari sprint prompt:
1. baca sprint doc
2. pilih task kecil
3. buat issue
4. tulis objective
5. tulis scope
6. tulis out of scope
7. tulis acceptance criteria
8. beri label
9. pilih milestone
10. masukkan ke project
11. assign owner
12. tentukan AI Agent status

## 11. AI Agent Workflow
Jelaskan step:
1. Pilih issue dengan `ai: ready`
2. Baca AGENTS.md
3. Baca SKILLS.md
4. Baca sprint task prompt
5. Copy task prompt ke AI Agent
6. Tambahkan konteks file/dokumen
7. Minta AI Agent implement
8. Review hasil
9. Jalankan test
10. Commit ke feature branch
11. Buat PR
12. Human review wajib

Sertakan task yang tidak boleh dikerjakan AI Agent:
- production secrets
- final security approval
- legal/compliance decision
- production deployment approval
- akses data asli sensitif
- perubahan arsitektur besar tanpa instruksi
- final decision untuk data privacy
- hotfix production tanpa human review

## 12. Pull Request Workflow
Jelaskan:
- kapan membuat PR
- PR title convention
- PR description
- required checklist
- reviewer
- CI
- comment resolution
- merge strategy

PR title format:
`type(scope): short description`

## 13. Code Review Workflow
Jelaskan:
- backend review
- frontend review
- mobile review
- infra review
- security review
- QA review

Sertakan checklist review:
- scope sesuai issue
- service boundary aman
- no cross-service DB query
- permission/scope check
- object-level authorization
- audit log
- event if required
- tests
- no sensitive logs
- docs updated if needed

## 14. QA Workflow
Jelaskan:
- test case creation
- test execution
- bug report
- bug severity
- regression test
- UAT
- QA sign-off

Bug severity:
- Critical
- High
- Medium
- Low

Jelaskan release blocking rule:
- Critical/High bug pada core flow memblokir production release.

## 15. Development Local Workflow
Jelaskan step:
1. clone repo
2. checkout develop
3. create feature branch
4. copy .env.example
5. run docker compose
6. run migration
7. run sqlc generate if needed
8. run service
9. run tests
10. commit
11. push
12. PR

Sertakan command contoh.

## 16. CI/CD Workflow
Jelaskan:
- CI on PR
- CI on develop/staging/main
- staging deploy
- production deploy
- manual approval
- rollback

Quality gate:
- lint
- test
- build
- docker build
- migration check
- proto/openapi/event check if changed

## 17. Release Workflow
Jelaskan:
- release candidate
- staging QA
- UAT
- release notes
- merge to main
- production approval
- release tag
- post-release verification

## 18. Hotfix Workflow
Jelaskan:
- kapan hotfix dipakai
- branch dari main
- PR ke main
- production deploy
- release tag
- back-merge ke staging dan develop

## 19. Documentation Workflow
Jelaskan kapan dokumen harus diupdate:
- API berubah
- proto berubah
- event berubah
- data model berubah
- UI flow berubah
- workflow berubah
- local setup berubah
- sprint scope berubah

Buat tabel:
| Change | Document to Update |
|---|---|

## 20. Definition of Ready
Tuliskan syarat issue siap dikerjakan.

## 21. Definition of Done
Tuliskan syarat task dianggap selesai.

## 22. Communication and Handoff
Jelaskan handoff antar:
- Product/Owner
- Backend
- Frontend
- Mobile
- QA
- DevOps
- AI Agent

## 23. Workflow Checklist
Buat checklist ringkas untuk:
- Before Starting Task
- Before Commit
- Before PR
- Before Merge
- Before Staging Release
- Before Production Release

## 24. Example End-to-End Workflow
Berikan contoh workflow:

`Implement Sprint 5 Task 5.5 — Bill Generation with Snapshots`

Jelaskan dari:
- issue creation
- label
- project status
- branch
- AI Agent prompt
- coding
- test
- PR
- review
- QA
- merge

## 25. Final Summary
Ringkas aturan workflow utama.

---

# 6. Writing Rules

- Gunakan Bahasa Indonesia.
- Buat praktis dan bisa dijalankan sebagai SOP.
- Jangan terlalu teoritis.
- Gunakan project `school-platform`.
- Jangan mengubah keputusan workflow yang sudah ada.
- Jangan menulis kode implementasi.
- Sertakan tabel jika membantu.
- Sertakan command contoh untuk workflow lokal dan Git.
- Format siap disimpan sebagai `docs/27-workflow.md`.
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

## Workflow Final Document Status

The final generated document is now available at:

```text
docs/27-workflow.md
```

Use this prompt file only when regenerating or updating the final document.

When making product, planning, workflow, sprint, GitHub issue, PR, QA/UAT, or implementation decisions, read the final document first.
