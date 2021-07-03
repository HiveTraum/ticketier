package ticket_attachment

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/h2non/filetype"
	"src/domain"
)

type ticketAttachmentEntity struct {
}

func NewTicketAttachmentEntity() domain.TicketAttachmentEntity {
	return &ticketAttachmentEntity{}
}

func (entity *ticketAttachmentEntity) NewAttachment(DTO *domain.CreateTicketAttachmentDTO) (*domain.TicketAttachment, error) {
	meta, err := filetype.Match(DTO.Payload)
	if err != nil {
		return nil, err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	fileName := id.String()
	if meta != filetype.Unknown {
		fileName = fmt.Sprintf("%s.%s", fileName, meta.Extension)
	}

	path := fmt.Sprintf("attachments/%s/%s", DTO.TicketID.String(), id.String())

	return &domain.TicketAttachment{
		ID:        id,
		TicketID:  DTO.TicketID,
		Path:      path,
		MimeType:  meta.MIME.Value,
		Extension: meta.Extension,
		Comment:   DTO.Comment,
		Payload:   DTO.Payload,
	}, nil
}

func (entity *ticketAttachmentEntity) NewAttachments(DTOs []*domain.CreateTicketAttachmentDTO) ([]*domain.TicketAttachment, error) {
	attachments := make([]*domain.TicketAttachment, len(DTOs))

	for i, DTO := range DTOs {
		attachment, err := entity.NewAttachment(DTO)
		if err != nil {
			return nil, err
		}

		attachments[i] = attachment
	}

	return attachments, nil
}
