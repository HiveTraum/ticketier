package subject

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"src/domain"
	"src/postgresql"
)

type subjectPostgreSQLRepository struct {
	*postgresql.DB
}

func NewSubjectPostgreSQLRepository(db *postgresql.DB) domain.SubjectRepository {
	return &subjectPostgreSQLRepository{DB: db}
}

func (repository *subjectPostgreSQLRepository) List(ctx context.Context) ([]*domain.Subject, error) {
	sql, _, err := sq.Select("id, title, parent_id, created_by").From("subjects").ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := repository.Pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}

	return scanRows(rows)
}

func (repository *subjectPostgreSQLRepository) Get(ctx context.Context, id uuid.UUID) (*domain.Subject, error) {
	sql, _, err := sq.Select("id, title, parent_id, created_by").From("subjects").Limit(1).ToSql()
	if err != nil {
		return nil, err
	}

	return scanRow(repository.DB.Pool.QueryRow(ctx, sql))
}

func (repository *subjectPostgreSQLRepository) Insert(ctx context.Context, subjects []*domain.Subject) error {
	return insert(ctx, repository.DB.Pool, subjects)
}

func (repository *subjectPostgreSQLRepository) InsertInTransaction(ctx context.Context, subjects []*domain.Subject, transactionID uuid.UUID) error {
	tx, err := repository.DB.GetTransaction(transactionID)
	if err != nil {
		return err
	}

	return insert(ctx, tx, subjects)
}

func insert(ctx context.Context, connection postgresql.Connection, subjects []*domain.Subject) error {
	batch := &pgx.Batch{}
	for _, s := range subjects {
		batch.Queue("INSERT INTO subjects(id, title, parent_id, created_by) VALUES ($1, $2, $3, $4);", s.ID, s.Title, s.ParentID, s.CreatedBy)
	}

	_, err := connection.SendBatch(ctx, batch).Exec()
	return err
}

func scanRow(row pgx.Row) (*domain.Subject, error) {
	s := &domain.Subject{}
	err := row.Scan(s.ID, s.Title, s.ParentID, s.CreatedBy)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return s, nil
}

func scanRows(rows pgx.Rows) ([]*domain.Subject, error) {
	var subjects []*domain.Subject

	for rows.Next() {
		subject, err := scanRow(rows)
		if err != nil {
			return nil, err
		}

		subjects = append(subjects, subject)
	}

	return subjects, nil
}
