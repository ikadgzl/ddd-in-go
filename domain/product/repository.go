package product

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ikadgzl/ddd-in-go/aggregate"
)

var (
	ErrProductNotFound      = errors.New("product not found")
	ErrProductAlreadyExists = errors.New("product already exists")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(uuid.UUID) (aggregate.Product, error)
	Add(aggregate.Product) error
	Update(aggregate.Product) error
	Delete(uuid.UUID) error
}
