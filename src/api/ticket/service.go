package ticket

import (
	"src/domain"
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
	_ticket, err := New(ticket)
	if err != nil {
		return nil, err
	}

	for _, answer := range ticket.Answers {
		answer.TicketID = _ticket.ID
	}

	for _, attachment := range ticket.Attachments {
		attachment.TicketID = _ticket.ID
	}

	_answers, err := service.ticketAnswerService.New(ticket.Answers)
	if err != nil {
		return nil, err
	}

	_attachments, err := service.ticketAttachmentService.New(ticket.Attachments)
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
