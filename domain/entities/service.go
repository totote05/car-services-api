package entities

import (
	"errors"
	"time"
)

type (
	Service struct {
		ID          ServiceID          `json:"id"`
		Name        string             `json:"name"`
		Rercurrence *ServiceRecurrence `json:"rercurrence"`
	}

	ServiceRecurrence struct {
		Kilometers *uint32        `json:"kilometers"`
		Interval   *time.Duration `json:"interval"`
	}

	ServiceID string
)

func (s Service) Validate() error {
	if s.Name == "" {
		return errors.New("empty name")
	}

	if err := s.Rercurrence.Validate(); err != nil {
		return err
	}

	return nil
}

func (r *ServiceRecurrence) Validate() error {
	if r != nil && r.Kilometers == nil && r.Interval == nil {
		return errors.New("shoul have at least one recurrence")
	}

	return nil
}
