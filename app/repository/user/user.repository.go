package userrespository

import (
	"context"

	usermodel "github.com/kevinsudut/tech-curriculum-workshops/app/model/user"
	"github.com/kevinsudut/tech-curriculum-workshops/lib/errors"
)

/* VULNERABILITY:
1. using %s instead of parameterized queries
*/
func (repo *UserRepository) InsertUser(ctx context.Context, user usermodel.User) (result usermodel.User, err error) {
	err = repo.db.QueryRowContext(ctx, queryInsertUser, user.Name, user.Email, user.Password, user.Role).Scan(&user.ID)
	if err != nil {
		return result, errors.AddTrace(err)
	}

	return user, nil
}

func (repo *UserRepository) GetUserByID(ctx context.Context, id int64) (result usermodel.User, err error) {
	err = repo.db.GetContext(ctx, &result, queryGetUserByID, id)
	if err != nil {
		return result, errors.AddTrace(err)
	}

	return result, nil
}

func (repo *UserRepository) GetUserByEmail(ctx context.Context, email string) (result usermodel.User, err error) {
	err = repo.db.GetContext(ctx, &result, queryGetUserByEmail, email)
	if err != nil {
		return result, errors.AddTrace(err)
	}

	return result, nil
}
