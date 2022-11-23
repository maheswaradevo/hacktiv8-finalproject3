package service

import (
	"context"
	"log"

	"github.com/go-playground/validator"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/task"
	"github.com/maheswaradevo/hacktiv8-finalproject3/pkg/errors"
)

type TaskServiceImpl struct {
	repo task.TaskRepository
}

func ProvideTaskService(repo task.TaskRepository) *TaskServiceImpl {
	return &TaskServiceImpl{
		repo: repo,
	}
}

func (tsk *TaskServiceImpl) CreateTask(ctx context.Context, data *dto.CreateTaskRequest, userID uint64) (res *dto.CreateTaskResponse, err error) {
	taskData := data.ToCommentEntity()
	taskData.UserID = userID
	validate := validator.New()
	validateError := validate.Struct(data)
	if validateError != nil {
		validateError = errors.ErrInvalidRequestBody
		log.Printf("[CreateTask] there's data that not through the validate process")
		return nil, validateError
	}
	taskID, err := tsk.repo.CreateTask(ctx, *taskData)
	if err != nil {
		log.Printf("[CreateTask] failed to store user data to database: %v", err)
		return
	}
	return dto.NewTaskCreateResponse(*taskData, userID, taskID), nil
}
