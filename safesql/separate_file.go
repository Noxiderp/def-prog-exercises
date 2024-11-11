package safesql

import (
	"context"
	"database/sql"package safesql

import (
    "context"
    "database/sql"
)

type DB struct {
    db *sql.DB
}

func (db *DB) QueryContext(ctx context.Context, query TrustedSQL, args ...any) (*Rows, error) {
    r, err := db.db.QueryContext(ctx, query.s, args...)
    return r, err
}

func (db *DB) ExecContext(ctx context.Context, query TrustedSQL, args ...any) (Result, error) {
    r, err := db.db.ExecContext(ctx, query.s, args...)
    return r, err
}

func Open(driverName, dataSourceName string) (*DB, error) {
    db, err := sql.Open(driverName, dataSourceName)
    if err != nil {
        return nil, err
    }
    return &DB{db: db}, nil
}

type Rows = sql.Rows

type Result = sql.Result
)

type DB struct {
	db *sql.DB
}

func (db *DB) QueryContext(ctx context.Context, query TrustedSQL, args ...any) (*Rows, error) {
	r, err := db.db.QueryContext(ctx, query.s, args...)
	return r, err
}

func (db *DB) ExecContext(ctx context.Context, query TrustedSQL, args ...any) (Result, error) {
	r, err := db.db.ExecContext(ctx, query.s, args...)
	return r, err
}

func Open(driverName, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}

type Rows = sql.Rows

type Result = sql.Result
