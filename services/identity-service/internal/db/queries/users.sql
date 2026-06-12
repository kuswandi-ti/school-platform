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
