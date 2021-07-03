package ticket

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgtype/pgxtype"
	"src/domain"
	"src/postgresql"
)

type ticketRepositoryPostgreSQL struct {
	*postgresql.DB
}

func NewTicketPostgreSQLRepository(db *postgresql.DB) domain.TicketRepository {
	return &ticketRepositoryPostgreSQL{DB: db}
}

func (repository *ticketRepositoryPostgreSQL) InsertInTransaction(ctx context.Context, ticket *domain.Ticket, transactionID uuid.UUID) error {
	tx, err := repository.DB.GetTransaction(transactionID)
	if err != nil {
		return err
	}

	return insert(ctx, tx, ticket)
}

func (repository *ticketRepositoryPostgreSQL) Insert(ctx context.Context, ticket *domain.Ticket) error {
	return insert(ctx, repository.DB.Pool, ticket)
}

func insert(ctx context.Context, db pgxtype.Querier, ticket *domain.Ticket) error {
	_, err := db.Exec(ctx, "INSERT INTO tickets(id, subject_id, created_by, created_at) VALUES ($1, $2, $3, $4)", ticket.ID, ticket.SubjectID, ticket.CreatedBy, ticket.CreatedAt)
	return err
}

func (repository *ticketRepositoryPostgreSQL) Get(id uuid.UUID) (*domain.Ticket, error) {
	panic("implement me")
}

func (repository *ticketRepositoryPostgreSQL) Select() ([]*domain.Ticket, error) {
	panic("implement me")
}
