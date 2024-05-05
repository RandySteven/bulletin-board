package responses

import (
	"task_mission/entities/models"
	"task_mission/enums"
	"time"
)

type (
	RewardDetailResponse struct {
		ID          uint64 `json:"id"`
		Name        string `json:"name"`
		Image       string `json:"image"`
		Description string `json:"description"`
		User        struct {
			ID       uint64 `json:"id"`
			Name     string `json:"name"`
			UserName string `json:"user_name"`
			Image    string `json:"image"`
		} `json:"user"`
		Task struct {
			ID          uint64           `json:"id"`
			Title       string           `json:"title"`
			Description string           `json:"description"`
			Image       string           `json:"image"`
			Status      enums.TaskStatus `json:"status"`
			ExpiredDate time.Time        `json:"expired_date"`
		} `json:"task"`
		Categories []CategoryResponse `json:"categories"`
		CreatedAt  time.Time          `json:"created_at"`
		UpdatedAt  time.Time          `json:"updated_at"`
		DeletedAt  *time.Time         `json:"deleted_at"`
	}

	RewardListResponse struct {
		ID          uint64     `json:"id"`
		Name        string     `json:"name"`
		Image       string     `json:"image"`
		Description string     `json:"description"`
		CreatedAt   time.Time  `json:"created_at"`
		UpdatedAt   time.Time  `json:"updated_at"`
		DeletedAt   *time.Time `json:"deleted_at"`
	}
)

func NewRewardDetailResponse(reward *models.Reward, user *models.User, profile *models.UserProfile, task *models.Task, categories []*models.Category) *RewardDetailResponse {
	categoryResponses := make([]CategoryResponse, 0)
	for _, category := range categories {
		categoryResponses = append(categoryResponses, CategoryResponse{
			ID:       category.ID,
			Category: category.Category,
		})
	}
	return &RewardDetailResponse{
		ID:          reward.ID,
		Name:        reward.Name,
		Image:       reward.Image,
		Description: reward.Description,
		User: struct {
			ID       uint64 `json:"id"`
			Name     string `json:"name"`
			UserName string `json:"user_name"`
			Image    string `json:"image"`
		}{ID: user.ID, Name: user.Name, UserName: user.UserName, Image: profile.Image},
		Task: struct {
			ID          uint64           `json:"id"`
			Title       string           `json:"title"`
			Description string           `json:"description"`
			Image       string           `json:"image"`
			Status      enums.TaskStatus `json:"status"`
			ExpiredDate time.Time        `json:"expired_date"`
		}{ID: task.ID, Title: task.Title, Description: task.Description, Image: task.Image, Status: task.Status, ExpiredDate: task.ExpiredDate.Local()},
		Categories: categoryResponses,
		CreatedAt:  reward.CreatedAt.Local(),
		UpdatedAt:  reward.UpdatedAt.Local(),
		DeletedAt:  reward.DeletedAt,
	}
}

func NewRewardListResponse(reward *models.Reward) *RewardListResponse {
	return &RewardListResponse{
		ID:          reward.ID,
		Name:        reward.Name,
		Image:       reward.Image,
		Description: reward.Description,
		CreatedAt:   reward.CreatedAt.Local(),
		UpdatedAt:   reward.UpdatedAt.Local(),
		DeletedAt:   reward.DeletedAt,
	}
}

func NewRewardListResponses(rewards []*models.Reward) []*RewardListResponse {
	result := make([]*RewardListResponse, len(rewards))
	for _, reward := range rewards {
		response := NewRewardListResponse(reward)
		result = append(result, response)
	}
	return result
}
