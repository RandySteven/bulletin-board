package repositories

import (
	"context"
	"database/sql"
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
	var user = &models.User{}
	return utils.FindAll[models.User](ctx, u.db, queries.SelectAllUser, user)
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

func (u *userRepository) Update(ctx context.Context, request *models.User) (result *models.User, err error) {
	//TODO implement me
	panic("implement me")
}

var _ repositories.IUserRepository = &userRepository{}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}
