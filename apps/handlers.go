package apps

import (
	handlers2 "task_mission/handlers"
	"task_mission/interfaces/handlers"
	"task_mission/pkg/db"
)

type Handlers struct {
	UserHandler handlers.IUserHandler
	TaskHandler handlers.ITaskHandler
}

func NewHandlers(repo *db.Repositories) *Handlers {
	usecases := NewUseCases(repo)
	return &Handlers{
		UserHandler: handlers2.NewUserHandler(usecases.UserUsecase),
		TaskHandler: handlers2.NewTaskHandler(usecases.TaskUsecase),
	}
}
