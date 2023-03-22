package usercontroller

import (
	"context"

	usermodel "github.com/kevinsudut/tech-curriculum-workshops/app/model/user"
)

type UserController struct {
	userRepository UserRepositoryItf
}

type UserRepositoryItf interface {
	InsertUser(ctx context.Context, user usermodel.User) (usermodel.User, error)
	GetUserByEmail(ctx context.Context, email string) (usermodel.User, error)
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `jsoan:"password"`
}

type RegisterResponse struct {
	Message string         `json:"message"`
	User    usermodel.User `json:"user"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `jsoan:"password"`
}

type LoginResponse struct {
	Message string         `json:"message"`
	User    usermodel.User `json:"user"`
}
