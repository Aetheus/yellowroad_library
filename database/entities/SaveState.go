package entities

import (
	"yellowroad_library/database"
	"time"
)

type SaveState struct {
	Id 				int					`json:"id" gorm:"primary_key"`
	State 			database.Jsonb		`json:"state"` /* in the shape of: {history: [{hp:9}, {hp:5},...], cursor:1} */

	CreatedBy 		int					`json:"created_by"`
	Creator			User				`json:"creator,omitempty" gorm:"foreignkey:CreatedBy"`

	BookId    		int					`json:"book_id"`
	Book			Book				`json:"book,omitempty" gorm:"foreignkey:BookId"`

	//housekeeping attributes
	CreatedAt 		time.Time			`json:"created_at"`
	UpdatedAt 		time.Time			`json:"updated_at"`
}

//const for association requests
const STORYSAVE_CREATOR = "CreatedBy"
const STORYSAVE_BOOK = "Book"

func (this SaveState) TableName() string{
	return "save_states"
}

type StorySave_CreationForm struct {
	State		*database.Jsonb
	BookId		*int
	CreatedBy	*int
}

func (this StorySave_CreationForm) Apply(save *SaveState){
	if(this.State != nil) { save.State  = *this.State }
	if(this.BookId != nil) { save.BookId = *this.BookId }
	if(this.CreatedBy != nil) { save.CreatedBy = *this.CreatedBy }
}
