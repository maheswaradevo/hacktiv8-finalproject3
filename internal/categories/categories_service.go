package categories

import (
	"context"

	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/dto"
)

type CategoriesService interface {
	CreateCategories(ctx context.Context, data *dto.CreateCategoriesRequest, role string) (res *dto.CreateCategoriesResponse, err error)
	ViewCategories(ctx context.Context) (*dto.ViewAllCategoriesResponse, error)
	UpdateCategories(ctx context.Context, categoryID uint64, role string, data *dto.EditCategoriesRequest) (*dto.EditCategoriesResponse, error)
	DeleteCategories(ctx context.Context, categoryID uint64, role string) (*dto.DeleteCategoriesResponse, error)
}
