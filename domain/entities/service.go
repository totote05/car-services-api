package entities

import (
	"errors"
	"time"
)

var (
	ErrServiceHasEmptyName       = errors.New("empty name")
	ErrServiceHasEmptyRecurrence = errors.New("should have at least one recurrence")
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

func (s *Service) Validate() error {
	if s == nil {
		return nil
	}

	if s.Name == "" {
		return ErrServiceHasEmptyName
	}

	if err := s.Rercurrence.Validate(); err != nil {
		return err
	}

	return nil
}

func (r *ServiceRecurrence) Validate() error {
	if r != nil && r.Kilometers == nil && r.Interval == nil {
		return ErrServiceHasEmptyRecurrence
	}

	return nil
}
