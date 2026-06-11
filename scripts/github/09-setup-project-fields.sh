#!/usr/bin/env bash
set -euo pipefail

: "${GITHUB_OWNER:?Set GITHUB_OWNER}"
: "${PROJECT_NUMBER:?Set PROJECT_NUMBER from gh project list --owner $GITHUB_OWNER}"

create_single_select_field() {
  local name="$1"
  shift
  echo "Creating field: $name"
  # gh project field-create supports single-select fields in recent GitHub CLI versions.
  # If this fails, create field manually in GitHub Project UI.
  gh project field-create "$PROJECT_NUMBER" \
    --owner "$GITHUB_OWNER" \
    --name "$name" \
    --data-type SINGLE_SELECT \
    "$@" || {
      echo "Could not create field '$name' via CLI. Create it manually in GitHub Project UI."
    }
}

create_text_field() {
  local name="$1"
  echo "Creating text field: $name"
  gh project field-create "$PROJECT_NUMBER" \
    --owner "$GITHUB_OWNER" \
    --name "$name" \
    --data-type TEXT || {
      echo "Could not create text field '$name' via CLI. Create it manually in GitHub Project UI."
    }
}

create_number_field() {
  local name="$1"
  echo "Creating number field: $name"
  gh project field-create "$PROJECT_NUMBER" \
    --owner "$GITHUB_OWNER" \
    --name "$name" \
    --data-type NUMBER || {
      echo "Could not create number field '$name' via CLI. Create it manually in GitHub Project UI."
    }
}

# Some gh versions require --single-select-options "A,B,C".
# If your gh version does not support this syntax, create options manually in UI.
create_single_select_field "Status" --single-select-options "Backlog,Ready,In Progress,In Review,QA,Blocked,Done"
create_single_select_field "Sprint" --single-select-options "Sprint 0,Sprint 1,Sprint 2,Sprint 3,Sprint 4,Sprint 5,Sprint 6,Sprint 7,Sprint 8,Sprint 9,Sprint 10,MVP Release"
create_single_select_field "Priority" --single-select-options "Critical,High,Medium,Low"
create_single_select_field "Area" --single-select-options "api-gateway,identity,school-core,admission,academic,finance,communication,reporting,web-admin,mobile,infra,docs,security,observability,ci-cd,file-management"
create_single_select_field "Type" --single-select-options "feature,bug,chore,docs,refactor,test,security,infra,spike,hotfix"
create_number_field "Estimate"
create_single_select_field "Risk" --single-select-options "Low,Medium,High,Breaking Change,Migration,Data Sensitive"
create_single_select_field "Platform" --single-select-options "Backend,Web,Mobile,Infra,Docs,QA,Product"
create_single_select_field "AI Agent" --single-select-options "Ready,Needs Context,Generated,Needs Human Review,Do Not Use"
create_text_field "Target Release"

cat <<'EOF'

Owner field:
GitHub Project already has assignee/people-related fields depending on configuration.
If needed, create Owner field manually in the GitHub Project UI.

Next:
bash scripts/github/10-project-views-manual-guide.sh
EOF
