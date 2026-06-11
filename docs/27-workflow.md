# Workflow — School Platform

Project: `school-platform`  
Document Type: Workflow / Standard Operating Procedure  
Target Audience: Product Owner, Backend Developer, Frontend Developer, Mobile Developer, QA, DevOps, AI Agent  
Status: Draft for MVP Execution  
Repository Target Path: `docs/27-workflow.md`

---

## 1. Purpose

Dokumen ini adalah SOP kerja harian untuk project `school-platform`.

Tujuan dokumen ini adalah memastikan seluruh pekerjaan product, development, QA, DevOps, dan AI Agent berjalan dengan alur yang konsisten, aman, terdokumentasi, dan dapat dilacak.

Dokumen ini mengatur:

```text
- penggunaan GitHub Issues
- penggunaan GitHub Project
- penggunaan Labels dan Milestones
- branch workflow
- Pull Request workflow
- Code Review workflow
- QA/UAT workflow
- AI Agent workflow
- CI/CD workflow
- Release workflow
- Hotfix workflow
- Documentation workflow
```

Workflow utama project:

```text
Planning
→ GitHub Issue
→ Feature Branch
→ Local Development
→ Test
→ Pull Request
→ Review
→ CI
→ Merge to develop
→ Promote to staging
→ QA/UAT
→ Merge to main
→ Production Approval
→ Production Release
```

Prinsip utama: tidak ada perubahan langsung ke branch protected, semua pekerjaan harus melalui issue, PR, review, CI, dan QA sesuai tingkat risikonya.

---

## 2. Workflow Principles

### 2.1 Local-First Development

Semua perubahan dikembangkan dan diuji di local terlebih dahulu sebelum dibuat Pull Request.

Rules:

```text
- Developer checkout dari develop.
- Developer membuat feature/fix/docs/chore branch.
- Developer menjalankan dependency lokal dengan Docker Compose.
- Developer menjalankan test yang relevan sebelum push.
- Developer tidak menggunakan production secret atau production data di local.
```

### 2.2 PR-Based Workflow

Semua perubahan wajib melalui Pull Request.

Tidak boleh:

```text
- direct push ke develop
- direct push ke staging
- direct push ke main
- merge tanpa CI
- merge tanpa review
```

### 2.3 Protected Branch

Branch berikut harus diproteksi:

```text
develop
staging
main
```

Protected branch wajib menggunakan:

```text
- required Pull Request
- required approval
- required status checks
- conversation resolution
- force push disabled
- branch deletion disabled
```

### 2.4 Small Task

Task harus kecil, jelas, dan bisa direview.

Satu issue idealnya menghasilkan satu Pull Request.

Jika task terlalu besar:

```text
- pecah menjadi beberapa issue
- buat dependency antar issue
- hindari PR besar yang sulit direview
```

### 2.5 Service Boundary Discipline

Project menggunakan Go microservices dengan database per service.

Aturan wajib:

```text
- service hanya boleh query database miliknya sendiri
- cross-service communication menggunakan gRPC atau domain events
- API Gateway tidak boleh berisi business logic
- Reporting Service hanya membaca reporting_db/read model
```

### 2.6 Test Before PR

Developer wajib menjalankan test lokal yang relevan sebelum membuat PR.

Minimum:

```text
- go test untuk backend service yang berubah
- lint/typecheck/test/build untuk web jika berubah
- flutter analyze/test untuk mobile jika berubah
- docker compose config jika compose berubah
```

### 2.7 AI Agent as Assistant, Not Final Authority

AI Agent boleh membantu implementasi task kecil, membuat draft kode, membuat test, membuat dokumentasi, dan mempercepat pekerjaan teknis.

AI Agent bukan final authority.

Output AI Agent wajib direview manusia.

### 2.8 Human Review Required

Review manusia wajib untuk:

```text
- authentication
- authorization
- object-level authorization
- finance calculation
- payment verification
- file privacy
- signed URL
- audit log
- event contract
- database migration
- report card publish/lock
- production workflow
```

### 2.9 QA Sign-Off

QA sign-off wajib sebelum production release.

Critical/High bug pada core flow memblokir release.

### 2.10 Production Approval

Production deploy hanya dari `main` dan wajib manual approval melalui GitHub Environment.

### 2.11 Documentation as Source of Truth

Jika implementasi mengubah API, proto, event, data model, workflow, UI flow, local setup, atau sprint scope, dokumen terkait wajib diupdate.

---

## 3. Roles in Workflow

| Role | Responsibility | Main Output |
|---|---|---|
| Product Owner / Reviewer | Menentukan scope, prioritas, acceptance criteria, release readiness, dan UAT direction | PRD, backlog, issue acceptance, release sign-off |
| Backend Developer | Mengembangkan Go services, API Gateway, gRPC, database, event, permission, audit, dan backend tests | Service code, migrations, proto, OpenAPI, event contracts, tests |
| Frontend Developer | Mengembangkan Next.js web admin, UI flow, form validation, API integration, dan frontend tests | Web screens, components, API hooks, tests |
| Mobile Developer | Mengembangkan Flutter mobile app untuk parent/student/guru quick flow | Mobile screens, API client, local secure storage, tests |
| QA | Membuat test case, menjalankan QA/UAT, membuat bug report, regression test, dan sign-off | QA scenarios, bug reports, UAT result, release recommendation |
| Infrastructure/DevOps | Menyiapkan Docker Compose, CI/CD, environments, secrets, observability, backup, restore, deployment | CI workflow, deploy workflow, monitoring, backup/restore docs |
| AI Agent | Membantu task kecil yang jelas scope-nya sesuai AGENTS.md dan SKILLS.md | Draft code, tests, docs, checklist, prompt output |

---

## 4. GitHub Repository Workflow

### 4.1 Branch Types

