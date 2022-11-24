package dto

import "github.com/maheswaradevo/hacktiv8-finalproject3/internal/model"

type CreateTaskRequest struct {
	Title string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	CategoryID uint64 `json:"category_id"`
}

func (dto *CreateTaskRequest) ToTaskEntity() (cmt *model.Task) {
	cmt = &model.Task{
		Title: dto.Title,
		Description: dto.Description,
		CategoryID: dto.CategoryID,
	}
	return
}