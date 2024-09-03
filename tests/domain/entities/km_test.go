package entities_test

import (
	"testing"

	"car-services-api.totote05.ar/domain/entities"
	"car-services-api.totote05.ar/tests/dsl"
	"github.com/stretchr/testify/assert"
)

func TestValidateKm(t *testing.T) {
	invalidKm := dsl.NewInvalidKmFive()
	validKm := dsl.NewValidKmOne()
	assert := assert.New(t)
	suite := []struct {
		name     string
		input    *entities.Km
		expected error
	}{
		{"nil km should not have error", nil, nil},
		{"zero value km fail with message 'km has zero value'", &invalidKm, entities.ErrKmHasZeroValue},
		{"valid km should have not error", &validKm, nil},
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
