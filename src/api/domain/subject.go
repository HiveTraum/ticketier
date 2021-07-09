package domain

import (
	"context"
	"github.com/google/uuid"
)

type Subject struct {
	ID        uuid.UUID
	Title     string
	ParentID  uuid.UUID
	CreatedBy uuid.UUID
	Fields    []*SubjectField
}

type CreateSubjectDTO struct {
	Title     string                   `json:"title,omitempty"`
	ParentID  uuid.UUID                `json:"parentID,omitempty"`
	CreatedBy uuid.UUID                `json:"createdBy,omitempty"`
	Fields    []*CreateSubjectFieldDTO `json:"fields,omitempty"`
}

type SubjectEntity interface {
	NewSubjects(DTOs []*CreateSubjectDTO) ([]*Subject, error)
	NewSubject(DTO *CreateSubjectDTO) (*Subject, error)
}

type SubjectRepository interface {
	Insert(ctx context.Context, subjects []*Subject) error
	InsertInTransaction(ctx context.Context, subjects []*Subject, transactionID uuid.UUID) error
	Get(ctx context.Context, id uuid.UUID) (*Subject, error)
	List(ctx context.Context) ([]*Subject, error)

	Begin(ctx context.Context) (uuid.UUID, error)
	Rollback(ctx context.Context, transactionID uuid.UUID) error
	Commit(ctx context.Context, transactionID uuid.UUID) error
}

type SubjectService interface {
	Create(ctx context.Context, DTOs []*CreateSubjectDTO) ([]*Subject, error)
	List(ctx context.Context) ([]*Subject, error)
}
