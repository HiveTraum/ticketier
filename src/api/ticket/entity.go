package ticket

import (
	"github.com/google/uuid"
	"src/domain"
	"time"
)

type ticketEntity struct {
}

func NewTicketEntity() domain.TicketEntity {
	return &ticketEntity{}
}

func (entity *ticketEntity) NewTicket(DTO *domain.CreateTicketDTO, subject *domain.Subject) (*domain.Ticket, error) {
	if subject == nil {
		return nil, domain.SubjectNotFound(DTO.SubjectID)
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &domain.Ticket{
		ID:        id,
		SubjectID: DTO.SubjectID,
		CreatedBy: DTO.CreatedBy,
		CreatedAt: time.Now(),
	}, nil
}
