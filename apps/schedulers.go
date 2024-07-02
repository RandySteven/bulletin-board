package apps

import (
	"task_mission/pkg/db"
	"task_mission/pkg/scheduler"
	"task_mission/usecases"
)

func NewSchedulers(repo *db.Repositories) (usecase scheduler.UsecaseDependency) {
	usecase = scheduler.UsecaseDependency{
		TaskUsecase: usecases.NewTaskUsecase(
			repo.TaskRepository,
			repo.RewardRepository,
			repo.RewardCategoryRepository,
			repo.TaskRewardRepository,
			repo.UserRepository,
			repo.UserProfileRepository,
			repo.UserTaskRepository,
			repo.UnitOfWork,
		),
	}
	return
}
