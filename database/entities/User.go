package entities

import (
	"time"
	"yellowroad_library/utils"
)

type User struct {
	Username string
	Password string `json:"-"` //this excludes it from being accidentally serialized into a JSON and shown to users
	Email    string

	//housekeeping attributes
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt utils.NullTime
}
