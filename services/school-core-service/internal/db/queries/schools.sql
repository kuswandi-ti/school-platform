-- name: GetFoundationByID :one
SELECT * FROM foundations WHERE id = $1;

-- name: ListSchoolsByFoundation :many
SELECT * FROM schools WHERE foundation_id = $1 ORDER BY school_code, id;

-- name: GetSchoolByScope :one
SELECT * FROM schools WHERE id = $1 AND foundation_id = $2;

-- name: CreateSchool :one
INSERT INTO schools (
    id, foundation_id, school_code, name, school_level, npsn, address,
    phone, email, logo_file_id, status
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: UpdateSchool :one
UPDATE schools SET
    school_code = $3,
    name = $4,
    school_level = $5,
    npsn = $6,
    address = $7,
    phone = $8,
    email = $9,
    logo_file_id = $10,
    status = $11,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND foundation_id = $2
RETURNING *;

-- name: CreateAuditLog :exec
INSERT INTO audit_logs (
    id, foundation_id, school_id, actor_user_id, actor_role, action, module,
    entity_type, entity_id, old_values_json, new_values_json, metadata_json,
    ip_address, user_agent, request_id, correlation_id, occurred_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17);

-- name: CreateOutboxEvent :exec
INSERT INTO outbox_events (
    id, event_id, event_type, event_version, aggregate_type, aggregate_id, payload_json
) VALUES ($1, $2, $3, $4, $5, $6, $7);
