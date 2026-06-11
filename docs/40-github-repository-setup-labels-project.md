# GitHub Repository Setup, Labels, and Project — School Platform

Project: `school-platform`  
Document Type: GitHub Repository Setup & Project Management Guide  
Target Audience: Product Owner, Developer, QA, DevOps, AI Agent  
Repository Target Path: `docs/40-github-repository-setup-labels-project.md`

---

## 1. Purpose

Dokumen ini menjelaskan langkah praktis untuk menyiapkan GitHub repository, GitHub Labels, Milestones, GitHub Project, Issue workflow, dan Pull Request workflow untuk project `school-platform`.

Dokumen ini merupakan panduan operasional yang merangkum pembahasan sebelumnya terkait:

```text
- setup GitHub repository
- branch workflow
- branch protection
- repository support files
- GitHub Labels
- GitHub Milestones
- GitHub Project fields
- GitHub Project views
- issue lifecycle
- Pull Request workflow
- QA/UAT tracking
- AI Agent task tracking
- release readiness tracking
```

Dokumen ini harus selaras dengan:

```text
AGENTS.md
SKILLS.md
docs/11-github-repository-rules.md
docs/25-github-project-management.md
docs/26-development-plan.md
docs/27-workflow.md
```

---

## 2. Target GitHub Repository

Recommended repository name:

```text
school-platform
```

Repository type:

```text
Private repository
```

Alasan private:

```text
- project internal yayasan
- mengandung domain pendidikan
- akan memiliki struktur data anak/orang tua/guru/keuangan
- perlu kontrol akses developer dan reviewer
```

Recommended repository description:

```text
Internal multi-unit school foundation platform for TK, SD, SMP, and SMA operations.
```

---

## 3. Repository Setup Checklist

Bagian ini menjabarkan checklist setup repository menjadi langkah operasional yang bisa dijalankan manual atau menggunakan script di folder:

```text
scripts/github/
```

> Catatan penting:
>
> - Script menggunakan GitHub CLI (`gh`) dan Git.
> - Jalankan dari mesin lokal yang sudah login ke GitHub CLI.
> - Untuk branch protection dan environments, akun GitHub harus memiliki permission admin pada repository.
> - Untuk GitHub Project v2, beberapa konfigurasi view mungkin tetap perlu diselesaikan manual melalui UI GitHub karena dukungan CLI/API dapat berbeda tergantung akun/organisasi.

---

### 3.1 Prerequisites

Install dan login GitHub CLI:

```bash
gh auth login
gh auth status
```

Pastikan Git tersedia:

```bash
git --version
```

Set variable project:

```bash
export GITHUB_OWNER="<org-or-user>"
export REPO_NAME="school-platform"
export DEFAULT_BRANCH="main"
```

Contoh untuk user personal:

```bash
export GITHUB_OWNER="kuswandi-ti"
export REPO_NAME="school-platform"
export DEFAULT_BRANCH="main"
```

Jika menggunakan organization:

```bash
export GITHUB_OWNER="nama-organisasi"
export REPO_NAME="school-platform"
export DEFAULT_BRANCH="main"
```

---

### 3.2 Buat repository `school-platform`

Checklist:

```text
- [ ] Buat repository `school-platform`
```

Manual via GitHub UI:

```text
GitHub → New Repository → Repository name: school-platform
```

Via GitHub CLI:

```bash
gh repo create "$GITHUB_OWNER/$REPO_NAME" \
  --private \
  --description "Internal multi-unit school foundation platform for TK, SD, SMP, and SMA operations." \
  --add-readme
```

Script:

```bash
bash scripts/github/01-create-repository.sh
```

---

### 3.3 Set repository sebagai private

Checklist:

```text
- [ ] Set repository sebagai private
```

Jika repository dibuat menggunakan command di atas, repository sudah private.

Jika repository sudah terlanjur public:

```bash
gh repo edit "$GITHUB_OWNER/$REPO_NAME" --visibility private --accept-visibility-change-consequences
```

Script:

```bash
bash scripts/github/01-create-repository.sh
```

---

### 3.4 Tambahkan README.md

Checklist:

```text
- [ ] Tambahkan README.md
```

Jika `--add-readme` digunakan saat create repository, README sudah dibuat.

Jika belum ada, buat file:

```bash
cat > README.md <<'EOF'
# School Platform

Internal multi-unit school foundation platform for TK, SD, SMP, and SMA operations.

## Documentation

See `docs/README.md`.
EOF

git add README.md
git commit -m "docs(readme): add project readme"
git push origin main
```

Script:

```bash
bash scripts/github/02-bootstrap-repository-files.sh
```

---

### 3.5 Tambahkan `.gitignore`

Checklist:

```text
- [ ] Tambahkan .gitignore
```

Minimum `.gitignore`:

```gitignore
# Environment
.env
.env.*
!.env.example

# Secrets
*.pem
*.key
*.p12
*.pfx
id_rsa
id_ed25519

# OS / Editor
.DS_Store
Thumbs.db
.vscode/
.idea/

# Go
bin/
coverage.out

# Node / Next.js
node_modules/
.next/
out/
dist/

# Flutter
.dart_tool/
build/
.flutter-plugins
.flutter-plugins-dependencies

# Logs
*.log

# Local data
tmp/
storage/
```

