package entities

import (
	"time"
	"yellowroad_library/utils"
)

/*

   title text NOT NULL,
   body text NOT NULL,
   book_id integer NOT NULL,

   creator_id integer
*/

type Chapter struct {
	Title     string
	Body      string
	BookId    int
	CreatorId int

	//housekeeping attributes
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt utils.NullTime
}
