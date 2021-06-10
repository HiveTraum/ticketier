package domain

import (
	"github.com/google/uuid"
)

type TicketAnswerValue struct {
	Type   FieldType `json:"type,omitempty"`
	Number *float64  `json:"number,omitempty"`
	String *string   `json:"string,omitempty"`
	Flag   *bool     `json:"flag,omitempty"`
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

type TicketAnswerRepository interface {
	Get(id uuid.UUID) (*TicketAnswer, error)
}

type TicketAnswerService interface {
	Create(answers []*CreateTicketAnswerDTO) ([]*TicketAnswer, error)
}
