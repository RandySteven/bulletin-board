package repositories

import "context"

type UnitOfWork interface {
	Begin(ctx context.Context) (UnitOfWork, error)
	Rollback() error
	Commit() error
}
