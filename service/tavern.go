package service

import (
	"log"

	"github.com/google/uuid"
)

type TavernConfig func(t *Tavern) error

type Tavern struct {
	OrderService *OrderService
}

func NewTavern(cfgs ...TavernConfig) (*Tavern, error) {
	t := &Tavern{}

	for _, cfg := range cfgs {
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}

	return t, nil
}

func WithOrderService(os *OrderService) TavernConfig {
	return func(t *Tavern) error {
		t.OrderService = os
		return nil
	}
}

func (t *Tavern) Order(customer uuid.UUID, productIds []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer, productIds)
	if err != nil {
		return err
	}

	log.Printf("The total price is %s for customer: %f,", customer, price)
	return nil
}