Script:

```bash
bash scripts/github/02-bootstrap-repository-files.sh
```

---

### 3.6 Tambahkan LICENSE jika diperlukan

Checklist:

```text
- [ ] Tambahkan LICENSE jika diperlukan
```

Untuk private internal project, LICENSE opsional.

Recommended rule:

```text
- Jika repository tetap private internal yayasan: LICENSE boleh tidak dibuat.
- Jika repository akan dibuka sebagai open source: pilih license resmi seperti MIT/Apache-2.0/GPL sesuai keputusan owner.
```

Jika ingin menambahkan placeholder internal:

```bash
cat > LICENSE <<'EOF'
Copyright (c) 2026

This repository is private and intended for internal use only.
No license is granted for public use, distribution, or modification unless explicitly stated by the repository owner.
EOF

git add LICENSE
git commit -m "docs(repository): add internal license notice"
git push origin main
```

Script optional:

```bash
ADD_INTERNAL_LICENSE=true bash scripts/github/02-bootstrap-repository-files.sh
```

---

### 3.7 Buat branch `develop`

Checklist:

```text
- [ ] Buat branch `develop`
```

Command:

```bash
git checkout main
git pull origin main
git checkout -b develop
git push -u origin develop
```

Script:

```bash
bash scripts/github/03-create-branches.sh
```

---

### 3.8 Buat branch `staging`

Checklist:

```text
- [ ] Buat branch `staging`
```

Command:

```bash
git checkout main
git pull origin main
git checkout -b staging
git push -u origin staging
```

Script:

```bash
bash scripts/github/03-create-branches.sh
```

---

### 3.9 Gunakan `main` sebagai production branch

Checklist:

```text
- [ ] Gunakan `main` sebagai production branch
```

Set default branch ke `main`:

```bash
gh repo edit "$GITHUB_OWNER/$REPO_NAME" --default-branch main
```

Rule:

```text
- main = production branch
- staging = QA/UAT release candidate
- develop = development integration
```

Script:

```bash
bash scripts/github/03-create-branches.sh
```

---

### 3.10 Aktifkan branch protection untuk `develop`

Checklist:

```text
- [ ] Aktifkan branch protection untuk `develop`
```

Recommended rule:

```text
- Require pull request before merge
- Require at least 1 approval
- Require status checks
- Require conversation resolution
- Block force push
- Block delete branch
```

Script:

```bash
bash scripts/github/04-setup-branch-protection.sh
```

---

### 3.11 Aktifkan branch protection untuk `staging`

Checklist:

```text
- [ ] Aktifkan branch protection untuk `staging`
```

Recommended rule:

```text
- Require PR from develop
- Require CI pass
- Require conversation resolution
- Require QA/UAT validation before promote to main
- Block force push
- Block delete branch
```

Script:

```bash
bash scripts/github/04-setup-branch-protection.sh
```

---

### 3.12 Aktifkan branch protection untuk `main`

Checklist:

```text
- [ ] Aktifkan branch protection untuk `main`
```

Recommended rule:

```text
- Require PR from staging
- Require CI pass
- Require production approval via GitHub Environment
- Require conversation resolution
- Block force push
- Block delete branch
```

Script:

```bash
bash scripts/github/04-setup-branch-protection.sh
```

---

### 3.13 Tambahkan GitHub Environments: staging dan production

Checklist:

```text
- [ ] Tambahkan GitHub Environments: staging dan production
```

Create environments:

```bash
gh api \
  --method PUT \
  -H "Accept: application/vnd.github+json" \
  "/repos/$GITHUB_OWNER/$REPO_NAME/environments/staging"

gh api \
  --method PUT \
  -H "Accept: application/vnd.github+json" \
  "/repos/$GITHUB_OWNER/$REPO_NAME/environments/production"
```

Script:

```bash
bash scripts/github/05-setup-environments.sh
```

---

### 3.14 Tambahkan required reviewers untuk production environment

Checklist:

```text
- [ ] Tambahkan required reviewers untuk production environment
```

Set reviewer username:

```bash
export PRODUCTION_REVIEWER_USER="kuswandi-ti"
```

Jalankan:

```bash
bash scripts/github/05-setup-environments.sh
```

Catatan:

```text
- Required reviewer membutuhkan GitHub Environment protection rules.
- Untuk organization/team reviewer, sesuaikan reviewer type di script.
- Jika API gagal karena permission/plan, set manual via:
  Repository → Settings → Environments → production → Required reviewers
```

---

### 3.15 Tambahkan CODEOWNERS

Checklist:

```text
- [ ] Tambahkan CODEOWNERS
```

File path:

```text
.github/CODEOWNERS
```

Contoh:

```text
* @kuswandi-ti

/docs/ @kuswandi-ti
/.github/ @kuswandi-ti
/services/api-gateway/ @kuswandi-ti
/services/identity-service/ @kuswandi-ti
/services/school-core-service/ @kuswandi-ti
/services/finance-service/ @kuswandi-ti
/apps/web-admin/ @kuswandi-ti
/apps/mobile-app/ @kuswandi-ti
```

Script:

```bash
bash scripts/github/02-bootstrap-repository-files.sh
```

---

