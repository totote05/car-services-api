package usecases

import (
	"context"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
)

type GetKm struct {
	kmAdapter      adapters.Km
	vehicleAdapter adapters.Vehicle
}

func NewGetKm(kmAdapter adapters.Km, vehicleAdapter adapters.Vehicle) GetKm {
	return GetKm{
		kmAdapter:      kmAdapter,
		vehicleAdapter: vehicleAdapter,
	}
}

func (r GetKm) Execute(ctx context.Context, vehicleID entities.VehicleID, kmID entities.KmID) (*entities.Km, error) {
	vehicle, err := r.vehicleAdapter.Get(ctx, vehicleID)
	if err != nil {
		return nil, err
	}

	km, err := r.kmAdapter.Get(ctx, vehicle.ID, kmID)
	if err != nil {
		return nil, err
	}

	return km, nil
}
