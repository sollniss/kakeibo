package mysql

import (
	"context"
	"database/sql"

	"github.com/sollniss/kakeibo/domain"
)

type transferWriter struct {
	conn *sql.DB
	transferWriterQueries
}

// transferQueries holds all single (atomic) queries and is not exposed
type transferWriterQueries struct {
	db dbtx
}

// NewSendJobWriter creates a new user repository with read only permissions (CQRS reader).
func NewTransferWriter(db *DB) *transferWriter {
	if db == nil {
		panic("db can not be nil")
	}
	return &transferWriter{
		conn: db.db,
		transferWriterQueries: transferWriterQueries{
			db: db.db,
		},
	}
}

// Add adds a new transfer to the database.
func (q *transferWriterQueries) Add(ctx context.Context, t domain.Transfer) error {
	_, err := q.db.ExecContext(ctx,
		`INSERT INTO transfer
			(transfer_id, person, amount, comment, category, transfer_date)
		VALUES
			(?, ?, ?, ?, ?)`,
		t.ID, t.Person, t.Amount, t.Comment, t.Category, t.Date)

	if err != nil {
		return err
	}

	return nil
}

// Update updates an existing transfer in the database.
func (q *transferWriterQueries) Update(ctx context.Context, t domain.Transfer) error {
	res, err := q.db.ExecContext(ctx,
		`UPDATE transfer
		SET
			person = ?,
			amount = ?,
			comment = ?,
			category = ?,
			transfer_date = ?
		WHERE
			transfer_id = ?`,
		t.Person, t.Amount, t.Comment, t.Category, t.Date, t.ID)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return domain.ErrNotFound
	}

	return nil
}

// Delete deletes an existing transfer from the database.
func (q *transferWriterQueries) Delete(ctx context.Context, id domain.ID) error {
	res, err := q.db.ExecContext(ctx,
		`DELETE FROM transfer WHERE transfer_id = ?`, id)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return domain.ErrNotFound
	}

	return nil
}
