# School Platform Documentation

Project: `school-platform`  
Purpose: Documentation index, reading guide, and AI Agent documentation map.

---

## Quick Links

- [Documentation Index](INDEX.md)
- [Product Requirement Document](25-product-requirement-document.md)
- [Development Plan](26-development-plan.md)
- [Workflow](27-workflow.md)
- [GitHub Project Management](25-github-project-management.md)
- [Sprint 0 Plan](29-sprint-0-plan.md)
- [Sprint 0 Task Prompts](13-sprint-0-task-prompts.md)

---

## 1. Overview

This folder contains the technical, product, planning, workflow, and AI Agent documentation for `school-platform`.

`school-platform` is an internal school foundation management system for TK, SD, SMP, and SMA.

The MVP supports:

```text
- Identity & Access
- School Core
- File Management + Import Excel
- PPDB
- Finance/SPP manual payment
- Academic Basic
- Report Card / E-Rapor Basic
- Communication / Notification
- Reporting Dashboard
- Security, Observability, Backup, and UAT Hardening
```

The documentation is designed to support:

```text
- Product planning
- Technical architecture
- Development planning
- GitHub workflow
- AI Agent implementation
- QA/UAT
- DevOps and deployment preparation
```

---

## 2. Root-Level AI Agent Files

Before reading detailed docs, AI Agent should read the root-level files:

```text
../AGENTS.md
../SKILLS.md
```

Purpose:

| File | Purpose |
|---|---|
| `../AGENTS.md` | Mandatory project rules for AI Agent |
| `../SKILLS.md` | Operational workflows by task type |

Recommended AI Agent reading order:

```text
1. AGENTS.md
2. SKILLS.md
3. docs/README.md
4. active task/sprint document
```

---

## Mandatory Before Coding

AI Agent and developers must read these documents before starting coding work:

```text
../AGENTS.md
../SKILLS.md
docs/README.md
25-product-requirement-document.md
26-development-plan.md
27-workflow.md
25-github-project-management.md
09-ai-agent-rules.md
08-coding-standard.md
active sprint plan
active sprint task prompt
```

For Sprint 0, the active references are:

```text
29-sprint-0-plan.md
13-sprint-0-task-prompts.md
```

Use [INDEX.md](INDEX.md) for the condensed document map.

---

## 3. Core Technical Documents

Read these first for architecture, service boundary, data model, API, event, UI, testing, and coding rules.

| File | Purpose |
|---|---|
| `01-technical-architecture.md` | Main technical architecture and system-wide decisions |
| `02-service-boundary.md` | Service ownership, data boundary, and cross-service rules |
| `03-data-model-mvp.md` | MVP data model per service |
| `04-api-contract.md` | External REST and internal gRPC API rules |
| `05-event-contract.md` | RabbitMQ event contract, envelope, naming, outbox, retry, and DLQ |
| `06-ui-screen-user-flow.md` | Web/mobile screen list and user flows |
| `07-test-plan-acceptance-criteria.md` | Test strategy, acceptance criteria, UAT, and QA rules |
| `08-coding-standard.md` | Coding conventions, Go service structure, frontend/mobile conventions |
| `09-ai-agent-rules.md` | Detailed rules for AI Agent implementation |
| `10-sprint-backlog-mvp.md` | MVP sprint roadmap |
| `11-github-repository-rules.md` | GitHub branch, PR, CI/CD, release, and hotfix rules |
| `12-ai-agent-sprint-prompts.md` | General AI Agent sprint prompts |
| `24-local-development-guide.md` | Local development setup guide |

---

## 4. Sprint Task Prompt Documents

These documents are used when AI Agent starts writing code for a specific task.

Use one task prompt at a time.

