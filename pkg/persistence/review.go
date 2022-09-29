package persistence

import (
	"context"
	"database/sql"

	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/domain/repository"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/io"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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
	reviews, err := models.Reviews(qm.Select("id", "text")).All(ctx, r.dbOpen)
	if err != nil {
		return []*models.Review{}, err
	}
	return reviews, nil
}

func (r *ReviewRepository) GetReview(ctx context.Context, reviewID string) (*models.Review, error) {
	// CASE1
	// id, _ := strconv.ParseInt(reviewID, 10, 64)
	// reviews, err := models.FindReview(ctx, r.dbOpen, id)

	//CASE2
	reviews, err := models.Reviews(qm.Where("id = ?", reviewID)).One(ctx, r.dbOpen)
	if err != nil {
		return &models.Review{}, err
	}
	return reviews, nil
}
