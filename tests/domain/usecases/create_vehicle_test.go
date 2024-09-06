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

func TestNewVehicleWithoutPlateShouldFail(t *testing.T) {
	ctx := context.Background()
	vehicle := dsl.NewInvalidVehicle()

	usecase := usecases.NewCreateVehicle(nil)
	_, err := usecase.Execute(ctx, vehicle)

	assert.ErrorIs(t, err, usecases.ErrInvalidVehicleData)
}

func TestNewVehicleWithDuplicatedPlateShoulFail(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	vehicle := dsl.NewValidCreateVehicle()

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("FindByPlate", ctx, vehicle.Plate).Return(nil, adapters.ErrNotFound).Once()
	vehicleAdapter.On("Save", ctx, dsl.AnythingOfType(vehicle)).Return(nil).Once()
	vehicleAdapter.On("FindByPlate", ctx, vehicle.Plate).Return(&vehicle, nil).Once()

	usecase := usecases.NewCreateVehicle(vehicleAdapter)

	_, err := usecase.Execute(ctx, vehicle)
	assert.Nil(err)

	_, err = usecase.Execute(ctx, vehicle)
	assert.ErrorIs(err, usecases.ErrDuplicatedVehicle)
}

func TestNewVehicleAdapterFailOnFindByPlate(t *testing.T) {
	ctx := context.Background()
	vehicle := dsl.NewValidCreateVehicle()

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("FindByPlate", ctx, vehicle.Plate).Return(nil, adapters.ErrGetting).Once()

	usecase := usecases.NewCreateVehicle(vehicleAdapter)
	_, err := usecase.Execute(ctx, vehicle)

	assert.ErrorIs(t, err, adapters.ErrGetting)
}

func TestNewVehicleAdapterFailOnSave(t *testing.T) {
	ctx := context.Background()
	vehicle := dsl.NewValidCreateVehicle()

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("FindByPlate", ctx, vehicle.Plate).Return(nil, adapters.ErrNotFound).Once()
	vehicleAdapter.On("Save", ctx, dsl.AnythingOfType(vehicle)).Return(adapters.ErrPersisting)

	usecase := usecases.NewCreateVehicle(vehicleAdapter)
	_, err := usecase.Execute(ctx, vehicle)

	assert.ErrorIs(t, err, adapters.ErrPersisting)
}

func TestNewVehicleShouldReturnNotEmptyID(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	vehicle := dsl.NewValidCreateVehicle()

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("FindByPlate", ctx, vehicle.Plate).Return(&vehicle, adapters.ErrNotFound).Once()
	vehicleAdapter.On("Save", ctx, dsl.AnythingOfType(vehicle)).Return(nil)

	usecase := usecases.NewCreateVehicle(vehicleAdapter)
	createdVehicle, err := usecase.Execute(ctx, vehicle)

	assert.Nil(err)
	assert.NotEmpty(createdVehicle.ID)
}
