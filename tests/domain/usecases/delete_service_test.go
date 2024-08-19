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

func TestDeleteServiceFailGettingService(t *testing.T) {
	ctx := context.Background()
	service := dsl.NewValidServiceOne()

	serviceAdapter := mocks.NewService(t)
	serviceAdapter.On("Get", ctx, service.ID).Return(nil, adapters.ErrGetting)

	usecase := usecases.NewDeleteService(serviceAdapter)
	err := usecase.Execute(ctx, service.ID)

	assert.ErrorIs(t, err, adapters.ErrGetting)
}

func TestDeleteServiceFailServiceNotFound(t *testing.T) {
	ctx := context.Background()
	service := dsl.NewValidServiceOne()

	serviceAdapter := mocks.NewService(t)
	serviceAdapter.On("Get", ctx, service.ID).Return(nil, adapters.ErrNotFound)

	usecase := usecases.NewDeleteService(serviceAdapter)
	err := usecase.Execute(ctx, service.ID)

	assert.ErrorIs(t, err, adapters.ErrNotFound)
}

func TestDeleteServiceFailDeleting(t *testing.T) {
	ctx := context.Background()
	service := dsl.NewValidServiceOne()

	serviceAdapter := mocks.NewService(t)
	serviceAdapter.On("Get", ctx, service.ID).Return(service, nil)
	serviceAdapter.On("Delete", ctx, service.ID).Return(adapters.ErrPersisting)

	usecase := usecases.NewDeleteService(serviceAdapter)
	err := usecase.Execute(ctx, service.ID)

	assert.ErrorIs(t, err, adapters.ErrPersisting)
}

func TestDeleteServiceSuccess(t *testing.T) {
	ctx := context.Background()
	service := dsl.NewValidServiceOne()

	serviceAdapter := mocks.NewService(t)
	serviceAdapter.On("Get", ctx, service.ID).Return(service, nil)
	serviceAdapter.On("Delete", ctx, service.ID).Return(nil)

	usecase := usecases.NewDeleteService(serviceAdapter)
	err := usecase.Execute(ctx, service.ID)

	assert.NoError(t, err)
}
