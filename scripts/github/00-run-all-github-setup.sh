#!/usr/bin/env bash
set -euo pipefail

: "${GITHUB_OWNER:?Set GITHUB_OWNER, e.g. export GITHUB_OWNER=kuswandi-ti}"
: "${REPO_NAME:=school-platform}"

echo "==> Running GitHub setup for $GITHUB_OWNER/$REPO_NAME"

bash "$(dirname "$0")/01-create-repository.sh"
bash "$(dirname "$0")/02-bootstrap-repository-files.sh"
bash "$(dirname "$0")/03-create-branches.sh"
bash "$(dirname "$0")/04-setup-branch-protection.sh"
bash "$(dirname "$0")/05-setup-environments.sh"
bash "$(dirname "$0")/06-setup-labels.sh"
bash "$(dirname "$0")/07-setup-milestones.sh"
bash "$(dirname "$0")/08-setup-project.sh"

cat <<'EOF'

Next:
1. Run: gh project list --owner "$GITHUB_OWNER"
2. Get the project number for "School Platform MVP"
3. Run:
   PROJECT_NUMBER="<project-number>" bash scripts/github/09-setup-project-fields.sh
4. Run:
   bash scripts/github/10-project-views-manual-guide.sh

EOF
