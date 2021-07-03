package domain

import (
	"context"
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

type TicketEntity interface {
	NewTicket(DTO *CreateTicketDTO, subject *Subject) (*Ticket, error)
}

type TicketRepository interface {
	Insert(ctx context.Context, ticket *Ticket) error
	InsertInTransaction(ctx context.Context, ticket *Ticket, transactionID uuid.UUID) error
	Get(id uuid.UUID) (*Ticket, error)
	Select() ([]*Ticket, error)

	Begin(ctx context.Context) (uuid.UUID, error)
	Commit(ctx context.Context, transactionID uuid.UUID) error
	Rollback(ctx context.Context, transactionID uuid.UUID) error
}

type TicketService interface {
	Create(ctx context.Context, DTO *CreateTicketDTO) (*Ticket, error)
}
