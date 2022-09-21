package v1

import (
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/persistence"
	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.Logger
	repo   *persistence.Repositories
}

func NewHandler(logger *zap.Logger, repositories *persistence.Repositories) *Handler {
	return &Handler{
		logger: logger,
		repo:   repositories,
	}
}
