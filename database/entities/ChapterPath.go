package entities

import (
	"time"
	"yellowroad_library/database"
)

type ChapterPath struct {
	FromChapterId int
	FromChapter Chapter `gorm:"ForeignKey:FromChapterId"`

	ToChapterId int
	ToChapter Chapter `gorm:"ForeignKey:ToChapterId"`

	Effects database.Jsonb
	Requirements database.Jsonb

	//housekeeping attributes
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt database.NullTime
}

var ChapterPathAssociations = []string{
	"ToChapter", "FromChapter",
}

//for GORM
func (ChapterPath) TableName () string {
	return "chapter_paths"
}

type ChapterPathForm struct {
	FromChapterId *int
	ToChapterId *int
	Effects *database.Jsonb
	Requirements *database.Jsonb
}

func (this ChapterPathForm) Apply(path *ChapterPath){
	if (this.FromChapterId != nil) { path.FromChapterId = *this.FromChapterId }
	if (this.ToChapterId != nil) { path.ToChapterId = *this.ToChapterId }
	if (this.Effects != nil) { path.Effects = *this.Effects }
	if (this.Requirements != nil) { path.Requirements = *this.Requirements }
}