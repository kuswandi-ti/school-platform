#!/usr/bin/env bash
set -euo pipefail

: "${GITHUB_OWNER:?Set GITHUB_OWNER}"
: "${REPO_NAME:=school-platform}"

REPO="$GITHUB_OWNER/$REPO_NAME"

create_milestone() {
  local title="$1"
  local description="$2"

  if gh api "/repos/$REPO/milestones?state=all" --jq '.[].title' | grep -Fx "$title" >/dev/null 2>&1; then
    echo "Milestone already exists: $title"
  else
    gh api \
      --method POST \
      -H "Accept: application/vnd.github+json" \
      "/repos/$REPO/milestones" \
      -f title="$title" \
      -f description="$description"
  fi
}

create_milestone "Sprint 0 — Project Foundation" "Repository, local environment, service skeleton, CI, and docs foundation"
create_milestone "Sprint 1 — Identity & Access" "Authentication, authorization, actor context, and API Gateway auth"
create_milestone "Sprint 2 — School Core" "Foundation, school, student, guardian, teacher, class, and assignment master data"
create_milestone "Sprint 3 — File Management + Import Excel" "Private file management and Excel import"
create_milestone "Sprint 4 — PPDB" "Admission period, applicant workflow, document verification, and conversion"
create_milestone "Sprint 5 — Finance / SPP" "Manual payment, billing, fee policy, verification, and receipt"
create_milestone "Sprint 6 — Academic Basic" "Curriculum, subject, schedule, attendance"
create_milestone "Sprint 7 — Report Card / E-Rapor Basic" "Score, grade book, report card publish and lock"
create_milestone "Sprint 8 — Communication / Notification" "Announcement and notification system"
create_milestone "Sprint 9 — Reporting Dashboard" "Reporting service and dashboard projections"
create_milestone "Sprint 10 — Security, Observability, Backup, UAT Hardening" "Hardening before pilot/production"
create_milestone "MVP Release" "MVP release readiness and production release tracking"

echo "Milestones configured."
