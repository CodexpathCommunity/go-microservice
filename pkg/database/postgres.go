package database

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type db struct {
	ctx  context.Context
	conn *pgxpool.Pool
}

// New creates a new postgresql database instance
func NewPostgreSQL(ctx context.Context, connString string) (SQL, error) {
	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &db{
		ctx:  ctx,
		conn: conn,
	}, nil
}

func (p *db) Exec(sql string, args ...interface{}) (int64, error) {
	result, err := p.conn.Exec(p.ctx, sql, args...)
	return result.RowsAffected(), err
}

func (p *db) Query(sql string, args ...interface{}) (Rows, error) {
	rows, err := p.conn.Query(p.ctx, sql, args...)
	return newDatabaseRows(rows), err
}

func (p *db) QueryRow(sql string, args ...interface{}) Row {
	row := p.conn.QueryRow(p.ctx, sql, args...)
	return newDatabaseRow(row)
}

func (p *db) BeginTx() (Transaction, error) {
	tx, err := p.conn.BeginTx(p.ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}

	return newDatabaseTransaction(p.ctx, tx), nil
}

type dbRow struct {
	row pgx.Row
}

func newDatabaseRow(row pgx.Row) Row {
	return &dbRow{
		row: row,
	}
}

func (p *dbRow) Scan(dest ...interface{}) error {
	return p.row.Scan(dest...)
}

type dbRows struct {
	rows pgx.Rows
}

func newDatabaseRows(rows pgx.Rows) Rows {
	return &dbRows{
		rows: rows,
	}
}

func (p *dbRows) Scan(dest ...interface{}) error {
	return p.rows.Scan(dest...)
}

func (p *dbRows) Next() bool {
	return p.rows.Next()
}

func (p *dbRows) Close() {
	p.rows.Close()
}

func (p *dbRows) Err() error {
	return p.rows.Err()
}

type transaction struct {
	ctx context.Context
	tx  pgx.Tx
}

func newDatabaseTransaction(ctx context.Context, tx pgx.Tx) *transaction {
	return &transaction{
		ctx: ctx,
		tx:  tx,
	}
}

func (p *transaction) Commit() error {
	return p.tx.Commit(p.ctx)
}

func (p *transaction) Rollback() error {
	return p.tx.Rollback(p.ctx)
}

func (p *transaction) Exec(sql string, args ...interface{}) (int64, error) {
	result, err := p.tx.Exec(p.ctx, sql, args...)
	return result.RowsAffected(), err
}

func (p *transaction) Query(sql string, args ...interface{}) (Rows, error) {
	rows, err := p.tx.Query(p.ctx, sql, args...)
	return newDatabaseRows(rows), err
}

func (p *transaction) QueryRow(sql string, args ...interface{}) Row {
	row := p.tx.QueryRow(p.ctx, sql, args...)
	return newDatabaseRow(row)
}