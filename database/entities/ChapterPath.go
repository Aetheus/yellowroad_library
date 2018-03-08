package entities

import (
	"time"
	"yellowroad_library/database"
	"github.com/lib/pq"
)

type ChapterPath struct {
	FromChapterId int 				`json:"from_chapter_id"`
	FromChapter Chapter 			`json:"from_chapter,omitempty" gorm:"ForeignKey:FromChapterId"`

	ToChapterId int 				`json:"to_chapter_id"`
	ToChapter Chapter 				`json:"to_chapter,omitempty" gorm:"ForeignKey:ToChapterId"`

	Effects database.Jsonb			`json:"effects"`
	Requirements database.Jsonb		`json:"requirements"`
	Description string				`json:"description"`

	//housekeeping attributes
	ID        int					`json:"id"`
	CreatedAt time.Time				`json:"created_at"`
	UpdatedAt time.Time				`json:"updated_at"`
	DeletedAt pq.NullTime		`json:"deleted_at"`
}

const ASSOC_CHAPTERPATH_TOCHAPTER = "ToChapter"
const ASSOC_CHAPTERPATH_FROMCHAPTER = "FromChapter"

//for GORM
func (ChapterPath) TableName () string {
	return "chapter_paths"
}

type ChapterPath_CreationForm struct {
	FromChapterId *int 				`json:"from_chapter_id"`
	ToChapterId *int 				`json:"to_chapter_id"`
	Effects *database.Jsonb 		`json:"effects"`
	Requirements *database.Jsonb 	`json:"requirements"`
	Description *string 			`json:"description"`
}

func (this ChapterPath_CreationForm) Apply(path *ChapterPath){
	if (this.FromChapterId != nil) { path.FromChapterId = *this.FromChapterId }
	if (this.ToChapterId != nil) { path.ToChapterId = *this.ToChapterId }
	if (this.Effects != nil) { path.Effects = *this.Effects }
	if (this.Requirements != nil) { path.Requirements = *this.Requirements }
	if (this.Description != nil) { path.Description = *this.Description }
}

type ChapterPath_UpdateForm struct {
	Effects *database.Jsonb			`json:"effects"`
	Requirements *database.Jsonb	`json:"requirements"`
	Description *string				`json:"description"`
}
func (this ChapterPath_UpdateForm) Apply(path *ChapterPath){
	if (this.Effects != nil) { path.Effects = *this.Effects }
	if (this.Requirements != nil) { path.Requirements = *this.Requirements }
	if (this.Description != nil) { path.Description = *this.Description }
}