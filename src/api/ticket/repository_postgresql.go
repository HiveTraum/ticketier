package ticket

import (
	"github.com/google/uuid"
	"src/domain"
)

type ticketRepositoryPostgreSQL struct {

}

func NewTicketPostgreSQLRepository() domain.TicketRepository {
	return &ticketRepositoryPostgreSQL{}
}

func (repository *ticketRepositoryPostgreSQL) Save(ticket *domain.Ticket) error {
	panic("implement me")
}

func (repository *ticketRepositoryPostgreSQL) Get(id uuid.UUID) (*domain.Ticket, error) {
	panic("implement me")
}

func (repository *ticketRepositoryPostgreSQL) Fetch() ([]*domain.Ticket, error) {
	panic("implement me")
}
