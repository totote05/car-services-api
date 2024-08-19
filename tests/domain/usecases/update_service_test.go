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

func TestUpdateServiceWithoutNameShouldFail(t *testing.T) {
	ctx := context.Background()
	service := dsl.NewInvalidService()

	usecase := usecases.NewUpdateService(nil)
	_, err := usecase.Execute(ctx, *service)

	assert.ErrorIs(t, err, usecases.ErrInvalidServiceData)
}

func TestUpdateServiceAdapterFailOnGet(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	service := dsl.NewValidServiceOne()

	serviceAdapter := mocks.NewService(t)
	serviceAdapter.On("FindByName", ctx, service.Name).Return(nil, adapters.ErrNotFound).Once()
	serviceAdapter.On("Get", ctx, service.ID).Return(nil, adapters.ErrGetting).Once()

	usecase := usecases.NewUpdateService(serviceAdapter)
	updatedService, err := usecase.Execute(ctx, *service)

	assert.Nil(updatedService)
	assert.ErrorIs(err, adapters.ErrGetting)
}

func TestUpdateServiceNotFound(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	service := dsl.NewValidServiceOne()

	serviceAdapter := mocks.NewService(t)
	serviceAdapter.On("FindByName", ctx, service.Name).Return(nil, adapters.ErrNotFound).Once()
	serviceAdapter.On("Get", ctx, service.ID).Return(nil, adapters.ErrNotFound)

	usecase := usecases.NewUpdateService(serviceAdapter)
	updatedService, err := usecase.Execute(ctx, *service)

	assert.Nil(updatedService)
	assert.ErrorIs(err, adapters.ErrNotFound)
}

func TestUpdateServiceFailOnSave(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	service := dsl.NewValidServiceOne()

	serviceAdapter := mocks.NewService(t)
	serviceAdapter.On("FindByName", ctx, service.Name).Return(nil, adapters.ErrNotFound).Once()
	serviceAdapter.On("Get", ctx, service.ID).Return(service, nil)
	serviceAdapter.On("Save", ctx, *service).Return(adapters.ErrPersisting).Once()

	usecase := usecases.NewUpdateService(serviceAdapter)
	updatedService, err := usecase.Execute(ctx, *service)

	assert.Nil(updatedService)
	assert.ErrorIs(err, adapters.ErrPersisting)
}

func TestUpdateServiceFailOnFindByName(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	service := dsl.NewValidServiceOne()

	serviceAdapter := mocks.NewService(t)
	serviceAdapter.On("FindByName", ctx, service.Name).Return(nil, adapters.ErrGetting).Once()

	usecase := usecases.NewUpdateService(serviceAdapter)
	updatedService, err := usecase.Execute(ctx, *service)

	assert.Nil(updatedService)
	assert.ErrorIs(err, adapters.ErrGetting)
}

func TestUpdateServiceSuccess(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	service := dsl.NewValidServiceOne()

	serviceAdapter := mocks.NewService(t)
	serviceAdapter.On("FindByName", ctx, service.Name).Return(nil, adapters.ErrNotFound)
	serviceAdapter.On("Get", ctx, service.ID).Return(service, nil)
	serviceAdapter.On("Save", ctx, *service).Return(nil).Once()

	usecase := usecases.NewUpdateService(serviceAdapter)
	updatedService, err := usecase.Execute(ctx, *service)

	assert.Equal(service, updatedService)
	assert.Nil(err)
}

func TestUpdateServiceFailByDuplicatedName(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	service := dsl.NewValidServiceOne()
	serviceTwo := dsl.NewValidServiceTwo()

	serviceAdapter := mocks.NewService(t)
	serviceAdapter.On("FindByName", ctx, service.Name).Return([]entities.Service{*serviceTwo}, nil)

	usecase := usecases.NewUpdateService(serviceAdapter)
	updatedService, err := usecase.Execute(ctx, *service)

	assert.Nil(updatedService)
	assert.ErrorIs(err, usecases.ErrDuplicatedService)
}