### 3.16 Tambahkan `pull_request_template.md`

Checklist:

```text
- [ ] Tambahkan pull_request_template.md
```

File path:

```text
.github/pull_request_template.md
```

Isi minimal:

```markdown
## Summary

## Related Issue

Closes #

## Type

- [ ] feature
- [ ] bug
- [ ] docs
- [ ] chore
- [ ] refactor
- [ ] test
- [ ] security
- [ ] infra

## Checklist

- [ ] PR title follows `type(scope): short description`
- [ ] CI pass
- [ ] Tests added/updated
- [ ] Docs updated if needed
- [ ] No secrets committed
- [ ] Permission/scope checked if relevant
```

Script:

```bash
bash scripts/github/02-bootstrap-repository-files.sh
```

---

### 3.17 Tambahkan issue templates

Checklist:

```text
- [ ] Tambahkan issue templates
```

File path:

```text
.github/ISSUE_TEMPLATE/feature_task.yml
.github/ISSUE_TEMPLATE/bug_report.yml
.github/ISSUE_TEMPLATE/ai_agent_task.yml
.github/ISSUE_TEMPLATE/security_review.yml
.github/ISSUE_TEMPLATE/qa_uat.yml
```

Script:

```bash
bash scripts/github/02-bootstrap-repository-files.sh
```

---

### 3.18 Tambahkan GitHub Actions CI workflow

Checklist:

```text
- [ ] Tambahkan GitHub Actions CI workflow
```

File path:

```text
.github/workflows/ci.yml
```

Minimal CI harus melakukan:

```text
- repository check
- secret check
- YAML validation
- Go check jika go.mod ada
- Web check jika apps/web-admin/package.json ada
- Mobile check jika apps/mobile-app/pubspec.yaml ada
- Docker Compose validation jika docker-compose.yml ada
```

Script:

```bash
bash scripts/github/02-bootstrap-repository-files.sh
```

---

### 3.19 Tambahkan labels

Checklist:

```text
- [ ] Tambahkan labels
```

Label groups:

```text
type
area
sprint
priority
status
ai
risk
review
```

Script:

```bash
bash scripts/github/06-setup-labels.sh
```

---

### 3.20 Tambahkan milestones Sprint 0 sampai Sprint 10

Checklist:

```text
- [ ] Tambahkan milestones Sprint 0 sampai Sprint 10
```

Milestones:

```text
Sprint 0 — Project Foundation
Sprint 1 — Identity & Access
Sprint 2 — School Core
Sprint 3 — File Management + Import Excel
Sprint 4 — PPDB
Sprint 5 — Finance / SPP
Sprint 6 — Academic Basic
Sprint 7 — Report Card / E-Rapor Basic
Sprint 8 — Communication / Notification
Sprint 9 — Reporting Dashboard
Sprint 10 — Security, Observability, Backup, UAT Hardening
MVP Release
```

Script:

```bash
bash scripts/github/07-setup-milestones.sh
```

---

### 3.21 Buat GitHub Project `School Platform MVP`

Checklist:

```text
- [ ] Buat GitHub Project `School Platform MVP`
```

Command:

```bash
gh project create --owner "$GITHUB_OWNER" --title "School Platform MVP"
```

Script:

```bash
bash scripts/github/08-setup-project.sh
```

Catatan:

```text
- Jika owner adalah organization, pastikan akun memiliki permission membuat project di organization.
- Jika command gagal, buat manual via GitHub UI:
  GitHub → Projects → New project → School Platform MVP
```

---

### 3.22 Tambahkan GitHub Project fields

Checklist:

```text
- [ ] Tambahkan GitHub Project fields
```

Fields:

```text
Status
Sprint
Priority
Area
Type
Owner
Estimate
Risk
Platform
AI Agent
Target Release
```

Script:

```bash
PROJECT_NUMBER="<project-number>" bash scripts/github/09-setup-project-fields.sh
```

Cara mendapatkan project number:

```bash
gh project list --owner "$GITHUB_OWNER"
```

Catatan:

```text
- GitHub Project v2 fields dapat dibuat via GitHub CLI.
- Beberapa field/options mungkin perlu disesuaikan manual melalui UI jika CLI/API berubah atau permission terbatas.
```

---

### 3.23 Tambahkan GitHub Project views

Checklist:

```text
- [ ] Tambahkan GitHub Project views
```

Recommended views:

```text
MVP Board
Sprint Board
Backlog Table
By Area
By Priority
QA/UAT
AI Agent Tasks
Release Readiness
```

Status:

```text
- GitHub CLI/API untuk Project views dapat berbeda tergantung versi dan permission.
- Jika script tidak dapat membuat views, buat manual via UI GitHub Project.
```

Manual setup:

```text
GitHub Project → + New view
```

Buat views:

| View | Layout | Filter/Group |
|---|---|---|
| MVP Board | Board | Group by Status |
| Sprint Board | Board | Filter Sprint = active sprint, group by Status |
| Backlog Table | Table | Status != Done |
| By Area | Board/Table | Group by Area |
| By Priority | Table | Group by Priority |
| QA/UAT | Board/Table | Status = QA or label status: needs-qa |
| AI Agent Tasks | Table | AI Agent = Ready or label ai: ready |
| Release Readiness | Board/Table | Target Release = MVP |

