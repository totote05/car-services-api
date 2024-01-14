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

func TestGetServicesSuccess(t *testing.T) {
	ctx := context.Background()

	expected := dsl.NewValidServiceList()
	serviceAdapter := mocks.NewService(t)
	serviceAdapter.On("GetAll", ctx).Return(expected, nil)

	usecase := usecases.NewGetServices(serviceAdapter)

	list, err := usecase.Execute(ctx)

	assert.Nil(t, err)
	assert.Equal(t, expected, list)
}

func TestGetServicesFailAdapter(t *testing.T) {
	ctx := context.Background()

	serviceAdapter := mocks.NewService(t)
	serviceAdapter.On("GetAll", ctx).Return(nil, adapters.ErrGetting)

	usecase := usecases.NewGetServices(serviceAdapter)
	list, err := usecase.Execute(ctx)

	assert.ErrorIs(t, err, adapters.ErrGetting)
	assert.Nil(t, list)
}
