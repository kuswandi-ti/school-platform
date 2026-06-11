# 11 — GitHub Repository Rules and Branch Protection

Project: `school-platform`  
Status: Final decision for MVP  
Scope: GitHub repository workflow, branch rules, pull request rules, CI checks, environments, release, and hotfix process.

---

## 1. Purpose

This document defines the GitHub repository rules for the `school-platform` project.

The goals are:

- keep code changes controlled
- prevent direct production changes
- enforce CI before merge
- support staging QA/UAT
- protect production branch
- separate staging and production secrets
- make AI Agent output reviewable through Pull Requests
- keep release and hotfix workflow traceable

---

## 2. Core Decision

The repository uses three protected main branches:

```text
develop
staging
main
```

Workflow:

```text
feature/* → develop → staging → main/production
```

All changes must go through:

```text
Pull Request → CI → Review → Merge
```

Production deployment only happens from:

```text
main
```

Production deployment must require:

```text
manual approval
```

---

## 3. Branch Roles

## 3.1 develop

Purpose:

```text
Daily development integration branch.
```

Used for:

```text
- integrating feature branches
- running CI
- developer-level validation
- early integration testing
```

Allowed source branches:

```text
feature/*
fix/*
chore/*
docs/*
refactor/*
```

Rules:

```text
No direct push.
All changes via Pull Request.
CI must pass.
At least 1 approval required.
Force push disabled.
Branch deletion disabled.
```

---

## 3.2 staging

Purpose:

```text
Release candidate branch for QA, staging deployment, and UAT.
```

Used for:

```text
- staging deployment
- QA testing
- UAT
- regression testing
- release candidate validation
```

Allowed source branch:

```text
develop
```

Allowed exception:

```text
hotfix/* only when needed and approved.
```

Rules:

```text
No direct push.
All changes via Pull Request.
CI must pass.
QA sign-off required before merge to main.
Force push disabled.
Branch deletion disabled.
```

---

## 3.3 main

Purpose:

```text
Production branch.
```

Used for:

```text
- production deployment
- release tagging
- production rollback reference
```

Allowed source branch:

```text
staging
```

Allowed exception:

```text
hotfix/* for urgent production fix.
```

Rules:

```text
No direct push.
All changes via Pull Request.
CI must pass.
QA sign-off required.
Infrastructure/DevOps approval required.
Production environment manual approval required.
Force push disabled.
Branch deletion disabled.
```

---

## 4. Branch Protection Rules

## 4.1 Protection for develop

Recommended GitHub branch protection:

```text
Require a pull request before merging
Require at least 1 approval
Dismiss stale approvals when new commits are pushed
Require status checks to pass before merging
Require branches to be up to date before merging
Restrict force pushes
Do not allow branch deletion
```

Required checks:

```text
backend-go-test
backend-go-lint
frontend-web-lint
frontend-web-typecheck
frontend-web-build
flutter-analyze-test
proto-check if proto changes
docker-build if service/app changes
```

---

## 4.2 Protection for staging

Recommended GitHub branch protection:

```text
Require a pull request before merging
Require at least 1 approval
Require QA sign-off
Require status checks to pass before merging
Require branches to be up to date before merging
Restrict force pushes
Do not allow branch deletion
```

Additional rule:

```text
Only release candidate changes from develop should be merged into staging.
```

---

## 4.3 Protection for main

Recommended GitHub branch protection:

```text
Require a pull request before merging
Require at least 1–2 approvals
Require QA sign-off
Require Infrastructure/DevOps approval
Require status checks to pass before merging
Require production deployment approval through GitHub Environment
Require branches to be up to date before merging
Restrict who can push
Restrict force pushes
Do not allow branch deletion
```

Additional rule:

```text
main must always represent production-ready code.
```

---

## 5. Branch Naming Convention

Use these branch prefixes:

