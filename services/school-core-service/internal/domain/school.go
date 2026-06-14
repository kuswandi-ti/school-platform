package domain

import (
	"github.com/google/uuid"
	"time"
)

type Actor struct {
	UserID        uuid.UUID
	FoundationID  uuid.UUID
	SchoolID      *uuid.UUID
	Roles         []string
	Permissions   []string
	RequestID     string
	CorrelationID string
	IPAddress     *string
	UserAgent     *string
}
type Foundation struct {
	ID             uuid.UUID
	FoundationCode string
	Name           string
	LegalName      *string
	Address        *string
	Phone          *string
	Email          *string
	LogoFileID     *uuid.UUID
	Timezone       string
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
type School struct {
	ID           uuid.UUID
	FoundationID uuid.UUID
	SchoolCode   string
	Name         string
	SchoolLevel  string
	NPSN         *string
	Address      *string
	Phone        *string
	Email        *string
	LogoFileID   *uuid.UUID
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
