-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(50),
    password_hash TEXT NOT NULL,
    display_name VARCHAR(150) NOT NULL,
    avatar_file_id UUID,
    status VARCHAR(50) NOT NULL,
    last_login_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT users_status_check CHECK (status IN ('active', 'inactive', 'locked', 'invited'))
);

CREATE UNIQUE INDEX users_email_unique_idx ON users (LOWER(email));
CREATE INDEX users_phone_idx ON users (phone) WHERE phone IS NOT NULL;
CREATE INDEX users_status_idx ON users (status);

CREATE TABLE roles (
    id UUID PRIMARY KEY,
    code VARCHAR(100) NOT NULL,
    name VARCHAR(150) NOT NULL,
    description TEXT,
    is_system BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT roles_code_unique UNIQUE (code)
);

CREATE TABLE permissions (
    id UUID PRIMARY KEY,
    code VARCHAR(150) NOT NULL,
    name VARCHAR(150) NOT NULL,
    description TEXT,
    module VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT permissions_code_unique UNIQUE (code)
);

CREATE INDEX permissions_module_idx ON permissions (module);

CREATE TABLE role_permissions (
    id UUID PRIMARY KEY,
    role_id UUID NOT NULL REFERENCES roles (id) ON DELETE CASCADE,
    permission_id UUID NOT NULL REFERENCES permissions (id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT role_permissions_role_permission_unique UNIQUE (role_id, permission_id)
);

CREATE INDEX role_permissions_role_id_idx ON role_permissions (role_id);
CREATE INDEX role_permissions_permission_id_idx ON role_permissions (permission_id);

CREATE TABLE user_role_assignments (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    role_id UUID NOT NULL REFERENCES roles (id) ON DELETE RESTRICT,
    foundation_id UUID NOT NULL,
    school_id UUID,
    class_id UUID,
    student_id UUID,
    employee_id UUID,
    subject_id UUID,
    scope_json JSONB,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    starts_at TIMESTAMPTZ,
    ends_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT user_role_assignments_time_range_check CHECK (
        ends_at IS NULL OR starts_at IS NULL OR ends_at > starts_at
    )
);

CREATE UNIQUE INDEX user_role_assignments_active_scope_unique_idx
    ON user_role_assignments (
        user_id,
        role_id,
        foundation_id,
        school_id,
        class_id,
        student_id,
        employee_id,
        subject_id
    ) NULLS NOT DISTINCT
    WHERE is_active;
CREATE INDEX user_role_assignments_user_id_idx ON user_role_assignments (user_id);
CREATE INDEX user_role_assignments_role_id_idx ON user_role_assignments (role_id);
CREATE INDEX user_role_assignments_foundation_id_idx ON user_role_assignments (foundation_id);
CREATE INDEX user_role_assignments_school_id_idx ON user_role_assignments (school_id) WHERE school_id IS NOT NULL;
CREATE INDEX user_role_assignments_active_idx ON user_role_assignments (is_active);

CREATE TABLE user_devices (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    device_uid VARCHAR(255) NOT NULL,
    platform VARCHAR(50) NOT NULL,
    device_name VARCHAR(150),
    fcm_token TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    last_seen_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT user_devices_platform_check CHECK (platform IN ('web', 'android', 'ios')),
    CONSTRAINT user_devices_user_device_unique UNIQUE (user_id, device_uid)
);

CREATE INDEX user_devices_user_id_idx ON user_devices (user_id);
CREATE INDEX user_devices_active_idx ON user_devices (is_active);

CREATE TABLE user_sessions (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    refresh_token_hash TEXT NOT NULL,
    device_id UUID REFERENCES user_devices (id) ON DELETE SET NULL,
    ip_address VARCHAR(100),
    user_agent TEXT,
    expires_at TIMESTAMPTZ NOT NULL,
    revoked_at TIMESTAMPTZ,
    last_used_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT user_sessions_refresh_token_hash_unique UNIQUE (refresh_token_hash),
    CONSTRAINT user_sessions_expiry_check CHECK (expires_at > created_at),
    CONSTRAINT user_sessions_revocation_check CHECK (revoked_at IS NULL OR revoked_at >= created_at)
);

CREATE INDEX user_sessions_user_id_idx ON user_sessions (user_id);
CREATE INDEX user_sessions_device_id_idx ON user_sessions (device_id) WHERE device_id IS NOT NULL;
CREATE INDEX user_sessions_expires_at_idx ON user_sessions (expires_at);
CREATE INDEX user_sessions_active_idx ON user_sessions (user_id, expires_at) WHERE revoked_at IS NULL;

CREATE TABLE identity_audit_logs (
    id UUID PRIMARY KEY,
    foundation_id UUID NOT NULL,
    school_id UUID,
    actor_user_id UUID REFERENCES users (id) ON DELETE SET NULL,
    actor_role VARCHAR(100),
    action VARCHAR(150) NOT NULL,
    module VARCHAR(100) NOT NULL,
    entity_type VARCHAR(100) NOT NULL,
    entity_id UUID NOT NULL,
    old_values_json JSONB,
    new_values_json JSONB,
    metadata_json JSONB,
    ip_address VARCHAR(100),
    user_agent TEXT,
    request_id VARCHAR(100) NOT NULL,
    correlation_id VARCHAR(100) NOT NULL,
    occurred_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX identity_audit_logs_foundation_id_idx ON identity_audit_logs (foundation_id);
CREATE INDEX identity_audit_logs_school_id_idx ON identity_audit_logs (school_id) WHERE school_id IS NOT NULL;
CREATE INDEX identity_audit_logs_actor_user_id_idx ON identity_audit_logs (actor_user_id) WHERE actor_user_id IS NOT NULL;
CREATE INDEX identity_audit_logs_action_idx ON identity_audit_logs (action);
CREATE INDEX identity_audit_logs_entity_idx ON identity_audit_logs (entity_type, entity_id);
CREATE INDEX identity_audit_logs_occurred_at_idx ON identity_audit_logs (occurred_at);

-- +goose Down
DROP TABLE IF EXISTS identity_audit_logs;
DROP TABLE IF EXISTS user_sessions;
DROP TABLE IF EXISTS user_devices;
DROP TABLE IF EXISTS user_role_assignments;
DROP TABLE IF EXISTS role_permissions;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS users;
