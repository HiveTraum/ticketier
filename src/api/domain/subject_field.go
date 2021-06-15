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

type SubjectFieldQuery struct {
	ID []uuid.UUID
}

type SubjectFieldRepository interface {
	GetBySubjectID(subjectID uuid.UUID) ([]*SubjectField, error)
	Get(id uuid.UUID) (*SubjectField, error)
	Fetch(query *SubjectFieldQuery) ([]*SubjectField, error)
	FetchByIdentifiers(identifiers []uuid.UUID) (map[uuid.UUID]*SubjectField, error)
}