| Branch | Purpose |
|---|---|
| `develop` | Integrasi development harian |
| `staging` | QA/UAT dan release candidate |
| `main` | Production-ready branch |
| `feature/*` | Fitur baru atau task implementasi |
| `fix/*` | Perbaikan bug non-production |
| `hotfix/*` | Perbaikan urgent dari production/main |
| `docs/*` | Perubahan dokumentasi |
| `chore/*` | Maintenance, dependency, cleanup |
| `refactor/*` | Refactor tanpa perubahan behavior |
| `test/*` | Penambahan/perubahan test |

### 4.2 Merge Rules

| Source Branch | Target Branch | Purpose |
|---|---|---|
| `feature/*` | `develop` | Merge fitur/task selesai |
| `fix/*` | `develop` | Merge bug fix regular |
| `docs/*` | `develop` | Merge dokumentasi |
| `chore/*` | `develop` | Merge maintenance |
| `refactor/*` | `develop` | Merge refactor |
| `test/*` | `develop` | Merge test |
| `develop` | `staging` | Promote release candidate ke staging |
| `staging` | `main` | Promote production release |
| `hotfix/*` | `main` | Urgent production fix |
| `main` | `staging` dan `develop` | Back-merge hotfix |

### 4.3 Daily Branch Workflow

```bash
git checkout develop
git pull origin develop
git checkout -b feature/sprint-5-bill-generation-snapshot
```

Setelah selesai:

```bash
git status
git add .
git commit -m "feat(finance): add bill generation snapshot"
git push origin feature/sprint-5-bill-generation-snapshot
```

Lalu buat Pull Request ke `develop`.

---

## 5. Branch Protection Workflow

### 5.1 develop

Protection untuk `develop`:

```text
- Pull Request required
- minimum 1 approval
- required status checks must pass
- conversation must be resolved
- force push disabled
- branch deletion disabled
```

Merge ke `develop` hanya boleh jika:

```text
- issue jelas
- PR template terisi
- CI pass
- reviewer approve
- scope sesuai issue
```

### 5.2 staging

Protection untuk `staging`:

```text
- Pull Request required
- minimum 1 approval
- CI required
- QA/UAT required for release candidate
- conversation must be resolved
- force push disabled
- branch deletion disabled
```

Merge ke `staging` hanya boleh jika:

```text
- develop sudah stabil
- release candidate siap
- blocking issue diketahui
- QA siap melakukan validasi
```

### 5.3 main

Protection untuk `main`:

```text
- Pull Request required
- approval required
- CI required
- production GitHub Environment approval required
- conversation must be resolved
- force push disabled
- branch deletion disabled
```

Merge ke `main` hanya boleh jika:

```text
- staging QA/UAT pass
- no Critical/High bug pada core flow
- release notes siap
- rollback plan siap
- production approval diberikan
```

---

## 6. GitHub Labels Workflow

Labels digunakan untuk tracking, filtering, prioritas, review, risiko, dan AI Agent workflow.

### 6.1 Type Labels

| Label | Kapan Dipakai |
|---|---|
| `type: feature` | Fitur baru atau implementation task |
| `type: bug` | Bug, regression, unexpected behavior |
| `type: chore` | Maintenance, dependency, housekeeping |
| `type: docs` | Dokumentasi |
| `type: refactor` | Refactor tanpa perubahan behavior |
| `type: test` | Penambahan/perubahan test |
| `type: security` | Security/privacy hardening atau fix |
| `type: infra` | Infra, Docker, CI/CD, deployment |
| `type: spike` | Research, proof of concept, technical exploration |
| `type: hotfix` | Urgent production fix |

### 6.2 Area Labels

| Label | Kapan Dipakai |
|---|---|
| `area: api-gateway` | API Gateway |
| `area: identity` | Identity & Access |
| `area: school-core` | School Core |
| `area: admission` | PPDB |
| `area: academic` | Academic Basic / Report Card |
| `area: finance` | Finance / SPP |
| `area: communication` | Communication / Notification |
| `area: reporting` | Reporting Dashboard |
| `area: web-admin` | Next.js web admin |
| `area: mobile` | Flutter mobile app |
| `area: infra` | Infra/deployment |
| `area: docs` | Documentation |
| `area: security` | Security/privacy |
| `area: observability` | Logs, metrics, monitoring |
| `area: ci-cd` | GitHub Actions, release workflow |
| `area: file-management` | File upload, storage, import |

### 6.3 Sprint Labels

Gunakan label sprint untuk tracking cepat:

```text
sprint: 0
sprint: 1
sprint: 2
sprint: 3
sprint: 4
sprint: 5
sprint: 6
sprint: 7
sprint: 8
sprint: 9
sprint: 10
```

### 6.4 Priority Labels

| Label | Arti |
|---|---|
| `priority: critical` | Blocker, production risk, security/data loss, release blocker |
| `priority: high` | Core MVP atau bug berat |
| `priority: medium` | Penting tetapi tidak memblokir |
| `priority: low` | Minor improvement atau cleanup |

### 6.5 Status Labels

| Label | Arti |
|---|---|
| `status: ready` | Issue siap dikerjakan |
| `status: in-progress` | Sedang dikerjakan |
| `status: blocked` | Terblokir |
| `status: needs-review` | Butuh review |
| `status: needs-qa` | Butuh QA |
| `status: qa-passed` | QA pass |
| `status: done` | Selesai |

### 6.6 AI Agent Labels

| Label | Arti |
|---|---|
| `ai: ready` | Siap dikerjakan/dibantu AI Agent |
| `ai: needs-context` | Butuh konteks tambahan |
| `ai: generated` | Output dibuat/dibantu AI Agent |
| `ai: needs-human-review` | Wajib human review |
| `ai: do-not-use-agent` | AI Agent tidak boleh digunakan |

### 6.7 Risk Labels

| Label | Arti |
|---|---|
| `risk: low` | Risiko rendah |
| `risk: medium` | Risiko sedang |
| `risk: high` | Risiko tinggi |
| `risk: breaking-change` | Ada potensi breaking change |
| `risk: migration` | Menyentuh migration/data |
| `risk: data-sensitive` | Menyentuh Restricted/Confidential data |

