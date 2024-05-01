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

func (r *relationRepository) FindRelationWithUser(ctx context.Context, userId uint) (result []*models.Relation, err error) {
	rows, err := r.db.QueryContext(ctx, queries.SelectUserRelations, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		relation := &models.Relation{}
		rows.Scan(relation.ID, relation.UserID, relation.FriendID, relation.RelationStatus, relation.CreatedAt, relation.UpdatedAt, relation.DeletedAt)
		result = append(result, relation)
	}
	return result, nil
}

func (r *relationRepository) FindRelationWithFriend(ctx context.Context, userId uint, friendId uint) (result *models.Relation, err error) {
	result = &models.Relation{}
	err = r.db.QueryRowContext(ctx, queries.SelectUserFriendRelation, userId, friendId).Scan(
		&result.ID,
		&result.UserID,
		&result.FriendID,
		&result.RelationStatus,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *relationRepository) Save(ctx context.Context, request *models.Relation) (result *uint64, err error) {
	return utils.Save[models.Relation](ctx, r.db, queries.InsertIntoRelation, request.UserID, request.FriendID)
}

func (r *relationRepository) FindAll(ctx context.Context) (result []*models.Relation, err error) {
	var relation = &models.Relation{}
	return utils.FindAll[models.Relation](ctx, r.db, queries.SelectAllRelations, relation)
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
