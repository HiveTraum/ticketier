package domain

import "github.com/google/uuid"

type SubjectField struct {
	ID             uuid.UUID
	SubjectID      uuid.UUID
	Title          string
	Required       bool
	ProgrammaticID string
	Order          int
	Type           FieldType
}

type SubjectFieldRepository interface {
	GetBySubjectID(subjectID uuid.UUID) ([]*SubjectField, error)
	Get(id uuid.UUID) (*SubjectField, error)
}
