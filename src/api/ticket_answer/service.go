package ticket_answer

import (
	"context"
	"src/domain"
)

type ticketAnswerService struct {
	entity                 domain.TicketAnswerEntity
	repository             domain.TicketAnswerRepository
	subjectFieldRepository domain.SubjectFieldRepository
}

func NewTicketAnswerService(entity domain.TicketAnswerEntity, repository domain.TicketAnswerRepository, subjectFieldRepository domain.SubjectFieldRepository) domain.TicketAnswerService {
	return &ticketAnswerService{entity: entity, repository: repository, subjectFieldRepository: subjectFieldRepository}
}

func (service *ticketAnswerService) Create(ctx context.Context, DTOs []*domain.CreateTicketAnswerDTO) ([]*domain.TicketAnswer, error) {
	subjectFields, err := service.subjectFieldRepository.SelectByTicketAnswers(ctx, DTOs)
	if err != nil {
		return nil, err
	}

	answers, err := service.entity.NewAnswers(DTOs, subjectFields)
	if err != nil {
		return nil, err
	}

	err = service.repository.Insert(ctx, answers)
	if err != nil {
		return nil, err
	}

	return answers, nil
}
