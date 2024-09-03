package usecases

import "errors"

var (
	ErrInvalidVehicleData = errors.New("invalid vehicle data")
	ErrDuplicatedVehicle  = errors.New("duplicated vehicle")
	ErrInvalidServiceData = errors.New("invalid service data")
	ErrDuplicatedService  = errors.New("duplicated service")
	ErrKmNotFound         = errors.New("km not found")
)
