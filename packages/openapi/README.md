# OpenAPI Contracts

External REST API contracts live here.

The API Gateway exposes REST/JSON APIs under:

```text
/api/v1
```

Sprint 0 includes only a minimal skeleton and health/ping examples. Do not define full business endpoints until the related sprint task is active.

## Update Rules

- Update `api-gateway.v1.yaml` when an external REST endpoint changes.
- Keep user-facing messages aligned with `docs/copywriting` when applicable.
- Use English for paths, operation IDs, schema names, and technical fields.
- Keep the standard response envelope: `data`, `meta`, `error`.
- Do not expose tokens, passwords, or Confidential details in schemas.
- Coordinate OpenAPI changes with frontend, mobile, backend, and QA.
