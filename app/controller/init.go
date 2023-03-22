package controller

import (
	bookcontroller "github.com/kevinsudut/tech-curriculum-workshops/app/controller/book"
	reviewcontroller "github.com/kevinsudut/tech-curriculum-workshops/app/controller/review"
	usercontroller "github.com/kevinsudut/tech-curriculum-workshops/app/controller/user"
	"github.com/kevinsudut/tech-curriculum-workshops/app/repository"
)

type Controller struct {
	UserController   *usercontroller.UserController
	BookController   *bookcontroller.BookController
	ReviewController *reviewcontroller.ReviewController
}

func Init(
	repository *repository.Repository,
) *Controller {
	return &Controller{
		UserController: usercontroller.Init(
			repository.UserRepository,
		),
		BookController: bookcontroller.Init(
			repository.BookRepository,
		),
		ReviewController: reviewcontroller.Init(
			repository.ReviewRepository,
			repository.BookRepository,
			repository.UserRepository,
		),
	}
}
