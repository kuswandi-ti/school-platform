#!/usr/bin/env bash
set -euo pipefail

: "${GITHUB_OWNER:?Set GITHUB_OWNER}"
: "${PROJECT_TITLE:=School Platform MVP}"

if gh project list --owner "$GITHUB_OWNER" --format json --jq '.[].title' | grep -Fx "$PROJECT_TITLE" >/dev/null 2>&1; then
  echo "Project already exists: $PROJECT_TITLE"
else
  gh project create --owner "$GITHUB_OWNER" --title "$PROJECT_TITLE"
fi

echo "Projects:"
gh project list --owner "$GITHUB_OWNER"

cat <<'EOF'

Next:
1. Find project number for "School Platform MVP" above.
2. Run:
   PROJECT_NUMBER="<number>" bash scripts/github/09-setup-project-fields.sh
EOF
