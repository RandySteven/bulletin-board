package repositories

import (
	"context"
)

type UnitOfWork interface {
	Begin(ctx context.Context) (UnitOfWork, error)
	Rollback() error
	Commit() error
	NewUserRepository() IUserRepository
	NewRoleRepository() IRoleRepository
	NewCategoryRepository() ICategoryRepository
	NewUserProfileRepository() IUserProfileRepository
	NewUserRoleRepository() IUserRoleRepository
	NewTaskRepository() ITaskRepository
	NewRewardRepository() IRewardRepository
	NewTaskRewardRepository() ITaskRewardRepository
	NewRewardCategoryRepository() IRewardCategoryRepository
	NewUserCreditRepository() IUserCreditRepository
	NewUserTaskRepository() IUserTaskRepository
	NewRelationRepository() IRelationRepository
	NewCreditRepository() ICreditRepository
}