```text
feature/<area>-<short-description>
fix/<area>-<short-description>
hotfix/<short-description>
chore/<short-description>
docs/<short-description>
refactor/<area>-<short-description>
test/<area>-<short-description>
```

Examples:

```text
feature/finance-generate-bills
feature/admission-applicant-verification
feature/academic-report-card-publish
fix/identity-refresh-token-rotation
fix/academic-grade-scope-check
hotfix/payment-verification-error
chore/docker-compose-observability
docs/api-contract-finance
refactor/school-core-student-service
test/finance-payment-verification
```

Rules:

```text
Use English branch names.
Use lowercase.
Use hyphen separator.
Keep branch name short but descriptive.
```

---

## 6. Commit Message Convention

Use simple Conventional Commits.

Format:

```text
<type>(optional-scope): <message>
```

Types:

```text
feat
fix
chore
docs
refactor
test
ci
build
perf
security
```

Examples:

```text
feat(finance): add manual payment verification
fix(academic): prevent teacher from editing published grade
chore(infra): add staging deploy workflow
docs(api): update finance payment contract
test(school-core): add student scope tests
security(identity): add refresh token reuse detection
```

Rules:

```text
Use English.
Use present tense.
Keep message clear.
Do not include secrets or sensitive data in commit message.
```

---

## 7. Pull Request Rules

Every Pull Request must include:

```text
Summary
Affected area
What changed
How to test
Screenshots/video for UI changes
Migration notes if any
API/proto/event contract changes if any
Risk and rollback notes if relevant
```

Minimum PR checklist:

```text
- [ ] Code follows service boundary
- [ ] No cross-service database query
- [ ] Permission/scope checked
- [ ] Object-level authorization checked if resource by ID
- [ ] Audit log added for sensitive action
- [ ] Event published if required
- [ ] Tests added/updated
- [ ] OpenAPI/proto/event docs updated if contract changed
- [ ] No sensitive data in logs
- [ ] Migration is service-owned and reviewed
```

---

## 8. Pull Request Template

Suggested file:

```text
.github/pull_request_template.md
```

Template:

```md
## Summary

Describe what changed.

## Affected Area

- [ ] Backend
- [ ] API Gateway
- [ ] Frontend Web
- [ ] Mobile
- [ ] Infrastructure
- [ ] Database Migration
- [ ] API Contract
- [ ] gRPC/Proto Contract
- [ ] Event Contract
- [ ] Security/Permission
- [ ] Documentation

## Changes

- 
- 
- 

## How to Test

```bash
# commands
```

## Checklist

- [ ] CI passes
- [ ] Tests added/updated
- [ ] Permission/scope checked
- [ ] Object-level authorization checked
- [ ] Audit log added for sensitive action
- [ ] Event published if required
- [ ] OpenAPI/proto/event docs updated if needed
- [ ] Migration reviewed if any
- [ ] No secrets committed
- [ ] No sensitive data logged

## Screenshots / Video

Attach if UI changed.

## Risk

Describe risk.

## Rollback Plan

Describe rollback plan if relevant.
```

---

## 9. CODEOWNERS

Use CODEOWNERS to assign reviewers by area.

Suggested file:

```text
.github/CODEOWNERS
```

Example:

```text
/services/                       @backend-team
/services/api-gateway/            @backend-team @infra-team
/services/identity-service/        @backend-team
/services/school-core-service/     @backend-team
/services/admission-service/       @backend-team
/services/academic-service/        @backend-team
/services/finance-service/         @backend-team
/services/communication-service/   @backend-team
/services/reporting-service/       @backend-team

/packages/proto/                  @backend-team @frontend-team
/packages/openapi/                @backend-team @frontend-team @qa-team
/packages/events/                 @backend-team @qa-team

/apps/web-admin/                  @frontend-team
/apps/mobile-app/                 @frontend-team

/infra/                           @infra-team
/deploy/                          @infra-team
/.github/workflows/               @infra-team

/docs/                            @qa-team
```

