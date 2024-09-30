package usecases_test

import (
	"context"
	"testing"

	"car-services-api.totote05.ar/domain/entities"
	"car-services-api.totote05.ar/domain/usecases"
	"car-services-api.totote05.ar/tests/dsl"
	"car-services-api.totote05.ar/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetRegisteredServices(t *testing.T) {
	ctx := context.Background()
	anError := assert.AnError
	assert := assert.New(t)
	vehicle := dsl.NewValidVehicleOne()
	suite := []struct {
		name       string
		vehicle    *entities.Vehicle
		vehicleErr error
		expected   []entities.ServiceRegister
		err        error
	}{
		{
			name:       "vehicle service fails",
			vehicleErr: anError,
			err:        anError,
		},
		{
			name:    "service register service fails",
			vehicle: &vehicle,
			err:     anError,
		},
		{
			name:    "get all services successfully",
			vehicle: &vehicle,
			expected: []entities.ServiceRegister{
				dsl.NewValidServiceRegister(),
			},
		},
	}

	for _, test := range suite {
		t.Run(test.name, func(t *testing.T) {
			// t.Parallel()
			vehicleAdapter := mocks.NewVehicle(t)
			vehicleAdapter.On("Get", ctx, vehicle.ID).Return(test.vehicle, test.vehicleErr)

			serviceRegisterAdapter := mocks.NewServiceRegister(t)
			if test.vehicleErr == nil {
				serviceRegisterAdapter.On("GetAll", ctx, vehicle.ID).Return(test.expected, test.err)
			}

			usecase := usecases.NewGetRegisteredServices(serviceRegisterAdapter, vehicleAdapter)
			result, err := usecase.Execute(ctx, vehicle.ID)

			assert.Equal(test.expected, result)
			assert.Equal(test.err, err)
		})
	}
}
