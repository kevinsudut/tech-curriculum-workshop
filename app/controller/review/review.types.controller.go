package reviewcontroller

import (
	"context"
	"time"

	bookmodel "github.com/kevinsudut/tech-curriculum-workshops/app/model/book"
	reviewmodel "github.com/kevinsudut/tech-curriculum-workshops/app/model/review"
	usermodel "github.com/kevinsudut/tech-curriculum-workshops/app/model/user"
)

type ReviewController struct {
	reviewRepository ReviewRepositoryItf
	bookRepository   BookRepositoryItf
	userRepository   UserRepositoryItf
}

type ReviewRepositoryItf interface {
	InsertReview(ctx context.Context, review reviewmodel.Review) (reviewmodel.Review, error)
	UpdateReviewStatus(ctx context.Context, reviewID int64, status int) error
	GetReviewByID(ctx context.Context, id int64) (reviewmodel.Review, error)
	GetActiveReviewByBook(ctx context.Context, bookID int64) ([]reviewmodel.Review, error)
}

type BookRepositoryItf interface {
	GetBookByID(ctx context.Context, id int64) (result bookmodel.Book, err error)
}

type UserRepositoryItf interface {
	GetUserByID(ctx context.Context, id int64) (result usermodel.User, err error)
}

type GetReviewByBookRequest struct {
	BookID int64 `json:"book_id"`
}

type GetReviewByBookResponse struct {
	Reviews []ReviewResponse `json:"reviews"`
}

type ReviewResponse struct {
	ID          int64          `json:"id"`
	User        usermodel.User `json:"user"`
	Book        bookmodel.Book `json:"book"`
	Rate        int            `json:"rate"`
	Content     string         `json:"content"`
	Status      int            `json:"status"`
	CreatedTime time.Time      `json:"created_time"`
}

type PostReviewRequest struct {
	BookID  int64  `json:"book_id"`
	Rate    int    `json:"rate"`
	Content string `json:"content"`
}

type PostReviewResponse struct {
	Message string             `json:"message"`
	Review  reviewmodel.Review `json:"review"`
}

type DeleteReviewRequest struct {
	ReviewID int64 `json:"review_id"`
}

type DeleteReviewResponse struct {
	Message string `json:"message"`
}