| Sprint | File | Focus |
|---|---|---|
| Sprint 0 | `13-sprint-0-task-prompts.md` | Project Foundation |
| Sprint 1 | `14-sprint-1-task-prompts.md` | Identity & Access |
| Sprint 2 | `15-sprint-2-task-prompts.md` | School Core |
| Sprint 3 | `16-sprint-3-task-prompts.md` | File Management + Import Excel |
| Sprint 4 | `17-sprint-4-task-prompts.md` | PPDB |
| Sprint 5 | `18-sprint-5-task-prompts.md` | Finance / SPP |
| Sprint 6 | `19-sprint-6-task-prompts.md` | Academic Basic |
| Sprint 7 | `20-sprint-7-task-prompts.md` | Report Card / E-Rapor Basic |
| Sprint 8 | `21-sprint-8-task-prompts.md` | Communication / Notification |
| Sprint 9 | `22-sprint-9-task-prompts.md` | Reporting Dashboard |
| Sprint 10 | `23-sprint-10-task-prompts.md` | Security, Observability, Backup, and UAT Hardening |

Usage:

```text
Copy one task prompt from the active sprint file
→ provide it to AI Agent
→ ask AI Agent to implement only that task
→ review output
→ run tests
→ create PR
```

---

## 5. PRD, Development Plan, Workflow, and Sprint Planning Prompts

These documents are prompt templates for generating higher-level planning documents.

| File | Purpose | Expected Output |
|---|---|---|
| `25-prd-prompt.md` | Prompt for AI Agent to create Product Requirement Document | `docs/25-product-requirement-document.md` |
| `26-development-plan-prompt.md` | Prompt for AI Agent to create Development Plan | `docs/26-development-plan.md` |
| `27-workflow-prompt.md` | Prompt for AI Agent to create Workflow / SOP document | `docs/27-workflow.md` |
| `28-ai-agent-sprint-planning-prompts.md` | Prompt pack for AI Agent to create Sprint Plan documents | `docs/29-sprint-0-plan.md` through `docs/39-sprint-10-plan.md` |

These files do **not** instruct AI Agent to write code directly.

They are used to generate planning documents before coding starts.

---

## 6. Planning-to-Implementation Flow

Recommended project documentation flow:

```text
25-prd-prompt.md
→ docs/25-product-requirement-document.md

26-development-plan-prompt.md
→ docs/26-development-plan.md

27-workflow-prompt.md
→ docs/27-workflow.md

28-ai-agent-sprint-planning-prompts.md
→ docs/29-sprint-0-plan.md
→ docs/30-sprint-1-plan.md
→ docs/31-sprint-2-plan.md
→ ...
→ docs/39-sprint-10-plan.md

Sprint Plan
→ GitHub Issues
→ Sprint Task Prompt
→ AI Agent Coding
→ Pull Request
→ Review
→ QA
→ Merge
```

Simplified flow:

```text
PRD
→ Development Plan
→ Workflow
→ Sprint Planning
→ GitHub Issues
→ Task Prompt
→ Coding
→ Pull Request
```

---

## 7. Difference Between Prompt Types

| Document Type | Purpose | Used For Coding? |
|---|---|---|
| `25-prd-prompt.md` | Generate PRD | No |
| `26-development-plan-prompt.md` | Generate Development Plan | No |
| `27-workflow-prompt.md` | Generate Workflow/SOP | No |
| `28-ai-agent-sprint-planning-prompts.md` | Generate Sprint Plan per sprint | No |
| `12-ai-agent-sprint-prompts.md` | General sprint direction | Indirect |
| `13–23-sprint-*-task-prompts.md` | Detailed implementation task prompts | Yes |

Coding starts from:

```text
13-sprint-0-task-prompts.md
14-sprint-1-task-prompts.md
15-sprint-2-task-prompts.md
...
23-sprint-10-task-prompts.md
```

---

## 8. Recommended Reading Order by Role

## 8.1 Product Owner / Reviewer

```text
docs/README.md
25-prd-prompt.md
26-development-plan-prompt.md
27-workflow-prompt.md
10-sprint-backlog-mvp.md
06-ui-screen-user-flow.md
07-test-plan-acceptance-criteria.md
```

