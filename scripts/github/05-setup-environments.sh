#!/usr/bin/env bash
set -euo pipefail

: "${GITHUB_OWNER:?Set GITHUB_OWNER}"
: "${REPO_NAME:=school-platform}"

create_environment() {
  local env_name="$1"
  echo "Creating/updating environment: $env_name"
  gh api \
    --method PUT \
    -H "Accept: application/vnd.github+json" \
    "repos/$GITHUB_OWNER/$REPO_NAME/environments/$env_name" \
    --input - <<EOF
{}
EOF
}

create_environment "staging"
create_environment "production"

if [ -n "${PRODUCTION_REVIEWER_USER:-}" ]; then
  echo "Resolving reviewer user id for $PRODUCTION_REVIEWER_USER"
  reviewer_id="$(gh api "users/$PRODUCTION_REVIEWER_USER" --jq '.id')"

  echo "Setting production required reviewer"
  gh api \
    --method PUT \
    -H "Accept: application/vnd.github+json" \
    "repos/$GITHUB_OWNER/$REPO_NAME/environments/production" \
    --input - <<EOF
{
  "wait_timer": 0,
  "reviewers": [
    {
      "type": "User",
      "id": $reviewer_id
    }
  ],
  "deployment_branch_policy": null
}
EOF
else
  cat <<'EOF'
PRODUCTION_REVIEWER_USER is not set.

To configure required reviewer:
export PRODUCTION_REVIEWER_USER="github-username"
bash scripts/github/05-setup-environments.sh

Or configure manually:
Repository → Settings → Environments → production → Required reviewers
EOF
fi

echo "Environments configured."
