package domain

import (
	"context"
	"github.com/google/uuid"
)

type SubjectField struct {
	ID             uuid.UUID
	SubjectID      uuid.UUID
	Title          string
	Required       bool
	ProgrammaticID string
	Order          int
	Type           FieldType
}

type CreateSubjectFieldDTO struct {
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

type SubjectFieldEntity interface {
	NewFields(DTOs []*CreateSubjectFieldDTO) ([]*SubjectField, error)
	NewField(DTO *CreateSubjectFieldDTO) (*SubjectField, error)
}

type SubjectFieldRepository interface {
	GetBySubjectID(ctx context.Context, subjectID uuid.UUID) ([]*SubjectField, error)
	Get(ctx context.Context, id uuid.UUID) (*SubjectField, error)
	Select(ctx context.Context, query *SubjectFieldQuery) ([]*SubjectField, error)
	SelectByTicketAnswers(ctx context.Context, answers []*CreateTicketAnswerDTO) ([]*SubjectField, error)
	Insert(ctx context.Context, fields []*SubjectField) error
	InsertInTransaction(ctx context.Context, fields []*SubjectField, transactionID uuid.UUID) error
}

type SubjectFieldService interface {
	Create(ctx context.Context, DTOs []*CreateSubjectFieldDTO) ([]*SubjectField, error)
}
