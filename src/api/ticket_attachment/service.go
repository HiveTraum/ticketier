package ticket_attachment

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/h2non/filetype"
	"src/domain"
)

type ticketAttachmentService struct {
	fileRepository domain.FileRepository
}

func NewTicketAttachmentService(fileRepository domain.FileRepository) domain.TicketAttachmentService {
	return &ticketAttachmentService{fileRepository: fileRepository}
}

func (service *ticketAttachmentService) Create(attachments []*domain.CreateTicketAttachmentDTO) ([]*domain.TicketAttachment, error) {
	_attachments := make([]*domain.TicketAttachment, len(attachments))

	for i, attachment := range attachments {

		meta, err := filetype.Match(attachment.Payload)
		if err != nil {
			return nil, err
		}

		id := uuid.New()

		fileName := id.String()
		if meta != filetype.Unknown {
			fileName = fmt.Sprintf("%s.%s", fileName, meta.Extension)
		}

		path := fmt.Sprintf("attachments/%s/%s", attachment.TicketID.String(), id.String())

		url, err := service.fileRepository.Upload(path, attachment.Payload)
		if err != nil {
			return nil, err
		}

		_attachments[i] = &domain.TicketAttachment{
			ID:        id,
			TicketID:  attachment.TicketID,
			Path:      path,
			MimeType:  meta.MIME.Value,
			Extension: meta.Extension,
			URL:       url,
			Comment:   attachment.Comment,
		}
	}

	return _attachments, nil
}
