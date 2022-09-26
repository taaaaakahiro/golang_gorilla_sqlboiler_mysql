package handler

import (
	"context"

	v1 "github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/handler/v1"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/persistence"
	"go.uber.org/zap"
)

type Handler struct {
	V1 *v1.Handler
	// Version *version.Handler
}

func NewHandler(ctx context.Context, logger *zap.Logger, repo *persistence.Repositories, ver string) *Handler {
	h := &Handler{
		V1: v1.NewHandler(ctx, logger.Named("v1"), repo),
		// Version: version.NewHandler(logger.Named("version"), ver),
	}

	return h
}
