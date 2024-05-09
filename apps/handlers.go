package apps

import (
	handlers2 "task_mission/handlers"
	"task_mission/interfaces/handlers"
	"task_mission/pkg/db"
)

type Handlers struct {
	UserHandler     handlers.IUserHandler
	TaskHandler     handlers.ITaskHandler
	RelationHandler handlers.IRelationHandler
	RewardHandler   handlers.IRewardHandler
	CategoryHandler handlers.ICategoryHandler
}

func NewHandlers(repo *db.Repositories, service *Services) *Handlers {
	usecases := NewUseCases(repo)
	return &Handlers{
		UserHandler:     handlers2.NewUserHandler(usecases.UserUsecase),
		TaskHandler:     handlers2.NewTaskHandler(usecases.TaskUsecase),
		RelationHandler: handlers2.NewRelationHandler(usecases.RelationUsecase),
		RewardHandler:   handlers2.NewRewardHandler(usecases.RewardUsecase),
		CategoryHandler: handlers2.NewCategoryHandler(usecases.CategoryUsecase),
	}
}
