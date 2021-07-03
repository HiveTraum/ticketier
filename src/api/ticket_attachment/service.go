package ticket_attachment

import (
	"context"
	"src/domain"
)

type ticketAttachmentService struct {
	entity         domain.TicketAttachmentEntity
	repository     domain.TicketAttachmentRepository
	fileRepository domain.FileRepository
}

func NewTicketAttachmentService(entity domain.TicketAttachmentEntity, repository domain.TicketAttachmentRepository, fileRepository domain.FileRepository) domain.TicketAttachmentService {
	return &ticketAttachmentService{entity: entity, repository: repository, fileRepository: fileRepository}
}

func (service *ticketAttachmentService) Create(ctx context.Context, DTOs []*domain.CreateTicketAttachmentDTO) ([]*domain.TicketAttachment, error) {
	attachments, err := service.entity.NewAttachments(DTOs)
	if err != nil {
		return nil, err
	}

	err = service.fileRepository.UploadAttachments(ctx, attachments)
	if err != nil {
		return nil, err
	}

	err = service.repository.Insert(ctx, attachments)
	if err != nil {
		return nil, err
	}

	return attachments, nil
}
