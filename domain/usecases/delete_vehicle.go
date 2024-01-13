package usecases

import (
	"context"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
)

type (
	DeleteVehicle struct {
		vehicleAdapter adapters.Vehicle
	}
)

func NewDeleteVehicle(vehicleAdapter adapters.Vehicle) DeleteVehicle {
	return DeleteVehicle{
		vehicleAdapter: vehicleAdapter,
	}
}

func (uc DeleteVehicle) Execute(ctx context.Context, vehicleID entities.VehicleID) error {
	if _, err := uc.vehicleAdapter.Get(ctx, vehicleID); err != nil {
		return err
	}

	return uc.vehicleAdapter.Delete(ctx, vehicleID)
}
