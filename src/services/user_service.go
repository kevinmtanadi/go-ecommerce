package services

import (
	"context"
	"errors"
	"go-backend-template/src/constants"
	"go-backend-template/src/dtos"
	"go-backend-template/src/libraries"
	"go-backend-template/src/model"

	"github.com/sarulabs/di"
)

type UserService interface {
	Register(ctx context.Context, params dtos.RegisterReq) (data dtos.RegisterRes, err error)
	Verify(ctx context.Context, params dtos.VerifyUserReq) (data dtos.VerifyUserRes, err error)
}

type UserServiceImpl struct {
	service *BaseService
	library *libraries.Library
}

func NewUserService(ioc di.Container) UserService {
	return &UserServiceImpl{
		service: NewBaseService(ioc),
		library: ioc.Get(constants.LIBRARY).(*libraries.Library),
	}
}

func (s *UserServiceImpl) Register(ctx context.Context, params dtos.RegisterReq) (data dtos.RegisterRes, err error) {
	var response dtos.RegisterRes

	// Check if user with same username or email exists
	user, err := s.library.User.GetUserDataByUsername(ctx, params.Username)
	if user.ID > 0 {
		return response, errors.New("username already exists")
	}

	user, err = s.library.User.GetUserDataByEmail(ctx, params.Email)
	if user.ID > 0 {
		return response, errors.New("email already exists")
	}

	hashedPassword, salt, err := s.library.User.EncryptPassword(params.Password)
	if err != nil {
		return response, err
	}

	newUser := model.User{
		Username:   params.Username,
		Email:      params.Email,
		Password:   hashedPassword,
		Salt:       salt,
		IsVerified: 0,
	}

	err = s.library.User.Create(ctx, newUser)
	return response, err
}

func (s *UserServiceImpl) Verify(ctx context.Context, params dtos.VerifyUserReq) (data dtos.VerifyUserRes, err error) {
	var response dtos.VerifyUserRes

	err = s.library.User.Update(ctx, model.User{
		ID:         uint(params.UserID),
		IsVerified: 1,
	})

	return response, err
}
