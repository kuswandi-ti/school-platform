# 25 — GitHub Project Management

Project: `school-platform`  
Purpose: GitHub Labels, Milestones, GitHub Project, Issue Templates, PR tracking, and Sprint execution rules.

---

## 1. Purpose

Dokumen ini menjelaskan standar penggunaan GitHub untuk project `school-platform`.

Dokumen ini mencakup:

```text
- GitHub Labels
- GitHub Milestones
- GitHub Project
- GitHub Project fields
- GitHub Project views
- Issue workflow
- Pull Request workflow
- AI Agent workflow
- QA/UAT workflow
- Sprint tracking
- Release readiness tracking
```

Tujuannya agar seluruh pekerjaan di project `school-platform` bisa dilacak secara rapi dari:

```text
Planning
→ GitHub Issues
→ Development
→ Pull Request
→ Review
→ QA/UAT
→ Staging
→ Production
```

---

## 2. Repository Management Scope

GitHub digunakan untuk:

```text
- source code management
- documentation management
- task tracking
- sprint tracking
- issue tracking
- PR review
- CI/CD gate
- release tracking
- AI Agent task coordination
```

CI/CD gate harus path-aware; web dan mobile checks skip dengan bersih jika folder app target belum tersedia.

Repository utama:

```text
school-platform
```

Branch utama:

```text
develop   = development integration
staging   = QA/UAT/staging release candidate
main      = production
```

Branch workflow:

```text
feature/* → develop
fix/*     → develop
docs/*    → develop
chore/*   → develop
develop   → staging
staging   → main
hotfix/*  → main → back-merge staging/develop
```

---

## 3. GitHub Labels

Labels digunakan untuk mengelompokkan issue, PR, task, bug, sprint, area, priority, risk, review, dan status AI Agent.

Label harus konsisten agar GitHub Project bisa difilter dengan mudah.

---

## 4. Label Categories

## 4.1 Type Labels

| Label | Usage |
|---|---|
| `type: feature` | Fitur baru atau implementasi task |
| `type: bug` | Bug, regression, atau unexpected behavior |
| `type: chore` | Maintenance, cleanup, dependency, repo housekeeping |
| `type: docs` | Dokumentasi |
| `type: refactor` | Refactor tanpa perubahan behavior |
| `type: test` | Penambahan/perubahan test |
| `type: security` | Security hardening, vulnerability fix, privacy fix |
| `type: infra` | Infrastructure, Docker, deployment, CI/CD |
| `type: spike` | Research/technical exploration |
| `type: hotfix` | Perbaikan urgent untuk production |

---

## 4.2 Area Labels

| Label | Usage |
|---|---|
| `area: api-gateway` | API Gateway |
| `area: identity` | Identity & Access service |
| `area: school-core` | School Core service |
| `area: admission` | PPDB / Admission service |
| `area: academic` | Academic service |
| `area: finance` | Finance / SPP service |
| `area: communication` | Communication / Notification service |
| `area: reporting` | Reporting service |
| `area: web-admin` | Next.js web admin |
| `area: mobile` | Flutter mobile app |
| `area: infra` | Infrastructure, Docker, deploy |
| `area: docs` | Documentation |
| `area: security` | Security/privacy/cross-cutting |
| `area: observability` | Logging, metrics, monitoring |
| `area: ci-cd` | GitHub Actions / CI/CD |
| `area: file-management` | File upload, storage, signed URL, import file |

---

## 4.3 Sprint Labels

| Label | Sprint |
|---|---|
| `sprint: 0` | Project Foundation |
| `sprint: 1` | Identity & Access |
| `sprint: 2` | School Core |
| `sprint: 3` | File Management + Import Excel |
| `sprint: 4` | PPDB |
| `sprint: 5` | Finance / SPP |
| `sprint: 6` | Academic Basic |
| `sprint: 7` | Report Card / E-Rapor Basic |
| `sprint: 8` | Communication / Notification |
| `sprint: 9` | Reporting Dashboard |
| `sprint: 10` | Security, Observability, Backup, UAT Hardening |

