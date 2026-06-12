package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Email        string
	Phone        *string
	PasswordHash string
	DisplayName  string
	AvatarFileID *uuid.UUID
	Status       string
	LastLoginAt  *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
