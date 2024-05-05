package usecases

import (
	"context"
	"task_mission/apperror"
	"task_mission/entities/dtos/requests"
	"task_mission/entities/dtos/responses"
)

type ICategoryUsecase interface {
	CreateCategory(ctx context.Context, request *requests.CategoryRequest) (result *responses.CategoryResponse, customErr *apperror.CustomError)
	GetAllCategories(ctx context.Context) (result []*responses.CategoryResponse, customErr *apperror.CustomError)
	GetCategoryById(ctx context.Context, categoryId uint64) (result *responses.CategoryDetailResponse, customErr *apperror.CustomError)
}
