package dsl

import "car-services-api.totote05.ar/domain/entities"

func NewValidServiceOne() entities.Service {
	return entities.Service{
		ID:   "123",
		Name: "Service one",
	}
}

func NewValidServiceTwo() entities.Service {
	return entities.Service{
		ID:   "456",
		Name: "Service Two",
	}
}

func NewValidServiceList() []entities.Service {
	return []entities.Service{
		NewValidServiceOne(),
		NewValidServiceTwo(),
	}
}
