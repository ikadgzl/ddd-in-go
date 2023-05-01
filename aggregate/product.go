package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ikadgzl/ddd-in-go/entity"
)

type Product struct {
	item     *entity.Item
	price    float64
	quantity uint
}

var (
	ErrMissingValue = errors.New("missing one or more values")
)

func NewProduct(name, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValue
	}

	item := &entity.Item{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
	}

	return Product{
		item:     item,
		price:    price,
		quantity: 0,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *entity.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}
