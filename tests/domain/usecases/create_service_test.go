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

func TestCreateServiceNeverExecuteBecauseServiceIsNil(t *testing.T) {
	// coverage hack
	var service *entities.Service
	assert.Nil(t, service.Validate())
}

func TestCreateServiceShouldFailByEmptyName(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	service := dsl.NewInvalidService()

	usecase := usecases.NewCreateService(nil)
	result, err := usecase.Execute(ctx, *service)

	assert.Nil(result)
	assert.ErrorIs(err, entities.ErrServiceHasEmptyName)
}

func TestCreateServiceShouldFailByServiceError(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	serviceAdapter := mocks.NewService(t)
	serviceAdapter.On("Save", ctx, dsl.AnythingOfType(entities.Service{})).Return(adapters.ErrPersisting)

	service := dsl.NewValidService()

	usecase := usecases.NewCreateService(serviceAdapter)
	result, err := usecase.Execute(ctx, *service)

	assert.Nil(result)
	assert.ErrorIs(err, adapters.ErrPersisting)
}

func TestCreateServiceSuccess(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	service := dsl.NewValidServiceOne()

	serviceAdapter := mocks.NewService(t)
	serviceAdapter.On("Save", ctx, dsl.AnythingOfType(entities.Service{})).Return(nil)

	usecase := usecases.NewCreateService(serviceAdapter)
	result, err := usecase.Execute(ctx, *service)

	assert.Nil(err)
	assert.NotEmpty(result.ID)
}

func TestCreateServiceFailByInvalidRecurrences(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	service := dsl.NewValidServiceWithInvalidRecurrence()

	usecase := usecases.NewCreateService(nil)
	result, err := usecase.Execute(ctx, *service)

	assert.Nil(result)
	assert.ErrorIs(err, entities.ErrServiceHasEmptyRecurrence)
}
