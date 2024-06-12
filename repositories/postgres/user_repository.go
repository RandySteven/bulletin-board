package postgres_repositories

import (
	"context"
	"database/sql"
	"errors"
	"task_mission/entities/models"
	"task_mission/interfaces/repositories"
	"task_mission/queries"
	"task_mission/utils"
)

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) Save(ctx context.Context, request *models.User) (result *uint64, err error) {
	return utils.Save[models.User](ctx, u.db, queries.InsertUser, &request.Name, &request.UserName, &request.DateOfBirth, request.Gender.ToString())
}

func (u *userRepository) FindAll(ctx context.Context) (result []*models.User, err error) {
	return utils.FindAll[models.User](ctx, u.db, queries.SelectAllUser)
}

func (u *userRepository) FindVerifyUser(ctx context.Context, id uint64, isVerified bool) (result *models.User, err error) {
	result = &models.User{}
	err = u.db.QueryRowContext(ctx, queries.SelectUserVerifyCondition, id, isVerified).Scan(&result.IsVerified)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userRepository) Find(ctx context.Context, id uint64) (result *models.User, err error) {
	result = &models.User{}
	err = utils.FindByID[models.User](ctx, u.db, queries.SelectUserByID, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userRepository) Delete(ctx context.Context, id uint64) (err error) {
	return utils.Delete[models.User](ctx, u.db, queries.DeleteUser, id)
}

func (u *userRepository) Update(ctx context.Context, request *models.User) (*models.User, error) {
	err := utils.Update[models.User](ctx, u.db, queries.UpdateUser, &request.Name, &request.UserName, &request.DateOfBirth, &request.Gender, &request.IsVerified, &request.ID)
	if err != nil {
		return nil, err
	}
	return request, nil
}

var _ repositories.IUserRepository = &userRepository{}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}
