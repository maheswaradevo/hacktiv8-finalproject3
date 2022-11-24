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
	taskData := data.ToTaskEntity()
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

func (tsk *TaskServiceImpl) ViewTask(ctx context.Context) (dto.ViewTasksResponse, error) {
	count, err := tsk.repo.CountTask(ctx)

	if err != nil {
		log.Printf("[ViewTask] failed to count the task, err: %v", err)
		return nil, err
	}
	if count == 0 {
		err = errors.ErrDataNotFound
		log.Printf("[ViewTask] no data exists in the database: %v", err)
		return nil, err
	}
	res, err := tsk.repo.ViewTask(ctx)
	if err != nil {
		log.Printf("[ViewTask] failed to view the task, err: %v", err)
		return nil, err
	}
	return dto.NewViewTasksResponse(res), nil
}

func (tsk *TaskServiceImpl) UpdateTaskStatus(ctx context.Context, taskID uint64, userID uint64, data *dto.EditTaskStatusRequest) (*dto.EditTaskStatusResponse, error) {
	editedTaskStatus := data.ToTaskEntity()

	check, err := tsk.repo.CheckTask(ctx, taskID, userID)
	if err != nil {
		log.Printf("[UpdateTaskStatus] failed to check task with, userID: %v, err: %v", userID, err)
		return nil, err
	}
	if !check {
		err = errors.ErrDataNotFound
		log.Printf("[UpdateTaskStatus] no task in userID: %v", userID)
		return nil, err
	}
	err = tsk.repo.UpdateTaskStatus(ctx, *editedTaskStatus, taskID, userID)
	if err != nil {
		log.Printf("[UpdateTaskStatus] failed to update task status, err: %v", err)
		return nil, err
	}
	task, err := tsk.repo.GetTaskByID(ctx, taskID, userID)
	if err != nil {
		log.Printf("[UpdateTaskStatus] failed to get task, err: %v", err)
		return nil, err
	}
	return task, nil
}

func (tsk *TaskServiceImpl) UpdateTaskCategory(ctx context.Context, taskID uint64, userID uint64, data *dto.EditTaskCategoryRequest) (*dto.EditTaskStatusResponse, error) {
	editedTaskCategory := data.ToTaskEntity()

	check, err := tsk.repo.CheckTask(ctx, taskID, userID)
	if err != nil {
		log.Printf("[UpdateTaskStatus] failed to check task with, userID: %v, err: %v", userID, err)
		return nil, err
	}
	if !check {
		err = errors.ErrDataNotFound
		log.Printf("[UpdateTaskStatus] no task in userID: %v", userID)
		return nil, err
	}
	err = tsk.repo.UpdateTaskCategory(ctx, *editedTaskCategory, taskID, userID)
	if err != nil {
		log.Printf("[UpdateTaskStatus] failed to update task status, err: %v", err)
		return nil, err
	}
	task, err := tsk.repo.GetTaskByID(ctx, taskID, userID)
	if err != nil {
		log.Printf("[UpdateTaskStatus] failed to get task, err: %v", err)
		return nil, err
	}
	return task, nil
}

func (tsk *TaskServiceImpl) DeleteTask(ctx context.Context, taskID uint64, userID uint64) (*dto.DeleteTaskResponse, error) {
	check, err := tsk.repo.CheckTask(ctx, taskID, userID)
	if err != nil {
		log.Printf("[DeleteTask] failed to check task with, userID: %v, err: %v", userID, err)
		return nil, err
	}
	if !check {
		err = errors.ErrDataNotFound
		log.Printf("[DeleteTask] no task in userID: %v", userID)
		return nil, err
	}

	err = tsk.repo.DeleteTask(ctx, taskID, userID)
	if err != nil {
		log.Printf("[DeleteTask] failed to delete task, id: %v", taskID)
		return nil, err
	}
	message := "Your task has been successfully deleted"
	return dto.NewDeleteTaskResponse(message), nil
}
