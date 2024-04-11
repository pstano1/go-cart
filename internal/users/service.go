package users

import (
	"context"

	"github.com/pstano1/go-cart/protopb/users"
)

type UsersService struct {
	users.UnimplementedUsersServiceServer
}

func NewService() UsersService {
	return UsersService{}
}

func (u *UsersService) GetUser(_ context.Context, req *users.GetUserRequest) (*users.GetUserResponse, error) {
	return &users.GetUserResponse{}, nil
}

func (u *UsersService) RegisterUser(_ context.Context, req *users.RegisterUserRequest) (*users.RegisterUserResponse, error) {
	return &users.RegisterUserResponse{}, nil
}
