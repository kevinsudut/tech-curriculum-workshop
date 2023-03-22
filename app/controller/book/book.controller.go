package bookcontroller

import (
	"context"

	"github.com/kevinsudut/tech-curriculum-workshops/lib/errors"
)

func (bc *BookController) GetAllBook(ctx context.Context, request *GetAllBookRequest) (response *GetAllBookResponse, err error) {
	books, err := bc.bookRepository.GetAllBook(ctx)
	if err != nil {
		return response, errors.AddTrace(err)
	}

	return &GetAllBookResponse{
		Books: books,
	}, nil
}

func (bc *BookController) GetBookByID(ctx context.Context, request *GetBookByIDRequest) (response *GetBookByIDResponse, err error) {
	err = request.Validate()
	if err != nil {
		return response, errors.AddTrace(err)
	}

	book, err := bc.bookRepository.GetBookByID(ctx, request.ID)
	if err != nil {
		return response, errors.AddTrace(err)
	}

	return &GetBookByIDResponse{
		Book: book,
	}, nil
}

func (bc *BookController) SearchBook(ctx context.Context, request *SearchBookRequest) (response *SearchBookResponse, err error) {
	err = request.Validate()
	if err != nil {
		return response, errors.AddTrace(err)
	}

	books, err := bc.bookRepository.GetBookByTitle(ctx, request.Query)
	if err != nil {
		return response, errors.AddTrace(err)
	}

	return &SearchBookResponse{
		Books: books,
	}, nil
}
