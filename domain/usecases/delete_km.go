package usecases

import (
	"context"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
)

type DeleteKm struct {
	kmAdapter      adapters.Km
	vehicleAdapter adapters.Vehicle
}

func NewDeleteKm(kmAdapter adapters.Km, vehicleAdapter adapters.Vehicle) DeleteKm {
	return DeleteKm{
		kmAdapter:      kmAdapter,
		vehicleAdapter: vehicleAdapter,
	}
}

func (u DeleteKm) Execute(ctx context.Context, vehicleID entities.VehicleID, kmID entities.KmID) error {
	vehicle, err := u.vehicleAdapter.Get(ctx, vehicleID)
	if err != nil {
		return err
	}

	km, err := u.kmAdapter.Get(ctx, vehicle.ID, kmID)
	if err != nil {
		return err
	}

	return u.kmAdapter.Delete(ctx, vehicle.ID, km.ID)
}
