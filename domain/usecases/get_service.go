package usecases

import (
	"context"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
)

type (
	GetService struct {
		serviceAdapter adapters.Service
	}
)

func NewGetService(serviceAdapter adapters.Service) GetService {
	return GetService{
		serviceAdapter: serviceAdapter,
	}
}

func (u GetService) Execute(ctx context.Context, serviceID entities.ServiceID) (*entities.Service, error) {
	return u.serviceAdapter.Get(ctx, serviceID)
}
