package bookrepository

import (
	"context"

	bookmodel "github.com/kevinsudut/tech-curriculum-workshops/app/model/book"
	"github.com/kevinsudut/tech-curriculum-workshops/lib/errors"
)

func (repo *BookRepository) GetAllBook(ctx context.Context) (result []bookmodel.Book, err error) {
	err = repo.db.SelectContext(ctx, &result, queryGetAllBook)
	if err != nil {
		return result, errors.AddTrace(err)
	}

	return result, nil
}

func (repo *BookRepository) GetBookByID(ctx context.Context, id int64) (result bookmodel.Book, err error) {
	err = repo.db.GetContext(ctx, &result, queryGetBookByID, id)
	if err != nil {
		return result, errors.AddTrace(err)
	}

	return result, nil
}

func (repo *BookRepository) GetBookByTitle(ctx context.Context, title string) (result []bookmodel.Book, err error) {
	err = repo.db.SelectContext(ctx, &result, queryGetBookByTitle, title)
	if err != nil {
		// Don't log queries..
		// fmt.Println(queryGetBookByTitle, title)
		return result, errors.AddTrace(err)
	}

	return result, nil
}
