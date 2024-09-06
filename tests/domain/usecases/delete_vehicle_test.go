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

func TestDeleteVehicleFailGettingVehicle(t *testing.T) {
	ctx := context.Background()
	vehicle := dsl.NewValidVehicleOne()

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("Get", ctx, vehicle.ID).Return(nil, adapters.ErrGetting)

	usecase := usecases.NewDeleteVehicle(vehicleAdapter)
	err := usecase.Execute(ctx, vehicle.ID)

	assert.ErrorIs(t, err, adapters.ErrGetting)
}

func TestDeleteVehicleFailVehicleNotFound(t *testing.T) {
	ctx := context.Background()
	vehicle := dsl.NewValidVehicleOne()

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("Get", ctx, vehicle.ID).Return(nil, adapters.ErrNotFound)

	usecase := usecases.NewDeleteVehicle(vehicleAdapter)
	err := usecase.Execute(ctx, vehicle.ID)

	assert.ErrorIs(t, err, adapters.ErrNotFound)
}

func TestDeleteVehicleFailDeleting(t *testing.T) {
	ctx := context.Background()
	vehicle := dsl.NewValidVehicleOne()

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("Get", ctx, vehicle.ID).Return(&vehicle, nil)
	vehicleAdapter.On("Delete", ctx, vehicle.ID).Return(adapters.ErrPersisting)

	usecase := usecases.NewDeleteVehicle(vehicleAdapter)
	err := usecase.Execute(ctx, vehicle.ID)

	assert.ErrorIs(t, err, adapters.ErrPersisting)
}

func TestDeleteVehicleSuccess(t *testing.T) {
	ctx := context.Background()
	vehicle := dsl.NewValidVehicleOne()

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("Get", ctx, vehicle.ID).Return(&vehicle, nil)
	vehicleAdapter.On("Delete", ctx, vehicle.ID).Return(nil)

	usecase := usecases.NewDeleteVehicle(vehicleAdapter)
	err := usecase.Execute(ctx, vehicle.ID)

	assert.Nil(t, err)
}
