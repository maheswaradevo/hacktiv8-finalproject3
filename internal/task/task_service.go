package task

import (
	"context"

	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/dto"
)

type TaskService interface {
	CreateTask(ctx context.Context, data *dto.CreateTaskRequest, userID uint64) (res *dto.CreateTaskResponse, err error)
	ViewTask(ctx context.Context) (dto.ViewTasksResponse, error)
	UpdateTaskStatus(ctx context.Context, taskID uint64, userID uint64, data *dto.EditTaskStatusRequest) (*dto.EditTaskStatusResponse, error)
	UpdateTaskCategory(ctx context.Context, taskID uint64, userID uint64, data *dto.EditTaskCategoryRequest) (*dto.EditTaskStatusResponse, error)
	DeleteTask(ctx context.Context, taskID uint64, userID uint64) (*dto.DeleteTaskResponse, error)
}
