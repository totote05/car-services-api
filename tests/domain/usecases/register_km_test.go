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

func TestRegisterKm(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	vehicle := dsl.NewValidVehicleOne()
	suite := []struct {
		name         string
		vehicle      *entities.Vehicle
		km           entities.Km
		list         []entities.Km
		shouldSave   bool
		err          error
		serviceError error
		vehicleError error
	}{
		{
			name:         "register service should fail when km service fails",
			vehicle:      &vehicle,
			km:           dsl.NewValidKmOne(),
			list:         []entities.Km{},
			shouldSave:   false,
			err:          adapters.ErrNotFound,
			serviceError: adapters.ErrNotFound,
		},
		{
			name:         "register first km successfully",
			vehicle:      &vehicle,
			km:           dsl.NewValidKmOne(),
			list:         []entities.Km{},
			shouldSave:   true,
			err:          nil,
			serviceError: nil,
		},
		{
			name:         "register second km successfully",
			vehicle:      &vehicle,
			km:           dsl.NewValidKmTwo(),
			list:         []entities.Km{dsl.NewValidKmOne()},
			shouldSave:   true,
			err:          nil,
			serviceError: nil,
		},
		{
			name:    "register older km with newer value should fail",
			vehicle: &vehicle,
			km:      dsl.NewInvalidKm(),
			list: []entities.Km{
				dsl.NewValidKmOne(),
				dsl.NewValidKmTwo(),
			},
			shouldSave:   false,
			err:          usecases.ErrInvalidKmData,
			serviceError: nil,
		},
		{
			name:         "register should fail when vehicle service fails",
			vehicle:      nil,
			km:           dsl.NewValidKmOne(),
			list:         nil,
			shouldSave:   false,
			err:          adapters.ErrNotFound,
			serviceError: nil,
			vehicleError: adapters.ErrNotFound,
		},
		{
			name:    "register should fail when km service fails",
			vehicle: &vehicle,
			km:      dsl.NewInvalidKmTwo(),
			list: []entities.Km{
				dsl.NewValidKmOne(),
				dsl.NewValidKmTwo(),
			},
			shouldSave:   false,
			err:          usecases.ErrInvalidKmData,
			serviceError: nil,
		},
	}

	for _, test := range suite {
		t.Run(test.name, func(t *testing.T) {
			vehicleAdapter := mocks.NewVehicle(t)
			vehicleAdapter.On("Get", ctx, vehicle.ID).Return(test.vehicle, test.vehicleError)

			kmAdapter := mocks.NewKm(t)
			if test.shouldSave {
				kmAdapter.On("Save", ctx, vehicle.ID, test.km).Return(test.err)
			}

			if test.vehicleError == nil {
				kmAdapter.On("GetAll", ctx, vehicle.ID).Return(test.list, test.serviceError)
			}

			usecase := usecases.NewRegisterKm(kmAdapter, vehicleAdapter)

			err := usecase.Execute(ctx, vehicle.ID, test.km)

			assert.Equal(test.err, err)
		})
	}
}
