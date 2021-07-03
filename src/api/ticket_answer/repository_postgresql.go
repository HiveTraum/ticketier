package ticket_answer

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"src/domain"
	"src/postgresql"
)

type ticketAnswerRepositoryPostgreSQL struct {
	db *postgresql.DB
}

func NewTicketAnswerRepositoryPostgreSQL(db *postgresql.DB) domain.TicketAnswerRepository {
	return &ticketAnswerRepositoryPostgreSQL{db: db}
}

func (repository *ticketAnswerRepositoryPostgreSQL) Get(ctx context.Context, id uuid.UUID) (*domain.TicketAnswer, error) {
	panic("implement me")
}

func (repository *ticketAnswerRepositoryPostgreSQL) Insert(ctx context.Context, answers []*domain.TicketAnswer) error {
	return insert(ctx, repository.db.Pool, answers)
}

func (repository *ticketAnswerRepositoryPostgreSQL) InsertInTransaction(ctx context.Context, answers []*domain.TicketAnswer, transactionID uuid.UUID) error {
	tx, err := repository.db.GetTransaction(transactionID)
	if err != nil {
		return err
	}

	return insert(ctx, tx, answers)
}

func insert(ctx context.Context, db postgresql.Connection, answers []*domain.TicketAnswer) error {
	batch := &pgx.Batch{}
	for _, a := range answers {
		batch.Queue("INSERT INTO ticket_answers(id, ticket_id, title, programmatic_id, value) VALUES ($1, $2, $3, $4, $5)", a.ID, a.TicketID, a.Title, a.ProgrammaticID, a.Value)
	}

	_, err := db.SendBatch(ctx, batch).Exec()
	return err
}
