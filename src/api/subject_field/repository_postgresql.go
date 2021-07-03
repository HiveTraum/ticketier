package subject_field

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"src/domain"
	"src/postgresql"
)

type subjectFieldPostgreSQLRepository struct {
	db *postgresql.DB
}

func NewSubjectFieldPostgreSQLRepository(db *postgresql.DB) domain.SubjectFieldRepository {
	return &subjectFieldPostgreSQLRepository{db: db}
}

func (repository *subjectFieldPostgreSQLRepository) GetBySubjectID(ctx context.Context, subjectID uuid.UUID) ([]*domain.SubjectField, error) {
	panic("implement me")
}

func (repository *subjectFieldPostgreSQLRepository) Get(ctx context.Context, id uuid.UUID) (*domain.SubjectField, error) {
	panic("implement me")
}

func (repository *subjectFieldPostgreSQLRepository) Select(ctx context.Context, query *domain.SubjectFieldQuery) ([]*domain.SubjectField, error) {
	sql := sq.Select("id", "subject_id", "title", "required", "programmatic_id", "order", "type").From("subject_fields")
	if query != nil {
		if len(query.ID) > 0 {
			predicate, args := postgresql.In("id", query.ID)
			sql = sql.Where(predicate, args...)
		}
	}

	_sql, args, err := sql.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := repository.db.Pool.Query(ctx, _sql, args...)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	return scanRows(rows)
}

func (repository *subjectFieldPostgreSQLRepository) SelectByTicketAnswers(ctx context.Context, answers []*domain.CreateTicketAnswerDTO) ([]*domain.SubjectField, error) {
	if len(answers) <= 0 {
		return make([]*domain.SubjectField, 0), nil
	}

	identifiers := make([]uuid.UUID, len(answers))
	for i, answer := range answers {
		identifiers[i] = answer.SubjectFieldID
	}

	return repository.Select(ctx, &domain.SubjectFieldQuery{ID: identifiers})
}

func (repository *subjectFieldPostgreSQLRepository) Insert(ctx context.Context, fields []*domain.SubjectField) error {
	return insert(ctx, repository.db.Pool, fields)
}

func (repository *subjectFieldPostgreSQLRepository) InsertInTransaction(ctx context.Context, fields []*domain.SubjectField, transactionID uuid.UUID) error {
	tx, err := repository.db.GetTransaction(transactionID)
	if err != nil {
		return err
	}

	return insert(ctx, tx, fields)
}

func insert(ctx context.Context, connection postgresql.Connection, fields []*domain.SubjectField) error {
	batch := &pgx.Batch{}

	for _, f := range fields {
		batch.Queue("INSERT INTO subject_fields(id, subject_id, title, required, programmatic_id, \"order\", type) VALUES ($1, $2, $3, $4, $5, $6, $7);", f.ID, f.SubjectID, f.Title, f.Required, f.ProgrammaticID, f.Order, f.Type)
	}

	_, err := connection.SendBatch(ctx, batch).Exec()
	return err
}

func scanRow(row pgx.Row) (*domain.SubjectField, error) {
	f := &domain.SubjectField{}

	err := row.Scan(f.ID, f.SubjectID, f.Title, f.Required, f.ProgrammaticID, f.Order, f.Type)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func scanRows(rows pgx.Rows) ([]*domain.SubjectField, error) {
	var fields []*domain.SubjectField

	for rows.Next() {
		field, err := scanRow(rows)
		if err != nil {
			return nil, err
		}

		fields = append(fields, field)
	}

	return fields, nil
}
