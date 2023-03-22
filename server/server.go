package server

import (
	"context"

	"github.com/kevinsudut/tech-curriculum-workshops/app/controller"
	bookcontroller "github.com/kevinsudut/tech-curriculum-workshops/app/controller/book"
	reviewcontroller "github.com/kevinsudut/tech-curriculum-workshops/app/controller/review"
	usercontroller "github.com/kevinsudut/tech-curriculum-workshops/app/controller/user"
	"github.com/kevinsudut/tech-curriculum-workshops/app/repository"
	"github.com/kevinsudut/tech-curriculum-workshops/lib/database"
	"github.com/kevinsudut/tech-curriculum-workshops/lib/errors"
)

type Server struct {
	UserController   UserControllerItf
	BookController   BookControllerItf
	ReviewController ReviewControllerItf
}

type UserControllerItf interface {
	Register(ctx context.Context, request *usercontroller.RegisterRequest) (response *usercontroller.RegisterResponse, err error)
	Login(ctx context.Context, request *usercontroller.LoginRequest) (response *usercontroller.LoginResponse, err error)
}

type BookControllerItf interface {
	GetAllBook(ctx context.Context, request *bookcontroller.GetAllBookRequest) (response *bookcontroller.GetAllBookResponse, err error)
	GetBookByID(ctx context.Context, request *bookcontroller.GetBookByIDRequest) (response *bookcontroller.GetBookByIDResponse, err error)
	SearchBook(ctx context.Context, request *bookcontroller.SearchBookRequest) (response *bookcontroller.SearchBookResponse, err error)
}

type ReviewControllerItf interface {
	GetReviewByBook(ctx context.Context, request *reviewcontroller.GetReviewByBookRequest) (response *reviewcontroller.GetReviewByBookResponse, err error)
	PostReview(ctx context.Context, request *reviewcontroller.PostReviewRequest) (response *reviewcontroller.PostReviewResponse, err error)
	DeleteReview(ctx context.Context, request *reviewcontroller.DeleteReviewRequest) (response *reviewcontroller.DeleteReviewResponse, err error)
}

func Init() (*Server, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, errors.AddTrace(err)
	}

	controller := controller.Init(
		repository.Init(
			db,
		),
	)

	return &Server{
		UserController:   controller.UserController,
		BookController:   controller.BookController,
		ReviewController: controller.ReviewController,
	}, nil
}
