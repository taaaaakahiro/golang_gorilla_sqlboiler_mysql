package repository

import (
	"context"

	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/models"
)

// use case models created by sqlboiler
type IReviewRepository interface {
	ListReviews(ctx context.Context) ([]*models.Review, error)
}
