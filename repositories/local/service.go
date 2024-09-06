package local

import (
	"context"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
)

type (
	service struct {
		storage map[entities.ServiceID]entities.Service
	}
)

func NewService() adapters.Service {
	ld := getLocalData()
	storage := map[entities.ServiceID]entities.Service{}

	for _, service := range ld.Services {
		storage[service.ID] = service
	}

	return &service{
		storage: storage,
	}
}

func (r *service) FindByName(ctx context.Context, Name string) ([]entities.Service, error) {
	for _, item := range r.storage {
		if item.Name == Name {
			return []entities.Service{item}, nil
		}
	}

	return nil, adapters.ErrNotFound
}

func (r *service) Get(ctx context.Context, ID entities.ServiceID) (*entities.Service, error) {
	if service, ok := r.storage[ID]; ok {
		return &service, nil
	}

	return nil, adapters.ErrNotFound
}

func (r *service) GetAll(ctx context.Context) ([]entities.Service, error) {
	list := []entities.Service{}
	for _, item := range r.storage {
		list = append(list, item)
	}

	return list, nil
}

func (r *service) Save(ctx context.Context, service entities.Service) error {
	r.storage[service.ID] = service
	return nil
}

func (r *service) Delete(ctx context.Context, ID entities.ServiceID) error {
	delete(r.storage, ID)
	return nil
}
