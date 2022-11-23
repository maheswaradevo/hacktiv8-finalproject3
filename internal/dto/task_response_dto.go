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