Helper script:

```bash
bash scripts/github/10-project-views-manual-guide.sh
```

Script ini mencetak panduan manual karena pembuatan Project views sering lebih stabil dilakukan via UI.

---

### 3.24 One-command setup sequence

Jika repository sudah di-clone dan environment variable sudah diset:

```bash
export GITHUB_OWNER="<org-or-user>"
export REPO_NAME="school-platform"
export PRODUCTION_REVIEWER_USER="kuswandi-ti"

bash scripts/github/00-run-all-github-setup.sh
```

Recommended execution order:

```text
1. 01-create-repository.sh
2. git clone repository
3. 02-bootstrap-repository-files.sh
4. 03-create-branches.sh
5. 04-setup-branch-protection.sh
6. 05-setup-environments.sh
7. 06-setup-labels.sh
8. 07-setup-milestones.sh
9. 08-setup-project.sh
10. 09-setup-project-fields.sh
11. 10-project-views-manual-guide.sh
```

---

### 3.25 Verification checklist

Setelah setup selesai:

```bash
gh repo view "$GITHUB_OWNER/$REPO_NAME"
gh label list --repo "$GITHUB_OWNER/$REPO_NAME"
gh api "/repos/$GITHUB_OWNER/$REPO_NAME/branches/main/protection"
gh api "/repos/$GITHUB_OWNER/$REPO_NAME/environments"
gh project list --owner "$GITHUB_OWNER"
```

Checklist validasi:

```text
- [ ] Repository private
- [ ] main/develop/staging tersedia
- [ ] main adalah default branch
- [ ] branch protection aktif
- [ ] staging environment tersedia
- [ ] production environment tersedia
- [ ] production required reviewer aktif atau diset manual
- [ ] CODEOWNERS tersedia
- [ ] PR template tersedia
- [ ] issue templates tersedia
- [ ] CI workflow tersedia
- [ ] labels tersedia
- [ ] milestones tersedia
- [ ] GitHub Project tersedia
- [ ] GitHub Project fields tersedia
- [ ] GitHub Project views tersedia atau dibuat manual
```


## 4. Recommended Repository Structure

Struktur repository awal:

```text
school-platform/
├── .github/
│   ├── CODEOWNERS
│   ├── pull_request_template.md
│   ├── workflows/
│   │   └── ci.yml
│   └── ISSUE_TEMPLATE/
│       ├── feature_task.yml
│       ├── bug_report.yml
│       ├── ai_agent_task.yml
│       ├── security_review.yml
│       └── qa_uat.yml
├── apps/
│   ├── web-admin/
│   └── mobile-app/
├── services/
│   ├── api-gateway/
│   ├── identity-service/
│   ├── school-core-service/
│   ├── admission-service/
│   ├── academic-service/
│   ├── finance-service/
│   ├── communication-service/
│   └── reporting-service/
├── packages/
│   ├── proto/
│   ├── openapi/
│   ├── events/
│   └── shared-go/
├── infra/
├── deploy/
├── docs/
├── scripts/
├── AGENTS.md
├── SKILLS.md
├── README.md
├── Makefile
├── docker-compose.yml
└── .gitignore
```

---

## 5. Branch Strategy

Branch utama:

| Branch | Fungsi |
|---|---|
| `develop` | Integrasi development harian |
| `staging` | QA/UAT dan release candidate |
| `main` | Production branch |

Branch kerja:

| Branch Pattern | Fungsi | Target PR |
|---|---|---|
| `feature/*` | Fitur baru atau task implementasi | `develop` |
| `fix/*` | Bug fix non-production | `develop` |
| `docs/*` | Dokumentasi | `develop` |
| `chore/*` | Maintenance / housekeeping | `develop` |
| `refactor/*` | Refactor tanpa perubahan behavior | `develop` |
| `test/*` | Penambahan/perubahan test | `develop` |
| `hotfix/*` | Perbaikan urgent production | `main` |

Branch flow:

```text
feature/* → develop
fix/*     → develop
docs/*    → develop
chore/*   → develop
refactor/* → develop
test/*    → develop

develop   → staging
staging   → main

hotfix/*  → main
main      → staging
main      → develop
```

---

## 6. Initial Git Commands

Jika repository sudah dibuat di GitHub:

```bash
git clone git@github.com:<org-or-user>/school-platform.git
cd school-platform
```

Buat branch utama:

```bash
git checkout main
git pull origin main

git checkout -b develop
git push origin develop

git checkout main
git checkout -b staging
git push origin staging
```

Kembali ke develop:

```bash
git checkout develop
```

Contoh membuat feature branch:

```bash
git checkout develop
git pull origin develop
git checkout -b feature/sprint-0-monorepo-structure
```

Push branch:

```bash
git push origin feature/sprint-0-monorepo-structure
```

---

## 7. Branch Protection Rules

### 7.1 develop Protection

Recommended rules:

```text
- Require a pull request before merging
- Require at least 1 approval
- Require status checks to pass
- Require branches to be up to date before merging
- Require conversation resolution before merging
- Block force pushes
- Block branch deletion
```

Purpose:

```text
- menjaga kualitas development integration
- mencegah direct push
- memastikan CI dan review berjalan
```

---

### 7.2 staging Protection

