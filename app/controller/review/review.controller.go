package reviewcontroller

import (
	"context"
	"time"

	reviewmodel "github.com/kevinsudut/tech-curriculum-workshops/app/model/review"
	usermodel "github.com/kevinsudut/tech-curriculum-workshops/app/model/user"
	"github.com/kevinsudut/tech-curriculum-workshops/config"
	"github.com/kevinsudut/tech-curriculum-workshops/lib/errors"
)

func (rc *ReviewController) GetReviewByBook(ctx context.Context, request *GetReviewByBookRequest) (response *GetReviewByBookResponse, err error) {
	err = request.Validate(ctx, rc.bookRepository)
	if err != nil {
		return response, errors.AddTrace(err)
	}

	reviews, err := rc.reviewRepository.GetActiveReviewByBook(ctx, request.BookID)
	if err != nil {
		return response, errors.AddTrace(err)
	}

	response = &GetReviewByBookResponse{}

	response.Reviews = make([]ReviewResponse, 0, len(reviews))

	for _, review := range reviews {
		book, err := rc.bookRepository.GetBookByID(ctx, review.BookID)
		if err != nil {
			return response, errors.AddTrace(err)
		}

		user, err := rc.userRepository.GetUserByID(ctx, review.UserID)
		if err != nil {
			return response, errors.AddTrace(err)
		}

		response.Reviews = append(response.Reviews, ReviewResponse{
			ID:          review.ID,
			Book:        book,
			User:        user,
			Rate:        review.Rate,
			Content:     review.Content,
			Status:      review.Status,
			CreatedTime: review.CreatedTime,
		})
	}

	return response, nil
}

func (rc *ReviewController) PostReview(ctx context.Context, request *PostReviewRequest) (response *PostReviewResponse, err error) {
	err = request.Validate(ctx, rc.bookRepository)
	if err != nil {
		return response, errors.AddTrace(err)
	}

	auth, _ := ctx.Value(config.SESSION_AUTHENTICATION).(*usermodel.User)

	review, err := rc.reviewRepository.InsertReview(ctx, reviewmodel.Review{
		UserID:      auth.ID,
		BookID:      request.BookID,
		Rate:        request.Rate,
		Content:     request.Content,
		CreatedTime: time.Now(),
		Status:      reviewmodel.StatusActive,
	})
	if err != nil {
		return response, errors.AddTrace(err)
	}

	return &PostReviewResponse{
		Message: "Successfully post a review",
		Review:  review,
	}, nil
}

func (rc *ReviewController) DeleteReview(ctx context.Context, request *DeleteReviewRequest) (response *DeleteReviewResponse, err error) {
	err = request.Validate(ctx, rc.reviewRepository)
	if err != nil {
		return response, errors.AddTrace(err)
	}

	err = rc.reviewRepository.UpdateReviewStatus(ctx, request.ReviewID, reviewmodel.StatusInactive)
	if err != nil {
		return response, errors.AddTrace(err)
	}

	return &DeleteReviewResponse{
		Message: "Successfully delete a review",
	}, nil
}
