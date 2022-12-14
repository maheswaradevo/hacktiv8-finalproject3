package dto

import "github.com/maheswaradevo/hacktiv8-finalproject3/internal/model"

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	CategoryID  uint64 `json:"category_id"`
}

func (dto *CreateTaskRequest) ToTaskEntity() (cmt *model.Task) {
	cmt = &model.Task{
		Title:       dto.Title,
		Description: dto.Description,
		CategoryID:  dto.CategoryID,
	}
	return
}

type EditTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (dto *EditTaskRequest) ToTaskEntity() *model.TaskUserJoined {
	return &model.TaskUserJoined{
		Task: model.Task{
			Title:       dto.Title,
			Description: dto.Description,
		},
	}
}

type EditTaskStatusRequest struct {
	Status bool `json:"status"`
}

func (dto *EditTaskStatusRequest) ToTaskEntity() *model.TaskUserJoined {
	return &model.TaskUserJoined{
		Task: model.Task{
			Status: dto.Status,
		},
	}
}

type EditTaskCategoryRequest struct {
	CategoryID uint64 `json:"category_id"`
}

func (dto *EditTaskCategoryRequest) ToTaskEntity() *model.TaskUserJoined {
	return &model.TaskUserJoined{
		Task: model.Task{
			CategoryID: dto.CategoryID,
		},
	}
}