Recommended rules:

```text
- Require a pull request before merging
- Require at least 1 approval
- Require status checks to pass
- Require conversation resolution before merging
- Block force pushes
- Block branch deletion
- Require QA/UAT readiness before promoting to main
```

Purpose:

```text
- staging menjadi release candidate
- semua perubahan dari develop harus divalidasi di staging
```

---

### 7.3 main Protection

Recommended rules:

```text
- Require a pull request before merging
- Require approval
- Require status checks to pass
- Require conversation resolution before merging
- Block force pushes
- Block branch deletion
- Require production GitHub Environment approval
```

Purpose:

```text
- main hanya berisi production-ready code
- production deploy hanya dari main
- release membutuhkan approval manual
```

---

## 8. GitHub Environments

Buat GitHub Environments:

```text
staging
production
```

### 8.1 staging Environment

Digunakan untuk:

```text
- deploy dari branch staging
- QA/UAT
- release candidate validation
```

Recommended secrets:

```text
STAGING_DATABASE_URL
STAGING_REDIS_URL
STAGING_RABBITMQ_URL
STAGING_MINIO_ENDPOINT
STAGING_MINIO_ACCESS_KEY
STAGING_MINIO_SECRET_KEY
STAGING_JWT_SECRET
```

---

### 8.2 production Environment

Digunakan untuk:

```text
- deploy dari branch main
- production release
```

Production environment wajib:

```text
- required reviewer
- manual approval
- protected secrets
```

Recommended secrets:

```text
PRODUCTION_DATABASE_URL
PRODUCTION_REDIS_URL
PRODUCTION_RABBITMQ_URL
PRODUCTION_OBJECT_STORAGE_ENDPOINT
PRODUCTION_OBJECT_STORAGE_ACCESS_KEY
PRODUCTION_OBJECT_STORAGE_SECRET_KEY
PRODUCTION_JWT_SECRET
PRODUCTION_DEPLOY_KEY
```

Rules:

```text
- AI Agent tidak boleh menangani production secrets.
- Production deploy wajib manual approval.
- Jangan menyimpan secret di repository.
```

---

## 9. Required Repository Support Files

Repository harus memiliki:

| File | Purpose |
|---|---|
| `.github/CODEOWNERS` | Menentukan reviewer otomatis |
| `.github/pull_request_template.md` | Standar PR description dan checklist |
| `.github/workflows/ci.yml` | CI workflow |
| `.github/ISSUE_TEMPLATE/feature_task.yml` | Template feature/task |
| `.github/ISSUE_TEMPLATE/bug_report.yml` | Template bug report |
| `.github/ISSUE_TEMPLATE/ai_agent_task.yml` | Template task untuk AI Agent |
| `.github/ISSUE_TEMPLATE/security_review.yml` | Template security/privacy review |
| `.github/ISSUE_TEMPLATE/qa_uat.yml` | Template QA/UAT |
| `AGENTS.md` | Aturan AI Agent |
| `SKILLS.md` | Skill/workflow AI Agent |
| `docs/README.md` | Index dokumentasi |

---

## 10. Pull Request Rules

Semua perubahan wajib melalui Pull Request.

PR minimal harus:

```text
- link ke GitHub Issue
- menggunakan PR template
- target branch benar
- CI pass
- minimal 1 approval
- conversation resolved
- tidak mengandung secret
- tidak mengandung .env
- tidak keluar dari scope issue
```

Format PR title:

```text
type(scope): short description
```

Contoh:

```text
feat(identity): add refresh token rotation
fix(finance): prevent duplicate bill generation
docs(workflow): add github setup guide
chore(ci): add repository checks
test(academic): add attendance scope tests
```

---

## 11. GitHub Labels

GitHub Labels digunakan untuk mengelompokkan issue dan PR berdasarkan tipe, area, sprint, prioritas, status, risiko, review, dan AI Agent workflow.

Label taxonomy yang direkomendasikan:

```text
type
area
sprint
priority
status
ai
risk
review
```

---

## 12. Type Labels

| Label | Description |
|---|---|
| `type: feature` | Fitur baru atau implementation task |
| `type: bug` | Bug, regression, unexpected behavior |
| `type: chore` | Maintenance, dependency, housekeeping |
| `type: docs` | Dokumentasi |
| `type: refactor` | Refactor tanpa perubahan behavior |
| `type: test` | Penambahan/perubahan test |
| `type: security` | Security/privacy hardening atau fix |
| `type: infra` | Infra, Docker, deployment, CI/CD |
| `type: spike` | Research, proof of concept, technical exploration |
| `type: hotfix` | Urgent production fix |

---

## 13. Area Labels

| Label | Description |
|---|---|
| `area: api-gateway` | API Gateway |
| `area: identity` | Identity & Access |
| `area: school-core` | School Core |
| `area: admission` | PPDB / Admission |
| `area: academic` | Academic Basic / Report Card |
| `area: finance` | Finance / SPP |
| `area: communication` | Communication / Notification |
| `area: reporting` | Reporting Dashboard |
| `area: web-admin` | Next.js web admin |
| `area: mobile` | Flutter mobile app |
| `area: infra` | Infrastructure / deployment |
| `area: docs` | Documentation |
| `area: security` | Security / privacy |
| `area: observability` | Logs, metrics, monitoring |
| `area: ci-cd` | GitHub Actions / CI/CD |
| `area: file-management` | File upload, storage, import |

