package service

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ikadgzl/ddd-in-go/aggregate"
)

func Test_Tavern(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	c, err := aggregate.NewCustomer("Ilker")
	if err != nil {
		t.Fatal(err)
	}

	err = os.cr.Add(c)
	if err != nil {
		t.Fatal(err)
	}

	orders := []uuid.UUID{
		products[0].GetID(),
		products[1].GetID(),
		products[2].GetID(),
	}

	err = tavern.Order(c.GetID(), orders)
	if err != nil {
		t.Fatal(err)
	}

}
