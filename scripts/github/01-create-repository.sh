#!/usr/bin/env bash
set -euo pipefail

: "${GITHUB_OWNER:?Set GITHUB_OWNER, e.g. export GITHUB_OWNER=kuswandi-ti}"
: "${REPO_NAME:=school-platform}"
: "${REPO_DESCRIPTION:=Internal multi-unit school foundation platform for TK, SD, SMP, and SMA operations.}"

if ! command -v gh >/dev/null 2>&1; then
  echo "GitHub CLI 'gh' is required."
  exit 1
fi

gh auth status >/dev/null

if gh repo view "$GITHUB_OWNER/$REPO_NAME" >/dev/null 2>&1; then
  echo "Repository already exists: $GITHUB_OWNER/$REPO_NAME"
  gh repo edit "$GITHUB_OWNER/$REPO_NAME" --visibility private --accept-visibility-change-consequences || true
else
  gh repo create "$GITHUB_OWNER/$REPO_NAME" \
    --private \
    --description "$REPO_DESCRIPTION" \
    --add-readme
fi

gh repo edit "$GITHUB_OWNER/$REPO_NAME" --default-branch main || true

echo "Repository ready: $GITHUB_OWNER/$REPO_NAME"
