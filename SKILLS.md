# SKILLS.md

# School Platform AI Agent Skills

This file defines operational workflows by task type.

## Product Documentation Skill

Use this skill when creating or updating the PRD.

Required documents:

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

Output:

```text
docs/25-product-requirement-document.md
```

Rules:

- Focus on product requirements, users, scope, user journeys, acceptance criteria, success metrics, and risks.
- Do not write implementation code.
- Do not add features outside MVP.
- Mark assumptions and open questions clearly.

## Development Planning Skill

Use this skill when creating or updating the Development Plan.

Required documents:

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

Output:

```text
docs/26-development-plan.md
```

Rules:

- Focus on sprint delivery, dependencies, quality gates, testing strategy, release plan, AI Agent usage, and risks.
- Do not write implementation code.
- Keep the plan aligned with Sprint 0 through Sprint 10.

## Workflow/SOP Skill

Use this skill when creating or updating workflow/SOP documentation.

Required documents:

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/27-workflow-prompt.md
docs/11-github-repository-rules.md
docs/24-local-development-guide.md
docs/09-ai-agent-rules.md
```

Output:

```text
docs/27-workflow.md
```

Rules:

- Focus on GitHub workflow, issue lifecycle, PR workflow, code review, QA, CI/CD, release, hotfix, documentation update rules, and AI Agent handoff.
- Do not write implementation code.
- Keep workflow consistent with protected branches and PR-only delivery.

## Sprint Planning Skill

Use this skill when creating Sprint Plan documents before GitHub issues are created.

Required documents:

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/28-ai-agent-sprint-planning-prompts.md
docs/10-sprint-backlog-mvp.md
active sprint task prompt from docs/13 through docs/23
```

Output examples:

```text
docs/29-sprint-0-plan.md
docs/30-sprint-1-plan.md
...
docs/39-sprint-10-plan.md
```

Rules:

- Produce sprint-level planning, not code.
- Include objective, scope, out of scope, user stories, task breakdown, dependencies, API/proto/event impact, data impact, permissions, audit, test plan, acceptance criteria, GitHub issue plan, and handoff notes.
- If sprint plan already exists, update it instead of duplicating it.

## Coding Prompt Usage Skill

Use this skill only when starting implementation.

Required documents:

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

Rules:

- Use `docs/13` through `docs/23` task prompt files for coding.
- Use only one task prompt at a time.
- Do not use `docs/25` through `docs/28` as direct coding instructions.
- If a sprint plan exists, use it for context, but implement only the selected task scope.
- End every coding response with summary, changed files, tests run, and risks/notes.

## GitHub Project Management Skill

Use this skill when creating or updating GitHub project management, repository governance, issue templates, PR templates, CI workflow, or related documentation.

Required documents:

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

Outputs may include:

```text
.github/CODEOWNERS
.github/pull_request_template.md
.github/workflows/ci.yml
.github/ISSUE_TEMPLATE/*.yml
docs/25-github-project-management.md
docs/README.md
docs/11-github-repository-rules.md
```

Rules:

- Keep GitHub workflow aligned with `feature/* → develop → staging → main`.
- Do not weaken branch protection requirements.
- Do not remove CI requirements.
- Do not remove QA/UAT sign-off for staging or production readiness.
- Do not use AI Agent for production secrets or final production approval.
- If labels, milestones, project fields, issue templates, or PR templates change, update `docs/25-github-project-management.md` and `docs/README.md`.

## Final Planning Documents Usage Skill

Use this skill when working after the PRD, Development Plan, and Workflow documents have been generated.

Required documents:

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/25-product-requirement-document.md
docs/26-development-plan.md
docs/27-workflow.md
docs/25-github-project-management.md
```

Rules:

- Treat `docs/25-product-requirement-document.md` as product scope and requirement source.
- Treat `docs/26-development-plan.md` as sprint/development execution source.
- Treat `docs/27-workflow.md` as daily SOP source.
- Treat `docs/25-github-project-management.md` as GitHub tracking and repository management source.
- Use prompt files only when regenerating or updating the final documents.
- If a task changes PRD scope, sprint/development plan, workflow, or GitHub tracking, update the corresponding final document in the same PR or create a linked documentation issue.

## Sprint Execution and GitHub Setup Usage Skill

Use this skill after Sprint Plan documents and the GitHub setup guide have been generated.

Required documents:

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

Rules:

- Use `docs/29` through `docs/39` for sprint-level scope, dependencies, issue plan, risks, and exit criteria.
- Use `docs/13` through `docs/23` for coding prompt details only after a GitHub issue is selected.
- Use `docs/40-github-repository-setup-labels-project.md` when creating or changing repository setup, GitHub Labels, GitHub Milestones, GitHub Project fields/views, branch protection, or GitHub Environments.
- If repository setup, labels, project fields, or workflow changes, update `docs/40-github-repository-setup-labels-project.md`, `docs/25-github-project-management.md`, `docs/27-workflow.md`, and `docs/README.md` as needed.

## Git Commit and PR Convention Skill

Use this skill whenever creating commit messages, PR titles, PR descriptions, release notes, or AI Agent implementation summaries.

Required references:

```text
docs/08-coding-standard.md
docs/27-workflow.md
docs/40-github-repository-setup-labels-project.md
.github/pull_request_template.md
```

Commit/PR title format:

```text
type(scope): short description
```

Allowed common types:

```text
feat
fix
docs
chore
refactor
test
ci
build
perf
security
revert
```

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

Rules:

- Use English for commit messages and PR titles.
- Keep subject short and specific.
- Do not mix unrelated changes in one commit.
- Security, migration, and breaking changes must include context in commit body or PR notes.
- Squash merge commit must also follow this convention.

## GitHub Setup Scripts Skill

Use this skill when creating or updating GitHub repository setup automation.

Required references:

```text
docs/40-github-repository-setup-labels-project.md
scripts/github/README.md
scripts/github/*.sh
docs/25-github-project-management.md
docs/27-workflow.md
docs/11-github-repository-rules.md
```

Rules:

- Keep documentation and scripts synchronized.
- Script names must be descriptive and ordered by setup sequence.
- Scripts must be safe to re-run where possible.
- Scripts must validate required environment variables.
- Scripts must not contain real tokens, private keys, database passwords, or production secrets.
- Scripts must include manual fallback notes when GitHub CLI/API limitations exist.
- Project view creation may remain manual if GitHub CLI/API support is unstable.
