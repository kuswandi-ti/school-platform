package numbering

type Scope struct {
	FoundationID string
	SchoolID     *string
	SystemKey    string
	PeriodKey    string
}

type GeneratedNumber struct {
	Value string
	Scope Scope
}
