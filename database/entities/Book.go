package entities

import (
	"time"
	"yellowroad_library/utils"
)

type Book struct {
	Title          string
	Description    string
	FirstChapterId int
	CreatorId      int
	Permissions    string

	//housekeeping attributes
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt utils.NullTime
}
