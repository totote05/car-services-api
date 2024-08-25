package usecases

import (
	"context"
	"errors"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
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

func (r RegisterKm) Execute(ctx context.Context, vehicleID entities.VehicleID, km entities.Km) error {
	vehicle, err := r.vehicleAdapter.Get(ctx, vehicleID)
	if err != nil {
		return err
	}

	kms, err := r.kmAdapter.GetAll(ctx, vehicle.ID)
	if err != nil {
		return err
	}

	if len(kms) > 0 {
		for _, k := range kms {
			if (km.Date.After(k.Date) && km.Value < k.Value) ||
				(km.Date.Before(k.Date) && km.Value > k.Value) {
				return ErrInvalidKmData
			}
		}
	}

	return r.kmAdapter.Save(ctx, vehicle.ID, km)
}
