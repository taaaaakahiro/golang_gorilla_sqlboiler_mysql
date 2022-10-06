package persistence

import (
	"database/sql"

	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/domain/repository"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/io"
)

type Repositories struct {
	db     *io.SQLDatabase
	User   repository.IUserRepository
	Review repository.IReviewRepository
}

func NewRepositories(db *io.SQLDatabase) (*Repositories, error) {
	return &Repositories{
		db:     db,
		User:   NewUserRepository(db),
		Review: NewReviewRepository(db),
	}, nil
}

func (r *Repositories) GetDatabase() *sql.DB {
	return r.db.Database
}
