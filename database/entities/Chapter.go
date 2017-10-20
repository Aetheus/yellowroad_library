package entities

import (
	"time"
	"yellowroad_library/database"
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
	Book	  *Book `gorm:"ForeignKey:BookId" json:",omitempty"`

	CreatorId int
	Creator	  *User	`gorm:"ForeignKey:CreatorId" json:",omitempty"`

	//housekeeping attributes
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt database.NullTime
}

var ChapterAssociations = []string{
	"Book",
	"Creator",
}

//fields that we allow for updating
type Chapter_UpdateForm struct {
	Title *string
	Body *string
}
func (this Chapter_UpdateForm) Apply(chapter *Chapter){
	if(this.Title != nil) { chapter.Title = *this.Title }
	if(this.Body != nil) { chapter.Body = *this.Body }
}

//fields that we allow during creation
type Chapter_CreationForm struct {
	Title *string
	Body *string
	BookId *int

	//optional
	FromChapterPath *ChapterPath_CreationForm
}
func (this Chapter_CreationForm) Apply(chapter *Chapter){
	if(this.Title != nil) { chapter.Title = *this.Title }
	if(this.Body != nil) { chapter.Body = *this.Body }
	if(this.BookId != nil) { chapter.BookId = *this.BookId }
}