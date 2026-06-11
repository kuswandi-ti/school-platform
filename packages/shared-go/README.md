# Shared Go

Shared Go utilities for `school-platform` services.

This module contains generic cross-service helpers and conventions only. It must not contain service-specific business logic, service database access, or domain ownership decisions.

## Packages

```text
logger     slog setup helper
config     small environment config helpers
response   standard REST response envelope
errors     standard application error codes and type
context    actor/request context placeholder types
audit      audit record placeholder types
events     domain event envelope placeholder types
messaging  message publisher/consumer interfaces
numbering  document numbering placeholder types
files      file metadata placeholder types
```

## Rules

- Keep packages generic.
- Avoid circular dependencies.
- Do not depend on service-owned databases.
- Do not add business workflows here.
- Do not log tokens, passwords, or Confidential data.
- Extend only when more than one service needs the same convention.
