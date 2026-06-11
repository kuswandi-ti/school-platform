package events

import "time"

type Envelope struct {
	EventID       string    `json:"event_id"`
	EventType     string    `json:"event_type"`
	EventVersion  int       `json:"event_version"`
	SourceService string    `json:"source_service"`
	OccurredAt    time.Time `json:"occurred_at"`
	PublishedAt   time.Time `json:"published_at"`
	RequestID     string    `json:"request_id"`
	CorrelationID string    `json:"correlation_id"`
	Actor         Actor     `json:"actor"`
	Tenant        Tenant    `json:"tenant"`
	Entity        Entity    `json:"entity"`
	Payload       any       `json:"payload"`
	Metadata      any       `json:"metadata"`
}

type Actor struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

type Tenant struct {
	FoundationID string  `json:"foundation_id"`
	SchoolID     *string `json:"school_id"`
}

type Entity struct {
	EntityType string `json:"entity_type"`
	EntityID   string `json:"entity_id"`
}
