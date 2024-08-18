package local

import (
	"context"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
)

type (
	Service struct {
		storage map[entities.ServiceID]entities.Service
	}
)

func NewService() adapters.Service {
	return &Service{
		storage: map[entities.ServiceID]entities.Service{},
	}
}

func (r *Service) FindByName(ctx context.Context, Name string) ([]entities.Service, error) {
	panic("unimplemented")
}

func (r *Service) Get(ctx context.Context, ID entities.ServiceID) (*entities.Service, error) {
	if service, ok := r.storage[ID]; ok {
		return &service, nil
	}

	return nil, adapters.ErrNotFound
}

func (r *Service) GetAll(ctx context.Context) ([]entities.Service, error) {
	list := []entities.Service{}
	for _, item := range r.storage {
		list = append(list, item)
	}

	return list, nil
}

func (r *Service) Save(ctx context.Context, service entities.Service) error {
	r.storage[service.ID] = service
	return nil
}