## 8.2 Backend Developer

```text
../AGENTS.md
../SKILLS.md
01-technical-architecture.md
02-service-boundary.md
03-data-model-mvp.md
04-api-contract.md
05-event-contract.md
08-coding-standard.md
09-ai-agent-rules.md
10-sprint-backlog-mvp.md
active sprint task prompt
```

## 8.3 Frontend Developer

```text
../AGENTS.md
../SKILLS.md
01-technical-architecture.md
04-api-contract.md
06-ui-screen-user-flow.md
08-coding-standard.md
09-ai-agent-rules.md
active sprint task prompt
```

## 8.4 Mobile Developer

```text
../AGENTS.md
../SKILLS.md
01-technical-architecture.md
04-api-contract.md
06-ui-screen-user-flow.md
08-coding-standard.md
09-ai-agent-rules.md
24-local-development-guide.md
active sprint task prompt
```

## 8.5 QA

```text
01-technical-architecture.md
06-ui-screen-user-flow.md
07-test-plan-acceptance-criteria.md
10-sprint-backlog-mvp.md
11-github-repository-rules.md
27-workflow-prompt.md
active sprint plan
active sprint task prompt
```

## 8.6 Infrastructure / DevOps

```text
01-technical-architecture.md
08-coding-standard.md
10-sprint-backlog-mvp.md
11-github-repository-rules.md
24-local-development-guide.md
13-sprint-0-task-prompts.md
23-sprint-10-task-prompts.md
```

## 8.7 AI Agent

```text
../AGENTS.md
../SKILLS.md
09-ai-agent-rules.md
08-coding-standard.md
active sprint plan
active sprint task prompt
task-specific technical doc
```

---

## 9. Recommended Reading Order by Task Type

## 9.1 Product Requirement Document Task

```text
../AGENTS.md
docs/README.md
25-prd-prompt.md
01-technical-architecture.md
02-service-boundary.md
06-ui-screen-user-flow.md
10-sprint-backlog-mvp.md
```

## 9.2 Development Plan Task

```text
../AGENTS.md
docs/README.md
26-development-plan-prompt.md
01-technical-architecture.md
02-service-boundary.md
10-sprint-backlog-mvp.md
11-github-repository-rules.md
24-local-development-guide.md
```

## 9.3 Workflow / SOP Task

```text
../AGENTS.md
../SKILLS.md
docs/README.md
27-workflow-prompt.md
11-github-repository-rules.md
24-local-development-guide.md
```

## 9.4 Sprint Planning Task

```text
../AGENTS.md
../SKILLS.md
docs/README.md
28-ai-agent-sprint-planning-prompts.md
10-sprint-backlog-mvp.md
active sprint task prompt
```

## 9.5 Backend Service Task

```text
../AGENTS.md
../SKILLS.md
02-service-boundary.md
03-data-model-mvp.md
08-coding-standard.md
09-ai-agent-rules.md
active sprint task prompt
```

## 9.6 API Task

```text
../AGENTS.md
../SKILLS.md
04-api-contract.md
08-coding-standard.md
09-ai-agent-rules.md
active sprint task prompt
```

## 9.7 Event Task

```text
../AGENTS.md
../SKILLS.md
05-event-contract.md
08-coding-standard.md
09-ai-agent-rules.md
active sprint task prompt
```

## 9.8 UI Task

```text
../AGENTS.md
../SKILLS.md
06-ui-screen-user-flow.md
04-api-contract.md
08-coding-standard.md
active sprint task prompt
```

## 9.9 Testing Task

```text
../AGENTS.md
../SKILLS.md
07-test-plan-acceptance-criteria.md
09-ai-agent-rules.md
active sprint task prompt
```

## 9.10 GitHub / CI/CD Task

```text
../AGENTS.md
../SKILLS.md
11-github-repository-rules.md
24-local-development-guide.md
27-workflow-prompt.md
```

---

## 10. MVP Sprint Order

