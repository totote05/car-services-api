package entities

import (
	"errors"
	"time"
)

var (
	ErrKmHasZeroValue = errors.New("km has zero value")
)

type (
	Km struct {
		ID    KmID      `json:"id"`
		Value uint64    `json:"value"`
		Date  time.Time `json:"date"`
	}

	KmID string
)

func (k *Km) Validate() error {
	if k != nil && k.Value == 0 {
		return ErrKmHasZeroValue
	}

	return nil
}
