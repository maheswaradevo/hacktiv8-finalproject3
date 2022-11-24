package service

import (
	"context"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/auth"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/global/config"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/global/utils"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/model"
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
		err := errors.ErrEmailFormat
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

func (auth *Service) Login(ctx context.Context, data *dto.UserSignInRequest) (*dto.UserSignInResponse, error) {
	userLogin := data.ToEntity()

	userCred, err := auth.repo.CheckEmail(ctx, userLogin.Email)
	if err != nil {
		log.Printf("[Login] failed to fetch user with email: %v, err: %v", userLogin.Email, err)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userCred.Password), []byte(userLogin.Password))
	if err != nil {
		err = errors.ErrMismatchedHashAndPassword
		log.Printf("[Login] wrong password, err: %v", err)
		return nil, err
	}
	token, err := auth.createAccessToken(userCred)
	if err != nil {
		log.Printf("[Login] failed to create new token, err: %v", err)
		return nil, err
	}
	return dto.NewUserSignInResponse(token), nil
}

func (auth *Service) UpdateAccount(ctx context.Context, userID uint64, data *dto.UserUpdateAccountRequest) (*dto.UserUpdateAccountResponse, error) {
	userUpdate := data.ToEntity()

	exists, err := auth.repo.CheckEmail(ctx, userUpdate.Email)
	if err != nil {
		log.Printf("[UpdateAccount] failed to fetch email user, %v", userUpdate.Email)
		return nil, err
	}
	if exists == nil {
		err = errors.ErrDataNotFound
		log.Printf("[UpdateAccount] data with userID %v not found", userID)
		return nil, err
	}

	err = auth.repo.UpdateAccount(ctx, *userUpdate, userID)
	if err != nil {
		log.Printf("[UpdateAccount] failed to update user")
	}
	return dto.NewUserUpdateAccount(*userUpdate, userID), nil
}

func (auth *Service) DeleteAccount(ctx context.Context, userID uint64) (*dto.UserDeleteAccountResponse, error) {
	exists, err := auth.repo.FindUserByID(ctx, userID)
	if err != nil {
		log.Printf("[DeleteAccount] failed to check user, id: %v, err: %v", userID, err)
		return nil, err
	}
	if !exists {
		err = errors.ErrNotFound
		log.Printf("[DeleteAccount] user not found, id: %v", userID)
		return nil, err
	}

	msg, err := auth.repo.DeleteAccount(ctx, userID)
	if err != nil {
		log.Printf("[DeleteAccount] failed to delete account on id %v, err : %v", userID, err)
		return nil, err
	}

	return dto.NewUserDeleteAccountResponse(msg), nil
}

func (auth *Service) createAccessToken(user *model.User) (string, error) {
	cfg := config.GetConfig()

	claim := jwt.MapClaims{
		"authorized": true,
		"exp":        time.Now().Add(time.Hour * 8).Unix(),
		"user_id":    user.UserID,
		"user_role":  user.Role,
	}

	token := jwt.NewWithClaims(cfg.JWT_SIGNING_METHOD, claim)
	signedToken, err := token.SignedString([]byte(cfg.API_SECRET_KEY))
	if err != nil {
		log.Printf("[createAccessToken] failed to create new token, err: %v", err)
		return "", err
	}
	return signedToken, nil
}
