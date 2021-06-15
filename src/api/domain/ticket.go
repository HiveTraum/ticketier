package domain

import (
	"github.com/google/uuid"
	"time"
)

type Ticket struct {
	ID          uuid.UUID           `json:"id,omitempty"`
	SubjectID   uuid.UUID           `json:"subjectID,omitempty"`
	CreatedBy   uuid.UUID           `json:"createdBy,omitempty"`
	CreatedAt   time.Time           `json:"createdAt"`
	Answers     []*TicketAnswer     `json:"answers,omitempty"`
	Attachments []*TicketAttachment `json:"attachments,omitempty"`
}

type CreateTicketDTO struct {
	SubjectID   uuid.UUID                    `json:"subjectID,omitempty"`
	CreatedBy   uuid.UUID                    `json:"createdBy,omitempty"`
	Answers     []*CreateTicketAnswerDTO     `json:"answers,omitempty"`
	Attachments []*CreateTicketAttachmentDTO `json:"attachments,omitempty"`
}

type TicketRepository interface {
	Save(ticket *Ticket) error
	Get(id uuid.UUID) (*Ticket, error)
	Fetch() ([]*Ticket, error)
}

type TicketService interface {
	Create(ticket *CreateTicketDTO) (*Ticket, error)
}