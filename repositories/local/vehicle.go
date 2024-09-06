package local

import (
	"context"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
)

type (
	vehicle struct {
		storage map[entities.VehicleID]entities.Vehicle
	}
)

func NewVehicle() adapters.Vehicle {
	ld := getLocalData()
	storage := map[entities.VehicleID]entities.Vehicle{}

	for _, vehicle := range ld.Vehicle {
		storage[vehicle.ID] = vehicle.Vehicle
	}

	return &vehicle{
		storage: storage,
	}
}

func (r *vehicle) Delete(ctx context.Context, ID entities.VehicleID) error {
	delete(r.storage, ID)
	return nil
}

func (r *vehicle) Get(ctx context.Context, ID entities.VehicleID) (*entities.Vehicle, error) {
	if value, ok := r.storage[ID]; ok {
		return &value, nil
	}
	return nil, adapters.ErrNotFound
}

func (r *vehicle) GetAll(ctx context.Context) ([]entities.Vehicle, error) {
	vehicles := []entities.Vehicle{}
	for _, vehicle := range r.storage {
		vehicles = append(vehicles, vehicle)
	}
	return vehicles, nil
}

func (r *vehicle) Save(ctx context.Context, vehicle entities.Vehicle) error {
	r.storage[vehicle.ID] = vehicle
	return nil
}

func (r *vehicle) FindByPlate(ctx context.Context, plate string) (*entities.Vehicle, error) {
	for _, vehicle := range r.storage {
		if vehicle.Plate == plate {
			return &vehicle, nil
		}
	}
	return nil, adapters.ErrNotFound
}
