package repositories

import (
	"context"
	"database/sql"
	"task_mission/entities/models"
	"task_mission/interfaces/repositories"
	"task_mission/queries"
	"task_mission/utils"
)

type userProfileRepository struct {
	db *sql.DB
}

func (u *userProfileRepository) FindByEmail(ctx context.Context, email string) (result *models.UserProfile, err error) {
	result = &models.UserProfile{}
	err = u.db.QueryRowContext(ctx, queries.SelectUserProfileByEmail, email).Scan(
		&result.ID, &result.Email, &result.Password, &result.Image, &result.UserID,
		&result.CreatedAt, &result.UpdatedAt, &result.DeletedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userProfileRepository) Save(ctx context.Context, request *models.UserProfile) (result *uint64, err error) {
	return utils.Save[models.UserProfile](ctx, u.db, queries.InsertUserProfile, &request.Email, &request.Password, &request.Image, &request.UserID)
}

func (u *userProfileRepository) FindAll(ctx context.Context) (result []*models.UserProfile, err error) {
	//var userProfile = &models.UserProfile{}
	return utils.FindAll[models.UserProfile](ctx, u.db, queries.SelectAllUserProfiles)
}

func (u *userProfileRepository) FindByUserID(ctx context.Context, userId uint64) (result *models.UserProfile, err error) {
	result = &models.UserProfile{}
	err = utils.FindByID[models.UserProfile](ctx, u.db, queries.SelectUserProfileByUserID, userId, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userProfileRepository) Find(ctx context.Context, id uint64) (result *models.UserProfile, err error) {

	//err = utils.FindByID[models.UserProfile](ctx, u.db, queries.SelectUserProfileByID, id)
	return
}

func (u *userProfileRepository) Delete(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userProfileRepository) Update(ctx context.Context, request *models.UserProfile) (result *models.UserProfile, err error) {
	//TODO implement me
	panic("implement me")
}

var _ repositories.IUserProfileRepository = &userProfileRepository{}

func NewUserProfileRepository(db *sql.DB) *userProfileRepository {
	return &userProfileRepository{db: db}
}
