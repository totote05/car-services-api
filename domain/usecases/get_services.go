package usecases

import (
	"context"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
)

type (
	GetServices interface {
		Execute(ctx context.Context) ([]entities.Service, error)
	}
	getServices struct {
		serviceAdapter adapters.Service
	}
)

func NewGetServices(serviceAdapter adapters.Service) GetServices {
	return getServices{
		serviceAdapter: serviceAdapter,
	}
}

func (uc getServices) Execute(ctx context.Context) ([]entities.Service, error) {
	list, err := uc.serviceAdapter.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}
