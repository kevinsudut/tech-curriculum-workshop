package reviewrepository

import (
	"github.com/jmoiron/sqlx"
)

func Init(db *sqlx.DB) *ReviewRepository {
	return &ReviewRepository{
		db: db,
	}
}
