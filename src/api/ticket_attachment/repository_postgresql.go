package ticket_attachment

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"src/domain"
	"src/postgresql"
)

type ticketAttachmentRepositoryPostgreSQL struct {
	db *postgresql.DB
}

func NewTicketAttachmentRepositoryPostgreSQL(db *postgresql.DB) domain.TicketAttachmentRepository {
	return &ticketAttachmentRepositoryPostgreSQL{db: db}
}

func (repository *ticketAttachmentRepositoryPostgreSQL) Get(ctx context.Context, id uuid.UUID) (*domain.TicketAttachment, error) {
	panic("implement me")
}

func (repository *ticketAttachmentRepositoryPostgreSQL) Insert(ctx context.Context, attachments []*domain.TicketAttachment) error {
	return insert(ctx, repository.db.Pool, attachments)
}

func (repository *ticketAttachmentRepositoryPostgreSQL) InsertInTransaction(ctx context.Context, attachments []*domain.TicketAttachment, transactionID uuid.UUID) error {
	tx, err := repository.db.GetTransaction(transactionID)
	if err != nil {
		return err
	}

	return insert(ctx, tx, attachments)
}

func insert(ctx context.Context, db postgresql.Connection, attachments []*domain.TicketAttachment) error {
	batch := &pgx.Batch{}
	for _, a := range attachments {
		batch.Queue("INSERT INTO ticket_attachments(id, ticket_id, comment, path, mimetype, extension) VALUES ($1, $2, $3, $4, $5, $6)", a.ID, a.TicketID, a.Comment, a.Path, a.MimeType, a.Extension)
	}

	_, err := db.SendBatch(ctx, batch).Exec()
	return err
}
