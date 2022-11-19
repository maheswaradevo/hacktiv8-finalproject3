package dto

import (
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/model"
)

type UserRegisterRequest struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

func (dto *UserRegisterRequest) ToEntity() (u *model.User) {
	u = &model.User{
		FullName: dto.FullName,
		Email:    dto.Email,
		Password: dto.Password,
		Role:     "Member",
	}
	return
}