### 6.8 Review Labels

| Label | Reviewer Fokus |
|---|---|
| `review: backend` | Backend review |
| `review: frontend` | Frontend review |
| `review: mobile` | Mobile review |
| `review: infra` | Infra/DevOps review |
| `review: qa` | QA review |
| `review: security` | Security/privacy review |
| `review: product` | Product/business review |

---

## 7. GitHub Project Workflow

### 7.1 Project Name

```text
School Platform MVP
```

### 7.2 Project Fields

| Field | Purpose |
|---|---|
| Status | Menunjukkan lifecycle issue |
| Sprint | Sprint/milestone issue |
| Priority | Urgensi pekerjaan |
| Area | Modul/platform terdampak |
| Type | Jenis pekerjaan |
| Owner | Penanggung jawab |
| Estimate | Estimasi relatif |
| Risk | Level risiko |
| Platform | Backend, Web, Mobile, Infra, Docs, QA, Product |
| AI Agent | Status penggunaan AI Agent |
| Target Release | MVP atau release tertentu |

### 7.3 Project Views

| View | Fungsi |
|---|---|
| MVP Board | Board utama semua pekerjaan MVP |
| Sprint Board | Board sprint aktif |
| Backlog Table | Grooming backlog dan prioritas |
| By Area | Melihat pekerjaan berdasarkan modul/platform |
| By Priority | Fokus pada Critical/High |
| QA/UAT | Melihat issue yang butuh QA/UAT |
| AI Agent Tasks | Melihat task yang siap dibantu AI Agent |
| Release Readiness | Melacak kesiapan staging/production release |

### 7.4 Status Update Rule

Setiap issue harus diupdate statusnya saat berpindah tahap.

Contoh:

```text
Ready → In Progress saat developer mulai kerja
In Progress → In Review saat PR dibuat
In Review → QA saat PR merged dan butuh QA
QA → Done saat QA pass
Any status → Blocked saat ada blocker
```

---

## 8. Milestone Workflow

Gunakan milestone untuk mengikat issue dengan sprint.

| Milestone | Fokus |
|---|---|
| Sprint 0 — Project Foundation | Repo, CI, Docker, service skeleton |
| Sprint 1 — Identity & Access | Auth, JWT, RBAC, ABAC/scope |
| Sprint 2 — School Core | Master data yayasan/sekolah/siswa/guru/kelas |
| Sprint 3 — File Management + Import Excel | Private file dan import data awal |
| Sprint 4 — PPDB | Admission workflow |
| Sprint 5 — Finance / SPP | Tagihan dan pembayaran manual |
| Sprint 6 — Academic Dasar | Jadwal dan absensi |
| Sprint 7 — Report Card / E-Rapor Dasar | Nilai dan rapor |
| Sprint 8 — Communication / Notification | Announcement dan notification |
| Sprint 9 — Reporting Dashboard | Projection dan dashboard |
| Sprint 10 — Security, Observability, Backup, UAT Hardening | Hardening dan readiness |
| MVP Release | Final release readiness |

Rules:

```text
- Setiap issue sprint wajib punya milestone.
- PR harus menutup/mengacu pada issue di milestone yang sama.
- Milestone ditutup setelah issue blocking selesai.
```

---

## 9. Issue Workflow

### 9.1 Lifecycle

```text
Backlog
→ Ready
→ In Progress
→ In Review
→ QA
→ Done
```

Exception:

```text
Blocked
```

### 9.2 Status Rules

| Status | Arti | Siapa Mengubah | Syarat Pindah Status |
|---|---|---|---|
| Backlog | Candidate work, belum siap | Product Owner/Reviewer | Scope dan acceptance criteria dilengkapi |
| Ready | Siap dikerjakan | Product Owner/Reviewer/Lead | Owner mulai kerja |
| In Progress | Sedang dikerjakan | Developer/AI Agent operator | PR dibuat |
| In Review | PR sedang direview | Developer/Reviewer | PR approved dan merged atau perlu revisi |
| QA | Butuh QA/UAT | Developer/QA | QA pass atau bug dibuat |
| Done | Selesai | QA/Product Owner/Lead | DoD terpenuhi |
| Blocked | Terblokir | Siapa pun yang menemukan blocker | Blocker terselesaikan |

### 9.3 Blocked Issue Rule

Jika issue blocked:

```text
- tambahkan label status: blocked
- jelaskan blocker di komentar
- sebutkan dependency
- tag owner/reviewer yang dapat membantu
- update Project Status = Blocked
```

---

## 10. Task Creation Workflow

Cara membuat task dari sprint prompt:

1. Baca dokumen sprint terkait.
2. Pilih task kecil yang jelas.
3. Buat issue menggunakan template yang tepat.
4. Tulis objective.
5. Tulis scope.
6. Tulis out of scope.
7. Tulis acceptance criteria.
8. Beri label type, area, sprint, priority, risk, review.
9. Pilih milestone sprint.
10. Masukkan ke GitHub Project `School Platform MVP`.
11. Assign owner.
12. Tentukan AI Agent status.

### 10.1 Issue Template Selection

| Kondisi | Template |
|---|---|
| Feature/task biasa | `.github/ISSUE_TEMPLATE/feature_task.yml` |
| Bug/regression | `.github/ISSUE_TEMPLATE/bug_report.yml` |
| Task untuk AI Agent | `.github/ISSUE_TEMPLATE/ai_agent_task.yml` |
| Security/privacy review | `.github/ISSUE_TEMPLATE/security_review.yml` |
| QA/UAT | `.github/ISSUE_TEMPLATE/qa_uat.yml` |

---

## 11. AI Agent Workflow

### 11.1 AI Agent Execution Steps

