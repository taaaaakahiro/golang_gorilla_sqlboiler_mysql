package persistence

import (
	"database/sql"
	"log"

	errs "github.com/pkg/errors"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/domain/entity"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/domain/repository"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/io"
)

type UserRepo struct {
	database *io.SQLDatabase
}

var _ repository.IUserRepository = (*UserRepo)(nil)

func NewUserRepository(db *io.SQLDatabase) *UserRepo {
	return &UserRepo{
		database: db,
	}
}

func (r UserRepo) ListUsers() ([]entity.User, error) {
	query := "SELECT id, name FROM users ORDER BY id DESC"
	stmtOut, err := r.database.Prepare(query)
	if err != nil {
		return nil, errs.WithStack(err)
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query()
	if err != nil {
		return nil, errs.WithStack(err)
	}

	users := make([]entity.User, 0)
	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		user := entity.User{}

		err = rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil, errs.WithStack(err)
		}

		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r UserRepo) User(userId int) (entity.User, error) {
	user := entity.User{}

	query := "SELECT id, name FROM user WHERE id = ?"
	stmtOut, err := r.database.Prepare(query)
	if err != nil {
		return user, err
	}
	defer stmtOut.Close()

	err = stmtOut.QueryRow(userId).Scan(&user.Id, &user.Name)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			log.Println("row not found")
			return user, err
		default:
			return user, err
		}
	}
	return user, nil
}
