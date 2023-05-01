package service

import (
	"log"

	"github.com/google/uuid"
	"github.com/ikadgzl/ddd-in-go/aggregate"
	"github.com/ikadgzl/ddd-in-go/domain/customer"
	cm "github.com/ikadgzl/ddd-in-go/domain/customer/memory"
	"github.com/ikadgzl/ddd-in-go/domain/product"
	pm "github.com/ikadgzl/ddd-in-go/domain/product/memory"
)

type OrderConfig func(os *OrderService) error

type OrderService struct {
	cr customer.CustomerRepository
	pr product.ProductRepository
}

func NewOrderService(cfgs ...OrderConfig) (*OrderService, error) {
	os := &OrderService{}

	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfig {
	return func(os *OrderService) error {
		os.cr = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfig {
	cr := cm.New()

	return WithCustomerRepository(cr)
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfig {
	pr := pm.New()

	return func(os *OrderService) error {
		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}

		os.pr = pr
		return nil
	}
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	c, err := o.cr.Get(customerID)
	if err != nil {
		return 0, err
	}

	var products []aggregate.Product
	var total float64

	for _, id := range productIDs {
		p, err := o.pr.GetByID(id)
		if err != nil {
			return 0, err
		}

		products = append(products, p)
		total += p.GetPrice()
	}

	log.Printf("Customer with id: %s has ordered %d products", c.GetID(), len(products))

	return total, nil
}