---

## 4.4 Priority Labels

| Label | Meaning |
|---|---|
| `priority: critical` | Blocker, production risk, security/data loss, release blocker |
| `priority: high` | Important MVP core flow or severe bug |
| `priority: medium` | Important but not blocking |
| `priority: low` | Nice-to-have, cleanup, minor improvement |

Rules:

```text
Critical/High bugs in MVP core flows block release.
```

---

## 4.5 Status Labels

| Label | Meaning |
|---|---|
| `status: ready` | Issue siap dikerjakan |
| `status: in-progress` | Sedang dikerjakan |
| `status: blocked` | Terblokir dependency/decision |
| `status: needs-review` | Butuh review teknis/product/security |
| `status: needs-qa` | Butuh QA/UAT |
| `status: qa-passed` | QA sudah pass |
| `status: done` | Selesai |

Status utama tetap dikelola di GitHub Project field `Status`. Label status hanya dipakai untuk filtering cepat.

---

## 4.6 AI Agent Labels

| Label | Meaning |
|---|---|
| `ai: ready` | Task siap dikerjakan AI Agent |
| `ai: needs-context` | Task belum cukup jelas untuk AI Agent |
| `ai: generated` | Output dibuat/dibantu AI Agent |
| `ai: needs-human-review` | Wajib human review |
| `ai: do-not-use-agent` | Jangan gunakan AI Agent |

Gunakan `ai: do-not-use-agent` untuk:

```text
- production secrets
- final security approval
- legal/compliance decision
- production deployment approval
- akses data asli sensitif
- perubahan arsitektur besar tanpa approval
- hotfix production tanpa human review
```

---

## 4.7 Risk Labels

| Label | Meaning |
|---|---|
| `risk: low` | Risiko rendah |
| `risk: medium` | Risiko sedang |
| `risk: high` | Risiko tinggi |
| `risk: breaking-change` | Berpotensi breaking change |
| `risk: migration` | Menyentuh migration/data |
| `risk: data-sensitive` | Menyentuh data Restricted/Confidential |

---

## 4.8 Review Labels

| Label | Reviewer Focus |
|---|---|
| `review: backend` | Backend review |
| `review: frontend` | Web frontend review |
| `review: mobile` | Mobile review |
| `review: infra` | Infrastructure/DevOps review |
| `review: qa` | QA review |
| `review: security` | Security/privacy review |
| `review: product` | Product/business flow review |

---

## 5. Recommended Label Setup Command

Jika menggunakan GitHub CLI, labels bisa dibuat dengan script.

Contoh:

```bash
gh label create "type: feature" --color "1f883d" --description "New feature or implementation task"
gh label create "type: bug" --color "d1242f" --description "Bug or regression"
gh label create "type: docs" --color "0969da" --description "Documentation"
gh label create "type: security" --color "8250df" --description "Security or privacy"
gh label create "type: infra" --color "6f42c1" --description "Infrastructure or CI/CD"

gh label create "priority: critical" --color "b60205" --description "Release blocker or critical issue"
gh label create "priority: high" --color "d93f0b" --description "High priority"
gh label create "priority: medium" --color "fbca04" --description "Medium priority"
gh label create "priority: low" --color "0e8a16" --description "Low priority"

gh label create "ai: ready" --color "5319e7" --description "Ready for AI Agent"
gh label create "ai: needs-context" --color "c5def5" --description "Needs more context before AI Agent"
gh label create "ai: generated" --color "bfd4f2" --description "Generated or assisted by AI Agent"
gh label create "ai: needs-human-review" --color "f9d0c4" --description "Requires human review"
gh label create "ai: do-not-use-agent" --color "000000" --description "Do not use AI Agent"
```

Atau buat labels manual melalui:

