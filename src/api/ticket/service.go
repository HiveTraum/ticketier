package ticket

import (
	"github.com/google/uuid"
	"src/domain"
	"time"
)

type ticketService struct {
	repository              domain.TicketRepository
	ticketAnswerService     domain.TicketAnswerService
	ticketAttachmentService domain.TicketAttachmentService
	subjectFieldRepository  domain.SubjectFieldRepository
}

func NewTicketService(repository domain.TicketRepository) domain.TicketService {
	return &ticketService{repository: repository}
}

func (service *ticketService) Create(ticket *domain.CreateTicketDTO) (*domain.Ticket, error) {
	ticketID := uuid.New()

	_ticket := &domain.Ticket{
		ID:        ticketID,
		SubjectID: ticket.SubjectID,
		CreatedBy: ticket.CreatedBy,
		CreatedAt: time.Now(),
	}

	for _, answer := range ticket.Answers {
		answer.TicketID = ticketID
	}

	_answers, err := service.ticketAnswerService.Create(ticket.Answers)
	if err != nil {
		return nil, err
	}

	_attachments, err := service.ticketAttachmentService.Create(ticket.Attachments)
	if err != nil {
		return nil, err
	}

	err = service.repository.Save(_ticket)
	if err != nil {
		return nil, err
	}

	_ticket.Answers = _answers
	_ticket.Attachments = _attachments

	return _ticket, nil
}