1. Pilih issue dengan label `ai: ready`.
2. Baca `AGENTS.md`.
3. Baca `SKILLS.md`.
4. Baca `docs/README.md`.
5. Baca `docs/09-ai-agent-rules.md`.
6. Baca `docs/08-coding-standard.md`.
7. Baca sprint plan jika tersedia.
8. Baca sprint task prompt terkait.
9. Copy task prompt ke AI Agent.
10. Tambahkan konteks file/dokumen.
11. Minta AI Agent implement.
12. Review hasil AI Agent.
13. Jalankan test.
14. Commit ke feature branch.
15. Buat PR.
16. Human review wajib.

### 11.2 Required AI Agent Prompt Context

Minimum context:

```text
Project: school-platform
Branch/task:
Issue:
Required docs:
- AGENTS.md
- SKILLS.md
- docs/README.md
- docs/09-ai-agent-rules.md
- docs/08-coding-standard.md
- active sprint plan if available
- active sprint task prompt
Scope:
Out of scope:
Acceptance criteria:
Expected output:
```

### 11.3 AI Agent Must Not Work On

AI Agent tidak boleh mengerjakan:

```text
- production secrets
- final security approval
- legal/compliance decision
- production deployment approval
- akses data asli sensitif
- perubahan arsitektur besar tanpa instruksi
- final decision untuk data privacy
- hotfix production tanpa human review
```

### 11.4 AI Agent Output Review

Reviewer harus memeriksa:

```text
- apakah scope sesuai issue
- apakah tidak ada fitur tambahan
- apakah service boundary aman
- apakah tidak ada cross-service DB query
- apakah permission/scope diterapkan
- apakah object-level authorization ada
- apakah audit/event/file/privacy dicek
- apakah test cukup
- apakah dokumentasi diupdate jika perlu
```

---

## 12. Pull Request Workflow

### 12.1 Kapan Membuat PR

PR dibuat setelah:

```text
- task scope selesai
- test lokal relevan pass
- tidak ada perubahan out of scope
- dokumentasi diupdate jika perlu
- commit sudah rapi
```

### 12.2 PR Title Convention

Format:

```text
type(scope): short description
```

Examples:

```text
feat(finance): add bill generation snapshot
fix(identity): reject reused refresh token
docs(workflow): add daily development SOP
chore(ci): add repository validation workflow
test(academic): add attendance scope tests
```

### 12.3 PR Description

PR wajib menggunakan:

```text
.github/pull_request_template.md
```

Isi minimal:

```text
- summary
- related issue
- type of change
- affected area
- scope
- out of scope
- tests
- security/permission checklist
- documentation checklist
- rollback plan jika relevan
```

### 12.4 Reviewer

Reviewer dipilih berdasarkan area:

| Area | Reviewer |
|---|---|
| Backend service | Backend reviewer |
| Web admin | Frontend reviewer |
| Mobile app | Mobile reviewer |
| Infrastructure/CI/CD | DevOps reviewer |
| Permission/security/privacy | Security reviewer |
| QA/UAT | QA reviewer |
| Product flow | Product reviewer |

### 12.5 CI

PR tidak boleh merge jika CI gagal.

### 12.6 Comment Resolution

Semua review comment harus:

```text
- dijawab
- diperbaiki
- atau diberi alasan teknis
```

Conversation harus resolved sebelum merge.

### 12.7 Merge Strategy

Recommended:

```text
Squash merge untuk feature/fix kecil
Merge commit untuk release branch jika diperlukan
```

Commit message hasil squash harus jelas.

---

## 13. Code Review Workflow

### 13.1 Backend Review

Checklist:

```text
- service boundary aman
- tidak ada cross-service DB query
- gRPC/event digunakan sesuai kebutuhan
- handler/service/repository terpisah
- migration aman
- sqlc query aman
- decimal digunakan untuk finance
- permission/scope diterapkan
- object-level authorization ada
- audit log untuk aksi sensitif
- tests cukup
```

### 13.2 Frontend Review

Checklist:

```text
- UI sesuai flow
- labels Bahasa Indonesia
- form validation ada
- API error handling jelas
- permission-based UI diterapkan
- loading/empty/error state ada
- tidak menyimpan token secara tidak aman
- tests jika relevan
```

### 13.3 Mobile Review

Checklist:

```text
- secure storage digunakan untuk token
- API error handling jelas
- UI mobile-friendly
- role/scope behavior benar
- no offline write MVP
- tests jika relevan
```

### 13.4 Infra Review

Checklist:

```text
- CI tidak expose secrets
- Docker Compose valid
- env separation jelas
- deployment aman
- rollback plan tersedia
- observability/backup tidak melemah
```

### 13.5 Security Review

Checklist:

```text
- authentication aman
- authorization benar
- object-level authorization ada
- field-level access jika perlu
- no sensitive logs
- file private by default
- signed URL aman
- audit log tersedia
- secrets tidak masuk repo
```

### 13.6 QA Review

Checklist:

```text
- acceptance criteria dapat diuji
- test scenario jelas
- regression impact diketahui
- bug severity tepat
- release blocker ditandai
```

### 13.7 General Review Checklist

```text
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
```

---

## 14. QA Workflow

### 14.1 Test Case Creation

QA membuat test case dari:

```text
- PRD
- sprint plan
- issue acceptance criteria
- UI flow docs
- API contract
- event contract
- known risks
```

### 14.2 Test Execution

QA menjalankan test di staging atau environment yang disepakati.

Test execution mencakup:

```text
- positive flow
- negative flow
- validation
- permission/scope
- object-level access
- regression
- audit/event/file check jika relevan
```

### 14.3 Bug Report

Bug dibuat menggunakan:

```text
.github/ISSUE_TEMPLATE/bug_report.yml
```

Bug wajib mencakup:

```text
- environment
- steps to reproduce
- expected result
- actual result
- evidence
- severity
- affected area
```

### 14.4 Bug Severity

