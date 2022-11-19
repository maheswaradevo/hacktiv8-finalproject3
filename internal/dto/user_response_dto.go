package dto

import (
	"time"

	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/model"
)

type UserRegisterResponse struct {
	UserID    uint64    `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserSignInResponse struct {
	AccessToken string `json:"token"`
}

func NewUserRegisterResponse(u model.User) *UserRegisterResponse {
	return &UserRegisterResponse{
		UserID:    u.UserID,
		FullName:  u.FullName,
		Email:     u.Email,
		CreatedAt: time.Now(),
	}
}

func NewUserSignInResponse(ac string) *UserSignInResponse {
	return &UserSignInResponse{AccessToken: ac}
}
