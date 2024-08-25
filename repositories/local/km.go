package local

import (
	"context"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
)

type km struct {
	storage map[entities.VehicleID][]entities.Km
}

func NewKm() adapters.Km {
	return &km{
		storage: map[entities.VehicleID][]entities.Km{},
	}
}

func (k *km) GetAll(ctx context.Context, vehicleID entities.VehicleID) ([]entities.Km, error) {
	if _, ok := k.storage[vehicleID]; !ok {
		return []entities.Km{}, nil
	}

	return k.storage[vehicleID], nil
}

func (k *km) Save(ctx context.Context, vehicleID entities.VehicleID, km entities.Km) error {
	if _, ok := k.storage[vehicleID]; !ok {
		k.storage[vehicleID] = []entities.Km{}
	}

	k.storage[vehicleID] = append(k.storage[vehicleID], km)

	return nil
}

func (k *km) Get(ctx context.Context, vehicleID entities.VehicleID, kmID entities.KmID) (*entities.Km, error) {
	if _, ok := k.storage[vehicleID]; !ok {
		return nil, adapters.ErrNotFound
	}

	for _, km := range k.storage[vehicleID] {
		if km.ID == kmID {
			return &km, nil
		}
	}

	return nil, adapters.ErrNotFound
}

func (k *km) Update(ctx context.Context, vehicleID entities.VehicleID, km entities.Km) error {
	if _, ok := k.storage[vehicleID]; !ok {
		return adapters.ErrNotFound
	}

	for i, value := range k.storage[vehicleID] {
		if value.ID == km.ID {
			k.storage[vehicleID][i] = km
			return nil
		}
	}

	return adapters.ErrNotFound
}
