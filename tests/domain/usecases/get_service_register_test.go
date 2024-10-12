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

func TestGetServiceRegister(t *testing.T) {
	ctx := context.Background()
	anError := assert.AnError
	assert := assert.New(t)
	vehicle := dsl.NewValidVehicleOne()
	serviceRegister := dsl.NewValidServiceRegister()
	suite := []struct {
		name         string
		vehicle      *entities.Vehicle
		vehicleError error
		expected     *entities.ServiceRegister
		expectedErr  error
	}{
		{
			name:         "should return an error when vehicle adapter returns an error",
			vehicle:      nil,
			vehicleError: anError,
			expectedErr:  anError,
		},
		{
			name:        "should return an error when service register adapter returns an error",
			vehicle:     &vehicle,
			expectedErr: anError,
		},
		{
			name:     "should return a service register",
			vehicle:  &vehicle,
			expected: &serviceRegister,
		},
	}

	for _, test := range suite {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			vehicleAdapter := mocks.NewVehicle(t)
			vehicleAdapter.On("Get", ctx, vehicle.ID).Return(test.vehicle, test.vehicleError)

			serviceRegisterAdapter := mocks.NewServiceRegister(t)
			if test.vehicleError == nil {
				serviceRegisterAdapter.On("Get", ctx, vehicle.ID, serviceRegister.ID).Return(test.expected, test.expectedErr)
			}

			usecase := usecases.NewGetServiceRegister(serviceRegisterAdapter, vehicleAdapter)
			result, err := usecase.Execute(ctx, vehicle.ID, serviceRegister.ID)

			assert.Equal(test.expected, result)
			assert.Equal(test.expectedErr, err)
		})
	}
}
