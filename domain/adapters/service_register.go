package adapters

import (
	"context"

	"car-services-api.totote05.ar/domain/entities"
)

type ServiceRegister interface {
	Save(ctx context.Context, serviceRegister entities.ServiceRegister) error
	GetAll(ctx context.Context, vehicleID entities.VehicleID) ([]entities.ServiceRegister, error)
	Get(ctx context.Context, vehicleID entities.VehicleID, serviceRegisterID entities.ServiceRegisterID) (*entities.ServiceRegister, error)
}
