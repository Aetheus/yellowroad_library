package entities

import (
	"time"
	"yellowroad_library/database"
	"github.com/lib/pq"
	"net/http"
	"github.com/asaskevich/govalidator"
	"yellowroad_library/utils/app_error"
)

type ChapterPath struct {
	FromChapterId int 				`json:"from_chapter_id"`
	FromChapter Chapter 			`json:"from_chapter,omitempty" gorm:"foreignkey:FromChapterId"`

	ToChapterId int 				`json:"to_chapter_id"`
	ToChapter Chapter 				`json:"to_chapter,omitempty" gorm:"foreignkey:ToChapterId"`

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
	FromChapterId *int 				`json:"from_chapter_id" valid:"required~FromChapterId is a required field"`
	ToChapterId *int 				`json:"to_chapter_id"   valid:"required~ToChapterId is a required field"`
	Effects *database.Jsonb 		`json:"effects"         valid:"required~Effects is a required field"`
	Requirements *database.Jsonb 	`json:"requirements"    valid:"required~Requirements is a required field"`
	Description string 				`json:"description"     valid:"required~Description is a required field,stringlength(1|255)~Description should be between 1 to 255 characters"`
}

func (this ChapterPath_CreationForm) Validate() (app_error.AppError){
	if _,err := govalidator.ValidateStruct(this); err != nil {
		return app_error.New(http.StatusUnprocessableEntity,"",err.Error())
	}

	return nil
}

func (this ChapterPath_CreationForm) Apply(path *ChapterPath) (app_error.AppError){
	if err := this.Validate(); err != nil{
		return err
	}

	path.FromChapterId = *this.FromChapterId
	path.ToChapterId = *this.ToChapterId
	path.Effects = *this.Effects
	path.Requirements = *this.Requirements
	path.Description = this.Description

	return nil
}

type ChapterPath_UpdateForm struct {
	Effects 		*database.Jsonb		`json:"effects"      valid:"optional"`
	Requirements 	*database.Jsonb		`json:"requirements" valid:"optional"`
	Description 	*string				`json:"description"  valid:"optional,stringlength(1|255)~Description should be between 1 to 255 characters"`
}
func (this ChapterPath_UpdateForm) Apply(path *ChapterPath){
	if (this.Effects != nil) { path.Effects = *this.Effects }
	if (this.Requirements != nil) { path.Requirements = *this.Requirements }
	if (this.Description != nil) { path.Description = *this.Description }
}