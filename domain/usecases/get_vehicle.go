package usecases

import (
	"context"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
)

type (
	GetVehicle struct {
		vehicleAdapter adapters.Vehicle
	}
)

func NewGetVehicle(vehicleAdapter adapters.Vehicle) GetVehicle {
	return GetVehicle{
		vehicleAdapter: vehicleAdapter,
	}
}

func (uc GetVehicle) Execute(ctx context.Context, ID entities.VehicleID) (*entities.Vehicle, error) {
	return uc.vehicleAdapter.Get(ctx, ID)
}
