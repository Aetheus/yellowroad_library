package entities

import (
	"time"
	"yellowroad_library/database"
	"github.com/lib/pq"
)

type Book struct {
	Title          string			`json:"title"`
	Description    string			`json:"description"`
	Permissions    string			`json:"permissions"`

	FirstChapterId database.NullInt `json:"first_chapter_id" sql:"DEFAULT:null"` //when first creating a book, you won't have a first chapter
	FirstChapter   *Chapter 		`json:"first_chapter,omitempty" gorm:"ForeignKey:FirstChapterId"`

	CreatorId      int				`json:"creator_id"`
	Creator		   *User			`json:"creator,omitempty" gorm:"ForeignKey:CreatorId"`

	//housekeeping attributes
	ID        int					`json:"id"`
	CreatedAt time.Time				`json:"created_at"`
	UpdatedAt time.Time				`json:"updated_at"`
	DeletedAt pq.NullTime		`json:"deleted_at"`
}

var BookAssociations = []string{
	"FirstChapter",
	"Creator",
}


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