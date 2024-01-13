package usecases

import (
	"context"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
)

type (
	GetVehicles struct {
		vehicleAdapter adapters.Vehicle
	}
)

func NewGetVehicles(vehicleAdapter adapters.Vehicle) GetVehicles {
	return GetVehicles{
		vehicleAdapter: vehicleAdapter,
	}
}

func (uc GetVehicles) Execute(ctx context.Context) ([]entities.Vehicle, error) {
	return uc.vehicleAdapter.GetAll(ctx)
}
