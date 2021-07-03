package ticket

import (
	"context"
	"github.com/google/uuid"
	"src/domain"
)

type ticketService struct {
	entity                   domain.TicketEntity
	repository               domain.TicketRepository
	answerEntity             domain.TicketAnswerEntity
	attachmentEntity         domain.TicketAttachmentEntity
	attachmentRepository     domain.TicketAttachmentRepository
	attachmentFileRepository domain.FileRepository
	answerRepository         domain.TicketAnswerRepository
	subjectFieldRepository   domain.SubjectFieldRepository
	subjectRepository        domain.SubjectRepository
}

func NewTicketService(entity domain.TicketEntity, repository domain.TicketRepository, answerEntity domain.TicketAnswerEntity, answerRepository domain.TicketAnswerRepository, attachmentEntity domain.TicketAttachmentEntity, attachmentRepository domain.TicketAttachmentRepository, attachmentFileRepository domain.FileRepository, subjectFieldRepository domain.SubjectFieldRepository, subjectRepository domain.SubjectRepository) domain.TicketService {
	return &ticketService{entity: entity, repository: repository, answerEntity: answerEntity, answerRepository: answerRepository, attachmentEntity: attachmentEntity, attachmentRepository: attachmentRepository, attachmentFileRepository: attachmentFileRepository, subjectFieldRepository: subjectFieldRepository, subjectRepository: subjectRepository}
}

func (service *ticketService) Create(ctx context.Context, DTO *domain.CreateTicketDTO) (*domain.Ticket, error) {
	subject, err := service.subjectRepository.Get(ctx, DTO.SubjectID)
	if err != nil {
		return nil, err
	}

	ticket, err := service.entity.NewTicket(DTO, subject)
	if err != nil {
		return nil, err
	}

	for _, DTO := range DTO.Answers {
		DTO.TicketID = ticket.ID
	}

	for _, DTO := range DTO.Attachments {
		DTO.TicketID = ticket.ID
	}

	subjectFields, err := service.subjectFieldRepository.SelectByTicketAnswers(ctx, DTO.Answers)
	if err != nil {
		return nil, err
	}

	answers, err := service.answerEntity.NewAnswers(DTO.Answers, subjectFields)
	if err != nil {
		return nil, err
	}

	attachments, err := service.attachmentEntity.NewAttachments(DTO.Attachments)
	if err != nil {
		return nil, err
	}

	err = service.attachmentFileRepository.UploadAttachments(ctx, attachments)

	transaction, err := service.repository.Begin(ctx)
	if err != nil {
		return nil, err
	}

	err = service.save(ctx, ticket, answers, attachments, transaction)
	if err != nil {
		err = service.repository.Rollback(ctx, transaction)
		if err != nil {
			return nil, err
		}

		return nil, err
	}

	return ticket, service.repository.Commit(ctx, transaction)
}

func (service *ticketService) save(ctx context.Context, ticket *domain.Ticket, answers []*domain.TicketAnswer, attachments []*domain.TicketAttachment, transactionID uuid.UUID) error {
	err := service.repository.InsertInTransaction(ctx, ticket, transactionID)
	if err != nil {
		return err
	}

	err = service.answerRepository.InsertInTransaction(ctx, answers, transactionID)
	if err != nil {
		return err
	}

	return service.attachmentRepository.InsertInTransaction(ctx, attachments, transactionID)
}
