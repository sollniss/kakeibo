package mysql

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/sollniss/kakeibo/domain"
)

type transferReader struct {
	conn *sql.DB
	transferReaderQueries
}

// transferQueries holds all single (atomic) queries and is not exposed
type transferReaderQueries struct {
	db dbtx
}

// NewSendJobReader creates a new user repository with read only permissions (CQRS reader).
func NewTransferReader(db *DB) *transferReader {
	if db == nil {
		panic("db can not be nil")
	}
	return &transferReader{
		conn: db.db,
		transferReaderQueries: transferReaderQueries{
			db: db.db,
		},
	}
}

// IncomesForMonth returns all incomes for the given month and year.
func (q *transferReaderQueries) IncomesForMonth(ctx context.Context, month time.Month, year int) ([]domain.Transfer, error) {
	rows, err := q.db.QueryContext(ctx,
		`SELECT
			transfer_id, person, amount, comment, category, transfer_date
		FROM
			transfer
		WHERE
			WHERE amount > 0 AND MONTH(transfer_date) = ? AND YEAR(transfer_date) = ?`, month, year)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// TODO: return empty slice instead?
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	var transfers []domain.Transfer
	for rows.Next() {
		var t domain.Transfer
		err = rows.Scan(&t.ID, &t.Person, &t.Amount, &t.Comment, &t.Category, &t.Date)
		if err != nil {
			return nil, err
		}
		transfers = append(transfers, t)
	}

	return transfers, nil
}

// ExpensesForMonth returns all expenses for the given month and year.
func (q *transferReaderQueries) ExpensesForMonth(ctx context.Context, month time.Month, year int) ([]domain.Transfer, error) {
	rows, err := q.db.QueryContext(ctx,
		`SELECT
			transfer_id, person, amount, comment, category, transfer_date
		FROM
			transfer
		WHERE
			WHERE amount < 0 AND MONTH(transfer_date) = ? AND YEAR(transfer_date) = ?`, month, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transfers []domain.Transfer
	for rows.Next() {
		var t domain.Transfer
		err = rows.Scan(&t.ID, &t.Person, &t.Amount, &t.Comment, &t.Category, &t.Date)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				// TODO: return empty slice instead?
				return nil, nil
			}
			return nil, err
		}
		transfers = append(transfers, t)
	}

	return transfers, nil
}
