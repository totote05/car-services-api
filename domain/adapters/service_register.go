package adapters

import (
	"context"

	"car-services-api.totote05.ar/domain/entities"
)

type ServiceRegister interface {
	Save(ctx context.Context, serviceRegister entities.ServiceRegister) error
}
