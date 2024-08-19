package dsl

import (
	"time"

	"car-services-api.totote05.ar/domain/entities"
)

func NewInvalidService() *entities.Service {
	return &entities.Service{}
}

func NewValidServiceWithInvalidRecurrence() *entities.Service {
	return &entities.Service{
		Name:        "dummy",
		Rercurrence: &entities.ServiceRecurrence{},
	}
}

func NewValidService() *entities.Service {
	return &entities.Service{
		Name: "dummy",
	}
}

func NewValidServiceOne() *entities.Service {
	var (
		kilometers uint32        = 10_000
		interval   time.Duration = time.Hour * 24 * 365
	)
	return &entities.Service{
		ID:   "123",
		Name: "Service one",
		Rercurrence: &entities.ServiceRecurrence{
			Kilometers: &kilometers,
			Interval:   &interval,
		},
	}
}

func NewValidServiceTwo() *entities.Service {
	var kilometers uint32 = 60_000
	return &entities.Service{
		ID:   "456",
		Name: "Service Two",
		Rercurrence: &entities.ServiceRecurrence{
			Kilometers: &kilometers,
		},
	}
}

func NewValidServiceList() []entities.Service {
	return []entities.Service{
		*NewValidServiceOne(),
		*NewValidServiceTwo(),
	}
}
