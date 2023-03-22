package usercontroller

import (
	"context"

	usermodel "github.com/kevinsudut/tech-curriculum-workshops/app/model/user"
	"github.com/kevinsudut/tech-curriculum-workshops/lib/errors"
)

func (uc *UserController) Register(ctx context.Context, request *RegisterRequest) (response *RegisterResponse, err error) {
	err = request.Validate(ctx, uc.userRepository)
	if err != nil {
		return response, errors.AddTrace(err)
	}

	user, err := uc.userRepository.InsertUser(ctx, usermodel.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Role:     "general",
	})
	if err != nil {
		return response, errors.AddTrace(err)
	}

	return &RegisterResponse{
		Message: "Successfully register new user",
		User:    user,
	}, nil
}

func (uc *UserController) Login(ctx context.Context, request *LoginRequest) (response *LoginResponse, err error) {
	user, err := request.Validate(ctx, uc.userRepository)
	if err != nil {
		return response, errors.AddTrace(err)
	}

	return &LoginResponse{
		Message: "Successfully login",
		User:    user,
	}, nil
}
