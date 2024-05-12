package apps

import (
	"task_mission/interfaces/usecases"
	"task_mission/pkg/db"
	usecases2 "task_mission/usecases"
)

type Usecases struct {
	UserUsecase     usecases.IUserUsecase
	TaskUsecase     usecases.ITaskUsecase
	RelationUsecase usecases.IRelationUsecase
	RewardUsecase   usecases.IRewardUsecase
	CategoryUsecase usecases.ICategoryUsecase
	CreditUsecase   usecases.ICreditUseCase
}

func NewUseCases(repo *db.Repositories) *Usecases {
	return &Usecases{
		UserUsecase:     usecases2.NewUserUsecase(repo.UnitOfWork, repo.UserRepository, repo.UserProfileRepository, repo.UserRoleRepository, repo.UserCreditRepository),
		TaskUsecase:     usecases2.NewTaskUsecase(repo.TaskRepository, repo.RewardRepository, repo.RewardCategoryRepository, repo.TaskRewardRepository, repo.UserRepository, repo.UserProfileRepository, repo.UserTaskRepository, repo.UnitOfWork),
		RelationUsecase: usecases2.NewRelationUsecase(repo.UnitOfWork, repo.RelationRepository, repo.UserRepository, repo.CreditRepository),
		RewardUsecase:   usecases2.NewRewardUsecase(repo.UnitOfWork, repo.RewardRepository, repo.TaskRepository, repo.UserRepository, repo.UserProfileRepository, repo.CategoryRepository, repo.RewardCategoryRepository, repo.TaskRewardRepository),
		CategoryUsecase: usecases2.NewCategoryUsecase(repo.CategoryRepository, repo.RewardRepository, repo.RewardCategoryRepository),
		CreditUsecase:   usecases2.NewCreditUsecase(repo.UserRepository, repo.CreditRepository, repo.UserCreditRepository),
	}
}
