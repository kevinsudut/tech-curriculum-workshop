package reviewrepository

import (
	"context"

	"github.com/jmoiron/sqlx"
	reviewmodel "github.com/kevinsudut/tech-curriculum-workshops/app/model/review"
)

type ReviewRepository struct {
	db *sqlx.DB
}

type ReviewRepositoryItf interface {
	InsertReview(ctx context.Context, review reviewmodel.Review) (reviewmodel.Review, error)
	UpdateReviewStatus(ctx context.Context, reviewID int64, status int) error
	GetReviewByID(ctx context.Context, id int64) (reviewmodel.Review, error)
	GetActiveReviewByBook(ctx context.Context, bookID int64) ([]reviewmodel.Review, error)
}
