package usecases

import (
	"context"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
)

type (
	UpdateService struct {
		serviceAdapter adapters.Service
	}
)

func NewUpdateService(serviceAdapter adapters.Service) UpdateService {
	return UpdateService{serviceAdapter: serviceAdapter}
}

func (u UpdateService) Execute(ctx context.Context, service entities.Service) (*entities.Service, error) {
	if err := service.Validate(); err != nil {
		return nil, err
	}

	if _, err := u.serviceAdapter.Get(ctx, service.ID); err != nil {
		return nil, err
	}

	if err := u.serviceAdapter.Save(ctx, service); err != nil {
		return nil, err
	}

	return &service, nil
}
