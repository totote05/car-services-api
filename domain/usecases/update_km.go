package usecases

import (
	"context"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
)

type UpdateKm struct {
	kmAdapter      adapters.Km
	vehicleAdapter adapters.Vehicle
}

func NewUpdateKm(kmAdapter adapters.Km, vehicleAdapter adapters.Vehicle) UpdateKm {
	return UpdateKm{
		kmAdapter:      kmAdapter,
		vehicleAdapter: vehicleAdapter,
	}
}

func (u UpdateKm) Execute(
	ctx context.Context,
	vehicleID entities.VehicleID,
	kmID entities.KmID,
	km entities.Km,
) (*entities.Km, error) {
	if err := km.Validate(); err != nil {
		return nil, err
	}

	vehicle, err := u.vehicleAdapter.Get(ctx, vehicleID)
	if err != nil {
		return nil, err
	}

	list, err := u.kmAdapter.GetAll(ctx, vehicle.ID)
	if err != nil {
		return nil, err
	}

	var found bool
	for _, k := range list {
		isDifferentKm := k.ID != kmID
		isSameDate := k.Date.Equal(km.Date)
		isSameValue := k.Value == km.Value
		isAfter := km.Date.After(k.Date) && km.Value < k.Value
		isBefore := km.Date.Before(k.Date) && km.Value > k.Value

		if isDifferentKm && (isSameDate || isSameValue || isAfter || isBefore) {
			return nil, ErrInvalidKmData
		}
		if k.ID == kmID {
			found = true
		}
	}

	if !found {
		return nil, ErrKmNotFound
	}

	if err := u.kmAdapter.Save(ctx, vehicle.ID, km); err != nil {
		return nil, err
	}

	return &km, nil
}
