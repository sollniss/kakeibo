package mysql

import (
	"context"
	"database/sql"
	"time"
)

type dbtx interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...any) *sql.Row
}

// DB holds the database connection.
type DB struct {
	db  *sql.DB
	dsn string
}

// NewDB creates a new DB object.
// Panics if the provided DSN is empty.
func NewDB(dsn string) *DB {
	if dsn == "" {
		panic("dsn required")
	}

	db := &DB{
		dsn: dsn,
	}
	return db
}

// Open opens a new database connection and pings it.
// The following default are used for the connection.
//
//	SetMaxOpenConns(25)
//	SetMaxIdleConns(25)
//	SetConnMaxLifetime(5 * time.Minute)
//
// The ping timeout is 3 seconds.
func (db *DB) Open() (err error) {
	db.db, err = sql.Open("mysql", db.dsn)
	if err != nil {
		return err
	}

	// https://www.alexedwards.net/blog/configuring-sqldb
	db.db.SetMaxOpenConns(25)
	db.db.SetMaxIdleConns(25)
	db.db.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return db.db.PingContext(ctx)
}
