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

func TestDeleteKmFailGettingVehicle(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	vehicle := dsl.NewValidVehicleOne()
	km := dsl.NewValidKmOne()

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("Get", ctx, vehicle.ID).Return(nil, adapters.ErrGetting)
	kmAdapter := mocks.NewKm(t)

	usecase := usecases.NewDeleteKm(kmAdapter, vehicleAdapter)
	err := usecase.Execute(ctx, vehicle.ID, km.ID)

	assert.ErrorIs(err, adapters.ErrGetting)
}

func TestDeleteKmFailGettingKm(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	vehicle := dsl.NewValidVehicleOne()
	km := dsl.NewValidKmOne()

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("Get", ctx, vehicle.ID).Return(&vehicle, nil)
	kmAdapter := mocks.NewKm(t)
	kmAdapter.On("Get", ctx, vehicle.ID, km.ID).Return(nil, adapters.ErrGetting)

	usecase := usecases.NewDeleteKm(kmAdapter, vehicleAdapter)
	err := usecase.Execute(ctx, vehicle.ID, km.ID)

	assert.ErrorIs(err, adapters.ErrGetting)
}

func TestDeleteKmFailKmNotFound(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	vehicle := dsl.NewValidVehicleOne()
	km := dsl.NewValidKmOne()

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("Get", ctx, vehicle.ID).Return(&vehicle, nil)
	kmAdapter := mocks.NewKm(t)
	kmAdapter.On("Get", ctx, vehicle.ID, km.ID).Return(nil, adapters.ErrNotFound)

	usecase := usecases.NewDeleteKm(kmAdapter, vehicleAdapter)
	err := usecase.Execute(ctx, vehicle.ID, km.ID)

	assert.ErrorIs(err, adapters.ErrNotFound)
}

func TestDeleteKmFailDeleting(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	vehicle := dsl.NewValidVehicleOne()
	km := dsl.NewValidKmOne()

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("Get", ctx, vehicle.ID).Return(&vehicle, nil)
	kmAdapter := mocks.NewKm(t)
	kmAdapter.On("Get", ctx, vehicle.ID, km.ID).Return(&km, nil)
	kmAdapter.On("Delete", ctx, vehicle.ID, km.ID).Return(adapters.ErrPersisting)

	usecase := usecases.NewDeleteKm(kmAdapter, vehicleAdapter)
	err := usecase.Execute(ctx, vehicle.ID, km.ID)

	assert.ErrorIs(err, adapters.ErrPersisting)
}

func TestDeleteKmSuccess(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	vehicle := dsl.NewValidVehicleOne()
	km := dsl.NewValidKmOne()

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("Get", ctx, vehicle.ID).Return(&vehicle, nil)
	kmAdapter := mocks.NewKm(t)
	kmAdapter.On("Get", ctx, vehicle.ID, km.ID).Return(&km, nil)
	kmAdapter.On("Delete", ctx, vehicle.ID, km.ID).Return(nil)

	usecase := usecases.NewDeleteKm(kmAdapter, vehicleAdapter)
	err := usecase.Execute(ctx, vehicle.ID, km.ID)

	assert.Nil(err)
}
