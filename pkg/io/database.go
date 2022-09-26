package io

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	errs "github.com/pkg/errors"
)

type MySQLSettings interface {
	DSN() string
	MaxOpenConns() int
	MaxIdleConns() int
	ConnsMaxLifetime() int
}

type SQLDatabase struct {
	database *sql.DB
}

func NewDatabase(setting MySQLSettings) (*SQLDatabase, *sql.DB, error) {
	db, err := sql.Open("mysql", setting.DSN())
	if err != nil {
		return nil, nil, errs.WithStack(err)
	}

	// check config
	if setting.MaxOpenConns() <= 0 {
		return nil, nil, errs.WithStack(errs.New("require set max open conns"))
	}
	if setting.MaxIdleConns() <= 0 {
		return nil, nil, errs.WithStack(errs.New("require set max idle conns"))
	}
	if setting.ConnsMaxLifetime() <= 0 {
		return nil, nil, errs.WithStack(errs.New("require set conns max lifetime"))
	}
	db.SetMaxOpenConns(setting.MaxOpenConns())
	db.SetMaxIdleConns(setting.MaxIdleConns())
	db.SetConnMaxLifetime(time.Duration(setting.ConnsMaxLifetime()) * time.Second)

	return &SQLDatabase{database: db}, db, nil
}

func (d *SQLDatabase) Begin() (*sql.Tx, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	tx, err := d.database.BeginTx(ctx, &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  false,
	})
	return tx, cancel, err
}

func (d *SQLDatabase) Close() error {
	return d.database.Close()
}

func (d *SQLDatabase) Prepare(query string) (*sql.Stmt, error) {
	if d.database == nil {
		return nil, errDoesNotDB()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt, err := d.database.PrepareContext(ctx, query)

	return stmt, err
}

func (d *SQLDatabase) Exec(query string, args ...interface{}) (sql.Result, error) {
	if d.database == nil {
		return nil, errDoesNotDB()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := d.database.ExecContext(ctx, query, args)

	return res, err
}

func errDoesNotDB() error {
	return errs.New("database does not exist. Please Open() first")
}
