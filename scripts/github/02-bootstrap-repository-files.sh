#!/usr/bin/env bash
set -euo pipefail

: "${GITHUB_OWNER:?Set GITHUB_OWNER}"
: "${REPO_NAME:=school-platform}"
: "${ADD_INTERNAL_LICENSE:=false}"

ROOT="$(pwd)"

mkdir -p .github/ISSUE_TEMPLATE .github/workflows docs scripts

if [ ! -f README.md ]; then
cat > README.md <<'EOF'
# School Platform

Internal multi-unit school foundation platform for TK, SD, SMP, and SMA operations.

## Documentation

See `docs/README.md`.
EOF
fi

cat > .gitignore <<'EOF'
# Environment
.env
.env.*
!.env.example

# Secrets
*.pem
*.key
*.p12
*.pfx
id_rsa
id_ed25519

# OS / Editor
.DS_Store
Thumbs.db
.vscode/
.idea/

# Go
bin/
coverage.out

# Node / Next.js
node_modules/
.next/
out/
dist/

# Flutter
.dart_tool/
build/
.flutter-plugins
.flutter-plugins-dependencies

# Logs
*.log

# Local data
tmp/
storage/
EOF

if [ "$ADD_INTERNAL_LICENSE" = "true" ] && [ ! -f LICENSE ]; then
cat > LICENSE <<'EOF'
Copyright (c) 2026

This repository is private and intended for internal use only.
No license is granted for public use, distribution, or modification unless explicitly stated by the repository owner.
EOF
fi

cat > .github/CODEOWNERS <<EOF
* @$GITHUB_OWNER

/docs/ @$GITHUB_OWNER
/.github/ @$GITHUB_OWNER
/services/api-gateway/ @$GITHUB_OWNER
/services/identity-service/ @$GITHUB_OWNER
/services/school-core-service/ @$GITHUB_OWNER
/services/admission-service/ @$GITHUB_OWNER
/services/academic-service/ @$GITHUB_OWNER
/services/finance-service/ @$GITHUB_OWNER
/services/communication-service/ @$GITHUB_OWNER
/services/reporting-service/ @$GITHUB_OWNER
/apps/web-admin/ @$GITHUB_OWNER
/apps/mobile-app/ @$GITHUB_OWNER
EOF

cat > .github/pull_request_template.md <<'EOF'
## Summary

## Related Issue

Closes #

## Type

- [ ] feature
- [ ] bug
- [ ] docs
- [ ] chore
- [ ] refactor
- [ ] test
- [ ] security
- [ ] infra

## Checklist

- [ ] PR title follows `type(scope): short description`
- [ ] CI pass
- [ ] Tests added/updated
- [ ] Docs updated if needed
- [ ] No secrets committed
- [ ] Permission/scope checked if relevant
- [ ] Object-level authorization checked if relevant
- [ ] Audit/event/file/privacy checked if relevant
EOF

cat > .github/ISSUE_TEMPLATE/feature_task.yml <<'EOF'
name: Feature / Task
description: Create a feature or implementation task
title: "Sprint N Task X.Y — "
labels: ["type: feature", "status: ready"]
body:
  - type: textarea
    id: objective
    attributes:
      label: Objective
    validations:
      required: true
  - type: textarea
    id: scope
    attributes:
      label: Scope
    validations:
      required: true
  - type: textarea
    id: acceptance
    attributes:
      label: Acceptance Criteria
    validations:
      required: true
EOF

cat > .github/ISSUE_TEMPLATE/bug_report.yml <<'EOF'
name: Bug Report
description: Report a bug or regression
title: "fix(scope): "
labels: ["type: bug"]
body:
  - type: textarea
    id: bug
    attributes:
      label: Bug Description
    validations:
      required: true
  - type: textarea
    id: steps
    attributes:
      label: Steps to Reproduce
    validations:
      required: true
  - type: textarea
    id: expected
    attributes:
      label: Expected Result
    validations:
      required: true
EOF

cat > .github/ISSUE_TEMPLATE/ai_agent_task.yml <<'EOF'
name: AI Agent Task
description: Task prepared for AI Agent assistance
title: "Sprint N Task X.Y — "
labels: ["ai: ready", "ai: needs-human-review"]
body:
  - type: textarea
    id: context
    attributes:
      label: Required Context
    validations:
      required: true
  - type: textarea
    id: instruction
    attributes:
      label: AI Agent Instruction
    validations:
      required: true
  - type: textarea
    id: guardrails
    attributes:
      label: Guardrails
    validations:
      required: true
EOF

cat > .github/ISSUE_TEMPLATE/security_review.yml <<'EOF'
name: Security Review
description: Security or privacy review task
title: "security(scope): "
labels: ["type: security", "review: security"]
body:
  - type: textarea
    id: area
    attributes:
      label: Area
    validations:
      required: true
  - type: textarea
    id: risks
    attributes:
      label: Risks
    validations:
      required: true
  - type: textarea
    id: checklist
    attributes:
      label: Review Checklist
    validations:
      required: true
EOF

cat > .github/ISSUE_TEMPLATE/qa_uat.yml <<'EOF'
name: QA / UAT
description: QA or UAT validation task
title: "test(scope): "
labels: ["type: test", "review: qa"]
body:
  - type: textarea
    id: scenario
    attributes:
      label: Scenario
    validations:
      required: true
  - type: textarea
    id: steps
    attributes:
      label: Test Steps
    validations:
      required: true
  - type: textarea
    id: expected
    attributes:
      label: Expected Result
    validations:
      required: true
EOF

cat > .github/workflows/ci.yml <<'EOF'
name: CI

on:
  pull_request:
    branches: [develop, staging, main]
  push:
    branches: [develop, staging, main]

jobs:
  repository-check:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Check required files
        run: |
          test -f README.md
          test -f .gitignore
          test -f .github/pull_request_template.md
          test -f .github/CODEOWNERS
      - name: Block committed secrets
        run: |
          if find . -name ".env" -o -name "*.pem" -o -name "*.key" | grep -v ".git"; then
            echo "Secret-like files found. Remove them."
            exit 1
          fi

  yaml-check:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Install PyYAML
        run: python -m pip install pyyaml
      - name: Validate YAML
        run: |
          python - <<'PY'
          import pathlib, yaml
          for p in list(pathlib.Path(".github").rglob("*.yml")) + list(pathlib.Path(".github").rglob("*.yaml")):
              with p.open("r", encoding="utf-8") as f:
                  yaml.safe_load(f)
              print("valid", p)
          PY
EOF

if git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
  git add README.md .gitignore .github
  if [ "$ADD_INTERNAL_LICENSE" = "true" ]; then
    git add LICENSE
  fi
  if ! git diff --cached --quiet; then
    git commit -m "chore(repository): add github support files"
  else
    echo "No repository support file changes to commit."
  fi
fi

echo "Repository support files bootstrapped."
