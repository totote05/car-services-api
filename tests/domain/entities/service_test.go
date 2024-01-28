package entities_test

import (
	"testing"

	"car-services-api.totote05.ar/domain/entities"
	"car-services-api.totote05.ar/tests/dsl"
	"github.com/stretchr/testify/assert"
)

func TestValidateService(t *testing.T) {
	assert := assert.New(t)
	suite := []struct {
		name     string
		input    *entities.Service
		expected error
	}{
		{"nil service should not have error", nil, nil},
		{"service with empty name should have error", dsl.NewInvalidService(), entities.ErrServiceHasEmptyName},
		{"service with valid name and nil recurrences should not have error", dsl.NewValidService(), nil},
		{"service with valid name and empty recurrences should have error", dsl.NewValidServiceWithInvalidRecurrence(), entities.ErrServiceHasEmptyRecurrence},
		{"service with valid name and recurrence should not have error", dsl.NewValidServiceOne(), nil},
	}

	for _, test := range suite {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := test.input.Validate()
			if test.expected != nil {
				assert.ErrorIs(err, test.expected)
			} else {
				assert.Nil(err)
			}
		})
	}
}