| Severity | Meaning | Release Impact |
|---|---|---|
| Critical | Sistem tidak bisa digunakan, data loss, security breach, core flow mati | Block production |
| High | Core flow terganggu atau data penting salah | Block production jika core flow |
| Medium | Gangguan penting tetapi ada workaround | Tidak selalu block |
| Low | Minor UI/cosmetic/edge case | Tidak block |

### 14.5 Regression Test

Regression test wajib dilakukan sebelum staging → main.

Minimal regression:

```text
- login
- core data access
- PPDB core flow
- Finance bill/payment flow
- Academic attendance/score flow
- Report publish/view flow
- Notification flow
- Dashboard scope
```

### 14.6 UAT

UAT dilakukan oleh Product Owner/Reviewer dan user perwakilan jika tersedia.

UAT harus menggunakan issue template:

```text
.github/ISSUE_TEMPLATE/qa_uat.yml
```

### 14.7 QA Sign-Off

QA sign-off diperlukan sebelum production release.

Rule:

```text
Critical/High bug pada core flow memblokir production release.
```

---

## 15. Development Local Workflow

### 15.1 Step-by-Step

1. Clone repo.
2. Checkout `develop`.
3. Create feature branch.
4. Copy `.env.example`.
5. Run Docker Compose.
6. Run migration.
7. Run `sqlc generate` jika query berubah.
8. Run service.
9. Run tests.
10. Commit.
11. Push.
12. Create PR.

### 15.2 Command Example

```bash
git clone git@github.com:<org-or-user>/school-platform.git
cd school-platform

git checkout develop
git pull origin develop

git checkout -b feature/sprint-5-bill-generation-snapshot
```

Copy env:

```bash
cp .env.example .env
```

Run dependencies:

```bash
docker compose up -d
docker compose ps
docker compose logs -f
```

Run migration example:

```bash
cd services/finance-service
goose -dir migrations postgres "$FINANCE_DATABASE_URL" up
```

Run sqlc if needed:

```bash
sqlc generate
```

Run service:

```bash
go run ./cmd/server
```

Run tests:

```bash
go test ./...
go vet ./...
gofmt -w .
```

Web:

```bash
cd apps/web-admin
npm install
npm run lint
npm run test
npm run build
```

Mobile:

```bash
cd apps/mobile-app
flutter pub get
flutter analyze
flutter test
```

Commit and push:

```bash
git status
git add .
git commit -m "feat(finance): add bill generation snapshot"
git push origin feature/sprint-5-bill-generation-snapshot
```

---

## 16. CI/CD Workflow

### 16.1 CI on PR

CI berjalan saat PR menuju:

```text
develop
staging
main
```

CI checks:

```text
- repository check
- no .env/secrets committed
- YAML validation
- Go fmt/vet/test
- Web lint/typecheck/test/build
- Flutter analyze/test
- Docker Compose config
```

Web dan mobile checks harus mendeteksi folder app dari repository root dan skip dengan bersih jika app belum dibuat.

### 16.2 CI on develop/staging/main

CI juga berjalan saat push/merge ke:

```text
develop
staging
main
```

### 16.3 Staging Deploy

Staging deploy dilakukan dari branch:

```text
staging
```

Requirements:

```text
- CI pass
- PR develop → staging approved
- release candidate notes
```

### 16.4 Production Deploy

Production deploy dilakukan dari branch:

```text
main
```

Requirements:

```text
- PR staging → main approved
- CI pass
- QA sign-off
- no Critical/High core bug
- release notes
- rollback plan
- GitHub Environment manual approval
```

### 16.5 Manual Approval

Production environment harus membutuhkan reviewer approval.

### 16.6 Rollback

Rollback dapat dilakukan dengan:

```text
- redeploy previous image/tag
- disable feature via config jika tersedia
- run rollback migration jika aman
- restore backup jika terjadi data corruption
```

### 16.7 Quality Gate

| Gate | Kapan |
|---|---|
| lint | PR dan branch protected |
| test | PR dan branch protected |
| build | PR dan branch protected |
| docker build/config | Jika Docker berubah |
| migration check | Jika migration berubah |
| proto check | Jika protobuf berubah |
| openapi check | Jika API contract berubah |
| event check | Jika event contract berubah |

---

## 17. Release Workflow

### 17.1 Release Candidate

Release candidate dibuat dari:

```text
develop → staging
```

### 17.2 Staging QA

QA menjalankan:

```text
- smoke test
- regression test
- module test
- permission/scope test
- UAT scenario
```

### 17.3 UAT

UAT dilakukan untuk alur utama MVP.

Core flow:

```text
- login
- setup school core
- import data
- PPDB conversion
- bill generation
- payment verification
- attendance/score/report card
- announcement/notification
- dashboard
```

### 17.4 Release Notes

Release notes harus mencakup:

```text
- changes
- bug fixes
- migration notes
- known issues
- rollback plan
```

### 17.5 Merge to Main

Jika QA/UAT pass:

```text
staging → main
```

### 17.6 Production Approval

Production deploy membutuhkan manual approval.

### 17.7 Release Tag

Setelah deploy production:

```bash
git tag v0.1.0
git push origin v0.1.0
```

### 17.8 Post-Release Verification

Verifikasi setelah release:

```text
- app reachable
- login works
- core APIs healthy
- workers running
- RabbitMQ queues normal
- DB connected
- file access works
- logs/metrics visible
```

---

## 18. Hotfix Workflow

### 18.1 Kapan Hotfix Dipakai

Hotfix hanya untuk:

```text
- production outage
- critical auth/security issue
- data corruption risk
- finance critical bug
- core flow production blocker
```

### 18.2 Branch dari Main

```bash
git checkout main
git pull origin main
git checkout -b hotfix/fix-payment-verification
```

### 18.3 PR ke Main

Hotfix PR target:

```text
main
```

PR wajib mencakup:

```text
- root cause
- fix summary
- test evidence
- risk
- rollback
```

### 18.4 Production Deploy

Deploy production setelah:

