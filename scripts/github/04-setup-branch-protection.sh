#!/usr/bin/env bash
set -euo pipefail

: "${GITHUB_OWNER:?Set GITHUB_OWNER}"
: "${REPO_NAME:=school-platform}"

setup_branch_protection() {
  local branch="$1"
  local approvals="$2"

  echo "Setting branch protection for $branch"

  gh api \
    --method PUT \
    -H "Accept: application/vnd.github+json" \
    "/repos/$GITHUB_OWNER/$REPO_NAME/branches/$branch/protection" \
    --input - <<EOF
{
  "required_status_checks": null,
  "enforce_admins": true,
  "required_pull_request_reviews": {
    "required_approving_review_count": $approvals,
    "dismiss_stale_reviews": true,
    "require_code_owner_reviews": false,
    "require_last_push_approval": false
  },
  "restrictions": null,
  "required_conversation_resolution": true,
  "allow_force_pushes": false,
  "allow_deletions": false,
  "block_creations": false,
  "required_linear_history": false,
  "lock_branch": false,
  "allow_fork_syncing": true
}
EOF
}

setup_branch_protection "develop" 1
setup_branch_protection "staging" 1
setup_branch_protection "main" 1

echo "Branch protection configured."
