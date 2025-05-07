package db

type CRUDTable[T any] interface {
	GetAll() ([]T, error)
	GetByID(id string) (T, error)
	Create(item *T) error
	Update(item *T) error
	Delete(id string) error
}
