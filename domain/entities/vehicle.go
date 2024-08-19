package entities

import "errors"

var (
	ErrVehicleHasEmptyPlate = errors.New("empty plate")
)

type (
	Vehicle struct {
		ID    VehicleID `json:"id"`
		Plate string    `json:"plate"`
	}
	VehicleID string
)

func (v *Vehicle) Validate() error {
	if v != nil && v.Plate == "" {
		return ErrVehicleHasEmptyPlate
	}

	return nil
}
