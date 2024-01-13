package usecases

import (
	"context"
	"errors"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
	"github.com/aidarkhanov/nanoid"
)

type (
	CreateVehicle struct {
		vehicleAdapter adapters.Vehicle
	}
)

func NewCreateVehicle(vehicleAdapter adapters.Vehicle) CreateVehicle {
	return CreateVehicle{
		vehicleAdapter,
	}
}

func (uc CreateVehicle) Execute(ctx context.Context, vehicle entities.Vehicle) (*entities.Vehicle, error) {
	if err := vehicle.Validate(); err != nil {
		return nil, ErrInvalidVehicleData
	}

	_, err := uc.vehicleAdapter.FindByPlate(ctx, vehicle.Plate)

	if err == nil {
		return nil, ErrDuplicatedVehicle
	} else if !errors.Is(adapters.ErrNotFound, err) {
		return nil, err
	}

	vehicle.ID = entities.VehicleID(nanoid.New())

	if err := uc.vehicleAdapter.Save(ctx, vehicle); err != nil {
		return nil, err
	}

	return &vehicle, nil
}
