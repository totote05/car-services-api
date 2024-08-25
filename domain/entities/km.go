package entities

import "time"

type (
	Km struct {
		ID    KmID      `json:"id"`
		Value uint64    `json:"value"`
		Date  time.Time `json:"date"`
	}

	KmID string
)
