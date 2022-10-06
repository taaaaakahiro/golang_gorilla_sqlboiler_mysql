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

type UserRepository struct {
	database *io.SQLDatabase
	dbOpen   *sql.DB
}

var _ repository.IUserRepository = (*UserRepository)(nil)

func NewUserRepository(db *io.SQLDatabase) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (r *UserRepository) ListUsers(ctx context.Context) ([]*entity.User, error) {
	modelUsers, err := models.Users().All(ctx, r.dbOpen)
	if err != nil {
		return []*entity.User{}, errs.WithStack(err)
	}
	var users []*entity.User
	for _, modelUser := range modelUsers {
		users = append(users, &entity.User{ID: int(modelUser.ID), Name: modelUser.Name})
	}

	return users, nil
}

func (r *UserRepository) GetUser(userID string) (entity.User, error) {
	user := entity.User{}

	query := "SELECT id, name FROM users WHERE id = ?"
	stmtOut, err := r.database.Prepare(query)
	if err != nil {
		return user, err
	}
	defer stmtOut.Close()

	err = stmtOut.QueryRow(userID).Scan(&user.ID, &user.Name)
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
