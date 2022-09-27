package repository

import (
	"context"

	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/domain/entity"
)

// use case entity
type IUserRepository interface {
	ListUsers(ctx context.Context) ([]*entity.User, error)
	User(userId int) (entity.User, error)
}
