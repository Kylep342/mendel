package db

import "context"

// CRUDTable is an interface for go_model-to-db_record mapping as T
//
//	T - a table schema in GO as a struct
type CRUDTable[T any] interface {
	GetAll(ctx context.Context) ([]T, error)
	GetByID(ctx context.Context, id string) (T, error)
	Create(ctx context.Context, item *T) error
	Update(ctx context.Context, item *T) error
	Delete(ctx context.Context, id string) error
}
