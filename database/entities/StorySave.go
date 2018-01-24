package entities

import (
	"yellowroad_library/database"
	"time"
)

type StorySave struct {
	Token 			string				`json:"token" gorm:"primary_key"`
	Save 			database.Jsonb		`json:"save"`

	CreatedBy 		int					`json:"created_by"`
	Creator			User				`json:"creator,omitempty" gorm:"ForeignKey:CreatedBy"`

	BookId    		int					`json:"book_id"`
	Book			Book				`json:"book,omitempty" gorm:"ForeignKey:BookId"`

	ChapterId 		int					`json:"chapter_id"`
	Chapter			Chapter 			`json:"chapter,omitempty" gorm:"ForeignKey:ChapterId"`

	//housekeeping attributes
	CreatedAt 		time.Time			`json:"created_at"`
}

//const for association requests
const STORYSAVE_CREATOR = "CreatedBy"
const STORYSAVE_BOOK = "Book"
const STORYSAVE_CHAPTER = "Chapter"

func (this StorySave) TableName() string{
	return "story_saves"
}

type StorySave_CreationForm struct {
	Token 		*string
	Save		*database.Jsonb
	BookId		*int
	ChapterId	*int
}

func (this StorySave_CreationForm) Apply(save *StorySave){
	if(this.Token != nil) { save.Token = *this.Token }
	if(this.Save != nil) { save.Save  = *this.Save }
	if(this.BookId != nil) { save.BookId = *this.BookId }
	if(this.ChapterId != nil) {save.ChapterId = *this.ChapterId}
}
