package domain

import "github.com/google/uuid"

type Subject struct {
	ID        uuid.UUID
	Title     string
	ParentID  uuid.UUID
	CreatedBy uuid.UUID
}
