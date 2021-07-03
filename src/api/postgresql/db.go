package postgresql

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Connection interface {
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row
}

type DB struct {
	Pool         *pgxpool.Pool
	transactions map[uuid.UUID]pgx.Tx
}

func New(uri string) (*DB, error) {
	pool, err := pgxpool.Connect(context.Background(), uri)
	if err != nil {
		return nil, err
	}

	return &DB{
		Pool:         pool,
		transactions: map[uuid.UUID]pgx.Tx{},
	}, nil
}

func (postgresql *DB) Begin(ctx context.Context) (uuid.UUID, error) {
	transactionID, err := uuid.NewRandom()
	if err != nil {
		return uuid.Nil, err
	}

	tx, err := postgresql.Pool.Begin(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	postgresql.transactions[transactionID] = tx
	return transactionID, nil
}

func (postgresql *DB) Rollback(ctx context.Context, transactionID uuid.UUID) error {
	if tx, ok := postgresql.transactions[transactionID]; !ok {
		return TransactionNotFound
	} else {
		delete(postgresql.transactions, transactionID)
		return tx.Rollback(ctx)
	}
}

func (postgresql *DB) Commit(ctx context.Context, transactionID uuid.UUID) error {
	if tx, ok := postgresql.transactions[transactionID]; !ok {
		return TransactionNotFound
	} else {
		delete(postgresql.transactions, transactionID)
		return tx.Commit(ctx)
	}
}

func (postgresql *DB) GetTransaction(transactionID uuid.UUID) (pgx.Tx, error) {
	if tx, ok := postgresql.transactions[transactionID]; !ok {
		return nil, TransactionNotFound
	} else {
		return tx, nil
	}
}

func (postgresql *DB) Close() {
	postgresql.Pool.Close()
}
