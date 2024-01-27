package usecases_test

import (
	"context"
	"testing"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
	"car-services-api.totote05.ar/domain/usecases"
	"car-services-api.totote05.ar/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateServiceShouldFailByEmptyName(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	service := entities.Service{}

	usecase := usecases.NewCreateService(nil)
	result, err := usecase.Execute(ctx, service)

	assert.Nil(result)
	assert.ErrorContains(err, "empty name")
}

func TestCreateServiceShouldFailByServiceError(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	serviceAdapter := mocks.NewService(t)
	serviceAdapter.On("Save", ctx, mock.Anything).Return(adapters.ErrPersisting)

	service := entities.Service{
		Name: "dummy service",
	}

	usecase := usecases.NewCreateService(serviceAdapter)
	result, err := usecase.Execute(ctx, service)

	assert.Nil(result)
	assert.ErrorIs(err, adapters.ErrPersisting)
}

func TestCreateServiceSuccess(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	service := entities.Service{
		Name: "dummy service",
	}

	serviceAdapter := mocks.NewService(t)
	serviceAdapter.On("Save", ctx, mock.AnythingOfType("entities.Service")).Return(nil)

	usecase := usecases.NewCreateService(serviceAdapter)
	result, err := usecase.Execute(ctx, service)

	assert.Nil(err)
	assert.NotEmpty(result.ID)
}
