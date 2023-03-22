package reviewcontroller

import (
	"context"

	"github.com/kevinsudut/tech-curriculum-workshops/lib/errors"
)

func (r GetReviewByBookRequest) Validate(ctx context.Context, bookRepository BookRepositoryItf) (err error) {
	if r.BookID <= 0 {
		return errors.AddTrace(errors.MissingParam)
	}

	_, err = bookRepository.GetBookByID(ctx, r.BookID)
	if err != nil {
		return errors.AddTrace(err)
	}

	return nil
}

func (r PostReviewRequest) Validate(ctx context.Context, bookRepository BookRepositoryItf) (err error) {
	if r.BookID <= 0 || r.Content == "" || r.Rate <= 0 {
		return errors.AddTrace(errors.MissingParam)
	}

	if len(r.Content) < 10 || len(r.Content) > 200 || r.Rate > 10 {
		return errors.AddTrace(errors.InvalidParam)
	}

	_, err = bookRepository.GetBookByID(ctx, r.BookID)
	if err != nil {
		return errors.AddTrace(err)
	}

	return nil
}

func (r DeleteReviewRequest) Validate(ctx context.Context, reviewRepository ReviewRepositoryItf) (err error) {
	if r.ReviewID <= 0 {
		return errors.AddTrace(errors.MissingParam)
	}

	_, err = reviewRepository.GetReviewByID(ctx, r.ReviewID)
	if err != nil {
		return errors.AddTrace(err)
	}

	return nil
}
