package postgres_repositories

import (
	"context"
	"database/sql"
	"task_mission/entities/models"
	"task_mission/interfaces/repositories"
	"task_mission/queries"
	"task_mission/utils"
)

type userTaskRepository struct {
	db *sql.DB
}

func (u *userTaskRepository) Save(ctx context.Context, request *models.UserTask) (result *uint64, err error) {
	return utils.Save[models.UserTask](ctx, u.db, queries.InsertUserTask, request.UserID, request.TaskID)
}

func (u *userTaskRepository) FindAll(ctx context.Context) (result []*models.UserTask, err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userTaskRepository) Find(ctx context.Context, id uint64) (result *models.UserTask, err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userTaskRepository) Delete(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userTaskRepository) Update(ctx context.Context, request *models.UserTask) (result *models.UserTask, err error) {
	//TODO implement me
	panic("implement me")
}

var _ repositories.IUserTaskRepository = &userTaskRepository{}

func NewUserTaskRepository(db *sql.DB) *userTaskRepository {
	return &userTaskRepository{db: db}
}
