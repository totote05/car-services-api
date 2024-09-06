package usecases_test

import (
	"context"
	"testing"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/usecases"
	"car-services-api.totote05.ar/tests/dsl"
	"car-services-api.totote05.ar/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAdapterGetVehicleFail(t *testing.T) {
	ctx := context.Background()
	vehicleID := dsl.NewValidVehicleOne().ID

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("Get", ctx, vehicleID).Return(nil, adapters.ErrGetting)

	usecase := usecases.NewGetVehicle(vehicleAdapter)
	_, err := usecase.Execute(ctx, vehicleID)

	assert.ErrorIs(t, err, adapters.ErrGetting)
}

func TestGetVehicleNotFound(t *testing.T) {
	ctx := context.Background()
	vehicleID := dsl.NewValidVehicleOne().ID

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("Get", ctx, vehicleID).Return(nil, adapters.ErrNotFound)

	usecase := usecases.NewGetVehicle(vehicleAdapter)
	_, err := usecase.Execute(ctx, vehicleID)

	assert.ErrorIs(t, err, adapters.ErrNotFound)
}

func TestGetVehicleSuccess(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	vehicle := dsl.NewValidVehicleOne()

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("Get", ctx, vehicle.ID).Return(&vehicle, nil)

	usecase := usecases.NewGetVehicle(vehicleAdapter)
	result, err := usecase.Execute(ctx, vehicle.ID)

	assert.Nil(err)
	assert.Equal(&vehicle, result)
}
