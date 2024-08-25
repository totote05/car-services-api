package entities_test

import (
	"testing"

	"car-services-api.totote05.ar/domain/entities"
	"car-services-api.totote05.ar/tests/dsl"
	"github.com/stretchr/testify/assert"
)

func TestValidateVehicle(t *testing.T) {
	invalidVehicle := dsl.NewInvalidVehicle()
	validVehicle := dsl.NewValidVehicleOne()
	assert := assert.New(t)
	suite := []struct {
		name     string
		input    *entities.Vehicle
		expected error
	}{
		{"nil vehicle should not have error", nil, nil},
		{"empty plate fail with message 'empty plate'", &invalidVehicle, entities.ErrVehicleHasEmptyPlate},
		{"valid vehicle should have not error", &validVehicle, nil},
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
