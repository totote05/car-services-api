package usecases_test

import (
	"context"
	"testing"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
	"car-services-api.totote05.ar/domain/usecases"
	"car-services-api.totote05.ar/tests/dsl"
	"car-services-api.totote05.ar/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetKm(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	vehicle := dsl.NewValidVehicleOne()
	km := dsl.NewValidKmOne()
	suite := []struct {
		name         string
		vehicle      *entities.Vehicle
		km           *entities.Km
		vehicleError error
		kmError      error
		err          error
	}{
		{
			name:         "get km successfully",
			vehicle:      &vehicle,
			km:           &km,
			vehicleError: nil,
			kmError:      nil,
			err:          nil,
		},
		{
			name:         "should fail when vehicle service fails",
			vehicle:      nil,
			km:           nil,
			vehicleError: adapters.ErrNotFound,
			kmError:      nil,
			err:          adapters.ErrNotFound,
		},
		{
			name:         "should fail when km service fails",
			vehicle:      &vehicle,
			km:           nil,
			vehicleError: nil,
			kmError:      adapters.ErrNotFound,
			err:          adapters.ErrNotFound,
		},
	}

	for _, test := range suite {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			vehicleAdapter := mocks.NewVehicle(t)
			vehicleAdapter.On("Get", ctx, vehicle.ID).Return(test.vehicle, test.vehicleError)

			kmAdapter := mocks.NewKm(t)
			if test.vehicleError == nil {
				kmAdapter.On("Get", ctx, vehicle.ID, km.ID).Return(test.km, test.kmError)
			}

			usecase := usecases.NewGetKm(kmAdapter, vehicleAdapter)
			result, err := usecase.Execute(ctx, vehicle.ID, km.ID)

			assert.Equal(test.err, err)
			assert.Equal(test.km, result)
		})
	}
}
