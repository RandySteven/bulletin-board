package repositories

import (
	"context"
	"database/sql"
	"task_mission/entities/models"
	"task_mission/interfaces/repositories"
	"task_mission/queries"
	"task_mission/utils"
)

type taskRewardRepository struct {
	db *sql.DB
}

func (t *taskRewardRepository) FindByTaskId(ctx context.Context, taskId uint64) (result *models.TaskReward, err error) {
	result = &models.TaskReward{}
	err = t.db.QueryRowContext(ctx, queries.SelectTaskRewardByTaskID.ToString(), taskId).Scan(
		&result.ID,
		&result.TaskID,
		&result.RewardID,
		&result.CreatedAt, &result.UpdatedAt, &result.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t *taskRewardRepository) Save(ctx context.Context, request *models.TaskReward) (result *uint64, err error) {
	return utils.Save[models.TaskReward](ctx, t.db, queries.InsertIntoTaskReward, request.TaskID, request.RewardID)
}

func (t *taskRewardRepository) FindAll(ctx context.Context) (result []*models.TaskReward, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *taskRewardRepository) Find(ctx context.Context, id uint64) (result *models.TaskReward, err error) {

	return
}

func (t *taskRewardRepository) Delete(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (t *taskRewardRepository) Update(ctx context.Context, request *models.TaskReward) (result *models.TaskReward, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *taskRewardRepository) FindByRewardId(ctx context.Context, rewardId uint64) (result *models.TaskReward, err error) {
	result = &models.TaskReward{}
	err = t.db.QueryRowContext(ctx, queries.SelectTaskRewardByRewardID.ToString(), rewardId).Scan(
		&result.ID,
		&result.TaskID,
		&result.RewardID,
		&result.CreatedAt, &result.UpdatedAt, &result.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NewTaskRewardRepository(db *sql.DB) *taskRewardRepository {
	return &taskRewardRepository{db: db}
}

var _ repositories.ITaskRewardRepository = &taskRewardRepository{}
