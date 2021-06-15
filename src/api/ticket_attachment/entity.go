package ticket_attachment

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/h2non/filetype"
	"src/domain"
)

func New(dto *domain.CreateTicketAttachmentDTO) (*domain.TicketAttachment, error) {
	meta, err := filetype.Match(dto.Payload)
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

	path := fmt.Sprintf("attachments/%s/%s", dto.TicketID.String(), id.String())

	return &domain.TicketAttachment{
		ID:        id,
		TicketID:  dto.TicketID,
		Path:      path,
		MimeType:  meta.MIME.Value,
		Extension: meta.Extension,
		Comment:   dto.Comment,
		Payload:   dto.Payload,
	}, nil
}
