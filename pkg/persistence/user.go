package persistence

import (
	"context"
	"database/sql"
	"log"

	errs "github.com/pkg/errors"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/domain/entity"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/domain/repository"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/io"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/models"
)

type UserRepo struct {
	database *io.SQLDatabase
	dbOpen   *sql.DB
}

var _ repository.IUserRepository = (*UserRepo)(nil)

func NewUserRepository(db *io.SQLDatabase, dbOpen *sql.DB) *UserRepo {
	return &UserRepo{
		database: db,
		dbOpen:   dbOpen,
	}
}

func (r UserRepo) ListUsers(ctx context.Context) (*[]*entity.User, error) {
	modelUsers, err := models.Users().All(ctx, r.dbOpen)
	if err != nil {
		return nil, errs.WithStack(err)
	}
	var users []*entity.User
	for _, modelUser := range modelUsers {
		users = append(users, &entity.User{Id: int(modelUser.ID), Name: modelUser.Name})
	}

	return &users, nil
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
