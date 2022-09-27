package v1

import (
	"context"

	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/persistence"
	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.Logger
	repo   *persistence.Repositories
	ctx    context.Context
}

func NewHandler(ctx context.Context, logger *zap.Logger, repositories *persistence.Repositories) *Handler {
	return &Handler{
		logger: logger,
		repo:   repositories,
		ctx:    ctx,
	}
}
