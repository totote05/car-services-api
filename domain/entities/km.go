package entities

import (
	"errors"
	"time"
)

const (
	KmValidationSuccess KmValidation = iota
	KmValidationInvalid
	KmValidationNotFound
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

	KmID         string
	KmList       []Km
	KmValidation int
)

func (k *Km) Validate() error {
	if k != nil && k.Value == 0 {
		return ErrKmHasZeroValue
	}

	return nil
}

func (list KmList) ValidateConsistency(km Km) KmValidation {
	if len(list) == 0 && km.ID == "" {
		return KmValidationSuccess
	}

	var found bool
	for _, k := range list {
		isDifferentKm := k.ID != km.ID
		isSameDate := k.Date.Equal(km.Date)
		isSameValue := k.Value == km.Value
		isAfter := km.Date.After(k.Date) && km.Value < k.Value
		isBefore := km.Date.Before(k.Date) && km.Value > k.Value

		if isDifferentKm && (isSameDate || isSameValue || isAfter || isBefore) {
			return KmValidationInvalid
		}
		if k.ID == km.ID {
			found = true
		}
	}

	if !found && km.ID != "" {
		return KmValidationNotFound
	}

	return KmValidationSuccess
}
