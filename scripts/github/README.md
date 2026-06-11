# GitHub Setup Scripts

These scripts support `docs/40-github-repository-setup-labels-project.md`.

## Prerequisites

```bash
gh auth login
gh auth status
git --version
```

Set variables:

```bash
export GITHUB_OWNER="<org-or-user>"
export REPO_NAME="school-platform"
export PRODUCTION_REVIEWER_USER="kuswandi-ti"
```

## Recommended Order

```bash
bash scripts/github/01-create-repository.sh
# clone repository if not already inside it
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

## One-command Setup

```bash
bash scripts/github/00-run-all-github-setup.sh
```

Project fields/views may still need manual adjustment in GitHub UI depending on account permissions and GitHub CLI version.
