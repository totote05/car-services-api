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

func TestGetVehiclesAdapterFail(t *testing.T) {
	ctx := context.Background()

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("GetAll", ctx).Return(nil, adapters.ErrGetting)

	usecase := usecases.NewGetVehicles(vehicleAdapter)
	_, err := usecase.Execute(ctx)

	assert.ErrorIs(t, err, adapters.ErrGetting)
}

func TestGetVehiclesSuccess(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	vehicles := dsl.NewValidVehicleList()

	vehicleAdapter := mocks.NewVehicle(t)
	vehicleAdapter.On("GetAll", ctx).Return(vehicles, nil)

	usecase := usecases.NewGetVehicles(vehicleAdapter)
	list, err := usecase.Execute(ctx)

	assert.Nil(err)
	assert.Equal(vehicles, list)
}
