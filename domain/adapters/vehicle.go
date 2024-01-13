package adapters

import (
	"context"

	"car-services-api.totote05.ar/domain/entities"
)

type (
	Vehicle interface {
		Save(ctx context.Context, vehicle entities.Vehicle) error
		Delete(ctx context.Context, ID entities.VehicleID) error
		Get(ctx context.Context, ID entities.VehicleID) (*entities.Vehicle, error)
		GetAll(ctx context.Context) ([]entities.Vehicle, error)
		FindByPlate(ctc context.Context, plate string) (*entities.Vehicle, error)
	}
)