Implementation order:

```text
Sprint 0  : Project Foundation
Sprint 1  : Identity & Access
Sprint 2  : School Core
Sprint 3  : File Management + Import Excel
Sprint 4  : PPDB
Sprint 5  : Finance / SPP
Sprint 6  : Academic Basic
Sprint 7  : Report Card / E-Rapor Basic
Sprint 8  : Communication / Notification
Sprint 9  : Reporting Dashboard
Sprint 10 : Security, Observability, Backup, and UAT Hardening
```

Do not implement later sprint features early unless explicitly approved.

---

## 11. Suggested Future Generated Documents

The prompt files 25–28 are used to generate these final documents:

| Planned File | Generated From | Purpose |
|---|---|---|
| `25-product-requirement-document.md` | `25-prd-prompt.md` | Final PRD |
| `26-development-plan.md` | `26-development-plan-prompt.md` | Final Development Plan |
| `27-workflow.md` | `27-workflow-prompt.md` | Final Workflow / SOP |
| `29-sprint-0-plan.md` | `28-ai-agent-sprint-planning-prompts.md` | Sprint 0 Plan |
| `30-sprint-1-plan.md` | `28-ai-agent-sprint-planning-prompts.md` | Sprint 1 Plan |
| `31-sprint-2-plan.md` | `28-ai-agent-sprint-planning-prompts.md` | Sprint 2 Plan |
| `32-sprint-3-plan.md` | `28-ai-agent-sprint-planning-prompts.md` | Sprint 3 Plan |
| `33-sprint-4-plan.md` | `28-ai-agent-sprint-planning-prompts.md` | Sprint 4 Plan |
| `34-sprint-5-plan.md` | `28-ai-agent-sprint-planning-prompts.md` | Sprint 5 Plan |
| `35-sprint-6-plan.md` | `28-ai-agent-sprint-planning-prompts.md` | Sprint 6 Plan |
| `36-sprint-7-plan.md` | `28-ai-agent-sprint-planning-prompts.md` | Sprint 7 Plan |
| `37-sprint-8-plan.md` | `28-ai-agent-sprint-planning-prompts.md` | Sprint 8 Plan |
| `38-sprint-9-plan.md` | `28-ai-agent-sprint-planning-prompts.md` | Sprint 9 Plan |
| `39-sprint-10-plan.md` | `28-ai-agent-sprint-planning-prompts.md` | Sprint 10 Plan |

---

## 12. Documentation Update Rules

Update documentation when implementation changes contract, architecture, workflow, or planning.

| Change | Document to Update |
|---|---|
| API endpoint changed | `04-api-contract.md` and `packages/openapi` |
| gRPC proto changed | `packages/proto` |
| Event added/changed | `05-event-contract.md` and `packages/events` |
| Data model changed | `03-data-model-mvp.md` |
| UI screen changed | `06-ui-screen-user-flow.md` |
| Testing strategy changed | `07-test-plan-acceptance-criteria.md` |
| Coding convention changed | `08-coding-standard.md` |
| AI Agent rule changed | `09-ai-agent-rules.md`, `../AGENTS.md`, or `../SKILLS.md` |
| Sprint plan changed | `10-sprint-backlog-mvp.md` and active sprint plan |
| GitHub workflow changed | `11-github-repository-rules.md` and `27-workflow.md` |
| Local setup changed | `24-local-development-guide.md` |
| New prompt document added | `docs/README.md` |
| PRD changed | `25-product-requirement-document.md` |
| Development Plan changed | `26-development-plan.md` |
| Workflow changed | `27-workflow.md` |

---

## 13. Non-Negotiable Reminders

```text
- Do not query another service database.
- Do not put business logic in API Gateway.
- Do not skip permission/scope checks.
- Do not skip object-level authorization.
- Do not skip tests.
- Do not log tokens/passwords/Confidential data.
- Do not use float for finance.
- Do not expose private files publicly.
- Do not implement features outside sprint scope.
- Do not use AI Agent for production secrets or final production approval.
```

