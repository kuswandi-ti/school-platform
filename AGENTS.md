# AGENTS.md

# School Platform AI Agent Rules

This file defines mandatory rules for AI Agents working on `school-platform`.

## Mandatory Source of Truth

AI Agent must treat these files as source of truth:

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/01-technical-architecture.md
docs/02-service-boundary.md
docs/03-data-model-mvp.md
docs/04-api-contract.md
docs/05-event-contract.md
docs/06-ui-screen-user-flow.md
docs/07-test-plan-acceptance-criteria.md
docs/08-coding-standard.md
docs/09-ai-agent-rules.md
docs/10-sprint-backlog-mvp.md
docs/11-github-repository-rules.md
docs/12-ai-agent-sprint-prompts.md
docs/13-sprint-0-task-prompts.md through docs/23-sprint-10-task-prompts.md
docs/24-local-development-guide.md
docs/25-prd-prompt.md
docs/26-development-plan-prompt.md
docs/27-workflow-prompt.md
docs/28-ai-agent-sprint-planning-prompts.md
```

## Planning Prompt Documents

Use these prompt documents before implementation:

| Document | Use When | Expected Output |
|---|---|---|
| `docs/25-prd-prompt.md` | Creating Product Requirement Document | `docs/25-product-requirement-document.md` |
| `docs/26-development-plan-prompt.md` | Creating Development Plan | `docs/26-development-plan.md` |
| `docs/27-workflow-prompt.md` | Creating Workflow/SOP | `docs/27-workflow.md` |
| `docs/28-ai-agent-sprint-planning-prompts.md` | Creating Sprint Plan documents | `docs/29-sprint-0-plan.md` through `docs/39-sprint-10-plan.md` |

These are not direct coding prompts.

## Planning vs Coding Prompt Rule

Use this flow:

```text
PRD
→ Development Plan
→ Workflow
→ Sprint Planning
→ GitHub Issues
→ Sprint Task Prompt
→ Coding
→ Pull Request
```

Rules:

- Use `docs/25-prd-prompt.md` for PRD generation.
- Use `docs/26-development-plan-prompt.md` for Development Plan generation.
- Use `docs/27-workflow-prompt.md` for Workflow/SOP generation.
- Use `docs/28-ai-agent-sprint-planning-prompts.md` for sprint planning.
- Use `docs/13` through `docs/23` sprint task prompt files only for implementation/coding.
- If a sprint plan exists, read it before coding.
- Do not use planning prompts as direct coding instructions.

## Mandatory Reading Order by Work Type

### PRD Work

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/25-prd-prompt.md
docs/01-technical-architecture.md
docs/02-service-boundary.md
docs/06-ui-screen-user-flow.md
docs/10-sprint-backlog-mvp.md
```

### Development Plan Work

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/26-development-plan-prompt.md
docs/01-technical-architecture.md
docs/02-service-boundary.md
docs/10-sprint-backlog-mvp.md
docs/11-github-repository-rules.md
docs/24-local-development-guide.md
```

### Workflow/SOP Work

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/27-workflow-prompt.md
docs/11-github-repository-rules.md
docs/24-local-development-guide.md
docs/09-ai-agent-rules.md
```

### Sprint Planning Work

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/28-ai-agent-sprint-planning-prompts.md
docs/10-sprint-backlog-mvp.md
active sprint task prompt
```

### Coding/Implementation Work

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/09-ai-agent-rules.md
docs/08-coding-standard.md
active sprint plan if available
active sprint task prompt from docs/13 through docs/23
task-specific technical documents
```

## Non-Negotiable Rules

- Do not query another service database.
- Do not put business logic in API Gateway.
- Do not skip permission/scope checks.
- Do not skip object-level authorization.
- Do not skip tests.
- Do not log tokens, passwords, or Confidential data.
- Do not use float for finance.
- Do not expose private files publicly.
- Do not implement features outside sprint scope.
- Do not use AI Agent for production secrets or final production approval.

