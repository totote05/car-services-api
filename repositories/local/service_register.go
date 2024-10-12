package local

import (
	"context"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
)

type ServiceRegister struct {
	storage map[entities.VehicleID][]entities.ServiceRegister
}

func NewServiceRegister() adapters.ServiceRegister {
	ld := getLocalData()
	storage := map[entities.VehicleID][]entities.ServiceRegister{}

	for _, vehicle := range ld.Vehicle {
		storage[vehicle.ID] = vehicle.ServiceRegisters
	}

	return ServiceRegister{
		storage: storage,
	}
}

func (r ServiceRegister) GetAll(ctx context.Context, vehicleID entities.VehicleID) ([]entities.ServiceRegister, error) {
	if _, ok := r.storage[vehicleID]; !ok {
		return []entities.ServiceRegister{}, nil
	}

	return r.storage[vehicleID], nil
}

func (r ServiceRegister) Save(ctx context.Context, serviceRegister entities.ServiceRegister) error {
	if _, ok := r.storage[serviceRegister.VehicleID]; !ok {
		r.storage[serviceRegister.VehicleID] = []entities.ServiceRegister{}
	}

	for i, value := range r.storage[serviceRegister.VehicleID] {
		if value.ID == serviceRegister.ID {
			r.storage[serviceRegister.VehicleID][i] = serviceRegister
			return nil
		}
	}

	r.storage[serviceRegister.VehicleID] = append(r.storage[serviceRegister.VehicleID], serviceRegister)

	return nil
}

func (r ServiceRegister) Get(ctx context.Context, vehicleID entities.VehicleID, serviceRegisterID entities.ServiceRegisterID) (*entities.ServiceRegister, error) {
	if _, ok := r.storage[vehicleID]; !ok {
		return nil, adapters.ErrNotFound
	}

	for _, serviceRegister := range r.storage[vehicleID] {
		if serviceRegister.ID == serviceRegisterID {
			return &serviceRegister, nil
		}
	}

	return nil, adapters.ErrNotFound
}
