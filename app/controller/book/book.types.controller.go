package bookcontroller

import (
	"context"

	bookmodel "github.com/kevinsudut/tech-curriculum-workshops/app/model/book"
)

type BookController struct {
	bookRepository BookRepositoryItf
}

type BookRepositoryItf interface {
	GetAllBook(ctx context.Context) ([]bookmodel.Book, error)
	GetBookByID(ctx context.Context, id int64) (result bookmodel.Book, err error)
	GetBookByTitle(ctx context.Context, title string) (result []bookmodel.Book, err error)
}

type GetAllBookRequest struct{}

type GetAllBookResponse struct {
	Books []bookmodel.Book `json:"books"`
}

type GetBookByIDRequest struct {
	ID int64 `json:"id"`
}

type GetBookByIDResponse struct {
	Book bookmodel.Book `json:"book"`
}

type SearchBookRequest struct {
	Query string `json:"query"`
}

type SearchBookResponse struct {
	Books []bookmodel.Book `json:"books"`
}
