package entities

import (
	"time"
	"gopkg.in/guregu/null.v3"
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
	Book	  *Book 			`json:"book,omitempty" gorm:"foreignkey:BookId"`

	CreatorId int				`json:"creator_id"`
	Creator	  *User				`json:"creator,omitempty" gorm:"foreignkey:CreatorId"`

	PathsAway []ChapterPath		`json:"paths_away" gorm:"foreignkey:FromChapterId;association_foreignkey:ID"`

	//housekeeping attributes
	ID        int				`json:"id"`
	CreatedAt time.Time			`json:"created_at"`
	UpdatedAt time.Time			`json:"updated_at"`
	DeletedAt null.Time			`json:"deleted_at"`
}

const ASSOC_CHAPTER_BOOK = "Book"
const ASSOC_CHAPTER_CREATOR = "Creator"
const ASSOC_CHAPTER_PATHS_AWAY = "PathsAway"


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