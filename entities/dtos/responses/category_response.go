package responses

import (
	"task_mission/entities/models"
	"time"
)

type (
	CategoryResponse struct {
		ID        uint64     `json:"id"`
		Category  string     `json:"category"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}

	CategoryDetailResponse struct {
		ID       uint64 `json:"id"`
		Category string `json:"category"`
		Rewards  []*struct {
			ID          uint64 `json:"id"`
			Name        string `json:"name"`
			Image       string `json:"image"`
			Description string `json:"description"`
		}
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}
)

func NewCategoryListResponse(category *models.Category) *CategoryResponse {
	return &CategoryResponse{
		ID:        category.ID,
		Category:  category.Category,
		CreatedAt: category.CreatedAt.Local(),
		UpdatedAt: category.UpdatedAt.Local(),
		DeletedAt: category.DeletedAt,
	}
}

func NewCategoryListResponses(categoryList []*models.Category) []*CategoryResponse {
	var result []*CategoryResponse
	for _, category := range categoryList {
		result = append(result, &CategoryResponse{
			ID:        category.ID,
			Category:  category.Category,
			CreatedAt: category.CreatedAt.Local(),
			UpdatedAt: category.UpdatedAt.Local(),
			DeletedAt: category.DeletedAt,
		})
	}
	return result
}

func NewCategoryDetailResponse(category *models.Category, rewards []*models.Reward) *CategoryDetailResponse {
	var rewardResponses []*struct {
		ID          uint64 `json:"id"`
		Name        string `json:"name"`
		Image       string `json:"image"`
		Description string `json:"description"`
	}

	for _, reward := range rewards {
		rewardResponses = append(rewardResponses, &struct {
			ID          uint64 `json:"id"`
			Name        string `json:"name"`
			Image       string `json:"image"`
			Description string `json:"description"`
		}{ID: reward.ID, Name: reward.Name, Image: reward.Image, Description: reward.Description})
	}

	return &CategoryDetailResponse{
		ID:        category.ID,
		Category:  category.Category,
		Rewards:   rewardResponses,
		CreatedAt: category.CreatedAt.Local(),
		UpdatedAt: category.UpdatedAt.Local(),
		DeletedAt: category.DeletedAt,
	}
}
