package repositories

import (
	"context"
	"database/sql"
	"log"
	"task_mission/entities/models"
	"task_mission/interfaces/repositories"
	"task_mission/queries"
	"task_mission/utils"
)

type taskRepository struct {
	db *sql.DB
}

func (t *taskRepository) Save(ctx context.Context, request *models.Task) (result *uint64, err error) {
	return utils.Save[models.Task](ctx, t.db, queries.InsertTask, request.Title, request.Description, request.Image, request.UserID, request.ExpiredDate, request.Status)
}

func (t *taskRepository) FindAll(ctx context.Context) (result []*models.Task, err error) {
	var task = &models.Task{}
	result, err = utils.FindAll[models.Task](ctx, t.db, queries.SelectAllTasks, task)
	if err != nil {
		return nil, err
	}
	log.Println("result : ", result)
	return result, nil
}

func (t *taskRepository) Find(ctx context.Context, id uint64) (result *models.Task, err error) {
	//utils.FindByID[models.Task](ctx, t.db, queries.SelectTaskByID, id)
	return
}

func (t *taskRepository) Delete(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (t *taskRepository) Update(ctx context.Context, request *models.Task) (result *models.Task, err error) {
	//TODO implement me
	panic("implement me")
}

var _ repositories.ITaskRepository = &taskRepository{}

func NewTaskRepository(db *sql.DB) *taskRepository {
	return &taskRepository{db: db}
}
