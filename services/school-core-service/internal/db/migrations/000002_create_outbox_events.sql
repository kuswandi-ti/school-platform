-- +goose Up
CREATE TABLE outbox_events (
    id UUID PRIMARY KEY,
    event_id UUID NOT NULL,
    event_type VARCHAR(150) NOT NULL,
    event_version INT NOT NULL,
    aggregate_type VARCHAR(100) NOT NULL,
    aggregate_id UUID NOT NULL,
    payload_json JSONB NOT NULL,
    status VARCHAR(30) NOT NULL DEFAULT 'pending',
    retry_count INT NOT NULL DEFAULT 0,
    next_retry_at TIMESTAMPTZ,
    published_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT outbox_events_event_id_unique UNIQUE (event_id),
    CONSTRAINT outbox_events_version_check CHECK (event_version > 0),
    CONSTRAINT outbox_events_status_check CHECK (status IN ('pending', 'processing', 'published', 'failed')),
    CONSTRAINT outbox_events_retry_count_check CHECK (retry_count >= 0)
);

CREATE INDEX outbox_events_pending_idx
    ON outbox_events (status, next_retry_at, created_at)
    WHERE status IN ('pending', 'failed');
CREATE INDEX outbox_events_aggregate_idx ON outbox_events (aggregate_type, aggregate_id);

-- +goose Down
DROP TABLE IF EXISTS outbox_events;
