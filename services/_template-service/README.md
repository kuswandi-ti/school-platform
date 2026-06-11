# Template Service

Reusable Go service template for future `school-platform` microservices.

This template provides:

- environment-based config loading
- structured JSON logging with `slog`
- HTTP server startup
- `GET /healthz`
- `GET /readyz`
- request ID and correlation ID middleware
- graceful shutdown on `SIGINT` and `SIGTERM`
- Dockerfile for service image builds

It intentionally does not include business logic, database repositories, authentication, gRPC, or domain modules.

## Local Run

Copy local environment values:

```bash
cp .env.example .env
```

Run the service:

```bash
go run ./cmd/server
```

Default endpoint:

```text
http://localhost:8081
```

Health checks:

```bash
curl http://localhost:8081/healthz
curl http://localhost:8081/readyz
```

## Tests

```bash
go test ./...
```

Tests are included for the health handlers. Run them before copying this template into a concrete service.

## Docker

Build:

```bash
docker build -t school-platform-template-service .
```

Run:

```bash
docker run --rm -p 8081:8081 --env-file .env school-platform-template-service
```

## Copying This Template

When creating a new service:

1. Copy `services/_template-service` to `services/<service-name>`.
2. Update `go.mod` module name.
3. Update `SERVICE_NAME` in `.env.example`.
4. Keep business logic inside `internal/domain` and `internal/usecase` when those folders are introduced.
5. Keep transport concerns out of business logic.
