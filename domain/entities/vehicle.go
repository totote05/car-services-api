package entities

import "errors"

type (
	Vehicle struct {
		ID    VehicleID `json:"id"`
		Plate string    `json:"plate"`
	}
	VehicleID string
)

func (v Vehicle) Validate() error {
	if v.Plate == "" {
		return errors.New("empty plate")
	}

	return nil
}