```text
Repository → Issues → Labels → New label
```

---

## 6. GitHub Milestones

Milestone digunakan untuk mengelompokkan issue/PR berdasarkan sprint.

Recommended milestones:

| Milestone | Purpose |
|---|---|
| `Sprint 0 — Project Foundation` | Fondasi repo, local dev, CI, service skeleton |
| `Sprint 1 — Identity & Access` | Auth, RBAC, ABAC/scope foundation |
| `Sprint 2 — School Core` | Master data yayasan/sekolah/siswa/guru/class |
| `Sprint 3 — File Management + Import Excel` | Private file + import data awal |
| `Sprint 4 — PPDB` | Admission process |
| `Sprint 5 — Finance / SPP` | Fee, bill, manual payment |
| `Sprint 6 — Academic Basic` | Curriculum, subject, schedule, attendance |
| `Sprint 7 — Report Card / E-Rapor Basic` | Grade and report card workflow |
| `Sprint 8 — Communication / Notification` | Announcement and notification |
| `Sprint 9 — Reporting Dashboard` | Reporting read model and dashboard |
| `Sprint 10 — Security, Observability, Backup, UAT Hardening` | Production readiness |
| `MVP Release` | Final MVP release readiness |

Rules:

```text
- Every sprint issue must have a milestone.
- PR should link to an issue in the same milestone.
- Milestone can be closed only after all blocking issues are done.
```

---

