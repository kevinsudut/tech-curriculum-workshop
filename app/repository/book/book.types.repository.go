package bookrepository

import (
	"context"

	"github.com/jmoiron/sqlx"
	bookmodel "github.com/kevinsudut/tech-curriculum-workshops/app/model/book"
)

type BookRepository struct {
	db *sqlx.DB
}

type BookRepositoryItf interface {
	GetAllBook(ctx context.Context) ([]bookmodel.Book, error)
	GetBookByID(ctx context.Context, id int64) (result bookmodel.Book, err error)
	GetBookByTitle(ctx context.Context, title string) (result []bookmodel.Book, err error)
}
