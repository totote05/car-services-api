package usecases

import (
	"context"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
	"github.com/aidarkhanov/nanoid"
)

type RegisterService struct {
	serviceRegisterAdapter adapters.ServiceRegister
	vehicleAdapter         adapters.Vehicle
	serviceAdapter         adapters.Service
	kmAdapter              adapters.Km
}

func NewRegisterService(
	serviceRegisterAdapter adapters.ServiceRegister,
	vehicleAdapter adapters.Vehicle,
	serviceAdapter adapters.Service,
	kmAdapter adapters.Km,
) RegisterService {
	return RegisterService{
		serviceRegisterAdapter: serviceRegisterAdapter,
		vehicleAdapter:         vehicleAdapter,
		serviceAdapter:         serviceAdapter,
		kmAdapter:              kmAdapter,
	}
}

func (r RegisterService) Execute(
	ctx context.Context,
	vehicleID entities.VehicleID,
	serviceID entities.ServiceID,
	km entities.Km,
) (*entities.ServiceRegister, error) {
	vehicle, err := r.vehicleAdapter.Get(ctx, vehicleID)
	if err != nil {
		return nil, err
	}

	service, err := r.serviceAdapter.Get(ctx, serviceID)
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

	if km.ID == "" {
		km.ID = entities.KmID(nanoid.New())
	}

	if err := r.kmAdapter.Save(ctx, vehicle.ID, km); err != nil {
		return nil, err
	}

	registerService := entities.ServiceRegister{
		ID:        entities.ServiceRegisterID(nanoid.New()),
		VehicleID: vehicle.ID,
		ServiceID: service.ID,
		Km:        km,
	}

	if err := r.serviceRegisterAdapter.Save(ctx, registerService); err != nil {
		return nil, err
	}

	return &registerService, nil
}
