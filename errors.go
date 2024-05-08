package goapipattern

import (
	"errors"
)

var (
	ErrNotFound       = errors.New("not found")
	ErrParentNotFound = errors.New("parent not found")
)
