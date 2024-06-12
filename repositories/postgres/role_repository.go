package postgres_repositories

import (
	"context"
	"database/sql"
	"task_mission/entities/models"
	"task_mission/interfaces/repositories"
	"task_mission/queries"
	"task_mission/utils"
)

type roleRepository struct {
	db *sql.DB
}

func (r roleRepository) Save(ctx context.Context, request *models.Role) (result *uint64, err error) {
	return utils.Save[models.Role](ctx, r.db, queries.InsertRole, request.ID, request.Role)
}

func (r roleRepository) FindAll(ctx context.Context) (result []*models.Role, err error) {
	return
}

func (r roleRepository) Find(ctx context.Context, id uint64) (result *models.Role, err error) {
	return
}

func (r roleRepository) Delete(ctx context.Context, id uint64) (err error) {
	return
}

func (r roleRepository) Update(ctx context.Context, request *models.Role) (result *models.Role, err error) {
	return
}

func NewRoleRepository(db *sql.DB) *roleRepository {
	return &roleRepository{
		db: db,
	}
}

var _ repositories.IRoleRepository = &roleRepository{}
