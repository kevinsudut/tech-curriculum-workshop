package reviewrepository

import (
	"context"

	reviewmodel "github.com/kevinsudut/tech-curriculum-workshops/app/model/review"
	"github.com/kevinsudut/tech-curriculum-workshops/lib/errors"
)

func (repo *ReviewRepository) InsertReview(ctx context.Context, review reviewmodel.Review) (result reviewmodel.Review, err error) {
	err = repo.db.QueryRowContext(ctx, queryInsertReview, review.UserID, review.BookID, review.Rate, review.Content, review.Status).Scan(&review.ID)
	if err != nil {
		return result, errors.AddTrace(err)
	}

	return review, nil
}

func (repo *ReviewRepository) UpdateReviewStatus(ctx context.Context, reviewID int64, status int) (err error) {
	result, err := repo.db.ExecContext(ctx, queryUpdateReviewStatus, status, reviewID)
	if err != nil {
		return errors.AddTrace(err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return errors.AddTrace(err)
	}

	if affected <= 0 {
		return errors.AddTrace(errors.BadRequest)
	}

	return nil
}

func (repo *ReviewRepository) GetReviewByID(ctx context.Context, id int64) (result reviewmodel.Review, err error) {
	err = repo.db.GetContext(ctx, &result, queryGetReviewByID, id)
	if err != nil {
		return result, errors.AddTrace(err)
	}

	return result, nil
}

func (repo *ReviewRepository) GetActiveReviewByBook(ctx context.Context, bookID int64) (result []reviewmodel.Review, err error) {
	err = repo.db.SelectContext(ctx, &result, queryGetReviewByBookAndStatus, bookID, reviewmodel.StatusActive)
	if err != nil {
		return result, errors.AddTrace(err)
	}

	return result, nil
}
