-- name: CreateUser :one
INSERT INTO users (
    id,
    email,
    phone,
    password_hash,
    display_name,
    avatar_file_id,
    status
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, email, phone, password_hash, display_name, avatar_file_id, status,
    last_login_at, created_at, updated_at;

-- name: GetUserByEmail :one
SELECT id, email, phone, password_hash, display_name, avatar_file_id, status,
    last_login_at, created_at, updated_at
FROM users
WHERE LOWER(email) = LOWER($1)
LIMIT 1;

-- name: GetUserByID :one
SELECT id, email, phone, password_hash, display_name, avatar_file_id, status,
    last_login_at, created_at, updated_at
FROM users
WHERE id = $1;

-- name: UpdateUserLastLogin :one
UPDATE users
SET last_login_at = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, email, phone, password_hash, display_name, avatar_file_id, status,
    last_login_at, created_at, updated_at;

-- name: UpdateUserStatus :one
UPDATE users
SET status = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, email, phone, password_hash, display_name, avatar_file_id, status,
    last_login_at, created_at, updated_at;

-- name: CreateUserSession :one
INSERT INTO user_sessions (
    id,
    user_id,
    refresh_token_hash,
    device_id,
    ip_address,
    user_agent,
    expires_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING id;

-- name: GetUserSessionByRefreshHash :one
SELECT id, user_id, refresh_token_hash, device_id, ip_address, user_agent,
    expires_at, revoked_at, last_used_at, created_at, updated_at
FROM user_sessions
WHERE refresh_token_hash = $1;

-- name: RevokeUserSessionForRotation :one
UPDATE user_sessions
SET revoked_at = $2,
    last_used_at = $2,
    updated_at = $2
WHERE id = $1
  AND revoked_at IS NULL
  AND expires_at > $2
RETURNING user_id, device_id;
