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
	Title     string			`json:"title"`
	Body      string			`json:"body"`

	BookId    int				`json:"book_id"`
	Book	  *Book 			`json:"book,omitempty" gorm:"ForeignKey:BookId"`

	CreatorId int				`json:"creator_id"`
	Creator	  *User				`json:"creator,omitempty" gorm:"ForeignKey:CreatorId"`

	//housekeeping attributes
	ID        int				`json:"id"`
	CreatedAt time.Time			`json:"created_at"`
	UpdatedAt time.Time			`json:"updated_at"`
	DeletedAt database.NullTime	`json:"deleted_at"`
}

var ChapterAssociations = []string{
	"Book",
	"Creator",
}

//fields that we allow for updating
type Chapter_UpdateForm struct {
	Title *string	`json:"title"`
	Body *string	`json:"body"`
}
func (this Chapter_UpdateForm) Apply(chapter *Chapter){
	if(this.Title != nil) { chapter.Title = *this.Title }
	if(this.Body != nil) { chapter.Body = *this.Body }
}

//fields that we allow during creation
type Chapter_CreationForm struct {
	Title *string	`json:"title"`
	Body *string	`json:"body"`
	BookId *int		`json:"book_id"`
}
func (this Chapter_CreationForm) Apply(chapter *Chapter){
	if(this.Title != nil) { chapter.Title = *this.Title }
	if(this.Body != nil) { chapter.Body = *this.Body }
	if(this.BookId != nil) { chapter.BookId = *this.BookId }
}

type Chapter_And_Path_CreationForm struct {
	ChapterForm Chapter_CreationForm
	ChapterPathForm ChapterPath_CreationForm
}