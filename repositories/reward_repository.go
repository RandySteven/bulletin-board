package repositories

import (
	"context"
	"database/sql"
	"task_mission/entities/models"
	"task_mission/interfaces/repositories"
	"task_mission/queries"
	"task_mission/utils"
)

type rewardRepository struct {
	db *sql.DB
}

func (r *rewardRepository) Save(ctx context.Context, request *models.Reward) (result *uint64, err error) {
	return utils.Save[models.Reward](ctx, r.db, queries.InsertIntoReward, &request.Name, &request.Description, &request.Image, &request.UserID)
}

func (r *rewardRepository) FindAll(ctx context.Context) (result []*models.Reward, err error) {
	var reward = &models.Reward{}
	result, err = utils.FindAll[models.Reward](ctx, r.db, queries.SelectAllRewards, reward)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *rewardRepository) Find(ctx context.Context, id uint64) (result *models.Reward, err error) {
	err = utils.FindByID[models.Reward](ctx, r.db, queries.SelectRewardByID, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *rewardRepository) Delete(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r *rewardRepository) Update(ctx context.Context, request *models.Reward) (result *models.Reward, err error) {
	//TODO implement me
	panic("implement me")
}

func NewRewardRepository(db *sql.DB) *rewardRepository {
	return &rewardRepository{db: db}
}

var _ repositories.IRewardRepository = &rewardRepository{}
