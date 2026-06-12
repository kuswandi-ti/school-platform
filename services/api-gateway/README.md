# API Gateway

Custom Go API Gateway for `school-platform`.

The gateway is the single external REST/JSON entrypoint for Next.js and Flutter. It owns edge concerns such as request routing, response standardization, request ID and correlation ID propagation, logging, CORS, and future REST-to-gRPC mapping.

Do not put domain business logic or service-owned database queries in the API Gateway.

## Current Endpoints

```text
GET /healthz
GET /readyz
GET /api/v1/ping
```

`/api/v1/ping` returns the standard response envelope:

```json
{
  "data": {
    "message": "pong"
  },
  "meta": null,
  "error": null
}
```

## Local Run

```bash
cp .env.example .env
go run ./cmd/server
```

Default URL:

```text
http://localhost:8080
```

The local example uses `HTTP_PORT=8080`. Set `HTTP_ADDR` only when you need a full bind address such as `127.0.0.1:8080`.

Check:

```bash
curl http://localhost:8080/healthz
curl http://localhost:8080/readyz
curl http://localhost:8080/api/v1/ping
```

## Tests

```bash
go test ./...
```

## Docker

```bash
docker build -t school-platform-api-gateway .
docker run --rm -p 8080:8080 --env-file .env school-platform-api-gateway
```

## Future Work

Sprint 1 will add authentication and Identity Service integration. Later service clients should be placed behind explicit REST-to-gRPC adapters and must not introduce domain business rules into the gateway.
