package domain

import (
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
}

type CreateTicketAttachmentDTO struct {
	TicketID uuid.UUID `json:"ticketID,omitempty"`
	Payload  []byte    `json:"payload,omitempty"`
	Comment  string    `json:"comment,omitempty"`
}

type TicketAttachmentRepository interface {
	Get(id uuid.UUID) (*TicketAttachment, error)
}

type TicketAttachmentService interface {
	Create(attachments []*CreateTicketAttachmentDTO) ([]*TicketAttachment, error)
}
