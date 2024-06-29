package apps

import (
	handlers2 "task_mission/handlers"
	"task_mission/interfaces/handlers"
	"task_mission/pkg/db"
)

type Handlers struct {
	DevHandler      handlers.IDevHandler
	UserHandler     handlers.IUserHandler
	TaskHandler     handlers.ITaskHandler
	RelationHandler handlers.IRelationHandler
	RewardHandler   handlers.IRewardHandler
	CategoryHandler handlers.ICategoryHandler
	CreditHandler   handlers.ICreditHandler
	ChatHandler     handlers.IChatHandler
}

func NewHandlers(repo *db.Repositories, service *Services) *Handlers {
	usecases := NewUseCases(repo)
	return &Handlers{
		DevHandler:      handlers2.NewDevHandler(),
		UserHandler:     handlers2.NewUserHandler(usecases.UserUsecase),
		TaskHandler:     handlers2.NewTaskHandler(usecases.TaskUsecase),
		RelationHandler: handlers2.NewRelationHandler(usecases.RelationUsecase),
		RewardHandler:   handlers2.NewRewardHandler(usecases.RewardUsecase),
		CategoryHandler: handlers2.NewCategoryHandler(usecases.CategoryUsecase),
		CreditHandler:   handlers2.NewCreditHandler(usecases.CreditUsecase),
		ChatHandler:     handlers2.NewChatHandler(usecases.ChatUsecase),
	}
}