---

## 14. Final Notes

This documentation set is the source of truth for MVP planning and implementation.

When in doubt:

```text
1. Check AGENTS.md.
2. Check SKILLS.md.
3. Check docs/README.md.
4. Check the relevant technical document.
5. Check the active sprint plan.
6. Check the active sprint task prompt.
7. Keep the task small.
```

Recommended execution order:

```text
PRD
→ Development Plan
→ Workflow
→ Sprint Plan
→ GitHub Issues
→ AI Agent Task Prompt
→ Implementation
→ Pull Request
→ Review
→ QA
→ Merge
```

## GitHub Setup and Project Management

The repository now includes GitHub governance and tracking files.

| File | Purpose |
|---|---|
| `25-github-project-management.md` | GitHub Labels, Milestones, Project fields/views, issue lifecycle, PR tracking, QA/UAT tracking, AI Agent task tracking |
| `.github/CODEOWNERS` | Code ownership and automatic reviewer assignment |
| `.github/pull_request_template.md` | Pull Request checklist and review template |
| `.github/workflows/ci.yml` | CI workflow for repository checks, YAML checks, Go, Web, Mobile, and Docker Compose validation |
| `.github/ISSUE_TEMPLATE/feature_task.yml` | Feature/task issue template |
| `.github/ISSUE_TEMPLATE/bug_report.yml` | Bug report issue template |
| `.github/ISSUE_TEMPLATE/ai_agent_task.yml` | AI Agent task issue template |
| `.github/ISSUE_TEMPLATE/security_review.yml` | Security/privacy review issue template |
| `.github/ISSUE_TEMPLATE/qa_uat.yml` | QA/UAT issue template |

GitHub setup flow:

```text
CODEOWNERS
→ Pull Request Template
→ CI Workflow
→ Issue Templates
→ Labels
→ Milestones
→ GitHub Project
→ Sprint Issues
→ PR Review
→ QA/UAT
→ Release Readiness
```

Numbering note:

```text
25-prd-prompt.md
25-github-project-management.md
```

Both files are valid and intentionally separate:

- `25-prd-prompt.md` is a prompt template for generating the PRD.
- `25-github-project-management.md` is a GitHub project management guide.


Reference `docs/25-github-project-management.md` for GitHub labels, milestones, project fields/views, issue templates, PR tracking, QA/UAT tracking, and release readiness.

## Final MVP Planning Documents

The following final planning documents have been generated and should be used as active source-of-truth references.

| File | Type | Purpose |
|---|---|---|
| `25-product-requirement-document.md` | Final document | Product Requirement Document / PRD |
| `26-development-plan.md` | Final document | MVP technical delivery and development plan |
| `27-workflow.md` | Final document | Daily workflow / SOP |
| `25-github-project-management.md` | Final document | GitHub Project, labels, milestones, issues, PR, QA/UAT, release readiness |

Prompt files remain available for regeneration or updates:

| File | Type | Purpose |
|---|---|---|
| `25-prd-prompt.md` | Prompt template | Generate/update PRD |
| `26-development-plan-prompt.md` | Prompt template | Generate/update Development Plan |
| `27-workflow-prompt.md` | Prompt template | Generate/update Workflow/SOP |
| `28-ai-agent-sprint-planning-prompts.md` | Prompt template | Generate Sprint 0–10 plan documents |

