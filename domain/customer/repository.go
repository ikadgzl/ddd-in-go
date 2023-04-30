package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ikadgzl/ddd-in-go/aggregate"
)

var (
	ErrCustomerNotFound       = errors.New("the customer was not found in the repository")
	ErrFailedToAddCustomer    = errors.New("failed to add the customer")
	ErrFailedToUpdateCustomer = errors.New("failed to update the customer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(uuid.UUID) error
}
