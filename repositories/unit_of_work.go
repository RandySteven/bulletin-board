package repositories

import (
	"context"
	"database/sql"
	"task_mission/interfaces/repositories"
)

type (
	unitOfWork struct {
		conn interface{}
	}
)

func NewUnitOfWork(db *sql.DB) *unitOfWork {
	return &unitOfWork{conn: db}
}

func (uow *unitOfWork) Begin(ctx context.Context) (repositories.UnitOfWork, error) {
	switch c := uow.conn.(type) {
	case *sql.DB:
		tx, err := uow.conn.(*sql.DB).BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		uow.conn = tx
	case *sql.Tx:
		uow.conn = c
	}
	return uow, nil
}

func (uow *unitOfWork) Rollback() error {
	tx, ok := uow.conn.(*sql.Tx)
	if !ok {
		return nil
	}
	return tx.Rollback()
}

func (uow *unitOfWork) Commit() error {
	tx, ok := uow.conn.(*sql.Tx)
	if !ok {
		return nil
	}
	return tx.Commit()
}

func (uow *unitOfWork) NewUserRepository() repositories.IUserRepository {
	return NewUserRepository(uow.conn.(*sql.DB))
}

func (uow *unitOfWork) NewRoleRepository() repositories.IRoleRepository {
	return NewRoleRepository(uow.conn.(*sql.DB))
}

func (uow *unitOfWork) NewCategoryRepository() repositories.ICategoryRepository {
	return NewCategoryRepository(uow.conn.(*sql.DB))
}

func (uow *unitOfWork) NewUserProfileRepository() repositories.IUserProfileRepository {
	return NewUserProfileRepository(uow.conn.(*sql.DB))
}

func (uow *unitOfWork) NewUserRoleRepository() repositories.IUserRoleRepository {
	return NewUserRoleRepository(uow.conn.(*sql.DB))
}

func (uow *unitOfWork) NewTaskRepository() repositories.ITaskRepository {
	return NewTaskRepository(uow.conn.(*sql.DB))
}

func (uow *unitOfWork) NewRewardRepository() repositories.IRewardRepository {
	return NewRewardRepository(uow.conn.(*sql.DB))
}

func (uow *unitOfWork) NewTaskRewardRepository() repositories.ITaskRewardRepository {
	return NewTaskRewardRepository(uow.conn.(*sql.DB))
}

func (uow *unitOfWork) NewRewardCategoryRepository() repositories.IRewardCategoryRepository {
	return NewRewardCategoryRepository(uow.conn.(*sql.DB))
}

func (uow *unitOfWork) NewUserCreditRepository() repositories.IUserCreditRepository {
	return NewUserCreditRepository(uow.conn.(*sql.DB))
}

func (uow *unitOfWork) NewUserTaskRepository() repositories.IUserTaskRepository {
	return NewUserTaskRepository(uow.conn.(*sql.DB))
}

func (uow *unitOfWork) NewRelationRepository() repositories.IRelationRepository {
	return NewRelationRepository(uow.conn.(*sql.DB))
}

func (uow *unitOfWork) NewCreditRepository() repositories.ICreditRepository {
	return NewCreditRepository(uow.conn.(*sql.DB))
}

var _ repositories.UnitOfWork = &unitOfWork{}
