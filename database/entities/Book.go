package entities

import (
	"time"
	"yellowroad_library/database"
)

type Book struct {
	Title          string
	Description    string
	Permissions    string

	FirstChapterId database.NullInt  `sql:"DEFAULT:null"` //when first creating a book, you won't have a first chapter
	FirstChapter   *Chapter `gorm:"ForeignKey:FirstChapterId" json:",omitempty"`

	CreatorId      int
	Creator		   *User	`gorm:"ForeignKey:CreatorId" json:",omitempty"`

	//housekeeping attributes
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt database.NullTime
}

var BookAssociations = []string{
	"FirstChapter",
	"Creator",
}


type BookCreationForm struct {
	Title *string
	Description *string
	//FirstChapterId *int
}
func (this BookCreationForm) Apply(book *Book){
	if (this.Title != nil) { book.Title = *this.Title }
	if (this.Description != nil ) {book.Description = *this.Description}
	//if (this.FirstChapterId != nil ) {book.FirstChapterId = database.NullInt{Int : *this.FirstChapterId}}
}

type BookUpdateForm struct {
	Title *string
	Description *string
}
func (this BookUpdateForm) Apply(book *Book){
	if (this.Title != nil) { book.Title = *this.Title }
	if (this.Description != nil ) {book.Description = *this.Description}
}