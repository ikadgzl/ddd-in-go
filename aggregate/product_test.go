package aggregate_test

import (
	"errors"
	"testing"

	"github.com/ikadgzl/ddd-in-go/aggregate"
)

func TestProduct_NewProduct(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		description string
		price       float64
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Invalid name, empty",
			name:        "",
			description: "d",
			price:       0,
			expectedErr: aggregate.ErrMissingValue,
		},
		{
			test:        "Invalid description, empty",
			name:        "n",
			description: "",
			price:       0,
			expectedErr: aggregate.ErrMissingValue,
		},
		{
			test:        "Valid product",
			name:        "Beer",
			description: "Beverage",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewProduct(tc.name, tc.description, tc.price)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
