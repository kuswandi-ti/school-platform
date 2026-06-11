package audit

import "time"

type Record struct {
	ActorUserID   string
	FoundationID  string
	SchoolID      *string
	Action        string
	EntityType    string
	EntityID      string
	OldValues     any
	NewValues     any
	Reason        string
	RequestID     string
	CorrelationID string
	OccurredAt    time.Time
}
