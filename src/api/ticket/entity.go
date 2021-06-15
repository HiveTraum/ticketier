package ticket

import (
	"github.com/google/uuid"
	"src/domain"
	"time"
)

func New(dto *domain.CreateTicketDTO) (*domain.Ticket, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &domain.Ticket{
		ID:        id,
		SubjectID: dto.SubjectID,
		CreatedBy: dto.CreatedBy,
		CreatedAt: time.Now(),
	}, nil
}
