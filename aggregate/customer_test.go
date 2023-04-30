package aggregate_test

import (
	"errors"
	"testing"

	"github.com/ikadgzl/ddd-in-go/aggregate"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Invalid name, empty",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		},
		{
			test:        "Valid name",
			name:        "Ilker Adiguzel",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
