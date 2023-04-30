package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/ikadgzl/ddd-in-go/aggregate"
	"github.com/ikadgzl/ddd-in-go/domain/customer"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (m *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := m.customers[id]; ok {
		return customer, nil
	}

	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (m *MemoryRepository) Add(c aggregate.Customer) error {
	if _, ok := m.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists %w", customer.ErrFailedToAddCustomer)
	}

	m.Lock()
	m.customers[c.GetID()] = c
	m.Unlock()

	return nil
}

func (m *MemoryRepository) Update(c aggregate.Customer) error {
	if _, ok := m.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist %w", customer.ErrFailedToUpdateCustomer)
	}

	m.Lock()
	m.customers[c.GetID()] = c
	m.Unlock()

	return nil
}
