package bookrepository

import (
	"github.com/jmoiron/sqlx"
)

func Init(db *sqlx.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}
