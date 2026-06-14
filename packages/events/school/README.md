# School Events

School Core event schemas.

## `school.school.created` v1

Published through the School Core transactional outbox after a school is created.
The payload contains only the school identifier, code, name, level, and status.

Schema: `school.school.created.v1.schema.json`
