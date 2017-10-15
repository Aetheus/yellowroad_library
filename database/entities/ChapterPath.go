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