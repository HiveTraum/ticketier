package ticket_attachment

import (
	"src/domain"
)

type ticketAttachmentService struct {
	repository     domain.TicketAttachmentRepository
	fileRepository domain.FileRepository
}

func NewTicketAttachmentService(repository domain.TicketAttachmentRepository, fileRepository domain.FileRepository) domain.TicketAttachmentService {
	return &ticketAttachmentService{repository: repository, fileRepository: fileRepository}
}

func (service *ticketAttachmentService) New(attachments []*domain.CreateTicketAttachmentDTO) ([]*domain.TicketAttachment, error) {
	_attachments := make([]*domain.TicketAttachment, len(attachments))

	for i, attachment := range attachments {
		_attachment, err := New(attachment)
		if err != nil {
			return nil, err
		}

		url, err := service.fileRepository.URL(_attachment.Path)
		if err != nil {
			return nil, err
		}

		_attachment.URL = url
		_attachments[i] = _attachment
	}

	return _attachments, nil
}

func (service *ticketAttachmentService) Create(attachments []*domain.CreateTicketAttachmentDTO) ([]*domain.TicketAttachment, error) {
	_attachments, err := service.New(attachments)
	if err != nil {
		return nil, err
	}

	err = service.fileRepository.Upload(ticketAttachmentsToUploads(_attachments))
	if err != nil {
		return nil, err
	}

	err = service.repository.Save(_attachments)
	if err != nil {
		return nil, err
	}

	return _attachments, nil
}

func ticketAttachmentsToUploads(attachments []*domain.TicketAttachment) []*domain.UploadFileDTO {
	files := make([]*domain.UploadFileDTO, len(attachments))
	for i, attachment := range attachments {
		files[i] = &domain.UploadFileDTO{
			Payload: attachment.Payload,
			Path:    attachment.Path,
		}
	}

	return files
}
