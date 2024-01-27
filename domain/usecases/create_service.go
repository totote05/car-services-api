package usecases

import (
	"context"

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
		return nil, err
	}

	service.ID = entities.ServiceID(nanoid.New())

	err := s.serviceAdapter.Save(ctx, service)
	if err != nil {
		return nil, err
	}

	return &service, nil
}
