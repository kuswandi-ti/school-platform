-- name: GetRoleByCode :one
SELECT id, code, name, description, is_system, created_at, updated_at
FROM roles
WHERE code = $1;

-- name: CreateUserRoleAssignment :one
INSERT INTO user_role_assignments (
    id,
    user_id,
    role_id,
    foundation_id,
    school_id,
    class_id,
    student_id,
    employee_id,
    subject_id,
    scope_json,
    starts_at,
    ends_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
)
RETURNING id, user_id, role_id, foundation_id, school_id, class_id, student_id,
    employee_id, subject_id, scope_json, is_active, starts_at, ends_at, created_at, updated_at;

-- name: CreateIdentityAuditLog :exec
INSERT INTO identity_audit_logs (
    id,
    foundation_id,
    school_id,
    actor_user_id,
    actor_role,
    action,
    module,
    entity_type,
    entity_id,
    old_values_json,
    new_values_json,
    metadata_json,
    ip_address,
    user_agent,
    request_id,
    correlation_id,
    occurred_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17
);

-- name: GetUserContextAssignments :many
SELECT
    ura.id AS assignment_id,
    ura.user_id,
    r.code AS role_code,
    ura.foundation_id,
    ura.school_id,
    ura.class_id,
    ura.student_id,
    ura.employee_id,
    ura.subject_id,
    ura.scope_json,
    ura.starts_at,
    ura.ends_at,
    COALESCE(
        ARRAY_AGG(DISTINCT p.code ORDER BY p.code) FILTER (WHERE p.code IS NOT NULL),
        ARRAY[]::VARCHAR[]
    )::TEXT[] AS permission_codes
FROM user_role_assignments ura
JOIN roles r ON r.id = ura.role_id
LEFT JOIN role_permissions rp ON rp.role_id = r.id
LEFT JOIN permissions p ON p.id = rp.permission_id
WHERE ura.user_id = $1
  AND ura.is_active = TRUE
  AND (ura.starts_at IS NULL OR ura.starts_at <= $2)
  AND (ura.ends_at IS NULL OR ura.ends_at > $2)
GROUP BY ura.id, r.code
ORDER BY r.code, ura.foundation_id, ura.school_id, ura.class_id, ura.student_id, ura.subject_id;
