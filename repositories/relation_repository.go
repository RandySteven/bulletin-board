package repositories

import (
	"context"
	"database/sql"
	"task_mission/entities/models"
	"task_mission/interfaces/repositories"
	"task_mission/queries"
	"task_mission/utils"
)

type relationRepository struct {
	db *sql.DB
}

func (r *relationRepository) Save(ctx context.Context, request *models.Relation) (result *uint64, err error) {
	return utils.Save[models.Relation](ctx, r.db, queries.InsertIntoRelation, request.UserID, request.FriendID)
}

func (r *relationRepository) FindAll(ctx context.Context) (result []*models.Relation, err error) {
	return
}

func (r *relationRepository) Find(ctx context.Context, id uint64) (result *models.Relation, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *relationRepository) Delete(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r *relationRepository) Update(ctx context.Context, request *models.Relation) (result *models.Relation, err error) {
	//TODO implement me
	panic("implement me")
}

func NewRelationRepository(db *sql.DB) *relationRepository {
	return &relationRepository{db: db}
}

var _ repositories.IRelationRepository = &relationRepository{}
