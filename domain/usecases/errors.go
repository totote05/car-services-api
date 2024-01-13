package usecases

import "errors"

var (
	ErrInvalidVehicleData = errors.New("invalid vehicle data")
	ErrDuplicatedVehicle  = errors.New("duplicated vehicle")
)
