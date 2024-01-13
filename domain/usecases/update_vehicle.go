package usecases

import (
	"context"
	"errors"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
)

type (
	UpdateVehicle struct {
		vehicleAdapter adapters.Vehicle
	}
)

func NewUpdateVehicle(vehicleAdapter adapters.Vehicle) UpdateVehicle {
	return UpdateVehicle{
		vehicleAdapter: vehicleAdapter,
	}
}

func (uc UpdateVehicle) Execute(ctx context.Context, vehicle entities.Vehicle) (*entities.Vehicle, error) {
	if err := vehicle.Validate(); err != nil {
		return nil, ErrInvalidVehicleData
	}

	if _, err := uc.vehicleAdapter.Get(ctx, vehicle.ID); err != nil {
		return nil, err
	}

	findedVehicle, err := uc.vehicleAdapter.FindByPlate(ctx, vehicle.Plate)
	if err != nil && !errors.Is(err, adapters.ErrNotFound) {
		return nil, err
	}

	if findedVehicle != nil && findedVehicle.ID != vehicle.ID {
		return nil, ErrDuplicatedVehicle
	}

	if err := uc.vehicleAdapter.Save(ctx, vehicle); err != nil {
		return nil, err
	}

	return &vehicle, nil
}
