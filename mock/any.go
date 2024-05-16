package mock

import (
	"database/sql/driver"
	"time"
)

type Any struct{}

func (a Any) Match(v driver.Value) bool {
	return true
}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type AnyTimeAfter struct {
	Value time.Time
}

func (a AnyTimeAfter) Match(v driver.Value) bool {
	actual, ok := v.(time.Time)
	if !ok {
		return false
	}
	return actual.After(a.Value)
}
