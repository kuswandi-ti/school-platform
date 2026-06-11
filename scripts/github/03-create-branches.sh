#!/usr/bin/env bash
set -euo pipefail

: "${GITHUB_OWNER:?Set GITHUB_OWNER}"
: "${REPO_NAME:=school-platform}"

git fetch origin

git checkout main
git pull origin main

if git ls-remote --exit-code --heads origin develop >/dev/null 2>&1; then
  echo "Remote branch develop already exists."
else
  git checkout -b develop
  git push -u origin develop
fi

git checkout main
if git ls-remote --exit-code --heads origin staging >/dev/null 2>&1; then
  echo "Remote branch staging already exists."
else
  git checkout -b staging
  git push -u origin staging
fi

git checkout develop || git checkout main

gh repo edit "$GITHUB_OWNER/$REPO_NAME" --default-branch main || true

echo "Branches ready: main, develop, staging"
