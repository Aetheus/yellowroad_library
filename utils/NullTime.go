package utils

import "time"

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}
