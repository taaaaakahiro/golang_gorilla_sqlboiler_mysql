package repository

import (
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/domain/entity"
)

type IUserRepository interface {
	ListUsers() ([]entity.User, error)
	User(userId int) (entity.User, error)
}