```text
- CI pass
- urgent review approved
- production approval
```

### 18.5 Release Tag

```bash
git tag v0.1.1
git push origin v0.1.1
```

### 18.6 Back-Merge

Setelah hotfix masuk main:

```bash
git checkout staging
git pull origin staging
git merge main
git push origin staging

git checkout develop
git pull origin develop
git merge main
git push origin develop
```

Atau gunakan PR:

```text
main → staging
main → develop
```

---

## 19. Documentation Workflow

Dokumentasi wajib diupdate jika ada perubahan kontrak, arsitektur, workflow, data model, atau scope.

| Change | Document to Update |
|---|---|
| Arsitektur berubah | `docs/01-technical-architecture.md` |
| Service boundary berubah | `docs/02-service-boundary.md` |
| Data model berubah | `docs/03-data-model-mvp.md` |
| REST API berubah | `docs/04-api-contract.md` |
| gRPC/proto berubah | `docs/04-api-contract.md`, `packages/proto/` |
| Event berubah | `docs/05-event-contract.md`, `packages/events/` |
| UI/user flow berubah | `docs/06-ui-screen-user-flow.md` |
| Acceptance criteria/test berubah | `docs/07-test-plan-acceptance-criteria.md` |
| Coding standard berubah | `docs/08-coding-standard.md` |
| AI Agent rule berubah | `docs/09-ai-agent-rules.md`, `AGENTS.md`, `SKILLS.md` |
| Sprint scope berubah | `docs/10-sprint-backlog-mvp.md`, sprint plan terkait |
| GitHub workflow berubah | `docs/11-github-repository-rules.md`, `docs/25-github-project-management.md`, `docs/27-workflow.md` |
| Local setup berubah | `docs/24-local-development-guide.md` |
| PR/Issue template berubah | `docs/25-github-project-management.md`, `docs/27-workflow.md` |
| CI/CD berubah | `docs/11-github-repository-rules.md`, `docs/25-github-project-management.md`, `docs/27-workflow.md` |
| PRD berubah | `docs/25-product-requirement-document.md` |
| Development Plan berubah | `docs/26-development-plan.md` |

Rules:

```text
- Dokumentasi diupdate di PR yang sama jika perubahan terkait langsung.
- Jika update dokumentasi besar, boleh dibuat PR terpisah dengan issue terkait.
- AI Agent boleh membantu draft dokumentasi, tetapi human review tetap wajib.
```

---

## 20. Definition of Ready

Issue siap dikerjakan jika:

```text
- objective jelas
- scope jelas
- out of scope jelas
- acceptance criteria jelas
- role/user impact jelas
- area ditentukan
- priority ditentukan
- risk ditentukan
- milestone ditentukan
- dependency diketahui
- owner ditentukan atau siap diassign
- AI Agent status jelas
- required docs/prompts diketahui
```

Checklist:

```text
- [ ] Objective clear
- [ ] Scope clear
- [ ] Out of scope clear
- [ ] Acceptance criteria defined
- [ ] Labels assigned
- [ ] Milestone assigned
- [ ] GitHub Project fields filled
- [ ] Dependencies known
- [ ] AI Agent status set
```

---

## 21. Definition of Done

Task dianggap selesai jika:

```text
- scope issue selesai
- tidak ada pekerjaan out of scope
- tests ditambahkan/diupdate
- tests pass
- lint/format pass
- CI pass
- permission/scope dicek jika relevan
- object-level authorization dicek jika relevan
- audit log ditambahkan jika relevan
- event contract diupdate jika relevan
- dokumentasi diupdate jika relevan
- PR approved
- QA pass jika diperlukan
- issue ditutup atau dipindah Done
```

Checklist:

```text
- [ ] Implementation matches issue scope
- [ ] Tests pass
- [ ] CI pass
- [ ] Review approved
- [ ] Docs updated if needed
- [ ] QA passed if needed
- [ ] No Critical/High bug remains for the flow
```

---

## 22. Communication and Handoff

### 22.1 Product Owner / Reviewer → Developer

Handoff harus mencakup:

```text
- objective
- scope
- out of scope
- acceptance criteria
- priority
- risk
- UI/business notes
```

### 22.2 Backend → Frontend

Handoff harus mencakup:

```text
- endpoint
- request/response
- error response
- permission behavior
- status enum
- sample payload
```

### 22.3 Backend → Mobile

Handoff harus mencakup:

```text
- API contract
- auth behavior
- token refresh behavior
- file upload/download rule
- mobile-specific constraints
```

### 22.4 Backend → QA

Handoff harus mencakup:

```text
- feature summary
- test data
- edge cases
- permission/scope cases
- audit/event expected behavior
```

### 22.5 DevOps → Developer

Handoff harus mencakup:

```text
- env vars
- service port
- Docker Compose changes
- migration command
- health check
- deploy notes
```

### 22.6 AI Agent → Human Reviewer

AI Agent output harus mencakup:

```text
- summary
- changed files
- tests run
- notes/risks
- assumptions
- missing context if any
```

---

## 23. Workflow Checklist

### 23.1 Before Starting Task

```text
- [ ] Issue exists
- [ ] Issue is Ready
- [ ] Scope is clear
- [ ] Acceptance criteria are clear
- [ ] Milestone and labels are set
- [ ] Branch target is clear
- [ ] Required docs are read
- [ ] AI Agent status is clear
```

### 23.2 Before Commit

```text
- [ ] Code matches scope
- [ ] No unrelated changes
- [ ] No secrets
- [ ] No .env
- [ ] Formatting applied
- [ ] Tests run locally
```

### 23.3 Before PR

```text
- [ ] Branch up to date
- [ ] PR template filled
- [ ] Related issue linked
- [ ] Tests listed
- [ ] Risk documented
- [ ] Docs updated if needed
```

### 23.4 Before Merge

```text
- [ ] CI pass
- [ ] Review approved
- [ ] Conversations resolved
- [ ] No Critical issue introduced
- [ ] QA required? If yes, moved to QA
```

