# MinIO

Local MinIO configuration for `school-platform`.

MinIO provides S3-compatible object storage for local development. Files are private by default, and access to private files must go through backend authorization and signed URLs in later service tasks.

## Local Ports

```text
API    : http://localhost:9000
Console: http://localhost:9001
```

## Local Credentials

Defaults are defined in `.env.example`:

```text
MINIO_ROOT_USER=school_local_minio
MINIO_ROOT_PASSWORD=school_local_minio_password
```

These credentials are for local development only. Do not use them in staging or production.

## Verify Locally

When Docker Compose is running:

```bash
docker compose ps minio
docker compose logs minio
```

Open the console:

```text
http://localhost:9001
```

Use the local credentials from `.env.example`:

```text
Username: school_local_minio
Password: school_local_minio_password
```

Bucket creation is intentionally not included in this task. Add buckets in a later storage setup task or through service-specific setup scripts.
