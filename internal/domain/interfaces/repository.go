package interfaces

import "github.com/google/uuid"

type Repository[ET any, FT any] interface {
	Add(entity ET) error
	Delete(Id uuid.UUID) error
	Update(entity ET) error
	Get(filter FT) ([]ET, error)
}