## 7. GitHub Project

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
- sprint tracking
- backlog management
- QA/UAT tracking
- AI Agent task tracking
- release readiness
```

---

## 8. GitHub Project Fields

Create these custom fields:

| Field | Type | Values |
|---|---|---|
| `Status` | Single select | Backlog, Ready, In Progress, In Review, QA, Blocked, Done |
| `Sprint` | Single select | Sprint 0 to Sprint 10, MVP Release |
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

## 9. GitHub Project Views

## 9.1 MVP Board

Purpose:

```text
Main execution board for all MVP work.
```

Configuration:

```text
Layout: Board
Group by: Status
Columns: Backlog, Ready, In Progress, In Review, QA, Blocked, Done
```

---

## 9.2 Sprint Board

Purpose:

```text
Daily sprint tracking.
```

Configuration:

```text
Layout: Board
Filter: Sprint = active sprint
Group by: Status
```

Recommended visible fields:

```text
Priority
Area
Owner
Estimate
Risk
AI Agent
```

---

## 9.3 Backlog Table

Purpose:

```text
Backlog grooming and planning.
```

Configuration:

```text
Layout: Table
Filter: Status != Done
Sort: Priority desc
```

Visible fields:

```text
Title
Status
Sprint
Priority
Area
Type
Owner
Estimate
Risk
AI Agent
```

---

## 9.4 By Area

Purpose:

```text
Track work grouped by module/service.
```

Configuration:

```text
Layout: Board or Table
Group by: Area
```

---

## 9.5 By Priority

Purpose:

```text
Focus on Critical/High work.
```

Configuration:

```text
Layout: Table
Group by: Priority
Sort: Priority desc
```

---

## 9.6 QA / UAT

Purpose:

```text
Track issues ready for QA/UAT.
```

Configuration:

```text
Filter: Status = QA OR label:"status: needs-qa"
Group by: Area
```

---

## 9.7 AI Agent Tasks

Purpose:

```text
Track tasks that can be handled by AI Agent.
```

Configuration:

```text
Filter: AI Agent = Ready OR label:"ai: ready"
```

Recommended visible fields:

```text
Title
Status
Sprint
Area
Priority
Risk
Owner
AI Agent
```

---

## 9.8 Release Readiness

Purpose:

```text
Track readiness for staging/production release.
```

Configuration:

```text
Filter: Target Release = MVP OR Sprint = MVP Release
Group by: Status
```

Must highlight:

```text
- Critical/High open bugs
- Security review issues
- QA/UAT sign-off issues
- Production readiness issues
- Backup/restore issues
```

---

## 10. Issue Templates

Use these issue templates:

| File | Purpose |
|---|---|
| `.github/ISSUE_TEMPLATE/feature_task.yml` | Feature/task implementation |
| `.github/ISSUE_TEMPLATE/bug_report.yml` | Bug/regression report |
| `.github/ISSUE_TEMPLATE/ai_agent_task.yml` | AI Agent-ready task |
| `.github/ISSUE_TEMPLATE/security_review.yml` | Security/privacy review |
| `.github/ISSUE_TEMPLATE/qa_uat.yml` | QA/UAT scenario or sign-off |

Rules:

```text
- Use feature_task.yml for normal implementation work.
- Use ai_agent_task.yml when AI Agent will be used directly.
- Use bug_report.yml for defects.
- Use security_review.yml for sensitive/security/privacy work.
- Use qa_uat.yml for QA/UAT scenarios and sign-off.
```

---

## 11. Issue Lifecycle

Issue status flow:

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

---

## 12. Issue Status Rules

| Status | Meaning | Entry Rule | Exit Rule |
|---|---|---|---|
| Backlog | Candidate work | Created but not ready | Groomed and accepted |
| Ready | Ready to work | Scope and acceptance criteria complete | Assigned and started |
| In Progress | Being worked on | Developer/AI Agent starts work | PR opened |
| In Review | PR under review | PR opened and linked | PR approved/merged or needs changes |
| QA | Needs QA/UAT | Feature merged to staging or ready for QA | QA passed or bug created |
| Blocked | Cannot proceed | Dependency/decision missing | Blocker resolved |
| Done | Completed | PR merged and QA passed if needed | Closed |

---

## 13. Definition of Ready

An issue is Ready when:

```text
- objective is clear
- scope is clear
- out of scope is clear
- acceptance criteria are defined
- area is defined
- priority is defined
- milestone is defined
- dependencies are known
- required docs/prompts are identified
- AI Agent status is clear
```

---

## 14. Definition of Done

An issue is Done when:

```text
- implementation matches scope
- tests added/updated
- tests pass
- lint/format pass
- permission/scope checked if relevant
- object-level authorization checked if relevant
- audit/event/file/privacy handled if relevant
- documentation updated if relevant
- PR reviewed
- QA passed if required
- no Critical/High bug remains for the issue
```

---

## 15. AI Agent Issue Rules

AI Agent can work only on issues where:

```text
- scope is small and clear
- acceptance criteria are explicit
- required docs are listed
- AI Agent label/status is ready
- task does not involve production secrets
- task does not require final legal/security/compliance decision
- task does not change architecture without explicit approval
```

Required docs for AI Agent implementation:

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/09-ai-agent-rules.md
docs/08-coding-standard.md
active sprint plan if available
active sprint task prompt
task-specific technical documents
```

Planning prompt docs:

```text
docs/25-prd-prompt.md
docs/26-development-plan-prompt.md
docs/27-workflow-prompt.md
docs/28-ai-agent-sprint-planning-prompts.md
```

Coding prompt docs:

```text
docs/13-sprint-0-task-prompts.md
docs/14-sprint-1-task-prompts.md
...
docs/23-sprint-10-task-prompts.md
```

---

## 16. Pull Request Rules

Every PR must:

```text
- link to an issue
- use PR template
- target the correct branch
- pass CI
- receive required review
- update docs if needed
- not include secrets
- not include .env files
```

PR title format:

```text
type(scope): short description
```

Examples:

```text
feat(identity): add refresh token rotation
fix(finance): prevent duplicate bill generation
docs(workflow): add github project management guide
chore(ci): add repository check workflow
```

---

## 17. PR Target Rules

