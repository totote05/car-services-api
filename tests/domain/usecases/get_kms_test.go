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

func TestGetKms(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	vehicle := dsl.NewValidVehicleOne()
	suite := []struct {
		name         string
		list         []entities.Km
		vehicle      *entities.Vehicle
		err          error
		vehicleError error
		kmError      error
	}{
		{
			name: "get all kms successfully",
			list: []entities.Km{
				dsl.NewValidKmOne(),
			},
			vehicle:      &vehicle,
			err:          nil,
			vehicleError: nil,
			kmError:      nil,
		},
		{
			name:         "should fail when vehicle service fails",
			list:         nil,
			vehicleError: adapters.ErrNotFound,
			kmError:      nil,
			err:          adapters.ErrNotFound,
		},
		{
			name:         "should fail when km service fails",
			list:         nil,
			vehicle:      &vehicle,
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
				kmAdapter.On("GetAll", ctx, vehicle.ID).Return(test.list, test.kmError)
			}

			usecase := usecases.NewGetKms(kmAdapter, vehicleAdapter)
			list, err := usecase.Execute(ctx, vehicle.ID)
			assert.Equal(test.list, list)
			assert.Equal(test.err, err)
		})
	}
}
