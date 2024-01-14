package adapters

import (
	"context"

	"car-services-api.totote05.ar/domain/entities"
)

type (
	Service interface {
		GetAll(ctx context.Context) ([]entities.Service, error)
		Get(ctx context.Context, ID entities.ServiceID) (*entities.Service, error)
		Save(ctx context.Context, service entities.Service) error
		FindByName(ctx context.Context, Name string) ([]entities.Service, error)
	}
)
