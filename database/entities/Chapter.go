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
	Book	  *Book `gorm:"ForeignKey:BookId"`

	CreatorId int
	Creator	  *User	`gorm:"ForeignKey:CreatorId"`

	//housekeeping attributes
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt utils.NullTime
}

var ChapterAssociations = []string{
	"Book",
	"Creator",
}

//fields that we allow to edit in our handlers (e.g: for the "update" routes of CRUD)
type ChapterForm struct {
	Title *string
	Body *string
	BookId *int
}
func (this ChapterForm) Apply(chapter *Chapter){
	if(this.Title != nil) { chapter.Title = *this.Title }
	if(this.Body != nil) { chapter.Body = *this.Body }
	if(this.BookId != nil) { chapter.BookId = *this.BookId }
}