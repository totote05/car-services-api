package dsl

import "car-services-api.totote05.ar/domain/entities"

func NewValidServiceRegister() entities.ServiceRegister {
	return entities.ServiceRegister{
		ID:        "1",
		VehicleID: NewValidVehicleOne().ID,
		ServiceID: NewValidServiceOne().ID,
		Km:        NewValidKmOne(),
	}
}
