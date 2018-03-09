package entities

import (
	"time"
	"gopkg.in/guregu/null.v3"
)

type Book struct {
	Title          string			`json:"title"`
	Description    string			`json:"description"`
	Permissions    string			`json:"permissions"`

	FirstChapterId null.Int 		`json:"first_chapter_id" sql:"DEFAULT:null"` //when first creating a book, you won't have a first chapter
	FirstChapter   *Chapter 		`json:"first_chapter,omitempty" gorm:"foreignkey:FirstChapterId"`

	//Chapters 		[]Chapter		`gorm:"ForeignKey:BookId;AssociationForeignKey:ID"`
	ChapterCount 	int				`json:"chapter_count" sql:"-"`

	CreatorId      int				`json:"creator_id"`
	Creator		   *User			`json:"creator,omitempty" gorm:"foreignkey:CreatorId"`

	//housekeeping attributes
	ID        int					`json:"id"`
	CreatedAt time.Time				`json:"created_at"`
	UpdatedAt time.Time				`json:"updated_at"`
	DeletedAt null.Time				`json:"deleted_at"`
}

//Constants for Gorm Association queries
const ASSOC_BOOK_FIRST_CHAPTER = "FirstChapter"
const ASSOC_BOOK_CREATOR = "Creator"
const ASSOC_BOOK_ASSOC_CHAPTERS = "Chapters"


type Book_CreationForm struct {
	Title *string			`json:"title"`
	Description *string		`json:"description"`
	//FirstChapterId *int
}
func (this Book_CreationForm) Apply(book *Book){
	if (this.Title != nil) { book.Title = *this.Title }
	if (this.Description != nil ) {book.Description = *this.Description}
	//if (this.FirstChapterId != nil ) {book.FirstChapterId = database.NullInt{Int : *this.FirstChapterId}}
}

type Book_UpdateForm struct {
	Title *string			`json:"title"`
	Description *string		`json:"description"`
}
func (this Book_UpdateForm) Apply(book *Book){
	if (this.Title != nil) { book.Title = *this.Title }
	if (this.Description != nil ) {book.Description = *this.Description}
}