---

## 14. Sprint Labels

| Label | Sprint |
|---|---|
| `sprint: 0` | Sprint 0 — Project Foundation |
| `sprint: 1` | Sprint 1 — Identity & Access |
| `sprint: 2` | Sprint 2 — School Core |
| `sprint: 3` | Sprint 3 — File Management + Import Excel |
| `sprint: 4` | Sprint 4 — PPDB |
| `sprint: 5` | Sprint 5 — Finance / SPP |
| `sprint: 6` | Sprint 6 — Academic Basic |
| `sprint: 7` | Sprint 7 — Report Card / E-Rapor Basic |
| `sprint: 8` | Sprint 8 — Communication / Notification |
| `sprint: 9` | Sprint 9 — Reporting Dashboard |
| `sprint: 10` | Sprint 10 — Security, Observability, Backup, UAT Hardening |

---

## 15. Priority Labels

| Label | Description | Release Impact |
|---|---|---|
| `priority: critical` | Blocker, production risk, security/data loss, release blocker | Blocks release |
| `priority: high` | Core MVP flow or severe bug | Blocks release if core flow |
| `priority: medium` | Important but not release-blocking | May be deferred |
| `priority: low` | Minor improvement or cleanup | Usually non-blocking |

Rule:

```text
Critical/High bug pada core flow memblokir production release.
```

---

## 16. Status Labels

| Label | Description |
|---|---|
| `status: ready` | Issue siap dikerjakan |
| `status: in-progress` | Sedang dikerjakan |
| `status: blocked` | Terblokir |
| `status: needs-review` | Butuh review |
| `status: needs-qa` | Butuh QA |
| `status: qa-passed` | QA pass |
| `status: done` | Selesai |

Catatan:

```text
Status utama tetap dikelola di GitHub Project field `Status`.
Status label digunakan untuk filter cepat.
```

---

## 17. AI Agent Labels

| Label | Description |
|---|---|
| `ai: ready` | Task siap dibantu AI Agent |
| `ai: needs-context` | Task belum cukup jelas untuk AI Agent |
| `ai: generated` | Output dibuat/dibantu AI Agent |
| `ai: needs-human-review` | Wajib human review |
| `ai: do-not-use-agent` | AI Agent tidak boleh digunakan |

Gunakan `ai: do-not-use-agent` untuk:

```text
- production secrets
- final security approval
- legal/compliance decision
- production deployment approval
- akses data asli sensitif
- perubahan arsitektur besar tanpa instruksi
- hotfix production tanpa human review
```

---

## 18. Risk Labels

| Label | Description |
|---|---|
| `risk: low` | Risiko rendah |
| `risk: medium` | Risiko sedang |
| `risk: high` | Risiko tinggi |
| `risk: breaking-change` | Berpotensi breaking change |
| `risk: migration` | Menyentuh migration/data |
| `risk: data-sensitive` | Menyentuh Restricted/Confidential data |

---

## 19. Review Labels

| Label | Reviewer Focus |
|---|---|
| `review: backend` | Backend review |
| `review: frontend` | Frontend review |
| `review: mobile` | Mobile review |
| `review: infra` | Infrastructure/DevOps review |
| `review: qa` | QA review |
| `review: security` | Security/privacy review |
| `review: product` | Product/business review |

---

## 20. Recommended GitHub Label Setup Commands

Gunakan GitHub CLI:

```bash
gh auth login
gh repo set-default <org-or-user>/school-platform
```

### 20.1 Type Labels

```bash
gh label create "type: feature" --color "1f883d" --description "New feature or implementation task"
gh label create "type: bug" --color "d1242f" --description "Bug or regression"
gh label create "type: chore" --color "6e7781" --description "Maintenance or housekeeping"
gh label create "type: docs" --color "0969da" --description "Documentation"
gh label create "type: refactor" --color "a2eeef" --description "Refactor without behavior change"
gh label create "type: test" --color "0e8a16" --description "Test changes"
gh label create "type: security" --color "8250df" --description "Security or privacy"
gh label create "type: infra" --color "6f42c1" --description "Infrastructure or CI/CD"
gh label create "type: spike" --color "d4c5f9" --description "Research or exploration"
gh label create "type: hotfix" --color "b60205" --description "Urgent production fix"
```

### 20.2 Priority Labels

```bash
gh label create "priority: critical" --color "b60205" --description "Release blocker or critical issue"
gh label create "priority: high" --color "d93f0b" --description "High priority"
gh label create "priority: medium" --color "fbca04" --description "Medium priority"
gh label create "priority: low" --color "0e8a16" --description "Low priority"
```

### 20.3 AI Agent Labels

```bash
gh label create "ai: ready" --color "5319e7" --description "Ready for AI Agent"
gh label create "ai: needs-context" --color "c5def5" --description "Needs more context before AI Agent"
gh label create "ai: generated" --color "bfd4f2" --description "Generated or assisted by AI Agent"
gh label create "ai: needs-human-review" --color "f9d0c4" --description "Requires human review"
gh label create "ai: do-not-use-agent" --color "000000" --description "Do not use AI Agent"
```