| Source Branch | Target Branch |
|---|---|
| `feature/*` | `develop` |
| `fix/*` | `develop` |
| `docs/*` | `develop` |
| `chore/*` | `develop` |
| `develop` | `staging` |
| `staging` | `main` |
| `hotfix/*` | `main` |

Hotfix rule:

```text
hotfix/* → main
main → staging
main → develop
```

---

## 18. Branch Protection Summary

Recommended branch protection:

| Branch | Rule |
|---|---|
| `develop` | PR required, 1 approval, CI required, no force push, no delete |
| `staging` | PR required, 1 approval, CI required, QA/UAT required, no force push, no delete |
| `main` | PR required, approval required, CI required, production approval, no force push, no delete |

---

## 19. Sprint Execution Workflow

For each sprint:

```text
1. Read sprint plan.
2. Create issues from sprint plan.
3. Add milestone.
4. Add labels.
5. Add to GitHub Project.
6. Move Ready issues to active sprint.
7. Assign owner.
8. Execute issue.
9. Open PR.
10. Review PR.
11. Merge to develop.
12. Promote develop to staging.
13. Run QA/UAT.
14. Promote staging to main when approved.
```

---

## 20. GitHub Project Usage by Sprint

At sprint start:

```text
- filter Backlog by sprint
- confirm Ready issues
- confirm priority
- confirm owner
- confirm estimate
- confirm dependencies
- confirm AI Agent status
```

During sprint:

```text
- update Status daily
- keep blockers visible
- link PRs to issues
- move completed work to QA/Done
```

At sprint end:

```text
- close done issues
- document unfinished issues
- move carry-over issues
- review risks
- update sprint handoff notes
```

---

## 21. QA/UAT Tracking

QA issues should use:

```text
.github/ISSUE_TEMPLATE/qa_uat.yml
```

QA labels:

```text
review: qa
status: needs-qa
```

QA/UAT status:

```text
- Ready for QA
- QA In Progress
- QA Passed
- QA Failed
- Blocked
```

Release blocker:

```text
Critical/High bug in MVP core flow blocks release.
```

---

## 22. Security Review Tracking

Security review issues should use:

```text
.github/ISSUE_TEMPLATE/security_review.yml
```

Security labels:

```text
type: security
review: security
risk: high
risk: data-sensitive
```

Security review is required for:

```text
- authentication
- authorization
- object-level access
- finance/payment verification
- file download/export
- Restricted/Confidential data
- role assignment
- production deployment workflow
```

---

## 23. Release Readiness Tracking

Before production release, confirm:

```text
- no Critical/High bug remains open
- QA sign-off complete
- UAT sign-off complete if required
- security review complete
- CI pass
- staging verified
- backup/restore tested
- rollback plan documented
- production approval ready
```

Use view:

```text
Release Readiness
```

---

## 24. Documentation Tracking

Documentation updates should be tracked as issues when they affect:

```text
- architecture
- service boundary
- data model
- API contract
- event contract
- UI flow
- test plan
- coding standard
- AI Agent rules
- sprint plan
- workflow/SOP
- local development
```

Use label:

```text
type: docs
area: docs
```

---

## 25. Recommended GitHub Setup Order

Setup order:

```text
1. Create repository.
2. Create develop, staging, main branches.
3. Add branch protection.
4. Add GitHub Environments: staging and production.
5. Add CODEOWNERS.
6. Add PR template.
7. Add Issue Templates.
8. Add CI workflow.
9. Create labels.
10. Create milestones.
11. Create GitHub Project: School Platform MVP.
12. Create Project fields.
13. Create Project views.
14. Create Sprint 0 issues.
15. Start Sprint 0 implementation.
```

---

## 26. Recommended Initial Sprint 0 Issues

Suggested initial issues:

