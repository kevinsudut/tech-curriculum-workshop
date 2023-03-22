package userrespository

import (
	"context"

	"github.com/jmoiron/sqlx"
	usermodel "github.com/kevinsudut/tech-curriculum-workshops/app/model/user"
)

type UserRepository struct {
	db *sqlx.DB
}

type UserRepositoryItf interface {
	InsertUser(ctx context.Context, user usermodel.User) (usermodel.User, error)
	GetUserByID(ctx context.Context, id int64) (result usermodel.User, err error)
	GetUserByEmail(ctx context.Context, email string) (usermodel.User, error)
}