If GitHub teams are not available yet, use individual usernames.

Rules:

```text
Backend changes require backend review.
Frontend changes require frontend review.
Infrastructure/deployment changes require Infrastructure/DevOps review.
API/proto changes require backend + frontend review.
Testing/UAT documents should involve QA.
```

---

## 10. Required CI Checks

CI must be path-aware where possible.

Recommended checks:

```text
backend-go-test
backend-go-lint
frontend-web-lint
frontend-web-typecheck
frontend-web-build
flutter-analyze-test
proto-check
openapi-check
event-schema-check
docker-build
migration-check
secret-scan
```

Examples:

### Backend service changed

Required:

```text
backend-go-test
backend-go-lint
docker-build for changed service
migration-check if migration exists
```

### Web app changed

Required:

```text
frontend-web-lint
frontend-web-typecheck
frontend-web-build
```

### Mobile app changed

Required:

```text
flutter-analyze-test
```

### Proto changed

Required:

```text
proto-check
affected backend tests
API Gateway tests
frontend contract review
```

### OpenAPI changed

Required:

```text
openapi-check
frontend API client compatibility check if available
QA review
```

### Event schema changed

Required:

```text
event-schema-check
affected consumer tests
```

---

## 11. GitHub Environments

Use GitHub Environments:

```text
staging
production
```

Optional:

```text
development
```

---

## 11.1 Staging Environment

Rules:

```text
Deploy from staging branch.
Use staging secrets only.
Can deploy automatically after merge to staging.
Used for QA/UAT.
```

Recommended protection:

```text
Optional approval by QA/Infra for staging deploy.
```

---

## 11.2 Production Environment

Rules:

```text
Deploy only from main branch.
Use production secrets only.
Requires manual approval.
Approver: Infrastructure/DevOps or owner teknis.
```

Production secrets must not be available to staging workflows.

Production deployment must not run from:

```text
develop
staging
feature/*
fix/*
```

---

## 12. Secrets Management in GitHub

Rules:

```text
Do not commit .env files.
Do not commit private keys.
Do not commit production credentials.
Use GitHub Environment Secrets.
Separate staging and production secrets.
Production secrets only available to production environment.
```

Recommended secrets:

```text
STAGING_HOST
STAGING_USER
STAGING_SSH_KEY
STAGING_DATABASE_URL
STAGING_RABBITMQ_URL
STAGING_REDIS_URL
STAGING_MINIO_ACCESS_KEY
STAGING_MINIO_SECRET_KEY

PRODUCTION_HOST
PRODUCTION_USER
PRODUCTION_SSH_KEY
PRODUCTION_DATABASE_URL
PRODUCTION_RABBITMQ_URL
PRODUCTION_REDIS_URL
PRODUCTION_MINIO_ACCESS_KEY
PRODUCTION_MINIO_SECRET_KEY
```

Enable if available:

```text
Secret scanning
Dependabot alerts
Dependabot security updates
```

---

## 13. Deployment Rules

### Staging Deployment

Trigger:

```text
Merge to staging
```

Flow:

```text
CI pass
→ build Docker images
→ push to container registry
→ deploy staging
→ run health checks
→ QA/UAT
```

### Production Deployment

Trigger:

```text
Merge to main
```

Flow:

```text
CI pass
→ build Docker images
→ push to container registry
→ wait for GitHub Environment production approval
→ deploy production
→ run health checks
→ create release tag
```

Production deploy must be manually approved.

---

## 14. Container Image Tagging

Use commit SHA for traceability.

Recommended image tags:

```text
ghcr.io/<org>/school-api-gateway:<commit-sha>
ghcr.io/<org>/school-identity-service:<commit-sha>
ghcr.io/<org>/school-finance-service:<commit-sha>
```

Optional additional tags:

```text
staging-latest
production-latest
v0.1.0
v1.0.0
```

Rule:

```text
Production deployment should reference immutable commit SHA tag.
```