Recommended reading order for implementation work:

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/25-product-requirement-document.md
docs/26-development-plan.md
docs/27-workflow.md
docs/25-github-project-management.md
docs/09-ai-agent-rules.md
docs/08-coding-standard.md
active sprint plan if available
active sprint task prompt
```

Numbering note:

```text
25-prd-prompt.md
25-product-requirement-document.md
25-github-project-management.md
```

These files intentionally share prefix `25-` but have different purposes.

## Sprint Plan Documents

Final sprint planning documents are available as the implementation planning source for Sprint 0 through Sprint 10.

| File | Sprint | Purpose |
|---|---|---|
| `29-sprint-0-plan.md` | Sprint 0 | Project Foundation |
| `30-sprint-1-plan.md` | Sprint 1 | Identity & Access |
| `31-sprint-2-plan.md` | Sprint 2 | School Core |
| `32-sprint-3-plan.md` | Sprint 3 | File Management + Import Excel |
| `33-sprint-4-plan.md` | Sprint 4 | PPDB |
| `34-sprint-5-plan.md` | Sprint 5 | Finance / SPP |
| `35-sprint-6-plan.md` | Sprint 6 | Academic Basic |
| `36-sprint-7-plan.md` | Sprint 7 | Report Card / E-Rapor Basic |
| `37-sprint-8-plan.md` | Sprint 8 | Communication / Notification |
| `38-sprint-9-plan.md` | Sprint 9 | Reporting Dashboard |
| `39-sprint-10-plan.md` | Sprint 10 | Security, Observability, Backup, UAT Hardening |

Usage:

```text
Sprint Plan Document
→ GitHub Issues
→ Sprint Task Prompt
→ Implementation Branch
→ Pull Request
→ QA/UAT
```

## GitHub Repository Setup Guide

Use this document for repository setup, labels, milestones, GitHub Project fields/views, branch protection, and GitHub Environments:

```text
40-github-repository-setup-labels-project.md
```

Relationship with other GitHub documents:

```text
25-github-project-management.md = ongoing project management guide
40-github-repository-setup-labels-project.md = practical repository setup and labels/project setup guide
27-workflow.md = daily SOP and workflow execution guide
```

## Git Commit Convention

The Git Commit Convention has been added as an official engineering standard.

Primary references:

| File | Purpose |
|---|---|
| `08-coding-standard.md` | Engineering standard for commit format |
| `27-workflow.md` | Daily workflow usage for commit and PR |
| `40-github-repository-setup-labels-project.md` | GitHub setup alignment with labels and PR workflow |
| `.github/pull_request_template.md` | PR checklist and validation |

Required format:

```text
type(scope): short description
```

Examples:

```text
feat(identity): add refresh token rotation
fix(finance): prevent duplicate bill generation
docs(workflow): add git commit convention
chore(ci): add repository validation workflow
test(academic): add attendance scope tests
```

Use the same format for PR title.

## GitHub Setup Scripts

The GitHub repository setup guide now has supporting scripts.

Primary document:

```text
40-github-repository-setup-labels-project.md
```

Supporting scripts:

| Script | Purpose |
|---|---|
| `../scripts/github/00-run-all-github-setup.sh` | Run setup sequence |
| `../scripts/github/01-create-repository.sh` | Create private repository |
| `../scripts/github/02-bootstrap-repository-files.sh` | Add README, `.gitignore`, CODEOWNERS, PR template, issue templates, CI |
| `../scripts/github/03-create-branches.sh` | Create `develop` and `staging`, set `main` as default |
| `../scripts/github/04-setup-branch-protection.sh` | Configure branch protection |
| `../scripts/github/05-setup-environments.sh` | Configure staging/production environments |
| `../scripts/github/06-setup-labels.sh` | Create/update labels |
| `../scripts/github/07-setup-milestones.sh` | Create milestones Sprint 0–10 and MVP Release |
| `../scripts/github/08-setup-project.sh` | Create GitHub Project |
| `../scripts/github/09-setup-project-fields.sh` | Create GitHub Project fields where supported |
| `../scripts/github/10-project-views-manual-guide.sh` | Print manual guide for Project views |
| `../scripts/github/README.md` | Script usage guide |

Recommended usage:

```bash
export GITHUB_OWNER="<org-or-user>"
export REPO_NAME="school-platform"
export PRODUCTION_REVIEWER_USER="kuswandi-ti"

bash scripts/github/00-run-all-github-setup.sh
```
