package memory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/ikadgzl/ddd-in-go/aggregate"
	"github.com/ikadgzl/ddd-in-go/domain/product"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (m *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product

	for _, product := range m.products {
		products = append(products, product)
	}

	return products, nil
}

func (m *MemoryProductRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := m.products[id]; ok {
		return product, nil
	}

	return aggregate.Product{}, product.ErrProductNotFound
}

func (m *MemoryProductRepository) Add(p aggregate.Product) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.products[p.GetID()]; ok {
		return product.ErrProductAlreadyExists
	}

	m.products[p.GetID()] = p
	return nil
}

func (m *MemoryProductRepository) Update(p aggregate.Product) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.products[p.GetID()]; !ok {
		return product.ErrProductNotFound
	}

	m.products[p.GetID()] = p
	return nil
}

func (m *MemoryProductRepository) Delete(id uuid.UUID) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.products[id]; !ok {
		return product.ErrProductNotFound
	}

	delete(m.products, id)
	return nil
}
