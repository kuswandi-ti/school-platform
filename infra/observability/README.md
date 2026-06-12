# Observability

This folder contains the local observability skeleton for `school-platform`.

Sprint 0 keeps observability intentionally small:

- structured JSON logs with `slog`
- `request_id` and `correlation_id` in request logs
- reserved `/metrics` endpoint on the API Gateway and Go service template
- optional Docker Compose profile for Prometheus, Grafana, and Loki

## Current Baseline

The current local observability baseline is log-first:

- API Gateway logs in structured JSON
- Go service template logs in structured JSON
- startup logs include `service` and `environment`
- access logs include `method`, `path`, `status`, `duration_ms`, `request_id`, and `correlation_id`

This baseline is enough for Sprint 0 local troubleshooting without introducing a full monitoring stack by default.

## Optional Local Stack

The root `docker-compose.yml` includes an optional `observability` profile with:

- `prometheus`
- `grafana`
- `loki`

These services are disabled unless you start the profile explicitly.

Example:

```bash
docker compose --profile observability up -d
```

Stop it with:

```bash
docker compose --profile observability down
```

Recommended local URLs when the profile is enabled:

```text
Prometheus : http://localhost:9090
Grafana    : http://localhost:3001
Loki       : http://localhost:3100
```

## Metrics Path

`/metrics` is a reserved placeholder path for future Prometheus-compatible metrics output.

Sprint 0 does not implement:

- Prometheus exposition libraries
- custom counters and histograms
- distributed tracing
- alert rules
- production dashboards

The placeholder exists so future services can adopt a consistent path without changing local conventions later.

## Future Direction

Later sprints can extend this folder with:

- Prometheus scrape targets for running services
- Grafana dashboards
- Loki log shipping configuration
- OpenTelemetry collector config
- alert rules for staging/production

Those additions should stay local-safe by default and must not introduce production secrets.
