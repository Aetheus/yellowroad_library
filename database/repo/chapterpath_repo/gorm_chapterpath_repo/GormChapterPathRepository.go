package gorm_chapterpath_repo

import (
	"yellowroad_library/database/repo/chapterpath_repo"
	"yellowroad_library/utils/app_error"
	"github.com/jinzhu/gorm"
	"yellowroad_library/database/entities"
	"net/http"
	"errors"
)

type GormChapterPathRepository struct {
	db *gorm.DB
}
var _ chapterpath_repo.ChapterPathRepository = GormChapterPathRepository{}

func New(db *gorm.DB) GormChapterPathRepository{
	return GormChapterPathRepository{
		db : db,
	}
}

func (this GormChapterPathRepository) FindById(id int) (entities.ChapterPath, app_error.AppError) {
	var chapterPath entities.ChapterPath

	queryResult := this.db.
						Preload(entities.ASSOC_CHAPTERPATH_FROMCHAPTER).
						Preload(entities.ASSOC_CHAPTERPATH_TOCHAPTER).
						Where("id = ?", id).First(&chapterPath)

	if queryResult.Error != nil {
		var returnedErr app_error.AppError
		if queryResult.RecordNotFound() {
			returnedErr = app_error.Wrap(queryResult.Error).
				SetHttpCode(http.StatusNotFound).
				SetEndpointMessage("No path found")
		} else {
			returnedErr = app_error.Wrap(queryResult.Error)
		}
		return chapterPath, returnedErr
	}

	return chapterPath, nil
}

func (this GormChapterPathRepository) FindByChapterIds(fromChapterId int, toChapterId int) (entities.ChapterPath, app_error.AppError) {
	var chapterPath entities.ChapterPath
	var err app_error.AppError

	queryResult := this.db.
						Preload(entities.ASSOC_CHAPTERPATH_FROMCHAPTER).
						Preload(entities.ASSOC_CHAPTERPATH_TOCHAPTER).
						Where("from_chapter_id = ? AND to_chapter_id = ?", fromChapterId, toChapterId).
						First(&chapterPath)

	if queryResult.Error != nil {
		if queryResult.RecordNotFound() {
			err = app_error.Wrap(queryResult.Error).SetHttpCode(http.StatusNotFound).SetEndpointMessage("No path found")
		} else {
			err = app_error.Wrap(queryResult.Error)
		}
		return chapterPath, err
	}

	return chapterPath, nil
}

func (this GormChapterPathRepository) Update(chapter_path *entities.ChapterPath) app_error.AppError {
	queryResult := this.db.
					Set("gorm:save_associations", false).	//no magic! let the individual objects be saved on their own!
					Save(chapter_path)

	if queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}

func (this GormChapterPathRepository) Insert(chapter_path *entities.ChapterPath) app_error.AppError {
	queryResult := this.db.
		Set("gorm:save_associations", false).	//no magic! let the individual objects be saved on their own!
		Create(chapter_path)

	if queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}

func (this GormChapterPathRepository) Delete(chapter_path *entities.ChapterPath) app_error.AppError {
	if (chapter_path.ID == 0){
		err := errors.New("Invalid primary key value of 0 while attempting to delete")
		return app_error.Wrap(err)
	}

	if queryResult := this.db.Delete(chapter_path); queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}

