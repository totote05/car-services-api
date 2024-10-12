package adapters

import (
	"context"

	"car-services-api.totote05.ar/domain/entities"
)

type Km interface {
	Save(ctx context.Context, vehicleID entities.VehicleID, km entities.Km) error
	GetAll(ctx context.Context, vehicleID entities.VehicleID) (entities.KmList, error)
	Get(ctx context.Context, vehicleID entities.VehicleID, kmID entities.KmID) (*entities.Km, error)
	Delete(ctx context.Context, vehicleID entities.VehicleID, kmID entities.KmID) error
}