## GitHub Project Management and Repository Support Files

AI Agent must recognize these GitHub management files as part of the project workflow source of truth:

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

Use `docs/25-github-project-management.md` when working on:

```text
- GitHub Labels
- GitHub Milestones
- GitHub Project fields/views
- Issue workflow
- PR workflow
- AI Agent issue lifecycle
- QA/UAT tracking
- release readiness tracking
```

Important numbering note:

```text
docs/25-prd-prompt.md
docs/25-github-project-management.md
```

Both files intentionally use prefix `25-` because they serve different purposes. Do not treat them as duplicates. `25-prd-prompt.md` is a prompt template for PRD generation. `25-github-project-management.md` is the GitHub project management guide.

## GitHub Management Work Reading Order

For GitHub repository/project management work, read:

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/11-github-repository-rules.md
docs/25-github-project-management.md
docs/27-workflow-prompt.md
.github/CODEOWNERS
.github/pull_request_template.md
.github/workflows/ci.yml
.github/ISSUE_TEMPLATE/
```

## Final MVP Planning Documents

The project now has final generated planning documents in addition to prompt templates.

Final documents:

```text
docs/25-product-requirement-document.md
docs/26-development-plan.md
docs/27-workflow.md
docs/25-github-project-management.md
```

Prompt generator documents:

```text
docs/25-prd-prompt.md
docs/26-development-plan-prompt.md
docs/27-workflow-prompt.md
docs/28-ai-agent-sprint-planning-prompts.md
```

Usage rule:

- Use `docs/25-product-requirement-document.md` as the product requirement source of truth.
- Use `docs/26-development-plan.md` as the MVP implementation planning source of truth.
- Use `docs/27-workflow.md` as the daily workflow/SOP source of truth.
- Use `docs/25-github-project-management.md` as the GitHub Project, Labels, Milestones, Issues, PR, QA/UAT, and release tracking source of truth.
- Use `docs/25-prd-prompt.md`, `docs/26-development-plan-prompt.md`, and `docs/27-workflow-prompt.md` only when regenerating or updating those final documents.
- Use `docs/28-ai-agent-sprint-planning-prompts.md` to generate sprint plan documents `docs/29` through `docs/39`.

## Updated Reading Order After Final Planning Docs

For product or requirement work, read:

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/25-product-requirement-document.md
docs/01-technical-architecture.md
docs/02-service-boundary.md
docs/03-data-model-mvp.md
docs/04-api-contract.md
docs/06-ui-screen-user-flow.md
docs/10-sprint-backlog-mvp.md
```

For development planning or sprint execution work, read:

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/25-product-requirement-document.md
docs/26-development-plan.md
docs/27-workflow.md
docs/25-github-project-management.md
docs/10-sprint-backlog-mvp.md
active sprint plan if available
active sprint task prompt
```

For GitHub workflow, issue, PR, CI/CD, QA/UAT, or release process work, read:

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/11-github-repository-rules.md
docs/25-github-project-management.md
docs/26-development-plan.md
docs/27-workflow.md
.github/CODEOWNERS
.github/pull_request_template.md
.github/workflows/ci.yml
.github/ISSUE_TEMPLATE/
```