| Issue Title | Type | Area | Priority | Labels | Milestone |
|---|---|---|---|---|---|
| Sprint 0 Task 0.1 — Create Monorepo Structure | feature | infra | high | `type: feature`, `area: infra`, `sprint: 0`, `ai: ready` | Sprint 0 |
| Sprint 0 Task 0.2 — Add Docker Compose Local Dependencies | feature | infra | high | `type: infra`, `area: infra`, `sprint: 0` | Sprint 0 |
| Sprint 0 Task 0.3 — Create Go Service Template | feature | backend | high | `type: feature`, `area: api-gateway`, `sprint: 0` | Sprint 0 |
| Sprint 0 Task 0.4 — Create API Gateway Skeleton | feature | api-gateway | high | `type: feature`, `area: api-gateway`, `sprint: 0` | Sprint 0 |
| Sprint 0 Task 0.5 — Add Basic CI | infra | ci-cd | high | `type: infra`, `area: ci-cd`, `sprint: 0` | Sprint 0 |
| Sprint 0 Task 0.6 — Add Repository Docs and Local Guide | docs | docs | medium | `type: docs`, `area: docs`, `sprint: 0` | Sprint 0 |

---

## 27. Final Rules

```text
- Every task must be an issue.
- Every PR must link to an issue.
- Every issue must have label, milestone, and project field.
- Every sprint task must have acceptance criteria.
- Every AI Agent task must include required context docs.
- Every PR must pass CI before merge.
- Every release must pass QA/UAT.
- Every production deployment must be approved.
```

---

## 28. Related Files

Repository support files:

```text
.github/CODEOWNERS
.github/pull_request_template.md
.github/workflows/ci.yml
.github/ISSUE_TEMPLATE/feature_task.yml
.github/ISSUE_TEMPLATE/bug_report.yml
.github/ISSUE_TEMPLATE/ai_agent_task.yml
.github/ISSUE_TEMPLATE/security_review.yml
.github/ISSUE_TEMPLATE/qa_uat.yml
```

Documentation files:

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/11-github-repository-rules.md
docs/27-workflow.md
docs/28-ai-agent-sprint-planning-prompts.md
```

## Alignment With PRD, Development Plan, and Workflow

This GitHub Project Management document must stay aligned with:

```text
docs/25-product-requirement-document.md
docs/26-development-plan.md
docs/27-workflow.md
```

Rules:

- If PRD scope changes, update labels/milestones/issues if impacted.
- If Development Plan sprint structure changes, update milestones, sprint labels, and GitHub Project fields/views.
- If Workflow changes, update issue lifecycle, PR workflow, QA/UAT workflow, release workflow, and hotfix workflow.
- If GitHub templates change, update this document and `docs/README.md`.

## Setup Guide Reference

The practical setup guide for GitHub repository, labels, milestones, GitHub Project fields/views, branch protection, and GitHub Environments is:

```text
docs/40-github-repository-setup-labels-project.md
```

Use this document when:

```text
- creating the repository
- creating labels
- creating milestones
- creating GitHub Project fields
- creating GitHub Project views
- configuring branch protection
- configuring staging/production GitHub Environments
```

Keep this GitHub Project Management guide aligned with the setup guide.

## Git Commit Convention Alignment

GitHub Issues, Pull Requests, and commits should be traceable through consistent naming.

Commit and PR title format:

```text
type(scope): short description
```

Alignment rules:

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

References:

```text
docs/08-coding-standard.md
docs/27-workflow.md
docs/40-github-repository-setup-labels-project.md
```

## GitHub Setup Scripts

Initial GitHub setup is supported by scripts in:

```text
scripts/github/
```

Use scripts for initial or repeatable setup:

```text
scripts/github/06-setup-labels.sh
scripts/github/07-setup-milestones.sh
scripts/github/08-setup-project.sh
scripts/github/09-setup-project-fields.sh
scripts/github/10-project-views-manual-guide.sh
```

Important:

- Labels and milestones should stay aligned with this GitHub Project Management document.
- Project fields should stay aligned with the fields defined in this document.
- Project views may require manual setup via GitHub UI; use `10-project-views-manual-guide.sh`.
