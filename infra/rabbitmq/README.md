# RabbitMQ

Local RabbitMQ configuration for `school-platform`.

RabbitMQ is used for asynchronous domain events. The planned event exchange is:

```text
domain.events
```

Exchange and queue declarations are owned by later service or infrastructure tasks. This Sprint 0 Docker Compose task only starts RabbitMQ with the management UI.

## Local Ports

```text
AMQP          : localhost:5672
Management UI: http://localhost:15672
```

## Local Credentials

Defaults are defined in `.env.example`:

```text
RABBITMQ_DEFAULT_USER=school_local
RABBITMQ_DEFAULT_PASS=school_local_password
```

These credentials are for local development only. Do not use them in staging or production.

## Verify Locally

When Docker Compose is running:

```bash
docker compose ps rabbitmq
docker compose logs rabbitmq
```

Open the management UI:

```text
http://localhost:15672
```