For coding work, read:

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/25-product-requirement-document.md
docs/26-development-plan.md
docs/27-workflow.md
docs/09-ai-agent-rules.md
docs/08-coding-standard.md
active sprint plan if available
active sprint task prompt from docs/13 through docs/23
task-specific technical documents
```

## Numbering Note

There are multiple documents using prefix `25-`:

```text
docs/25-prd-prompt.md
docs/25-product-requirement-document.md
docs/25-github-project-management.md
```

They are not duplicates.

- `25-prd-prompt.md` is a prompt to generate the PRD.
- `25-product-requirement-document.md` is the final PRD.
- `25-github-project-management.md` is the GitHub Project Management guide.

## Final Sprint Plans and GitHub Setup Document

The project now has final Sprint Plan documents and a dedicated GitHub setup guide.

Final sprint planning documents:

```text
docs/29-sprint-0-plan.md
docs/30-sprint-1-plan.md
docs/31-sprint-2-plan.md
docs/32-sprint-3-plan.md
docs/33-sprint-4-plan.md
docs/34-sprint-5-plan.md
docs/35-sprint-6-plan.md
docs/36-sprint-7-plan.md
docs/37-sprint-8-plan.md
docs/38-sprint-9-plan.md
docs/39-sprint-10-plan.md
```

GitHub setup guide:

```text
docs/40-github-repository-setup-labels-project.md
```

Usage rule:

- Use `docs/29-sprint-0-plan.md` through `docs/39-sprint-10-plan.md` as sprint-level planning source of truth.
- Use `docs/13-sprint-0-task-prompts.md` through `docs/23-sprint-10-task-prompts.md` only for task-level implementation prompts.
- Use `docs/40-github-repository-setup-labels-project.md` when setting up repository, labels, milestones, GitHub Project fields/views, branch protection, and GitHub Environments.
- Use `docs/25-github-project-management.md` for ongoing GitHub Project management and workflow governance.

## Updated Reading Order After Sprint Plans and GitHub Setup Guide

For sprint execution work, read:

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/25-product-requirement-document.md
docs/26-development-plan.md
docs/27-workflow.md
docs/25-github-project-management.md
docs/40-github-repository-setup-labels-project.md
active sprint plan from docs/29 through docs/39
active sprint task prompt from docs/13 through docs/23
```

For GitHub repository setup work, read:

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/11-github-repository-rules.md
docs/25-github-project-management.md
docs/27-workflow.md
docs/40-github-repository-setup-labels-project.md
.github/CODEOWNERS
.github/pull_request_template.md
.github/workflows/ci.yml
.github/ISSUE_TEMPLATE/
```

For coding work, read:

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/25-product-requirement-document.md
docs/26-development-plan.md
docs/27-workflow.md
docs/25-github-project-management.md
active sprint plan from docs/29 through docs/39
active sprint task prompt from docs/13 through docs/23
task-specific architecture/API/event/data-model documents
```

## Git Commit Convention Reference

The official Git Commit Convention is now documented in:

```text
docs/08-coding-standard.md
docs/27-workflow.md
docs/40-github-repository-setup-labels-project.md
```

Required convention:

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

AI Agent rules:

- AI Agent-generated commit suggestions must follow the convention.
- AI Agent must not suggest vague commit messages such as `update file`, `fix bug`, or `changes`.
- If the task includes migration, security, or breaking change, AI Agent must include that context in the commit body or PR notes.
- PR title should follow the same convention.

## GitHub Setup Scripts Reference

The repository setup guide now includes executable helper scripts under:

```text
scripts/github/
```

Official script list:

```text
scripts/github/00-run-all-github-setup.sh
scripts/github/01-create-repository.sh
scripts/github/02-bootstrap-repository-files.sh
scripts/github/03-create-branches.sh
scripts/github/04-setup-branch-protection.sh
scripts/github/05-setup-environments.sh
scripts/github/06-setup-labels.sh
scripts/github/07-setup-milestones.sh
scripts/github/08-setup-project.sh
scripts/github/09-setup-project-fields.sh
scripts/github/10-project-views-manual-guide.sh
scripts/github/README.md
```

AI Agent rules:

- Do not execute scripts against a real GitHub repository without explicit user instruction.
- Do not create, read, or modify production secrets.
- Do not configure production deployment approval without human confirmation.
- When modifying repository setup, update `docs/40-github-repository-setup-labels-project.md` and `scripts/github/` together.
- If GitHub CLI behavior changes, document the limitation and prefer manual fallback steps in `docs/40`.
