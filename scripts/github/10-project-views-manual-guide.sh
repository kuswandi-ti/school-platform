#!/usr/bin/env bash
set -euo pipefail

cat <<'EOF'
GitHub Project Views Manual Setup Guide
======================================

Open:
GitHub → Projects → School Platform MVP

Create these views:

1. MVP Board
   - Layout: Board
   - Group by: Status
   - Columns: Backlog, Ready, In Progress, In Review, QA, Blocked, Done

2. Sprint Board
   - Layout: Board
   - Filter: Sprint = active sprint
   - Group by: Status
   - Visible fields: Priority, Area, Owner, Estimate, Risk, AI Agent

3. Backlog Table
   - Layout: Table
   - Filter: Status != Done
   - Sort: Priority descending

4. By Area
   - Layout: Board or Table
   - Group by: Area

5. By Priority
   - Layout: Table
   - Group by: Priority
   - Sort: Priority descending

6. QA/UAT
   - Layout: Board or Table
   - Filter: Status = QA OR label:"status: needs-qa"
   - Group by: Area

7. AI Agent Tasks
   - Layout: Table
   - Filter: AI Agent = Ready OR label:"ai: ready"

8. Release Readiness
   - Layout: Board or Table
   - Filter: Target Release = MVP OR Sprint = MVP Release
   - Highlight:
     - Critical/High open bugs
     - security review issues
     - QA/UAT sign-off issues
     - backup/restore readiness
     - rollback readiness

Reason:
GitHub Project view automation is less stable across accounts/org permissions.
Manual setup is safer and more transparent.
EOF
