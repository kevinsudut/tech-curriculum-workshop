package usercontroller

import (
	"context"
	"net/mail"

	usermodel "github.com/kevinsudut/tech-curriculum-workshops/app/model/user"
	"github.com/kevinsudut/tech-curriculum-workshops/lib/errors"
	"golang.org/x/crypto/bcrypt"
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

	// VULNERABILITY: do not use MD5
	// passwordHash := md5.New()
	// io.WriteString(passwordHash, r.Password)

	bytes, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		return result, errors.AddTrace(err)
	}

	// if result.Password != fmt.Sprintf("%x", passwordHash.Sum(nil)) {
	if err := bcrypt.CompareHashAndPassword(bytes, []byte(result.Password)); err != nil {
		return result, errors.AddTrace(errors.InvalidUserPassword)
	}

	return result, nil
}