### 20.4 Sprint Labels

```bash
for i in 0 1 2 3 4 5 6 7 8 9 10; do
  gh label create "sprint: $i" --color "ededed" --description "Sprint $i"
done
```

---

## 21. GitHub Milestones

Buat milestones berikut:

```text
Sprint 0 — Project Foundation
Sprint 1 — Identity & Access
Sprint 2 — School Core
Sprint 3 — File Management + Import Excel
Sprint 4 — PPDB
Sprint 5 — Finance / SPP
Sprint 6 — Academic Basic
Sprint 7 — Report Card / E-Rapor Basic
Sprint 8 — Communication / Notification
Sprint 9 — Reporting Dashboard
Sprint 10 — Security, Observability, Backup, UAT Hardening
MVP Release
```

Rules:

```text
- Setiap issue sprint wajib punya milestone.
- PR harus menutup atau mengacu pada issue di milestone yang sama.
- Milestone ditutup setelah semua blocking issue selesai.
```

---

## 22. GitHub Project Setup

Recommended GitHub Project name:

```text
School Platform MVP
```

Project type:

```text
GitHub Projects / Project v2
```

Purpose:

```text
- backlog management
- sprint tracking
- QA/UAT tracking
- AI Agent task tracking
- release readiness
```

---

## 23. GitHub Project Fields

Tambahkan fields berikut:

| Field | Type | Values |
|---|---|---|
| `Status` | Single select | Backlog, Ready, In Progress, In Review, QA, Blocked, Done |
| `Sprint` | Single select | Sprint 0, Sprint 1, Sprint 2, Sprint 3, Sprint 4, Sprint 5, Sprint 6, Sprint 7, Sprint 8, Sprint 9, Sprint 10, MVP Release |
| `Priority` | Single select | Critical, High, Medium, Low |
| `Area` | Single select | api-gateway, identity, school-core, admission, academic, finance, communication, reporting, web-admin, mobile, infra, docs, security, observability, ci-cd, file-management |
| `Type` | Single select | feature, bug, chore, docs, refactor, test, security, infra, spike, hotfix |
| `Owner` | User | GitHub assignee |
| `Estimate` | Number | 1, 2, 3, 5, 8 |
| `Risk` | Single select | Low, Medium, High, Breaking Change, Migration, Data Sensitive |
| `Platform` | Single select | Backend, Web, Mobile, Infra, Docs, QA, Product |
| `AI Agent` | Single select | Ready, Needs Context, Generated, Needs Human Review, Do Not Use |
| `Target Release` | Text / Single select | MVP, v0.1, v1.0 |

---

## 24. GitHub Project Views

### 24.1 MVP Board

Purpose:

```text
Board utama seluruh pekerjaan MVP.
```

Configuration:

```text
Layout: Board
Group by: Status
Columns: Backlog, Ready, In Progress, In Review, QA, Blocked, Done
```

---

### 24.2 Sprint Board

Purpose:

```text
Tracking sprint aktif.
```

Configuration:

```text
Layout: Board
Filter: Sprint = active sprint
Group by: Status
```

Visible fields:

```text
Priority
Area
Owner
Estimate
Risk
AI Agent
```

---

### 24.3 Backlog Table

Purpose:

```text
Backlog grooming dan planning.
```

Configuration:

```text
Layout: Table
Filter: Status != Done
Sort: Priority descending
```

---

### 24.4 By Area

Purpose:

```text
Melihat pekerjaan per modul/service/app.
```

Configuration:

```text
Layout: Board or Table
Group by: Area
```

---

### 24.5 By Priority

Purpose:

```text
Fokus pada Critical/High issue.
```

Configuration:

```text
Layout: Table
Group by: Priority
Sort: Priority descending
```

---

### 24.6 QA/UAT

Purpose:

```text
Melihat issue yang siap atau sedang QA/UAT.
```

Configuration:

```text
Filter: Status = QA OR label:"status: needs-qa"
Group by: Area
```

---

### 24.7 AI Agent Tasks

Purpose:

```text
Melihat task yang bisa dibantu AI Agent.
```

Configuration:

```text
Filter: AI Agent = Ready OR label:"ai: ready"
```

---

### 24.8 Release Readiness

Purpose:

```text
Melacak kesiapan staging/production release.
```

Configuration:

```text
Filter: Target Release = MVP OR Sprint = MVP Release
Group by: Status
```

Highlight:

```text
- Critical/High open bugs
- security review issues
- QA/UAT sign-off issues
- backup/restore readiness
- rollback readiness
```

---

## 25. Issue Lifecycle

Lifecycle issue:

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

| Status | Meaning | Entry Rule | Exit Rule |
|---|---|---|---|
| Backlog | Candidate work | Issue dibuat tapi belum siap | Scope dan AC lengkap |
| Ready | Siap dikerjakan | DoR terpenuhi | Owner mulai kerja |
| In Progress | Sedang dikerjakan | Developer/AI Agent operator mulai kerja | PR dibuat |
| In Review | PR direview | PR opened | PR approved/merged atau revisi |
| QA | Butuh QA/UAT | Feature merged/ready for QA | QA pass atau bug dibuat |
| Done | Selesai | DoD terpenuhi | Issue closed |
| Blocked | Terblokir | Ada dependency/decision missing | Blocker resolved |

