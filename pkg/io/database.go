package io

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Database *sql.DB
}

func NewDatabase(Dsn string) (*DBConfig, error) {
	conn, err := sql.Open("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		return nil, err
	}
	db := &DBConfig{
		Database: conn,
	}
	return db, nil
}
