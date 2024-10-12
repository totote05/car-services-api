package usecases

import (
	"context"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
)

type GetServiceRegister struct {
	serviceRegisterAdapter adapters.ServiceRegister
	vehicleAdapter         adapters.Vehicle
}

func NewGetServiceRegister(serviceRegisterAdapter adapters.ServiceRegister, vehicleAdapter adapters.Vehicle) *GetServiceRegister {
	return &GetServiceRegister{
		serviceRegisterAdapter: serviceRegisterAdapter,
		vehicleAdapter:         vehicleAdapter,
	}
}

func (g *GetServiceRegister) Execute(
	ctx context.Context,
	vehicleID entities.VehicleID,
	serviceRegisterID entities.ServiceRegisterID,
) (*entities.ServiceRegister, error) {
	vehicle, err := g.vehicleAdapter.Get(ctx, vehicleID)
	if err != nil {
		return nil, err
	}

	return g.serviceRegisterAdapter.Get(ctx, vehicle.ID, serviceRegisterID)
}
