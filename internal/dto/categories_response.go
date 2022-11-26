package dto

import (
	"time"

	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/model"
)

type CreateCategoriesResponse struct {
	CategoryID uint64    `json:"id"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewCategoriesCreateResponse(ctg model.Categories, role string, categoryID uint64) *CreateCategoriesResponse {
	return &CreateCategoriesResponse{
		CategoryID: categoryID,
		Type:       ctg.Type,
		CreatedAt:  time.Now(),
	}
}

type ViewCategoriesResponse struct {
	CategoryID uint64                     `json:"id"`
	Type       string                     `json:"type"`
	UpdatedAt  time.Time                  `json:"updated_at"`
	CreatedAt  time.Time                  `json:"created_at"`
	Task       ViewCategoriesTaskResponse `json:"task"`
}

type ViewCategoriesTaskResponse struct {
	TaskID      uint64    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      uint64    `json:"user_id"`
	CategoryID  uint64    `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ViewAllCategoriesResponse []*ViewCategoriesResponse

func NewViewCategoriesResponse(ctg model.CategoriesUserJoined) *ViewCategoriesResponse {
	return &ViewCategoriesResponse{
		CategoryID: ctg.Categories.CategoryID,
		Type:       ctg.Categories.Type,
		CreatedAt:  ctg.Categories.CreatedAt,
		UpdatedAt:  ctg.Categories.UpdatedAt,
		Task: ViewCategoriesTaskResponse{
			TaskID:      ctg.Task.TaskID,
			Title:       ctg.Task.Title,
			Description: ctg.Task.Description,
			UserID:      ctg.Task.UserID,
			CategoryID:  ctg.Task.CategoryID,
			CreatedAt:   ctg.Task.CreatedAt,
			UpdatedAt:   ctg.Task.UpdatedAt,
		},
	}
}

func NewViewAllCategoriesResponse(ctg model.PeopleCategoriesJoined) *ViewAllCategoriesResponse {
	var viewAllCategoriesResponse ViewAllCategoriesResponse

	for idx := range ctg {
		peopleCategories := NewViewCategoriesResponse(*ctg[idx])
		viewAllCategoriesResponse = append(viewAllCategoriesResponse, peopleCategories)
	}
	return &viewAllCategoriesResponse
}

type EditCategoriesResponse struct {
	CategoryID uint64    `json:"id"`
	Type       string    `json:"type"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewEditCategoriesResponse(ctg model.Categories) *EditCategoriesResponse {
	return &EditCategoriesResponse{
		CategoryID: ctg.CategoryID,
		Type:       ctg.Type,
		UpdatedAt:  ctg.UpdatedAt,
	}
}

type DeleteCategoriesResponse struct {
	Message string `json:"message"`
}

func NewDeleteCategoriesResponse(message string) *DeleteCategoriesResponse {
	return &DeleteCategoriesResponse{
		Message: message,
	}
}
