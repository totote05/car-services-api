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

func TestUpdateVehicleWithoutPlateShouldFail(t *testing.T) {
	ctx := context.Background()
	vehicle := entities.Vehicle{}

	usecase := usecases.NewUpdateVehicle(nil)
	_, err := usecase.Execute(ctx, vehicle)

	assert.ErrorIs(t, err, usecases.ErrInvalidVehicleData)
}

func TestUpdateVehicleAdapterFailOnGet(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	vehicle := dsl.NewValidVehicleOne()

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("Get", ctx, vehicle.ID).Return(nil, adapters.ErrGetting).Once()

	usecase := usecases.NewUpdateVehicle(vehicleAdapter)
	updatedVehicle, err := usecase.Execute(ctx, vehicle)

	assert.Nil(updatedVehicle)
	assert.ErrorIs(err, adapters.ErrGetting)
}

func TestUpdateVehicleNotFound(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	vehicle := dsl.NewValidVehicleOne()

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("Get", ctx, vehicle.ID).Return(nil, adapters.ErrNotFound)

	usecase := usecases.NewUpdateVehicle(vehicleAdapter)
	updatedVehicle, err := usecase.Execute(ctx, vehicle)

	assert.Nil(updatedVehicle)
	assert.ErrorIs(err, adapters.ErrNotFound)
}

func TestUpdateVehicleFailFindByPlate(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	vehicle := dsl.NewValidVehicleOne()

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("Get", ctx, vehicle.ID).Return(&vehicle, nil)
	vehicleAdapter.On("FindByPlate", ctx, vehicle.Plate).Return(nil, adapters.ErrGetting)

	usecase := usecases.NewUpdateVehicle(vehicleAdapter)
	updatedVehicle, err := usecase.Execute(ctx, vehicle)

	assert.Nil(updatedVehicle)
	assert.ErrorIs(err, adapters.ErrGetting)
}

func TestUpdateVehicleFailByDuplicatedPlate(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	vehicle := dsl.NewValidVehicleOne()
	toUpdate := dsl.UpdateValidVehicle(vehicle)
	vehicleTwo := dsl.UpdateValidVehicle(dsl.NewValidVehicleTwo())

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("Get", ctx, vehicle.ID).Return(&vehicle, nil)
	vehicleAdapter.On("FindByPlate", ctx, toUpdate.Plate).Return(&vehicleTwo, nil)

	usecase := usecases.NewUpdateVehicle(vehicleAdapter)
	updateVehicle, err := usecase.Execute(ctx, toUpdate)

	assert.Nil(updateVehicle)
	assert.ErrorIs(err, usecases.ErrDuplicatedVehicle)
}

func TestUpdateVehicleAdapterFailOnSave(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	vehicle := dsl.NewValidVehicleOne()
	toUpdate := dsl.UpdateValidVehicle(dsl.NewValidVehicleOne())

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("Get", ctx, vehicle.ID).Return(&vehicle, nil).Once()
	vehicleAdapter.On("FindByPlate", ctx, toUpdate.Plate).Return(nil, adapters.ErrNotFound).Once()
	vehicleAdapter.On("Save", ctx, toUpdate).Return(adapters.ErrPersisting)

	usecase := usecases.NewUpdateVehicle(vehicleAdapter)
	updatedVehicle, err := usecase.Execute(ctx, toUpdate)

	assert.Nil(updatedVehicle)
	assert.ErrorIs(err, adapters.ErrPersisting)
}

func TestUpdateVehicleSuccess(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	vehicle := dsl.NewValidVehicleOne()
	toUpdate := dsl.UpdateValidVehicle(dsl.NewValidVehicleOne())

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("Get", ctx, vehicle.ID).Return(&vehicle, nil).Once()
	vehicleAdapter.On("FindByPlate", ctx, toUpdate.Plate).Return(nil, adapters.ErrNotFound).Once()
	vehicleAdapter.On("Save", ctx, toUpdate).Return(nil)

	usecase := usecases.NewUpdateVehicle(vehicleAdapter)
	updatedVehicle, err := usecase.Execute(ctx, toUpdate)

	assert.Nil(err)
	assert.Equal(&toUpdate, updatedVehicle)
}
