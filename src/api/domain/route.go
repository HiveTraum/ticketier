package domain

import "github.com/google/uuid"

type Route struct {
	ID         uuid.UUID
	ExecutorID uuid.UUID
}
