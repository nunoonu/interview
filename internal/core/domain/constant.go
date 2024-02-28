package domain

import "errors"

var (
	ErrorNotFound = errors.New("not found")
	ErrorConflict = errors.New("conflict")
)
