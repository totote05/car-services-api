package entities

import "time"

type (
	Service struct {
		ID          ServiceID         `json:"id"`
		Name        string            `json:"name"`
		Rercurrence ServiceRecurrence `json:"rercurrence"`
	}

	ServiceRecurrence struct {
		Kilometers *uint32        `json:"kilometers"`
		Interval   *time.Duration `json:"interval"`
	}

	ServiceID string
)
