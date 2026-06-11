package actorcontext

type ActorContext struct {
	UserID        string
	FoundationID  string
	SchoolID      *string
	Roles         []string
	Permissions   []string
	Scope         map[string]any
	RequestID     string
	CorrelationID string
}

type RequestContext struct {
	RequestID     string
	CorrelationID string
}
