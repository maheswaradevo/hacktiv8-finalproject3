package auth

import (
	"context"

	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/dto"
)

type UserService interface {
	RegisterUser(ctx context.Context, data *dto.UserRegisterRequest) (*dto.UserRegisterResponse, error)
	Login(ctx context.Context, data *dto.UserSignInRequest) (*dto.UserSignInResponse, error)
}
