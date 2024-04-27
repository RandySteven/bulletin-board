package apps

import (
	"task_mission/interfaces/usecases"
	"task_mission/pkg/db"
	usecases2 "task_mission/usecases"
)

type Usecases struct {
	UserUsecase usecases.IUserUsecase
	TaskUsecase usecases.ITaskUsecase
}

func NewUseCases(repo *db.Repositories) *Usecases {
	return &Usecases{
		UserUsecase: usecases2.NewUserUsecase(repo.UnitOfWork, repo.UserRepository, repo.UserProfileRepository, repo.UserRoleRepository, repo.UserCreditRepository),
		TaskUsecase: usecases2.NewTaskUsecase(repo.TaskRepository, repo.RewardRepository, repo.RewardCategoryRepository, repo.TaskRewardRepository, repo.UserRepository, repo.UserTaskRepository, repo.UnitOfWork),
	}
}
