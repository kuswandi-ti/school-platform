#!/usr/bin/env bash
set -euo pipefail

: "${GITHUB_OWNER:?Set GITHUB_OWNER}"
: "${REPO_NAME:=school-platform}"

REPO="$GITHUB_OWNER/$REPO_NAME"

create_label() {
  local name="$1"
  local color="$2"
  local description="$3"

  if gh label list --repo "$REPO" --search "$name" --json name --jq '.[].name' | grep -Fx "$name" >/dev/null 2>&1; then
    gh label edit "$name" --repo "$REPO" --color "$color" --description "$description"
  else
    gh label create "$name" --repo "$REPO" --color "$color" --description "$description"
  fi
}

# type
create_label "type: feature" "1f883d" "New feature or implementation task"
create_label "type: bug" "d1242f" "Bug or regression"
create_label "type: chore" "6e7781" "Maintenance or housekeeping"
create_label "type: docs" "0969da" "Documentation"
create_label "type: refactor" "a2eeef" "Refactor without behavior change"
create_label "type: test" "0e8a16" "Test changes"
create_label "type: security" "8250df" "Security or privacy"
create_label "type: infra" "6f42c1" "Infrastructure or CI/CD"
create_label "type: spike" "d4c5f9" "Research or exploration"
create_label "type: hotfix" "b60205" "Urgent production fix"

# area
for area in api-gateway identity school-core admission academic finance communication reporting web-admin mobile infra docs security observability ci-cd file-management shared-go proto openapi events; do
  create_label "area: $area" "ededed" "Area: $area"
done

# sprint
for i in 0 1 2 3 4 5 6 7 8 9 10; do
  create_label "sprint: $i" "ededed" "Sprint $i"
done

# priority
create_label "priority: critical" "b60205" "Release blocker or critical issue"
create_label "priority: high" "d93f0b" "High priority"
create_label "priority: medium" "fbca04" "Medium priority"
create_label "priority: low" "0e8a16" "Low priority"

# status
for status in ready in-progress blocked needs-review needs-qa qa-passed done; do
  create_label "status: $status" "cfd3d7" "Status: $status"
done

# ai
create_label "ai: ready" "5319e7" "Ready for AI Agent"
create_label "ai: needs-context" "c5def5" "Needs more context before AI Agent"
create_label "ai: generated" "bfd4f2" "Generated or assisted by AI Agent"
create_label "ai: needs-human-review" "f9d0c4" "Requires human review"
create_label "ai: do-not-use-agent" "000000" "Do not use AI Agent"

# risk
create_label "risk: low" "0e8a16" "Low risk"
create_label "risk: medium" "fbca04" "Medium risk"
create_label "risk: high" "d93f0b" "High risk"
create_label "risk: breaking-change" "b60205" "Breaking change"
create_label "risk: migration" "5319e7" "Migration or data risk"
create_label "risk: data-sensitive" "8250df" "Touches Restricted or Confidential data"

# review
for review in backend frontend mobile infra qa security product; do
  create_label "review: $review" "c2e0c6" "Needs $review review"
done

echo "Labels configured."
