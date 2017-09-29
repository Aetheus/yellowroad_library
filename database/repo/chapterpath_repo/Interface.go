package chapterpath_repo

import (
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
)

type ChapterPathRepository interface {
	FindById(int) (entities.ChapterPath, app_error.AppError)
	Update(*entities.ChapterPath) app_error.AppError
	Insert(*entities.ChapterPath) app_error.AppError
	Delete(*entities.ChapterPath) app_error.AppError
}