---

## 15. Merge Strategy

Recommended merge strategy:

```text
feature/* → develop     : squash merge
develop → staging       : merge commit
staging → main          : merge commit
hotfix/* → main         : merge commit
```

Reason:

```text
Feature history stays clean in develop.
Release history remains traceable in staging/main.
```

---

## 16. Release Tagging

Every production release must have a tag.

Format:

```text
vMAJOR.MINOR.PATCH
```

Examples:

```text
v0.1.0
v0.2.0
v1.0.0
```

MVP internal examples:

```text
v0.1.0-mvp
v0.2.0-staging
v1.0.0
```

Release notes must include:

```text
new features
bug fixes
database migrations
known issues
rollback notes
```

---

## 17. Hotfix Workflow

Used only for critical production issues.

Flow:

```text
main
→ create hotfix/*
→ fix issue
→ PR to main
→ CI pass
→ review
→ production manual approval
→ deploy production
→ back-merge main to staging
→ back-merge staging to develop
```

Example branch:

```text
hotfix/payment-verification-error
```

Rules:

```text
Hotfix still uses PR unless extreme emergency.
Hotfix must be back-merged to staging and develop.
Hotfix must include test if possible.
Hotfix release must be tagged.
```

---

## 18. Release Candidate Workflow

Regular release flow:

```text
feature/* → develop
develop → staging
QA/UAT on staging
staging → main
production approval
production deploy
release tag
```

Before staging → main:

```text
CI pass
QA sign-off
No Critical/High bug
Migration reviewed
Backup/snapshot plan if migration is risky
Release notes prepared
Rollback plan prepared
```

---

## 19. Bug Workflow

Bug severity:

```text
Critical
High
Medium
Low
```

Flow:

```text
QA/User reports bug
→ triage severity
→ assign owner
→ fix branch
→ PR
→ CI
→ review
→ QA verify
→ close bug
```

Rules:

```text
Critical/High bugs in MVP core flows block production release.
Medium/Low may be released only if accepted by QA/Product Owner.
```

---

## 20. AI Agent Workflow Rules

AI Agent must work through branches and PR-ready changes.

AI Agent task output must include:

```text
summary of changes
files created/modified
tests added
how to run tests
migration notes if any
contract updates if any
known limitations
```

AI Agent must not:

```text
commit directly to protected branches
skip tests
change architecture without instruction
create huge multi-module changes without task scope
modify unrelated files
```

Good AI Agent task:

```text
Implement fee type CRUD in finance-service with migration, sqlc query, usecase, API mapping, permission check, audit log, and tests.
```

Bad AI Agent task:

```text
Build all finance features.
```

---

## 21. Documentation Update Rules

If a PR changes contract or architecture, update related docs.

Examples:

```text
API endpoint changed        → docs/04-api-contract.md and packages/openapi
Event added/changed         → docs/05-event-contract.md and packages/events
Data model changed          → docs/03-data-model-mvp.md
Service boundary changed    → docs/02-service-boundary.md
Sprint/task changed         → docs/10-sprint-backlog-mvp.md
GitHub workflow changed     → docs/11-github-repository-rules.md
```

---

## 22. Repository Security Rules

Repository must include:

```text
.gitignore
.env.example
.env.local.example
README.md
CODEOWNERS
Pull Request template
GitHub Actions workflows
```

`.gitignore` must exclude:

```text
.env
.env.*
!.env.example
node_modules
dist
build
coverage
*.log
tmp
.DS_Store
private keys
```

Rules:

```text
No production secrets in repo.
No production data dump in repo.
No student personal data sample in repo unless anonymized/synthetic.
No backup files in repo.
```

---

## 23. Minimum Repository Setup Checklist

Before coding Sprint 0 is considered done:

