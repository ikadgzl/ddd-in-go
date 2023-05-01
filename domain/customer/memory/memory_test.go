package memory

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/ikadgzl/ddd-in-go/aggregate"
	"github.com/ikadgzl/ddd-in-go/domain/customer"
)

func TestMemory_GetCustom(t *testing.T) {
	type testCase struct {
		test        string
		id          uuid.UUID
		expectedErr error
	}

	c, err := aggregate.NewCustomer("Ilker")
	if err != nil {
		t.Fatal(err)
	}

	id := c.GetID()

	repo := New()
	repo.customers[id] = c

	testCases := []testCase{
		{
			test:        "no customer by id",
			id:          uuid.MustParse("de26c30a-e78d-11ed-a05b-0242ac120003"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			test:        "customer by id",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := repo.Get(tc.id)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
