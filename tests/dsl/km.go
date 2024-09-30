package dsl

import (
	"time"

	"car-services-api.totote05.ar/domain/entities"
)

var initialDate = time.Now()

func UpdateWith(km entities.Km, value uint64, time time.Duration) entities.Km {
	return entities.Km{
		ID:    km.ID,
		Value: value,
		Date:  km.Date.Add(time),
	}
}

func NewValidKm() entities.Km {
	return entities.Km{
		Value: 1000,
		Date:  initialDate,
	}
}

func NewValidKmOne() entities.Km {
	return entities.Km{
		ID:    entities.KmID("1"),
		Value: 1000,
		Date:  initialDate,
	}
}

func NewValidKmTwo() entities.Km {
	return entities.Km{
		ID:    entities.KmID("2"),
		Value: 1050,
		Date:  initialDate.Add(time.Hour),
	}
}

func NewInvalidKm() entities.Km {
	return entities.Km{
		ID:    entities.KmID("3"),
		Value: 1001,
		Date:  initialDate.Add(-time.Hour),
	}
}

func NewInvalidKmTwo() entities.Km {
	return entities.Km{
		ID:    entities.KmID("4"),
		Value: 1060,
		Date:  initialDate.Add(time.Minute * 30),
	}
}

func NewInvalidKmThree() entities.Km {
	return entities.Km{
		ID:    entities.KmID("5"),
		Value: 1000,
		Date:  initialDate.Add(time.Minute * 30),
	}
}

func NewInvalidKmFour() entities.Km {
	return entities.Km{
		ID:    entities.KmID("6"),
		Value: 1060,
		Date:  initialDate,
	}
}

func NewInvalidKmFive() entities.Km {
	return entities.Km{
		Date: initialDate.Add(time.Hour),
	}
}
