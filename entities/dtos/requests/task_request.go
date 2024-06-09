package requests

import (
	"task_mission/entities/models"
	"task_mission/utils"
)

type (
	RewardRequest struct {
		Name        string `json:"name" form:"name"`
		Image       string `json:"image" form:"image"`
		Description string `json:"description" form:"description"`
	}

	TaskRequest struct {
		Title       string `json:"title" validate:"required,min:3,max:36" form:"title"`
		Description string `json:"description" validate:"required" form:"description"`
		Image       string `json:"image" form:"image"`
		ExpiredDate string `json:"expired_date" form:"expired_date"`
	}

	CategoriesRequest struct {
		IDs []uint64 `json:"ids"`
	}

	CreateTaskRequest struct {
		Task       TaskRequest       `json:"task" validate:"required" form:"task"`
		Reward     RewardRequest     `json:"reward" validate:"required" form:"reward"`
		Categories CategoriesRequest `json:"categories" validate:"required" form:"categories"`
	}
)

func (request *CreateTaskRequest) ConvertTask() *models.Task {
	taskRequest := request.Task
	expiredDate, err := utils.StringToDate(request.Task.ExpiredDate)
	if err != nil {
		return nil
	}
	task := &models.Task{
		Title:       taskRequest.Title,
		Description: taskRequest.Description,
		Image:       taskRequest.Image,
		ExpiredDate: expiredDate,
	}
	return task
}

func (request *CreateTaskRequest) ConvertReward() *models.Reward {
	rewardRequest := request.Reward
	reward := &models.Reward{
		Name:        rewardRequest.Name,
		Image:       rewardRequest.Image,
		Description: rewardRequest.Description,
	}
	return reward
}
