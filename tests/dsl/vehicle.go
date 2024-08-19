package dsl

import "car-services-api.totote05.ar/domain/entities"

func NewInvalidVehicle() *entities.Vehicle {
	return &entities.Vehicle{}
}

func NewValidCreateVehicle() *entities.Vehicle {
	return &entities.Vehicle{
		Plate: "ABC 123",
	}
}

func NewValidVehicleOne() *entities.Vehicle {
	return &entities.Vehicle{
		ID:    "123",
		Plate: "ABC 123",
	}
}

func NewValidVehicleTwo() *entities.Vehicle {
	return &entities.Vehicle{
		ID:    "456",
		Plate: "DEF 456",
	}
}

func UpdateValidVehicle(vehicle entities.Vehicle) entities.Vehicle {
	vehicle.Plate = "GHI 789"
	return vehicle
}

func NewValidVehicleList() []entities.Vehicle {
	return []entities.Vehicle{
		*NewValidCreateVehicle(),
		*NewValidVehicleTwo(),
	}
}
