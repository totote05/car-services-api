package usecases

import (
	"context"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
)

type (
	DeleteService struct {
		serviceAdapter adapters.Service
	}
)

func NewDeleteService(serviceAdapter adapters.Service) DeleteService {
	return DeleteService{
		serviceAdapter: serviceAdapter,
	}
}

func (u DeleteService) Execute(ctx context.Context, serviceID entities.ServiceID) error {
	if _, err := u.serviceAdapter.Get(ctx, serviceID); err != nil {
		return err
	}

	return u.serviceAdapter.Delete(ctx, serviceID)
}