### 23.5 Before Staging Release

```text
- [ ] develop stable
- [ ] CI pass
- [ ] Release candidate notes ready
- [ ] Known issues documented
- [ ] QA plan ready
```

### 23.6 Before Production Release

```text
- [ ] QA/UAT pass
- [ ] No Critical/High core bug open
- [ ] Security review done if needed
- [ ] Backup done if risky release
- [ ] Rollback plan ready
- [ ] Release notes ready
- [ ] Production approval obtained
```

---

## 24. Example End-to-End Workflow

Example:

```text
Implement Sprint 5 Task 5.5 — Bill Generation with Snapshots
```

### 24.1 Issue Creation

Create issue using:

```text
.github/ISSUE_TEMPLATE/feature_task.yml
```

Issue title:

```text
Sprint 5 Task 5.5 — Bill Generation with Snapshots
```

Labels:

```text
type: feature
area: finance
sprint: 5
priority: critical
risk: data-sensitive
risk: migration
review: backend
review: security
ai: ready
```

Milestone:

```text
Sprint 5 — Finance / SPP
```

Project fields:

```text
Status: Ready
Sprint: Sprint 5
Priority: Critical
Area: finance
Type: feature
Platform: Backend
Risk: Data Sensitive
AI Agent: Ready
Target Release: MVP
```

### 24.2 Objective

```text
Generate monthly/periodic SPP bills for students using fee scheme and fee policy, and store immutable bill snapshot.
```

### 24.3 Scope

```text
- bill generation service
- idempotency check
- bill_items snapshot
- fee policy snapshot
- decimal/NUMERIC amount
- audit log
- finance event published
- unit/integration tests
```

### 24.4 Out of Scope

```text
- payment gateway
- bank reconciliation
- full accounting ledger
- WhatsApp notification
```

### 24.5 Branch

```bash
git checkout develop
git pull origin develop
git checkout -b feature/sprint-5-bill-generation-snapshot
```

### 24.6 AI Agent Prompt

Use issue prompt with context:

```text
Read:
- AGENTS.md
- SKILLS.md
- docs/README.md
- docs/09-ai-agent-rules.md
- docs/08-coding-standard.md
- docs/18-sprint-5-task-prompts.md
- docs/26-development-plan.md
- docs/27-workflow.md

Task:
Implement only Sprint 5 Task 5.5 — Bill Generation with Snapshots.

Rules:
- no float for money
- use decimal/NUMERIC
- bill generation must be idempotent
- bill must store snapshot
- no cross-service DB query
- add audit log
- add tests
```

### 24.7 Coding

Developer or AI Agent-assisted developer implements:

```text
- migrations
- sqlc queries
- service method
- handler
- tests
- event publisher if required
```

### 24.8 Test

Run:

```bash
cd services/finance-service
gofmt -w .
go vet ./...
go test ./... -count=1
```

If API/web affected:

```bash
cd apps/web-admin
npm run lint
npm run test
npm run build
```

### 24.9 Pull Request

PR title:

```text
feat(finance): add bill generation with snapshots
```

PR target:

```text
develop
```

PR description must include:

```text
Closes #<issue-id>
Summary
Scope
Out of scope
Tests
Security/permission checklist
Finance checklist
Rollback plan
```

### 24.10 Review

Review focus:

```text
- decimal/NUMERIC
- idempotency
- snapshot correctness
- no cross-service DB query
- audit log
- permission/scope
- test coverage
```

### 24.11 QA

Move issue to QA if UI/API flow needs validation.

QA validates:

```text
- bill generated once
- duplicate generation rejected/skipped
- amount snapshot unchanged after policy update
- only authorized Bendahara can generate
- audit log exists
```

### 24.12 Merge

If approved and CI pass:

```text
feature/* → develop
```

If included in release candidate:

```text
develop → staging
```

After QA/UAT:

```text
staging → main
```

---

## 25. Final Summary

Workflow utama `school-platform`:

```text
Issue
→ Branch
→ Local Development
→ Test
→ Pull Request
→ CI
→ Human Review
→ Merge to develop
→ Staging
→ QA/UAT
→ Main
→ Production Approval
→ Release
```

Aturan utama:

