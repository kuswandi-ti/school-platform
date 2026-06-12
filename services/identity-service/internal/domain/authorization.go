package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type RoleAssignment struct {
	ID           uuid.UUID
	Role         string
	FoundationID uuid.UUID
	SchoolID     *uuid.UUID
	ClassID      *uuid.UUID
	StudentID    *uuid.UUID
	EmployeeID   *uuid.UUID
	SubjectID    *uuid.UUID
	Scope        json.RawMessage
	StartsAt     *time.Time
	EndsAt       *time.Time
	Permissions  []string
}

type UserContext struct {
	UserID      uuid.UUID
	Roles       []string
	Permissions []string
	Assignments []RoleAssignment
}