---

## 26. Definition of Ready

Issue siap dikerjakan jika:

```text
- objective jelas
- scope jelas
- out of scope jelas
- acceptance criteria jelas
- area jelas
- priority jelas
- risk jelas
- milestone jelas
- dependency diketahui
- owner siap
- AI Agent status jelas
- required docs/prompts diketahui
```

---

## 27. Definition of Done

Task selesai jika:

```text
- implementation sesuai scope
- tidak ada out-of-scope feature
- tests ditambahkan/diupdate
- tests pass
- lint/format pass
- CI pass
- permission/scope dicek jika relevan
- object-level authorization dicek jika relevan
- audit/event/file/privacy dicek jika relevan
- dokumentasi diupdate jika relevan
- PR approved
- QA pass jika diperlukan
- issue ditutup atau dipindah Done
```

---

## 28. AI Agent Workflow

AI Agent boleh digunakan untuk task kecil dan jelas.

Workflow:

```text
1. Pilih issue dengan label `ai: ready`.
2. Baca AGENTS.md.
3. Baca SKILLS.md.
4. Baca docs/README.md.
5. Baca dokumen sprint/task terkait.
6. Jalankan prompt sesuai scope issue.
7. Review hasil AI Agent.
8. Jalankan test.
9. Commit ke feature branch.
10. Buat PR.
11. Human review wajib.
```

AI Agent tidak boleh:

```text
- menangani production secrets
- final security approval
- legal/compliance decision
- production deployment approval
- akses data asli sensitif
- perubahan arsitektur besar tanpa instruksi
- final decision data privacy
- hotfix production tanpa human review
```

---

## 29. QA/UAT Workflow

QA issue menggunakan template:

```text
.github/ISSUE_TEMPLATE/qa_uat.yml
```

QA flow:

```text
Ready for QA
→ QA In Progress
→ QA Passed / QA Failed
→ Done or Bug Created
```

Bug severity:

| Severity | Release Impact |
|---|---|
| Critical | Block production |
| High | Block production jika core flow |
| Medium | Bisa deferred jika ada workaround |
| Low | Tidak block release |

Release rule:

```text
Critical/High bug pada core flow memblokir production release.
```

---

## 30. Release Readiness Workflow

Sebelum production release, pastikan:

```text
- staging QA/UAT pass
- no Critical/High core bug
- CI pass
- security review selesai jika diperlukan
- backup/restore siap
- rollback plan tersedia
- release notes tersedia
- production approval diberikan
```

Production deploy hanya dari:

```text
main
```

Production deploy wajib:

```text
manual approval via GitHub Environment
```

---

## 31. Hotfix Workflow

Hotfix digunakan untuk production issue urgent.

Flow:

```text
main
→ hotfix/*
→ PR to main
→ production deploy
→ release tag
→ back-merge main to staging
→ back-merge main to develop
```

Hotfix PR wajib mencakup:

```text
- root cause
- scope
- tests
- risk
- rollback
```

---

## 32. Recommended Setup Order

Urutan setup yang disarankan:

```text
1. Create repository `school-platform`
2. Create branches: develop, staging, main
3. Add .gitignore and README
4. Add AGENTS.md and SKILLS.md
5. Add docs/README.md
6. Add CODEOWNERS
7. Add PR template
8. Add Issue Templates
9. Add CI workflow
10. Add branch protection
11. Add GitHub Environments
12. Add Labels
13. Add Milestones
14. Create GitHub Project `School Platform MVP`
15. Add Project fields
16. Add Project views
17. Create Sprint 0 issues
18. Start Sprint 0 implementation
```

---

## 33. Final Checklist

```text
- [ ] Repository created as private
- [ ] develop/staging/main branches created
- [ ] Branch protection enabled
- [ ] GitHub Environments created
- [ ] CODEOWNERS added
- [ ] PR template added
- [ ] Issue templates added
- [ ] CI workflow added
- [ ] Labels created
- [ ] Milestones created
- [ ] GitHub Project created
- [ ] GitHub Project fields created
- [ ] GitHub Project views created
- [ ] Sprint 0 issues created
- [ ] Team understands workflow
```

---

## 34. Related Documents

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/11-github-repository-rules.md
docs/25-github-project-management.md
docs/26-development-plan.md
docs/27-workflow.md
docs/29-sprint-0-plan.md
docs/30-sprint-1-plan.md sampai docs/39-sprint-10-plan.md
.github/CODEOWNERS
.github/pull_request_template.md
.github/workflows/ci.yml
.github/ISSUE_TEMPLATE/*.yml
```

## 35. Git Commit Convention

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


### Alignment with Labels and PR

Commit `type` should generally align with GitHub `type:*` labels:

| Commit Type | GitHub Label |
|---|---|
| `feat` | `type: feature` |
| `fix` | `type: bug` |
| `docs` | `type: docs` |
| `chore` | `type: chore` |
| `refactor` | `type: refactor` |
| `test` | `type: test` |
| `security` | `type: security` |
| `ci` / `build` | `type: infra` or `area: ci-cd` |
| `revert` | use original related type plus PR explanation |

PR title should follow the same convention:

```text
type(scope): short description
```
