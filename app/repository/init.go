package repository

import (
	"github.com/jmoiron/sqlx"
	bookrepository "github.com/kevinsudut/tech-curriculum-workshops/app/repository/book"
	reviewrepository "github.com/kevinsudut/tech-curriculum-workshops/app/repository/review"
	userrepository "github.com/kevinsudut/tech-curriculum-workshops/app/repository/user"
)

type Repository struct {
	UserRepository   userrepository.UserRepositoryItf
	BookRepository   bookrepository.BookRepositoryItf
	ReviewRepository reviewrepository.ReviewRepositoryItf
}

func Init(
	db *sqlx.DB,
) *Repository {
	return &Repository{
		UserRepository:   userrepository.Init(db),
		BookRepository:   bookrepository.Init(db),
		ReviewRepository: reviewrepository.Init(db),
	}
}