```text
- Semua pekerjaan harus punya issue.
- Semua perubahan harus lewat Pull Request.
- develop, staging, dan main harus protected.
- CI wajib pass sebelum merge.
- Review manusia wajib.
- QA sign-off wajib sebelum production.
- Production deploy hanya dari main.
- Production deploy wajib manual approval.
- Hotfix berasal dari main dan wajib back-merge.
- AI Agent hanya assistant, bukan final authority.
- AI Agent tidak boleh menangani secrets atau production approval.
- Tidak boleh cross-service database query.
- API Gateway tidak boleh berisi business logic.
- Dokumentasi harus diupdate jika kontrak/workflow/scope berubah.

## Sprint Plan and GitHub Setup References

For sprint execution, use the active sprint plan:

```text
docs/29-sprint-0-plan.md through docs/39-sprint-10-plan.md
```

For practical repository setup, GitHub Labels, GitHub Milestones, GitHub Project fields/views, branch protection, and GitHub Environments, use:

```text
docs/40-github-repository-setup-labels-project.md
```

Workflow rule:

```text
Sprint Plan
→ GitHub Issue
→ Feature Branch
→ Local Development
→ Pull Request
→ Review
→ CI
→ QA/UAT
→ Release
```

## 26. Git Commit Convention

Project `school-platform` menggunakan format commit berbasis Conventional Commit agar riwayat perubahan mudah dibaca, mudah ditelusuri ke GitHub Issue/PR, dan dapat mendukung release notes.

### Format Commit

Gunakan format:

```text
type(scope): short description
```

Contoh:

```text
feat(identity): add refresh token rotation
fix(finance): prevent duplicate bill generation
docs(workflow): add git commit convention
chore(ci): add repository validation workflow
test(academic): add attendance scope tests
refactor(api-gateway): simplify request context middleware
```

### Commit Type

Gunakan `type` berikut:

| Type | Kapan Dipakai |
|---|---|
| `feat` | Penambahan fitur baru |
| `fix` | Perbaikan bug |
| `docs` | Perubahan dokumentasi |
| `chore` | Maintenance, dependency, housekeeping |
| `refactor` | Refactor tanpa perubahan behavior |
| `test` | Penambahan/perubahan test |
| `ci` | Perubahan GitHub Actions atau CI/CD |
| `build` | Build system, Docker, dependency packaging |
| `perf` | Peningkatan performa |
| `security` | Security/privacy hardening atau fix |
| `revert` | Revert commit sebelumnya |

### Commit Scope

Gunakan `scope` sesuai area project.

Recommended scopes:

```text
api-gateway
identity
school-core
admission
academic
finance
communication
reporting
web-admin
mobile
infra
ci
docs
security
observability
file-management
shared-go
proto
openapi
events
```

Contoh:

```text
feat(school-core): add student class assignment
fix(admission): prevent duplicate applicant conversion
security(file-management): enforce signed url authorization
ci(repository): block committed env files
docs(github): add label setup guide
```

### Subject / Short Description

Rules:

```text
- Gunakan Bahasa Inggris.
- Gunakan lowercase setelah titik dua.
- Gunakan kalimat pendek dan jelas.
- Gunakan imperative mood jika memungkinkan.
- Maksimal disarankan 72 karakter.
- Jangan akhiri dengan titik.
```

Contoh baik:

```text
feat(finance): add bill generation snapshot
```

Contoh kurang baik:

```text
feat(finance): Added the billing feature.
update file
fix bug
```

### Commit Body

Gunakan commit body jika perubahan butuh penjelasan tambahan.

Format:

```text
type(scope): short description

Longer explanation of what changed and why.
Mention constraints, risk, or migration notes if needed.
```

Contoh:

```text
feat(finance): add bill generation snapshot

Store applied fee policy snapshot on bill items to prevent
future fee policy changes from changing historical bills.
```

### Issue Reference

Jika commit terkait issue tertentu, tambahkan reference di body atau footer.

```text
Refs #123
```

Untuk menutup issue, gunakan di PR description, bukan wajib di commit:

```text
Closes #123
```

### Breaking Change

Jika ada breaking change, gunakan footer:

```text
BREAKING CHANGE: explain what changed and migration required
```

Contoh:

```text
refactor(identity): change actor context payload

BREAKING CHANGE: API Gateway must forward actor context using X-Actor-Context header.
```

Breaking change harus diberi label:

```text
risk: breaking-change
```

dan wajib human review.

### Migration Commit

Jika commit menambahkan migration, gunakan scope service terkait.

Contoh:

```text
feat(identity): add users and sessions migrations
feat(finance): add student bills migrations
```

Migration commit wajib mencantumkan risiko jika menyentuh data existing.

### Security Commit

Untuk perubahan security/privacy:

```text
security(identity): hash refresh tokens at rest
security(file-management): restrict signed url access by scope
```

Security commit wajib review dengan label:

```text
review: security
```

### Commit Granularity

Rules:

```text
- Satu commit sebaiknya merepresentasikan satu perubahan logis.
- Jangan campur refactor besar dengan feature.
- Jangan campur formatting massal dengan business logic.
- Jangan commit file hasil build/cache kecuali memang diperlukan.
- Jangan commit .env, private key, secret, dump database, atau file sensitif.
```

### Commit Before PR

Sebelum push:

```bash
git status
git diff
git add .
git commit -m "feat(scope): short description"
git push origin feature/branch-name
```

Jika perlu body:

```bash
git commit
```

lalu tulis:

```text
feat(scope): short description

Detailed explanation.
```

### Relationship with PR Title

PR title sebaiknya mengikuti format yang sama:

```text
type(scope): short description
```

Jika branch memiliki banyak commit kecil, PR tetap harus memiliki title yang mewakili hasil akhir.

### Squash Merge

Jika menggunakan squash merge, pastikan squash commit mengikuti format:

```text
type(scope): short description
```

Contoh:

```text
feat(identity): implement login and refresh token flow
```

### Commit Review Checklist

Sebelum push/PR:

```text
- [ ] Commit message mengikuti format `type(scope): description`
- [ ] Scope sesuai area yang berubah
- [ ] Subject jelas dan singkat
- [ ] Tidak ada secret atau .env
- [ ] Tidak ada file build/cache yang tidak perlu
- [ ] Commit tidak mencampur perubahan tidak terkait
- [ ] Breaking change diberi catatan jika ada
- [ ] Migration/security change diberi konteks jika perlu
```

## GitHub Setup Script Workflow

When initializing or updating the repository setup, use:

```text
docs/40-github-repository-setup-labels-project.md
scripts/github/
```

Recommended workflow:

```text
Review docs/40
→ Set environment variables
→ Run script step by step
→ Verify repository settings
→ Commit generated support files
→ Open PR
→ Human review
```

Recommended command sequence:

```bash
export GITHUB_OWNER="<org-or-user>"
export REPO_NAME="school-platform"
export PRODUCTION_REVIEWER_USER="kuswandi-ti"

bash scripts/github/01-create-repository.sh
bash scripts/github/02-bootstrap-repository-files.sh
bash scripts/github/03-create-branches.sh
bash scripts/github/04-setup-branch-protection.sh
bash scripts/github/05-setup-environments.sh
bash scripts/github/06-setup-labels.sh
bash scripts/github/07-setup-milestones.sh
bash scripts/github/08-setup-project.sh
PROJECT_NUMBER="<project-number>" bash scripts/github/09-setup-project-fields.sh
bash scripts/github/10-project-views-manual-guide.sh
```

Do not run setup scripts against production repository settings without human confirmation.
