package usecases

import (
	"context"
	"errors"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
	"github.com/aidarkhanov/nanoid"
)

type (
	CreateSerice struct {
		serviceAdapter adapters.Service
	}
)

func NewCreateService(serviceAdapter adapters.Service) CreateSerice {
	return CreateSerice{
		serviceAdapter: serviceAdapter,
	}
}

func (s CreateSerice) Execute(ctx context.Context, service entities.Service) (*entities.Service, error) {
	if err := service.Validate(); err != nil {
		return nil, ErrInvalidServiceData
	}

	founded, err := s.serviceAdapter.FindByName(ctx, service.Name)
	if err != nil && !errors.Is(adapters.ErrNotFound, err) {
		return nil, err
	}

	if len(founded) > 0 {
		return nil, ErrDuplicatedService
	}

	service.ID = entities.ServiceID(nanoid.New())

	if err := s.serviceAdapter.Save(ctx, service); err != nil {
		return nil, err
	}

	return &service, nil
}
