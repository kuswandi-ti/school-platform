# Redis

Local Redis configuration for `school-platform`.

Redis is available for local cache, rate limiting, and temporary infrastructure needs introduced by later service tasks.

## Local Port

```text
localhost:6379
```

## Verify Locally

When Docker Compose is running:

```bash
docker compose ps redis
docker compose logs redis
```

Ping Redis:

```bash
docker compose exec redis redis-cli ping
```

Redis local data is stored in the named Docker volume `redis_data`.
