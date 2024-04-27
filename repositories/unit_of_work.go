package repositories

import (
	"context"
	"database/sql"
	"task_mission/interfaces/repositories"
)

type unitOfWork struct {
	conn interface{}
}

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

var _ repositories.UnitOfWork = &unitOfWork{}
