package repositories

import (
	"context"
	"database/sql"
	"task_mission/entities/models"
	"task_mission/enums"
	"task_mission/interfaces/repositories"
	"task_mission/queries"
	"task_mission/utils"
)

type relationRepository struct {
	db *sql.DB
}

func (r *relationRepository) FindUserFollowings(ctx context.Context, userId uint64) (result []*models.Relation, err error) {
	return r.findUserRelation(ctx, queries.SelectFollowingRelations, userId)
}

func (r *relationRepository) FindRelationWithFriend(ctx context.Context, userId uint64, friendId uint64) (result *models.Relation, err error) {
	result = &models.Relation{}
	err = r.db.QueryRowContext(ctx, queries.SelectUserFriendRelation.ToString(), userId, friendId).Scan(
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

func (r *relationRepository) FindUserFollowers(ctx context.Context, userId uint64) (result []*models.Relation, err error) {
	return r.findUserRelation(ctx, queries.SelectFollowersRelations, userId)
}

func (r *relationRepository) Save(ctx context.Context, request *models.Relation) (result *uint64, err error) {
	return utils.Save[models.Relation](ctx, r.db, queries.InsertIntoRelation, request.UserID, request.FriendID, enums.Following)
}

func (r *relationRepository) FindAll(ctx context.Context) (result []*models.Relation, err error) {
	var relation = &models.Relation{}
	return utils.FindAll[models.Relation](ctx, r.db, queries.SelectAllRelations, relation)
}

func (r *relationRepository) Find(ctx context.Context, id uint64) (result *models.Relation, err error) {
	result = &models.Relation{}
	err = utils.FindByID[models.Relation](ctx, r.db, queries.SelectRelationByID, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *relationRepository) Delete(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r *relationRepository) Update(ctx context.Context, request *models.Relation) (result *models.Relation, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *relationRepository) findUserRelation(ctx context.Context, query queries.GoQuery, userId uint64) (result []*models.Relation, err error) {
	rows, err := r.db.QueryContext(ctx, query.ToString(), userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		relation := &models.Relation{}
		rows.Scan(
			&relation.ID,
			&relation.UserID,
			&relation.FriendID,
			&relation.RelationStatus,
			&relation.CreatedAt,
			&relation.UpdatedAt,
			&relation.DeletedAt,
		)
		result = append(result, relation)
	}
	return result, nil
}

func NewRelationRepository(db *sql.DB) *relationRepository {
	return &relationRepository{db: db}
}

var _ repositories.IRelationRepository = &relationRepository{}
