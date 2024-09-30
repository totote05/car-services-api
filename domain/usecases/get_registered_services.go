package usecases

import (
	"context"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
)

type GetRegisteredServices struct {
	serviceRegisterAdapter adapters.ServiceRegister
	vehicleAdapter         adapters.Vehicle
}

func NewGetRegisteredServices(serviceRegisterAdapter adapters.ServiceRegister, vehicleAdapter adapters.Vehicle) *GetRegisteredServices {
	return &GetRegisteredServices{
		serviceRegisterAdapter: serviceRegisterAdapter,
		vehicleAdapter:         vehicleAdapter,
	}
}

func (g *GetRegisteredServices) Execute(ctx context.Context, vehicleID entities.VehicleID) ([]entities.ServiceRegister, error) {
	_, err := g.vehicleAdapter.Get(ctx, vehicleID)
	if err != nil {
		return nil, err
	}

	return g.serviceRegisterAdapter.GetAll(ctx, vehicleID)
}
