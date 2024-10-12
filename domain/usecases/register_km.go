package usecases

import (
	"context"
	"errors"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
	"github.com/aidarkhanov/nanoid"
)

var (
	ErrInvalidKmData = errors.New("invalid km data")
)

type RegisterKm struct {
	kmAdapter      adapters.Km
	vehicleAdapter adapters.Vehicle
}

func NewRegisterKm(kmAdapter adapters.Km, vehicleAdapter adapters.Vehicle) RegisterKm {
	return RegisterKm{
		kmAdapter:      kmAdapter,
		vehicleAdapter: vehicleAdapter,
	}
}

func (r RegisterKm) Execute(ctx context.Context, vehicleID entities.VehicleID, km entities.Km) (*entities.Km, error) {
	if err := km.Validate(); err != nil {
		return nil, err
	}

	vehicle, err := r.vehicleAdapter.Get(ctx, vehicleID)
	if err != nil {
		return nil, err
	}

	kms, err := r.kmAdapter.GetAll(ctx, vehicle.ID)
	if err != nil {
		return nil, err
	}

	if kms.ValidateConsistency(km) != entities.KmValidationSuccess {
		return nil, ErrInvalidKmData
	}

	km.ID = entities.KmID(nanoid.New())
	err = r.kmAdapter.Save(ctx, vehicle.ID, km)
	if err != nil {
		return nil, err
	}

	return &km, nil
}
