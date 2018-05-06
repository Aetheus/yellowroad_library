package entities

import (
	"time"
	"gopkg.in/guregu/null.v3"
	"encoding/json"
	"yellowroad_library/utils/app_error"
	"net/http"
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
	ChapterForm 	*Chapter_CreationForm		`json:"chapter"`
	ChapterPathForm *ChapterPath_CreationForm	`json:"chapter_path"`
}

func (form *Chapter_And_Path_CreationForm) UnmarshalJSON(data []byte) (error) {
	unmarshalTarget := struct {
		ChapterForm 	*Chapter_CreationForm		`json:"chapter"`
		ChapterPathForm *ChapterPath_CreationForm	`json:"chapter_path"`
	}{}

	err := json.Unmarshal(data, &unmarshalTarget)
	if (err != nil){
		return err
	} else if unmarshalTarget.ChapterForm == nil {
		return app_error.New(http.StatusBadRequest,
			"",
			"'chapter' is a required field!")
	} else if unmarshalTarget.ChapterPathForm == nil {
		return app_error.New(http.StatusBadRequest,
			"",
			"'chapter_path' is a required field!")
	} else {
		form.ChapterForm = unmarshalTarget.ChapterForm
		form.ChapterPathForm = unmarshalTarget.ChapterPathForm
	}

	return nil
}