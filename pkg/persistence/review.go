package persistence

import (
	"context"
	"database/sql"

	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/domain/repository"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/io"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/models"
)

type ReviewRepository struct {
	database *io.SQLDatabase
	dbOpen   *sql.DB
}

var _ repository.IReviewRepository = (*ReviewRepository)(nil)

func NewReviewRepository(db *io.SQLDatabase, dbOpen *sql.DB) *ReviewRepository {
	return &ReviewRepository{
		database: db,
		dbOpen:   dbOpen,
	}
}
func (r *ReviewRepository) ListReviews(ctx context.Context) ([]*models.Review, error) {
	reviews, err := models.Reviews().All(ctx, r.dbOpen)
	if err != nil {
		return []*models.Review{}, err
	}
	return reviews, nil
}
