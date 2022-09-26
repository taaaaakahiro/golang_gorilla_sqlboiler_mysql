package persistence

import (
	"database/sql"

	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/domain/repository"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/io"
)

type Repositories struct {
	User repository.IUserRepository
}

func NewRepositories(db *io.SQLDatabase, dbOpen *sql.DB) (*Repositories, error) {
	return &Repositories{
		User: NewUserRepository(db, dbOpen),
	}, nil
}
