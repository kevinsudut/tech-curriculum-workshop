package usercontroller

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"net/mail"

	usermodel "github.com/kevinsudut/tech-curriculum-workshops/app/model/user"
	"github.com/kevinsudut/tech-curriculum-workshops/lib/errors"
)

func (r RegisterRequest) Validate(ctx context.Context, userRepository UserRepositoryItf) error {
	if r.Name == "" || r.Email == "" || r.Password == "" {
		return errors.AddTrace(errors.MissingParam)
	}

	if len(r.Name) < 3 || len(r.Email) < 5 || len(r.Password) < 8 {
		return errors.AddTrace(errors.InvalidParam)
	}

	if _, err := mail.ParseAddress(r.Email); err != nil {
		return errors.AddTrace(errors.InvalidParam)
	}

	user, err := userRepository.GetUserByEmail(ctx, r.Email)
	if err != nil && !errors.IsNotFound(err) {
		return errors.AddTrace(err)
	}

	if user.ID > 0 {
		return errors.AddTrace(errors.AlreadyRegistered)
	}

	return nil
}

func (r LoginRequest) Validate(ctx context.Context, userRepository UserRepositoryItf) (result usermodel.User, err error) {
	if r.Email == "" || r.Password == "" {
		return result, errors.AddTrace(errors.MissingParam)
	}

	result, err = userRepository.GetUserByEmail(ctx, r.Email)
	if err != nil && !errors.IsNotFound(err) {
		return result, errors.AddTrace(err)
	}

	if result.ID <= 0 {
		return result, errors.AddTrace(errors.EmailDoesNotExists)
	}

	passwordHash := md5.New()
	io.WriteString(passwordHash, r.Password)

	if result.Password != fmt.Sprintf("%x", passwordHash.Sum(nil)) {
		return result, errors.AddTrace(errors.InvalidUserPassword)
	}

	return result, nil
}
