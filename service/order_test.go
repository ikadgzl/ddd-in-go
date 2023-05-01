package service

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ikadgzl/ddd-in-go/aggregate"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "Nice beer", 8.99)
	if err != nil {
		t.Fatal(err)
	}

	burger, err := aggregate.NewProduct("Burger", "Fills you up", 14.99)
	if err != nil {
		t.Fatal(err)
	}

	chips, err := aggregate.NewProduct("Chips", "Very crispy", 4.99)
	if err != nil {
		t.Fatal(err)
	}

	return []aggregate.Product{
		beer, burger, chips,
	}
}

func TestOrder_NewOrderService(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	c, err := aggregate.NewCustomer("Ilker")
	if err != nil {
		t.Error(err)
	}

	err = os.cr.Add(c)
	if err != nil {
		t.Error(err)
	}

	orders := []uuid.UUID{
		products[0].GetID(),
		products[1].GetID(),
		products[2].GetID(),
	}

	_, err = os.CreateOrder(c.GetID(), orders)
	if err != nil {
		t.Error(err)
	}
}
