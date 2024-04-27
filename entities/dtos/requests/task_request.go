package requests

import (
	"task_mission/entities/models"
	"task_mission/utils"
)

type RewardRequest struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

type TaskRequest struct {
	Title       string `json:"title" validate:"required,min:3,max:36"`
	Description string `json:"description" validate:"required"`
	Image       string `json:"image"`
	ExpiredDate string `json:"expired_date"`
}

type CategoriesRequest struct {
	IDs []uint64 `json:"ids"`
}

type CreateTaskRequest struct {
	Task       TaskRequest       `json:"task" validate:"required"`
	Reward     RewardRequest     `json:"reward" validate:"required"`
	Categories CategoriesRequest `json:"categories" validate:"required"`
}

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
