package postgres_repositories

import (
	"context"
	"database/sql"
	"task_mission/entities/models"
	"task_mission/interfaces/repositories"
	"task_mission/queries"
	"task_mission/utils"
)

type userRoleRepository struct {
	db *sql.DB
}

func (u *userRoleRepository) FindByUserID(ctx context.Context, userId uint64) (result *models.UserRole, err error) {
	result = &models.UserRole{}
	err = u.db.QueryRowContext(ctx, queries.SelectUserRoleByUserID, userId).Scan(&result.ID, &result.UserID, &result.RoleID, &result.CreatedAt, &result.UpdatedAt, &result.DeletedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userRoleRepository) Save(ctx context.Context, request *models.UserRole) (result *uint64, err error) {
	return utils.Save[models.UserRole](ctx, u.db, queries.InsertUserRole, &request.UserID, &request.RoleID)
}

func (u *userRoleRepository) FindAll(ctx context.Context) (result []*models.UserRole, err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRoleRepository) Find(ctx context.Context, id uint64) (result *models.UserRole, err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRoleRepository) Delete(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRoleRepository) Update(ctx context.Context, request *models.UserRole) (result *models.UserRole, err error) {
	//TODO implement me
	panic("implement me")
}

var _ repositories.IUserRoleRepository = &userRoleRepository{}

func NewUserRoleRepository(db *sql.DB) *userRoleRepository {
	return &userRoleRepository{
		db: db,
	}
}
