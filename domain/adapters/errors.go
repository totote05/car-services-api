package adapters

import "errors"

var (
	ErrNotFound   = errors.New("not found")
	ErrPersisting = errors.New("error saving")
	ErrGetting    = errors.New("error getting")
)