```text
- [ ] Branches develop, staging, main exist
- [ ] Branch protection enabled for develop
- [ ] Branch protection enabled for staging
- [ ] Branch protection enabled for main
- [ ] Pull Request required
- [ ] CI required checks configured
- [ ] Force push disabled for protected branches
- [ ] Delete branch disabled for protected branches
- [ ] GitHub Environments staging and production created
- [ ] Production environment requires approval
- [ ] Secrets separated by environment
- [ ] CODEOWNERS added
- [ ] PR template added
- [ ] Basic CI workflow added
- [ ] Container registry configured
```

---

## 24. Final Summary

Repository rules:

```text
develop = feature integration
staging = QA/UAT release candidate
main = production
```

Workflow:

```text
feature/* → develop → staging → main/production
```

Mandatory controls:

```text
Pull Request
CI
Review
QA sign-off for staging/main release
Infrastructure approval for production
GitHub Environment production manual approval
CODEOWNERS
Release tags
Hotfix back-merge
```

Production must never be deployed directly from feature, develop, or staging without passing main and manual approval.

## Related GitHub Management Files

This repository rule document is supported by the following files:

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

Use `docs/25-github-project-management.md` as the detailed guide for:

```text
- labels
- milestones
- GitHub Project fields
- GitHub Project views
- issue templates
- issue lifecycle
- PR tracking
- AI Agent task tracking
- QA/UAT tracking
- release readiness
```

If any `.github` workflow/template file changes, update this document and `docs/25-github-project-management.md`.

## Alignment With Final Planning Documents

This repository rules document must stay aligned with:

```text
docs/25-product-requirement-document.md
docs/26-development-plan.md
docs/27-workflow.md
docs/25-github-project-management.md
```

Update this file when any of the following changes:

```text
- branch workflow
- branch protection
- PR rule
- CI/CD requirement
- issue template
- PR template
- CODEOWNERS
- label taxonomy
- GitHub Project field/view
- QA/UAT workflow
- release/hotfix workflow
```

For detailed day-to-day SOP, use:

```text
docs/27-workflow.md
```

For detailed GitHub Project/Labels/Milestones/Issue/PR management, use:

```text
docs/25-github-project-management.md
```

## Related GitHub Setup Guide

For practical repository setup steps, GitHub Labels, GitHub Milestones, GitHub Project fields/views, branch protection, and GitHub Environments, use:

```text
docs/40-github-repository-setup-labels-project.md
```

This document should stay aligned with:

```text
docs/25-github-project-management.md
docs/27-workflow.md
docs/40-github-repository-setup-labels-project.md
.github/CODEOWNERS
.github/pull_request_template.md
.github/workflows/ci.yml
.github/ISSUE_TEMPLATE/
```

Update this file if branch rules, PR rules, CI/CD rules, GitHub Project setup, or release workflow changes.

## Git Commit Convention

All commits and PR titles should follow the project convention:

```text
type(scope): short description
```

Primary references:

```text
docs/08-coding-standard.md
docs/27-workflow.md
docs/40-github-repository-setup-labels-project.md
```

Examples:

```text
feat(identity): add refresh token rotation
fix(finance): prevent duplicate bill generation
docs(workflow): add git commit convention
chore(ci): add repository validation workflow
test(academic): add attendance scope tests
```

Repository rule:

- PR title should follow the same format as commit convention.
- Squash merge commit must follow the same format.
- Breaking changes must be documented in PR description and commit body if relevant.
- Security and migration commits must include review context.

## Repository Setup Scripts

Repository setup automation is maintained under:

```text
scripts/github/
```

Use these scripts for repeatable setup of:

```text
- private repository
- repository support files
- develop/staging/main branches
- branch protection
- staging/production environments
- labels
- milestones
- GitHub Project
- GitHub Project fields
```

Rules:

- Scripts must not weaken branch protection.
- Scripts must not remove required review/approval rules.
- Scripts must not contain secrets.
- If a script changes repository policy, update this document and `docs/40-github-repository-setup-labels-project.md`.
