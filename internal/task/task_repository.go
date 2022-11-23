package task

import (
	"context"

	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/model"
)


type TaskRepository interface {
	CreateTask(ctx context.Context, data model.Task) (taskID uint64, err error)
	CheckTask(ctx context.Context, categoryID uint64) (bool, error)
}