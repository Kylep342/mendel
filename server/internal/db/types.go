package db

import "context"

type CRUDTable[T interface{}] interface {
	GetAll(ctx context.Context) ([]T, error)
	GetByID(ctx context.Context, id string) (T, error)
	Create(ctx context.Context, item *T) error
	Update(ctx context.Context, item *T) error
	Delete(ctx context.Context, id string) error
}
