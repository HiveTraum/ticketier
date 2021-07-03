package domain

import (
	"context"
	"github.com/google/uuid"
	"net/url"
)

type TicketAttachment struct {
	ID        uuid.UUID `json:"id,omitempty"`
	TicketID  uuid.UUID `json:"ticketID,omitempty"`
	Comment   string    `json:"comment,omitempty"`
	Path      string    `json:"path,omitempty"`
	MimeType  string    `json:"mimeType,omitempty"`
	Extension string    `json:"extension,omitempty"`
	URL       url.URL   `json:"url"`
	Payload   []byte
}

type CreateTicketAttachmentDTO struct {
	TicketID uuid.UUID `json:"ticketID,omitempty"`
	Payload  []byte    `json:"payload,omitempty"`
	Comment  string    `json:"comment,omitempty"`
}

type TicketAttachmentEntity interface {
	NewAttachment(DTO *CreateTicketAttachmentDTO) (*TicketAttachment, error)
	NewAttachments(DTOs []*CreateTicketAttachmentDTO) ([]*TicketAttachment, error)
}

type TicketAttachmentRepository interface {
	Get(ctx context.Context, id uuid.UUID) (*TicketAttachment, error)
	Insert(ctx context.Context, attachments []*TicketAttachment) error
	InsertInTransaction(ctx context.Context, attachments []*TicketAttachment, transactionID uuid.UUID) error
}

type TicketAttachmentService interface {
	Create(ctx context.Context, attachments []*CreateTicketAttachmentDTO) ([]*TicketAttachment, error)
}
