package dsl

import (
	"time"

	"car-services-api.totote05.ar/domain/entities"
)

var initialDate = time.Now()

func NewValidKmOne() entities.Km {
	return entities.Km{
		Value: 1000,
		Date:  initialDate,
	}
}

func NewValidKmTwo() entities.Km {
	return entities.Km{
		Value: 1050,
		Date:  initialDate.Add(time.Hour),
	}
}

func NewInvalidKm() entities.Km {
	return entities.Km{
		Value: 1001,
		Date:  initialDate.Add(-time.Hour),
	}
}

func NewInvalidKmTwo() entities.Km {
	return entities.Km{
		Value: 1060,
		Date:  initialDate.Add(time.Minute * 30),
	}
}
