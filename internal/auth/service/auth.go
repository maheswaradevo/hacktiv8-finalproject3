package service

import (
	"context"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/auth"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/global/utils"
	"github.com/maheswaradevo/hacktiv8-finalproject3/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo auth.User
}

func NewAuthService(repo auth.User) *Service {
	return &Service{
		repo: repo,
	}
}

func (auth *Service) RegisterUser(ctx context.Context, data *dto.UserRegisterRequest) (*dto.UserRegisterResponse, error) {
	user := data.ToEntity()

	validate := validator.New()
	validateError := validate.Struct(data)
	if validateError != nil {
		validateError = errors.ErrInvalidRequestBody
		log.Printf("[RegisterUser] there's data that not through the validate process")
		return nil, validateError
	}

	isValid := utils.IsValidEmail(user.Email)

	if !isValid {
		err := errors.ErrInvalidRequestBody
		log.Printf("[RegisterUser] wrong email format, err: %v", err)
		return nil, err
	}
	exists, err := auth.repo.CheckEmail(ctx, user.Email)
	if err != nil && err != errors.ErrInvalidResources {
		log.Printf("[RegisterUser] failed to check duplicate email: %v", err)
		return nil, err
	}

	if exists != nil {
		err = errors.ErrUserExists
		log.Printf("[RegisterUser] user with email %v already existed", data.Email)
		return nil, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[RegisterUser] failed to created hashed password, err: %v", err)
		return nil, err
	}
	user.Password = string(hashed)

	userID, err := auth.repo.Save(ctx, *user)
	if err != nil {
		log.Printf("[RegisterUser] failed to save user, err: %v", err)
		return nil, err
	}
	user.UserID = userID
	return dto.NewUserRegisterResponse(*user), nil
}
