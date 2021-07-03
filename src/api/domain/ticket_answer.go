package domain

import (
	"context"
	"github.com/google/uuid"
)

type TicketAnswerValue struct {
	Type   FieldType `json:"type,omitempty"`
	Number float64   `json:"number,omitempty"`
	String string    `json:"string,omitempty"`
	Flag   bool      `json:"flag,omitempty"`
}

type TicketAnswer struct {
	ID             uuid.UUID          `json:"id,omitempty"`
	TicketID       uuid.UUID          `json:"ticketID,omitempty"`
	Title          string             `json:"title,omitempty"`
	ProgrammaticID string             `json:"programmaticID,omitempty"`
	Value          *TicketAnswerValue `json:"value,omitempty"`
}

type CreateTicketAnswerDTO struct {
	TicketID       uuid.UUID   `json:"ticketID,omitempty"`
	SubjectFieldID uuid.UUID   `json:"subjectFieldID,omitempty"`
	Value          interface{} `json:"value,omitempty"`
}

type TicketAnswerEntity interface {
	NewAnswers(answers []*CreateTicketAnswerDTO, subjectFields []*SubjectField) ([]*TicketAnswer, error)
	NewAnswer(dto *CreateTicketAnswerDTO, subjectField *SubjectField) (*TicketAnswer, error)
}

type TicketAnswerRepository interface {
	Get(ctx context.Context, id uuid.UUID) (*TicketAnswer, error)
	Insert(ctx context.Context, tickets []*TicketAnswer) error
	InsertInTransaction(ctx context.Context, tickets []*TicketAnswer, transactionID uuid.UUID) error
}

type TicketAnswerService interface {
	Create(ctx context.Context, answers []*CreateTicketAnswerDTO) ([]*TicketAnswer, error)
}
