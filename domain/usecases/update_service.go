package usecases

import (
	"context"
	"errors"

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
		return nil, ErrInvalidServiceData
	}

	founded, err := u.serviceAdapter.FindByName(ctx, service.Name)
	if err != nil && !errors.Is(err, adapters.ErrNotFound) {
		return nil, err
	}

	if len(founded) > 0 && founded[0].ID != service.ID {
		return nil, ErrDuplicatedService
	}

	if _, err := u.serviceAdapter.Get(ctx, service.ID); err != nil {
		return nil, err
	}

	if err := u.serviceAdapter.Save(ctx, service); err != nil {
		return nil, err
	}

	return &service, nil
}
