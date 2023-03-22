package userrespository

import (
	"github.com/jmoiron/sqlx"
)

func Init(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
