package gorm_chapter_repo

import (
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
	"github.com/jinzhu/gorm"
	"net/http"

	"yellowroad_library/database/repo/chapter_repo"
	"errors"
)

type GormChapterRepository struct {
	dbConn *gorm.DB
}
var _ chapter_repo.ChapterRepository = GormChapterRepository{}

func New(dbConn *gorm.DB) GormChapterRepository{
	return GormChapterRepository{
		dbConn : dbConn,
	}
}


func (this GormChapterRepository) FindById(id int) (entities.Chapter, app_error.AppError) {
	var chapter entities.Chapter

	queryResult := this.dbConn.
					Preload(entities.ASSOC_CHAPTER_BOOK).
					Preload(entities.ASSOC_CHAPTER_CREATOR).
					Preload(entities.ASSOC_CHAPTER_PATHS_AWAY).
					Where("id = ?",id).
					First(&chapter)

	if queryResult.Error != nil {
		var returnedErr app_error.AppError
		if queryResult.RecordNotFound() {
			returnedErr = app_error.Wrap(queryResult.Error).
				SetHttpCode(http.StatusNotFound).
				SetEndpointMessage("No chapter found")
		} else {
			returnedErr = app_error.Wrap(queryResult.Error)
		}
		return chapter, returnedErr
	}

	return chapter, nil
}

func (this GormChapterRepository) FindWithinBook(chapter_id int, book_id int) (entities.Chapter, app_error.AppError){
	var chapter entities.Chapter

	queryResult := this.dbConn.
					Preload(entities.ASSOC_CHAPTER_BOOK).
					Preload(entities.ASSOC_CHAPTER_CREATOR).
					Preload(entities.ASSOC_CHAPTER_PATHS_AWAY).
					Where("id = ? AND book_id = ?",chapter_id, book_id).
					First(&chapter)

	if queryResult.Error != nil {
		var returnedErr app_error.AppError
		if queryResult.RecordNotFound() {
			returnedErr = app_error.Wrap(queryResult.Error).
				SetHttpCode(http.StatusNotFound).
				SetEndpointMessage("No chapter found")
		} else {
			returnedErr = app_error.Wrap(queryResult.Error)
		}
		return chapter, returnedErr
	}

	return chapter, nil
}

func (this GormChapterRepository) Update(chapter *entities.Chapter) app_error.AppError {
	queryResult := this.dbConn.
						Set("gorm:save_associations", false).	//no magic! let the individual objects be saved on their own!
						Save(chapter)

	if queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}

func (this GormChapterRepository) Insert(chapter *entities.Chapter) app_error.AppError {
	queryResult := this.dbConn.
						Set("gorm:save_associations", false).	//no magic! let the individual objects be saved on their own!
						Create(chapter)

	if queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}

func (this GormChapterRepository) Delete(chapter *entities.Chapter) app_error.AppError {
	if (chapter.ID == 0 ){
		err := errors.New("Invalid primary key value of 0 while attempting to delete")
		return app_error.Wrap(err)
	}

	if queryResult := this.dbConn.Delete(chapter); queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}
