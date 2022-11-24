package dto

import (
	"time"

	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/model"
)

type CreateTaskResponse struct {
	TaskID      uint64    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	CategoryID  uint64    `json:"category_id"`
	UserID      uint64    `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type ViewTaskResponse struct {
	TaskID      uint64               `json:"id"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	Status      bool                 `json:"status"`
	CategoryID  uint64               `json:"category_id"`
	UserID      uint64               `json:"user_id"`
	UpdatedAt   time.Time            `json:"updated_at"`
	CreatedAt   time.Time            `json:"created_at"`
	User        ViewTaskUserResponse `json:"user"`
}

type ViewTaskUserResponse struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type ViewTasksResponse []*ViewTaskResponse

type DeleteTaskResponse struct {
	Message string `json:"message"`
}

func NewTaskCreateResponse(tsk model.Task, userID uint64, taskID uint64) *CreateTaskResponse {
	return &CreateTaskResponse{
		TaskID:      taskID,
		Title:       tsk.Title,
		Description: tsk.Description,
		CategoryID:  tsk.CategoryID,
		UserID:      userID,
		CreatedAt:   time.Now(),
	}
}

func NewViewTaskResponse(tsk model.TaskUserJoined) *ViewTaskResponse {
	return &ViewTaskResponse{
		TaskID:      tsk.Task.TaskID,
		Title:       tsk.Task.Title,
		Description: tsk.Task.Description,
		Status:      tsk.Task.Status,
		CategoryID:  tsk.Task.CategoryID,
		UserID:      tsk.Task.UserID,
		UpdatedAt:   tsk.Task.UpdatedAt,
		CreatedAt:   tsk.Task.CreatedAt,
		User: ViewTaskUserResponse{
			Email:    tsk.User.Email,
			FullName: tsk.User.FullName,
		},
	}
}

func NewViewTasksResponse(tsk model.PeopleTaskJoined) ViewTasksResponse {
	var viewTasksResponse ViewTasksResponse

	for idx := range tsk {
		peopleTask := NewViewTaskResponse(*tsk[idx])
		viewTasksResponse = append(viewTasksResponse, peopleTask)
	}
	return viewTasksResponse
}

func NewDeleteTaskResponse(message string) *DeleteTaskResponse {
	return &DeleteTaskResponse{
		Message: message,
	}
}
