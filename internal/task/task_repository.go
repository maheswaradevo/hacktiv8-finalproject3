package task

import (
	"context"

	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/model"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, data model.Task) (taskID uint64, err error)
	CheckCategory(ctx context.Context, categoryID uint64) (bool, error)
	ViewTask(ctx context.Context) (model.PeopleTaskJoined, error)
	CountTask(ctx context.Context) (int, error)
	CheckTask(ctx context.Context, taskID uint64, userID uint64) (bool, error)
	UpdateTaskStatus(ctx context.Context, reqData model.TaskUserJoined, taskID uint64, userID uint64) error
	DeleteTask(ctx context.Context, taskID uint64, userID uint64) error
	GetTaskByID(ctx context.Context, taskID uint64, userID uint64) (*dto.EditTaskStatusResponse, error)
}
