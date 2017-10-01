package entities

import (
	"time"
	"yellowroad_library/database"
)

type ChapterPath struct {
	FromChapterId int
	//FromChapters []Chapter `gorm:"ForeignKey:ToChapterId"`

	ToChapterId int
	ToChapters []Chapter `gorm:"ForeignKey:ToChapterId"`

	//housekeeping attributes
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt database.NullTime
}

var ChapterPathAssociations = []string{
	"ToChapterId",
}

//for GORM
func (ChapterPath) TableName () string {
	return "chapter_paths"
}