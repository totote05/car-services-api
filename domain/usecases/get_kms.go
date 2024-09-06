package usecases

import (
	"context"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
)

type GetKms struct {
	kmAdapter      adapters.Km
	vehicleAdapter adapters.Vehicle
}

func NewGetKms(kmAdapter adapters.Km, vehicleAdapter adapters.Vehicle) GetKms {
	return GetKms{
		kmAdapter:      kmAdapter,
		vehicleAdapter: vehicleAdapter,
	}
}

func (r GetKms) Execute(ctx context.Context, vehicleID entities.VehicleID) ([]entities.Km, error) {
	vehicle, err := r.vehicleAdapter.Get(ctx, vehicleID)
	if err != nil {
		return nil, err
	}

	return r.kmAdapter.GetAll(ctx, vehicle.ID)
}
