package repositories

import "context"

type IRepository[T any] interface {
	Save(ctx context.Context, request *T) (result *uint64, err error)
	FindAll(ctx context.Context) (result []*T, err error)
	Find(ctx context.Context, id uint64) (result *T, err error)
	Delete(ctx context.Context, id uint64) (err error)
	Update(ctx context.Context, request *T) (result *T, err error)
}
